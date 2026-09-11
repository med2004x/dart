# Lesson 8: xv6 Implementation Labs

## Objective

Bridge from OS theory (Lessons 1-7) to real kernel code by working directly in xv6, a small, teaching-oriented Unix-like kernel — reading actual kernel source before modifying it, and implementing at least one real system call or scheduling change end to end.

## Prerequisites

Lessons 1-7, completed. This lesson assumes familiarity with C (C track) and the process/scheduling/synchronization concepts already covered, and applies them to real, if simplified, kernel source rather than toy simulators.

## Learn

**Why a teaching kernel, not Linux itself.** The real Linux kernel is millions of lines of code, accumulated over three decades of hardware-support and performance work — invaluable to know how to navigate but a poor starting point for understanding the *fundamental* mechanisms this track has covered. xv6 (MIT's teaching kernel, originally modeled on 1970s Unix, rewritten for RISC-V) implements the core ideas from Lessons 1-7 — process creation, scheduling, virtual memory, system calls, a simple filesystem — in a few thousand lines total, small enough to read start to finish, while being real, runnable, bootable kernel code rather than a simulation.

**Read before you modify.** This is worth stating as a rule rather than just a suggestion: xv6's existing implementation of `fork()`, the scheduler, or the trap-handling path for syscalls already embodies correct, working solutions to problems closely related to what a given lab asks you to extend. Reading the existing code for the *related* mechanism before implementing something new (e.g. reading how an existing syscall is wired up before adding a new one) is dramatically more efficient than reverse-engineering the pattern from lab instructions alone, and it's standard practice in real kernel and systems development, not a crutch specific to learning.

**What "system call implementation" actually requires, end to end.** Adding a new syscall to xv6 (a common lab exercise) touches multiple layers you've studied separately in this track: a user-space wrapper function, a syscall number registered in a table, a kernel-space handler function that does the actual work, and — if the syscall touches process or memory state — code interacting with the exact process/page-table structures Lessons 1 and 4 discussed abstractly. Doing this once, completely, is what turns "I understand syscalls conceptually" into "I've implemented one."

## Attempt

MIT's 6.1810 (formerly 6.S081) course provides the canonical xv6 labs, freely available; work through them using the actual course materials rather than reconstructing equivalent labs from scratch, since the lab infrastructure (test harness, provided starter code) is part of what makes this exercise tractable. Reference: `https://github.com/mit-pdos/xv6-riscv` and the course's published lab assignments.

Minimum required labs, chosen to touch every major topic from Lessons 1-7:

1. **System call lab**: implement at least one new syscall end to end (a common starting lab adds something like a `trace` syscall or a simple system-information syscall). Before writing any new code, read xv6's existing syscall dispatch path (`syscall.c`, the syscall number table, and one existing simple syscall's implementation) and write a short note (a few sentences) describing the path a syscall takes from user-space invocation to kernel-space handler and back.

2. **Scheduling-related lab or exercise**: read xv6's scheduler implementation (`proc.c`'s `scheduler()` function) and identify which of the policies from Lesson 3 (FCFS, SJF, RR) it most closely resembles, or how it differs from all three. If the course offers a scheduling-modification lab, complete it; if not, write a short design note proposing one concrete, specific change to xv6's scheduler (e.g. adding basic priority levels) and explain, in terms of Lesson 3's turnaround/response-time tradeoffs, what effect you'd expect.

3. **Memory-related lab**: complete a lab involving xv6's page table code (labs commonly cover something like implementing `mmap`-style memory mapping, or lazy allocation) — this directly extends computer architecture Lesson 6 and OS Lesson 4's virtual memory coverage from reading about page tables to modifying real page-table-manipulation code.

4. For each of the three labs above, write a short design note (2-4 sentences each) *before* implementing: what you're changing, which existing code path it touches, and what specific test or observation will confirm it works correctly — this mirrors the actual engineering discipline of writing down intent before diving into an unfamiliar codebase's implementation.

## Verify

For each completed lab, run the course-provided test suite (xv6 labs typically include automated grading scripts) and report the actual pass/fail output — this lesson's "done" state is defined by the lab infrastructure's own tests passing, not self-assessment.

## Failure drill

Pick one lab you completed and deliberately introduce a regression related to a specific concept from this track — for example, in the syscall lab, remove the syscall number's registration in the dispatch table while leaving the handler function itself intact, and observe the specific failure mode (the syscall not being found/dispatched, likely surfacing as the wrong return value or an unimplemented-syscall path being hit, rather than your handler's code ever running). Explain, using OS Lesson 1's syscall-dispatch discussion, exactly why this specific break produces this specific symptom rather than some other kind of failure — connecting the abstract "syscalls go through a dispatch table" idea to the literal line of xv6 source code responsible for it.

## Transfer

Compare one specific piece of xv6's implementation (your choice — the syscall dispatch table, the scheduler, or the page-table code from whichever labs you completed) to the corresponding real Linux mechanism, at whatever level of detail you can find via documentation or source browsing (the Linux kernel source is public) — you're not expected to fully understand Linux's much more complex real implementation, just to identify one or two concrete differences (e.g. Linux's much larger, versioned syscall table with backward-compatibility constraints xv6 doesn't need to worry about) that illustrate what a teaching kernel deliberately simplifies away.

## Done when

You've completed and passed the automated tests for at least the three labs specified above, you wrote a design note before implementing each one (not after, as documentation-after-the-fact), and you can trace, using your own completed syscall lab as the concrete example, the full path a system call takes from user-space code through the kernel and back, citing actual xv6 source rather than describing it only in the abstract.
