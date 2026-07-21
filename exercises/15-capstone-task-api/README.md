# Exercise 15 - Capstone Task API

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Build a complete local task API and be able to explain every boundary:

```text
client
  -> middleware
  -> handler
  -> service
  -> repository
  -> JSON file storage
```

This is "production-shaped," not production-ready. File persistence, a shared
development API key, and single-process state impose real limits.

## Beginner Bridge: Capstone Means Combining Boundaries

The previous exercises each taught one boundary:

| Earlier exercise | Boundary learned |
|---|---|
| contact book | CRUD over a collection |
| file notes | memory to JSON file |
| errors and validation | failure as a returned value |
| HTTP basics | request to response |
| JSON API | bytes to structs and back |
| middleware | policy around handlers |
| service/repository | business rules separated from storage |

This capstone combines them. That does not mean writing a giant program from
top to bottom. It means proving each boundary separately, then connecting them.

The correct build order is center-out:

```text
domain rules -> storage -> repository -> service -> handler -> middleware -> server
```

If you start by making routes, you will probably hide business rules and file
logic inside handlers. That works for a demo and becomes painful immediately
when tests, persistence failures, or auth behavior need to change.

## Do Not Start With HTTP

Build from the center outward:

1. domain model and errors
2. storage behavior and failure policy
3. repository
4. service rules
5. handler contract
6. middleware
7. server assembly
8. end-to-end verification

Starting with routes encourages business and storage logic to accumulate inside
handlers.

## Required Structure

Keep one package to avoid package ceremony:

```text
main.go        dependency assembly, config, server lifecycle
task.go        Task and domain errors
storage.go     atomic JSON file read/write
repository.go  synchronized task collection and persistence
service.go     task use cases and rules
handler.go     HTTP/JSON adapter
middleware.go  logging and API-key policy
```

Tests should sit beside the code they verify:

```text
storage_test.go
repository_test.go
service_test.go
handler_test.go
```

## Domain Contract

```go
type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"createdAt"`
}
```

Rules:

- IDs are server-owned and never reused during one repository lifetime.
- title is trimmed and must not be empty.
- `CreatedAt` is server-owned and remains unchanged during update.
- create starts with `Done == false`.
- missing IDs produce a stable `ErrTaskNotFound`.

Write these as tests before service implementation.

## API Contract

| Method | Path | Auth | Success | Common failures |
|---|---|---|---:|---|
| GET | `/health` | public | 200 | 405 |
| POST | `/tasks` | required | 201 | 400, 401 |
| GET | `/tasks` | required | 200 | 401 |
| GET | `/tasks/{id}` | required | 200 | 400, 401, 404 |
| PUT | `/tasks/{id}` | required | 200 | 400, 401, 404 |
| DELETE | `/tasks/{id}` | required | 204 | 400, 401, 404 |

Use a consistent JSON error body for API failures, for example:

```json
{"error":"title is required"}
```

Do not expose file paths, stack traces, or internal error details.

## Persistence Is Part Of The Write

A dangerous flow:

```text
mutate in-memory tasks
attempt file save
save fails
return error
```

The caller sees failure, but memory contains the change. A later successful save
may persist an operation the caller was told failed.

Use transactional thinking:

```text
lock repository
copy current tasks
apply change to candidate copy
persist candidate copy
IF persistence fails
    leave current tasks unchanged
    return error
replace current tasks with candidate copy
unlock
return success
```

The file write and memory commit form one logical operation.

## Atomic File Replacement

Writing directly to `tasks.json` can leave corrupt partial content.

Safer sequence:

```text
encode complete data
write to temporary file in same directory
close/flush temporary file
rename temporary file over target
```

Renaming within the same filesystem is commonly atomic. Windows replacement
semantics require care when the target exists, so test the exact implementation
on Windows. If replacement cannot be made reliably, document the limitation.

Loading policy:

- missing file -> empty repository
- empty valid JSON array -> empty repository
- malformed JSON -> startup failure, never silent reset
- read permission failure -> startup failure

Silent recovery from malformed data risks permanent data loss.

## Repository Ownership

The repository owns:

- synchronized in-memory tasks
- next-ID calculation
- invoking persistence on writes
- rollback-on-save-failure behavior

It must not own:

- HTTP status codes
- request bodies
- title validation policy
- API key checks

Return copies of slices so handlers cannot mutate repository state without a
write operation.

## Service Ownership

The service owns use cases:

```text
CreateTask(title)
ListTasks()
GetTask(id)
UpdateTask(id, title, done)
DeleteTask(id)
```

Example create pseudocode:

```text
trim title
IF title is empty
    return title-required error
construct task with current time and done false
ask repository to create it
return created task or wrapped repository error
```

For deterministic service tests, inject a clock:

```go
type Clock func() time.Time
```

Production assembly passes `time.Now`; tests pass a fixed function.

## Handler Ownership

Handlers:

1. verify method and parse path
2. limit and decode JSON input
3. call exactly one service use case
4. map known errors to statuses
5. encode a consistent response

Handlers do not read files, lock mutexes, allocate IDs, or validate domain rules.

Use input types that exclude server-owned fields:

```go
type createTaskInput struct {
    Title string `json:"title"`
}

