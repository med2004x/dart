# Lesson 8: Concurrent Updates

## Objective

Detect and correctly reject conflicting concurrent updates to the same resource using optimistic concurrency control (ETags/version numbers), rather than letting the last write silently win and destroy an earlier, unseen change.

## Prerequisites

Database-internals Lesson 5 (transaction isolation — lost updates at the database layer are the same underlying problem this lesson solves at the API layer, one level up), Lesson 4 (HTTP semantics — this lesson uses conditional request headers built into HTTP for exactly this purpose).

## Learn

**The lost-update problem, restated at the API level.** Two clients both `GET /clinics/{id}`, both see the same current state, both independently decide on a change, both `PUT` their updated version back — the second `PUT` overwrites the first, silently discarding the first client's change, with neither client aware anything went wrong. This is precisely database-internals Lesson 5's lost-update anomaly, except now it's happening across two separate HTTP requests (each internally a properly isolated database transaction) rather than within one database transaction — meaning the database's own isolation guarantees don't automatically protect against it, since each individual `PUT` request is, from the database's perspective, a perfectly valid, self-consistent write.

**Optimistic concurrency control: detect conflicts instead of preventing them upfront.** Rather than locking the resource for the duration of a client's "think time" between GET and PUT (which would be both impractical over HTTP's stateless request model and a real availability cost, since one slow client would block others), optimistic concurrency lets both clients proceed freely but requires each write to state *which version it was based on* — and rejects the write if that's no longer the current version, meaning something else changed the resource in between.

**ETags and conditional requests, the standard HTTP mechanism for this.** The server includes an `ETag` header (an opaque version identifier — commonly a hash of the content, or an incrementing version number) in every `GET` response. When the client later sends a `PUT` or `PATCH`, it includes an `If-Match: <etag>` header with the version it read. The server compares this against the resource's *current* ETag: if they match, the update proceeds (and a new ETag is generated); if they don't match, the server rejects the update with `412 Precondition Failed`, telling the client explicitly "the resource changed since you last read it — refetch and reconsider your update" rather than silently overwriting.

**Why this is better than the alternative of pessimistic locking for most API use cases.** Pessimistic locking (acquiring an explicit lock before allowing any read, held until the corresponding write completes) would require holding server-side state across multiple HTTP requests from possibly-abandoned clients — a client that reads and then never writes back would leave the resource locked indefinitely without some additional timeout/cleanup mechanism. Optimistic concurrency requires no such held state — every request is independently stateless (consistent with HTTP's general design), and conflicts are simply detected and reported at write time rather than prevented at read time.

## Attempt

1. Implement `GET /resource/{id}` returning an `ETag` header (a simple version integer, or a hash of the resource's serialized content, is fine) alongside the resource body.

2. Implement `PUT /resource/{id}` requiring an `If-Match` header matching the resource's current ETag. On a match, apply the update and return a new ETag; on a mismatch, return `412 Precondition Failed` with a clear error body explaining the resource has changed.

3. Directly reproduce the lost-update scenario without protection first: simulate two clients both reading the same resource (same ETag), both independently modifying their local copy, and both submitting a `PUT` — but only the *first* one includes the correct, still-current `If-Match` value; have the second client send its now-stale ETag. Confirm the second client's request is correctly rejected with 412, rather than silently succeeding and overwriting the first client's change.

4. Implement the client-side recovery flow for a 412 response: on receiving 412, the client should refetch the current resource (getting a fresh ETag), decide how to reconcile its intended change with the new current state (for this exercise, a simple "reapply my change on top of the fresh data and retry" is sufficient — real systems sometimes need more sophisticated merge logic depending on the domain), and resubmit with the fresh ETag. Test this full retry-after-conflict flow end to end and confirm it eventually succeeds correctly.

## Verify

For step 3, show both clients' actual requests (including their `If-Match` headers) and the server's actual responses, confirming the first succeeds (200 with new ETag) and the second is correctly rejected (412), with the resource's final state matching only the first client's intended change, not silently reflecting a merged or overwritten result.

## Failure drill

Implement a version of `PUT` that *ignores* the `If-Match` header entirely (accepts and applies any update regardless of the client-provided version) and rerun step 3's exact concurrent-update scenario against this broken version. Confirm the second client's stale-based update now silently succeeds, overwriting the first client's change with no error or warning to either party — reproducing the original lost-update problem exactly, now with your own test as concrete evidence rather than an abstract description. Explain why this specific bug — an endpoint that accepts `If-Match` syntactically but doesn't actually enforce it — would be easy to miss in casual testing (a single client updating a resource with no concurrent conflict would never trigger the difference) and would only surface under real concurrent usage, exactly the kind of bug that's invisible in typical single-user manual testing but real and damaging under actual production traffic.

## Transfer

If TARDOC's clinic-configuration editing or Mahall's product-editing endpoints currently use plain `PUT`/`PATCH` with no `If-Match`/ETag protection, describe a realistic scenario in your own system where two legitimate concurrent edits to the same resource could occur (e.g. a clinic admin and a support staff member both editing the same clinic's settings around the same time), and state explicitly what silently happens today under that scenario without this lesson's protection — then describe what adding ETag-based optimistic concurrency would change about that outcome.

## Done when

Your `PUT` endpoint correctly enforces `If-Match` and rejects stale updates with 412, you've directly reproduced the lost-update problem both with and without protection using the same concurrent-client test scenario, and you've implemented a working client-side retry-after-conflict flow that successfully reconciles and resubmits after a detected conflict rather than just failing permanently.
