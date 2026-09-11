# Lesson 5: Latency Distributions and Tail Behavior

## Objective

Compute percentiles correctly from real latency data, and connect the math-for-engineering queueing lesson's utilization-driven blowup directly to why tail latency (p99) degrades disproportionately compared to average latency as a system approaches saturation.

## Prerequisites

Statistics track Lesson 1 (descriptive statistics — this lesson is a direct, applied extension of that lesson's percentile discussion), math-for-engineering queueing lesson (the mechanism behind why tails specifically get worse under load).

## Learn

**Why percentiles, not the mean, are the right way to characterize latency — restated with the specific mechanism now available.** Statistics track Lesson 1 established that skewed distributions (like latency) make the mean misleading. This lesson adds the "why" specific to latency: a request's latency is influenced by whether it happens to arrive during a queueing delay (math-for-engineering's `W = 1/(μ-λ)`), and queueing delay is *not* uniform across requests — most requests, arriving when the system has spare capacity, experience minimal queueing; a minority, arriving during a transient burst or when the system is momentarily near saturation, experience disproportionately larger delays. This is precisely why latency distributions have a long right tail, and why p99 (capturing that tail) tells you something p50 (the typical case) fundamentally cannot.

**Computing percentiles correctly, a genuinely easy detail to get wrong.** Given N sorted latency measurements, the p99 is (informally) "the value below which 99% of measurements fall" — but the exact interpolation method for a percentile that doesn't land exactly on an index (e.g. the 99th percentile of 1000 sorted values is somewhere between the 989th and 990th value) varies by convention (nearest-rank, linear interpolation, and others), and different tools/libraries can report meaningfully different numbers for the same raw data if they use different conventions — worth knowing this exists so you don't over-interpret small differences between two percentile calculations as a real change when they might just reflect different calculation methods.

**Why tail latency specifically worsens as utilization increases — connecting directly to the queueing formula.** As `ρ = λ/μ` approaches 1, the queueing lesson's `W = 1/(μ-λ)` grows without bound — but this growth doesn't affect every request equally. A system at moderate utilization (say `ρ=0.5`) has most requests experiencing near-zero queueing delay, with p50 staying low even as `ρ` climbs somewhat — but the tail (p99) starts degrading much earlier and much faster than the median does, because the tail specifically captures the requests unlucky enough to arrive during a transient burst, and bursts (even under an average load well below capacity) become more likely to cause a real backlog as the system's average headroom shrinks. This is why p99 latency is often the first, earliest warning sign of a system approaching saturation — well before the *average* latency shows an obviously alarming trend.

**Coordinated omission: a genuine, subtle measurement bias worth knowing about.** If a load-testing tool waits for one request to complete before issuing the next (a "closed" load-generation model), it systematically under-samples the slow tail — since a slow response delays when the *next* measurement can even begin, a naive closed-loop load generator effectively hides exactly the slow requests you most need to measure. Correctly measuring tail latency under load typically requires an "open" load model (issuing requests at a fixed target rate regardless of whether previous requests have completed yet) — a real, important, and genuinely non-obvious pitfall in load-testing methodology.

## Attempt

1. Generate a synthetic latency dataset that mimics a realistic long-tailed distribution — e.g. mostly small values (10-20ms) with an occasional large value (200-500ms), similar to statistics track Lesson 1's synthetic latency example. Compute p50, p95, and p99 from this data using at least one standard method (nearest-rank is simplest to implement correctly by hand), and report all three alongside the mean, confirming the mean sits somewhere between p50 and p95/p99, pulled upward by the tail.

2. Implement a simple M/M/1-style queueing simulator (reusing or adapting your math-for-engineering queueing lesson's simulator, if you built one) and run it at several different utilization levels (`ρ = 0.5, 0.7, 0.9, 0.95`), recording the full distribution of individual request wait times at each level (not just the average `W`). Compute p50 and p99 at each `ρ` level and report both, confirming p99 degrades much more sharply than p50 as `ρ` increases — direct, quantitative evidence of the tail-vs-median divergence from Learn.

3. Demonstrate coordinated omission directly: implement a naive closed-loop load generator (issue a request, wait for its response, then issue the next) against a service with occasional artificial slow responses, and compute the measured p99 from this closed-loop data. Then implement an open-loop load generator (issue requests at a fixed rate regardless of prior completion) against the same service, and compute p99 from *this* data. Compare the two p99 values and report the difference — the closed-loop version should under-report tail latency compared to the open-loop version.

4. Using two different percentile-calculation conventions (nearest-rank vs. linear interpolation — implement both, or use two different tools/libraries if convenient) on the same dataset from step 1, compare the resulting p99 values and report whether/how much they differ — confirming for yourself whether this specific dataset happens to be sensitive to calculation-method choice.

## Verify

Present your step 1 percentile table (mean, p50, p95, p99), your step 2 utilization-vs-percentile table showing the disproportionate p99 degradation, and your step 3 closed-loop vs. open-loop p99 comparison with actual numbers demonstrating coordinated omission's under-reporting effect.

## Failure drill

Take your step 3 closed-loop load generator's under-reported p99 and use it (incorrectly) to conclude your service's tail latency is acceptable, when your open-loop measurement from the same test shows it's actually considerably worse. Explain, using your own concrete numbers, why relying on a closed-loop load-testing tool (a common default in many simple benchmarking setups) could lead you to genuinely believe a system's tail latency is fine when real users — who don't wait politely for one response before making their next request, and whose requests arrive independent of your system's current backlog — would actually experience the worse, open-loop-measured behavior.

## Transfer

If TARDOC's or Mahall's load testing (if any currently exists) uses a simple sequential request pattern (implicitly closed-loop), describe what this lesson's coordinated-omission finding suggests about whether your current measured latency numbers might be under-representing real tail latency, and what changing to an open-loop load-generation approach would require.

## Done when

You've correctly computed p50/p95/p99 from real data and understand why they diverge from the mean for skewed distributions, you've directly measured — via your own queueing simulation — that p99 degrades disproportionately faster than p50 as utilization increases, and you've directly demonstrated coordinated omission's under-reporting effect by comparing real closed-loop versus open-loop measured p99 values on the same underlying service.
