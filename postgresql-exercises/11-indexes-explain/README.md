# Project 11 - Indexes And EXPLAIN

## Goal

Add one index for one measured query and prove its effect with query plans.

## Checkpoints

1. run `generate-data.sql`.
2. run `ANALYZE`.
3. capture `before.sql` output.
4. identify filters, sort, estimates, actual rows, and buffers.
5. design one matching index in `add-index.sql`.
6. run `ANALYZE` again.
7. capture the same query's after plan.
8. inspect index size.

## Main Query

```text
open tasks for one project
newest first
limit 20
```

The likely index begins with equality filters and then supports ordering.

## Failure Drills

1. add an index to a tiny table and expect it must be used.
2. put columns in an order that does not support the query.
3. create several redundant indexes and measure size.
4. use `EXPLAIN ANALYZE` on a mutation without rollback.

## Done Means

The index has a query, before/after plan, write/storage cost, and removal
criterion.

