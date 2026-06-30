# Project 11 - Webhooks

## Goal

Deliver signed events to customer endpoints with retries, deduplication, and an
observable failure policy.

## Event Envelope

```json
{
  "id": "evt_123",
  "type": "task.completed",
  "occurredAt": "2026-06-29T12:00:00Z",
  "data": {}
}
```

## Checkpoints

1. define versioned event types.
2. serialize exact bytes.
3. sign timestamp plus body with HMAC.
4. send event ID and signature headers.
5. use a strict delivery timeout.
6. retry only transient failures with backoff/jitter.
7. stop after bounded attempts.
8. move exhausted events to a dead-letter state.
9. let consumers deduplicate by event ID.
10. expose delivery status and replay controls.

## Security

- require HTTPS in production
- rotate secrets
- reject stale signatures
- avoid internal/private network destinations
- limit redirects and response body reads

## Failure Drills

1. receiver succeeds but response is lost.
2. receiver returns 400.
3. receiver returns 503.
4. receiver sleeps beyond timeout.
5. attacker replays old signed request.
6. URL targets an internal service.

## Done Means

Delivery is authenticated, duplicate-safe, bounded, observable, and manually
recoverable.

