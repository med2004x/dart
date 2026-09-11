# Lesson 2: Resource Modeling

## Objective

Model API resources (nouns, not actions) with correct hierarchy and relationships, and understand why REST's resource-oriented design, done well, produces URLs and operations that are predictable across an entire API rather than each endpoint being independently invented.

## Prerequisites

Lesson 1 (requirements — resource modeling is the structural backbone the requirements you wrote there need to map onto).

## Learn

**Resources are nouns; HTTP methods are the verbs.** A REST API models the domain as resources (e.g. `clinics`, `subscriptions`, `products`) and uses HTTP methods to express operations on them (`GET /clinics/{id}` — retrieve, `POST /clinics` — create, `PATCH /clinics/{id}` — partial update, `DELETE /clinics/{id}` — remove) rather than encoding the verb into the URL itself (`/getClinic`, `/createClinic`). This isn't just stylistic convention — it's what makes an API's *shape* predictable: once a client understands your resource model, they can reasonably guess how to interact with a resource type they haven't used yet, because the same verb-to-method mapping applies consistently across every resource, rather than needing separate documentation for each endpoint's idiosyncratic verb-in-URL naming.

**Collection vs. individual resource, and the URL pattern this implies.** `/clinics` represents the *collection* of clinics (GET lists them, POST creates a new one within it); `/clinics/{id}` represents one *specific* clinic (GET retrieves it, PATCH/PUT updates it, DELETE removes it). This two-level pattern — collection URL, individual-resource URL nested under it — is the standard shape, and deviating from it without a strong reason (e.g. inventing `/clinic-list` as a separate URL from `/clinics`) tends to confuse clients who've internalized the standard pattern from every other well-designed REST API they've used.

**Nested resources: modeling genuine ownership/containment relationships.** `/clinics/{id}/subscriptions` expresses "the subscriptions belonging to this specific clinic" — appropriate when the nested resource genuinely doesn't make sense (or isn't independently addressable) outside its parent's context. This is a real design decision, not automatic: a resource that *can* meaningfully exist and be referenced independently of any specific parent (e.g. a `product` that could belong to a `seller` but is still independently useful to fetch by its own ID without needing the seller's ID as context) is usually better modeled as a top-level resource with a reference field to its owner, rather than forced into a nested URL — nesting should reflect genuine structural containment, not just "this happens to relate to that."

**Why over- or under-nesting both cause real problems.** Over-nesting (e.g. `/clinics/{clinicId}/subscriptions/{subId}/payments/{paymentId}/receipts/{receiptId}`) produces unwieldy URLs and forces every client to always know the full ancestor chain even when they only care about the deepest resource — a genuine usability cost. Under-nesting (modeling everything as flat, unrelated top-level resources even when a real containment relationship exists) loses the ability to naturally express "give me all of X belonging to Y" as a single collection URL, forcing clients to filter client-side or via query parameters instead of a natural nested-collection request.

## Attempt

1. For TARDOC's domain (clinics, subscriptions, billing records, possibly transcription jobs — adjust to your actual domain model), sketch a resource model: list each resource type, and for each, state whether it can exist and be meaningfully referenced independently, or whether it's genuinely only meaningful in the context of a specific parent resource.

2. Based on step 1, design the URL structure: which resources are top-level (`/resource`), and which are nested under a parent (`/parent/{id}/resource`) — justify each nesting decision using the "genuine containment vs. independently addressable" distinction from Learn, not just "it seemed related."

3. For one resource in your model, write out the full set of standard operations it should support (list, get one, create, update, delete — not every resource needs all five; e.g. some might be read-only, or creation might only ever happen as a side effect of another operation, worth noting explicitly if so) and their corresponding HTTP method + URL pairs.

4. Deliberately design a bad, over-nested version of one part of your model (e.g. force something naturally independent into 3+ levels of unnecessary nesting) and a bad, under-nested version of another part (flatten something that has genuine containment structure into an unrelated top-level resource with no connection expressed in the URL). Write one sentence for each explaining specifically what a client would find awkward or confusing about interacting with it.

## Verify

Present your final resource model (step 1-3) as a table: resource name, top-level or nested-under-what, and the operations it supports with their HTTP method + URL — this table should be complete enough that someone unfamiliar with your domain could infer the correct URL for a new operation by following the established pattern.

## Failure drill

Take your step 4 over-nested bad example and trace through what happens to every client currently depending on it if you later decide to correct the design (move the resource to be top-level instead of deeply nested) — specifically, note that this is a breaking URL change, affecting every existing client, and connect this forward to Lesson 13 (versioning), which will cover how such changes are normally handled without breaking existing clients outright. Explain, in your own words, why getting resource modeling right *before* clients depend on your API (this lesson's whole point) is considerably cheaper than fixing a bad model after the fact, given that a URL structure change is a breaking change in a way that, say, adding a new optional field usually is not.

## Transfer

If Mahall's storefront API models products, sellers, and orders, describe explicitly, using this lesson's containment criteria, whether `orders` should be nested under `/sellers/{id}/orders`, under `/products/{id}/orders`, as a top-level `/orders` resource with reference fields to both seller and product, or some combination — and justify your choice using the genuine-relationship reasoning from Learn rather than an arbitrary preference, noting explicitly if an order can reasonably be said to have exactly one clear "owning" parent or whether it's better modeled with references to multiple related resources instead of forced into one nesting hierarchy.

## Done when

You've produced a complete resource model for a real domain with justified top-level/nested decisions (not arbitrary ones), you've deliberately constructed and critiqued both an over-nested and under-nested bad example, and you can explain why resource modeling decisions are expensive to change later, specifically because of their direct relationship to public URL structure that clients depend on.
