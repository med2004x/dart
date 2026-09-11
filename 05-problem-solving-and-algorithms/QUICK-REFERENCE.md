# Problem Solving Quick Reference

Use the direct algorithm first. The point is to understand state and proof
before choosing a named technique.

## Trace A Loop

```go
total := 0
for _, value := range values {
	total += value
}
```

Write a row for each iteration: input, state before, operation, state after.
The trace is the proof that the loop matches the pseudocode.

## Maps For Counts And Lookup

```go
counts := make(map[string]int)
for _, word := range words {
	counts[word]++
}
```

Missing integer map keys read as zero, which makes counting convenient. Still
decide whether a missing key and a stored zero mean different things.

## Sorting

```go
sort.Slice(items, func(i, j int) bool {
	return items[i].Score < items[j].Score
})
```

State the ordering rule before sorting. Decide how ties behave and whether the
input is allowed to be modified.

## Stack, Queue, And Search

```go
stack = append(stack, value)        // push
last := stack[len(stack)-1]         // peek
stack = stack[:len(stack)-1]        // pop

queue = append(queue, value)        // enqueue
next := queue[0]                    // front
queue = queue[1:]                   // dequeue
```

Check empty slices before indexing. For binary search, the input must satisfy
the ordering assumption. For BFS, track visited nodes so cycles terminate.

## Recursion And Dynamic Programming

Every recursive function needs:

1. a base case that stops recursion
2. a smaller subproblem
3. progress toward the base case

Dynamic programming is useful when subproblems overlap. Name the state clearly
before choosing a table or memoization map.

## Complexity

Identify the operation that repeats. One pass is usually `O(n)`. A nested scan
is often `O(n^2)`. A map lookup is average `O(1)`. Sorting is commonly
`O(n log n)`. State assumptions; Big-O is a growth description, not a runtime
promise.

## Best Practices

- Write input, output, constraints, and edge cases before code.
- Use a brute-force solution as a correctness reference when optimizing.
- Keep the algorithm separate from printing and input parsing.
- Test empty input, one item, duplicates, already ordered data, and boundaries.
- Do not call an algorithm understood until you can trace it by hand.

## Official Resources

- [Tour of Go](https://go.dev/tour/)
- [`sort`](https://pkg.go.dev/sort)
- [`container/heap`](https://pkg.go.dev/container/heap)
- [`container/list`](https://pkg.go.dev/container/list)
- [Go specification](https://go.dev/ref/spec)
