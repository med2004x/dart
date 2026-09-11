# Step 11: API Engineering

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- concurrent updates need a conflict rule
- versions, timestamps, or locks protect shared state
- the client must know when a write lost the race

## Beginner Bridge

Start from a plain HTTP handler or request struct. Then add the new contract rule only where it is needed.

### Before
```go
package main

import "fmt"

func main() {
    fmt.Println("write the contract, then the handler")
}
```

### After
```go
package main

import "fmt"

func helper() string {
    return "reject stale writes with a version rule"
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

func main() {
    fmt.Println("write with version check, reject stale updates")
}
```

### Expected output
```text
write with version check, reject stale updates
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| version | detects stale state |
| conflict | signals a lost race |
| client | must reload and retry |

## Design / Reasoning Before Syntax

1. Write the request the client sends.
2. Write the response the client should observe.
3. Choose the boundary checks that protect the handler.
4. Decide which status code describes each outcome.
5. Keep the contract and the code in lockstep.

This is the proof path. The code should match it instead of inventing a new shape after the fact.

## Your Program / Tasks

1. Describe the contract, then make the HTTP shape match it.
2. Write the observable request and response first.
3. Keep transport details separate from the business meaning.

## Build In Checkpoints

1. Name the resource and the action before writing the handler.
2. Choose the status code and response shape before the implementation.
3. Add one success path and one failure path.
4. Check that the boundary behavior matches the contract exactly.

## Failure Drills

1. Return 200 for everything. Why: the client cannot tell success from failure.
2. Blend validation, auth, and storage together. Why: the contract becomes unreadable.
3. Hide a breaking change inside the path or body. Why: old clients will fail with no warning.

## You Understand This When / Done Means

- Can you state the contract in one sentence?
- Can you point to the exact method, status, or field that proves each branch?
- Can you explain why a client would trust this API shape?
