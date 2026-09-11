# Systems Engineering Quick Reference

Use this page to translate a system requirement into an implementation and
evidence. The project README remains the source of the specific task.

## Ownership

```text
handler/controller -> HTTP parsing and response
service            -> business rules
repository         -> persistence
domain             -> state and domain behavior
```

Dependencies should point inward. Keep the core rule testable without HTTP,
PostgreSQL, or a network client.

## Context And Deadlines

```go
ctx, cancel := context.WithTimeout(parent, 500*time.Millisecond)
defer cancel()

if err := dependency.Call(ctx); err != nil {
	return fmt.Errorf("call dependency: %w", err)
}
```

Pass `context.Context` through external work. Always release the cancel
function. Decide which failures are retryable before writing a retry loop.

## Bounded Retries

```text
for attempt from 1 through maximum attempts
    call dependency with a deadline
    if success, return
    if error is not retryable, return
    wait with bounded exponential backoff and jitter
return the last error
```

Retries multiply load. Use them only for transient failures, cap the attempts
and total time, and require the operation to be safe to repeat.

## Channels And Worker Limits

```go
jobs := make(chan Job)
results := make(chan Result)

go worker(ctx, jobs, results)
```

Bound queues and worker counts. Decide what happens when the queue is full:
reject, block briefly, or shed work. An unbounded queue only moves failure into
memory.

## Transactions And Recovery

For multi-step writes, define the atomic boundary first. If a later step can
fail, earlier state must be rolled back or an explicit recovery record must
exist. Test interruption between steps, not only the successful path.

## Observability

Record enough structured data to answer: which request, which component, which
entity, how long, and what failed. Avoid passwords, tokens, and unnecessary
personal data. Measure rate, errors, latency percentiles, and saturation.

## Architecture Decisions

An ADR should state:

```text
context -> decision -> consequences -> rejected alternatives
```

Choose the simplest architecture that satisfies measured requirements. Add
distributed components only when an actual boundary, scale, or ownership need
requires them.

## Official Resources

- [Go `context`](https://pkg.go.dev/context)
- [Go `net/http`](https://pkg.go.dev/net/http)
- [Go `sync`](https://pkg.go.dev/sync)
- [Go `os/signal`](https://pkg.go.dev/os/signal)
- [Go race detector](https://go.dev/doc/articles/race_detector)
- [Go memory model](https://go.dev/ref/mem)
- [HTTP semantics, RFC 9110](https://www.rfc-editor.org/rfc/rfc9110)
