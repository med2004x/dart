# Project 12 - Asynchronous Job APIs

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Turn a long report request into an explicit job lifecycle.

## Contract

Create:

```text
POST /reports
-> 202 Accepted
Location: /jobs/{jobId}
```

Poll:

```text
GET /jobs/{jobId}
```

States:

```text
queued -> running -> succeeded
                 `-> failed
queued/running -> canceled
```

## Checkpoints

1. define allowed state transitions.
2. store job before returning 202.
3. make create idempotent.
4. let a bounded worker claim jobs.
5. record attempts and error category.
6. define cancellation semantics.
7. expire old results.
8. return result URL only after success.
9. expose queue depth and oldest-job age.

## Failure Drills

1. worker crashes after completing work but before marking success.
2. client repeats create.
3. cancellation arrives while work is finishing.
4. producers exceed workers for ten minutes.
5. result storage fails after job work.

## Done Means

Clients can determine accepted, running, failed, canceled, and completed states
without holding one HTTP request open.

