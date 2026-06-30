# Project 08 - Concurrent Updates And Preconditions

## Goal

Prevent a stale client from silently overwriting a newer task version.

## Scenario

```text
client A reads version 4
client B reads version 4
client A updates -> version 5
client B sends stale version 4
```

## Contract

Response:

```text
ETag: "task-7-v4"
```

Update request:

```text
If-Match: "task-7-v4"
```

The repository updates only when stored version matches.

## Checkpoints

1. add integer task version.
2. return ETag on reads.
3. require `If-Match` on updates.
4. parse only the documented ETag format.
5. update with `WHERE id = ? AND version = ?`.
6. increment version atomically.
7. return 412 for stale precondition.
8. test two concurrent writers.

## Failure Drills

1. read then write without a version condition.
2. trust a body version but ignore `If-Match` contract.
3. return 200 after zero affected rows.

## Done Means

A stale client receives a conflict signal and cannot overwrite unseen changes.

