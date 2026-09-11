# Lesson 1: Request Boundaries

## Objective

Define clear boundaries for what happens within a single request's scope versus what's deferred to background processing, and understand why conflating "everything the user's action triggers" with "everything that must happen before responding" is a common source of slow, fragile endpoints.

## Prerequisites

API engineering Lesson 12 (async jobs — this lesson covers the design thinking behind *when* to reach for that pattern, rather than the pattern's mechanics themselves).

## Learn

**What "request boundary" means, precisely.** When a user action (e.g. "create an order") triggers multiple downstream effects (charge a payment, send a confirmation email, update inventory, notify a fulfillment system), the request boundary question is: which of these must complete, successfully, before the client receives a response — and which can happen afterward, asynchronously, without the client waiting? Getting this wrong in either direction causes real problems: including too much inside the boundary makes the request slow and fragile (any one slow or failing downstream step blocks the entire response, even for effects the client doesn't need immediate confirmation of); including too little risks telling the client "success" before something that should have been guaranteed (e.g. payment actually succeeding) has actually happened.

**The core question: what does the client actually need confirmed before it can safely proceed?** If a client's next action depends on knowing whether the payment succeeded (e.g. it needs to show a different UI for success vs. failure), that specific check must be inside the request boundary. If a client doesn't need to know whether the confirmation email was sent successfully before proceeding with its own next steps, that effect can safely be deferred — it happens as a consequence of the request, but not as a blocking part of *answering* the request.

**Why this connects directly to API engineering Lesson 12's async pattern, but is a distinct, earlier design decision.** Lesson 12 covered the *mechanics* of accept-and-poll/webhook once you've decided something is async. This lesson is about the design judgment that comes first: deciding, for a specific business operation, which pieces belong in the synchronous path (and therefore need to complete fast and reliably within the request) and which belong outside it (and can be handed off to background processing, per distributed Lesson 7's queue patterns). Getting this boundary right is a domain-specific decision, not a mechanical rule — "should email sending block the response" doesn't have a universal answer, it depends on whether anything about that specific flow actually needs it to.

**The cost of an incorrectly wide boundary, concretely.** If your order-creation endpoint synchronously calls the payment processor, the inventory service, the email service, and a fulfillment webhook all before responding, the endpoint's latency is the *sum* of all four calls' latencies (or worse, blocked entirely if any one hangs), and a transient failure in the *least* critical of the four (say, the email service being briefly down) causes the entire, otherwise-successful order creation to fail or appear to fail to the client — a direct, avoidable coupling between something non-critical and the success of something that should have been independent.

## Attempt

1. Take a real multi-step operation from TARDOC or Mahall (e.g. TARDOC's clinic subscription activation, which per your project history involves the mailer being wired into lifecycle events) and list every distinct downstream effect it currently triggers.

2. For each effect listed in step 1, explicitly answer: does the client need to know, before receiving a response, whether this specific effect succeeded? Justify each answer in one sentence based on what the client would actually do differently depending on the outcome.

3. Based on step 2, redesign the operation's request boundary: which effects stay synchronous (inside the request, blocking the response), and which move to asynchronous background processing (e.g. via a queue, per distributed Lesson 7, or simply an async goroutine with its own error handling and retry logic, not tied to the client's connection).

4. Implement (or sketch in enough detail to be concrete, if a full implementation isn't practical here) the redesigned version: the synchronous path handles only the effects that genuinely need synchronous confirmation, and the asynchronous effects are dispatched to background processing with their own independent error handling — critically, a failure in an asynchronous effect must not cause the synchronous response to report failure, since the client already received a valid "the core operation succeeded" response by that point.

## Verify

Present your step 2 table (effect, does-client-need-to-know-synchronously, justification) and your step 3 redesigned boundary, and estimate the latency improvement your redesign would provide by comparing the sum of all four original synchronous calls' typical latencies against just the synchronous subset's latency in your redesign.

## Failure drill

Take your original, wide-boundary design (everything synchronous) and simulate one of the genuinely-non-critical effects (per your own step 2 analysis) failing or timing out — e.g. the email service being unreachable. Confirm the entire operation fails or hangs as a result, even though the core operation (e.g. the order/subscription itself) would have succeeded fine on its own. Then run the same failure scenario against your redesigned, narrower-boundary version and confirm the core operation now succeeds and responds promptly, with the email failure handled independently (logged, retried in the background, or otherwise surfaced without blocking the client). Explain, using this concrete before/after comparison, why an overly wide request boundary doesn't just cost latency in the average case, it actively creates unnecessary failure coupling between unrelated concerns.

## Transfer

For a second real operation in your own systems (different from the one used in steps 1-4), apply the same "does the client need synchronous confirmation of this effect" question to at least 3 downstream effects, and identify whether the current implementation's actual request boundary matches what your analysis suggests it should be — if there's a mismatch, describe specifically what's currently synchronous that your analysis suggests should be async, or vice versa.

## Done when

You've analyzed a real multi-effect operation and produced a justified, effect-by-effect boundary decision (not just an intuition), you've demonstrated concretely — with a real before/after failure simulation — that a poorly-drawn boundary creates unnecessary failure coupling, and you can explain the general principle (does the client need to know, synchronously, whether this specific effect succeeded) well enough to apply it to a new operation you haven't analyzed before.
