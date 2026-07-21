# Exercise 04 - Expense Tracker

If a syntax item is unfamiliar, use the [track quick reference](../../GO-QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## What You Are Learning

This exercise introduces maps and two forms of aggregation:

- total every value
- group values by a key, then total each group

The critical distinction is between raw records and an index or summary built
from those records.

## Beginner Bridge: Why A Map Exists

A slice answers: "What records do I have, in order?"

A map answers: "What value belongs to this key?"

For expenses, the slice stores every event:

```text
lunch, food, 1200
bus, transport, 300
coffee, food, 250
```

The map stores a summary:

```text
food -> 1450
transport -> 300
```

That is why the slice remains the source of truth and the map is calculated from
it. If you only keep the map, you lose the original expense names. If you only
keep the slice, answering "total by category" requires scanning everything each
time.

Similar programs:

| Program | Records | Map summary |
|---|---|---|
| Sales | `[]Sale` | revenue by day |
| Logs | `[]LogEntry` | count by level |
| Votes | `[]Vote` | votes by candidate |

The central move is `totals[key] += amount`: read the old value, add the new
amount, store the result back under the same key.

## Slice Versus Map

A slice stores records in order:

```go
expenses := []Expense{
    {Title: "lunch", Category: "food", Amount: 1200},
    {Title: "bus", Category: "transport", Amount: 300},
}
```

A map associates unique keys with values:

```go
totals := map[string]int{
    "food":      1200,
    "transport": 300,
}
```

Read `map[string]int` as "a map whose keys are strings and whose values are
integers."

Use the slice as the source of truth. Build the map as a summary.

## Map Lookup And Zero Values

Reading a missing `int` map entry returns zero:

```go
totals := map[string]int{}
fmt.Println(totals["food"]) // 0
```

That makes accumulation concise:

```go
totals["food"] = totals["food"] + 500
```

or:

```go
totals["food"] += 500
```

Trace:

| Expense | Previous category total | Amount | New total |
|---|---:|---:|---:|
| lunch / food | 0 | 1200 | 1200 |
| bus / transport | 0 | 300 | 300 |
| coffee / food | 1200 | 250 | 1450 |

## Maps Must Be Initialized

This declares a nil map:

```go
var totals map[string]int
```

Reading is allowed, but assigning a key causes a run-time panic.

Initialize before writing:

```go
totals := make(map[string]int)
```

or:

```go
totals := map[string]int{}
```

## Money As Integers

Binary floating-point cannot exactly represent many decimal fractions. Financial
programs normally store the smallest currency unit:

```text
12.50 currency units -> 1250 cents
```

This exercise uses `int`. A real financial system would also specify currency,
overflow policy, rounding rules, and possibly use a fixed-precision decimal
library.

## Pseudocode

Total all expenses:

```text
FUNCTION total(expenses)
    total = 0
    FOR each expense
        total = total plus expense amount
    RETURN total
```

Group by category:

```text
FUNCTION totals by category(expenses)
    create empty map
    FOR each expense
        key = expense category
        map[key] = map[key] plus expense amount
    RETURN map
```

## Worked Example: Votes By Candidate

```go
package main

import "fmt"

type Vote struct {
    District  string
    Candidate string
    Count     int
}

func votesByCandidate(votes []Vote) map[string]int {
    totals := make(map[string]int)

    for _, vote := range votes {
        totals[vote.Candidate] += vote.Count
    }

    return totals
}

func main() {
    votes := []Vote{
        {District: "north", Candidate: "A", Count: 12},
        {District: "south", Candidate: "B", Count: 9},
        {District: "west", Candidate: "A", Count: 7},
    }

    for candidate, total := range votesByCandidate(votes) {
        fmt.Println(candidate, total)
    }
}
```

Map iteration order is not guaranteed. Candidate A may print before or after B.
If stable output matters, collect and sort the keys before printing.

## Your Program

Build:

1. An expense type with title, category, and integer amount.
2. At least five expenses.
3. A function returning the grand total.
4. A function returning `map[string]int` totals by category.
5. Output for the grand total and every category.

## Build In Checkpoints

1. Create and print the expense slice.
2. Calculate one grand total.
3. Create an initialized empty map.
4. Add each expense to its category.
5. Print the map.
6. Add repeated and previously unseen categories.

```powershell
gofmt -w .\main.go
go run .
```

## Test Table

| Input | Expected |
|---|---|
| no expenses | total 0, empty map |
| one food expense of 500 | grand total 500, food 500 |
| food 500 + food 200 | grand total 700, food 700 |
| food 500 + travel 200 | two map keys |
| zero amount | totals unchanged numerically |
| negative amount | decide whether to reject in a later validation exercise |

This exercise does not yet return validation errors. Record the negative-amount
gap; do not hide it.

## Failure Experiments

1. Use a nil map and assign a key. Read the panic.
2. Replace `+=` with `=`. Explain why repeated categories lose prior values.
3. Run the program several times and observe map output order.
4. Use `float64` amounts such as `0.1` and inspect precision with many decimal
   places.

## You Understand This Exercise When

You can explain why the slice and map serve different purposes, trace a repeated
category, initialize maps correctly, and adapt the pattern to sales by day or
errors by endpoint.

References:

- [Maps](https://go.dev/tour/moretypes/19)
- [Mutating maps](https://go.dev/tour/moretypes/22)
