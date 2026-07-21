package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	outputDir      = flag.String("o", ".", "Output directory")
	convertMP3     = flag.Bool("mp3", false, "Convert to MP3 (requires ffmpeg)")
	concurrency    = flag.Int("c", 3, "Number of concurrent downloads")
	timeout        = flag.Duration("t", 10*time.Minute, "Per-download timeout")
	retries        = flag.Int("r", 3, "Max retries per download (exponential backoff)")
	maxBackoff    = flag.Duration("max-backoff", 60*time.Second, "Cap on retry backoff duration")
	verbose        = flag.Bool("v", false, "Verbose logging (debug level)")
	setupDeps     = flag.Bool("setup", true, "Auto-install missing dependencies when possible")
	allowInstalls = flag.Bool("install", true, "Allow installer actions when -setup is enabled")
)

type ytDlpRunner struct {
	execPath string
	desc     string
}

func main() {
	flag.Parse()

	if runtime.GOOS != "windows" {
		fmt.Fprintln(os.Stderr, "This version is Windows-only.")
		os.Exit(1)
	}

	logLevel := slog.LevelInfo
	if *verbose {
		logLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(logger)

	urls := normalizeURLs(flag.Args())
	if len(urls) == 0 {
		urls = normalizeURLs(promptURLs())
	}
	if len(urls) == 0 {
		slog.Error("No URLs provided")
		os.Exit(1)
	}

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		slog.Error("Cannot create output directory", "error", err)
		os.Exit(1)
	}

	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	runner, nodeSpec, ffmpegDir, err := ensureEnvironment(rootCtx, *setupDeps, *allowInstalls, *convertMP3)
	if err != nil {
		slog.Error("Dependency check failed", "error", err)
		os.Exit(1)
	}

	queue := make(chan string, len(urls))
	for _, u := range urls {
		queue <- u
	}
	close(queue)

	var (
		wg      sync.WaitGroup
		errList []error
		mu      sync.Mutex
	)

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range queue {
				select {
				case <-rootCtx.Done():
					mu.Lock()
					errList = append(errList, fmt.Errorf("%s: cancelled", url))
					mu.Unlock()
					continue
				default:
				}

				if err := downloadWithRetry(rootCtx, runner, nodeSpec, ffmpegDir, url, *outputDir, *convertMP3, *timeout, *retries, *maxBackoff); err != nil {
					mu.Lock()
					errList = append(errList, fmt.Errorf("%s: %w", url, err))
					mu.Unlock()
				}
			}
		}()
	}

	go func() {
		<-rootCtx.Done()
		slog.Warn("Interrupt received, finishing in-flight downloads; press Ctrl+C again to force quit")
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		slog.Error("Second interrupt received, exiting immediately")
		os.Exit(1)
	}()

	wg.Wait()

	if len(errList) > 0 {
		slog.Error("Some downloads failed", "error", errors.Join(errList...))
		os.Exit(1)
	}

	slog.Info("All downloads finished successfully")
}

func ensureEnvironment(ctx context.Context, setup, allowInstalls, wantMP3 bool) (ytDlpRunner, string, string, error) {
	runner, err := ensureYtDlp(ctx, setup, allowInstalls)
	if err != nil {
		return ytDlpRunner{}, "", "", err
	}

	nodeSpec, err := ensureNodeRuntime(ctx, setup, allowInstalls)
	if err != nil {
		return ytDlpRunner{}, "", "", err
	}

	ffmpegDir := ""
	if wantMP3 {
		ffmpegDir, err = ensureFFmpeg(ctx, setup, allowInstalls)
		if err != nil {
			return ytDlpRunner{}, "", "", err
		}
	}

	return runner, nodeSpec, ffmpegDir, nil
}

