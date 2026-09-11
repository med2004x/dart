# Lesson 4: HTTP Semantics for APIs

## Objective

Apply HTTP's method semantics (safety, idempotency) and status codes correctly and consistently across a real API, going beyond networking Lesson 4's transport-level mechanics into how an API's design should respect and communicate through these semantics.

## Prerequisites

Networking Lesson 4 (HTTP mechanics — request/response structure, status code categories), Lesson 2 (resource modeling — this lesson applies HTTP methods to the resources modeled there).

## Learn

**Safety and idempotency, applied at the API-design level (not just protocol trivia).** `GET` must be safe (no side effects) — an API that implements `GET /clinics/{id}/deactivate` (a side-effecting operation exposed via GET) violates this convention and creates real risk: browsers, proxies, and crawlers may prefetch or repeat GET requests, assuming safety, potentially triggering the side effect unintentionally. `PUT` and `DELETE` should be idempotent — calling `DELETE /clinics/{id}` twice should have the same end effect as calling it once (the clinic is deleted; a second call finding it already gone should return a clear "already deleted" response, typically still treated as success or a specific 404, not an error implying something went wrong). `POST` is neither safe nor generally idempotent by convention — which is exactly why distributed Lesson 2's request-ID-based idempotency mechanism matters specifically for `POST` operations representing non-idempotent creates, since HTTP itself provides no built-in protection there the way it conventionally implies for `PUT`/`DELETE`.

**Status codes as a precise vocabulary, not "200 or 500."** Reusing status codes consistently across your entire API (not inventing a different meaning for the same code in different endpoints) is what lets a generic HTTP client library, a monitoring system, or a new developer reading your API reliably interpret responses without reading every endpoint's specific documentation. Common, precise usage: `200 OK` (successful GET/PUT/PATCH with a body), `201 Created` (successful POST that created a resource — should include a `Location` header pointing to the new resource), `204 No Content` (successful operation with no body to return, e.g. a successful DELETE), `400 Bad Request` (the request itself is malformed — client's fault, not retriable without fixing the request), `401 Unauthorized` (authentication is missing or invalid — Lesson 9 covers this properly), `403 Forbidden` (authenticated, but not permitted to perform this action), `404 Not Found`, `409 Conflict` (the request conflicts with the current state — e.g. Lesson 8's concurrent-update conflict), `422 Unprocessable Entity` (well-formed request, but semantically invalid — e.g. valid JSON with a field value that fails a business rule), `429 Too Many Requests` (Lesson 10's rate limiting), `500 Internal Server Error` (something went wrong on the server side, not the client's fault).

**Why 4xx vs. 5xx distinction matters for client behavior, restated from networking Lesson 4 but now at the API-design level.** A well-designed API's status codes let a client's *generic* retry logic behave correctly without needing endpoint-specific knowledge: retry 5xx (the server might recover, or a load balancer might route the retry to a healthy instance), don't blindly retry 4xx (the request itself needs to change first) — an API that returns 500 for a client's malformed request (a common but genuinely incorrect design mistake) breaks this generic retry assumption and can cause clients to retry a request that will never succeed as written, wasting resources on both sides.

## Attempt

1. Audit one of your own existing API endpoints (TARDOC or Mahall) and check whether its HTTP method usage matches this lesson's safety/idempotency conventions — specifically, confirm no GET endpoint has side effects, and check whether your DELETE endpoints behave idempotently (calling delete on an already-deleted resource — what does it currently return, and is that appropriate per this lesson's guidance).

2. For the same endpoint (or a small set of related endpoints), produce a status-code table: every distinct outcome (success and each failure mode) mapped to a specific status code, and confirm you're not reusing the same status code for two semantically different outcomes, or using different codes for what's actually the same kind of outcome across similar endpoints.

3. Deliberately construct a bad example: an endpoint that returns `200 OK` even when the operation actually failed (with the actual error only discoverable by inspecting the response body), and rewrite it to return an appropriate 4xx or 5xx status code instead, with the error detail in the body. Test both versions with a simple client (`curl -w "%{http_code}"` or equivalent) and confirm the corrected version allows detecting failure from the status code alone, without needing to parse the body first.

4. Implement (or find and confirm) idempotent DELETE behavior for one real resource: call DELETE on a resource, confirm it succeeds; call DELETE again on the same (now already-deleted) resource ID, and confirm it returns a sensible, non-error-implying response (commonly still a 204, or a 404 clearly indicating "already gone" rather than "something broke") rather than a 500 or an ambiguous error.

## Verify

Present your step 2 status-code table for review, and for step 3, show actual `curl` output (status code and body) for both the "always 200" bad version and your corrected version, demonstrating the practical difference in what a status-code-only check can detect.

## Failure drill

Take your corrected step 3 endpoint (correct status codes) and write a small generic client function that retries any request receiving a 5xx response (up to 3 times with backoff) but does *not* retry on 4xx. Test it against both a genuinely transient 5xx scenario (simulate your server temporarily failing, e.g. by adding a deliberate `if firstAttempt { return 500 }` in test code) and a genuine 4xx scenario (send a deliberately malformed request). Confirm the generic retry logic correctly recovers from the transient 5xx (retries succeed) and correctly avoids wasting retries on the un-retriable 4xx (fails fast instead). Explain why this generic client logic — reusable across any endpoint that correctly follows the status code conventions — would have behaved incorrectly against your *original*, "always 200" bad-example endpoint from step 3, since it provides no status-code signal at all for the retry logic to act on.

## Transfer

If TARDOC's subscription-activation endpoint or Mahall's order-creation endpoint currently uses `POST`, describe explicitly, using distributed Lesson 2's idempotency-via-request-ID pattern and this lesson's "`POST` is not idempotent by HTTP convention" point together, whether that endpoint is currently safe against a client's naive retry-on-timeout behavior — and if not, what you'd need to add (a request-ID mechanism, per distributed Lesson 2) to make it genuinely safe, distinct from simply hoping clients don't retry.

## Done when

You've audited a real endpoint against HTTP method safety/idempotency conventions and identified any violations, you've produced a complete, consistent status-code table for a real set of endpoints and fixed at least one status-code misuse you found, and you've demonstrated — with actual working retry-logic code — why correct status code usage is what makes generic, endpoint-agnostic client retry logic possible at all.
