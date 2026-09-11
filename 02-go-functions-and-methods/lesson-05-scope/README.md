# Exercise 05 - Scope

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- locals live only inside the function that declared them
- later helpers only see values that were passed to them
- main is often the place where separate helpers are wired together

## Beginner Bridge

Start from a plain function, loop, or handler. Then add the new syntax only where it is needed.

### Before
```go
package main

import "fmt"

func main() {
    first := "Mina"
    last := "A."
    full := first + " " + last
    tax := 10
    subtotal := 90
    total := subtotal + tax
    fmt.Println(full)
    fmt.Println(total)
}
```

### After
```go
package main

import "fmt"

func buildLabel(first, last string) string {
    return first + " " + last
}

func addTax(subtotal, tax int) int {
    return subtotal + tax
}

func printReceipt(label string, total int) {
    fmt.Printf("%s -> %d\n", label, total)
}

func main() {
    first := "Mina"
    last := "A."
    subtotal := 90
    tax := 10

    label := buildLabel(first, last)
    total := addTax(subtotal, tax)
    printReceipt(label, total)
}
```

## Premade Helpers To Notice

- No special helper is needed here. The lesson is about where names live.
- `fmt.Printf` helps show the value flow across function boundaries.
- `strings.TrimSpace` is fine if the example needs input cleanup, but not as the core lesson.

## Worked Example

Use the same pattern in a different domain first. The names are different. The structure is the part to copy.

### Example code
```go
package main

import "fmt"

func buildLabel(first, last string) string {
    return first + " " + last
}

func addTax(subtotal, tax int) int {
    return subtotal + tax
}

func printTicket(label string, total int) {
    fmt.Printf("%s -> %d\n", label, total)
}

func main() {
    first := "Mina"
    last := "A."
    subtotal := 90
    tax := 10

    label := buildLabel(first, last)
    total := addTax(subtotal, tax)
    printTicket(label, total)
}
```

### Expected output
```text
Mina A. -> 100
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| buildLabel | creates one local result |
| addTax | creates the next local result |
| printTicket | uses both values only after they are wired together |

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
