# Lesson 3: Authentication and Authorization (Deep Dive)

## Objective

Design a complete authorization matrix for a real system with multiple roles and resource types, and rigorously test privilege boundaries — extending API engineering Lesson 9's endpoint-level treatment into a systematic, whole-system authorization model.

## Prerequisites

API engineering Lesson 9 (authentication/authorization basics — 401 vs 403, role vs. resource-level checks), system-engineering Lesson 10 (security architecture — trust boundaries and least privilege, applied here to the specific case of authorization design).

## Learn

**Identity, authentication, sessions, and authorization: four genuinely distinct concepts, often conflated.** Identity is *who* an entity claims to be (a user ID, an account). Authentication is the *process of verifying* that claim (a password check, a token validation). A session (or its stateless equivalent, a token) is *how a verified identity is remembered* across multiple requests, so authentication doesn't need to be repeated on every single interaction. Authorization (API engineering Lesson 9's core topic) is what a *specific, now-known* identity is permitted to *do*. Conflating these — e.g. treating "has a valid session" as equivalent to "is authorized for this specific action" — is exactly the category of mistake that produces the broken-access-control vulnerabilities API engineering Lesson 9's failure drill demonstrated.

**An authorization matrix: making every role-resource-action combination explicit, rather than implicit in scattered code.** For a system with roles (e.g. `admin`, `clinic-staff`, `patient` — adjust to your actual domain) and resource types (e.g. `clinic-records`, `billing-data`, `user-accounts`) and actions (`read`, `create`, `update`, `delete`), a matrix explicitly states, for every combination, whether it's permitted — making gaps and inconsistencies visible in one place, rather than requiring someone to trace through every individual endpoint's ad hoc authorization logic to reconstruct what the *intended* policy actually is. This is directly analogous to API engineering Lesson 3's point about writing the OpenAPI contract before implementation: writing the authorization matrix before or alongside implementation surfaces ambiguity and gaps early, rather than discovering them via a security incident.

**Capability-based versus role-based thinking, briefly.** Role-based access control (RBAC, API engineering Lesson 9's model) grants permissions based on a role label. A capability-based model instead grants specific, fine-grained permissions directly (e.g. "can read this specific clinic's billing data") independent of any named role — more flexible and precise, but with more bookkeeping overhead. Most real systems use a hybrid: RBAC for broad categories, with resource-ownership checks (API engineering Lesson 9's second check) providing the fine-grained, capability-like precision on top.

**Testing privilege boundaries deliberately, not just testing that authorized access works.** A genuinely thorough authorization test suite specifically tries to access resources/actions a given role/identity should *not* be able to reach — for every cell in your authorization matrix marked "not permitted," a test confirming that access attempt is actually rejected. This is a different, and often neglected, testing discipline compared to only testing the happy path (authorized users can do what they're supposed to) — API engineering Lesson 9's failure drill demonstrated exactly why: the vulnerability there was only detectable by testing cross-resource access specifically, not by testing that legitimate, same-resource access worked correctly.

## Attempt

1. For a real or plausible system (TARDOC's roles might include `admin`, `clinic-staff`, and possibly a billing-specific role; adjust to your actual domain), enumerate at least 3 distinct roles, 3 distinct resource types, and the standard CRUD actions, and produce a complete authorization matrix — every role/resource/action combination explicitly marked permitted or not, with no gaps left implicit.

2. Implement the authorization checks matching your matrix for at least one resource type across all 3 roles, using a combination of role-based and resource-ownership checks per API engineering Lesson 9's pattern.

3. Write a systematic test suite covering *every cell* of your matrix for the implemented resource type — not just "admin can do everything" and "patient can read their own records" (the happy paths), but explicitly testing every "not permitted" cell as well, confirming each is actually rejected with the correct status code (401 vs 403, per API engineering Lesson 9).

4. Deliberately introduce a gap between your matrix (the intended policy) and your implementation (e.g. forget to implement one specific role/resource/action check that your matrix says should be denied), and confirm your systematic test suite from step 3 actually catches this specific gap — a test that fails precisely because the implementation doesn't match the documented, intended policy.

## Verify

Present your complete authorization matrix, your test suite's actual pass/fail results covering every cell for at least one resource type, and confirm step 4's deliberately introduced gap was actually caught by your test suite (show the specific failing test and what it revealed).

## Failure drill

Take your authorization matrix and identify any cell where the *documented* policy and your *initial intuition* about what "should" be allowed differ — a genuinely useful signal that either your matrix has an error, or your intuition was based on an incomplete understanding of the actual intended policy. Investigate this specific discrepancy: which one is actually correct given the real business requirements, and what would happen if the implementation matched the wrong one? Explain why writing the matrix explicitly, as a document separate from the code, is specifically what surfaced this kind of discrepancy — if you'd gone straight to implementation without the explicit matrix step, this exact kind of "I thought it should work this way, but actually the policy requires something different" gap could easily have gone unnoticed until a real incident revealed it.

## Transfer

Audit TARDOC's or Mahall's actual current authorization logic (whatever roles and checks currently exist) against this lesson's matrix-first methodology: attempt to reconstruct what authorization matrix the *current, actual code* implicitly represents, by reading through the existing checks — and compare that reconstructed matrix against what you believe the *intended* policy actually is. Report any discrepancies you find, which represent exactly the kind of implicit, undocumented gap this lesson's explicit-matrix approach is designed to prevent going forward.

## Done when

You've produced a complete, explicit authorization matrix for a real system with no implicit gaps, you've implemented and systematically tested every cell of that matrix (not just the happy-path cells), you've demonstrated your test suite actually catches a deliberately introduced policy-implementation mismatch, and you've audited a real, existing system's authorization logic against this methodology and reported any real discrepancies found.
