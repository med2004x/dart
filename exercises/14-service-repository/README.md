# Exercise 14 - Handler, Service, And Repository

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

The same task API from exercise 12 now has enough responsibilities to separate:

```text
HTTP transport -> business rules -> storage
```

Layering is useful only when each boundary has a concrete reason. The reason
here is testability and change isolation:

- HTTP details can change without rewriting rules.
- Storage can change without rewriting rules.
- Rules can be tested without starting a server or database.

## Beginner Bridge: Layers Separate Reasons To Change

A handler changes when HTTP changes. A service changes when business rules
change. A repository changes when storage changes.

If one function parses JSON, validates titles, locks a slice, and writes a
response, it has too many reasons to change. That is how small programs become
hard to edit.

The clean flow is:

```text
handler:    translate HTTP into ordinary Go input
service:    decide whether the operation is allowed
repository: store and retrieve domain values
```

Similar layered programs:

| Domain | Handler owns | Service owns | Repository owns |
|---|---|---|---|
| Inventory | JSON/status | stock rules | item storage |
| Contacts | JSON/status | email/name rules | contact storage |
| Notes | JSON/status | note rules | note storage |

Do not add layers because they look professional. Add them because each layer
protects a different decision from leaking everywhere else.

## Dependency Direction

The intended dependency graph:

```text
handler -> service -> TaskRepository interface <- MemoryTaskRepository
   |
   +-> net/http and JSON
```

The service defines or consumes the repository behavior it needs. A concrete
memory repository implements that behavior.

Wrong direction:

```text
service imports HTTP status codes
service calls ORM/database APIs directly
repository decides whether titles are valid
handler edits repository fields
```

Those designs mix reasons to change.

## Responsibility Table

| Layer | Owns | Must not own |
|---|---|---|
| Handler | methods, paths, JSON, headers, statuses | title rules, slice mutation |
| Service | validation and use-case decisions | `http.Request`, JSON, DB queries |
| Repository | create/find/list/update/delete storage | HTTP statuses, title policy |
| Domain model | task data and intrinsic behavior | response writers, file paths |

An empty title is a business rule. Malformed JSON is a transport failure. A
missing stored ID is a repository fact that the service exposes as a domain
error and the handler maps to 404.

## Define Contracts Before Implementations

Domain model:

```go
type Task struct {
    ID    int
    Title string
    Done  bool
}
```

Repository behavior required by the service:

```go
type TaskRepository interface {
    Create(task Task) (Task, error)
    List() ([]Task, error)
    FindByID(id int) (Task, error)
    Update(task Task) (Task, error)
    Delete(id int) error
}
```

Do not expose ORM query builders or internal slices. Return domain values and
domain-relevant errors.

The exact signatures are a design decision. Write down who assigns IDs and
which layer owns concurrency before implementing.

## Constructor Injection

Make dependencies explicit:

```go
type TaskService struct {
    repository TaskRepository
}

func NewTaskService(repository TaskRepository) *TaskService {
    return &TaskService{repository: repository}
}
```

`main` assembles concrete objects:

```text
repository = new memory repository
service = new task service(repository)
handler = new task handler(service)
server = new HTTP server(handler routes)
```

This is dependency injection without a framework.

## Error Ownership And Mapping

Define stable errors at the domain/repository boundary:

```go
var ErrTaskNotFound = errors.New("task not found")
var ErrTitleRequired = errors.New("title is required")
```

Mapping belongs in the handler:

| Service result | HTTP response |
|---|---|
| success | operation-specific 2xx |
| `ErrTitleRequired` | 400 |
| `ErrTaskNotFound` | 404 |
| unknown internal error | 500 |

Use `errors.Is` so wrapped errors remain classifiable.

Do not compare exact error message strings and do not send unknown internal
errors to clients.

## Trace One Request

`PUT /tasks/7`:

```text
handler:
    parse ID 7
    decode update JSON
    call service.Update(7, input)

service:
    trim and validate title
    ask repository to find/update task
    return task or domain error

repository:
    lock storage
    locate ID 7
    mutate stored task or return not-found error

handler:
    map error to status, or encode updated task
```

At every arrow, only ordinary Go values and errors cross inward. HTTP objects do
not enter the service.

## Worked Example: Inventory Layers

Small analogous contracts:

```go
type Item struct {
    ID    int
    Name  string
    Stock int
}

type ItemRepository interface {
    Save(item Item) (Item, error)
    FindByID(id int) (Item, error)
}

type InventoryService struct {
    items ItemRepository
}

func (service InventoryService) AddStock(id, amount int) (Item, error) {
    if amount <= 0 {
        return Item{}, errors.New("stock amount must be positive")
    }

    item, err := service.items.FindByID(id)
    if err != nil {
        return Item{}, fmt.Errorf("find item: %w", err)
    }

    item.Stock += amount
    return service.items.Save(item)
}
```

The service owns the positive-amount rule. The repository owns retrieval and
saving. An HTTP handler would own parsing `id` and `amount`.

## Required Files

Keep one package for now:

```text
main.go        composition root and server startup
task.go        domain type and stable domain errors
repository.go  interface and memory implementation
service.go     use cases and business rules
handler.go     HTTP parsing and responses
```

Multiple files do not automatically create architecture. The import and call
directions do.

## Build In Checkpoints

1. Define the task model, errors, and repository interface.
2. Implement and directly exercise the memory repository.
3. Write service tests with a small fake repository.
4. Implement service rules until tests pass.
5. Add handlers and handler tests with `httptest`.
6. Assemble dependencies in `main`.
7. Run the full API contract from exercise 12.

```powershell
gofmt -w .
go test ./...
go vet ./...
go run .
```

## Test Boundaries

Service unit tests:

| Case | Infrastructure needed |
|---|---|
| blank title rejected | fake repository only |
| repository not found propagated | fake repository only |
| valid create calls repository | fake repository only |

Repository tests:

| Case | What it proves |
|---|---|
| create assigns/stores ID | storage behavior |
| update missing ID | stable not-found error |
| concurrent access | synchronization |

Handler tests:

| Case | What it proves |
|---|---|
| malformed JSON -> 400 | transport mapping |
| service not found -> 404 | error mapping |
| success -> JSON/status | API contract |

## Failure Experiments

1. Pass `*http.Request` into the service. List every way this couples business
   logic to HTTP.
2. Put title validation in the repository. Call the service with a fake
   repository and show that the rule disappears.
3. Return the repository's internal slice. Mutate it from a caller.
4. Map every service error to 400. Explain why storage failures are not client
   mistakes.
5. Create an interface with methods no service uses. Remove them.

## You Understand This Exercise When

You can assign every line to one owner, draw dependency direction, test rules
without HTTP or storage, and replace the memory repository without changing the
service.

References:

- [`net/http`](https://pkg.go.dev/net/http)
- [`net/http/httptest`](https://pkg.go.dev/net/http/httptest)
- [Go module layout](https://go.dev/doc/modules/layout)
- [`errors.Is`](https://pkg.go.dev/errors#Is)
