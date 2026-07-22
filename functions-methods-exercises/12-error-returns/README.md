# Exercise 12 - Error Returns

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- functions should return errors instead of hiding failures
- the caller checks the error before using the result
- boundary code should stop on invalid input quickly

## Beginner Bridge

Start from a plain function, loop, or handler. Then add the new syntax only where it is needed.

### Before
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    count, _ := strconv.Atoi("12")
    fmt.Println(count)
}
```

### After
```go
package main

import (
    "fmt"
    "strconv"
)

func parseCount(raw string) (int, error) {
    count, err := strconv.Atoi(raw)
    if err != nil {
        return 0, fmt.Errorf("parse count: %w", err)
    }
    return count, nil
}

func main() {
    count, err := parseCount("12")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(count)
}
```

## Premade Helpers To Notice

- `errors.New` for a simple static error
- `fmt.Errorf` when the error should include context
- `strings.TrimSpace` or `strconv.Atoi` when input cleanup and parsing are part of the boundary

## Worked Example

Use the same pattern in a different domain first. The names are different. The structure is the part to copy.

### Example code
```go
package main

import (
    "fmt"
    "strconv"
)

func parseCount(raw string) (int, error) {
    n, err := strconv.Atoi(raw)
    if err != nil {
        return 0, fmt.Errorf("parse count: %w", err)
    }
    return n, nil
}

func main() {
    count, err := parseCount("12")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(count)
}
```

### Expected output
```text
12
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| parseCount | returns an error when input is bad |
| main | checks the error immediately |
| fmt.Errorf | adds context |

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
