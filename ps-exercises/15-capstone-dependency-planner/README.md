# Project 15 - Dependency Planner Capstone

## Problem

Given jobs and dependencies, return an order that runs every prerequisite before
its dependent job.

Example:

```text
compile depends on generate
test depends on compile
deploy depends on test
```

Valid order:

```text
generate, compile, test, deploy
```

If A depends on B and B depends on A, no valid order exists.

## Algorithm: Topological Sort

You need:

- graph from prerequisite to dependent jobs
- indegree count: number of unfinished prerequisites per job
- queue of jobs with indegree zero

Pseudocode:

```text
build graph and indegree counts
enqueue every job with indegree zero

WHILE queue not empty
    remove one ready job
    append it to result
    FOR each dependent job
        decrease dependent indegree
        IF indegree becomes zero
            enqueue dependent

IF result count != job count
    dependency cycle exists
RETURN result
```

## Trace

```text
generate indegree 0
compile indegree 1
test indegree 1
deploy indegree 1
```

Processing `generate` makes `compile` ready, and so on.

## Complexity

Building and processing touches every job and dependency:
`O(V + E)` time and `O(V + E)` storage.

## Capstone Requirements

1. parse jobs and dependency pairs.
2. reject unknown job references.
3. reject self-dependency.
4. remove or reject duplicate dependency edges.
5. return deterministic order when several jobs are ready.
6. detect cycles.
7. return one useful cycle explanation as an extension.
8. test disconnected job groups.

## Use Prior Tools

- maps for graph/indegree
- queue for ready jobs
- sorting for deterministic ready order
- graph reasoning for cycles
- trace tables for debugging

## Failure Drills

1. omit isolated jobs.
2. count duplicate edges twice.
3. process a dependent before indegree reaches zero.
4. return partial order without reporting a cycle.

## Done Means

You can trace graph, indegrees, queue, result, and cycle decision for every test.

