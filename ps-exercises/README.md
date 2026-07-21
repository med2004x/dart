# Problem Solving And Algorithms Project Track

`PS` means problem solving.

You are not expected to know algorithms before starting. An algorithm is a
finite, precise sequence of steps that transforms input into output.

Example:

```text
input: [4, 7, 2]
goal: total
steps:
    start total at 0
    add each number
output: 13
```

Use the short [Quick Reference](QUICK-REFERENCE.md) when a Go or algorithm
pattern is unfamiliar. It gives generic syntax and proof habits without solving
the numbered projects.

The code is a translation of the steps.

## How To Solve A Problem

Use this order:

1. Restate the input and output.
2. Write examples by hand.
3. list invalid and edge inputs.
4. write direct pseudocode.
5. trace every variable.
6. translate to Go.
7. test normal, boundary, and failure cases.
8. measure complexity only after correctness.
9. choose a better algorithm only when needed.

## Complexity In Plain Language

Complexity describes how work or memory grows when input grows.

- `O(1)`: same amount of work regardless of input size
- `O(n)`: work grows roughly with number of items
- `O(n log n)`: common efficient sorting growth
- `O(n^2)`: compare many pairs; doubles can create about four times work

Big-O does not measure exact milliseconds. It describes growth.

## Project Map

| Project | Problem-solving tool |
|---|---|
| 01 Decompose And Trace | pseudocode and state tables |
| 02 Aggregation | one-pass totals/min/max |
| 03 Linear Search | find by scanning |
| 04 Frequency Maps | count/group by key |
| 05 Sorting | order records deliberately |
| 06 Two Pointers | coordinate two positions |
| 07 Sliding Window | reuse adjacent-range work |
| 08 Stack And Queue | LIFO/FIFO state |
| 09 Recursion | solve smaller copies safely |
| 10 Binary Search | halve sorted search space |
| 11 Intervals | sort and merge ranges |
| 12 Trees | hierarchical traversal |
| 13 Graphs | connected traversal/BFS |
| 14 Dynamic Programming | reuse overlapping results |
| 15 Capstone | dependency planner |

## Commands

```powershell
Set-Location C:\Users\pc\Documents\dart\ps-exercises
gofmt -w .
go test ./...
go vet ./...
```

## Rules

- Do not memorize code.
- Do not optimize before the direct solution works.
- Do not use a named algorithm you cannot trace.
- Do not claim complexity without identifying the repeated operation.
- Do not ignore empty input, duplicates, ties, or integer boundaries.

