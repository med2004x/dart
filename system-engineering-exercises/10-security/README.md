# Project 10 - Authentication, Authorization, And Abuse

## Goal

Prevent one authenticated user from reading or changing another user's task.
Client-provided IDs must never grant permission.

## Model

```go
type Task struct {
    ID      int
    OwnerID int
    Title   string
}
```

Authentication establishes trusted `actorID`. Authorization checks stored
`OwnerID`.

## Checkpoints

1. Build API-key authentication mapping keys to user IDs.
2. Store actor ID in request context.
3. implement `GetTask(actorID, taskID)`.
4. deny mismatched ownership.
5. apply method and body-size limits.
6. add per-identity request limiting.
7. ensure credentials never appear in logs.

## Required Test Matrix

| Actor | Task owner | Expected |
|---:|---:|---|
| 1 | 1 | success |
| 1 | 2 | denied |
| 2 | 1 | denied |
| none | any | 401 |
| invalid key | any | 401 |

Choose whether ownership mismatch returns 403 or privacy-preserving 404 and
document the information-leak tradeoff.

## Failure Drill

Accept `ownerId` from request JSON and authorize using it. Demonstrate how actor
1 submits owner 2. Then restore authorization from trusted stored data.

## Done Means

Changing path, query, body, or header resource IDs cannot grant access.

