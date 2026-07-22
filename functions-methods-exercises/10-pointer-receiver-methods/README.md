# Exercise 10 - Pointer Receiver Methods

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- a pointer receiver can mutate the original struct
- the method call changes shared state on purpose
- pointer receivers are the right choice for real updates

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

func (c Counter) add(n int) {
    c.Count += n
}

func main() {
    item := Counter{Name: "whiteboard marker", Count: 2}
    item.add(3)
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

func (c *Counter) add(n int) {
    c.Count += n
}

func main() {
    item := Counter{Name: "whiteboard marker", Count: 2}
    item.add(3)
    fmt.Println(item.Count)
}
```

## Premade Helpers To Notice

- `fmt.Printf` to prove the original struct changed
- `%+v` to make the field updates visible
- A pointer receiver is the important syntax, not a package helper.

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

func (m Meter) previewAdd(delta int) int {
    return m.Count + delta
}

func (m *Meter) add(delta int) {
    m.Count += delta
}

func main() {
    meter := Meter{Name: "whiteboard marker", Count: 2}

    fmt.Printf("start: %+v\n", meter)
    preview := meter.previewAdd(3)
    fmt.Printf("after value receiver preview: %+v, preview=%d\n", meter, preview)

    meter.add(3)
    fmt.Printf("after pointer receiver add: %+v\n", meter)
}
```

### Expected output
```text
start: {Name:whiteboard marker Count:2}
after value receiver preview: {Name:whiteboard marker Count:2}, preview=5
after pointer receiver add: {Name:whiteboard marker Count:5}
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| previewAdd | reads a copy and returns a new value |
| add | changes the original struct |
| the printout | shows the difference in behavior |

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
