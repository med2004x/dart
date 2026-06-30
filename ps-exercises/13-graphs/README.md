# Project 13 - Graphs And Breadth-First Search

## Problem

Find the fewest friendship steps between two users in an unweighted network.

## What A Graph Is

A graph contains:

- vertices/nodes: users
- edges: friendships

Unlike a tree, graphs can have cycles and several paths.

Adjacency list:

```text
A -> B, C
B -> A, D
C -> A, D
D -> B, C
```

## Breadth-First Search

BFS explores by distance:

```text
distance 0: start
distance 1: direct neighbors
distance 2: neighbors of neighbors
```

Use a queue and visited set.

Pseudocode:

```text
enqueue start
mark start visited

WHILE queue not empty
    current = dequeue
    IF current is target
        reconstruct path
    FOR each neighbor
        IF neighbor not visited
            mark visited immediately
            remember predecessor
            enqueue neighbor
return not found
```

Mark visited when enqueuing, not when dequeuing, to avoid repeated queue entries.

## Complexity

Each vertex and edge is processed at most a constant number of times:
`O(V + E)` time and `O(V)` extra space.

## Tasks

1. return whether path exists.
2. return shortest path.
3. handle start equals target.
4. handle disconnected graph.
5. test cycles and duplicate edges.
6. reject/ignore unknown nodes by documented policy.

## Done Means

You can draw queue contents, visited set, and predecessor map per step.

