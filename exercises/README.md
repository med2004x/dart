# Go Programming From First Principles

This folder is a practical course for learning how programs work, not a list of
syntax puzzles. The exercises begin with values in memory and end with an HTTP
API split into clear layers.

The goal is to move through three stages:

1. **Hope:** "I wrote some code and it printed the right thing."
2. **Reason:** "I can trace every value and explain why the result is right."
3. **Control:** "I can predict failures, test them, and design a similar program."

Running once is only stage one.

Use the workspace [Go Quick Reference](../GO-QUICK-REFERENCE.md) when syntax is
the blocker. It contains short generic examples and official documentation
links; the exercise READMEs remain focused on requirements and proof cases.

## Before Exercise 01

Install Go, then verify the installation in PowerShell:

```powershell
go version
go env GOROOT
go env GOPATH
```

`go version` should print an installed Go version. `GOROOT` is where Go itself
is installed. `GOPATH` is a workspace Go uses for downloaded tools and cached
module data. Your project does not need to live inside `GOPATH`.

Move into the exercise module and check that Go can see it:

```powershell
Set-Location C:\Users\pc\Documents\dart\exercises
Get-Content .\go.mod
go list ./...
```

The `go.mod` file names the module and records the Go language version. A module
is a group of related Go packages. Each numbered directory in this course is a
separate `main` package that can be run.

## What Happens When You Run Go Code

When you enter:

```powershell
Set-Location .\01-bank-account
go run .
```

Go does not execute `main.go` line by line as raw text. It:

1. Finds the module by looking for `go.mod`.
2. Collects all `.go` files in the current directory that belong to the current
   operating system and package.
3. Parses the source code.
4. Checks syntax, imports, names, and types.
5. Compiles a temporary executable.
6. Starts that executable at `main.main`.
7. Deletes the temporary executable after it exits.

That distinction explains two kinds of failure:

- A **compile-time failure** means Go could not build the program. Examples:
  misspelled names, wrong types, unused imports, and missing return values.
- A **run-time failure** means the program compiled and started, then encountered
  a bad state. Examples: dividing by zero, indexing outside a slice, or failing
  to open a file.

The compiler proving that code is legal does not prove that its behavior is
correct. Tests and deliberate edge cases do that.

## The Learning Loop

Use the same loop for every exercise.

### 1. Predict

Before running code, write down:

- the inputs
- the expected output
- the values that should change
- the values that must not change
- at least one invalid input

### 2. Write pseudocode

Pseudocode describes logic without Go syntax:

```text
FUNCTION find item by ID
    FOR each item in the list
        IF the item's ID equals the wanted ID
            RETURN the item and true
    RETURN an empty item and false
```

If the pseudocode is confused, Go syntax will not repair the design.

### 3. Translate one step at a time

Write the smallest useful part, format it, compile it, and run it:

```powershell
gofmt -w .\main.go
go run .
```

`gofmt` gives Go code one standard layout. Formatting early makes missing braces
and malformed blocks easier to see.

### 4. Trace

For a loop, make a table:

| Iteration | Current value | State before | State after |
|---:|---|---:|---:|
| 1 | 4 | 0 | 4 |
| 2 | 7 | 4 | 11 |
| 3 | 2 | 11 | 13 |

This is how you replace guessing with evidence.

### 5. Break it deliberately

Try empty input, zero, negative values, missing records, and malformed data.
Predict each result before running it. If the program behaves differently, trace
the first line where reality diverges from the prediction.

### 6. Explain it without the editor

You understand an exercise when you can explain:

- what data exists and where it is stored
- the type of every function input and output
- which values are copies and which values can mutate shared state
- every branch that can run
- what happens for invalid input
- why each imported package is needed

## Reading Compiler Errors

Read the first error first. Later errors are often consequences.

Example:

```text
.\main.go:14:20: cannot use "10" (untyped string constant) as int value
```

Read it in parts:

- `.\main.go`: file containing the problem
- `14:20`: line 14, column 20
- `cannot use "10"`: the value supplied
- `as int value`: the type the surrounding code requires

Do not randomly change nearby code. Go to the exact location, identify the
actual and required types, and decide where conversion or a better value belongs.

Useful diagnostic commands:

```powershell
go run .
go test ./...
go vet ./...
go doc fmt.Println
go doc builtin.append
```

`go vet` finds suspicious code that may compile. `go doc` explains packages,
types, and functions from the terminal.

## Exercise Map

| Exercise | Main idea | New responsibility |
|---|---|---|
| 01 Bank Account | structs and methods | protect state changes |
| 02 Auth Users | slices and search | report found/not found |
| 03 Scoreboard | loops and arithmetic | aggregate safely |
| 04 Expense Tracker | maps | group values by key |
| 05 Contact Book | CRUD | design complete operations |
| 06 Text Analyzer | strings and counting | normalize before comparing |
| 07 File Notes | files and JSON | persist data and handle I/O failure |
| 08 Errors and Validation | error values | separate rules from presentation |
| 09 Payment Interfaces | interfaces | depend on required behavior |
| 10 HTTP Basics | requests and responses | serve multiple clients |
| 11 JSON API | encoding and status codes | define an API contract |
| 12 CRUD API | routes and in-memory state | manage resources over HTTP |
| 13 Middleware Auth | handler composition | apply cross-cutting policy |
| 14 Service/Repository | layer boundaries | separate transport, rules, storage |
| 15 Capstone | complete request flow | combine and defend the design |
| 16 Product Catalog CRUD | slice CRUD repetition | make CRUD contracts automatic |
| 17 Inventory CRUD Methods | method receivers | mutate owned state correctly |
| 18 Notes CRUD With Errors | validation and errors | separate failure reasons |
| 19 Map-Backed CRUD | maps as stores | compare direct lookup with slice search |
| 20 CRUD Mini Project | combined CRUD design | choose methods, errors, and state ownership |

The order matters. Later exercises assume the earlier ideas are ordinary.

## What The README Must Do For You

Each exercise README is supposed to teach the idea before asking you to code it.
Use it in this order:

1. Read the beginner explanation until you can say the concept in plain words.
2. Read the pseudocode and trace table before touching Go syntax.
3. Run the small checkpoints one at a time.
4. Do the failure experiments on purpose.
5. Only then move to the next exercise.

Do not treat external Go documentation as the main teacher for these exercises.
The docs are useful references, but the exercise README should give you the
working mental model first: what the concept is, why it exists, what breaks, and
how a similar program would use the same pattern.

## PowerShell Reference

From any exercise directory:

```powershell
# Show files.
Get-ChildItem

# Read the current source.
Get-Content .\main.go

# Format all Go files in this directory.
gofmt -w .

# Compile and run the package.
go run .

# Build an executable without running it.
go build .

# Run all tests below the exercise module.
Set-Location C:\Users\pc\Documents\dart\exercises
go test ./...
```

For HTTP exercises, keep the server running in one PowerShell window and send
requests from a second window:

```powershell
Invoke-RestMethod -Uri http://localhost:8080/health
```

For status code and header details, use:

```powershell
Invoke-WebRequest -Uri http://localhost:8080/health
```

PowerShell aliases `curl` to different commands on some versions. Use
`curl.exe` when you specifically want the native curl command.

## Rules For Getting Help

Bring evidence, not "it does not work." Include:

1. the command you ran
2. the complete first error
3. the input that triggered it
4. the output you expected
5. the output you received
6. your current explanation of where the values diverge

That information turns debugging into a technical process.
