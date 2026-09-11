# Lesson 4: Graph Theory

## Objective

Formalize graphs (the mathematical structure, not the algorithms — those belong to the algorithms track's Lesson 13), and understand the definitions and properties that make graph algorithms provably correct.

## Prerequisites

Lesson 2 (relations — a graph is essentially a relation visualized).

## Learn

**Definitions.** A graph `G = (V, E)` consists of a vertex set `V` and an edge set `E` (pairs of vertices). *Undirected*: edges have no direction, `(u,v) = (v,u)`. *Directed* (digraph): edges are ordered pairs, `(u,v) ≠ (v,u)`. A **weighted** graph attaches a number to each edge (a cost, distance, or capacity).

**Degree.** In an undirected graph, `deg(v)` = number of edges touching `v`. Handshake lemma: `Σ deg(v) = 2|E|` (every edge contributes to exactly two vertices' degree counts), which means the sum of all degrees is always even — a useful sanity check when validating a graph's data.

**Paths, cycles, connectivity.** A *path* is a sequence of distinct vertices connected by edges. A *cycle* is a path that returns to its start. A graph is *connected* if there's a path between every pair of vertices. A **tree** is a connected graph with no cycles — exactly `|V| - 1` edges, and removing any edge disconnects it (this is a provable, not just observed, property).

**Directed Acyclic Graphs (DAGs).** A directed graph with no cycles. This is precisely the structure of a task dependency graph (Lesson 2's partial-order transfer task) — "must run before" relationships that don't loop back on themselves. **Topological sort** produces a linear ordering of vertices consistent with all edges (every edge points forward in the ordering) — this only exists if and only if the graph is acyclic, which is why cycle detection is the first check any dependency resolver (a package manager, a build system, a task scheduler) must perform before attempting to order anything.

**Why the "exists iff acyclic" fact matters practically:** if you've ever seen "circular dependency detected" as an error from a build tool, package manager, or ORM's migration system, that tool ran a cycle-detection check precisely because topological sort is undefined (impossible) on a graph with a cycle — there's no consistent "before" ordering when A depends on B which depends on A.

## Attempt

1. For the undirected graph with vertices `{A,B,C,D}` and edges `{(A,B),(B,C),(C,D),(D,A),(A,C)}`, compute the degree of every vertex, and verify the handshake lemma (`Σdeg(v) = 2|E|`) holds for your computed degrees.

2. Determine whether the graph in step 1 is connected (justify with an actual path between every pair, or identify a disconnected pair if none exists), and determine whether it contains a cycle (find one explicitly if so).

3. Model a small software dependency scenario as a DAG: 6 packages, where package X depends on package Y means a directed edge from X to Y (or Y to X — pick a direction convention and state it). Include at least one genuine multi-step dependency chain (A depends on B depends on C). Confirm by inspection that your graph has no cycles.

4. By hand, perform a topological sort on your DAG from step 3 (repeatedly: find a vertex with no incoming — or outgoing, depending on your convention — edges among the remaining graph, output it, remove it, repeat). Record the resulting order and confirm every dependency edge points in a direction consistent with your output order.

## Verify

For step 4, explicitly check every single edge in your original graph against your final topological order and confirm each one is satisfied (the "must come before" vertex actually appears before the "must come after" vertex in your output) — do this as an explicit line-by-line check, not just a visual scan.

## Failure drill

Add one edge to your DAG from step 3 that creates a cycle (e.g., make the "last" package in your dependency chain depend back on the "first" one). Attempt topological sort again by hand using the same removal process, and observe that at some point, no remaining vertex has zero incoming (or outgoing) edges — every remaining vertex has at least one dependency among the remaining set. Explain why this stuck state is the direct, hands-on evidence that topological sort is undefined on a cyclic graph, not just a rule you're told to accept.

## Transfer

If Lead Sourcer's pipeline stages, TARDOC's Celery task chains, or any deployment/build process you run has an implicit ordering ("this must happen before that"), draw it out as an actual directed graph (even informally) and check by inspection whether it's acyclic. State whether the tool you use (Celery, a Makefile, CI config) does this cycle-checking for you automatically or whether you're relying on manually getting the order right.

## Done when

You can compute vertex degrees and verify the handshake lemma on a graph you're given, you can construct a DAG and produce a valid topological sort by hand, and you can explain concretely (using your own failure-drill example) why a cycle makes topological sort impossible rather than just harder.
