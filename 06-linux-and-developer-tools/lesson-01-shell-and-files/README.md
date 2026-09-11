# Lesson 1: Shell, Files, and Permissions

## Objective

Reason precisely about how Linux represents files, permissions, and pipelines — not just run commands by memorized incantation, but explain what the shell and kernel are actually doing at each step.

## Prerequisites

None. This is the start of the Linux/tools track.

## Learn

**Everything is a file descriptor.** A process's open files (regular files, sockets, pipes, devices) are represented by small integers — file descriptors — that index into a per-process table the kernel maintains. By convention, fd 0 is stdin, fd 1 is stdout, fd 2 is stderr; every process starts with these three already open, inherited from its parent. `/proc/<pid>/fd/` exposes this table directly as a directory of symlinks, letting you inspect exactly what a running process has open, from outside the process itself, with no instrumentation required.

**Pipes connect file descriptors between processes.** `cmd1 | cmd2` creates a kernel pipe (a bounded, in-memory buffer with a read end and a write end) and rewires `cmd1`'s stdout (fd 1) to the pipe's write end, and `cmd2`'s stdin (fd 0) to the pipe's read end. `cmd1` has no idea its output isn't going to a terminal — it just writes to fd 1 as always; the shell did the redirection before `cmd1` even started. This is precisely why Unix pipelines compose so cleanly: every well-behaved program just reads stdin and writes stdout, ignorant of what's actually connected to those descriptors.

**Permissions, precisely.** Each file has an owner, a group, and permission bits for owner/group/other (read, write, execute — `rwx`), shown as `-rwxr-xr--` in `ls -l`. For directories, execute permission means "can enter/traverse" (needed to `cd` into it or access files inside by path), not "can run" — a very common point of confusion, since execute means something different for a directory than for a regular file.

**Redirection is just fd manipulation with syntax.** `cmd > file` opens `file` and duplicates its fd onto fd 1 before `cmd` runs (`2>&1` duplicates fd 2 onto whatever fd 1 currently points to — order matters, and `2>&1 > file` behaves differently from `> file 2>&1` for exactly this reason, since each redirection is applied left to right, changing what "current fd 1" means for the next redirection).

## Attempt

1. Run `ls -l /proc/self/fd/` (or, in a script, `/proc/$$/fd/` for the current shell) and identify which entries correspond to stdin/stdout/stderr. Then run the same command with output redirected to a file (`ls -l /proc/self/fd/ > out.txt`) and note how fd 1's target changed in the listing.

2. Build a 3-stage pipeline (e.g. `cat somefile.txt | grep pattern | wc -l`) and, while it's running (insert a `sleep` in the middle stage to give yourself time), inspect `/proc/<pid>/fd/` for each of the three processes in a separate terminal. Confirm that the middle process's stdin fd points to a pipe (shown as `pipe:[inode]`) matching the first process's stdout fd, and its stdout fd points to a different pipe matching the third process's stdin.

3. Demonstrate the redirection-order trap directly: run a command that writes to both stdout and stderr (e.g. `ls /exists /doesnotexist`), first with `> out.txt 2>&1` and then with `2>&1 > out.txt`, and compare what ends up in `out.txt` versus what appears on the terminal in each case. Explain the difference using the left-to-right fd duplication rule from Learn.

4. Create a file, then use `chmod` to set specific permission combinations, and predict before testing whether a different user (or `sudo -u` a different account, if available) could read, write, or execute it under each combination. Specifically test the directory-execute-vs-read distinction: create a directory with read but no execute permission (`chmod 600` on a directory, i.e. `rw-------`) and confirm you can list it might show contents via some methods but cannot `cd` into it or stat files inside by path.

## Verify

For step 3, show the actual contents of `out.txt` in both orderings side by side, and state in one sentence why they differ, referencing the specific fd each redirection duplicated and when.

## Failure drill

Run a command that reads a very large amount of output through a pipe into a slow consumer (e.g. `yes | head -1` is instant since `head` exits early, but try `yes | some_slow_processing_command` conceptually, or more safely: `seq 1 100000000 | sleep 5` won't demonstrate it well since sleep doesn't read stdin — instead use `seq 1 100000000 | (sleep 2; cat) | wc -l` and observe the pipeline pauses). Explain, using the "bounded, in-memory buffer" fact from Learn, why a fast producer writing into a pipe with no one reading yet will eventually block (the producer's `write()` call stalls) rather than buffering unboundedly in memory — this is the backpropagating flow control built into pipes, and it's the same fundamental mechanism (bounded buffer, blocking producer) that reappears in Go channels and in network flow control later in this curriculum.

## Transfer

In Go, `os.Stdin`, `os.Stdout`, and `os.Stderr` are just `*os.File` values wrapping fd 0, 1, 2 — identical to the shell's convention covered in this lesson. Write a tiny Go program that reads from stdin and writes to stdout with no special-casing, then run it in a shell pipeline (`cat file.txt | go run main.go | wc -l`) and confirm it works exactly like a compiled Unix filter program would, with no code in your program aware it's part of a pipeline.

## Done when

You can explain what a file descriptor is and inspect a running process's open descriptors via `/proc`, you've directly observed pipe wiring between processes rather than just trusting it happens, and you can correctly predict the output of differently-ordered redirection commands using the left-to-right duplication rule rather than by trial and error.