func ensureYtDlp(ctx context.Context, setup, allowInstalls bool) (ytDlpRunner, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		cacheDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "yt-dlp-cache")
	}
	localDir := filepath.Join(cacheDir, "yt-dlp")
	localExe := filepath.Join(localDir, "yt-dlp.exe")

	if isUsableYtDlp(ctx, localExe) {
		return ytDlpRunner{execPath: localExe, desc: "official cached yt-dlp.exe"}, nil
	}

	if setup && allowInstalls {
		if err := downloadOfficialYtDlp(ctx, localExe); err == nil && isUsableYtDlp(ctx, localExe) {
			return ytDlpRunner{execPath: localExe, desc: "official downloaded yt-dlp.exe"}, nil
		}
		slog.Warn("Official yt-dlp download failed, trying Windows package managers")
		if err := installYtDlpViaWindowsManagers(ctx); err == nil {
			if p, ok := findExeAfterInstall("yt-dlp.exe", []string{
				filepath.Join(os.Getenv("LOCALAPPDATA"), "Microsoft", "WinGet", "Links", "yt-dlp.exe"),
				filepath.Join(os.Getenv("ProgramFiles"), "yt-dlp", "yt-dlp.exe"),
				filepath.Join(os.Getenv("ProgramFiles(x86)"), "yt-dlp", "yt-dlp.exe"),
				filepath.Join(os.Getenv("USERPROFILE"), "scoop", "shims", "yt-dlp.exe"),
				filepath.Join(os.Getenv("ProgramData"), "chocolatey", "bin", "yt-dlp.exe"),
			}); ok && isUsableYtDlp(ctx, p) {
				return ytDlpRunner{execPath: p, desc: "package-managed yt-dlp.exe"}, nil
			}
		}
	}

	if p, ok := findExeAfterInstall("yt-dlp.exe", nil); ok && isUsableYtDlp(ctx, p) {
		return ytDlpRunner{execPath: p, desc: "yt-dlp.exe from PATH"}, nil
	}

	return ytDlpRunner{}, fmt.Errorf("yt-dlp not found; enable -setup or install it manually")
}

func ensureNodeRuntime(ctx context.Context, setup, allowInstalls bool) (string, error) {
	if p, ver, ok := findNodeExe(); ok && compareVersions(ver, "22.0.0") >= 0 {
		return "node:" + p, nil
	}

	if setup && allowInstalls {
		slog.Warn("Node.js missing or too old, attempting installation")
		if err := installNodeViaWindowsManagers(ctx); err == nil {
			if p, ver, ok := findNodeExe(); ok && compareVersions(ver, "22.0.0") >= 0 {
				return "node:" + p, nil
			}
		}
	}

	if p, ver, ok := findNodeExe(); ok {
		return "", fmt.Errorf("supported Node.js not available: found %s at %s, need >= 22.0.0", ver, p)
	}
	return "", fmt.Errorf("Node.js not found; install Node 22+ or run with -setup")
}

func ensureFFmpeg(ctx context.Context, setup, allowInstalls bool) (string, error) {
	if p, ok := findFFmpegExe(); ok {
		return filepath.Dir(p), nil
	}

	if setup && allowInstalls {
		slog.Warn("ffmpeg missing, attempting installation")
		if err := installFFmpegViaWindowsManagers(ctx); err == nil {
			if p, ok := findFFmpegExe(); ok {
				return filepath.Dir(p), nil
			}
		}
	}

	if p, ok := findFFmpegExe(); ok {
		return filepath.Dir(p), nil
	}
	return "", fmt.Errorf("ffmpeg is required for MP3 conversion")
}

func normalizeURLs(urls []string) []string {
	seen := make(map[string]struct{}, len(urls))
	out := make([]string, 0, len(urls))
	for _, u := range urls {
		u = strings.TrimSpace(u)
		if u == "" {
			continue
		}
		if _, ok := seen[u]; ok {
			slog.Warn("Skipping duplicate URL", "url", u)
			continue
		}
		seen[u] = struct{}{}
		out = append(out, u)
	}
	return out
}

func promptURLs() []string {
	fmt.Println("Enter YouTube links (one per line, empty line to finish):")
	scanner := bufio.NewScanner(os.Stdin)
	var urls []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		urls = append(urls, line)
	}
	return urls
}

func downloadWithRetry(
	ctx context.Context,
	runner ytDlpRunner,
	nodeSpec, ffmpegDir, url, outDir string,
	toMP3 bool,
	timeout time.Duration,
	maxRetries int,
	maxBackoff time.Duration,
) error {
	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		err := downloadOne(ctx, runner, nodeSpec, ffmpegDir, url, outDir, toMP3, timeout)
		if err == nil {
			return nil
		}
		lastErr = err
		if !isTransient(err) {
			slog.Warn("Non-transient error, giving up", "url", url, "error", err)
			return err
		}
		backoff := time.Duration(1<<uint(attempt)) * time.Second
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
		slog.Warn("Retrying download", "attempt", attempt+1, "maxRetries", maxRetries, "backoff", backoff, "url", url, "error", err)
		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return fmt.Errorf("all %d retries failed: %w", maxRetries, lastErr)
}

