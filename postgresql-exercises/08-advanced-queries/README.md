# Project 08 - NULL, CTEs, And Window Functions

## Goal

Choose advanced query tools based on result shape instead of writing one opaque
statement.

## Required Queries

1. tasks with no due date
2. users with at least one open task using `EXISTS`
3. projects with no tasks using `NOT EXISTS`
4. each task plus project task count
5. top two priority tasks per project
6. running task count ordered by creation

## Rules

`NULL` is missing/unknown:

```sql
due_at IS NULL
```

not:

```sql
due_at = NULL
```

A CTE names an intermediate result. A window function keeps individual rows
while calculating over related rows. `GROUP BY` collapses rows.

## Failure Drills

1. use `NOT IN` with a subquery that can return `NULL`.
2. use `GROUP BY` when individual tasks must remain.
3. apply `COALESCE` after an inner join removed the row.
4. rank without a deterministic tie-breaker.

## Done Means

You can explain the shape and purpose of every intermediate result.

