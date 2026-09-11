# Lesson 7: Idempotency

## Objective

Implement idempotency keys for a real, non-naturally-idempotent write endpoint (a `POST` that creates something with a real side effect), directly applying distributed Lesson 2's mechanism at the API-design level, where it's most commonly needed in practice.

## Prerequisites

Distributed Lesson 2 (RPC idempotency — this lesson applies the identical underlying mechanism to HTTP APIs specifically), Lesson 4 (HTTP semantics — POST's lack of built-in idempotency guarantee is exactly the gap this lesson fills).

## Learn

**Restating the core problem at the HTTP-API level.** A client sends `POST /payments` to charge a customer, the server processes it successfully, but the response is lost (network failure, client timeout) before the client receives confirmation. The client, having no way to distinguish "the server never got my request" from "the server processed it but I didn't get the response" (distributed Lesson 1's fundamental ambiguity), reasonably retries — and without protection, the customer gets charged twice. This is not a hypothetical concern; it's a routine, expected consequence of building any API over an unreliable network, and payment APIs specifically are the textbook case where getting this wrong has direct, serious real-world consequences.

**Idempotency keys, the standard HTTP-API pattern.** The client generates a unique key (typically a UUID) for each *logical* operation (not per HTTP attempt — the same key is reused across retries of the same logical request) and sends it in a header, conventionally `Idempotency-Key`. The server, upon receiving a request with an idempotency key it has already successfully processed, returns the *original* response (from its stored record) without re-executing the operation's side effects — functionally identical to distributed Lesson 2's request-ID mechanism, just expressed via an HTTP header instead of an RPC parameter, and now the standard, widely-recognized convention (Stripe's API popularized this exact pattern, and it's now common practice across payment and other side-effect-heavy APIs).

**What must be stored, and for how long.** The server needs a durable record (a database table, not just in-memory state — a server restart shouldn't lose idempotency protection, echoing distributed Lesson 7's durability requirement for at-least-once processing) mapping each seen idempotency key to the response that was returned for it, and ideally also the request body's hash (to detect and reject the case where a client reuses the same idempotency key for a *genuinely different* request body — a client error, but one worth actively detecting rather than silently returning a stale, mismatched response). Idempotency keys are typically retained for a bounded window (e.g. 24 hours), not forever — an operationally reasonable tradeoff, since retrying a multi-day-old failed request is rarely the intended use case, and unbounded retention would grow the tracking table indefinitely.

**Concurrent requests with the same key.** If two requests with the same idempotency key arrive genuinely concurrently (a client retrying very aggressively, or a bug causing duplicate near-simultaneous sends) before the first has finished processing, a correct implementation needs to handle this too — typically by having the second request either wait for the first to complete and then return its result, or fail fast with a specific "request already in progress" response, rather than both proceeding to execute the side effect concurrently, which would defeat the entire mechanism's purpose.

## Attempt

1. Implement a `POST` endpoint with a real (simulated) side effect — e.g. "create an order" that increments a counter or appends to a list representing "orders placed," standing in for a real payment charge. Confirm, without any idempotency protection, that sending the same logical request twice (two separate POST calls with identical body) produces two side effects (two orders created) — reproducing the base problem concretely before fixing it.

2. Add idempotency key support: accept an `Idempotency-Key` header, store a durable record (a database table: key → response, plus a hash of the request body) upon successful processing, and on a subsequent request with a previously-seen key, return the stored response without re-executing the side effect. Retest step 1's scenario with the same idempotency key on both calls and confirm only one order is now created.

3. Implement the mismatched-body detection from Learn: send two requests with the *same* idempotency key but *different* request bodies, and confirm your server detects this (via the stored request-body hash) and returns an explicit error (rather than silently returning the first request's now-mismatched stored response) — a genuinely different situation from a legitimate retry, and one that should be surfaced clearly to the caller as likely indicating a client-side bug.

4. Implement basic concurrent-request handling: using a lock or a "processing" marker stored alongside the idempotency key record, ensure that if a second request with the same key arrives *while the first is still being processed* (not yet completed and recorded), the second request either waits for the first to finish and returns its result, or fails fast with a clear "in progress, retry shortly" response — test this by adding an artificial delay to your side-effect logic and firing two genuinely concurrent requests with the same key to confirm your handling actually engages rather than allowing a race.

## Verify

For step 2, show the side-effect counter/list's actual state after the same-key retry test, confirming exactly one side effect occurred despite two HTTP requests. For step 4, show evidence (logs, or the final side-effect state) confirming the concurrent-request race was correctly handled rather than both requests executing the side effect.

## Failure drill

Remove your step 4 concurrent-request protection (revert to only checking for an already-*completed* idempotency key, with no protection against two requests arriving while the first is still mid-processing) and rerun the concurrent-request test with an artificial processing delay. Confirm both requests now proceed to execute the side effect (e.g. two orders created) despite sharing the same idempotency key, because neither request found a *completed* record yet at the moment each checked. Explain why this specific race — the gap between "request started" and "request's result durably recorded" — is a genuinely distinct failure mode from the sequential-retry case step 2 already handles correctly, and why simply checking for a completed record isn't sufficient; the mechanism needs to also account for requests that are concurrently in flight, not just ones that already fully finished.

## Transfer

If TARDOC's subscription-activation flow or any payment-adjacent operation in your own systems currently lacks idempotency-key protection, describe specifically, using this lesson's mismatched-body-detection and concurrent-request handling as the concrete requirements, what you'd need to add — and state explicitly what real-world failure mode (a client's network retry, a double-tap on a submit button before the UI disables it, a proxy or load balancer's own retry logic) is most likely to trigger a duplicate request against that specific endpoint in practice, grounding the abstract "idempotency matters" advice in your actual system's realistic retry sources.

## Done when

You've reproduced the duplicate-side-effect problem without protection, then implemented and verified idempotency-key-based deduplication actually prevents it across a real retry scenario, you've implemented and tested mismatched-body detection as a distinct case from a legitimate retry, and you've demonstrated — via the failure drill — the specific concurrent-request race that simple "check if already completed" logic misses, and why an in-progress marker is needed in addition to it.
