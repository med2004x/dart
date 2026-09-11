# Lesson 13: Versioning

## Objective

Design an API versioning strategy that lets you evolve the contract without breaking existing clients, and understand the real distinction between a breaking and non-breaking change — a distinction many teams get wrong in exactly the direction that quietly breaks production clients.

## Prerequisites

Lesson 1-3 (requirements, resource modeling, OpenAPI — versioning is fundamentally about managing changes to the contract those lessons establish), Lesson 2's failure drill already previewed why resource-model changes are breaking changes specifically.

## Learn

**What actually counts as a breaking change, precisely, not by intuition.** A breaking change is one that could cause a well-behaved, previously-working client to fail or behave incorrectly without any change on their end. This is more subtle than "did I change something" — removing a field is breaking (a client reading it will get an error or a missing value), renaming a field is breaking (same reason), changing a field's type is breaking (a client's deserialization may fail), adding a new *required* field to a request is breaking (existing clients aren't sending it). **Adding a new optional field to a response is generally non-breaking** — a well-behaved client should ignore fields it doesn't recognize (this is a design principle worth stating explicitly to client developers, not just assuming) — and adding a new endpoint entirely is non-breaking, since no existing client is calling it.

**Why "just add a field, it's harmless" isn't always true.** If a client's deserialization is strict (e.g. some JSON libraries or statically-typed deserializers reject unknown fields by default rather than ignoring them), even adding a field can break a poorly-configured client — worth knowing as a real edge case, though the general principle (well-behaved clients ignore unknown fields) still holds as the standard, correct client-side practice to design for and document as an expectation.

**Versioning strategies, and their tradeoffs.** URL-based versioning (`/v1/clinics`, `/v2/clinics`) is the most explicit and easiest for clients to understand and pin to, at the cost of needing to maintain multiple full route trees in parallel during a transition period. Header-based versioning (`Accept: application/vnd.api+json;version=2`) keeps URLs stable but is less discoverable/obvious to API consumers browsing documentation or debugging with simple tools. Whichever strategy you choose, the underlying discipline matters more than the specific mechanism: a clear policy for what counts as breaking (per the precise definition above), and a clear deprecation process (announcing an old version's end-of-life with enough lead time for clients to migrate, rather than removing it abruptly).

**Why breaking changes are sometimes still necessary, and how to handle them responsibly when they are.** Not every design mistake can be fixed non-breakingly forever — sometimes a genuinely wrong resource model (Lesson 2's failure drill) needs to change. The responsible approach: introduce the change as a new version, run both versions in parallel for a defined deprecation window, communicate clearly to known API consumers, and only remove the old version after that window — never simply changing the existing, already-depended-upon behavior in place, which silently breaks every client still using it with no warning.

## Attempt

1. Take one resource from your Lesson 2/3 work and design a "v2" of its response schema that includes a genuinely breaking change (e.g. renaming a field, or splitting one field into two) — write out both the v1 and v2 schemas explicitly side by side.

2. Implement URL-based versioning for this resource: `/v1/resource/{id}` continuing to serve the original schema, `/v2/resource/{id}` serving the new one, both running simultaneously against the same underlying data (adapting/transforming as needed so both versions correctly represent the same underlying state, just shaped differently).

3. Write a simple v1 client (a small script or test) that continues to work correctly and unmodified against `/v1/resource/{id}` even after your v2 endpoint exists — confirming the v1 contract genuinely hasn't changed for existing consumers, not just that a new v2 was added alongside it.

4. Add a new *optional* field to your v1 response (not a breaking change, per Learn) and confirm your existing v1 test client from step 3 still passes without modification — directly demonstrating the non-breaking case in contrast to step 1's deliberately breaking one, using the same testing approach for both so the comparison is apples-to-apples.

## Verify

Show your v1 and v2 schemas side by side highlighting the specific breaking difference, your working v1 client test passing against both the original v1 endpoint and the field-added v1 endpoint from step 4 (same client code, unmodified, both times), and confirm v2 correctly serves its new shape from the same underlying data source as v1.

## Failure drill

Take your step 1 breaking change and, instead of introducing it as a new v2 endpoint, modify the *existing* v1 endpoint in place to return the new, breaking shape directly (simulating the mistake of "just fixing" a v1 endpoint without going through a proper versioning process). Rerun your step 3 v1 client test (unmodified) against this now-changed "v1" endpoint and confirm it fails — directly reproducing, with your own test evidence, exactly the kind of silent breakage a real API's existing clients would experience if a breaking change were deployed without a proper versioning strategy. Explain why this failure, caught immediately by your own test in this exercise, would instead surface in production as a real client's application breaking unexpectedly, likely well after the change was deployed and by someone who has no visibility into what changed on the server side — the entire point of the versioning discipline in this lesson is converting this kind of silent, delayed, hard-to-diagnose production failure into an explicit, planned, and safely staged migration instead.

## Transfer

If TARDOC's or Mahall's API has ever needed a change that would have been breaking under this lesson's precise definition, describe what actually happened — was it introduced as a new version with a deprecation window, or applied in place (a reasonable thing to have done if there were no external clients yet depending on the API, worth noting honestly if that was the actual situation) — and state, for your current or near-future API surface, at what point (how many real external consumers, what business relationship) a breaking change would need this lesson's full versioning discipline rather than being safe to apply in place.

## Done when

You've precisely distinguished a breaking from a non-breaking change using this lesson's definition (not intuition) for two real examples, you've implemented working URL-based versioning where an old client continues functioning unmodified against v1 even as v2 exists, and you've directly reproduced — with your own failing test — what happens to an existing client when a breaking change is deployed without going through proper versioning, connecting that concrete, immediate test failure to what a real, delayed, hard-to-diagnose production incident would look like.
