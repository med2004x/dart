# Capstone 3: Distributed System

## Objective

Build a genuinely replicated, fault-tolerant service with real failure injection — the third terminal capstone, extending distributed-systems Lesson 8's own capstone to a larger, more deliberately hostile testing regime and requiring an explicit, written statement of exactly what consistency guarantees the system actually provides.

## Prerequisites

Distributed-systems track (all 8 lessons), networking track (for the RPC/partition-simulation tooling this capstone uses), system-engineering (reliability patterns — circuit breakers, graceful degradation, applicable to how clients interact with your distributed system).

## Learn

There is no new material. This capstone's specific, additional demand beyond distributed-systems Lesson 8: **a precise, written consistency-guarantee document** — not "it's fault-tolerant" as a vague claim, but a specific statement of exactly what a client can and cannot rely on (does a successful write guarantee it's readable by a subsequent read from any node? Only from the same node? Only after some bound? What happens to an in-flight write during a leader failover?) — and then *testing that document against your actual implementation*, treating any mismatch between documented and actual behavior as a real bug, exactly as system-engineering Lesson 15's architecture-review discipline would demand of any design claim.

**Why writing the guarantee down precisely is harder, and more valuable, than it sounds.** It's genuinely easy to build a replicated system that "seems to work" under casual testing while never having precisely specified what it actually guarantees under adversarial conditions — this capstone specifically forces the precision distributed-systems Lesson 4's replication lesson and Lesson 7's consistency discussion both built toward, applied to a system substantial enough that the precision actually matters.

## Attempt

Build a replicated key-value service (or extend distributed-systems Lesson 8's capstone, if you completed it, to this larger scope) with:

1. **Real consensus-based leader election and log replication** (distributed-systems Lesson 5, ideally via the MIT 6.5840 Raft implementation if you completed those labs, or your own from that lesson).
2. **A client interface with idempotent writes** (distributed-systems Lessons 2 and 7's idempotency pattern, API engineering Lesson 7's idempotency-key mechanism).
3. **A written consistency-guarantee document**, produced *before* your final round of testing (not written retroactively to match whatever you observe) — stating explicitly: what a client can rely on regarding read-after-write consistency, what happens to reads during a leader election, and what the system's behavior is during a network partition (distributed-systems Lesson 1's failure model, Lesson 5's majority-partition behavior).

## Verify

Run an extensive, deliberately adversarial failure-injection test suite — beyond distributed-systems Lesson 8's three minimum scenarios (minority failure, leader failure, partition), include at minimum:

1. **Rolling restarts**: restart nodes one at a time in sequence while under continuous client load, confirming the cluster remains available (majority intact) throughout, and correctly recovers each restarted node's state.
2. **Repeated partition/heal cycles**: partition and heal the cluster multiple times in succession (not just once), confirming correct behavior and correct catch-up after *each* cycle, not just the first.
3. **Client retries during failover**: send idempotent client requests that specifically straddle a leader election (the request is in flight when the leader fails), and confirm — using your idempotency mechanism — that the request is neither lost nor duplicated despite the mid-flight leadership change.

For each scenario, test your written consistency-guarantee document directly: does actual observed behavior match what you documented? Report any mismatch found as a genuine bug (in either the implementation or the document — determine which is actually wrong and fix that one).

## Failure drill

Deliberately construct the specific scenario your consistency document is *weakest* on, or most ambiguous about — for most simple designs, this is often "what exactly can a client observe about a write that was in-flight during a leader failover, before the client itself learns the failover happened." Test this scenario directly and rigorously, and if your document's claim doesn't precisely match observed behavior (a very realistic outcome — this is a genuinely hard case), revise the document to be accurate rather than aspirational. Explain why an accurate, even if less impressive-sounding, consistency guarantee is more valuable to a real client of your system than an inaccurate but stronger-sounding claim — a client that builds on an incorrect guarantee will eventually hit the gap in production, at a worse time than during your own deliberate testing.

## Transfer

Compare your capstone's actual, tested consistency guarantees against a real, production distributed system's documented guarantees (etcd's or CockroachDB's consistency documentation are reasonable references) — identify one specific place where the real system makes a stronger or more precisely-scoped guarantee than yours, and reason about what additional engineering (more extensive Raft edge-case handling, additional protocol layers) that stronger guarantee likely required.

## Done when

You've built a real, Raft-based (or equivalent) replicated service, written a precise consistency-guarantee document before your final testing round, run an extensive adversarial failure-injection suite beyond the minimum three scenarios, found and resolved at least one real mismatch between documented and actual behavior (via the failure drill's targeted weak-point testing), and honestly compared your guarantees against a real production system's documented ones.
