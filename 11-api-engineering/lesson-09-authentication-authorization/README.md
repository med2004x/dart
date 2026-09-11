# Lesson 9: Authentication and Authorization

## Objective

Implement token-based authentication and resource-level authorization correctly, and understand precisely why these are two separate, sequential checks (who are you, then what are you allowed to do) rather than one combined concept.

## Prerequisites

Networking Lesson 5 (TLS — authentication credentials must travel over an encrypted channel, or the whole mechanism is trivially compromised by network interception), Lesson 4 (HTTP semantics — 401 vs 403 status codes, used precisely here).

## Learn

**Authentication vs. authorization: two genuinely separate questions.** Authentication answers "who is making this request" (verifying an identity claim — a password, a token, a certificate). Authorization answers "is this specific, now-known identity allowed to perform this specific action on this specific resource." Conflating them is a common source of both bugs and confusing error responses — a request with an invalid token should get `401 Unauthorized` (authentication failed — we don't know who you are), while a request with a *valid* token belonging to someone who simply isn't permitted to do this specific thing should get `403 Forbidden` (we know who you are, and the answer is no) — these are different problems requiring different client responses (401 typically means "log in again / refresh your token"; 403 means "you're logged in correctly, but you're not allowed to do this, and retrying won't help").

**Token-based authentication, the common pattern for APIs (as opposed to session cookies, more common for traditional browser-rendered sites).** A client authenticates once (e.g. via a login endpoint exchanging credentials for a token) and includes that token on every subsequent request (typically an `Authorization: Bearer <token>` header). The server validates the token on each request — either by looking it up in a server-side store (a simple, revocable approach, but requiring a database check on every request) or by validating a self-contained signed token like a JWT (JSON Web Token — the token itself carries claims about the identity, cryptographically signed so the server can verify it hasn't been tampered with, without needing a database lookup on every request, at the cost of harder immediate revocation, since a signed token remains valid until its own expiration even if you wanted to invalidate it sooner).

**Authorization models, at the level relevant for most APIs.** Role-based access control (RBAC) — a user has one or more roles (e.g. `admin`, `clinic-staff`), and permissions are defined per role — is the most common, simplest starting point. Resource-level (or "ownership-based") authorization goes further: even with the right role, a user should typically only access resources they own or are explicitly granted access to (e.g. a clinic-staff role shouldn't automatically mean access to *every* clinic's data, only their own clinic's) — this distinction matters because role-only authorization is a genuinely common, serious real-world bug source: an endpoint checking "is this user a clinic-staff member" without also checking "and does this specific clinic belong to them" allows any clinic-staff user to access any clinic's data, not just their own — a classic broken-access-control vulnerability class.

**Why authorization checks belong close to the data access, not just at a route/middleware level.** A middleware-level check ("does this route require the `admin` role") can correctly gate access to an endpoint in general, but resource-level authorization ("does *this specific* clinic ID belong to *this specific* authenticated user") typically requires knowledge only available once you've identified which specific resource the request targets — meaning at least part of the authorization logic often needs to live in or near the handler itself, not purely in generic middleware, and skipping this specific-resource check in favor of only the general role check is exactly the vulnerability class described above.

## Attempt

1. Implement a minimal token-based authentication scheme: a login endpoint that verifies credentials and issues a token (a signed JWT is a reasonable, realistic choice, using a Go library like `golang-jwt/jwt` or equivalent), and middleware that validates the token on subsequent requests, rejecting requests with missing or invalid tokens with `401 Unauthorized`.

2. Implement role-based authorization: attach a role claim to the issued token, and add middleware or handler-level logic that checks the required role for a given endpoint, rejecting requests from an authenticated-but-wrong-role user with `403 Forbidden` (distinct from the 401 in step 1 — test both cases and confirm the status codes differ correctly).

3. Implement resource-level (ownership) authorization on top of role-based checks: for an endpoint like `GET /clinics/{id}/records`, in addition to checking the user has an appropriate role, explicitly check that the requested clinic ID actually belongs to (or is otherwise legitimately accessible by) the authenticated user — and test the specific vulnerability scenario from Learn: a user with the correct role attempting to access a *different* clinic's data than their own, confirming this is correctly rejected with 403 rather than incorrectly allowed just because the role check alone passed.

4. Test all three failure modes explicitly and confirm distinct, correct status codes: no token at all (401), valid token but wrong role for the endpoint (403), valid token and correct role but wrong resource ownership (403, but via the resource-level check from step 3, not the role check from step 2 — confirm your logs or a debugger trace shows which specific check actually rejected the request, to verify you're testing the right code path).

## Verify

Show actual request/response pairs for all three rejection scenarios from step 4 (missing token, wrong role, wrong resource ownership), confirming the correct status code for each, and confirm a legitimate request (correct token, correct role, correct resource ownership) succeeds.

## Failure drill

Deliberately implement the exact vulnerability described in Learn: an endpoint that checks only the user's role (e.g. "is this user clinic-staff") without checking whether the specific requested resource ID actually belongs to them. Using a real test with two different clinic-staff users (each associated with a different clinic), confirm user A can successfully access user B's clinic's data purely because both share the `clinic-staff` role, despite having no legitimate relationship to each other's clinic. Explain, using this concrete, successfully-reproduced vulnerability, why "the user has the right role" and "the user is allowed to access this specific resource" are genuinely different checks that must both be performed — and why a code review or test suite that only exercises the happy path (a user accessing their *own* resources) would never catch this bug, since it only manifests when specifically testing cross-resource access between two different legitimate users of the same role.

## Transfer

Audit one real endpoint in TARDOC or Mahall that returns clinic-specific or seller-specific data, and explicitly trace through its authorization logic (or lack thereof) to determine whether it currently performs both the role-level check and the resource-ownership check from this lesson, or only the former. If you find it's missing the resource-ownership check (a realistic possibility if authorization was added incrementally without this specific distinction in mind), describe exactly what a malicious or simply curious authenticated user of your system could currently access that they shouldn't be able to.

## Done when

You've implemented and correctly distinguished 401 (authentication failure) from 403 (authorization failure) in your own working code, you've implemented both role-based and resource-ownership-based authorization as genuinely separate checks, and you've deliberately reproduced the cross-resource-access vulnerability from Learn with two real test users, confirming you understand exactly why role-only authorization is insufficient and how to test for this specific class of bug going forward.
