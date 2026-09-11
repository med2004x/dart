# Lesson 2: Threads and Concurrency

## Objective

Understand threads as units of execution sharing an address space, deliberately produce a data race, and fix it correctly with synchronization — building the concrete intuition that makes "just add a mutex" an informed decision rather than a reflex.

## Prerequisites

Lesson 1 (processes — a thread is best understood as a lighter-weight sibling concept, sharing what a process would otherwise keep separate).

## Learn

**Thread vs. process, precisely.** A process has its own address space, file descriptor table, and other resources — fully isolated from other processes by default. A thread is a unit of execution *within* a process, sharing that process's address space and file descriptors with any other threads in the same process. Multiple threads can read and write the same memory directly, with no OS-mediated communication needed — which is both what makes threads efficient for cooperating work and exactly what makes them dangerous without discipline: nothing stops two threads from writing to the same memory location simultaneously.

**Data races, precisely.** A data race occurs when two or more threads access the same memory location concurrently, at least one access is a write, and there's no synchronization ordering the accesses. The outcome is undefined — not "some unpredictable but bounded result," but genuinely undefined behavior (same category as C track Lesson 5), because the CPU and compiler are both free to reorder, cache, or interleave operations in ways that produce results a sequential reading of the code would never suggest.

**Mutexes.** A mutual exclusion lock: `lock()` blocks until the lock is available, then acquires it; `unlock()` releases it. Code between `lock()` and `unlock()` (a critical section) is guaranteed to run without another thread concurrently executing the same critical section — this is the primary tool for turning an unsynchronized shared-memory access into a well-defined one.

**Why "just add a mutex everywhere" isn't sufficient understanding.** Locking too coarsely (one giant lock around everything) eliminates concurrency benefits entirely — threads spend all their time waiting for the lock instead of running in parallel, an outcome you can measure directly (Lesson 3 of computer architecture's pipeline analogy applies loosely here: serialized access is the concurrency equivalent of a structural hazard). Locking too finely, or inconsistently, risks missing an access path entirely (a race the developer didn't realize needed protection) or introducing deadlock (Lesson 7 of this track). The actual skill is identifying precisely which shared state needs protection and matching lock granularity to the real contention pattern.

## Attempt

1. Write a program (C with pthreads, or Go with goroutines — pick whichever your track context makes more natural, though C is more instructive here since Go's race detector will make the next steps almost too easy) with multiple threads/goroutines all incrementing a shared, unprotected counter many times (e.g. 4 threads, each incrementing a shared `int` 100,000 times). Run it several times and observe the final counter value is inconsistently *less* than the expected `4 × 100,000` — demonstrating the race directly, with a wrong number as concrete evidence rather than an abstract warning.

2. If using Go, run the same program with the race detector enabled (`go run -race main.go`) and observe it explicitly reporting the race, including the exact lines of the conflicting reads/writes — compare this precise diagnostic to C's lack of an equivalent built-in tool (ThreadSanitizer, `-fsanitize=thread`, is the closest C/C++ equivalent — use it if available on your platform).

3. Fix the race from step 1 by adding a mutex around the increment operation (`lock(); counter++; unlock();` or Go's `sync.Mutex`). Rerun multiple times and confirm the final counter value is now always exactly correct.

4. Deliberately lock too coarsely: wrap far more code than necessary in the same mutex (e.g. lock around an entire loop iteration that does unrelated, independent work in addition to the shared counter increment), and measure the wall-clock time with several concurrent threads/goroutines compared to a version where the lock only wraps the minimal critical section. Report the actual timing difference.

## Verify

For step 1, run the unsynchronized version at least 5 times and report all 5 final counter values — they should vary and should be less than the theoretical correct total, demonstrating the race is real and reproducibly wrong, not a one-off fluke.

## Failure drill

Take the fixed, mutex-protected version from step 3, and introduce a *second* shared variable that's incremented in the same critical section but *read* elsewhere in the program without holding the same lock (e.g. a separate reporting goroutine/thread that reads the counter periodically without locking). Run under the race detector (Go) or ThreadSanitizer (C) again and confirm it flags this new, partial-protection race even though the *write* path is correctly locked — the lesson being that protecting a write doesn't automatically protect all reads; every access path to shared state needs to agree on the same synchronization discipline, not just the "main" one you thought of first.

## Transfer

Compare pthread mutexes (C) to Go's `sync.Mutex` and to Go's alternative concurrency idiom of channels ("share memory by communicating, don't communicate by sharing memory," a phrase from Go's own documentation worth understanding rather than just repeating). Rewrite your step 3 counter-increment example using a Go channel instead of a mutex (e.g. a single goroutine owns the counter and receives increment requests over a channel, rather than multiple goroutines directly mutating shared state) and confirm it's also race-free under `-race`, then state in your own words what fundamentally different strategy this represents compared to the mutex approach — not just different syntax for the same idea.

## Done when

You've directly reproduced a data race with a measurably wrong output, then fixed it and confirmed correctness across repeated runs, and you've demonstrated — using the failure drill's partially-protected second variable — that locking one access path doesn't automatically protect every access path to the same shared state.
