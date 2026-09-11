# Lesson 8: Capstone — Network Stack Component

## Objective

Build one substantial networking component end to end, integrating sockets, framing, TCP semantics, and packet-level verification from Lessons 1-7, rather than another CRUD service — the networking track's equivalent of computer architecture Lesson 8 and C Lesson 6's integration capstones.

## Prerequisites

Lessons 1-7, completed. No new theory — this is deliberately an integration exercise where gaps between separately-understood pieces (framing, TCP semantics, connection handling) tend to surface under the pressure of building something that has to actually work correctly, not just in the lesson's narrow test case.

## Learn

There is no new material here. Stanford's CS144 (referenced throughout this track) builds a full user-space TCP implementation as its capstone project — a genuinely excellent, rigorous option if you want the deepest possible version of this exercise, and using it directly (rather than reinventing an equivalent from scratch) is a reasonable and encouraged choice, consistent with this curriculum's general principle of using strong existing courses where they exist rather than reconstructing worse versions of them.

If not using CS144 directly, the four component options below are designed to each force integration of a different subset of Lessons 1-7's material, so the choice should be made based on which combination you most want to solidify.

## Attempt

Choose **one** of the following components and build it to genuine, tested completion:

**Option A: A minimal reverse proxy / load balancer.** Extends Lesson 6's L7 balancer into something closer to production-shaped: support multiple backends with health checking (periodically probe each backend, stop routing to ones that fail), path-based routing (Lesson 6), and connection reuse to backends (don't open a new backend connection per request — pool and reuse them, directly applying Lesson 4's keep-alive concept in the reverse direction, proxy-to-backend rather than client-to-server).

**Option B: A TCP-based key-value store server.** A server accepting TCP connections (Lesson 1) speaking a simple custom protocol (design your own wire format — this is a good place to apply Lesson 1's length-prefixed framing deliberately) supporting `GET key`, `SET key value`, `DELETE key` operations, correctly handling multiple concurrent client connections (each in its own goroutine/thread, applying OS Lesson 2's concurrency concepts), with correct partial-read/partial-write handling per Lesson 1's core lesson.

**Option C: A DNS caching resolver.** Extends Lesson 3's minimal DNS client into a small caching proxy resolver: accepts DNS queries (UDP, Lesson 3's wire format), checks a local cache first (respecting TTL, per Lesson 3's caching discussion — an entry past its TTL must be treated as expired, not served stale), and if not cached (or expired), forwards to an upstream resolver, caches the response with its TTL, and returns it to the original client.

**Option D: A simplified reliable-transport protocol over UDP.** Extends Lesson 2's simplified retransmission exercise into something closer to a real minimal transport protocol: proper sequence numbering, cumulative acknowledgment, a sliding window (not just one-at-a-time send-and-wait), and basic timeout-based retransmission — essentially a small, deliberately simplified reimplementation of what TCP itself provides, which is exactly the kind of exercise that turns "I understand TCP's guarantees conceptually" (Lesson 2) into "I've built a system that provides those same guarantees myself."

Whichever option you choose, the requirements are the same:

1. **Write tests from invariants, not just examples.** For Option B, an invariant might be "a `GET` immediately following a `SET` for the same key always returns the value just set, regardless of how many other concurrent clients are also issuing requests" — a property to verify under concurrent load, not just a single-client happy-path test.
2. **Benchmark it.** Measure actual throughput and/or latency under a realistic load pattern (e.g. many concurrent clients for Option B, many concurrent proxied requests for Option A), and report real numbers, not estimates.
3. **Document limitations explicitly.** Every option above is a deliberately simplified version of a much more sophisticated real system (a real load balancer, a real distributed key-value store, a real DNS resolver, real TCP). Write a short, honest section listing what your implementation does *not* handle that a production version would need to (e.g. Option D's simplified protocol almost certainly doesn't implement real congestion control, per Lesson 2's discussion — say so explicitly, rather than implying the toy version is production-equivalent).

## Verify

Provide your actual benchmark numbers (throughput, latency, or whatever's most relevant to your chosen option) under at least two different load levels (e.g. 10 concurrent clients vs. 100), and confirm your invariant-based tests pass consistently across multiple runs, not just once.

## Failure drill

Take your chosen component and subject it to one deliberately adverse condition drawn from earlier lessons in this track: for Option A or B, use Lesson 7's artificial packet loss (`tc netem`) between client and server and confirm your component still behaves correctly (perhaps slower, but not incorrectly) under real, induced loss rather than only the clean, lossless conditions of local development. For Option C, deliberately query a domain that returns `NXDOMAIN` (Lesson 3's failure drill) through your caching resolver and confirm it correctly caches (or correctly chooses not to cache, per your design) the negative result rather than crashing or hanging. For Option D, deliberately introduce packet reordering (not just loss) and confirm your sequence-number-based reassembly still produces correct output, extending Lesson 2's failure drill about ordering being a separate guarantee from reliability.

## Transfer

Compare your chosen, deliberately simplified component against the real-world system it approximates (nginx/HAProxy for Option A, Redis for Option B, a real recursive resolver like Unbound for Option C, or real TCP for Option D) — pick one specific feature the real system has that yours doesn't, and explain in a few sentences why that feature exists (what real-world condition or requirement it's addressing) that your simplified version's scope didn't need to handle.

## Done when

Your chosen component passes invariant-based tests under concurrent/realistic load, you have real benchmark numbers at two different load levels, you've subjected it to at least one adverse network condition from earlier lessons and confirmed correct (if degraded) behavior, and you've written an honest, specific limitations section rather than an implied claim of production-readiness.
