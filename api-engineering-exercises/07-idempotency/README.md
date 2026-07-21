# Project 07 - Idempotent Creates

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Make repeated create requests return the original result instead of creating
duplicates.

## Failure Scenario

```text
client sends create
server commits task
response is lost
client retries
```

The retry is indistinguishable from a new request unless the client provides an
idempotency key.

## Required Contract

```text
Idempotency-Key: client-generated-value
```

Store atomically:

```text
caller scope
key
request fingerprint
status
response body
expiry
```

## Checkpoints

1. validate key format/length.
2. calculate a canonical request fingerprint.
3. create result and key record in one transaction.
4. replay stored response for matching retry.
5. return 409 when same key has different payload.
6. handle two concurrent requests with one key.
7. define retention and cleanup.

## Failure Drills

1. store key after creating the task.
2. use key without caller scope.
3. retry while first request is still processing.
4. reuse key with a different title.

## Done Means

All retries and races for one scoped key produce at most one business effect.

