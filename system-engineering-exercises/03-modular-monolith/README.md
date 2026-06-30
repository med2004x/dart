# Project 03 - Modular Monolith

## Goal

Build a task use case whose business rules run without HTTP, files, or a
database.

## Structure

```text
TaskService -> TaskRepository interface <- MemoryTaskRepository
```

The starter defines the domain type and repository interface. Complete the
repository and service.

## Rules

- title is trimmed
- blank title is rejected
- repository assigns IDs
- missing IDs return `ErrTaskNotFound`
- repository returns copies, not mutable internal slices

## Checkpoints

1. Implement memory create.
2. Implement find by ID.
3. Implement service title validation.
4. Add service tests with a fake repository.
5. Add repository tests.
6. Run the race detector.

```powershell
gofmt -w .
go test -v ./...
go test -race ./...
go vet ./...
```

## Failure Drill

Move title validation into the memory repository. Use a fake repository in a
service test and prove the rule disappears. Restore ownership to the service.

## Done Means

All rules are testable without transport or infrastructure.

