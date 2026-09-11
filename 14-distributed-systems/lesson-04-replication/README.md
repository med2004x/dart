# Lesson 4: Replication and State Machines

## Objective

Understand replicated state machines as the general pattern behind most fault-tolerant distributed systems, and directly experience the core consistency tradeoff: keeping replicas identical costs coordination, and skipping that coordination risks replicas disagreeing.

## Prerequisites

Lesson 3 (logical clocks — useful for reasoning about operation ordering across replicas), database-internals Lesson 6 (WAL — a replicated state machine's log is conceptually the same "durable, ordered record of operations" idea, now shared across multiple machines instead of surviving one machine's own crash).

## Learn

**The state machine replication pattern.** If every replica starts in the same initial state and applies the exact same sequence of deterministic operations in the exact same order, every replica ends up in the exact same state — this is the entire, deceptively simple idea underlying most database replication, distributed consensus systems (Lesson 5's Raft), and even blockchain-style systems at a conceptual level. The hard part isn't the idea itself, it's guaranteeing "the exact same sequence... in the exact same order" actually holds across machines that can fail, be partitioned (Lesson 1), and communicate over an unreliable, asynchronous network.

**Determinism is a strict requirement, not a nice-to-have.** If an "operation" isn't deterministic (e.g. it depends on the local wall-clock time, or a local random number generator seeded differently per machine), replicas that supposedly received the identical operation sequence can still diverge into different states — this is a genuinely common, subtle bug class in real replicated systems, and it's precisely why systems built on this pattern go out of their way to make operations deterministic (e.g. passing a fixed timestamp *as part of* the operation, generated once by whichever replica is leading, rather than each replica reading its own local clock when applying the operation).

**Leader-based replication (the common practical pattern).** Rather than every replica independently deciding operation order (which the CAP-theorem-adjacent tradeoffs discussed below make genuinely hard to do consistently), one replica is designated **leader** and is solely responsible for deciding the order of new operations; other replicas (**followers**) simply replicate the leader's log in order. This concentrates the ordering decision in one place, avoiding the harder distributed-agreement problem in the common case — until the leader itself fails, at which point a new leader must be elected (exactly what Raft, Lesson 5, provides a rigorous protocol for).

**Synchronous vs. asynchronous replication, and the tradeoff it represents.** Synchronous replication waits for a write to be acknowledged by follower replicas (some or all of them) before confirming success to the client — stronger consistency (a confirmed write really is on multiple machines already) at the cost of higher latency (waiting for the slowest required follower) and reduced availability if a required follower is unreachable. Asynchronous replication acknowledges the client immediately after the leader's own local write, propagating to followers afterward, in the background — lower latency and higher availability, but a real risk: if the leader crashes before a given write has actually propagated to any follower, that write can be lost entirely, even though the client was already told it succeeded.

## Attempt

1. Implement a minimal replicated state machine: a "leader" process maintaining a simple key-value store, with an append-only operation log (echoing database-internals Lesson 6's WAL), and one or more "follower" processes that receive the leader's log entries (over RPC, Lesson 2) and apply them in order to their own local copy of the state.

2. Test basic replication correctness: perform a sequence of writes against the leader, confirm each follower's local state exactly matches the leader's after all operations have propagated and been applied.

3. Simulate a follower crash and recovery: stop a follower process partway through a sequence of leader operations (so it misses some), then restart it and implement a simple catch-up mechanism (the follower requests all log entries after the last one it successfully applied, and replays them in order) — confirm the recovered follower's state correctly converges to match the leader's, including the entries it missed while down.

4. Implement both synchronous and asynchronous replication modes for the leader's writes, and measure the actual latency difference for a batch of writes under each mode (with at least one follower introducing artificial delay via a `sleep()` in its own operation-application logic, to make the synchronous mode's wait genuinely visible in your timing measurement rather than negligible).

## Verify

For step 2, report the exact final state (all key-value pairs) from the leader and from each follower after your test sequence, confirming byte-for-byte agreement. For step 4, report the actual measured latency numbers for synchronous versus asynchronous mode, with your artificial follower delay clearly identified as the cause of the difference.

## Failure drill

Using your asynchronous replication mode from step 4, perform a write against the leader, and — before it has propagated to any follower — simulate the leader crashing (kill the leader process or discard its state entirely) while at least one follower is still consistent with the leader's state *before* that specific write. Confirm the write is now permanently lost — no follower has it, and the leader (if restarted from scratch, with no separately-persisted log of its own) has no record of it either. Explain why this is not a bug in your implementation but the precise, expected cost of asynchronous replication's tradeoff — the client was told the write succeeded (since asynchronous mode acknowledges before propagation), yet the write is now gone, which is exactly the kind of "confirmed but actually lost" scenario synchronous replication (at the cost of the latency you measured in step 4) exists to prevent.

## Transfer

If TARDOC's PostgreSQL database uses any form of replication (even just for backup/disaster-recovery purposes, distinct from a full multi-node active setup), describe, using this lesson's synchronous/asynchronous distinction, what PostgreSQL's `synchronous_commit` and replication-related settings (referenced in database-internals Lesson 6's transfer task) are actually controlling in terms of this exact tradeoff, and state what business consequence (per your own project context — clinic billing data, specifically) would follow from choosing asynchronous replication and then experiencing exactly the failure-drill scenario above: a confirmed transaction whose data existed only on a leader that then failed before replicating.

## Done when

Your replicated state machine correctly converges follower state to match the leader across a sequence of operations, you've demonstrated a follower correctly catching up after a simulated crash and restart, and you've directly reproduced — not just read about — the specific data-loss scenario asynchronous replication risks, with your own measured latency numbers showing exactly what that risk was traded for.
