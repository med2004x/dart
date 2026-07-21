# Go Mastery Learning Workspace

This workspace teaches Go and backend systems through connected project tracks:

- [`exercises`](exercises): twenty staged programs built by the learner
- [`functions-methods-exercises`](functions-methods-exercises):
  fifteen focused projects for learning function calls, parameters, returns,
  methods, receivers, and when to choose each form
- [`go-mastery-examples`](go-mastery-examples): runnable worked examples
- [`system-engineering-exercises`](system-engineering-exercises):
  fifteen hands-on systems projects covering boundaries, capacity, reliability,
  security, deployments, evolution, and architecture review
- [`postgresql-exercises`](postgresql-exercises):
  seventeen cumulative PostgreSQL projects covering SQL, transactions,
  concurrency, performance, recovery, and Go integration
- [`api-engineering-exercises`](api-engineering-exercises):
  fifteen contract-first projects covering HTTP semantics, validation,
  pagination, idempotency, authorization, webhooks, and API operations
- [`ps-exercises`](ps-exercises):
  fifteen beginner problem-solving projects that teach every required algorithm
  from pseudocode and hand traces before Go implementation
- [`cs50`](cs50): CS50 setup and problem-set workspace for the C course,
  using WSL-based compiler tools on this Windows machine

Despite the parent folder name, this is a Go workspace.

For a compact syntax and resources guide, start with [Go Quick Reference](GO-QUICK-REFERENCE.md).

## Recommended Route

If programming itself is new:

1. Read the exercise track's sections on execution, compiler errors, pseudocode,
   tracing, and deliberate failure.
2. Complete `functions-methods-exercises` 01-08 alongside the first core
   exercises if functions and returns feel unclear.
3. Complete exercises 01 through 06 in order.
4. Use worked examples to study a pattern, then implement a different program
   using that pattern.
5. Continue through files, errors, interfaces, HTTP, layering, and the capstone.
6. Use PS projects 01-05 alongside the core Go exercises, then continue the
   algorithm projects in order.
7. Complete PostgreSQL before the database-backed API projects.
8. Complete API engineering before the final systems capstone and architecture
   review.

If JavaScript or Python is already familiar, move faster through Go exercises
01-06 but still complete their failure experiments.

## The Standard For Understanding

Code that runs once is not enough. For every program, be able to answer:

- What does the compiler verify before execution?
- Where is each value stored?
- Which statement changes state?
- Which values are copied?
- What are all valid and invalid inputs?
- Which failures occur at compile time and which at run time?
- What evidence proves the result is correct?
- How would the same algorithm appear in a different program?

If an answer depends on "Go handles it somehow," trace that part again.

## First Commands

```powershell
Set-Location C:\Users\pc\Documents\dart
go version
Get-Content .\exercises\README.md
Set-Location .\exercises\01-bank-account
go run .
```

Use [`exercises/README.md`](exercises/README.md) as the operating guide for the
course.
