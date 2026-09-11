# Exercise 14 - Higher-Order Functions

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- functions can be passed around like data
- the helper can accept behavior as an argument
- one loop can run many different operations

## Beginner Bridge

Start from a plain function, loop, or handler. Then add the new syntax only where it is needed.

### Before
```go
package main

import "fmt"

func main() {
    values := []int{1, 2, 3}
    for _, v := range values {
        fmt.Println(v * 2)
    }
}
```

### After
```go
package main

import "fmt"

func mapValues(values []int, fn func(int) int) []int {
    out := make([]int, 0, len(values))
    for _, v := range values {
        out = append(out, fn(v))
    }
    return out
}

func double(n int) int { return n * 2 }

func main() {
    values := []int{1, 2, 3}
    doubled := mapValues(values, double)
    fmt.Println(doubled)
}
```

## Premade Helpers To Notice

- `sort.Slice` as a real example of passing comparison behavior
- `strings.Map` or `strings.FieldsFunc` for behavior passed as a function
- Function values are the syntax you need, not a special package helper.

## Worked Example

Use the same pattern in a different domain first. The names are different. The structure is the part to copy.

### Example code
```go
package main

import "fmt"

func mapValues(values []int, fn func(int) int) []int {
    out := make([]int, 0, len(values))
    for _, value := range values {
        out = append(out, fn(value))
    }
    return out
}

func double(n int) int { return n * 2 }

func main() {
    values := []int{1, 2, 3}
    doubled := mapValues(values, double)
    fmt.Println(doubled)
}
```

### Expected output
```text
[2 4 6]
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| mapValues | accepts behavior as input |
| double | is passed into the helper |
| main | chooses the behavior |

## Design / Reasoning Before Syntax

1. Identify the data that moves between helpers.
2. Decide whether the next step should return a value, return an error, or mutate a receiver.
3. Write the helper that owns the rule, not the whole workflow.
4. Wire the return value into the next call in `main`.
5. Use a trace like: input -> helper one -> helper two -> printed result.

This is the proof path. The code should match it instead of inventing a new shape after the fact.

## Your Program / Tasks

1. Work the concept without hiding the flow.
2. Keep the wiring visible. Do not hide a value behind unrelated temporary state.
3. Use the local quick reference for syntax, but keep the reasoning in this README.

## Build In Checkpoints

1. Write the smallest direct helper first.
2. Add the next helper or receiver only after the data flow is visible.
3. Print or return the value at the end, not in the middle.
4. Check the behavior with one normal input and one boundary input.

## Failure Drills

1. Put the calculation inside the wrong helper. Why: the caller no longer sees the flow.
2. Ignore a returned value and pretend the program still changed. Why: the value never moved.
3. Choose a receiver form without checking whether the original state must change. Why: copy and mutation are not the same thing.

## You Understand This When / Done Means

- Can you explain why the value is stored, returned, or mutated in that exact place?
- Can you trace the call chain without jumping over a helper?
- Can you say which syntax feature owns the state change?
