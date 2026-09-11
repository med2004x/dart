# Exercise 02 - Parameters

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- values enter a function through parameters
- parameter names are local to that function
- input cleanup should happen before the real work starts

## Beginner Bridge

Start from a plain function, loop, or handler. Then add the new syntax only where it is needed.

### Before
```go
package main

import (
    "fmt"
    "strings"
)

func labelUser() string {
    name := "  Mina  "
    return strings.TrimSpace(name)
}

func main() {
    fmt.Println(labelUser())
}
```

### After
```go
package main

import (
    "fmt"
    "strings"
)

func labelUser(raw string) string {
    return strings.TrimSpace(raw)
}

func main() {
    raw := "  Mina  "
    clean := labelUser(raw)
    fmt.Println(clean)
}
```

## Premade Helpers To Notice

- `strings.TrimSpace` to sanitize raw text before passing it onward
- `strconv.Atoi` when a string must become a number
- `fmt.Errorf` when bad input should become a useful error later

## Worked Example

Use the same pattern in a different domain first. The names are different. The structure is the part to copy.

### Example code
```go
package main

import (
    "fmt"
    "strings"
)

func labelUser(raw string) string {
    return strings.TrimSpace(raw)
}

func main() {
    raw := "  cafe order  "
    clean := labelUser(raw)
    fmt.Println(clean)
}
```

### Expected output
```text
cafe order
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| raw string | enters through the parameter |
| labelUser | owns the cleanup |
| main | passes the value in |

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
