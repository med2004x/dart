# Lesson 10: Rate Limits

## Objective

Implement API rate limiting using the token bucket algorithm, understand why it's preferred over naive fixed-window counting, and correctly communicate limit state to clients via standard headers so well-behaved clients can self-regulate rather than just getting rejected.

## Prerequisites

Mathematics track's queueing lesson (math-for-engineering step — rate limiting is directly about controlling the arrival rate `λ` from that lesson's M/M/1 model, keeping it within what the system can sustain).

## Learn

**Why rate limiting exists.** Without it, a single misbehaving client (buggy retry loop, a scraper, or genuinely malicious traffic) can consume disproportionate server resources, degrading service for every other client — directly connecting to the math-for-engineering queueing lesson's `ρ→1` blowup: an unbounded arrival rate from even one source can push the whole system's utilization toward saturation, and rate limiting is the mechanism that caps `λ` per client before that happens.

**Fixed-window counting, and its real flaw.** The simplest approach: count requests per client within fixed time windows (e.g. "max 100 requests per minute," resetting the counter every minute boundary). The flaw: a client can send 100 requests in the last second of one window and another 100 in the first second of the next window — 200 requests in roughly 2 seconds, technically compliant with "100 per minute" in each window individually, but violating the actual intended rate by a wide margin at the window boundary. This is a genuine correctness gap, not a theoretical nitpick — it's exploitable by any client aware of the window boundaries.

**Token bucket: the standard fix.** Conceptually, each client has a "bucket" that holds up to some maximum number of tokens (the burst capacity), refilling at a steady rate (e.g. one token every 600ms for a 100-per-minute limit) up to that maximum. Each request consumes one token; if the bucket is empty, the request is rejected (or queued, depending on design). This smooths out the boundary-gaming flaw of fixed windows — since tokens refill continuously rather than resetting in a discrete jump, there's no window boundary to exploit — while still allowing legitimate short bursts (up to the bucket's max capacity) rather than forcing a perfectly uniform request pace, which real client usage patterns rarely follow anyway.

**Communicating limit state to clients: the standard headers.** `X-RateLimit-Limit` (the maximum, e.g. bucket capacity or window total), `X-RateLimit-Remaining` (tokens/requests left in the current state), `X-RateLimit-Reset` (when the limit will refresh/refill meaningfully) — and, when a request is rejected, `429 Too Many Requests` (Lesson 4's status code) with a `Retry-After` header telling the client specifically how long to wait before retrying. A well-behaved client can use these headers to self-regulate (slow down proactively as `Remaining` approaches zero) rather than only discovering the limit by hitting a 429 — good rate-limit design actively helps clients avoid being rate-limited, not just punishes them after the fact.

## Attempt

1. Implement fixed-window rate limiting (a simple per-client counter reset every N seconds) and demonstrate its boundary-gaming flaw directly: send a burst of requests timed to straddle a window boundary (most just before the reset, the rest just after) and confirm you can exceed the intended rate within a short overall time span, despite each individual window's count staying within its own limit.

2. Implement token bucket rate limiting: a per-client token count, a maximum capacity, and a refill rate (tokens added per unit time, capped at capacity). On each request, attempt to consume a token; reject with 429 if none available.

3. Repeat step 1's boundary-timed burst test against your token bucket implementation and confirm it correctly limits the effective rate across the boundary, unlike the fixed-window version — report the actual request counts allowed through in both implementations for the identical burst pattern, demonstrating the concrete difference.

4. Add the standard rate-limit response headers (`X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`) to every response (not just rejected ones), and `Retry-After` specifically on 429 responses. Write a simple client that reads `X-RateLimit-Remaining` and proactively slows its request rate as it approaches zero, confirming it can avoid ever receiving a 429 by self-regulating based on the headers alone.

## Verify

For step 3, report the exact request counts allowed through by each implementation for the identical boundary-straddling burst pattern, demonstrating the fixed-window version exceeding the intended rate and the token-bucket version correctly enforcing it.

## Failure drill

Implement rate limiting keyed only by client IP address (a common, simple default) and demonstrate its limitation directly: simulate multiple distinct logical clients sharing one IP (e.g. multiple users behind the same NAT/corporate network, or multiple browser tabs from the same machine) all counting against the same shared limit, causing one client's heavy usage to incorrectly exhaust the rate limit for other, unrelated clients sharing that IP. Then implement per-authenticated-user (or per-API-key) rate limiting instead (using Lesson 9's authentication token to identify the actual logical client, not just the network-level IP) and confirm this resolves the shared-IP problem, correctly isolating each real client's limit from the others. Explain why IP-based limiting, despite being simpler to implement (no authentication dependency), is a genuinely weaker mechanism for exactly this reason, and why per-identity limiting, once authentication (Lesson 9) is available, is generally the more correct default.

## Transfer

If TARDOC's or Mahall's API doesn't currently implement rate limiting, describe a realistic scenario specific to your own system where an unlimited API could cause real harm — e.g., a runaway client script hammering TARDOC's transcription-triggering endpoint, given the Groq API key rotation and cost implications mentioned in your project history — and state what rate-limit parameters (requests per time window, burst capacity) would be reasonable defaults for that specific endpoint, reasoning from the actual cost or resource constraint the limit is meant to protect against, not an arbitrary round number.

## Done when

You've directly demonstrated fixed-window rate limiting's boundary-gaming flaw with a real, exploitable test case, you've implemented and verified token bucket limiting correctly closes that gap using the same test scenario, you've implemented and tested the standard rate-limit response headers with a client that successfully self-regulates using them, and you've demonstrated the specific weakness of IP-based limiting versus identity-based limiting with a concrete shared-IP scenario.
