# Project 16 - Go PostgreSQL Repository

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Use PostgreSQL through one shared `database/sql` pool with contexts,
parameterized queries, complete row handling, and explicit transactions.

## Setup

```powershell
Set-Location .\16-go-repository
$env:DATABASE_URL = "postgres://task_app:local-app-password@localhost:5432/go_course?sslmode=disable"
go mod tidy
go run .
```

`sslmode=disable` is local-only. Production requires an explicit TLS policy.

## Checkpoints

1. open one shared pool.
2. validate connectivity with `PingContext`.
3. configure pool bounds.
4. implement `Create`.
5. implement `FindByID`.
6. implement `ListByProject`.
7. map `sql.ErrNoRows`.
8. close rows and check `rows.Err`.
9. add query deadlines.
10. add integration tests against a dedicated test database.

## Transaction Extension

Implement the account transfer from project 09 using `sql.Tx`. Every statement
must use `tx`, not `db`.

## Failure Drills

1. wrong password.
2. `SELECT pg_sleep(10)` with a short context.
3. unclosed rows with a small pool.
4. string-concatenated user input.
5. zero affected rows ignored.
6. one transaction statement sent through `db`.

## Done Means

The repository is the only layer containing SQL, and real integration tests
prove database behavior.

