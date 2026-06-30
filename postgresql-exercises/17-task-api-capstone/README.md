# Project 17 - PostgreSQL Task API Capstone

## Goal

Replace file persistence in the Go task capstone with a production-shaped
PostgreSQL repository and prove the complete request path.

## Architecture

```text
HTTP handler
-> TaskService
-> TaskRepository interface
<- PostgresTaskRepository
-> database/sql pool
-> PostgreSQL
```

Business logic imports neither HTTP nor SQL.

## Required Structure

```text
cmd/api/main.go
internal/task/model.go
internal/task/service.go
internal/task/repository.go
internal/postgres/task_repository.go
internal/httpapi/handler.go
migrations/
```

Do not create packages solely to match this diagram. Keep ownership clear and
avoid circular dependencies.

## Required Database Behavior

- versioned migrations
- named constraints
- parameterized SQL
- bounded query contexts
- transactions for multi-step writes
- configured and observed pool
- readiness reflects required database access
- graceful pool close
- tested backup/restore

## Required API Behavior

- project-scoped task CRUD
- ownership authorization
- deterministic pagination
- stable JSON errors
- 400, 401, documented 403/404, 409, and generic 500 mappings
- request IDs and structured completion logs

## Test Layers

Unit:

- service validation
- authorization
- error mapping

Integration:

- every repository method
- constraints
- transaction rollback
- concurrency
- migrations
- query timeout

End-to-end:

- authenticated CRUD
- pagination
- database unavailable
- pool saturation
- graceful shutdown
- restored database

## Failure Drills

1. stop PostgreSQL during a request.
2. hold a row lock beyond query timeout.
3. send duplicate concurrent inserts.
4. deploy an incompatible migration with old code running.
5. exhaust a deliberately small pool.
6. restore backup and run the API against it.

## Done Means

You can trace and prove:

```text
request
-> authorization
-> service rule
-> repository query/transaction
-> PostgreSQL constraint/commit
-> response and telemetry
```

