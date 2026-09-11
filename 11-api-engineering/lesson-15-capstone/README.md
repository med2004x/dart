# Lesson 15: Capstone — A Complete API

## Objective

Integrate Lessons 1-14 into one complete, production-shaped API, applying every principle covered — contract-first design, correct resource modeling, validation, idempotency, concurrency control, authentication/authorization, rate limiting, webhooks, async job handling, versioning, and operational endpoints — rather than a CRUD service missing the parts that only matter under real, adversarial, or high-scale conditions.

## Prerequisites

Lessons 1-14, completed. No new theory — this is deliberately an integration exercise, the api-engineering track's equivalent of the integration capstones throughout this curriculum.

## Learn

There is no new material. This lesson tests whether the individually-understood pieces from Lessons 1-14 actually compose correctly — a genuinely common gap: an API that correctly implements idempotency (Lesson 7) in isolation and correctly implements optimistic concurrency (Lesson 8) in isolation can still have a bug in how those two interact (e.g., does a retried idempotent request correctly re-check the `If-Match` precondition, or does it bypass that check on the cached-response path?) — exactly the kind of question only surfaces when building something that requires both mechanisms to work together, not when testing either alone.

## Attempt

Design and build a complete API for a real resource type from TARDOC or Mahall's actual domain (or a plausible extension of it) with this minimum required scope:

1. **Contract-first**: a complete OpenAPI spec (Lesson 3) written before implementation, covering all endpoints below, with reusable error-schema components (Lesson 3).

2. **Resource model** (Lesson 2): at least one top-level resource and one genuinely nested sub-resource, with a justified nesting decision.

3. **Standard CRUD with correct semantics** (Lesson 4): list (paginated per Lesson 6, using cursor-based pagination), get-one, create, update, delete — with correct HTTP methods, correct idempotency/safety properties per method, and correct status codes throughout.

4. **Validation and structured errors** (Lesson 5): boundary validation with field-level error detail for at least one endpoint accepting a nontrivial request body.

5. **Idempotency** (Lesson 7): idempotency-key support on your create endpoint specifically, since creation is the canonical non-naturally-idempotent `POST` case.

6. **Optimistic concurrency** (Lesson 8): ETag/`If-Match` support on your update endpoint, and — this is the integration point Learn specifically warned about — confirm your idempotency-key logic (item 5) and your `If-Match` logic (item 6) correctly compose if a single endpoint could theoretically need both (e.g. an idempotent update, if your design calls for one).

7. **Authentication and resource-level authorization** (Lesson 9): token-based auth, with both role-based and resource-ownership checks on at least one endpoint, correctly returning 401 vs 403 as appropriate.

8. **Rate limiting** (Lesson 10): token-bucket rate limiting per authenticated identity, with standard rate-limit response headers.

9. **At least one async operation** (Lesson 12): one endpoint representing a genuinely slow operation, using the accept-and-poll (or accept-and-webhook, Lesson 11) pattern rather than a long-held synchronous response.

10. **Versioning readiness** (Lesson 13): structure your implementation so a `/v2` of at least one endpoint could be introduced without breaking `/v1` — you don't need to actually build a v2, but your routing/handler structure should make clear how a v2 would be added without touching v1's code path.

11. **Operational endpoints** (Lesson 14): correctly distinguished liveness and readiness endpoints, with readiness checking your actual database dependency.

## Verify

Write integration tests exercising cross-lesson interactions specifically, not just each feature in isolation — for example: an idempotent create request that's retried after the resource was concurrently modified by someone else (does idempotency correctly return the original response, unaffected by the unrelated concurrent modification?), or a rate-limited client whose async job submission (item 9) correctly counts against their rate limit the same as any other request. Report your test results, and specifically call out any interaction you found that didn't compose as cleanly as you initially expected.

## Failure drill

Pick one specific cross-feature interaction from your capstone and deliberately break it — for example, make your rate limiter apply *before* authentication (checking IP-based limits before knowing who the authenticated user is, contradicting Lesson 10's failure drill about per-identity limiting being preferable to per-IP), or make your async job endpoint (item 9) bypass authorization checks that your synchronous endpoints correctly enforce (a realistic oversight, since async endpoints are sometimes implemented as an afterthought separate from the main authenticated request path). Demonstrate the resulting vulnerability or incorrect behavior with a real test, then fix it, and explain in your own words why integration points between separately-correct features are disproportionately likely to hide bugs, compared to bugs within any single feature's own isolated logic.

## Transfer

Compare your capstone API's design against a real, production API you've used or built (TARDOC's own API is the most relevant comparison, given your project history) — identify at least 2 specific principles from this track that your real, existing API is currently missing or handling incompletely, and describe concretely what you'd need to add, referencing the specific lesson and mechanism, to bring it in line with what this capstone required.

## Done when

Your capstone API implements all 11 required elements with real, passing tests — not just each element in isolation, but including at least one deliberately-tested cross-feature interaction — and you've found, demonstrated, and fixed at least one genuine integration bug via the failure drill, with an honest explanation of why it specifically occurred at the boundary between two otherwise-correct features rather than within either one alone.
