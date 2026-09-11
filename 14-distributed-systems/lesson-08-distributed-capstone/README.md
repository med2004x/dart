# Lesson 8: Capstone — Fault-Tolerant Service

## Objective

Integrate RPC, replication, failure handling, and observability into one working fault-tolerant service, and directly measure its behavior under deliberately injected node failures — the distributed-systems track's equivalent of the integration capstones in earlier tracks.

## Prerequisites

Lessons 1-7, completed. No new theory — this is deliberately an integration exercise, and distributed systems in particular tend to reveal integration gaps that no individual component's isolated testing would surface, since most real distributed bugs emerge specifically from the interaction between components under failure conditions, not from any single component in isolation.

## Learn

There is no new material here. MIT 6.5840 (referenced throughout this track) extends its Raft labs into a full fault-tolerant, sharded key-value store as its own capstone sequence — using that directly is a strong option if you want the most rigorous version of this exercise, with a battle-tested test harness specifically designed to inject exactly the kind of failures this lesson asks you to handle.

The specific engineering discipline this capstone tests, beyond correctness: **can you observe and reason about your own system's behavior under failure**, not just verify it eventually recovers. A fault-tolerant system that "self-heals" but gives you no visibility into what actually happened during the failure window is much harder to operate confidently in production than one you can actually watch and understand while it's degraded — this is why the requirements below explicitly include observability, not just fault tolerance itself.

## Attempt

Build a small key-value service with the following integrated properties, using your own implementations from Lessons 2-7 (or the MIT 6.5840 Raft-based KV store, if you're using that path):

1. **RPC-based client interface** (Lesson 2): clients interact with the service via RPC calls (`Get`, `Put`), with proper timeouts and request-ID-based idempotency for `Put` operations specifically (since, per Lesson 2's own example, a naive retry of a write operation risks duplication).

2. **Replicated, fault-tolerant backend** (Lessons 4-5): the service's state is replicated across at least 3 nodes using either your own Raft implementation from Lesson 5, or (if you completed the MIT labs) the course's Raft implementation — the service should continue serving requests correctly as long as a majority of nodes remain alive and connected, per Lesson 5's majority-based guarantees.

3. **At-least-once-safe request handling** (Lesson 7): if a client's request times out and it retries, the service should not produce a duplicate effect for a `Put` operation, directly reusing Lesson 2/7's idempotency pattern.

4. **Observability**: the service should log, at minimum, every leader election (including term numbers), every request received and whether it succeeded/failed/was deduplicated, and enough information to reconstruct, after the fact, exactly what happened during any failure-injection scenario you run — this is what makes the following verification step actually possible to do rigorously rather than just "it seemed to work."

## Verify

Run at least three distinct failure-injection scenarios against your running service, and for each, report both the client-observed behavior (did requests succeed, fail, or hang, and for how long) and the server-side log evidence explaining why:

- **Minority node failure**: kill one node out of a 3 (or 5) node cluster while continuously sending client requests. Confirm the service continues serving correctly (majority still available, per Lesson 5) with at most a brief interruption during any in-progress leader election, and report the actual measured interruption duration from your logs.
- **Leader failure specifically**: identify and kill the current leader (not an arbitrary node) while sending requests, and confirm a new leader is elected and service resumes, again reporting the actual measured interruption duration.
- **Network partition**: using `tc netem` or equivalent (networking Lesson 7), partition the cluster into a majority and minority group while sending requests to both groups' nodes, and confirm the majority group continues serving correctly while the minority group correctly refuses to serve (per Lesson 5's split-brain prevention), then heal the partition and confirm the minority nodes correctly catch up.

## Failure drill

Take your minority-node-failure scenario and, instead of a clean process kill, simulate a node that's merely *slow* rather than fully crashed (e.g. artificially delay its message processing significantly, per distributed Lesson 1's delay-vs-crash distinction) — confirm your system's behavior under this condition, and specifically check whether a slow-but-not-dead node causes any incorrect behavior (e.g. does it ever incorrectly get treated as failed and then cause a conflicting election, or does your timeout tuning handle this correctly). Report what you observed, and if you find a genuine bug or edge case here (a real possibility, since delay-vs-crash ambiguity, per Lesson 1, is exactly the hardest case), document it explicitly rather than papering over it — an honest "this is a known limitation of my current timeout configuration" is a more valuable capstone outcome than an unexamined claim of full correctness.

## Transfer

Compare your capstone's specific fault-tolerance guarantees (what failure scenarios it correctly handles, and any limitations you found in the failure drill) to what a production system like etcd or CockroachDB actually provides and guarantees in their own documentation — pick one specific area where your simplified version's guarantees are weaker or less tested than the real system's, and explain briefly what additional engineering (more extensive failure-injection testing, more careful timeout tuning, handling more failure-mode combinations) would be needed to close that gap.

## Done when

Your integrated service correctly handles all three verification scenarios (minority node failure, leader failure, network partition) with measured, logged evidence of correct behavior and bounded interruption time, your idempotency mechanism prevents duplicate writes under retry even during a leader election in progress, and you've honestly documented at least one limitation or edge case discovered via the failure drill's slow-node scenario, rather than claiming unqualified correctness.
