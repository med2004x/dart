# API Design And Engineering Project Track

This track teaches APIs as contracts and operated systems, not collections of
handlers.

Each project produces a runnable behavior or a reviewable contract. Complete
them in order.

Use the short [Quick Reference](QUICK-REFERENCE.md) when HTTP or Go handler
syntax is unfamiliar. It gives generic examples and keeps the project-specific
requirements here focused.

## Project Map

| Project | Main result |
|---|---|
| 01 Requirements | API problem brief |
| 02 Resource Modeling | ownership and URL model |
| 03 OpenAPI Contract | machine-readable contract |
| 04 HTTP Semantics | correct methods and statuses |
| 05 Validation And Errors | stable error behavior |
| 06 Query Design | pagination/filter/sort contract |
| 07 Idempotency | duplicate-safe creates |
| 08 Concurrent Updates | conflict-safe writes |
| 09 Authentication/Authorization | trusted identity and ownership |
| 10 Rate Limits | controlled overload |
| 11 Webhooks | signed, retryable events |
| 12 Asynchronous Jobs | 202 and job lifecycle |
| 13 Versioning | backward-compatible evolution |
| 14 Verification/Operations | tests, telemetry, shutdown |
| 15 Capstone | complete project/task API |

## Work Method

For every project:

1. Define the observable contract.
2. List invalid and failure cases.
3. Write examples before implementation.
4. Implement the smallest behavior.
5. Test at the HTTP boundary.
6. Trigger dependency and overload failures.
7. Record compatibility and security consequences.

## Required API Evidence

```text
request.txt
response.txt
test-output.txt
contract.yaml or contract.md
failure-explanation.md
```

Never include real credentials or personal data.

## Commands

```powershell
Set-Location C:\Users\pc\Documents\dart\api-engineering-exercises
gofmt -w .
go test ./...
go vet ./...
```

## Mastery Standard

You can explain:

- who owns each resource and field
- which client inputs are trusted
- every method and status code
- retry and duplicate behavior
- pagination stability
- concurrent update behavior
- authentication versus authorization
- compatibility policy
- rate, errors, latency, and saturation evidence

