# Lesson 1: Distributed Models and Failure

## Objective

Understand why "partial failure" — the defining problem of distributed systems — has no equivalent in single-machine programming, and build the vocabulary (crash, omission, partition, delay) to reason precisely about what can go wrong.

## Prerequisites

Networking Lesson 2 (TCP — network failures happen at exactly this layer), OS Lesson 1 (process crashes — a distributed system's "node failure" is this same event, just observed remotely and ambiguously).

## Learn

**Why distributed systems are qualitatively different, not just "more of the same."** In a single process, if a function call fails, you know it failed — an exception, an error return, a crash you can observe directly. Across a network, if you send a request and get no response, you face a fundamental ambiguity: did the remote server crash before processing it? Did it crash after processing it but before responding? Did the request never arrive? Did the response get lost on the way back? From the caller's perspective, all of these look identical — a timeout, no answer — but they have completely different implications for whether the operation actually happened. This ambiguity, not any specific technical mechanism, is the root of almost every hard problem this track covers.

**Failure modes, precisely distinguished:**
- **Crash failure**: a node stops entirely and stays stopped — the "clean" failure mode, and the one most textbook algorithms assume for simplicity, though real infrastructure doesn't always fail this cleanly.
- **Omission failure**: a message is sent but never arrives (or a node fails to send a response it should have) — distinct from a crash, since the node itself may be fine, just the message was lost.
- **Network partition**: a subset of nodes becomes unable to communicate with another subset, while each subset can still communicate internally — from inside either partition, the *other* partition's nodes appear to have crashed (indistinguishable, per the ambiguity above), even though they're actually alive and running.
- **Timing/delay failure**: a node or message is simply slow, not failed — arbitrarily delayed, but eventually correct. This is often the hardest to distinguish from a genuine crash within any bounded timeout, which is precisely why choosing timeout values is a real engineering tradeoff, not a formality.

**Synchronous vs. asynchronous system models.** A synchronous model assumes known bounds on message delay and processing time — if a response doesn't arrive within that bound, you can safely conclude failure. An asynchronous model makes no such assumption — a message could be arbitrarily (though finitely) delayed, meaning you can *never* safely distinguish "slow" from "crashed" using timeouts alone, only guess. Real networks are closer to asynchronous (no hard guaranteed bound, even if delays are usually small) — a foundational and genuinely inconvenient fact that shapes the guarantees any real distributed algorithm can actually provide, which is why the field so often talks in terms of "eventually" rather than "within N seconds."

**Why this determines what guarantees you can even ask for.** Given the ambiguity above, "did my write succeed" is not always answerable with certainty from the client's perspective after a timeout — this single fact is the root cause behind idempotency requirements (Lesson 2, RPC retries), the entire consensus problem (Lesson 5), and the CAP-theorem-style tradeoffs (Lesson 4) that recur throughout this track. Every later lesson's specific mechanism exists because of the fundamental ambiguity established here.

## Attempt

1. For a simple client-server request (e.g. your networking track's Lesson 4 HTTP server), enumerate every point at which a failure could occur between the client sending a request and receiving a response — the request in flight, the server crashing mid-processing, the response in flight, and the client itself. For each point, state whether the client, from a plain timeout alone, can distinguish it from any of the other points.

2. Using `tc netem` (networking Lesson 7's tool) or a similar mechanism, simulate a network partition between a client and server (block all traffic between them for a period, e.g. via a firewall rule or `tc` loss configuration set to 100%), and observe what your client-side code actually experiences — a timeout, a connection refused, or a hang, depending on your specific setup — and confirm this is observably identical to what a genuine server crash would produce from the client's perspective (test both scenarios and compare the client-side symptoms).

3. Design (on paper, no implementation required) a simple protocol for a client to append an item to a remote list, and explicitly write out: what happens if the request message is lost before reaching the server (omission), what happens if the server crashes after appending but before responding (a genuine ambiguity for the client), and what happens if the response is lost on the way back after a successful append. For each of these three cases, state what a naive "if no response, just retry" strategy would do, and whether that's safe or could cause a problem (duplicate append) — this sets up Lesson 2's idempotency discussion directly.

4. Write a short definition, in your own words, of the difference between the synchronous and asynchronous system models from Learn, and explicitly state which one better describes the real internet (or even a real datacenter network) and why — connect this to networking Lesson 2's discussion of TCP's own retransmission timeouts, which face exactly this same fundamental ambiguity at the transport-protocol level.

## Verify

For step 2, report the actual client-side symptom you observed (exception type, error message, or hang behavior) for both the simulated partition and a genuine server crash, and confirm explicitly whether they were distinguishable or identical from the client's vantage point.

## Failure drill

Take your step 3 "naive retry" design and actually reason through what happens if the client retries an append operation after *not* receiving a response, in the specific case where the server had actually already succeeded (the response was lost, not the original request) — trace explicitly that this naive approach results in the item being appended twice, a real, concrete duplicate-write bug, not a hypothetical one. Explain why this single scenario is the entire motivating case for idempotency (which Lesson 2 and Lesson 7 both address with specific mechanisms) — the retry itself isn't wrong (it's often the only reasonable response to an ambiguous timeout), but a retry without some means of detecting "this was already done" turns ambiguity into an actual, guaranteed data-correctness bug.

## Transfer

If TARDOC's Celery task queue (mentioned in your project history) ever retries a failed task, describe, using this lesson's crash/omission/partition/delay vocabulary, which specific failure mode a Celery worker becoming unresponsive actually represents from the scheduler's perspective, and state explicitly whether Celery's default retry behavior could produce the same duplicate-effect problem as this lesson's failure drill for a task that isn't naturally idempotent (e.g., a task that sends an email, or charges a payment, versus one that merely recomputes a value that's safe to recompute).

## Done when

You can name and distinguish the four failure modes covered here for a concrete scenario, you've directly observed that a network partition and a genuine crash are indistinguishable from a client's timeout-only perspective, and you can trace, using your own step 3/failure-drill reasoning, exactly why naive retry-on-timeout is unsafe without an additional mechanism — setting up the specific problem the next several lessons each address a piece of.
