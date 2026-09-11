# Lesson 1: Processes and System Calls

## Objective

Understand the process abstraction as the OS's mechanism for isolating and multiplexing programs, and the user/kernel boundary that every system call crosses.

## Prerequisites

C track Lesson 2 (memory model), Linux tools Lesson 2 (fork/exec/signals — this lesson goes deeper on the OS theory behind what that lesson used practically).

## Learn

**The process abstraction.** A process is the OS's illusion that a program has the entire machine to itself — its own virtual address space (computer architecture Lesson 6), its own view of CPU time (even on a single core, via scheduling, Lesson 3 of this track), its own file descriptor table (Linux tools Lesson 1). This illusion is what lets multiple unrelated programs run "simultaneously" on hardware with far fewer physical cores than running processes, each unaware of the others' existence unless they explicitly communicate (Lesson 6, IPC).

**The user/kernel boundary.** CPUs have privilege levels (rings, on x86) — user-mode code cannot directly access hardware, other processes' memory, or many CPU instructions; it can only do so by asking the kernel via a system call. A syscall triggers a controlled transition (a trap) into kernel mode, where the requested operation (reading a file, allocating memory, sending a signal) is actually performed with full privilege, and control returns to user mode afterward. This boundary is the entire basis of process isolation and OS security — user code physically cannot bypass it to touch another process's memory or the disk directly, only through the syscall interface the kernel chooses to expose.

**Process states.** A process is, at any moment, in one of a small number of states: *running* (actually executing on a CPU core right now), *ready* (runnable, waiting for the scheduler to give it a core), *blocked/waiting* (waiting on something — I/O completion, a signal, a lock — and cannot run even if given a core), or *terminated* (exited, but possibly still holding a slot in the process table as a "zombie" until its parent calls `wait()` to retrieve its exit status, tying directly back to Linux tools Lesson 2's `fork`/`wait` pattern).

**System calls, concretely.** Every "OS operation" a program does — reading a file, allocating memory beyond what's already mapped, creating a process, sending data over a socket — ultimately goes through a syscall. `strace` (Linux tools Lesson 5) works by intercepting exactly these transitions. The C standard library's `malloc`, `fopen`, `printf` are all user-space wrappers that eventually invoke syscalls (`mmap`/`brk`, `open`, `write`) — the library function and the syscall are different things, with the library often batching or buffering to avoid crossing the expensive user/kernel boundary on every single call.

## Attempt

1. Use `strace -c` (from Linux tools Lesson 5) on a simple program (e.g. one that reads a file and prints its contents) and identify which specific syscalls correspond to which C library calls you wrote (`fopen`→likely `openat`, `fread`→`read`, `printf`→eventually `write`). Note any syscalls you didn't expect (e.g. `mmap` calls from the C library's own startup or buffering).

2. Write a program that forks a child, and in the parent, poll the child's state via `/proc/<pid>/status` (look at the `State:` field) at different points — immediately after fork, while the child is doing work, and after the child exits but before the parent calls `wait()`. Confirm you can observe the "zombie" state (`Z`) in that last window, directly demonstrating that a terminated-but-unreaped process still occupies a process table entry.

3. Trace through, in writing, the full state transition sequence for a process that: starts, runs some computation, calls `read()` on a slow file/socket (blocking), receives the data and resumes, then voluntarily yields or is preempted by the scheduler, then eventually finishes and exits. Label each transition with the state before and after.

4. Deliberately create a zombie process that never gets reaped: fork a child that exits immediately, but have the parent sleep for a long time without calling `wait()`. While the parent sleeps, run `ps aux` (or check `/proc/<child_pid>/status`) in another terminal and confirm the child shows as a zombie (`<defunct>` in `ps` output) for the entire time the parent hasn't reaped it.

## Verify

For step 1, produce a table mapping each C library call you used to the actual syscall(s) `strace` shows it invoking, including any syscalls that appeared that weren't obviously tied to a line of your code (these are typically C runtime/libc startup or buffering behavior, worth noting as such).

## Failure drill

Extend step 4's zombie-creation scenario: fork *many* children in a loop (e.g. 50) from a parent that never calls `wait()` on any of them, and observe via `ps aux | grep defunct` (or counting zombie states in `/proc`) that all 50 remain as zombies simultaneously. Explain why this is a real, not just theoretical, resource leak — each zombie entry consumes a slot in the kernel's process table (a finite resource), and a long-running server process that forks children without ever reaping them will eventually exhaust available process table entries, a real historical class of production bug in daemon processes with buggy child-reaping logic.

## Transfer

In Go, you rarely see raw zombie processes because `os/exec`'s `Cmd.Wait()` (or the higher-level `Cmd.Run()`, which calls `Start()` then `Wait()` for you) handles reaping automatically as part of its normal API — but if you use `Cmd.Start()` without ever calling `Wait()`, you can still create the exact same zombie leak. Write a small Go program using `exec.Command(...).Start()` in a loop without calling `.Wait()`, and confirm (via `ps aux`) that this produces zombies in Go exactly as it did in your C example — demonstrating that Go's convenience doesn't eliminate the underlying OS-level requirement, it just makes it easy to forget the requirement even exists.

## Done when

You can map a program's library calls to the actual syscalls they invoke using `strace`, you've directly observed a zombie process existing in the process table and explained why it persists until reaped, and you can articulate why the user/kernel boundary — not just convention — is what actually prevents user code from directly touching hardware or other processes' memory.