type transientError struct {
	err    error
	stderr string
}

func (e *transientError) Error() string {
	if e.stderr == "" {
		return e.err.Error()
	}
	return fmt.Sprintf("%v: %s", e.err, e.stderr)
}

func (e *transientError) Unwrap() error { return e.err }

func downloadOne(parent context.Context, runner ytDlpRunner, nodeSpec, ffmpegDir, url, outDir string, toMP3 bool, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()

	audioFormat := "m4a"
	if toMP3 {
		audioFormat = "mp3"
	}

	args := []string{
		"-f", "bestaudio",
		"--extract-audio",
		"--audio-format", audioFormat,
		"--no-playlist",
		"--js-runtimes", nodeSpec,
		"--no-overwrites",
		"--print", "after_move:filepath",
		"-o", filepath.Join(outDir, "%(title)s.%(ext)s"),
		url,
	}

	cmd := exec.CommandContext(ctx, runner.execPath, args...)
	cmd.Env = addPathDir(os.Environ(), ffmpegDir)

	var stderrBuf bytes.Buffer
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("pipe creation: %w", err)
	}

	slog.Info("Starting download", "url", url, "runner", runner.desc)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start yt-dlp: %w", err)
	}

	outBytes, readErr := io.ReadAll(stdout)
	waitErr := cmd.Wait()

	if readErr != nil {
		return &transientError{err: fmt.Errorf("reading yt-dlp output: %w", readErr), stderr: stderrBuf.String()}
	}

	if waitErr != nil {
		if ctx.Err() != nil {
			return &transientError{err: fmt.Errorf("yt-dlp error: %w", ctx.Err()), stderr: stderrBuf.String()}
		}
		return &transientError{err: fmt.Errorf("yt-dlp error: %w", waitErr), stderr: stderrBuf.String()}
	}

	finalFile := strings.TrimSpace(string(outBytes))
	if finalFile != "" {
		slog.Info("Saved", "file", finalFile)
	} else if strings.Contains(strings.ToLower(stderrBuf.String()), "already") {
		slog.Info("Skipped (already exists)", "url", url)
	} else {
		slog.Info("Download complete (filename not reported)")
	}
	return nil
}

func addPathDir(env []string, extraDir string) []string {
	if extraDir == "" {
		return env
	}
	currentPath := os.Getenv("PATH")
	newPath := extraDir + ";" + currentPath
	out := make([]string, 0, len(env)+1)
	found := false
	for _, kv := range env {
		if strings.HasPrefix(strings.ToUpper(kv), "PATH=") {
			out = append(out, "PATH="+newPath)
			found = true
		} else {
			out = append(out, kv)
		}
	}
	if !found {
		out = append(out, "PATH="+newPath)
	}
	return out
}

func isTransient(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	msg := strings.ToLower(err.Error())
	for _, s := range []string{
		"http error 429",
		"too many requests",
		"http error 503",
		"http error 502",
		"connection reset",
		"temporary failure",
		"unable to download webpage",
		"signature solving failed",
		"challenge solving failed",
	} {
		if strings.Contains(msg, s) {
			return true
		}
	}
	return false
}

func isUsableYtDlp(ctx context.Context, path string) bool {
	if path == "" {
		return false
	}
	cmd := exec.CommandContext(ctx, path, "--version")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	return cmd.Run() == nil
}

func downloadOfficialYtDlp(ctx context.Context, dest string) error {
	const url = "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp.exe"

	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}

	tmp := dest + ".tmp"
	_ = os.Remove(tmp)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("download failed: %s", resp.Status)
	}

	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(io.MultiWriter(f, h), resp.Body); err != nil {
		_ = os.Remove(tmp)
		return err
	}

	if err := f.Close(); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	if err := os.Rename(tmp, dest); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}

func findExeAfterInstall(name string, extraCandidates []string) (string, bool) {
	if p, err := exec.LookPath(name); err == nil {
		return p, true
	}
	for _, p := range extraCandidates {
		if fileExists(p) {
			return p, true
		}
	}
	return "", false
}

