# Lesson 3: B+ Trees and Indexes

## Objective

Implement search and insert for a B+ tree — the data structure behind nearly every database index — and understand precisely why it's the right shape for disk-backed data, not just "a balanced tree."

## Prerequisites

Discrete math Lesson 4 (graphs/trees — general tree terminology), Lesson 2 (buffer pool — a real B+ tree's nodes are themselves pages managed by the buffer pool, though this lesson's implementation can work with in-memory nodes directly for simplicity, noting the connection).

## Learn

**Why not a binary search tree.** A binary search tree (2 children per node) gives O(log₂ n) lookup — fine in memory, where every node access costs roughly the same. On disk, every node access potentially costs a full page read (computer architecture Lesson 7's storage latency), and a binary tree over a large dataset has a *tall* tree — many levels, many potential disk reads per lookup. A B+ tree instead uses a much higher branching factor (hundreds of children per node, sized so each node fits exactly in one disk page), which produces a dramatically *shorter* tree for the same number of keys — a B+ tree over millions of rows commonly has only 3-4 levels, meaning a lookup costs only 3-4 page reads regardless of how large the dataset grows within reason, versus a binary tree's `log₂(millions)` ≈ 20+ levels/reads for the same data. This is a direct, deliberate application of "minimize the number of expensive disk-page accesses" as the design's primary optimization target, not raw comparison count.

**B+ tree structure specifically (as distinct from a general B-tree).** Internal nodes hold only keys and child pointers, used purely for navigation — no actual record data. All actual records (or, more commonly in practice, pointers to actual records/pages elsewhere) live in the **leaf nodes**, and leaf nodes are additionally linked together in a chain (each leaf points to the next) — this leaf-chain is specifically what makes range queries (`WHERE x BETWEEN a AND b`) efficient: once you've located the starting leaf via a normal tree search, you can scan forward through the linked leaves directly, without re-traversing the tree for each subsequent key.

**Node splitting on insert.** When inserting into a full node (leaf or internal) would exceed its capacity, the node splits into two, and the split propagates a new key up into the parent — if the parent is also full, it splits too, potentially propagating all the way up to the root, in which case the tree grows by one level (this is the specific mechanism that keeps a B+ tree *always* balanced — height only increases via root splits, uniformly across the whole tree, never via one branch growing deeper than another the way an unbalanced binary tree could).

**Why this matters directly for you.** Every `CREATE INDEX` in PostgreSQL (by default) builds exactly this structure. `EXPLAIN ANALYZE` showing "Index Scan" versus "Seq Scan" (computer architecture Lesson 7's transfer task previewed this) is choosing between using a B+ tree index (a handful of page reads via tree traversal, per this lesson) versus reading every page of the table sequentially — the B+ tree's shallow-height property is precisely why index scans on large tables are so much faster than sequential scans for selective queries.

## Attempt

1. Implement a B+ tree in Go with a small, deliberately low branching factor (e.g. max 4 keys per node — small enough to force splits quickly during testing, though a real implementation would use a much higher factor sized to a disk page). Implement `search(key)` (traverse from root to the correct leaf, returning whether the key exists and its associated value) and `insert(key, value)`.

2. Insert keys in increasing order (1, 2, 3, 4, 5, ...) and trace, at each insert, whether a split occurred and at which level. Confirm the tree's height only grows via root splits (verify by printing the tree structure — or at least its height — before and after each insert that triggers a split, and confirming height only increases when the root itself splits).

3. Implement the leaf-chain (each leaf node holding a pointer/reference to the next leaf in key order) and implement a `rangeSearch(low, high)` function that finds the starting leaf via a normal tree search, then scans forward via the leaf chain, collecting all keys in `[low, high]`. Test it against a tree with enough keys to span multiple leaves, and confirm it returns the correct, complete, correctly-ordered result.

4. Compare lookup cost directly: instrument your `search` function to count how many nodes it visits per lookup, and compare that count (for a tree holding, say, 1000 keys with your chosen branching factor) against `log₂(1000) ≈ 10` (what a binary tree would need). Report the actual difference and connect it explicitly to the branching-factor argument from Learn.

## Verify

For step 2, produce a small log (or diagram) showing the exact insert sequence and which inserts triggered splits, at which node, and whether the split propagated to the parent — enough detail that someone else could verify your split logic by hand-tracing the same sequence.

## Failure drill

Deliberately implement `rangeSearch` *without* the leaf-chain optimization — instead, for every key you want in the range, perform a fresh, independent tree search from the root. Measure and compare the total number of node visits for a range query spanning, say, 20 keys using this naive re-search approach versus your step 3 leaf-chain approach. Report the actual difference in node-visit count, and explain why the leaf-chain approach's advantage grows with the size of the range being queried — a range of 2 keys barely matters, but a range of thousands does, which is exactly the shape of query B+ tree indexes are specifically optimized to make cheap.

## Transfer

Run `EXPLAIN ANALYZE` on a real query against a table in TARDOC or Mahall's PostgreSQL database, once against an indexed column (`WHERE indexed_col = value`) and once against a non-indexed column with the same selectivity if possible (`WHERE non_indexed_col = value`). Compare the reported execution plans and actual timings, and explain the difference using this lesson's tree-height argument — the indexed query is doing a small, bounded number of B+ tree page reads (matching this lesson's step 4 measurement in shape, even though PostgreSQL's actual tree parameters differ from your toy implementation), while the non-indexed query is scanning every page of the table sequentially, with cost scaling with table size rather than staying roughly constant.

## Done when

Your B+ tree correctly implements search, insert-with-splitting, and leaf-chain range search, verified against a real multi-level tree (not just a single-node toy case), and you've directly measured and compared node-visit counts for both point lookups (versus a theoretical binary tree) and range queries (versus a naive re-search-per-key approach), connecting both measurements back to why B+ trees specifically are the standard index structure for disk-backed data.
