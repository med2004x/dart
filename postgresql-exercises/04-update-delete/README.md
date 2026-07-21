# Project 04 - Update And Delete Safely

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Mutate exactly the intended row set and verify what changed.

## Method

Before every mutation:

```text
write WHERE
run SELECT with the same WHERE
confirm target rows
run mutation in a transaction
use RETURNING
commit or roll back deliberately
```

## Checkpoints

Complete `operations.sql`:

1. rename one task.
2. mark one open task done only if it is open.
3. postpone open tasks in one project by one day.
4. delete one known test task.
5. rehearse a bulk update and roll it back.
6. verify zero-row results for missing IDs.

## Failure Drills

1. Omit `WHERE` inside a transaction and inspect before rollback.
2. update a missing ID and ignore affected row count.
3. write a delete without `RETURNING` and explain weaker evidence.

## Done Means

Every mutation identifies, previews, and reports its actual affected rows.

