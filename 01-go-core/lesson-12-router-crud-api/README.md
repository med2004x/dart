# Exercise 12 - Router And CRUD API

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

This exercise combines resource design, route parsing, JSON, and mutable state.
The hard part is not writing five handlers. It is making every method/path pair
have one unambiguous behavior.

## Beginner Bridge: CRUD Becomes HTTP Routes

Exercise 05 taught CRUD in memory. This exercise exposes the same idea over
HTTP.

| CRUD idea | In-memory contact book | HTTP task API |
|---|---|---|
| Create | `addContact` | `POST /tasks` |
| Read many | `printContacts` | `GET /tasks` |
| Read one | `findContactByID` | `GET /tasks/{id}` |
| Update | `updateContactEmail` | `PUT /tasks/{id}` |
| Delete | `deleteContact` | `DELETE /tasks/{id}` |

The new difficulty is that bad input arrives as text over the network. The path
`/tasks/abc` is not an integer ID. A missing task is different from a malformed
ID. A method can be wrong even when the path is right.

Do not start with handlers. Start with the route table. Every row in the table
must answer: what is the input, what changes, what returns, and what can fail?

## Start With The Resource Contract

A task:

```go
type Task struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}
```

Decide ownership:

- client supplies title
- client may supply done during update
- server assigns ID
- server owns the collection

API table:

| Method | Path | Meaning | Success |
|---|---|---|---:|
| POST | `/tasks` | create | 201 |
| GET | `/tasks` | list all | 200 |
| GET | `/tasks/{id}` | read one | 200 |
| PUT | `/tasks/{id}` | replace editable fields | 200 |
| DELETE | `/tasks/{id}` | delete | 204 |

For each row, define malformed input, missing resource, and unsupported method
before coding.

## Collection Routes And Item Routes

These are different route categories:

```text
/tasks    -> collection
/tasks/7  -> one item
```

An explicit standard-library setup:

```go
mux.HandleFunc("/tasks", tasksHandler)
mux.HandleFunc("/tasks/", taskByIDHandler)
```

Then dispatch by method inside each handler:

```text
tasks handler:
    POST -> create
    GET  -> list
    other -> 405

task-by-ID handler:
    parse ID
    GET    -> find
    PUT    -> update
    DELETE -> delete
    other  -> 405
```

Do not make one giant handler containing every branch and every business rule.

## Parse IDs Defensively

For `/tasks/7`:

```text
remove the exact "/tasks/" prefix -> "7"
reject empty remainder
reject an extra slash
convert "7" to integer 7
reject conversion error
reject IDs less than 1
```

Useful functions:

```go
strings.TrimPrefix(path, "/tasks/")
strings.Contains(rawID, "/")
strconv.Atoi(rawID)
```

Invalid syntax should return 400. A valid ID that is not stored should return
404. Those are different failures.

## In-Memory State And Concurrency

An HTTP server handles requests concurrently. Two requests can access state at
the same time. Unprotected global slices can race.

Keep related state together:

```go
type taskStore struct {
    mu     sync.RWMutex
    tasks  []Task
    nextID int
}
```

Conceptual rules:

- list/find take a read lock
- create/update/delete take a write lock
- unlock on every path, normally with `defer`
- do not expose the internal slice for callers to mutate

Concurrency is not the main lesson, but ignoring it would teach an HTTP bug.
Later, a repository will own this state.

Run the race detector once the API works:

```bash
go run -race .
```

The race detector may require a supported C toolchain on Windows. If unavailable,
record that limitation rather than claiming the state is race-tested.

## CRUD Algorithms

Create:

```text
validate trimmed title
lock state
construct task with next ID and done false
increment next ID
append task
unlock
return created task
```

Find:

```text
read-lock state
scan tasks for matching ID
return task and true, or empty task and false
```

Update:

```text
validate input
lock state
find matching index
replace editable fields
return updated task
if absent, report not found
```

Delete:

```text
lock state
find matching index
remove exactly that element
report whether it existed
```

## Slice Deletion

Given index `i`:

```go
tasks = append(tasks[:i], tasks[i+1:]...)
```

Meaning:

```text
everything before i + everything after i
```

This may reuse the original backing array. Do not keep and expose stale internal
slice references.

## Worked Example: In-Memory Bookmark Store

This is storage logic for a similar resource:

```go
type Bookmark struct {
    ID    int
    URL   string
    Label string
}

type bookmarkStore struct {
    nextID    int
    bookmarks []Bookmark
}

func newBookmarkStore() *bookmarkStore {
    return &bookmarkStore{nextID: 1}
}

func (store *bookmarkStore) create(url, label string) Bookmark {
    bookmark := Bookmark{
        ID:    store.nextID,
        URL:   url,
        Label: label,
    }
    store.nextID++
    store.bookmarks = append(store.bookmarks, bookmark)
    return bookmark
}

func (store *bookmarkStore) find(id int) (Bookmark, bool) {
    for _, bookmark := range store.bookmarks {
        if bookmark.ID == id {
            return bookmark, true
        }
    }
    return Bookmark{}, false
}
```

The example omits locking to focus on the algorithms. Your HTTP-owned state must
add synchronization because handlers can overlap.

## Your Program

Implement:

- public `GET /health`
- all five task operations in the API table
- server-assigned monotonic IDs
- title validation
- JSON request and response bodies
- 400, 404, 405, and successful status codes
- synchronized in-memory state

Data disappearing after restart is expected. Persistence is exercise 15.

## Build In Checkpoints

1. Define the task contract and JSON helpers.
2. Implement collection `POST`.
3. Implement collection `GET`.
4. Parse item IDs and implement item `GET`.
5. Implement update.
6. Implement delete.
7. Add locks and run concurrent requests.
8. Test every method/path/error combination.

```bash
gofmt -w .\main.go
go vet .
go run .
```

Example request script from a second PowerShell:

```bash
$created = Invoke-RestMethod `
    -Uri http://localhost:8080/tasks `
    -Method Post `
    -ContentType "application/json" `
    -Body (@{ title = "learn routing" } | ConvertTo-Json)

$created
Invoke-RestMethod -Uri http://localhost:8080/tasks
Invoke-RestMethod -Uri "http://localhost:8080/tasks/$($created.id)"
```

## Test Matrix

| Method | Path/input | Expected |
|---|---|---|
| POST | valid title | 201 |
| POST | blank title | 400 |
| POST | malformed JSON | 400 |
| GET | `/tasks` empty | 200 and JSON array |
| GET | existing item | 200 |
| GET | valid missing ID | 404 |
| GET | nonnumeric ID | 400 |
| GET | `/tasks/1/extra` | 400 or 404 by documented contract |
| PUT | existing ID | 200 and updated task |
| DELETE | existing ID | 204 and empty body |
| DELETE | same ID again | 404 |

## Failure Experiments

1. Trust a client-provided ID and create two tasks with the same ID.
2. Parse a bad ID while ignoring `Atoi`'s error. Show how it becomes zero.
3. Return the internal slice, mutate it elsewhere, and explain the ownership
   leak.
4. Remove locking and send concurrent creates. Use the race detector if
   available.
5. Send JSON after a 204 response. Explain why 204 means no response body.

## You Understand This Exercise When

You can derive route behavior from the contract table, distinguish malformed
from missing IDs, trace every state mutation, and explain why concurrent HTTP
state needs synchronization.

References:

- [`http.ServeMux`](https://pkg.go.dev/net/http#ServeMux)
- [`strings`](https://pkg.go.dev/strings)
- [`strconv`](https://pkg.go.dev/strconv)
- [`sync.RWMutex`](https://pkg.go.dev/sync#RWMutex)
