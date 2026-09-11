# Step 05: Problem Solving and Algorithms

If a syntax item is unfamiliar, use the local quick reference in this track. It contains syntax examples and helper notes without solving this project.

## What You Are Learning

- two pointers walk the data from one or both ends
- the pointers move based on what the current pair tells you
- the state is the gap between the pointers

## Beginner Bridge

Start from a direct loop or trace. Then add the pattern only where it is needed.

### Before
```go
package main

import "fmt"

func main() {
    fmt.Println("start with the direct version")
}
```

### After
```go
package main

import "fmt"

func helper() string {
    return "use two positions to narrow the work"
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

func hasPair(nums []int, target int) bool {
    left, right := 0, len(nums)-1
    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {
            return true
        }
        if sum < target {
            left++
            continue
        }
        right--
    }
    return false
}

func main() {
    nums := []int{1, 3, 4, 6}
    fmt.Println(hasPair(nums, 7))
}
```

### Expected output
```text
true
```

### Transfer the pattern, not the names

| In the example | In this exercise |
|---|---|
| left/right | move toward each other |
| sum | decides which pointer moves |
| return true | stops once the answer is found |

## Design / Reasoning Before Syntax

1. Write the input shape and the smallest useful state.
2. Choose the one variable that proves progress.
3. Trace every step with a table or a short list.
4. Keep the stop rule visible before you optimize.
5. Compare the final answer against the trace.

This is the proof path. The code should match it instead of inventing a new shape after the fact.

## Your Program / Tasks

1. Solve the pattern with a direct trace first.
2. Write the direct version first, even if it looks boring.
3. Prove the output with a trace table or a small hand walk.

## Build In Checkpoints

1. Define the input shape and the stop condition.
2. Write the smallest loop or recursive step.
3. Add the boundary case before you touch the edge cases.
4. Compare your result with a hand trace.

## Failure Drills

1. Start with a clever trick before the direct version works. Why: you will not know what you are proving.
2. Skip the trace table. Why: the state changes stay hidden.
3. Forget the not-found or empty-input case. Why: that is where the bug lives.

## You Understand This When / Done Means

- Can you describe the state that changes on every step?
- Can you explain why the algorithm stops?
- Can you predict the answer without running the code?
