# Project 04 - Transactions And Integrity

## Goal

Guarantee that a failed persistence operation leaves no visible in-memory
change.

## Contract

```go
type TaskWriter interface {
    Save([]Task) error
}
```

Repository create flow:

```text
lock repository
copy current tasks
append new task to candidate copy
save candidate copy
if save fails, return without changing current tasks
replace current tasks
unlock
```

## Checkpoints

1. Implement a successful writer.
2. Implement a writer returning `disk full`.
3. Implement repository create with candidate state.
4. Test successful commit.
5. Test failed save leaves count and next ID unchanged.
6. Test concurrent creates with `go test -race`.

## Failure Drill

Mutate repository memory before `Save`. The rollback test must fail. Restore the
commit-after-save order.

## PostgreSQL Extension

After PostgreSQL project 09, repeat the idea with an account transfer and verify
rollback preserves total balance.

## Done Means

The caller never sees a state change from an operation reported as failed.

