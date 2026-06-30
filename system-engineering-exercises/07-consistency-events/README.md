# Project 07 - Consistency And Events

## Goal

Build a source of truth plus a delayed search projection. Recover from delayed,
duplicate, and interrupted event delivery.

## Model

```text
task store + outbox -> worker -> search projection
```

The task store is authoritative. Search may be stale.

## Checkpoints

1. Create a task in the primary map.
2. Store an outbox event in the same locked operation.
3. Return before projection processing.
4. Run a worker that processes pending events after two seconds.
5. Mark events delivered only after projection update.
6. Make projection updates idempotent by task ID.
7. Restart the worker and process pending events.

## Required Tests

- primary read sees task immediately
- search does not promise immediate visibility
- projection eventually sees task
- duplicate event does not duplicate data
- worker interruption leaves event pending
- restarted worker catches up

```powershell
go test -v ./...
go test -race ./...
```

## Failure Drill

Publish only to an in-memory channel after creating the task. Stop the worker
between those operations. Explain why a committed task can lose its event.

## Done Means

The source of truth, stale-read contract, duplicate policy, and repair path are
explicit and tested.

