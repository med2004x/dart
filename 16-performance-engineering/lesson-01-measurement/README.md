# Lesson 1: Measurement Before Optimization

## Objective

Build disciplined, statistically sound benchmarks — accounting for warmup, variance, and environmental noise — as the mandatory prerequisite to any optimization work, directly extending Linux tools Lesson 5's "measure before optimizing" principle with the rigor real performance work requires.

## Prerequisites

Linux tools Lesson 5 (measurement discipline — this lesson goes deeper on doing it correctly), statistics track Lesson 1-2 (descriptive statistics and sampling — a benchmark result is a sample, and needs to be treated with the same rigor).

## Learn

**Why a single timed run is not a benchmark.** Running a piece of code once and timing it gives you one sample from a distribution of possible run times, affected by CPU scheduling noise (OS Lesson 3), cache state (computer architecture Lesson 4), background system load, and — critically for managed-runtime languages like Go — garbage collection timing and JIT/compiler warmup effects. A single measurement cannot distinguish real signal from this noise; statistics track Lesson 2's sampling-distribution lesson applies directly: you need enough repeated measurements to characterize the actual distribution, not just one draw from it.

**Warmup: why the first few iterations of a benchmark often don't represent steady-state behavior.** Many systems have a "cold start" cost — CPU caches not yet populated with the relevant data (computer architecture Lesson 4), a Go runtime's garbage collector not yet in steady operation, OS-level page cache (OS Lesson 4) not yet warm for files being accessed. Benchmark frameworks (Go's `testing.B`, for instance) typically run a warmup phase before the timed measurements begin, specifically to let these effects settle before recording numbers meant to represent normal, steady-state performance rather than one-time startup cost.

**Variance and confidence: reporting a range, not just a single number.** A responsible benchmark result reports not just a mean but some measure of spread (standard deviation, or a percentile range) across multiple runs, and ideally states how many independent runs/iterations that's based on. A performance claim like "this change made it faster" based on one before-run and one after-run, each a single measurement, cannot distinguish a real improvement from ordinary run-to-run variance — statistics track Lesson 3's hypothesis-testing framework applies directly here, even if informally: is the observed difference larger than what you'd expect from noise alone, given the variance you've actually measured?

**Recording the environment: what a benchmark result doesn't travel with unless you write it down.** CPU model, available cores, memory, OS, language/runtime version, and background load at the time of measurement all affect results — a benchmark number without this context is difficult to meaningfully compare against a later re-run (was a later "improvement" real, or did it happen on a less-loaded machine?), and difficult for anyone else to reproduce or sanity-check.

## Attempt

1. Write a small, deliberately simple function to benchmark (e.g. a sorting function, or a string-processing function) and run it using Go's `testing.B` benchmark framework (`go test -bench=.`), which automatically handles iteration count and reports operations-per-second along with per-operation timing.

2. Run the same benchmark 5 separate times (separate invocations of `go test -bench=.`, not just relying on the framework's internal iteration count) and record all 5 results. Compute the mean and standard deviation across these 5 runs (statistics track Lesson 1), and report both — not just the single "best" or most recent number.

3. Deliberately measure without proper warmup: implement a naive timing loop (not using `testing.B`, which handles this correctly) that times only the *first* iteration of your function, compare that single cold measurement against the steady-state average from step 2, and report the difference — if your function/environment shows a meaningful warmup effect, this should show a measurable discrepancy; if it doesn't, that's also a legitimate, reportable finding worth noting (not every workload has a strong warmup effect).

4. Document your benchmark's environment explicitly: CPU model/core count, available memory, Go version, OS, and note whether anything else was running on the machine during measurement (ideally, minimize background load for benchmark runs, and note explicitly if you couldn't).

## Verify

Present your 5-run results from step 2 with mean and standard deviation, your step 3 cold-vs-warm comparison with actual numbers, and your step 4 environment documentation — a complete, reproducible benchmark report, not just a single number.

## Failure drill

Take two trivially, functionally-identical versions of your benchmarked function (e.g. the exact same code, copy-pasted with no real change) and benchmark both using single-run measurements only (one run each, no repetition). Compare the two single numbers and observe that they likely differ somewhat, purely due to run-to-run noise — despite the code being identical. Then repeat using your step 2 methodology (5 runs each, mean and standard deviation) and confirm the two "different" functions' results now overlap within their reported variance, correctly indicating no real difference exists. Explain, using this direct before/after comparison, why a naive single-run comparison can produce a false conclusion ("version B is faster!") purely from noise, and why this specific failure mode is exactly what proper benchmark methodology (multiple runs, reported variance) is designed to prevent.

## Transfer

If you've previously made a performance claim about TARDOC or Mahall's code (e.g. "the new query is faster" or "the refactored function performs better") based on informal, single-measurement testing, revisit that claim using this lesson's methodology — rerun both the before and after versions with proper repetition and variance reporting, and state honestly whether the original claim holds up under this more rigorous standard, or whether it was within the range of ordinary measurement noise.

## Done when

You've produced a properly repeated, variance-reported benchmark for a real function rather than a single timed run, you've directly demonstrated (with your own identical-code test) how single-run comparisons can produce false conclusions that proper repeated measurement correctly avoids, and you've documented your benchmark's environment completely enough that someone else could meaningfully compare their own results against yours.
