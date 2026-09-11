# Exercise 03 - Scoreboard

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

This exercise combines data modeling with aggregation:

- a struct can contain a slice
- a loop can reduce many values into one total
- numeric types affect division
- an edge case must be handled before a dangerous operation
- one method can reuse another method

## Beginner Bridge: What Aggregation Means

Aggregation means reducing many values into one answer.

Examples:

- many scores become one average
- many expenses become one total
- many ratings become one rating summary
- many request times become one latency number

The program needs a memory variable that carries progress through the loop. That
variable is usually called an accumulator.

```text
total starts at 0
add first score
add second score
add third score
now total contains all previous work
```

In Go:

```go
total := 0
for _, score := range student.Scores {
    total += score
}
```

If `total` is inside the loop, it gets recreated every iteration and forgets the
previous scores. If you divide before checking for zero scores, the program can
crash. This exercise is about learning where state belongs and which edge case
must be handled before arithmetic.

## Model Variable-Length Data

A student can have any number of scores:

```go
type Student struct {
    Name   string
    Scores []int
}
```

`Scores []int` is better than `Score1`, `Score2`, and `Score3`. The algorithm
then works for zero, two, or one hundred scores without changing the type.

Example value:

```go
student := Student{
    Name:   "Sara",
    Scores: []int{80, 70, 90},
}
```

## Aggregation: Many Values Become One

To calculate an average, first calculate a total:

```text
start total at 0
visit 80 -> total becomes 80
visit 70 -> total becomes 150
visit 90 -> total becomes 240
divide 240 by 3
```

Trace:

| Iteration | Score | Total before | Total after |
|---:|---:|---:|---:|
| 1 | 80 | 0 | 80 |
| 2 | 70 | 80 | 150 |
| 3 | 90 | 150 | 240 |

The variable `total` is an **accumulator**. It remembers the result of previous
iterations.

## Integer Division Is Deliberately Strict

In Go:

```go
5 / 2 // 2
```

Both operands are integers, so the result is an integer. The fractional part is
discarded.

For a decimal average:

```go
float64(total) / float64(len(student.Scores))
```

Conversion does not change the original variables. It creates `float64` values
for this expression.

## Handle Empty Input Before Division

An empty slice has length zero. Division by zero is invalid.

Pseudocode:

```text
FUNCTION average(student)
    IF number of scores is zero
        RETURN 0

    total = 0
    FOR each score
        add score to total

    RETURN decimal total divided by number of scores
```

The early return proves that the later divisor is not zero.

## Reuse Behavior

`passed` should ask `average` for the result:

```text
FUNCTION passed(student)
    RETURN average(student) is at least 50
```

Do not duplicate the average loop in both methods. Duplicate logic can drift:
one copy may later handle empty input while the other does not.

In Go, comparison already returns a boolean:

```go
return student.average() >= 50
```

An `if` that returns `true` or `false` is legal but unnecessary.

## Worked Example: Product Rating

```go
package main

import "fmt"

type Product struct {
    Name    string
    Ratings []int
}

func (product Product) averageRating() float64 {
    if len(product.Ratings) == 0 {
        return 0
    }

    total := 0
    for _, rating := range product.Ratings {
        total += rating
    }

    return float64(total) / float64(len(product.Ratings))
}

func (product Product) recommended() bool {
    return product.averageRating() >= 4
}

func main() {
    product := Product{Name: "Keyboard", Ratings: []int{5, 4, 4}}
    fmt.Printf("%s: %.2f, recommended: %t\n",
        product.Name,
        product.averageRating(),
        product.recommended(),
    )
}
```

Important formatting verbs:

- `%s`: string
- `%.2f`: floating-point value with two digits after the decimal
- `%t`: boolean
- `\n`: newline

The example transfers directly to response-time averages, order totals, and
sensor measurements.

## Your Program

Build:

1. A student type with name and scores.
2. An average method returning `float64`.
3. A passed method using a threshold of 50.
4. At least three students.
5. A loop printing each name, average, and result.

## Build In Checkpoints

1. Create one student and print every score.
2. Add the accumulator and print the total.
3. Add the decimal division.
4. Handle the empty slice.
5. Add `passed`.
6. Create and print multiple students.

```bash
gofmt -w .\main.go
go run .
```

## Test Table

| Scores | Expected average | Passed? |
|---|---:|---|
| `80, 70, 90` | 80.00 | true |
| `40, 44, 50` | 44.67 | false |
| `50` | 50.00 | true |
| empty | 0.00 | false |
| `49, 50` | 49.50 | false |

## Failure Experiments

1. Remove the empty-slice guard and use no scores. Observe the failure.
2. Divide before converting to `float64`. Compare the lost fractional part.
3. Initialize `total` inside the loop. Explain why it forgets previous scores.
4. Hardcode division by 3, then add a fourth score. Explain the wrong result.

## You Understand This Exercise When

You can draw the accumulator trace, explain conversion and integer division,
defend the empty-input behavior, and apply the same pattern to a different list
of numbers.

References:

- [Slices](https://go.dev/tour/moretypes/7)
- [`range`](https://go.dev/tour/moretypes/16)
- [Methods](https://go.dev/tour/methods/1)
