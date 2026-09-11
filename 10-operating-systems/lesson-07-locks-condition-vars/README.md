# Lesson 7: Synchronization Primitives

## Objective

Go beyond the basic mutex from Lesson 2 into condition variables and semaphores, implement a correct bounded buffer (the canonical producer/consumer synchronization problem), and deliberately create and diagnose a deadlock.

## Prerequisites

Lesson 2 (threads, mutexes, and the basic data-race concept this lesson extends).

## Learn

**Why a mutex alone isn't enough for producer/consumer coordination.** A mutex protects a critical section from concurrent access, but says nothing about *waiting for a condition* — e.g. a consumer thread that finds a shared buffer empty needs to wait until a producer adds something, without holding the lock the whole time (which would prevent the producer from ever acquiring it to add anything). Busy-waiting (repeatedly checking-and-unlocking-and-relocking in a loop) technically works but wastes CPU continuously polling — condition variables exist to let a thread sleep efficiently until a specific condition becomes true, without polling.

**Condition variables.** A condition variable is always used together with a mutex. `wait(cv, mutex)` atomically releases the mutex and blocks the calling thread until another thread calls `signal(cv)` or `broadcast(cv)`, at which point it re-acquires the mutex before returning — the atomicity of "release and block" is essential; without it, there's a race window between checking a condition and starting to wait where a signal could be missed. The standard pattern is always to re-check the condition in a loop after waking (`while (!condition) wait(cv, mutex);`, not `if`), because spurious wakeups are a real possibility in most implementations, and because multiple waiters can be woken by a broadcast even though only one of them should actually proceed.

**Semaphores.** A counter-based synchronization primitive: `wait()`/`P()` decrements the counter (blocking if it would go negative), `signal()`/`V()` increments it. A semaphore initialized to 1 behaves like a mutex (binary semaphore); a semaphore initialized to N can be used to limit concurrent access to at most N resources at once (e.g. a connection pool limiter) — a use case a plain mutex can't directly express, since a mutex only ever represents "0 or 1 holders."

**Deadlock.** A situation where two or more threads are each waiting for a resource the other holds, with neither able to proceed. The classic minimal case: thread A holds lock 1 and waits for lock 2; thread B holds lock 2 and waits for lock 1 — neither will ever release what the other needs. The standard, general prevention strategy is **lock ordering**: establish a global, consistent order in which locks must always be acquired (e.g. always lock the lower-numbered resource first) across the entire codebase, so the circular-wait pattern above becomes structurally impossible — both threads would have to acquire lock 1 before lock 2, eliminating the possibility of the reversed acquisition order that creates the cycle.

## Attempt

1. Implement a **bounded buffer** (a fixed-capacity circular queue shared between producer and consumer threads) using a mutex and two condition variables: one signaled when the buffer becomes non-full (producers wait on this when the buffer is full), one signaled when the buffer becomes non-empty (consumers wait on this when the buffer is empty). Have multiple producer threads and multiple consumer threads running concurrently against the same buffer, and verify — after a run with a known total number of produced items — that the consumer(s) received exactly that many items, with no lost or duplicated items.

2. Deliberately remove the `while` loop re-check around your condition variable waits (using `if` instead) and, ideally with multiple consumers, try to reproduce a case where a spuriously or over-broadly woken thread proceeds when the condition it was waiting for isn't actually true anymore (a broadcast waking multiple consumers when only enough data exists for one of them). This can be timing-dependent and hard to reliably reproduce — document your attempt and reasoning even if you can't force a visible failure every run, and explain in your own words why the `while`-loop re-check pattern is considered mandatory practice regardless of whether your specific test happens to trigger the bug.

3. Deliberately construct the classic two-lock deadlock: two threads, two mutexes, thread A acquires lock 1 then tries to acquire lock 2 (with a deliberate small delay between the two acquisitions to make the race window reliably hit), thread B acquires lock 2 then tries to acquire lock 1. Run it and confirm the program hangs (both threads permanently blocked).

4. Diagnose the hung deadlock from step 3 using a debugger (Linux tools Lesson 4): attach to the hung process, get a backtrace for each thread (`thread apply all bt` in gdb, or the equivalent), and confirm you can see each thread is blocked inside a lock-acquisition call, waiting on the specific lock the other thread holds — using the debugger to directly diagnose the deadlock's structure rather than just observing "it hung."

## Verify

For step 1, report the exact count of items produced and consumed across at least 3 different runs with varying numbers of producer/consumer threads, confirming correctness holds consistently, not just once by luck.

## Failure drill

Fix the deadlock from step 3 by applying consistent lock ordering (both threads always acquire lock 1 before lock 2, regardless of which "logical" order the operation conceptually suggests). Rerun the same test many times (e.g. 100 iterations in a loop) and confirm it never hangs, in contrast to the original version which should have hung reliably or near-reliably given the deliberate delay you inserted. Explain why lock ordering is a *structural* fix (making the deadlock's precondition impossible) rather than a probabilistic one (like just hoping the timing doesn't align badly) — and why this matters when reasoning about correctness: a fix that merely makes a race window smaller still leaves a bug that could resurface under different timing (a different machine, a different load level), whereas eliminating the structural possibility doesn't depend on timing at all.

## Transfer

Go's standard library provides `sync.Cond` (condition variables) and buffered channels, which can express the bounded-buffer pattern from step 1 more idiomatically — a buffered channel with capacity N *is* essentially a bounded buffer with built-in blocking semantics for both full (send blocks) and empty (receive blocks) cases, without needing explicit condition variables at all. Rewrite your step 1 bounded buffer using a Go buffered channel instead of a mutex+condvar pair, confirm it behaves correctly under the same producer/consumer test, and state explicitly what specific synchronization complexity the channel abstraction eliminated compared to your explicit condition-variable implementation.

## Done when

Your bounded buffer implementation is correct under concurrent producer/consumer load across multiple runs, you've reliably reproduced a two-lock deadlock and diagnosed its exact structure using a debugger rather than just observing a hang, and you've fixed it with lock ordering and confirmed — across many repeated runs — that the fix is structural, not just a reduction in how often the bug shows up.
