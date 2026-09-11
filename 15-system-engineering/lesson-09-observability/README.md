# Lesson 9: Observability

## Objective

Implement the three pillars of observability — logs, metrics, and traces — in a way that lets you answer "why is this specific request slow/failing" after the fact, rather than only knowing *that* something is wrong in aggregate.

## Prerequisites

Linux tools Lesson 5 (measurement discipline — "measure before optimizing," now applied continuously in production rather than as a one-off profiling session), networking Lesson 7 (packet-level debugging — traces are the application-level analog of that lesson's request-decomposition idea).

## Learn

**Why "observability" is a distinct concept from "monitoring," worth the precise distinction.** Monitoring traditionally means watching a predefined set of known metrics/dashboards for known failure signatures — useful, but limited to problems you anticipated in advance. Observability specifically means having enough raw, detailed data (logs, metrics, traces) that you can answer *new, previously unanticipated* questions about your system's behavior after an incident, without having had to predict that exact question in advance — the difference between "I have a dashboard showing error rate" (monitoring) and "I can trace exactly which specific database query in which specific request caused this specific user's 3-second delay yesterday at 2:47pm" (observability).

**Logs: discrete, timestamped events, and why structure matters.** Unstructured log lines (`fmt.Println("processing request")`) are hard to search, filter, or aggregate at scale. Structured logs (JSON, or a structured logging library) with consistent fields (request ID, timestamp, severity, relevant context) let you query logs the way you'd query a database — "show me all errors for this specific request ID across every service it touched" — which is essentially impossible with unstructured text logs beyond simple grep-based searching.

**Metrics: aggregated, numeric measurements over time.** Counters (monotonically increasing, e.g. total requests served), gauges (a current value that can go up or down, e.g. current queue depth), and histograms (distribution of values, e.g. request latency — letting you compute p50/p95/p99, statistics track Lesson 1's point about percentiles mattering more than averages for latency, now operationalized). Metrics are cheap to store and query at scale (aggregated, not per-event), making them the right tool for dashboards and alerting — but they lose per-request detail, which is exactly what traces exist to preserve.

**Traces: following one request's full journey across every component it touches.** A trace assigns a unique ID to a request when it enters the system, and propagates that ID through every downstream call (every log line, every database query, every external API call related to that request tags itself with the trace ID) — letting you reconstruct, after the fact, the complete timeline of one specific request: which services it touched, how long each step took, and where the time actually went, directly extending networking Lesson 7's packet-level request-decomposition idea to the application layer, across possibly many services rather than one client-server hop.

**Why all three together, not just one.** Metrics tell you *something* is wrong (error rate spiked, p99 latency jumped) — an aggregate signal. Traces let you find *a specific example* of the problem and see its full journey. Logs, correlated via the trace ID, give you the detailed context at each step of that specific journey (what query ran, what parameters, what error message). Losing any one of the three leaves a real gap: metrics alone tell you something's wrong but not why; traces alone give you individual examples but no sense of scale/frequency; logs alone (especially unstructured, uncorrelated ones) are hard to connect into a coherent story across a multi-step request.

## Attempt

1. Implement structured logging for a real endpoint (JSON-formatted, with at minimum a request ID, timestamp, and severity level on every log line), and generate a request ID at the start of each request, propagating it through every subsequent log line for that request (via context, in Go's idiomatic style, or an equivalent mechanism in your language).

2. Implement basic metrics: a counter for total requests, a counter for errors, and a histogram for request latency (using a metrics library like Prometheus's client library, or a simple hand-rolled equivalent for this exercise). Expose them via a `/metrics` endpoint and confirm you can compute p50/p95/p99 latency from your histogram data.

3. Implement basic tracing: propagate a trace ID across at least 2 "hops" (e.g. an API handler calling a simulated downstream service, or a real database query) with each hop recording its own start/end timestamp tagged with the shared trace ID, and reconstruct, from your logged/traced data, a timeline showing exactly how long each hop took for one specific, real request.

4. Simulate a specific, hard-to-anticipate problem (e.g. a slow database query that only happens for a specific subset of requests matching some condition you didn't originally think to monitor for) and demonstrate that you can diagnose it using your logs/traces — find one specific slow request via your latency histogram (step 2's p99 tail), pull its trace ID, and use that trace ID to find the specific slow hop and its associated log context, reconstructing the full "why was this one slow" story after the fact, without having had a pre-built dashboard specifically for this exact problem.

## Verify

Present the actual reconstructed trace/log timeline from step 4 for one specific slow request, showing the specific hop where the time was spent and the associated structured log context explaining why (e.g. what query ran, what parameters), demonstrating the full observability chain (metric flagged it → trace located it → logs explained it) working end to end on your own data.

## Failure drill

Attempt the same step 4 diagnosis using *only* unstructured, non-request-correlated logs (revert to plain `fmt.Println`-style logging with no request ID propagation, mixing log lines from many concurrent requests together with no way to tell which lines belong to which request). Confirm you genuinely cannot reliably reconstruct which log lines correspond to your specific slow request once multiple requests are being processed concurrently — the log lines are real and present, but uncorrelated, making them practically useless for this specific diagnostic task. Explain, using this direct, hands-on comparison, why request-ID correlation (the specific thing structured logging plus tracing provides) is what actually makes logs useful for this class of diagnosis, not just "having logs" in some general sense.

## Transfer

If TARDOC's Celery task processing or API currently logs in an unstructured way (or doesn't propagate a request/task ID consistently across related log lines), describe what specific past debugging difficulty (a real one, if you can recall a specific incident, or a plausible one given the system's actual complexity) this lesson's structured-logging-plus-tracing approach would have made faster or possible where it currently isn't, and what the minimum first step toward this lesson's observability model would be for your actual codebase (likely: adopting a structured logging library and propagating a request ID, before tackling full distributed tracing).

## Done when

You've implemented structured, correlated logging and can filter to exactly one request's log lines using its request ID, you've implemented metrics sufficient to identify a p99-tail slow request, you've implemented basic tracing sufficient to reconstruct a multi-hop request's timeline, and you've directly demonstrated — via the failure drill's uncorrelated-logs comparison — why request correlation specifically, not just "having logs," is what makes this diagnostic workflow actually possible.
