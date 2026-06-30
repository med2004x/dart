# Project 03 - Insert And Select

## Goal

Create rows, retrieve generated IDs, filter sets, and define deterministic
ordering.

## Checkpoints

1. Complete and run `seed.sql`.
2. use `RETURNING` for every generated ID.
3. complete all prompts in `queries.sql`.
4. predict each result before execution.
5. verify result order explicitly.

## Required Queries

- all open tasks
- project 1 tasks ordered by priority then ID
- tasks due in the next three days
- tasks with no due date
- two newest tasks
- case-insensitive title search

## Rules

Prefer explicit columns:

```sql
SELECT id, title, status
FROM app.tasks;
```

Without `ORDER BY`, row order is not guaranteed.

Never retrieve a generated ID with `SELECT max(id)`. Concurrent inserts make
that unsafe.

## Failure Drills

1. Compare `due_at = NULL` with `due_at IS NULL`.
2. use `LIMIT` without `ORDER BY`.
3. guess IDs instead of using `RETURNING`.
4. select a column name that becomes ambiguous after a join.

## Done Means

You can state exact columns, rows, and ordering before running each query.

