# Project 12 - Architecture Decisions

## Goal

Compare synchronous and asynchronous notification delivery using measured
latency, failure behavior, and operating cost.

## Experiment

Synchronous:

```text
create task
wait 500 ms for provider
return
```

Asynchronous:

```text
create task and durable outbox entry
return
worker calls provider
```

## Checkpoints

1. Implement the synchronous path.
2. Measure 20 request latencies.
3. Force provider failure.
4. Implement the outbox path.
5. Measure response and delivery latency separately.
6. Force five provider failures.
7. Observe backlog and retry behavior.
8. Complete `ADR.md`.

## Decision Inputs

- API p95 target
- notification delivery target
- acceptable duplicate behavior
- maximum backlog age
- worker and queue operating cost
- recovery from provider outage

## Failure Drill

Stop the worker for five minutes while requests continue. Define the bound that
prevents unlimited backlog growth.

## Done Means

The ADR cites measured evidence and states what future evidence would reverse
the decision.

