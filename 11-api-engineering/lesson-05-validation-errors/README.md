# Lesson 5: Validation and Error Contracts

## Objective

Build validation that happens at the API boundary, before any business logic runs, and design an error response contract precise enough that clients can programmatically act on it — not just display an opaque message to a human.

## Prerequisites

Lesson 4 (HTTP semantics — this lesson's error responses use the status codes established there), Lesson 3 (OpenAPI — validation rules ideally derive from the same schema the spec already defines, avoiding duplicated, potentially drifting validation logic).

## Learn

**Why validation belongs at the boundary, not scattered through business logic.** If input validation is interleaved with business logic (a check here, a check there, throughout the handler and deeper call stack), it becomes hard to verify you've actually covered every required check, and harder still to keep the *documented* contract (Lesson 3's schema) in sync with what's *actually* enforced in code. Validating everything at the boundary — before any business logic executes — means business logic can then safely assume its inputs are well-formed, simplifying that code and making the validation itself auditable as one coherent unit rather than scattered assertions.

**Fail fast, and specifically: validate structurally invalid input before doing anything expensive.** A request with a missing required field or wrong type should be rejected immediately, before touching a database or calling an external service — this is both a correctness practice (why do expensive work for a request you already know is invalid) and, at scale, a real resource-protection measure (a flood of malformed requests shouldn't be able to exhaust database connections or downstream API quota before your validation layer even gets a chance to reject them).

**Structured, actionable error responses.** A good error response gives a client enough to act on programmatically, not just a human-readable string: a machine-readable error code (stable across API versions, suitable for `switch`/`if` logic in client code), a human-readable message (for logging/debugging, not for driving client logic), and — critically for validation errors specifically — field-level detail (which field failed, and why) when multiple fields could each independently be invalid, so a client presenting a form to a user can highlight the *specific* problematic fields rather than showing one generic "something's wrong" message for the whole request.

**400 vs. 422, a genuinely useful distinction some APIs skip.** `400 Bad Request` is appropriate for structurally malformed requests (invalid JSON, wrong type for a field, missing required field) — the request doesn't even conform to the expected shape. `422 Unprocessable Entity` is appropriate for requests that are structurally valid JSON matching the expected shape, but semantically invalid per a business rule (e.g. an end date before a start date, a well-formed but already-taken email address). This distinction, when used consistently, lets a client's error-handling logic distinguish "I sent malformed data" (a client bug) from "my data was well-formed but violates a business rule" (often something to show the end user directly, rather than a sign of a client-side programming bug).

## Attempt

1. For one endpoint accepting a request body (from your Lesson 2/3 resource model), implement boundary validation covering: required fields present, correct types, and at least one business-rule check (something beyond pure structural validation — e.g. a date range where end must be after start, or a string field that must match a specific format like an email). Ensure all validation runs before any database or business logic call.

2. Design and implement a structured error response format: a JSON body with a stable machine-readable `code` field, a human-readable `message` field, and (for validation failures specifically) a `fields` array or object identifying which specific field(s) failed and why, for cases where multiple fields might be invalid simultaneously.

3. Test your validation with multiple simultaneous violations in one request (e.g. two required fields missing at once) and confirm your error response reports *all* the violations in one response, not just the first one encountered — a client shouldn't need to fix one field, resubmit, discover a second violation, fix that, and resubmit again, when both could have been reported together the first time.

4. Implement the 400-vs-422 distinction explicitly: return 400 for a structurally malformed request (e.g. a required field entirely missing, or wrong JSON type), and 422 for a structurally valid request that fails a business rule (e.g. all required fields present and correctly typed, but the date-range rule from step 1 is violated). Test both cases and confirm your API returns the correct, distinct status code for each.

## Verify

For step 3, show the actual JSON error response for a request with 2+ simultaneous violations, confirming all violations are reported together with clear field-level attribution. For step 4, show both the 400 and 422 test cases' actual status codes and bodies side by side.

## Failure drill

Deliberately move one of your validation checks *after* a database call in your handler (e.g. check the business rule only after already querying for a related record) rather than fully at the boundary, and construct a test case where the invalid request would trigger an unnecessary, wasted database call before ultimately being rejected. Use a query counter or log statement to confirm the database call actually happened despite the request being invalid. Explain why this specific ordering — expensive work happening before validation completes — is exactly the resource-protection failure mode Learn warned about, and is a common, easy-to-introduce mistake when validation logic isn't kept clearly and completely separated from business logic at the very start of a handler.

## Transfer

If TARDOC's clinic-registration endpoint or Mahall's product-creation endpoint currently returns a generic error message on invalid input (a plain string, or an inconsistent ad hoc format that varies by which check failed), describe how you'd migrate it to this lesson's structured error format, and specifically identify at least 2 distinct business-rule violations (beyond basic required-field/type checks) that endpoint should validate for for but that a purely structural JSON-schema-based validator (Lesson 3) wouldn't catch on its own, requiring the additional business-logic validation layer this lesson covers.

## Done when

Your endpoint's validation runs entirely at the boundary before any business logic or database access, your error responses use a consistent, structured, machine-actionable format that correctly reports multiple simultaneous violations together, you correctly distinguish 400 from 422 for structural versus business-rule failures, and you've directly demonstrated — via the failure drill — why validation placed after expensive operations defeats its own resource-protection purpose.
