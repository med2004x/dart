# Project 12 - Trees

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- tree problems track parent and child relationships
- the traversal order changes what you can compute easily
- recursive or stack-based walks both work

## Beginner Bridge

Start from a direct loop or trace. Then add the pattern only where it is needed.

### Before
```go
package main

import "fmt"

func main() {
    fmt.Println("start with the direct version")
}
```

### After
```go
package main

import "fmt"

func helper() string {
    return "walk parent and child links with the right traversal"
}

func main() {
    fmt.Println(helper())
}
```

## Worked Example

Use the same pattern in a different domain first. The names are different. The structure is the part to copy.

### Example code
```go
package main

import "fmt"

type Node struct {
    Value string
    Left  *Node
    Right *Node
}

func count(n *Node) int {
    if n == nil {
        return 0
    }
    return 1 + count(n.Left) + count(n.Right)
}

func main() {
    root := &Node{Value: "root", Left: &Node{Value: "left"}}
    fmt.Println(count(root))
}
```

### Expected output
```text
2
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| Node | stores parent and child links |
| count | walks both branches |
| nil | stops recursion |

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
