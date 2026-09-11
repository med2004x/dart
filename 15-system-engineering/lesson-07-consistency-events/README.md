# Lesson 7: Consistency and Event-Driven Design

## Objective

Understand eventual consistency as a deliberate design choice (not a bug to be tolerated), and design an event-driven flow where different parts of a system converge to a consistent state asynchronously rather than requiring every consumer to be synchronously up to date.

## Prerequisites

Distributed Lesson 4 (replication — the synchronous/asynchronous tradeoff from that lesson is the same core idea, now applied to application-level event propagation rather than database replication specifically), distributed Lesson 7 (queues — the delivery mechanism event-driven design typically relies on).

## Learn

**Strong consistency vs. eventual consistency, as a real design spectrum, not "correct" vs. "incorrect."** Strong consistency means every reader sees the most recent write immediately — simple to reason about, but (per distributed Lesson 4) often requires synchronous coordination, with real latency and availability costs. Eventual consistency means readers *will* eventually see a given write, but there's a window (hopefully short, but not zero) during which different parts of the system may disagree about the current state. This isn't automatically worse — it's a legitimate, often necessary tradeoff for systems where the coordination cost of strong consistency (across services, not within one database transaction) would be prohibitive, and where the business domain can actually tolerate a brief inconsistency window.

**Event-driven design: propagating state changes as discrete events rather than direct synchronous calls.** Instead of Service A directly calling Service B, C, and D synchronously whenever something happens (which, per system-engineering Lesson 1's request-boundary discussion, couples A's success to all three downstream calls succeeding promptly), A publishes an event ("order created") to a queue or event bus, and B, C, D each independently consume that event and update their own state accordingly, on their own schedule, decoupled from A's original request lifecycle entirely. This directly extends Lesson 1's boundary-drawing lesson: an event-driven architecture is what "defer to background processing" often looks like when the deferred work is fanned out to multiple independent consumers rather than a single background task.

**Why the business domain determines whether eventual consistency is actually acceptable — this is not a purely technical decision.** For some data, even a brief inconsistency window is unacceptable (e.g. an account balance check immediately before authorizing a large withdrawal) — these need strong consistency, full stop, regardless of the coordination cost. For other data, a brief propagation delay is genuinely harmless (e.g. a "last seen" timestamp updating across a search index a few seconds after the actual event, or an analytics dashboard reflecting data with some intentional lag) — these are excellent, low-risk candidates for eventual consistency's benefits (looser coupling, better independent scalability of consumers). The engineering skill this lesson builds is correctly classifying which category a given piece of data falls into, not defaulting to either extreme universally.

**Idempotent event consumers, restated because it's non-negotiable here.** Since event delivery typically follows distributed Lesson 7's at-least-once guarantee (not exactly-once, for the reasons covered there), every event consumer needs to handle potential redelivery of the same event idempotently — this isn't optional for a correctly-designed event-driven system, it's a structural requirement given the delivery guarantee the underlying queue actually provides.

## Attempt

1. For a real or plausible operation in TARDOC or Mahall (e.g. "clinic subscription activated"), identify which downstream effects genuinely need strong, synchronous consistency (the client must know they succeeded before the response) versus which are acceptable candidates for eventual consistency via events (the client doesn't need to wait, and a brief propagation delay is genuinely harmless for that specific effect) — directly reusing system-engineering Lesson 1's boundary-drawing question, now applied specifically through the consistency lens.

2. Implement an event-driven flow for the eventually-consistent effects identified in step 1: the primary operation publishes an event to a queue (reuse distributed Lesson 7's queue implementation, or a simple in-memory equivalent for this exercise) upon completion, and at least 2 independent consumers subscribe to and process that event on their own schedule, each updating their own piece of state.

3. Measure the actual propagation delay in your implementation: from the moment the event is published to the moment each consumer has finished processing it and updated its own state, under normal (unloaded) conditions. Report the actual measured delay, and reason about whether that delay would be acceptable for the specific business effect it represents, given your step 1 classification.

4. Implement idempotent processing (distributed Lesson 2/7's pattern) in at least one of your event consumers, and test it against a simulated duplicate delivery of the same event (echoing distributed Lesson 7's exact test pattern), confirming the consumer's state update happens exactly once despite the redelivery.

## Verify

Present your step 1 classification table (effect, strong-consistency-required or eventual-consistency-acceptable, justification), your step 3 measured propagation delay, and confirm your step 4 idempotency test shows the consumer's final state correctly reflects a single processing despite duplicate delivery.

## Failure drill

Take one of your "eventual consistency is acceptable" effects from step 1 and construct a scenario where a user or client observably notices the temporary inconsistency window — for example, a client that creates a resource via the strongly-consistent primary path, then immediately queries a different part of the system (one of your eventually-consistent consumers) expecting to see the effect already reflected there, and finding it briefly absent because the event hasn't propagated and been processed yet. Confirm this observably happens in your implementation (a real, reproducible race between the primary operation completing and the consumer catching up), and explain why this specific user-visible gap is the actual, concrete cost of choosing eventual consistency for this effect — and why your step 1 classification needs to have genuinely accounted for this possibility (is a user plausibly going to check this specific downstream state within the propagation-delay window in a way that matters?) rather than assuming eventual consistency's cost is purely abstract.

## Transfer

If TARDOC's clinic dashboard or any reporting view aggregates data that's updated via background processing (e.g. Celery tasks, per your project history) rather than being read directly from the primary, strongly-consistent source, describe whether users of that dashboard would ever plausibly notice the propagation delay in a way that matters (e.g. checking immediately after performing an action that should be reflected), and whether the current design's consistency model matches what you'd classify, using this lesson's framework, as appropriate for that specific data.

## Done when

You've classified real operation effects into strong-consistency-required versus eventual-consistency-acceptable with genuine justification (not a default assumption), you've implemented and measured a real event-driven flow's actual propagation delay, you've implemented and tested idempotent event consumption against simulated duplicate delivery, and you've directly reproduced and explained a user-visible consequence of the eventual-consistency window your own design introduced.
