# Project 12 - Trees

## Problem

Represent folders and calculate total files, maximum depth, and search path.

## What A Tree Is

A tree is hierarchical:

```text
root
|-- documents
|   `-- reports
`-- pictures
```

Each node has children. A tree has one root and no cycles.

## Depth-First Search

DFS fully explores a child before moving to the next child.

Recursive pseudocode:

```text
FUNCTION count files(node)
    total = node's direct file count
    FOR each child
        total = total + count files(child)
    RETURN total
```

Preorder visits node before children. Postorder visits children before node.
Directory size/deletion often needs postorder; display often uses preorder.

## Complexity

Visit every node once: `O(n)` time. Recursive stack is `O(h)`, where `h` is
height.

## Tasks

1. total all file counts.
2. calculate maximum depth.
3. find a folder by name.
4. return full path.
5. flatten folders in preorder.
6. implement iterative DFS with an explicit stack.

## Failure Drills

1. accidentally insert a cycle and recurse forever.
2. return after searching only the first child.
3. confuse number of children with total descendants.

## Done Means

You can draw the visit order and distinguish node count, depth, and path.

