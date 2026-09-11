# Lesson 6: Partitioning and Sharding

## Objective

Implement consistent hashing and understand precisely why it minimizes key movement when nodes join or leave — the standard technique behind how large-scale systems distribute data across many machines without requiring a full reshuffle on every membership change.

## Prerequisites

Discrete math Lesson 3 (combinatorics — reasoning about key redistribution probabilities uses similar counting ideas), probability Lesson 1 (foundations — useful for reasoning about the "roughly even distribution" property of hashing).

## Learn

**Why naive hashing fails at scale.** The simplest way to distribute N keys across M nodes: `node = hash(key) % M`. This works fine until M changes (a node is added or removed, which is routine in any system that scales up/down or tolerates failures) — because `% M` changing to `% (M±1)` remaps *almost every key* to a different node, not just the keys that "should" move. For a large dataset, this means a huge, unnecessary data migration triggered by a single node joining or leaving — a real, serious operational problem for any system using this naive scheme at scale.

**Consistent hashing: the fix.** Instead of hashing keys directly to node *indices*, both keys and nodes are hashed onto the same circular space (typically visualized as a ring, e.g. hash values from 0 to 2³²-1 wrapping back to 0). A key belongs to the first node encountered walking clockwise around the ring from the key's own hash position. When a node is added, it only takes over a portion of the ring previously owned by its (now) clockwise neighbor — only the keys in that specific arc need to move, not the entire dataset. When a node is removed, only the keys it owned need to move, to its clockwise neighbor — again, a small, bounded fraction of the total data, not everything.

**Virtual nodes: fixing consistent hashing's own uneven-distribution problem.** With only a few real nodes placed randomly on the ring, the arcs each node owns can be quite uneven in size purely by chance (a small number of random points on a circle don't necessarily divide it evenly) — meaning some nodes could end up owning a disproportionate share of keys. The standard fix is **virtual nodes**: each physical node is hashed onto the ring multiple times (e.g. 100-200 virtual positions per physical node), which — by a law-of-large-numbers-style averaging effect (probability track, Lesson 4) — produces a much more even overall distribution across physical nodes, since each physical node's total ring coverage becomes an average over many smaller, more evenly-distributed arcs rather than one single large, luck-dependent arc.

**Why this matters directly.** Any system that shards data across multiple database instances, cache servers, or storage nodes (a common pattern once a single-node system outgrows what one machine can hold or serve) faces exactly this problem. Memcached client libraries, DynamoDB's internal partitioning, and Cassandra's ring-based architecture all use consistent hashing (or a close variant) specifically because of the key-movement-minimization property this lesson demonstrates directly.

## Attempt

1. Implement naive modulo-based sharding: `nodeIndex = hash(key) % numNodes`. Populate it with, say, 10,000 test keys across 4 nodes, then simulate adding a 5th node (recompute `% 5` for every key) and measure exactly what fraction of the 10,000 keys ended up mapped to a *different* node than before — report the actual percentage, and compare it to what you'd naively expect (should be a very large fraction, close to `(numNodes-1)/numNodes` of keys moving, not just a proportional 1/5).

2. Implement consistent hashing: hash both nodes and keys onto a numeric ring (e.g. using a 32-bit or 64-bit hash function), and implement the "find the first node clockwise from the key's position" lookup (a sorted list of node positions plus a binary search, discrete math/algorithms track, is a natural implementation approach). Populate it with the same 10,000 test keys across the same 4 initial nodes.

3. Simulate adding a 5th node to your consistent-hashing ring (insert its hash position(s) into your sorted node-position list) and measure the actual fraction of the 10,000 keys that moved to a different node — report the number, and confirm it's dramatically smaller than the naive modulo approach's result from step 1, and reasonably close to the theoretical expectation of roughly `1/numNodes` (since only the new node's owned arc's worth of keys should move).

4. Implement virtual nodes: instead of hashing each physical node once onto the ring, hash it multiple times (e.g. 150 virtual positions per physical node, each derived by hashing something like `nodeID + virtualIndex`). Rerun your key-distribution measurement across the 4 (then 5) physical nodes and report the actual key-count-per-physical-node balance with and without virtual nodes — confirm virtual nodes measurably improve the evenness of distribution (report something like the standard deviation or min/max spread of keys-per-node in both configurations).

## Verify

Produce a table comparing: naive modulo (step 1), plain consistent hashing (steps 2-3), and consistent hashing with virtual nodes (step 4) — for each, report the percentage of keys that moved when adding the 5th node, and the evenness of the resulting distribution across nodes. The results should clearly show consistent hashing dramatically reducing key movement compared to naive modulo, and virtual nodes measurably improving distribution evenness compared to plain consistent hashing.

## Failure drill

Implement consistent hashing with a deliberately *small* number of virtual nodes per physical node (e.g. just 2, instead of 150) and rerun the distribution-evenness measurement from step 4. Confirm the distribution is noticeably more uneven than with 150 virtual nodes (though still likely better than a single hash position with zero virtual nodes) — report the actual numbers. Explain, using the law-of-large-numbers-style reasoning from Learn, why increasing the virtual-node count specifically improves evenness — you're averaging over more independent random ring positions per physical node, which (echoing probability track Lesson 4's Central Limit Theorem discussion, in spirit if not exact mechanism) reduces the variance of each physical node's total ring coverage relative to its mean.

## Transfer

If TARDOC or Mahall ever needed to shard their PostgreSQL data across multiple database instances (not a current stated need per your project history, but a reasonable hypothetical scaling scenario), describe, using this lesson's key-movement measurements, why consistent hashing would be strongly preferable to naive modulo-based sharding specifically for minimizing the operational disruption (data migration volume, and the corresponding downtime or complexity) of adding a new database shard as the dataset grows, compared to what a naive scheme would require every single time capacity needed to increase.

## Done when

You've measured and directly compared naive modulo sharding's catastrophic key movement against consistent hashing's dramatically smaller movement when adding a node, using your own generated data rather than trusting the claim abstractly, and you've demonstrated that virtual nodes measurably improve distribution evenness, with numbers showing the effect scales with virtual-node count as the failure drill's comparison establishes.
