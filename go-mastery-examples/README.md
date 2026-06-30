# Go Mastery Worked Examples

These are reference examples, not exercises to memorize.

Use each example in four passes:

1. Read the code and predict every printed line.
2. Run it and compare the real result with the prediction.
3. Change one input or rule and predict again.
4. Rebuild the same idea in a different small program without looking.

## Before Running

Verify Go:

```powershell
go version
```

Each example is its own module. Run commands from that example's directory.

## 01 - Basics

Path: `01-basics`

Concepts:

- package and `main`
- variables and inferred types
- function calls
- formatted terminal output

Run:

```powershell
Set-Location .\01-basics
go run .
```

Study questions:

- Why does execution begin in `main`?
- Which variable types are inferred?
- Which errors would the compiler catch before execution?

Transfer program: create a temperature converter with one input variable, one
calculation function, and formatted output.

## 02 - Slices And Maps

Path: `02-slices-maps`

Concepts:

- ordered slices
- key/value maps
- `range`
- accumulating grouped totals

Run:

```powershell
Set-Location ..\02-slices-maps
go run .
```

Before running, trace each loop in a table. Then add a repeated value and predict
which count changes.

Transfer program: count support tickets by status.

## 03 - Testing

Path: `03-testing`

Concepts:

- production code and `_test.go` files
- table-driven test cases
- expected versus actual behavior
- test failure messages

Run:

```powershell
Set-Location ..\03-testing
go test -v ./...
```

Deliberately change one expected value and read the failure:

```text
test name -> input -> expected -> actual
```

Restore it afterward.

Transfer program: write tests for a shipping-cost function with zero, boundary,
and invalid inputs.

## 04 - Concurrency

Path: `04-concurrency`

Concepts:

- goroutines
- channels
- concurrent completion order
- waiting for results

Run:

```powershell
Set-Location ..\04-concurrency
go run .
go run -race .
```

Concurrency does not guarantee execution order. Separate these questions:

- Are operations allowed to overlap?
- Is shared memory synchronized?
- In what order are results consumed?

Transfer program: run three independent simulated checks and collect all results.
Do not add concurrency to work that is already fast and sequential without a
measured reason.

## 05 - HTTP API

Path: `05-http-api`

Concepts:

- handlers and routes
- JSON input/output
- status codes
- in-memory state

Start it:

```powershell
Set-Location ..\05-http-api
go run .
```

Use a second terminal:

```powershell
curl.exe -i http://localhost:8080/health
```

Trace one request:

```text
method/path -> route -> handler -> validation -> state -> response
```

Transfer program: create a two-route book API with health and one create
operation. Define the request and response contract before coding.

## 06 - Generics

Path: `06-generics`

Concepts:

- type parameters
- constraints
- one algorithm over multiple concrete types

Run:

```powershell
Set-Location ..\06-generics
go run .
```

Do not use generics merely to remove two readable lines. Use them when the
algorithm is genuinely the same for multiple types and type safety is preserved.

Transfer program: write a `contains` function for comparable values, then compare
its readability with two concrete implementations.

## Validate Every Example

From this folder, test and vet each module explicitly:

```powershell
$modules = Get-ChildItem -Directory

foreach ($module in $modules) {
    Push-Location $module.FullName
    try {
        Write-Host "Checking $($module.Name)"
        go test ./...
        go vet ./...
    }
    finally {
        Pop-Location
    }
}
```

One module's success does not prove another module builds. The loop changes into
each module and checks it independently.