func findNodeExe() (string, string, bool) {
	candidates := []string{
		filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "nodejs", "node.exe"),
		filepath.Join(os.Getenv("ProgramFiles"), "nodejs", "node.exe"),
		filepath.Join(os.Getenv("ProgramFiles(x86)"), "nodejs", "node.exe"),
		filepath.Join(os.Getenv("ProgramData"), "chocolatey", "bin", "node.exe"),
		filepath.Join(os.Getenv("USERPROFILE"), "scoop", "shims", "node.exe"),
	}
	if p, err := exec.LookPath("node.exe"); err == nil {
		if ver, ok := nodeVersion(p); ok {
			return p, ver, true
		}
	}
	for _, p := range candidates {
		if ver, ok := nodeVersion(p); ok {
			return p, ver, true
		}
	}
	return "", "", false
}

func findFFmpegExe() (string, bool) {
	candidates := []string{
		filepath.Join(os.Getenv("ProgramFiles"), "ffmpeg", "bin", "ffmpeg.exe"),
		filepath.Join(os.Getenv("ProgramFiles"), "FFmpeg", "bin", "ffmpeg.exe"),
		filepath.Join(os.Getenv("ProgramData"), "chocolatey", "bin", "ffmpeg.exe"),
		filepath.Join(os.Getenv("USERPROFILE"), "scoop", "shims", "ffmpeg.exe"),
	}
	if p, err := exec.LookPath("ffmpeg.exe"); err == nil {
		return p, true
	}
	for _, p := range candidates {
		if fileExists(p) {
			return p, true
		}
	}
	return "", false
}

func nodeVersion(path string) (string, bool) {
	if path == "" {
		return "", false
	}
	out, err := exec.Command(path, "--version").Output()
	if err != nil {
		return "", false
	}
	v := strings.TrimSpace(string(out))
	v = strings.TrimPrefix(v, "v")
	v = strings.TrimSpace(v)
	if v == "" {
		return "", false
	}
	return v, true
}

func compareVersions(a, b string) int {
	pa := parseVersion(a)
	pb := parseVersion(b)

	n := len(pa)
	if len(pb) > n {
		n = len(pb)
	}

	for i := 0; i < n; i++ {
		av := 0
		bv := 0
		if i < len(pa) {
			av = pa[i]
		}
		if i < len(pb) {
			bv = pb[i]
		}
		switch {
		case av < bv:
			return -1
		case av > bv:
			return 1
		}
	}
	return 0
}

func parseVersion(v string) []int {
	v = strings.TrimSpace(v)
	if v == "" {
		return []int{0}
	}
	fields := strings.FieldsFunc(v, func(r rune) bool {
		return r < '0' || r > '9'
	})

	out := make([]int, 0, len(fields))
	for _, f := range fields {
		if f == "" {
			continue
		}
		n, err := strconv.Atoi(f)
		if err == nil {
			out = append(out, n)
		}
	}
	if len(out) == 0 {
		return []int{0}
	}
	return out
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func installYtDlpViaWindowsManagers(ctx context.Context) error {
	if hasCommand("scoop") {
		return runInstaller(ctx, "scoop", "install", "yt-dlp")
	}
	if hasCommand("winget") {
		return runInstaller(ctx, "winget", "install", "yt-dlp", "--accept-source-agreements", "--accept-package-agreements", "--silent")
	}
	if hasCommand("choco") {
		return runInstaller(ctx, "choco", "install", "yt-dlp", "-y")
	}
	return fmt.Errorf("no supported Windows package manager found for yt-dlp")
}

func installNodeViaWindowsManagers(ctx context.Context) error {
	if hasCommand("scoop") {
		return runInstaller(ctx, "scoop", "install", "nodejs-lts")
	}
	if hasCommand("winget") {
		return runInstaller(ctx, "winget", "install", "OpenJS.NodeJS.LTS", "--accept-source-agreements", "--accept-package-agreements", "--silent")
	}
	if hasCommand("choco") {
		return runInstaller(ctx, "choco", "install", "nodejs-lts", "-y")
	}
	return fmt.Errorf("no supported Windows package manager found for Node.js")
}

func installFFmpegViaWindowsManagers(ctx context.Context) error {
	if hasCommand("scoop") {
		return runInstaller(ctx, "scoop", "install", "ffmpeg")
	}
	if hasCommand("winget") {
		return runInstaller(ctx, "winget", "install", "Gyan.FFmpeg", "--accept-source-agreements", "--accept-package-agreements", "--silent")
	}
	if hasCommand("choco") {
		return runInstaller(ctx, "choco", "install", "ffmpeg", "-y")
	}
	return fmt.Errorf("no supported Windows package manager found for ffmpeg")
}

func hasCommand(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func runInstaller(ctx context.Context, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}
	return nil
}
