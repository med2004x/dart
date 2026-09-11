# Lesson 6: Capacity Planning (Validated with Load Testing)

## Objective

Take system-engineering Lesson 2's theoretical capacity-planning numbers and validate them against a real load test, finding the actual bottleneck (which may not match your theoretical prediction) — closing the loop between queueing-theory estimation and empirical, measured reality.

## Prerequisites

System-engineering Lesson 2 (capacity planning fundamentals — this lesson validates that lesson's theoretical numbers empirically), performance track Lesson 1 (measurement rigor — load testing needs the same statistical discipline as any other benchmark), performance track Lesson 5 (coordinated omission — a real risk specifically in load-testing tools, directly relevant here).

## Learn

**Why theoretical capacity estimates need empirical validation, not just trust.** System-engineering Lesson 2's queueing-formula-based capacity estimate depends on assumptions (Poisson arrivals, exponential service times, no unmodeled bottleneck) that real systems only approximate, not satisfy exactly — a real system might have a bottleneck the simple M/M/1 model doesn't capture at all (a database connection pool limit, a downstream API rate limit, a specific slow code path only triggered under certain conditions) that no amount of queueing-formula reasoning alone would reveal. A load test is what confirms whether your theoretical estimate actually holds under real, measured conditions, or whether reality diverges from the model in a way that matters.

**Load testing methodology: avoiding performance track Lesson 5's coordinated omission trap specifically.** A load test that issues requests in a closed loop (wait for each response before sending the next) will systematically under-report the exact tail-latency degradation you're trying to detect as capacity is approached — an open-loop load generator (issuing requests at a fixed target rate regardless of prior completion, per performance track Lesson 5's exact point) is essential for a capacity-validating load test specifically, since the entire point is observing what happens as arrival rate approaches or exceeds the system's actual service capacity, which a closed-loop generator cannot accurately represent.

**Finding the actual bottleneck: increasing load until something breaks, and diagnosing what.** Rather than testing only at your theoretically-estimated capacity, a genuine capacity validation increases load progressively (ramping up request rate) until throughput stops increasing proportionally or latency (specifically p99, per performance track Lesson 5) degrades sharply — the point where this happens is your system's *actual* measured capacity, and the specific resource that's saturated at that point (CPU, per performance track Lesson 6's system-level tools; database connections; a downstream API's own rate limit) is your actual bottleneck, which may or may not match what your theoretical model assumed was the limiting factor.

**Why the gap between theoretical and measured capacity, when found, is itself valuable information.** If your load test reveals the actual bottleneck is a database connection pool limit at a load level well below your queueing-formula's theoretical CPU-bound estimate, that's a genuinely important finding — it means your capacity planning needs to account for that specific, previously unmodeled constraint, and it tells you exactly what to fix (increase the pool size, or address whatever's causing connections to be held longer than necessary) to actually reach your originally-estimated theoretical capacity.

## Attempt

1. Revisit your system-engineering Lesson 2 capacity-planning exercise (or redo it for a real, current service) and restate your theoretical capacity estimate (the `μ` you derived, and the resulting predicted `W` at your estimated peak `λ`).

2. Build (or use an existing tool, e.g. `vegeta`, `k6`, or a custom script) an open-loop load generator issuing requests at a controlled, fixed rate against your real or test service, explicitly avoiding performance track Lesson 5's coordinated-omission trap.

3. Run a progressive load test: start well below your theoretical capacity estimate, and increase the request rate in steps, measuring actual throughput and p50/p99 latency (performance track Lesson 5's methodology) at each step, until you observe the specific point where p99 latency degrades sharply or throughput stops scaling with increased request rate.

4. At the point of degradation identified in step 3, use performance track Lesson 6's system-level tools (`top`, `vmstat`, `iostat`) and/or application-level profiling (performance track Lesson 2) to identify the actual bottleneck resource — CPU, memory, disk I/O, database connections, or something else — and compare this measured, actual capacity and bottleneck against your step 1 theoretical prediction.

## Verify

Present your progressive load test results (throughput and p50/p99 latency at each tested load level) showing the actual degradation point, your system-level diagnosis identifying the real bottleneck resource, and an explicit comparison between your theoretical capacity estimate (step 1) and your actual measured capacity (step 3-4) — report the gap, whichever direction it goes, honestly.

## Failure drill

Run the exact same load test using a naive, *closed-loop* load generator instead of your proper open-loop one, at a load level you already know (from your correct, open-loop test) causes real p99 degradation. Compare the closed-loop test's reported p99 against your open-loop test's actual measured p99 at the same target load level, and report the discrepancy — directly, concretely reproducing performance track Lesson 5's coordinated omission finding, but now in the specific context of validating a capacity plan, where under-reporting tail latency could lead you to incorrectly conclude your system has more headroom than it actually does, a genuinely costly mistake if it leads to under-provisioning based on falsely reassuring closed-loop test results.

## Transfer

If TARDOC has never been load-tested against its actual, current infrastructure (a single Contabo VPS, per your project history), describe what a realistic load-testing plan would look like given system-engineering Lesson 2's theoretical capacity estimate for TARDOC specifically (using your own numbers from that lesson's transfer task, if you completed it) — what request rate you'd start ramping from, what you'd expect the actual bottleneck to be given the VPS's likely resource constraints, and what you'd want to measure to confirm or correct your theoretical estimate before actually needing that capacity in production.

## Done when

You've built and run a genuine open-loop load test avoiding coordinated omission, you've found your system's actual measured capacity and the specific real bottleneck resource limiting it, you've honestly compared this against your theoretical queueing-formula estimate and reported the gap in either direction, and you've directly demonstrated — via the failure drill — how a naive closed-loop load test would have under-reported the real tail-latency risk at the same load level.
