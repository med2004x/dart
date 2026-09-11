# Exercise 09 - Value Receiver Methods

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- a value receiver works on a copy of the struct
- the original struct does not change unless you assign a new value
- value receivers are good for read-only or copy-based behavior

## Beginner Bridge

Start from a plain function, loop, or handler. Then add the new syntax only where it is needed.

### Before
```go
package main

import "fmt"

type Counter struct {
    Name  string
    Count int
}

func (c Counter) bumpBy(n int) {
    c.Count += n
}

func (c *Counter) bumpPointer(n int) {
    c.Count += n
}

func main() {
    item := Counter{Name: "desk bell", Count: 2}
    item.bumpBy(3)
    fmt.Println(item.Count)
}
```

### After
```go
package main

import "fmt"

type Counter struct {
    Name  string
    Count int
}

func (c Counter) previewBump(n int) int {
    return c.Count + n
}

func main() {
    item := Counter{Name: "desk bell", Count: 2}
    next := item.previewBump(3)
    fmt.Println(item.Count)
    fmt.Println(next)
}
```

## Premade Helpers To Notice

- `fmt.Printf` to print the whole struct before and after the call
- `%+v` in `fmt.Printf` when you want field names in the output
- A value receiver is the helper itself here, not a package helper.

## Worked Example

Use the same pattern in a different domain first. The names are different. The structure is the part to copy.

### Example code
```go
package main

import "fmt"

type Meter struct {
    Name  string
    Count int
}

func (m Meter) addValue(delta int) {
    m.Count += delta
}

func (m *Meter) addPointer(delta int) {
    m.Count += delta
}

func main() {
    meter := Meter{Name: "desk bell", Count: 2}

    meter.addValue(3)
    fmt.Printf("after value receiver: %+v\n", meter)

    meter.addPointer(3)
    fmt.Printf("after pointer receiver: %+v\n", meter)
}
```

### Expected output
```text
after value receiver: {Name:desk bell Count:2}
after pointer receiver: {Name:desk bell Count:5}
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| value receiver call | works on a copy |
| pointer receiver call | updates the same struct |
| printed struct after each call | proves what changed |

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
