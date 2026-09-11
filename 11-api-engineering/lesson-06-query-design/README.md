# Lesson 6: Query Design (Filtering, Sorting, Pagination)

## Objective

Design collection-endpoint query parameters (filtering, sorting, pagination) that scale correctly as data grows, and understand why naive "return everything" or offset-based pagination both break down at real scale.

## Prerequisites

Lesson 2 (resource modeling — this lesson covers how clients query the collection endpoints modeled there), database-internals Lesson 3 (B+ trees/indexes — efficient filtering and sorting at the API level ultimately depends on appropriate indexes existing underneath).

## Learn

**Why "just return the whole collection" fails as data grows.** A `GET /clinics` endpoint that returns every single row, unpaginated, works fine during development with a handful of test records and becomes a real problem once the table has thousands or millions of rows — a slow, memory-heavy response for both server and client, and (per database-internals Lesson 4's query-execution discussion) potentially a full sequential table scan on the server side if no filtering narrows what's actually needed. Pagination isn't an optional nicety for large APIs, it's a structural requirement for a collection endpoint to remain viable as the underlying data grows.

**Offset-based pagination, and its real, well-known problem.** `GET /clinics?offset=100&limit=20` — conceptually simple, but has two genuine issues: performance (many databases must actually scan and discard the first `offset` rows before returning the requested page, meaning cost grows with how deep into the collection you're paginating, not staying constant per page — connecting directly to database-internals Lesson 4's sequential-scan cost discussion), and consistency (if rows are inserted or deleted between page requests, offset-based pagination can skip or duplicate items across pages, since "the 100th row" isn't a stable identity, it shifts as the underlying data changes).

**Cursor-based (keyset) pagination: the standard fix at scale.** Instead of an offset, the client provides a cursor — typically an encoded reference to the last item seen on the previous page (e.g. its ID, or a compound key if sorting by a non-unique field) — and the server returns the next N items *after* that specific position, using a `WHERE id > lastSeenID ORDER BY id LIMIT N` style query, which a B+ tree index (database-internals Lesson 3) can satisfy efficiently regardless of how deep into the collection you are, and which remains stable even if rows are inserted/deleted elsewhere in the collection, since the cursor references a specific, stable position rather than a shifting numeric offset.

**Filtering and sorting: keeping the query parameter surface predictable.** A consistent convention (e.g. `?status=active&sort=-created_at` for filtering by status and sorting by creation date descending) that applies uniformly across every collection endpoint is what lets a client reasonably guess how to filter/sort a resource type they haven't used before, echoing Lesson 2's point about predictable API shape. Critically, every filterable/sortable field needs a corresponding database index (database-internals Lesson 3) to avoid a full table scan per filtered/sorted query — exposing a filter parameter for an unindexed column is a real, easy-to-miss way to introduce a slow, unscalable query path into an otherwise well-designed API.

## Attempt

1. Implement offset-based pagination for a collection endpoint (`?offset=N&limit=M`) against a table with a meaningful number of test rows (at least a few thousand, generated synthetically if needed, to make performance differences actually measurable rather than negligible). Measure query response time at `offset=0` versus `offset=5000` and report the actual difference.

2. Implement cursor-based pagination for the same endpoint (`?after=<cursor>&limit=M`, where the cursor encodes the last-seen item's sort key), using a proper indexed query (`WHERE sort_key > decoded_cursor ORDER BY sort_key LIMIT M`) rather than an offset. Measure response time for an equivalent "deep" page (functionally similar position to your step 1 high-offset test) and compare against the offset-based version's timing at a similar depth.

3. Demonstrate the offset-pagination consistency bug directly: fetch page 1 (offset=0) of a collection, then insert a new row that would sort to the very beginning of the collection (or delete a row from early in the collection), then fetch what would be "page 2" using the original offset-based scheme, and confirm you've either skipped an item or seen a duplicate across the two page fetches, depending on which mutation you made — directly reproducing the correctness problem, not just the performance one.

4. Implement filtering and sorting for the same collection endpoint using a consistent query-parameter convention of your choosing, ensure the filtered/sorted columns have appropriate database indexes (database-internals Lesson 3), and use `EXPLAIN ANALYZE` (database-internals Lesson 4) to confirm your filtered queries are using an index scan, not a sequential scan.

## Verify

For step 1-2, report actual measured response times (offset-based at low vs. high offset, and cursor-based at an equivalent depth) — cursor-based should show roughly constant time regardless of depth, while offset-based should show measurably increasing time at higher offsets. For step 3, show the actual page 1 and page 2 results demonstrating the skip or duplicate.

## Failure drill

Implement a filter parameter for a column that deliberately has *no* index, and run `EXPLAIN ANALYZE` on a query using that filter against your large test table. Confirm it shows a sequential scan (database-internals Lesson 4's terminology) rather than an index scan, and measure the actual query time compared to an equivalent, properly indexed filter from step 4. Report the measured difference, and explain why exposing a filter parameter without ensuring a backing index exists is a real, easy design mistake — the API contract *looks* identical to a properly indexed filter from the client's perspective (both are just `?field=value` query parameters), but one scales and one doesn't, and this difference is completely invisible until someone actually measures it against a realistically large dataset, exactly as this drill just did.

## Transfer

If TARDOC's or Mahall's collection endpoints (e.g. listing clinics, listing products) currently use offset-based pagination, describe, using your own step 1-3 measurements as evidence, at what approximate collection size you'd expect the offset-based performance degradation to become noticeable to real users, and whether migrating to cursor-based pagination is worth prioritizing now versus deferring until the collection actually grows large enough for the difference to matter in practice — a genuine tradeoff worth reasoning about explicitly rather than either prematurely optimizing or ignoring a known future bottleneck.

## Done when

You've measured and directly compared offset-based versus cursor-based pagination performance at depth, using real data rather than trusting the general claim, you've reproduced the offset-pagination consistency bug with an actual skip or duplicate across pages, and you've confirmed via `EXPLAIN ANALYZE` that your filtering/sorting implementation actually uses indexes rather than assuming it does — plus directly measured the cost of a filter that doesn't.
