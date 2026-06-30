# Project 14 - API Verification And Operations

## Goal

Verify the API contract at unit, HTTP, integration, and end-to-end boundaries,
then make production failures observable.

## Test Layers

Unit:

- service rules
- cursor encoding
- authorization decisions

HTTP:

- method/path/status/header/body
- malformed and oversized input
- error mapping

Integration:

- PostgreSQL constraints
- transactions
- query deadlines
- migrations

End-to-end:

- authenticated workflow
- duplicate request
- concurrent update
- dependency outage
- restart/recovery

## Operational Controls

- request IDs
- structured completion logs
- request rate
- 4xx/5xx rates
- p50/p95/p99 latency
- active requests
- database pool saturation
- liveness/readiness
- graceful shutdown

## Checkpoints

1. complete `test-plan.md`.
2. create a contract test per endpoint.
3. inject one service error per mapping.
4. stop PostgreSQL during a request.
5. hold a lock beyond query timeout.
6. send concurrent duplicate requests.
7. stop during an in-flight request.
8. connect response request ID to logs.

## Done Means

Passed controls, confirmed defects, and untested risks are reported separately.

