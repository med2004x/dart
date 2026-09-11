# Lesson 4: Lock Contention and Parallelism

## Objective

Measure lock contention directly, understand false sharing as a subtle cache-coherency cost invisible to naive reasoning about "correct" concurrent code, and explain concretely why adding more goroutines/threads to a parallel workload eventually stops improving (and can even worsen) throughput.

## Prerequisites

OS Lesson 2 (mutexes and data races — this lesson measures the performance cost of correct synchronization, not just its correctness), computer architecture Lesson 4 (caching — false sharing is specifically a cache-coherency phenomenon).

## Learn

**Lock contention: the throughput cost of correct synchronization.** OS Lesson 2 established that a mutex correctly prevents data races. But a mutex also serializes access to its critical section — if many goroutines frequently contend for the same lock, they spend time waiting rather than running, and this waiting time is real, measurable throughput lost to synchronization overhead, not just an abstract concern. A profiler's blocking profile (Lesson 2 of this track) is precisely the tool for measuring this directly, rather than assuming a given lock is or isn't a bottleneck.

**False sharing: a cache-coherency cost that occurs even without any actual logical data conflict.** Computer architecture Lesson 4 established that CPU caches operate on cache lines (typically 64 bytes), not individual variables. If two different goroutines, running on different CPU cores, each modify *different* variables that happen to reside on the *same* cache line, the CPU's cache-coherency protocol treats this as if they were contending for the same data — each core's write invalidates the other core's cached copy of that line, forcing expensive re-fetches, even though the two goroutines are never actually touching the same logical variable and there's no data race by any correctness definition. This is a genuinely subtle, easy-to-miss performance bug: the code is completely correct, race-free, and passes every correctness test, yet performs far worse under real parallel load than the (correct) reasoning about data independence would suggest.

**Why parallel scaling flattens (and sometimes reverses) as you add more workers.** Amdahl's Law formalizes this: if a fraction `p` of a workload can be parallelized and `1-p` is inherently sequential (e.g. time spent acquiring/holding a shared lock, or coordinating results), the maximum possible speedup from N parallel workers is bounded by `1 / ((1-p) + p/N)` — as N grows large, this approaches `1/(1-p)`, a hard ceiling determined entirely by the sequential fraction, no matter how many additional workers you throw at it. Beyond this theoretical ceiling, *real* systems often see throughput actually *decrease* past some worker count, due to lock contention (this lesson's first topic) and false sharing (this lesson's second topic) both getting worse as more concurrent workers compete for the same limited synchronization resources.

## Attempt

1. Implement a shared counter incremented by many concurrent goroutines, protected by a single `sync.Mutex` (correct, per OS Lesson 2). Benchmark it (Lesson 1's methodology) across an increasing number of concurrent goroutines (e.g. 1, 2, 4, 8, 16, 32) all contending for the same lock, and plot/report throughput (operations completed per second) at each level.

2. Use Go's blocking profile (`go test -blockprofile=block.out -bench=.` or `runtime.SetBlockProfileRate`) to measure actual time spent waiting on the mutex at your highest contention level (e.g. 32 goroutines) from step 1, and report the measured blocking time as a percentage of total wall-clock time — direct, quantitative evidence of contention, not just an inferred effect from the throughput numbers alone.

3. Reproduce false sharing directly: create a small array of counters (e.g. `[4]int64`), have 4 separate goroutines each repeatedly increment a *different* element of the same array (no logical conflict — each goroutine only ever touches its own element), and benchmark this against a version where each goroutine's counter is instead padded to occupy its own separate cache line (e.g. wrapped in a struct with padding bytes added to push each counter onto its own 64-byte line). Report the actual measured throughput difference between the falsely-shared and properly-padded versions.

4. Fit Amdahl's Law to your step 1 data: estimate the sequential fraction `p` your measured throughput implies (by fitting your observed speedup-vs-worker-count curve against the formula), and use that fitted `p` to predict the theoretical maximum speedup achievable regardless of worker count — compare this prediction against your actual highest-worker-count measurement to see how closely real behavior tracked the theoretical model.

## Verify

Present your step 1 throughput-vs-worker-count table/plot showing the flattening (or reversal) pattern, your step 2 measured blocking-time percentage, your step 3 false-sharing vs. padded throughput comparison with actual numbers, and your step 4 fitted Amdahl's Law prediction alongside your real measured data.

## Failure drill

Take your step 3 false-sharing reproduction and, instead of padding to separate cache lines, attempt a naive-seeming "fix" of simply using 4 entirely separate variables (not in an array) instead of array elements — but declare them as consecutive local variables in a way that a compiler *might* still place adjacent in memory (this specific outcome is compiler/platform-dependent and worth investigating rather than assuming). Check, using tooling if available (or by reasoning about your platform's likely behavior) whether this naive fix actually resolves the false sharing or not, and explain why "not being in the same array" doesn't automatically guarantee separate cache lines — the actual guarantee only comes from explicit padding (or the compiler/allocator happening to place them favorably, which isn't something you should rely on without verifying). This is meant to surface a genuinely subtle point: false-sharing fixes need to reason about actual memory layout, not just source-level variable separation, which can be misleading about what's really happening at the cache-line level.

## Transfer

If TARDOC's Celery worker pool or any concurrent processing in your systems uses shared counters or shared mutable state across goroutines/workers (e.g. a shared rate-limiter state, or aggregate statistics updated from multiple concurrent request handlers), describe whether that shared state's access pattern more closely resembles your step 1 lock-contention scenario or your step 3 false-sharing scenario, and what this lesson's specific measurement tools (blocking profile for contention, cache-line-aware reasoning for false sharing) would tell you about whether it's currently a real bottleneck at your actual, current concurrency level.

## Done when

You've measured real lock-contention-driven throughput flattening across increasing worker counts, backed by an actual blocking-profile measurement of wait time, you've directly reproduced false sharing's performance cost despite the code being completely logically correct, and you've fit Amdahl's Law to real measured data and compared its prediction against your actual highest-concurrency measurement, understanding both the theoretical ceiling and the additional real-world factors (contention, false sharing) that can make actual behavior worse than the theoretical model alone predicts.
