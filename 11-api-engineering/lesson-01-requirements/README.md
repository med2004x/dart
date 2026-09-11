# Lesson 1: Requirements and API Design as a Contract

## Objective

Treat an API as a contract negotiated *before* implementation, not a byproduct of whatever the handler code happens to do — write down the request/response shape and behavioral guarantees first, then build code that matches, rather than the reverse.

## Prerequisites

Networking Lesson 4 (HTTP semantics — this track builds directly on request/response mechanics already covered there).

## Learn

**Why "contract first" matters, concretely, not just as good practice.** An API is a promise to every client that calls it — mobile apps, other backend services, third-party integrators — none of whom can see your implementation, only the request/response shape and the documented behavior. If you build the handler first and let the API shape emerge from whatever was convenient to implement, you tend to leak implementation details into the contract (e.g. an internal database column name showing up as a JSON field, or an error message format that changes whenever you refactor internal error handling) — clients then depend on those accidental details, and changing your implementation later silently becomes a breaking change for everyone depending on it, even though nothing about the *intended* contract changed.

**What "requirements" means at the API level, specifically.** Before writing a single handler, you should be able to answer, for each operation: what resource does this act on (Lesson 2 covers resource modeling properly), what does the client send, what does the client get back on success, what does the client get back on each distinct failure mode, and what side effects (if any) does this operation have that a client might reasonably need to know about (e.g. does it send an email, charge a payment, trigger a webhook). Writing these down *before* implementing is what Lesson 3 formalizes into an OpenAPI contract — this lesson is about the thinking that contract-writing requires, before worrying about the specific format.

**Distinguishing what the client needs to know from what's an implementation detail.** A genuinely common design mistake: exposing something in the API response because it happened to be convenient to include (e.g. an internal processing duration, a database-generated timestamp with implementation-specific precision, an internal status enum value that doesn't map cleanly to anything the client should reason about) rather than because a client actually needs it. Every field in a response is a promise you're making to maintain — the fewer accidental, implementation-driven fields you expose, the more freedom you retain to change your implementation later without breaking anyone.

**Failure modes are part of the contract too, not an afterthought.** A client needs to know, for a given operation, not just "it can fail" but specifically *how* it can fail and what each failure means for what they should do next (retry? fix their request? contact support?) — Lesson 4 and Lesson 5 formalize this further, but the discipline starts here: writing down the failure cases as part of requirements gathering, not discovering them ad hoc while writing `if` statements during implementation.

## Attempt

1. Pick a real feature from TARDOC or Mahall that involves an API endpoint you've already built (or plan to build) — for example, TARDOC's clinic subscription activation, or Mahall's product listing creation. Before looking at (or without regard to) your existing implementation, write a plain-language requirements document answering: what resource is being acted on, what does the client send, what's returned on success, and what are at least 3 distinct ways this specific operation could fail, with a plain-English description of what each failure means.

2. For the same operation, explicitly list every field you'd include in the success response, and for each field, justify in one sentence why the client actually needs it (not "it was easy to include" or "it's just there internally") — if you can't justify a field this way, mark it as a candidate for removal from the contract.

3. Compare your step 1-2 requirements document against your actual existing implementation (if one exists) for the same operation. Identify at least one place where the implementation exposes something that isn't clearly justified as something the client needs (an implementation detail that leaked into the contract), and describe what a cleaner version of the response would look like instead.

4. Write the requirements document for one operation you *haven't* built yet, following the same process, and treat it as the actual specification you'd hand to yourself (or a collaborator) before writing any code — this is the practice this lesson is building: requirements-first, not implementation-first, becoming your default working method rather than a one-off exercise.

## Verify

For step 3, show the specific field(s) you identified as implementation leakage, alongside your reasoning for why a client shouldn't need to depend on it, and what you'd change.

## Failure drill

Take your step 4 requirements document (for an operation you haven't built) and deliberately have someone else — or, if working alone, revisit it after a meaningful gap (a different day) with fresh eyes — attempt to implement a mock/stub version of the endpoint using *only* your written requirements, without asking you clarifying questions. Note every point where the requirements document turned out to be ambiguous or incomplete, requiring an assumption to be made that wasn't actually specified. Explain why this exercise — not just writing requirements, but testing whether they're actually sufficient for someone else to implement against without guessing — is the real test of whether "contract-first" thinking actually happened, versus just writing documentation that still leaves the real design decisions implicit.

## Transfer

If Lead Sourcer's pipeline (mentioned in your project history) exposes or will expose any API surface to other tools or services, describe how you'd apply this lesson's process to that surface specifically — what resource(s) it's acting on, and at least 2 concrete failure modes a caller of that API would need to know how to handle, distinct from whatever internal error handling Lead Sourcer's own scraping/scoring logic already does.

## Done when

You've written a complete requirements document for a real operation before referring to (or instead of referring to) its existing implementation, you've identified at least one real instance of implementation detail leaking into an existing API contract in your own code, and you've stress-tested a requirements document by having it (or a fresh read of it) actually attempted for implementation without your ongoing clarification, surfacing real gaps rather than assuming the document was complete.
