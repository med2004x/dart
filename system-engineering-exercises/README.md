# Systems Engineering Project Track

This folder turns systems concepts into projects you can build, break, measure,
and explain.

Do the projects in order. Each project adds one system responsibility while
keeping the previous ones understandable.

Use the short [Quick Reference](QUICK-REFERENCE.md) for the recurring Go and
systems patterns. It is intentionally brief; each project README defines the
actual experiment and evidence to produce.

## Working Method

For every project:

1. Read its `README.md`.
2. Write your prediction before running anything.
3. Run the starter.
4. Implement one checkpoint at a time.
5. Trigger every required failure.
6. Save sanitized evidence in an `evidence` directory.
7. Explain the root cause and recovery.

Required evidence:

```text
command.txt       exact command
expected.txt      prediction
actual.txt        observed result
explanation.md    why the behavior occurred
```

Do not store credentials, API keys, personal data, or full production logs.

## Project Map

| Project | Build | Main proof |
|---|---|---|
| 01 Request Boundaries | instrumented HTTP server | locate a failure boundary |
| 02 Requirements Capacity | capacity model | requirements become numbers |
| 03 Modular Monolith | service/repository core | rules run without HTTP |
| 04 Transactions Integrity | atomic repository | failed writes leave no state |
| 05 Timeouts Retries | bounded HTTP client | dependency waits are bounded |
| 06 Scaling Backpressure | limited worker server | overload is controlled |
| 07 Consistency Events | outbox projection | delayed/duplicate events recover |
| 08 Reliability Recovery | graceful server | in-flight work drains |
| 09 Observability | structured request telemetry | one request is traceable |
| 10 Security | ownership checks | client IDs cannot grant access |
| 11 Safe Deployments | compatible migration | old and new code overlap |
| 12 Architecture Decisions | measured ADR | design follows evidence |
| 13 System Evolution | staged task system | complexity follows requirements |
| 14 Design Capstone | complete vertical slice | happy and failure paths work |
| 15 Architecture Review | executed audit | findings cite evidence |

## Commands

Compile every Go starter:

```powershell
Set-Location C:\Users\pc\Documents\dart\system-engineering-exercises
go test ./...
go vet ./...
```

Run one project:

```powershell
Set-Location .\01-request-boundaries
go run .
```

Later projects may require PostgreSQL. Complete the corresponding projects in
[`../postgresql-exercises`](../postgresql-exercises) first.

## Mastery Standard

For each system, you must answer:

- What user goal does it serve?
- Which component owns each rule and datum?
- What are the measurable capacity and reliability targets?
- What happens when every dependency fails?
- How is overload bounded?
- Which signal identifies the failing boundary?
- How is state recovered?
- What evidence proves the answer?

