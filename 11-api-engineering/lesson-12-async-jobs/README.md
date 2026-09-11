# Lesson 12: Async Jobs

## Objective

Design an API pattern for operations too slow to complete within a single request/response cycle — accept-and-poll or accept-and-webhook — and understand why forcing a client to wait synchronously for a long-running operation is often the wrong design, not just a performance inconvenience.

## Prerequisites

Lesson 11 (webhooks — one of the two standard patterns this lesson covers for notifying completion), distributed Lesson 7 (queues — the underlying execution mechanism for the actual long-running work, conceptually).

## Learn

**Why synchronous request/response breaks down for slow operations.** An HTTP request held open while a server does genuinely slow work (e.g. TARDOC's transcription processing, mentioned throughout your project history) ties up a connection, risks client-side or intermediate-proxy timeouts (many HTTP clients and load balancers have default timeouts in the 30-60 second range, sometimes shorter), and gives the client no way to do anything else while waiting — if the operation could reasonably take minutes, forcing a client to hold a connection open that whole time is both fragile (any transient network blip during that window fails the entire operation) and a poor use of server-side connection resources at any real scale.

**The accept-and-poll pattern.** The client submits the request; the server immediately responds with `202 Accepted` and a URL for a **job resource** representing the in-progress operation (e.g. `Location: /jobs/{jobId}`), rather than the actual result. The client then polls `GET /jobs/{jobId}` periodically, receiving a status (`pending`, `processing`, `completed`, `failed`) and, once completed, either the result directly or a link to where it can be retrieved. This decouples job submission from result retrieval entirely — the client can disconnect and reconnect, poll from a different process, or simply check back later, none of which is possible with a long-held synchronous connection.

**The accept-and-webhook pattern.** Same initial `202 Accepted` and job resource, but instead of (or in addition to) polling, the client registers a callback URL (Lesson 11's webhook mechanism) to be notified when the job completes — appropriate when polling overhead is undesirable (e.g. very long-running jobs where frequent polling would be wasteful) or when the client architecture is itself event-driven rather than poll-based. This is strictly more complex to implement correctly (requiring all of Lesson 11's signing/retry considerations) but avoids the inherent latency and resource cost of repeated polling.

**Job status modeling, and why it needs to be a genuine state machine, not just a boolean "done" flag.** A job resource needs at least: pending (accepted, not yet started), processing (actively running), completed (successful, with a result available), failed (terminated with an error, with error detail available) — and the valid transitions between these states should be enforced (e.g. a job cannot go from `completed` back to `processing`), because client code polling for status needs to reliably distinguish "still working, keep waiting" from "done, here's your answer" from "failed, here's why, don't keep waiting" — collapsing these into a simpler "done: true/false" loses exactly the information a client needs to behave correctly in the failure case.

## Attempt

1. Implement the accept-and-poll pattern: a `POST /jobs` endpoint that accepts a request representing slow work (simulate the actual slowness with a `sleep()` in a background goroutine, rather than genuinely slow processing, so testing remains fast), immediately returns `202 Accepted` with a `Location` header pointing to `/jobs/{jobId}`, and processes the work asynchronously (in a separate goroutine, not blocking the response).

2. Implement `GET /jobs/{jobId}` returning the job's current status (`pending`/`processing`/`completed`/`failed`) and, once completed, the result. Write a client that polls this endpoint at a reasonable interval (e.g. every 500ms) until it observes `completed` or `failed`, and confirm it correctly retrieves the eventual result without ever needing to hold open the original submission request.

3. Implement failure handling: have your simulated background work occasionally (or deliberately, for one test case) fail, and confirm the job resource correctly transitions to `failed` with an appropriate error detail, distinct from `completed`, and that your polling client correctly distinguishes this case and stops polling rather than waiting indefinitely for a `completed` state that will never arrive.

4. Extend the same job system with a webhook option (Lesson 11): allow the initial `POST /jobs` request to optionally include a callback URL, and upon job completion (or failure), fire a signed webhook notification to that URL instead of (or in addition to) relying on polling. Test both the poll-only and webhook-notified paths against the same underlying job execution, confirming both correctly observe the eventual outcome.

## Verify

For step 3, show the actual job resource's state transitions over time for both a successful and a failing job (a sequence of `GET /jobs/{jobId}` responses showing `pending` → `processing` → `completed`/`failed`), confirming the state machine behaves correctly and the polling client correctly stops in both terminal cases.

## Failure drill

Modify your job system to allow an invalid state transition — specifically, allow a `completed` job to be reprocessed and transitioned back to `processing` (e.g. by not guarding against calling your "start processing" logic on a job that's already terminal). Construct a scenario where a delayed or duplicate trigger (analogous to distributed Lesson 7's at-least-once redelivery) causes the same job to be processed twice, and observe the job's result potentially changing or its status flickering between states in a way that would confuse a polling client that had already observed `completed` and stopped polling, now silently missing that it later reverted to `processing`. Explain why guarding state transitions explicitly (only allowing `pending`→`processing`→{`completed`,`failed`}, with terminal states genuinely terminal) is necessary for job status to be a reliable contract a client can safely stop watching once given a terminal state — directly connecting back to distributed Lesson 7's idempotency lesson, since this is the same underlying "at-least-once processing needs safeguards" problem in a new form.

## Transfer

If TARDOC's transcription pipeline is triggered via an API call, describe, using this lesson's job-resource pattern, whether it currently returns synchronously (holding the request open until transcription completes) or asynchronously (accept-and-poll or accept-and-webhook) — and if synchronous, describe specifically what would need to change (given the Groq-hosted Whisper processing and multi-key rotation mentioned in your project history, which suggests real, possibly variable processing latency) to migrate it to the async pattern this lesson covers, and what client-facing behavior change that migration would require callers to adapt to.

## Done when

Your accept-and-poll implementation correctly decouples job submission from result retrieval, your job state machine correctly distinguishes and reports success versus failure (not just a boolean done flag), you've extended it with a working webhook-notification alternative to polling, and you've demonstrated — via the failure drill — why unguarded state transitions on a job resource create a real reliability problem for any client relying on terminal states being genuinely final.
