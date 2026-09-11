# Lesson 2: Processes and Signals

## Objective

Understand process creation (`fork`/`exec`), exit codes, and the signal delivery mechanism well enough to write a program that shuts down cleanly instead of being killed mid-operation.

## Prerequisites

Lesson 1 (file descriptors — descriptors are inherited across `fork`, which matters for this lesson's process model).

## Learn

**`fork` and `exec`, the two-step process most shells use to run a command.** `fork()` creates a near-exact copy of the calling process — same memory contents (via copy-on-write, so it's cheap despite copying "everything" conceptually), same open file descriptors, different PID. The child process typically then calls one of the `exec` family (`execve`, etc.), which *replaces* the calling process's memory image with a new program — same PID, entirely different code and data. This two-step split is why a shell can set up redirections (Lesson 1) in the child *after* `fork` but *before* `exec` — the child still has the parent's fd table at that point, so `dup2()` calls to rewire fds happen in the narrow window between the two calls.

**Exit codes.** A process exits with a status code (0-255) that its parent can retrieve. Convention: 0 means success, nonzero means some kind of failure — this is a convention, not enforced by the kernel, but universally relied upon by shell scripting (`&&`, `||`, `if cmd; then`) and CI systems.

**Signals.** Asynchronous notifications delivered to a process — not something the receiving process requested at that exact moment, but something the kernel (or another process, via `kill`) delivers, interrupting normal execution flow. Key signals: `SIGTERM` (15) — polite request to terminate, catchable and ignorable; `SIGKILL` (9) — immediate, uncatchable, unignorable termination (the kernel just removes the process, no cleanup code runs); `SIGINT` (2) — what Ctrl-C sends; `SIGCHLD` — sent to a parent when a child exits, letting it know to call `wait()` and reap the child.

**Why graceful shutdown matters practically.** A process that installs a `SIGTERM` handler can flush buffers, close connections cleanly, finish an in-flight request, and exit on its own terms — this is exactly what container orchestrators (Docker, Kubernetes) rely on: they send `SIGTERM` first and wait a grace period, only escalating to `SIGKILL` if the process hasn't exited. A process with no `SIGTERM` handler gets the default action (usually immediate termination) — the same abrupt outcome as `SIGKILL`, just with an extra step the process didn't use.

## Attempt

1. Write a C program (or use `fork()` via a small wrapper if working in Go, which exposes `os/exec` instead — pick whichever language makes the fork/exec split easiest to observe directly; C is more instructive here since it exposes both syscalls explicitly) that calls `fork()`, and in the child, calls `execve` (or `execvp`) to run `/bin/ls`. In the parent, call `wait()` to retrieve the child's exit status and print it. Confirm the parent process's PID is unchanged before and after, while the child's PID (printed from inside the child, before `exec`) differs from the parent's.

2. Modify step 1 so the executed command deliberately fails (e.g. `ls /nonexistent`), and confirm the parent correctly retrieves and prints a nonzero exit status, distinguishing it from the success case.

3. Write a program that installs a handler for `SIGTERM` (using `signal()` or `sigaction()` in C, or `signal.Notify` on an `os.Signal` channel in Go) that prints a message and exits cleanly, rather than accepting the default termination. Run it in the background, send it `SIGTERM` via `kill <pid>`, and confirm your handler's message appears before the process exits — rather than the process just vanishing.

4. Repeat step 3 but send `SIGKILL` (`kill -9 <pid>`) instead. Confirm your handler does *not* run and the process terminates immediately with no cleanup message — directly demonstrating why `SIGKILL` cannot be caught, unlike `SIGTERM`.

## Verify

For steps 3-4, show the actual terminal output (or logs) confirming the handler ran for `SIGTERM` and did not run for `SIGKILL`, with the process's PID and a timestamp or ordering that makes the difference in behavior unambiguous.

## Failure drill

Modify step 3's `SIGTERM` handler to deliberately take a long time (e.g. `sleep(30)` inside the handler, simulating slow cleanup like draining a large in-flight request queue) before exiting. Send `SIGTERM`, then — simulating what a container orchestrator does after its grace period expires — send `SIGKILL` a few seconds later, before the handler finishes. Confirm the process is terminated immediately at the `SIGKILL`, mid-cleanup, with whatever the handler was doing left incomplete. Explain why this is exactly the real-world tension between graceful shutdown and deployment tooling's grace period: a `SIGTERM` handler that takes longer than the orchestrator's configured grace period provides no more protection than having no handler at all, since it gets `SIGKILL`ed before finishing anyway.

## Transfer

In Go, `os/exec`'s `Cmd.Start()` and `Cmd.Wait()} wrap `fork`+`exec`+`wait` into a higher-level API, and `signal.Notify(ch, syscall.SIGTERM)` wraps signal handler installation. Write a small Go program using `signal.Notify` to catch `SIGTERM` and shut down a simple long-running loop (e.g. a `for { ... }` with a `select` on both the signal channel and a ticker) cleanly, and confirm — using the same `kill` commands as steps 3-4 — that it behaves identically to your C signal handler in terms of catching `SIGTERM` but not `SIGKILL`.

## Done when

You can explain the fork/exec split and why redirection setup happens between them, you've directly observed the exit-status retrieval mechanism working for both success and failure cases, and you've demonstrated for yourself — not just read — the difference between a catchable `SIGTERM` and an uncatchable `SIGKILL`, including the real operational tension the failure drill demonstrates between cleanup time and grace-period limits.
