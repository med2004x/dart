# Project 14 - Systems Design Capstone

## Goal

Design and implement one critical vertical slice with its most important failure
path.

## Choose One

- URL shortener
- report generator
- inventory reservation
- idempotent payment request
- notification platform

Do not combine all five.

## Required Deliverables

1. `requirements.md`
2. `capacity.md`
3. `diagram.txt`
4. `contracts.md`
5. `data-model.md`
6. `failure-modes.md`
7. `security.md`
8. `observability.md`
9. `deployment.md`
10. runnable implementation and tests

## Required Runtime Controls

- startup configuration validation
- input validation
- authentication and authorization where applicable
- timeouts on external calls
- transaction for multi-step writes
- structured logs with request ID
- request rate, errors, latency, and saturation signals
- liveness and readiness
- graceful shutdown
- bounded overload behavior

## Required Tests

- happy path
- boundary inputs
- missing resource
- unauthorized access
- concurrent conflict
- duplicate request
- dependency timeout
- storage failure and rollback
- graceful shutdown

## Failure Review

For each failed test, write:

```text
root cause is X at boundary Y because Z
```

## Done Means

The running implementation matches the written contracts and failure table.

