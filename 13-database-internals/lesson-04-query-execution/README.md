# Lesson 4: Query Execution

## Objective

Build a minimal query execution pipeline (scan → filter → project) using the iterator pattern real database engines use, and learn to read a real `EXPLAIN` plan as a description of exactly this kind of operator tree.

## Prerequisites

Lesson 3 (B+ trees — this lesson's "indexed execution" comparison directly uses that structure), algorithms track (general familiarity with composing simple operations into pipelines).

## Learn

**The iterator model (Volcano-style execution, the standard architecture).** A query plan is a tree of operators (scan, filter, project, join, sort, aggregate...), and each operator implements a uniform interface: typically `open()`, `next()` (returns the next row, or signals "no more rows"), and `close()`. Crucially, a parent operator's `next()` call pulls from its child operator's `next()` — data flows one row at a time, pulled from the bottom of the tree up, rather than each operator processing its entire input in one batch and materializing the full intermediate result before the next operator starts. This row-at-a-time pull model is precisely why a query with a `LIMIT 10` can often stop early without processing the entire dataset — the top-level operator simply stops calling `next()` on its child once it has enough rows, and that "stop asking for more" propagates down through the whole tree.

**Scan.** The base of most query trees — either a **sequential scan** (read every page of the table in order, per Lesson 1's page format, yielding every row) or an **index scan** (use a B+ tree, Lesson 3, to jump directly to relevant rows without reading the whole table) — the choice between these two is exactly what a query planner decides based on selectivity estimates (roughly: how many rows will actually match, which determines whether the index-traversal overhead is worth it compared to just scanning everything).

**Filter.** Wraps a child operator, calling its `next()` repeatedly and only passing through rows that satisfy a predicate (`WHERE` clause) — implemented as a simple loop: pull from child, test the predicate, discard or pass through, repeat until child is exhausted or a passing row is found.

**Project.** Wraps a child operator, transforming each passed-through row to include only the requested columns (or computed expressions) — this is what `SELECT col1, col2` (as opposed to `SELECT *`) corresponds to at the execution level.

**Why the iterator model composes so cleanly.** Because every operator has the identical `next()` interface regardless of what it does internally, operators can be nested arbitrarily deep without any operator needing to know anything about what kind of operator its child or parent is — a filter doesn't care whether its child is a sequential scan or an index scan or another filter; it just calls `next()` and gets a row. This is the same "uniform interface enables composition" idea behind Unix pipes (Linux tools Lesson 1) and Go's `io.Reader`/`io.Writer` interfaces, applied at the level of database rows instead of bytes.

## Attempt

1. Implement the iterator interface (`Open()`, `Next() (row, hasMore)`, `Close()`, or the Go-idiomatic equivalent) and a `SeqScan` operator that iterates over an in-memory slice of rows (standing in for reading table pages, per Lesson 1) one at a time.

2. Implement a `Filter` operator that wraps any child iterator and a predicate function, calling the child's `Next()` and only returning rows where the predicate is true — confirm it correctly stops (returns `hasMore = false`) once the child is exhausted, not just once it happens to find no more matching rows in a finite lookahead.

3. Implement a `Project` operator that wraps a child iterator and a list of column indices/names to keep, transforming each row before passing it up. Compose all three: `Project(Filter(SeqScan(data), predicate), columns)`, and confirm the full pipeline produces correct results for a nontrivial test dataset and predicate.

4. Implement an `IndexScan` operator using your Lesson 3 B+ tree instead of a full sequential scan, for queries with an equality or range predicate on the indexed key. Run the same logical query (same predicate, same expected result) through both a `Filter(SeqScan(...))` pipeline and an `IndexScan(...)` pipeline, confirm they produce identical results, and measure the difference in "rows/nodes touched" between the two approaches for a selective query (one matching only a small fraction of the total dataset) — this is your own hands-on version of what a query planner is deciding when it chooses between a sequential scan and an index scan.

## Verify

For step 4, report the exact count of rows/nodes visited by each approach for the same selective query, and confirm the index-scan approach visited dramatically fewer — directly connecting your own measured numbers to the general principle rather than just citing it.

## Failure drill

Run your step 4 comparison again, but this time with a *non-selective* predicate — one matching most or all of the rows in the dataset (e.g. `WHERE x > 0` against data where nearly every row satisfies it). Confirm the index-scan approach no longer shows a clear advantage over the sequential scan — it may even be slower, since traversing the B+ tree structure plus following pointers back to full row data can cost more overhead than simply reading rows sequentially when nearly all of them are going to be returned anyway. Explain, using this concrete measured result, why "always use an index" is not correct general advice — index scans win specifically when selectivity is high (few matching rows), and this is exactly the kind of decision a real query planner's cost-based optimizer is making by estimating selectivity before choosing a plan, a decision your own comparison just made manually and empirically instead.

## Transfer

Run `EXPLAIN ANALYZE` on two real queries against a TARDOC or Mahall table — one with a highly selective `WHERE` clause and one with a non-selective one — and identify, in PostgreSQL's actual output, whether it chose an Index Scan or a Seq Scan for each, and whether that choice matches the pattern your own step 4/failure-drill measurements would predict. State explicitly whether PostgreSQL's choice surprised you in either case, and if so, what additional factor (e.g. how "stale" the table's collected statistics are, which PostgreSQL uses to estimate selectivity before actually running the query) might explain a plan choice that doesn't match your naive selectivity intuition.

## Done when

Your composed scan/filter/project pipeline produces correct results using the iterator pattern, you've built and correctly used an index-scan alternative to a sequential scan, and you've measured — not just asserted — that the index scan wins decisively on a selective query and loses or ties on a non-selective one, connecting this directly to how real query planners make the same tradeoff decision.
