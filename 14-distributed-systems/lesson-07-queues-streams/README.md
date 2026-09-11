# Lesson 7: Queues and Stream Processing

## Objective

Implement at-least-once message delivery and understand precisely why "at-least-once" (not "exactly-once," which is far harder to achieve genuinely) is the realistic default guarantee — and design idempotent processing to make that guarantee safe to work with in practice, directly extending distributed Lesson 1 and Lesson 2's duplicate-handling groundwork to the message-queue setting specifically.

## Prerequisites

Distributed Lesson 1 (failure models — the ambiguity around "did it actually get processed" reappears here in a new form), distributed Lesson 2 (idempotency via request IDs — this lesson applies the same core idea to queue consumers instead of RPC clients).

## Learn

**Why "exactly-once" delivery is much harder than it sounds, and usually not what's actually implemented.** For a message to be delivered exactly once, the system would need to atomically: receive/process the message, and durably record that it was processed, as a single indivisible operation — but per distributed Lesson 1's core ambiguity, a crash between "processed" and "durably recorded" is always possible, and there's no way to make these two things atomic across a genuine machine/process boundary without essentially building a mini-consensus protocol around every single message. Most real messaging systems (Kafka, RabbitMQ, SQS) instead offer **at-least-once** delivery by default: a message might be delivered and processed more than once (if the consumer crashes or fails to acknowledge after processing but before that acknowledgment is durably recorded), but it will never be silently dropped — which is a strictly easier, more achievable guarantee, and the one that shifts the "did this actually happen exactly once" burden onto the consumer's own processing logic.

**At-least-once, mechanically.** A consumer receives a message, processes it, then explicitly acknowledges (ACKs) it back to the queue system. If the consumer crashes (or the ACK is lost, per distributed Lesson 1's omission-failure category) before the ACK is durably recorded, the queue system will eventually redeliver the same message to some consumer (possibly the same one after it restarts, possibly a different one) — this redelivery is the "at least" in at-least-once, and it's a deliberate design choice favoring "never lose a message" over "never duplicate a message," on the reasoning that duplicates can be handled by the consumer (this lesson's core topic) while lost messages generally cannot be recovered at all.

**Idempotent processing: making at-least-once delivery safe.** Exactly the same idea as distributed Lesson 2's RPC request-ID mechanism, applied to queue consumers: track which message IDs have already been fully processed (in a durable store, e.g. a database table, not just in-memory state that a crash would lose), and when a redelivered message arrives, check that record before processing — if already processed, skip (or return the previously-computed result) rather than executing the operation's side effects a second time. Some operations are *naturally* idempotent regardless of this mechanism (e.g. "set user's email to X" — running it twice produces the same final state as running it once) — for these, explicit deduplication tracking may be unnecessary. Others are fundamentally not naturally idempotent (e.g. "charge $10," "send an email," "append to a list," per distributed Lesson 1's original example) and genuinely require the explicit tracking mechanism to be safe under at-least-once delivery.

## Attempt

1. Implement a minimal at-least-once queue: a producer appends messages to a durable log (a file, or reuse database-internals Lesson 6's WAL-style append-only log concept), and a consumer reads messages, processes them, and only advances its "last processed" checkpoint (also durably stored) *after* successfully completing processing — not before. Confirm that if you kill the consumer process mid-processing (before it advances the checkpoint) and restart it, it correctly redelivers/reprocesses the message it was working on when killed, rather than skipping it.

2. Implement a naturally-idempotent operation as your consumer's processing logic (e.g. "set key K's value to V" in a key-value store) and confirm that redelivery/reprocessing (per step 1's crash-and-restart test) produces no observable problem — the final state is correct either way, since reapplying the same "set" operation is harmless by the operation's own nature.

3. Implement a *non*-idempotent operation instead (e.g. "append value V to a list associated with key K," directly reusing distributed Lesson 1's example) and confirm, without any additional deduplication mechanism, that the crash-and-restart redelivery scenario from step 1 now produces a real, observable duplicate — the value appears twice in the list, even though logically it should only have been added once.

4. Add explicit deduplication to your step 3 consumer: track processed message IDs in a durable store (a simple table/file mapping message ID to "already processed"), check this record before processing each message, and skip reprocessing (or return a cached result) for already-seen IDs. Rerun the exact crash-and-restart scenario from step 3 and confirm the duplicate is now prevented — the list contains the value exactly once, matching step 2's naturally-idempotent case's correctness, but now achieved explicitly for an operation that isn't naturally idempotent.

## Verify

For step 3 and step 4, report the actual final list contents after the crash-and-restart test in both cases — step 3 should show the value duplicated, step 4 should show it appearing exactly once, giving you direct, concrete before/after evidence of the deduplication mechanism's effect.

## Failure drill

In your step 4 deduplication mechanism, deliberately introduce the same ordering bug distributed database-internals Lesson 6's WAL failure drill covered: record the "already processed" marker *before* actually applying the operation's effect (reversed from the correct order), then simulate a crash occurring between recording the marker and actually applying the effect. Confirm that after restart, the message is now considered "already processed" (so it will never be redelivered or reprocessed) even though its actual effect was never applied — a silently *lost* operation, which is arguably worse than the duplicate this lesson has otherwise been preventing, since a duplicate is at least visible/detectable while a silently skipped operation may not be. Explain why the correct ordering (apply the effect, then durably record it as processed — mirroring database-internals Lesson 6's "write-ahead" ordering principle, just applied in the opposite temporal direction here: record only after the effect is safely applied, not before) is what avoids this specific new failure mode.

## Transfer

If TARDOC's Celery-based task queue (mentioned throughout your project history — the hourly subscription sweep, the nightly audio retention purge) processes tasks that aren't naturally idempotent (e.g. sending a billing notification email, or incrementing a usage counter), describe, using this lesson's exact deduplication pattern, whether Celery's own task-retry mechanism could produce the same duplicate-processing risk this lesson demonstrated, and what a durable "already processed" tracking table (keyed on task ID, or some other uniquely-identifying value per logical task invocation) would need to look like to make those specific tasks safe under Celery's at-least-once-style retry behavior.

## Done when

You've directly demonstrated at-least-once delivery redelivering a message after a simulated consumer crash, you've shown the concrete difference in outcome between a naturally-idempotent and a non-idempotent operation under that same redelivery, and you've implemented and verified explicit deduplication fixing the non-idempotent case — while also demonstrating, via the failure drill, the specific new failure mode (silently lost operations) that results from getting the deduplication-marker ordering wrong.
