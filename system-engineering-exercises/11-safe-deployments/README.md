# Project 11 - Safe Deployments And Migrations

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- safe deploys reduce blast radius
- rollouts, health checks, and rollback rules need to be explicit
- the deployment path should fail small

## Beginner Bridge

Start from one service or one boundary. Then add the new control rule only where it is needed.

### Before
```go
package main

import "fmt"

func main() {
    fmt.Println("draw the boundary before the implementation")
}
```

### After
```go
package main

import "fmt"

func helper() string {
    return "roll out changes with canary and rollback rules"
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
    fmt.Println("canary, health check, rollback")
}
```

### Expected output
```text
canary, health check, rollback
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| canary | limits blast radius |
| health check | detects bad rollout |
| rollback | restores the known good version |

## Design / Reasoning Before Syntax

1. Write the boundary or control rule first.
2. Name the failure mode and the recovery path.
3. Decide which component owns the state change.
4. Add numbers or limits so the design can be checked.
5. Keep the plan easy to trace during review.

This is the proof path. The code should match it instead of inventing a new shape after the fact.

## Your Program / Tasks

1. Reason about the system before you write implementation code.
2. Draw the boundary or control rule before you write the implementation details.
3. Keep the load, failure, and recovery story visible in the text.

## Build In Checkpoints

1. Write the assumption or boundary rule first.
2. Add the simplest path that proves the idea.
3. Add the failure path or recovery path second.
4. Check that the design still makes sense when the load or failure grows.

## Failure Drills

1. Handwave the numbers. Why: design without numbers is theater.
2. Hide the boundary. Why: the caller and callee will fight over ownership.
3. Describe recovery only in prose. Why: the real recovery path must be concrete.

## You Understand This When / Done Means

- Can you explain where responsibility changes hands?
- Can you name the failure mode and the recovery path?
- Can you show the decision that keeps the system within its limits?
