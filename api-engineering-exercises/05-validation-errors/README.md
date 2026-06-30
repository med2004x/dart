# Project 05 - Validation And Error Contracts

## Goal

Reject malformed input at the HTTP boundary, enforce business rules in the
service, and return one stable error format.

## Error Shape

```json
{
  "error": "validation_failed",
  "message": "request is invalid",
  "requestId": "req-123",
  "fields": {
    "title": "title is required"
  }
}
```

Do not expose stack traces, SQL errors, file paths, or dependency secrets.

## Checkpoints

1. limit request body size.
2. require JSON content type.
3. reject malformed JSON.
4. reject unknown fields.
5. reject trailing second JSON value.
6. validate required title.
7. map known service errors.
8. map unknown errors to generic 500.
9. log internal error with request ID.

## Failure Matrix

- empty body
- malformed JSON
- wrong JSON type
- unknown field
- missing title
- whitespace title
- oversized body
- service not found
- unexpected repository failure

## Done Means

Every invalid request has deterministic status/body behavior and no internal
detail leakage.

