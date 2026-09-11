# Lesson 6: Scaling and Backpressure

## Objective

Understand the difference between vertical and horizontal scaling, and implement backpressure — a system's ability to signal "slow down" to its own callers rather than silently accepting more work than it can handle and degrading uncontrollably.

## Prerequisites

Math-for-engineering queueing lesson (this lesson's entire premise — a system approaching saturation — is that lesson's `ρ→1` scenario, now addressed with a specific mitigation), OS Lesson 6 (pipes and backpressure — this lesson generalizes that exact mechanism beyond a single OS pipe to whole-system design).

## Learn

**Vertical vs. horizontal scaling, and their different limits.** Vertical scaling means making a single instance more powerful (more CPU, more RAM, a faster disk) — simple to reason about (no new distributed-systems complexity), but has a hard ceiling (there's a maximum machine size available, and cost often grows faster than linearly as you approach the largest available instances) and provides no redundancy (one bigger machine is still a single point of failure). Horizontal scaling means running more instances and distributing load across them (networking Lesson 6's load balancing) — no inherent ceiling on total capacity, and provides redundancy as a side effect, but introduces real complexity: state must either be externalized (a shared database/cache, not in-process memory) or carefully partitioned (distributed Lesson 6's sharding), and every one of the distributed-systems track's failure-mode considerations now applies.

**Why unconstrained systems degrade badly, not gracefully, under overload.** Without any backpressure mechanism, a system facing more incoming work than it can process will typically just accept everything anyway — filling internal queues, buffers, or goroutine pools unboundedly — until it exhausts memory or some other hard resource limit, at which point it fails abruptly and often catastrophically (an out-of-memory crash, or every request becoming so slow that effectively nothing completes), rather than degrading in a controlled, recoverable way. This is a direct, practical consequence of the queueing lesson's `ρ→1` blowup: accepting unlimited work when `ρ` is already at or above 1 doesn't just mean slower service, it means unboundedly growing queues, since work is arriving faster than it can ever be drained.

**Backpressure: making overload an explicit, handleable signal instead of silent degradation.** Rather than accepting unlimited incoming work, a system with backpressure has some mechanism to say "I'm at capacity, don't send more right now" — this could be a bounded queue that rejects new work once full (rather than growing unboundedly), a rate limiter (API engineering Lesson 10) actively shedding excess load, or a load balancer routing away from instances reporting themselves as near-saturated. The key design shift: overload becomes a condition the system actively detects and responds to (by rejecting, queueing with a bound, or signaling upstream to slow down), rather than a condition that silently accumulates until something breaks.

**Why bounded queues specifically matter, connecting to OS Lesson 6.** A bounded queue (as opposed to an unbounded one) is precisely what makes backpressure possible at all — an unbounded queue can always accept one more item, deferring the overload problem rather than surfacing it, right up until memory itself is exhausted. This is the exact same principle as OS Lesson 6's pipe buffer: a fixed-size buffer forces the producer to eventually block (or, at the application level, to be told "no" and decide what to do about it) rather than allowing unlimited buffering that defers a resource problem instead of solving it.

## Attempt

1. Implement a simple work-processing service with an *unbounded* internal queue (e.g. an unbuffered channel that a goroutine keeps appending to, or a slice that grows without limit) accepting incoming "jobs" faster than a single worker can process them. Run it under sustained overload (submission rate deliberately higher than processing rate) and measure memory usage growth over time, confirming it grows without bound as the unbounded queue accumulates unprocessed work.

2. Implement the same service with a *bounded* queue (e.g. a buffered Go channel with a fixed capacity) and explicit rejection logic: when the queue is full, new submissions are rejected immediately (e.g. with an equivalent of API engineering's 429 status, if this is an HTTP-facing service) rather than accepted and queued indefinitely. Rerun the same sustained-overload test and confirm memory usage now stays bounded, with a measurable rejection rate instead of unbounded accumulation.

3. Implement a basic form of backpressure signaling to an upstream caller: have your bounded-queue service expose its current queue depth (or a simple busy/not-busy signal) via a status endpoint, and write a caller that checks this signal before submitting new work, voluntarily slowing its own submission rate when the callee reports itself near capacity — rather than the callee's rejection (step 2) being the only mechanism; this is the caller proactively cooperating, closer to how well-designed distributed systems are meant to behave.

4. Compare vertical and horizontal scaling for your bounded-queue service directly: first, increase the single worker's processing capacity (simulate this by reducing the artificial per-job processing delay, standing in for "a faster machine") and measure the new sustainable throughput before the queue starts filling under your test load; then, instead, run multiple worker instances behind a simple round-robin dispatcher (networking Lesson 6's load-balancing concept, simplified) and measure the new sustainable throughput that way — compare the two approaches' resulting capacity for the same relative "cost" (e.g. 2x processing speed on one worker vs. 2 workers at original speed).

## Verify

For steps 1-2, report actual measured memory usage over time for both the unbounded and bounded queue versions under identical sustained overload, showing the concrete difference. For step 4, report actual measured sustainable throughput for both the vertical (faster single worker) and horizontal (multiple workers) scaling approaches.

## Failure drill

Take your step 2 bounded-queue-with-rejection service and deliberately configure the queue capacity far too small (e.g. capacity of 1, when your actual expected burst size is much larger) relative to normal, healthy traffic patterns (not just overload conditions). Confirm your service now rejects a significant fraction of requests even under what should be entirely manageable, non-overload load — directly demonstrating that backpressure mechanisms themselves need correctly-tuned parameters (echoing Lesson 5's timeout-tuning lesson: too aggressive a bound causes unnecessary rejection, just as too short a timeout causes unnecessary failure) and aren't a free correctness improvement regardless of configuration — they trade one failure mode (unbounded resource growth) for a different one (premature rejection) if not sized appropriately for actual expected traffic patterns.

## Transfer

If TARDOC's Celery task queue (processing transcription jobs, per your project history) currently has any bound on queue depth or in-flight task count, describe what you know or can infer about its current configuration, and reason about whether it's more likely to currently exhibit the unbounded-queue failure mode (step 1's memory-growth risk) or a well-tuned bounded-backpressure behavior (step 2's controlled rejection) under a sustained burst of transcription requests larger than your workers can keep up with in real time.

## Done when

You've directly measured unbounded memory growth under sustained overload with no backpressure, then directly measured bounded, controlled behavior with backpressure implemented, using the same overload test for both, you've implemented a working proactive backpressure signal a caller actually uses to self-regulate, and you've demonstrated — via the failure drill — that backpressure parameters themselves require correct tuning, with an undersized bound causing its own distinct failure mode under entirely normal traffic.
