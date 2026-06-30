# Project 09 - Recursion

## Problem

Count all comments in a nested discussion where each comment can have replies.

## What Recursion Is

A recursive function calls itself on a smaller part of the same problem.

Data:

```text
comment
`-- replies
    `-- replies
```

Algorithm:

```text
FUNCTION count(comment)
    total = 1 for current comment
    FOR each reply
        total = total + count(reply)
    RETURN total
```

The base case is a comment with no replies. Its loop runs zero times and returns
1. A base case prevents infinite calls.

## Call Trace

```text
A
|-- B
`-- C
    `-- D
```

```text
count(B) = 1
count(D) = 1
count(C) = 1 + count(D) = 2
count(A) = 1 + count(B) + count(C) = 4
```

## Complexity

Every comment is visited once: `O(n)` time. Call-stack space is proportional to
maximum nesting depth: `O(h)`.

Deep untrusted nesting can exhaust stack resources. An explicit stack can
replace recursion.

## Tasks

1. count all comments.
2. find maximum depth.
3. flatten comments in preorder.
4. test one node, wide tree, deep chain, empty forest.
5. write an iterative stack version.

## Done Means

You can identify the smaller subproblem, base case, and maximum call depth.

