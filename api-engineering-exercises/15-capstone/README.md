# Project 15 - API Engineering Capstone

## Goal

Build a project/task API whose contract, persistence, security, duplicate
behavior, compatibility, and operational controls are proven.

## Required Operations

- project create/read/list
- project membership management
- task create/read/list/update/delete
- task assignment
- task completion event webhook
- asynchronous project report

## Required Engineering

- OpenAPI contract
- separate request/response schemas
- PostgreSQL constraints and migrations
- stable errors
- bounded bodies and page sizes
- cursor pagination
- create idempotency
- ETag update preconditions
- authentication and resource authorization
- rate and concurrency limits
- signed webhooks
- asynchronous job lifecycle
- request telemetry
- graceful shutdown
- backup/restore

## Deliverables

```text
openapi.yaml
migrations/
cmd/api/
internal/domain/
internal/service/
internal/postgres/
internal/httpapi/
tests/
runbook.md
compatibility.md
```

## Mandatory Failure Drills

1. repeated create after lost response
2. two stale concurrent updates
3. cross-project task ID
4. database timeout
5. pool saturation
6. webhook receiver failure
7. worker restart
8. migration with old server running
9. shutdown during request
10. restore into separate database

## Done Means

Every endpoint and background flow has a contract test, ownership check,
failure policy, telemetry signal, and recovery procedure.

