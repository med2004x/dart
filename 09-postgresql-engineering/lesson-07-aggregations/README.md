# Project 07 - Aggregations And Reports

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Reduce rows into trustworthy summaries without double-counting joins.

## Required Reports

1. task count by status
2. open task count by owner
3. project count per user, including zero
4. projects with at least two open tasks
5. completion percentage per project
6. overdue task count per owner

## Mental Model

```text
FROM/JOIN creates input rows
WHERE filters input rows
GROUP BY forms groups
aggregate reduces each group
HAVING filters completed groups
ORDER BY sorts output
```

Use `COUNT(child.id)`, not `COUNT(*)`, for empty parents from a left join.

## Checkpoints

1. predict source row count.
2. identify group key.
3. name every aggregate.
4. handle empty groups.
5. prevent division by zero with `NULLIF`.
6. test many-to-many multiplication.

## Failure Drills

1. use `WHERE COUNT(*) > 1`.
2. select an ungrouped column.
3. count `*` on an empty left-joined parent.
4. join tags before counting tasks and inspect inflated totals.

## Done Means

Every reported number can be traced to its input rows and grouping rule.

