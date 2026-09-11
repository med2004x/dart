# Lesson 6: IPC and Pipes

## Objective

Build a real producer/consumer pipeline using OS-level inter-process communication (pipes and/or local sockets), and handle backpressure and EOF correctly — the two edge cases that separate a toy IPC example from one that behaves correctly under real conditions.

## Prerequisites

Linux tools Lesson 1 (pipes at the shell level — this lesson builds the same mechanism programmatically), OS Lesson 2 (threads/synchronization — IPC between processes shares conceptual ground with synchronization between threads, applied across the process boundary instead of within one address space).

## Learn

**Why IPC exists at all.** Processes are isolated by design (Lesson 1's user/kernel boundary and per-process address spaces) — this isolation is a feature, not an obstacle, but it means processes that *need* to cooperate must use an explicit, OS-provided channel. Pipes, sockets, shared memory, and signals are different IPC mechanisms with different tradeoffs: pipes and sockets move data through the kernel with built-in flow control; shared memory is faster (no data copying through the kernel) but requires the processes to handle their own synchronization, since the OS provides no automatic ordering guarantees on shared memory access.

**Pipes, reprised programmatically.** Linux tools Lesson 1 covered shell-level pipes conceptually. Programmatically, `pipe()` (C) creates a pair of connected file descriptors — one for writing, one for reading — typically used together with `fork()` (OS Lesson 1) so a parent and child each end up with one end. Anonymous pipes (created with `pipe()`) only work between related processes (sharing a common ancestor that created the pipe before forking); **named pipes** (FIFOs, created with `mkfifo`) exist as filesystem entries and can connect unrelated processes.

**Backpressure.** A pipe has a finite kernel buffer (typically 64KB on Linux by default). If a producer writes faster than a consumer reads, the producer's `write()` calls eventually block once the buffer fills — this is backpressure, and it's a *feature*: without it, a fast producer and slow consumer pairing would require unbounded memory to buffer the gap, or would need the producer to implement its own flow control manually. This is the exact mechanism Linux tools Lesson 1's failure drill demonstrated at the shell level; this lesson has you build and observe it directly in code.

**EOF, and why it must be handled explicitly.** When a pipe's write end is closed (all writer file descriptors referencing it), subsequent reads on the read end return 0 bytes — end-of-file — rather than blocking forever. A consumer that doesn't correctly check for a zero-byte read result and instead loops assuming more data is always coming will hang indefinitely once the producer finishes and closes its end. Correctly detecting and acting on EOF is what lets a consumer know "the producer is genuinely done," as opposed to "temporarily has nothing to send right now" (which looks identical from a blocking read's perspective, until you specifically check the return value).

## Attempt

1. Write a C program that creates a pipe, forks, and has the child write a sequence of messages to the pipe's write end while the parent reads and prints them from the read end. Ensure the child closes the write end when done, and the parent's read loop correctly detects EOF (a `read()` returning 0) and exits its loop rather than blocking forever.

2. Modify step 1 so the parent (consumer) deliberately reads slowly (add a small `sleep()` between reads) while the child (producer) writes a large volume of data quickly with no delay. Use `strace` (Linux tools Lesson 5) on the producer process and observe its `write()` calls blocking — visible as calls that take a measurably longer time to return once the pipe buffer fills — directly observing backpressure in your own code, not just the shell-level pipeline from Linux tools Lesson 1.

3. Implement the same producer/consumer pattern using a Unix domain socket instead of a pipe (`socketpair()` in C, which creates a bidirectional connected pair, unlike a pipe's unidirectional pair) or, if working in Go, `net.Pipe()` or a real Unix socket via `net.Listen("unix", ...)`. Confirm equivalent producer/consumer behavior, and note in your own words one concrete difference between a pipe and a socket for this use case (sockets are bidirectional and support more sophisticated addressing/multiple clients; pipes are simpler and unidirectional).

4. Deliberately break EOF handling: modify your consumer to ignore the zero-byte return from `read()` and instead loop unconditionally trying to read more. Run it against a producer that finishes and closes its end, and observe the consumer hang indefinitely (you'll need to manually kill it) rather than exiting cleanly — a direct, hands-on demonstration of why explicit EOF handling isn't optional.

## Verify

For step 2, produce actual `strace` output (or timing measurements) showing at least one `write()` call from the producer taking noticeably longer than the others, correlating with the point where the pipe buffer would have filled given your artificial consumer delay — connect this specific observation back to the pipe buffer size mentioned in Learn.

## Failure drill

Take step 4's hung consumer and, before killing it, inspect it with `strace -p <pid>` attached to the already-running, hung process (rather than starting a fresh `strace` session) — confirm you can see it sitting in a blocking `read()` syscall with no data arriving, distinguishing "hung waiting for more input that will never come because the pipe's write end is closed but I'm not checking for that" from "hung waiting for input that's genuinely still coming." Explain why, from purely observing a blocked `read()` syscall via `strace`, you cannot tell these two cases apart without also knowing the producer's actual state — the consumer's hang is a symptom with (at least) two very different possible causes, and `strace` alone identifies the blocking call but not which of those causes applies.

## Transfer

In Go, `io.Reader`'s contract explicitly defines EOF as a returned `error` value equal to `io.EOF` (rather than a zero-length read the caller must separately check for) — a language-level API design choice specifically meant to make EOF handling harder to accidentally skip, compared to C's convention of a `read()` return value of 0 requiring the caller to remember to check for it as a special case. Write a small Go producer/consumer using an `io.Pipe()` (Go's in-memory, blocking pipe implementation, directly analogous to what you built in C) and confirm your consumer correctly handles `io.EOF` to know when the producer is done — then state explicitly, using your step 4 experience, what specific mistake Go's `io.EOF` error-value convention makes structurally harder to make by accident compared to C's raw zero-byte-read convention.

## Done when

You've built a working producer/consumer pipeline using OS-level IPC (not just language-level channels) and correctly handle EOF, you've directly observed backpressure via `strace` on a real blocking `write()` call rather than just reading about it, and you've deliberately broken EOF handling and observed the resulting hang, then explained why `strace` alone can't distinguish that specific failure from a legitimately-still-waiting consumer.
