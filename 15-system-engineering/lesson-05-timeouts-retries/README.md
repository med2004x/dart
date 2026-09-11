# Lesson 5: Timeouts and Retries

## Objective

Set timeouts and retry policies deliberately, based on the actual failure modes and cost tradeoffs involved, rather than copying default values — and understand cascading failure, the specific danger poorly-tuned timeouts/retries can cause across a system of interdependent services.

## Prerequisites

Distributed Lesson 1 (failure models — timeouts are the practical mechanism for handling that lesson's fundamental ambiguity about slow-vs-crashed), math-for-engineering queueing lesson (the `ρ→1` blowup this lesson's cascading-failure discussion directly reuses).

## Learn

**Why "just set a timeout" isn't sufficient design — the value matters, and it's a real tradeoff.** Too short a timeout: you'll frequently give up on requests that would have succeeded if you'd waited slightly longer, especially under normal, expected latency variance (not every slow response indicates a genuine problem). Too long a timeout: a caller waiting on a genuinely stuck downstream dependency ties up its own resources (a connection, a goroutine, a thread) for an excessive time, and if that caller itself has upstream callers waiting on *it*, the delay propagates — this propagation is exactly the mechanism behind cascading failure, covered below.

**Retry storms, and why "just retry on failure" is dangerous without more thought.** If a downstream service is struggling (e.g. approaching the queueing lesson's `ρ→1` saturation point) and starts responding slowly or with errors, naive retry logic across many callers can make the situation dramatically worse: every failed request gets retried, adding *more* load to an already-overloaded system — exactly the kind of positive feedback loop that can turn a temporary, recoverable slowdown into a full outage. This is a real, well-documented failure pattern in production systems, not a theoretical concern.

**Exponential backoff with jitter: the standard mitigation.** Instead of retrying immediately (worsening a retry storm) or at a fixed interval (which can cause many callers to retry in lockstep, still overwhelming a recovering service with a synchronized burst), exponential backoff increases the delay between successive retries (1s, 2s, 4s, 8s...), and **jitter** (adding a small random variation to each delay) spreads out retries from many simultaneous callers so they don't all hit the recovering service at the exact same moments — both together substantially reduce the retry storm risk compared to naive immediate or fixed-interval retries.

**Cascading failure, the pattern this lesson's tuning is meant to prevent.** Service A calls Service B, which calls Service C. If C becomes slow (per the queueing lesson, perhaps approaching saturation), B's calls to C start taking longer, tying up B's own resources (connections, threads/goroutines) waiting on C — if B has a fixed, limited pool of resources for handling A's requests, and enough of them are now tied up waiting on slow C, B itself becomes unable to serve A's requests promptly, even though B's own code has no bug — B has become collaterally overloaded purely because of C's slowness propagating upward. **Timeouts are what bound this propagation**: a well-chosen timeout on B's call to C ensures B gives up and frees its resources after a bounded time, rather than waiting indefinitely and being dragged down by C's problem.

## Attempt

1. For a real call chain in TARDOC or Mahall (or a plausible one — e.g. an API endpoint that calls a database, which itself might be under load), identify what timeout is currently configured (or would be, by a common default) at each hop, and reason about whether that value seems deliberately chosen or just a framework default.

2. Implement a small simulation: Service A calls Service B calls Service C, with C's response time controllable (simulate normal fast responses, then simulate C becoming very slow, e.g. via an artificial delay). Confirm that with no timeout at B (or an excessively long one), B's own resource usage (e.g. count of concurrently blocked goroutines/threads waiting on C) grows without bound as C slows down, directly demonstrating the cascading-failure mechanism from Learn.

3. Add a reasonable timeout at B's call to C, and rerun the same C-slowdown scenario. Confirm B now correctly bounds its own resource usage (goroutines waiting on C are capped at the timeout duration, not growing unboundedly) and can continue serving at least some of A's requests, even while C remains slow — directly demonstrating the timeout's mitigating effect using the same before/after comparison.

4. Implement exponential backoff with jitter for A's retries against B (simulating A retrying failed/timed-out calls), and demonstrate the retry-storm risk directly: with many simulated concurrent instances of A all retrying against a struggling B using naive immediate retries, measure the total request volume hitting B during B's recovery window; then repeat with your backoff-with-jitter implementation and compare the measured request volume during the same recovery window.

## Verify

For step 2-3, report the actual measured resource usage (blocked goroutine/thread count) over time for both the no-timeout and with-timeout versions under the same simulated C-slowdown, showing the concrete difference. For step 4, report the actual measured request volume hitting B during its recovery window for both naive-immediate-retry and backoff-with-jitter, showing the concrete reduction.

## Failure drill

Configure your step 3 timeout to be deliberately too short (much shorter than C's *normal*, healthy response time, not just its slowed-down time) and rerun the simulation with C responding at its normal, healthy speed. Confirm B now frequently times out on requests that would have succeeded, generating unnecessary failures and retries even though nothing is actually wrong with C — directly demonstrating the "too short" failure mode from Learn, using your own measured false-timeout rate as evidence, and connect this concretely to why timeout values need to be chosen based on actual measured normal-case latency (with reasonable margin), not picked arbitrarily small "to be safe."

## Transfer

If TARDOC's Celery task processing calls an external API (the Groq-hosted Whisper endpoint, per your project history, with multi-key rotation already implemented), describe what timeout and retry strategy would be appropriate there specifically — reasoning about the actual expected latency of that external call (transcription of audio is not instantaneous, so a timeout tuned for a typical fast API call would be inappropriately short here) and whether the existing multi-key rotation logic already provides some of this lesson's retry/backoff benefit, or whether it's solving a different problem (rate-limit avoidance across keys) that's complementary to, but not a substitute for, proper timeout/backoff tuning on each individual call.

## Done when

You've directly demonstrated cascading resource exhaustion from an unbounded downstream slowdown, then directly demonstrated a correctly-set timeout bounding that exhaustion, using the same before/after simulation, and you've directly measured the concrete reduction in request volume that exponential backoff with jitter provides over naive immediate retries during a simulated recovery window — plus demonstrated the opposite failure mode (a too-short timeout causing unnecessary failures under normal conditions) with your own measured false-timeout rate.
