# Lesson 2: Go Profiling

## Objective

Use Go's built-in `pprof` profiler to find real CPU, allocation, and blocking hot spots in a genuinely slow program, and make one targeted, measured change based on profiler evidence — directly applying Linux tools Lesson 5's profiling discipline with Go-specific tooling.

## Prerequisites

Lesson 1 (measurement discipline — profiling identifies *where* to optimize; this lesson's "re-measure after" step depends on Lesson 1's rigor to confirm a change actually helped), Linux tools Lesson 5 (the general profiling principle this lesson specializes for Go).

## Learn

**Why profiling, not intuition, should drive optimization effort.** Linux tools Lesson 5 established the core principle: intuition about where time is spent is often wrong. Go's `pprof` makes this concrete and precise for Go programs specifically: it can show CPU time by function (a sampling profiler, periodically capturing the call stack, building a statistical picture of where execution time actually goes), memory allocations by call site (which functions are allocating, and how much — directly relevant to Lesson 3's GC-pressure discussion), and goroutine blocking profiles (where goroutines are spending time waiting — on locks, channels, or I/O — rather than actually running).

**CPU profiling, mechanically.** `pprof.StartCPUProfile` (or `go test -cpuprofile=cpu.out -bench=.`) samples the call stack at a fixed frequency while the program runs, and `go tool pprof cpu.out` lets you explore the results — `top` shows functions by cumulative sampled time, and the `web` command (or `go tool pprof -http=:8080`) generates an interactive flame graph, where wider bars represent more time spent, letting you visually identify the actual hot path through your call graph rather than guessing from source code alone.

**Allocation profiling: finding where memory pressure comes from.** `go test -memprofile=mem.out -bench=.` records allocation counts and sizes by call site. This matters because excessive allocation doesn't just cost the allocation itself — it increases garbage collector work (Lesson 3 covers this specifically), which can dominate a program's actual CPU time even when no single function looks expensive in a CPU profile, since GC overhead is often distributed across the program's execution rather than attributed to one obvious hot function.

**Blocking profiling: finding where goroutines wait, not where they compute.** A CPU profile only samples when a goroutine is actively running on a core — exactly the limitation computer architecture Lesson 4/Linux tools Lesson 5 warned about for I/O-bound programs. Go's blocking profiler (`runtime.SetBlockProfileRate`, or `go tool pprof` with the block profile) specifically captures time spent blocked on synchronization (mutex contention, channel operations) — essential for diagnosing concurrency-related slowness that a CPU profile alone would show as "not much happening" rather than revealing the actual bottleneck.

**Making one change and re-measuring — the discipline that closes the loop.** After identifying a hot spot via profiling, make exactly one targeted change addressing it, then re-benchmark (Lesson 1's rigor) to confirm the change actually helped, by how much, and that it didn't regress something else — changing multiple things at once makes it impossible to attribute a measured improvement (or lack of one) to any specific change, defeating the entire purpose of profiling-driven optimization.

## Attempt

1. Write (or use) a deliberately inefficient Go function — e.g. one doing unnecessary repeated string concatenation in a loop (a classic Go performance anti-pattern, since naive `+=` string concatenation reallocates on every iteration) — and benchmark it using Lesson 1's methodology to establish a baseline.

2. Generate a CPU profile of your benchmark (`go test -cpuprofile=cpu.out -bench=.`) and use `go tool pprof -http=:8080 cpu.out` (or the `top`/`web` commands in the interactive CLI) to identify the actual hot function. Confirm it correctly points to your inefficient string-concatenation code, and report the actual percentage of total sampled time attributed to it.

3. Make exactly one targeted fix based on the profile (e.g. replace the naive concatenation with a `strings.Builder`, which avoids repeated reallocation), and re-benchmark using Lesson 1's proper multi-run methodology. Report the before/after performance numbers with variance, confirming a real, statistically meaningful improvement (not just noise, per Lesson 1's failure drill).

4. Generate a memory allocation profile (`go test -memprofile=mem.out -bench=.`) for both the before and after versions, and compare allocations-per-operation between them using `go tool pprof`'s allocation-focused views — confirm the fix reduced allocation count/volume as well as time, connecting the CPU-time improvement to its actual underlying mechanism (fewer, larger allocations instead of many small, repeated ones).

## Verify

Present your actual `pprof` output (a screenshot or text summary of the `top` view) identifying the hot function before your fix, your Lesson-1-rigorous before/after benchmark comparison with variance, and your before/after allocation profile comparison — real tool output, not just a description of what you'd expect to see.

## Failure drill

Make a *second*, unrelated change to your function at the same time as your intended fix (e.g. also reorder some unrelated logic, or add an unrelated micro-optimization you assumed was harmless) and re-benchmark. If the combined change shows an improvement, attempt to determine — by reverting just the second, unrelated change and re-benchmarking with only the original fix — whether the unrelated change contributed anything, hurt performance, or was neutral. Explain, using whatever you actually find, why changing multiple things simultaneously (even if the end result "seems better") undermines the specific causal claim profiling-driven optimization is supposed to let you make — you can no longer say with confidence which change (or combination) produced the measured effect, precisely the discipline Learn's final paragraph described as necessary and that this drill deliberately violated to make the cost concrete.

## Transfer

If TARDOC's transcription pipeline or Lead Sourcer's scoring logic has ever felt slow without a specific profiled cause identified, describe what a real `pprof` CPU and allocation profile run against that actual workload would likely reveal, given what you know about its structure (e.g., is it more likely CPU-bound in local processing, or dominated by allocation overhead from string/data processing, or largely blocked on external I/O per computer architecture Lesson 7's numbers) — and state what specific `pprof` command you'd run first to test that hypothesis, rather than guessing at an optimization to make.

## Done when

You've profiled a real Go program, correctly identified its actual hot spot via `pprof` (not by inspection alone), made and measured exactly one targeted fix with Lesson 1's rigor confirming a real improvement, and connected the CPU-time improvement to its underlying allocation-count reduction using a real memory profile comparison — plus directly experienced, via the failure drill, why isolating one change at a time is necessary for the causal claim "this fix helped" to actually be justified.