type updateTaskInput struct {
    Title string `json:"title"`
    Done  bool   `json:"done"`
}
```

## Middleware And Configuration

Use:

```text
X-API-Key: dev-secret
```

for local practice. Read it from an environment variable rather than hardcoding
it in the final capstone:

```powershell
$env:TASK_API_KEY = "dev-secret"
go run .
```

Startup must fail with a clear message if the key is missing.

Never log the header value. Logging should include method, path, status, duration,
and a request ID. JSON structured logs are the production target.

## Worked Example: Persistent Bookmark Service

The same design can build:

```text
Bookmark { ID, URL, Label, CreatedAt }
```

Request path:

```text
POST /bookmarks
-> API-key middleware checks credentials
-> handler decodes URL and label
-> service validates URL
-> repository creates candidate state
-> storage atomically writes bookmarks.json
-> repository commits memory
-> handler returns 201
```

Failure path:

```text
storage write fails
-> repository discards candidate state
-> service wraps storage error
-> handler logs internal detail with request ID
-> client receives generic 500 JSON
```

That is the same architecture with different domain rules. If your task design
cannot transfer to bookmarks without rewriting every layer, the boundaries are
wrong.

## Test-First Build Order

### Milestone 1: Storage

Write failing tests for:

- missing file loads empty
- valid file loads exact values
- malformed JSON returns error
- save then load round-trips values

### Milestone 2: Repository

Write failing tests for:

- create assigns next ID
- restart derives a safe next ID
- update preserves `CreatedAt`
- missing operations return `ErrTaskNotFound`
- save failure does not mutate memory
- returned lists cannot mutate internal state

### Milestone 3: Service

Write failing tests for:

- blank create/update titles fail
- clock controls `CreatedAt`
- repository errors remain classifiable

### Milestone 4: Handlers

Use `httptest` for:

- malformed JSON
- oversized bodies
- status and content type
- stable error body
- 204 has no body

### Milestone 5: Middleware And End-To-End

Verify public health, unauthorized requests, authorized CRUD, persistence after
restart, logs, and graceful stop.

Commands:

```powershell
gofmt -w .
go test ./...
go test -race ./...
go vet ./...
go run .
```

## Manual End-To-End Run

Set configuration and start:

```powershell
$env:TASK_API_KEY = "dev-secret"
go run .
```

Second PowerShell window:

```powershell
$headers = @{ "X-API-Key" = "dev-secret" }

$created = Invoke-RestMethod `
    -Uri http://localhost:8080/tasks `
    -Method Post `
    -Headers $headers `
    -ContentType "application/json" `
    -Body (@{ title = "understand persistence" } | ConvertTo-Json)

$created
Invoke-RestMethod -Uri http://localhost:8080/tasks -Headers $headers

$updateBody = @{
    title = "understand persistence boundaries"
    done = $true
} | ConvertTo-Json

Invoke-RestMethod `
    -Uri "http://localhost:8080/tasks/$($created.id)" `
    -Method Put `
    -Headers $headers `
    -ContentType "application/json" `
    -Body $updateBody
```

Stop with `Ctrl+C`, restart, and list tasks again. Persistence is proven only if
the task survives process restart.

## Failure Drills

Do not skip these:

1. Corrupt `tasks.json`; startup must fail without replacing it.
2. Make the save path unwritable; create must fail and list must show no ghost
   task.
3. Send no key, a wrong key, and a correct key.
4. Send malformed JSON, unknown fields, blank title, and an oversized body.
5. Request nonnumeric, zero, negative, missing, and existing IDs.
6. Send concurrent creates and run the race detector where supported.
7. Stop during requests and verify graceful shutdown behavior.

## Production Readiness Gap

Do not label this production-ready. It still lacks:

- a transactional database and multi-process consistency
- real identity and authorization
- secret manager integration and key rotation
- TLS termination policy
- rate limiting
- complete structured telemetry and alerting
- deployment, migration, backup, restore, and rollback procedures

The capstone is complete when its stated local contract is tested and understood,
not when those gaps are renamed as future details.

## You Understand This Exercise When

You can:

- trace success and every failure across all layers
- prove failed persistence does not mutate visible state
- replace file storage with a database adapter without changing HTTP or business
  rules
- explain middleware order and every status code
- reproduce behavior with tests, not just one manual run
- identify exactly why this design is not yet production-ready

References:

- [`net/http`](https://pkg.go.dev/net/http)
- [`encoding/json`](https://pkg.go.dev/encoding/json)
- [`os`](https://pkg.go.dev/os)
- [`errors`](https://pkg.go.dev/errors)
- [`time`](https://pkg.go.dev/time)
- [`httptest`](https://pkg.go.dev/net/http/httptest)
