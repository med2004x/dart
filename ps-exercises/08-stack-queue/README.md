# Project 08 - Stack And Queue

## Goal

Learn two data structures defined by removal order.

## Stack: Last In, First Out

Like plates:

```text
push A
push B
pop -> B
pop -> A
```

Use a slice:

```text
push: append
peek: last element
pop: read last, shorten slice
```

Problem: validate brackets in `"([{}])"`.

Algorithm:

```text
FOR each character
    IF opening bracket
        push it
    IF closing bracket
        IF stack empty
            invalid
        pop opening bracket
        IF pair does not match
            invalid
valid only if stack is empty
```

## Queue: First In, First Out

Like a service line:

```text
enqueue A
enqueue B
dequeue -> A
dequeue -> B
```

Problem: process jobs in arrival order.

## Complexity

Slice append/pop-at-end is normally `O(1)` amortized. Removing index zero by
shifting can be `O(n)`. A queue can keep a head index or use a ring buffer.

## Tasks

1. implement bracket validator.
2. implement queue with head index.
3. test empty, mismatched, nested, early close, leftover open.
4. test queue ordering and empty dequeue.
5. compact consumed queue storage safely.

## Done Means

You choose stack or queue by required removal order, not by name recognition.

