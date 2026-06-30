# Project 06 - Pagination, Filtering, And Sorting

## Goal

Design collection queries that are bounded, deterministic, and indexable.

## Contract

```text
GET /projects/{id}/tasks?
    status=open&
    assigneeId=7&
    sort=-createdAt&
    pageSize=25&
    cursor=...
```

## Checkpoints

1. define allowed filters.
2. define fixed sort allowlist.
3. set default and maximum page size.
4. sort by `createdAt DESC, id DESC`.
5. encode both values in an opaque cursor.
6. reject malformed cursors.
7. return `nextCursor` only when another page may exist.
8. define behavior when rows are inserted between page requests.

## Why Tie-Breakers Matter

Several tasks can share one timestamp. Adding ID produces a total order:

```text
createdAt DESC, id DESC
```

## Failure Drills

1. paginate with `LIMIT` and no order.
2. sort only by non-unique timestamp.
3. allow arbitrary client SQL column names.
4. permit unbounded page size.
5. compare offset and cursor behavior during concurrent inserts.

## Done Means

Repeated traversal has deterministic order, bounded cost, and documented
concurrent-change behavior.

