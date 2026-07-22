# Project 09 - Recursion

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- recursion needs a base case
- each call works on a smaller subproblem
- you need to trace call depth and return order

## Beginner Bridge

Start from a direct loop or trace. Then add the pattern only where it is needed.

### Before
```go
package main

import "fmt"

func main() {
    fmt.Println("count the nested items by hand")
}
```

### After
```go
package main

import "fmt"

func countLeaves(nodes []string) int {
    if len(nodes) == 0 {
        return 0
    }
    if len(nodes) == 1 {
        return 1
    }
    return countLeaves(nodes[:len(nodes)-1]) + 1
}

func main() {
    fmt.Println(countLeaves([]string{"a", "b", "c"}))
}
```

## Worked Example

Use the same pattern in a different domain first. The names are different. The structure is the part to copy.

### Example code
```go
package main

import "fmt"

func countLeaves(nodes []string) int {
    if len(nodes) == 0 {
        return 0
    }
    if len(nodes) == 1 {
        return 1
    }
    return countLeaves(nodes[:len(nodes)-1]) + 1
}

func main() {
    fmt.Println(countLeaves([]string{"a", "b", "c"}))
}
```

### Expected output
```text
3
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| base case | stops the call chain |
| smaller slice | shrinks the problem |
| return | combines the recursive result |

## Design / Reasoning Before Syntax

1. Write the input shape and the smallest useful state.
2. Choose the one variable that proves progress.
3. Trace every step with a table or a short list.
4. Keep the stop rule visible before you optimize.
5. Compare the final answer against the trace.

This is the proof path. The code should match it instead of inventing a new shape after the fact.

## Your Program / Tasks

1. Solve the pattern with a direct trace first.
2. Write the direct version first, even if it looks boring.
3. Prove the output with a trace table or a small hand walk.

## Build In Checkpoints

1. Define the input shape and the stop condition.
2. Write the smallest loop or recursive step.
3. Add the boundary case before you touch the edge cases.
4. Compare your result with a hand trace.

## Failure Drills

1. Start with a clever trick before the direct version works. Why: you will not know what you are proving.
2. Skip the trace table. Why: the state changes stay hidden.
3. Forget the not-found or empty-input case. Why: that is where the bug lives.

## You Understand This When / Done Means

- Can you describe the state that changes on every step?
- Can you explain why the algorithm stops?
- Can you predict the answer without running the code?
