# Lesson 11: Webhooks

## Objective

Design and implement outbound webhooks correctly — signed payloads for authenticity, retry logic for delivery failures, and idempotent receiver-side handling — understanding webhooks as essentially an inverted API call where your server becomes the client and must handle all the same reliability problems from the other side.

## Prerequisites

Distributed Lesson 1-2 (failure models and RPC idempotency — a webhook delivery has exactly the same ambiguity and duplicate-risk properties as any other network call), networking Lesson 5 (TLS — webhook payloads must be signed since the receiving endpoint, unlike a normal API call, has no other way to verify the request genuinely came from you).

## Learn

**What a webhook is, and why it inverts the usual client/server relationship.** Normally, a client calls your API. A webhook is your server calling *someone else's* endpoint to notify them of an event (e.g. "subscription activated," "payment received") — your server is now the client, and their endpoint is the server, with all the same reliability concerns from distributed Lesson 1 now applying to *your* outbound calls rather than calls made to you.

**Why webhook payloads must be signed.** Since a webhook is an unsolicited incoming request to the receiver's endpoint (not a response to something they initiated), the receiver has no inherent way to verify it genuinely came from you and wasn't forged by an attacker who discovered or guessed their webhook URL. The standard fix: sign the payload with a secret shared between you and the receiver (established at webhook-registration time), typically via HMAC — compute a signature over the payload using the shared secret, include it in a header (e.g. `X-Signature`), and the receiver recomputes the same signature independently and confirms it matches before trusting the payload's contents. This is a genuinely important security control, not a formality — an unsigned webhook endpoint is an open door for anyone who discovers the URL to inject fake events.

**Retry logic: your server now needs the same reliability handling any API client needs.** If the receiver's endpoint is down, slow, or returns an error, your webhook delivery attempt has failed — and per distributed Lesson 1, you generally can't be certain whether they actually processed it before failing to respond. Standard practice: retry with exponential backoff (increasing delay between attempts, to avoid hammering a struggling receiver) up to some maximum number of attempts, after which the delivery is considered permanently failed (typically logged/surfaced for manual investigation, since silently giving up on notifying about a real event like "payment received" is rarely acceptable).

**Idempotent receiver-side handling: the receiver's responsibility, not yours, but worth designing for.** Because retries can result in the same webhook event being delivered more than once (the same ambiguity as any other retry scenario), a well-designed webhook payload includes a unique event ID, letting the *receiver* deduplicate on their end (directly reusing distributed Lesson 2/7's idempotency pattern) — this is the receiver's implementation responsibility, but as the sender, including a stable, unique event ID in every payload is what makes that deduplication possible for them at all; omitting it forces every receiver to either accept duplicate processing risk or build their own fragile deduplication heuristics.

## Attempt

1. Implement a webhook sender: given an event (e.g. a simulated "subscription activated" event with some payload data), POST it to a configured receiver URL, with an HMAC signature (using a shared secret) included in a header, computed over the raw request body.

2. Implement a webhook receiver (a separate small server) that verifies the incoming signature against its own copy of the shared secret before processing the payload, rejecting (with an appropriate 4xx status) any request with a missing or incorrect signature. Test both a correctly signed request (accepted) and a tampered payload with a stale signature (rejected) to confirm signature verification is actually functioning, not just present in code but unused.

3. Implement retry logic on the sender side: if the receiver returns a 5xx or the connection times out, retry with exponential backoff (e.g. 1s, 2s, 4s, 8s delays) up to a maximum of 5 attempts, logging each attempt's outcome. Test it against a receiver that deliberately fails the first 2 attempts and succeeds on the 3rd, confirming your sender correctly persists through the failures and eventually succeeds, with the expected backoff timing between attempts.

4. Add a unique event ID to every webhook payload, and implement idempotent processing on the receiver side (reusing distributed Lesson 2/7's pattern: track processed event IDs, skip reprocessing duplicates). Simulate a duplicate delivery (send the same event ID twice, as would happen if a retry succeeded on the receiver's end but the sender didn't get the success response and retried anyway) and confirm the receiver correctly processes it only once.

## Verify

For step 2, show both the accepted (valid signature) and rejected (invalid/tampered signature) request/response pairs, confirming signature verification is genuinely functioning. For step 3, show your actual logged retry timing, confirming it matches the expected exponential backoff pattern.

## Failure drill

Implement a receiver that verifies the payload signature using a naive, non-constant-time string comparison (e.g. a simple `==` or `strcmp`-style comparison rather than a timing-safe comparison function) and research (via documentation on your language's cryptography library, e.g. Go's `crypto/hmac.Equal` or `subtle.ConstantTimeCompare`) why this specific implementation detail matters — a naive comparison can leak timing information about how many leading characters of the signature matched before it found a mismatch, which is (in principle, given a very large number of attempts) an exploitable timing side-channel for an attacker attempting to forge a valid signature byte by byte. You don't need to demonstrate an actual successful timing attack (genuinely difficult to do reliably in a simple lab exercise), but explain in your own words why "use a constant-time comparison for secret verification" is a real, specific, non-obvious security requirement — the difference between the naive and safe version is invisible in normal functional testing (both correctly accept valid signatures and reject invalid ones), which is exactly what makes this class of subtle security bug easy to introduce without noticing.

## Transfer

If TARDOC ever needs to notify an external system (e.g. a clinic's own practice-management software) of events like "invoice generated" or "payment received," describe how you'd apply this lesson's signing and retry pattern to that specific integration, and state explicitly what unique, stable identifier you'd use as the event ID (a database primary key, a UUID generated at event-creation time) to enable the receiving system's own deduplication, per this lesson's idempotent-receiver design.

## Done when

Your webhook sender correctly signs payloads and your receiver correctly verifies and rejects tampered ones, your sender's retry logic correctly persists through transient failures with proper exponential backoff, you've implemented and tested idempotent receiver-side processing using a unique event ID, and you can explain why constant-time signature comparison matters even though the difference is invisible to ordinary functional testing.
