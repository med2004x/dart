# Lesson 10: Security at the System Level

## Objective

Apply security thinking at the architecture level — trust boundaries, defense in depth, least privilege — as a design discipline distinct from (and prerequisite to) the specific vulnerability classes the security-engineering track covers in depth.

## Prerequisites

API engineering Lesson 9 (authentication/authorization — this lesson generalizes those endpoint-level concerns to whole-system architecture), networking Lesson 5 (TLS — one concrete instance of a trust boundary this lesson discusses more generally).

## Learn

**Trust boundaries: identifying exactly where "untrusted input becomes trusted" in your system.** Every system has boundaries where data crosses from an untrusted context (a public API request, user-uploaded content, a third-party webhook) into a trusted internal context (your database, your business logic assuming well-formed data). Security failures very often occur precisely at these boundaries — not because the internal logic is flawed, but because something crossed the boundary without adequate validation or sanitization, carrying an assumption of trust it hadn't actually earned. Explicitly mapping where your system's trust boundaries are — not assuming they're obvious — is the first step to reasoning about where validation, authentication, and sanitization actually need to happen.

**Defense in depth: not relying on any single control to be perfect.** A well-designed system doesn't assume its input validation (API engineering Lesson 5) will catch every malicious input, or that its authentication (API engineering Lesson 9) will never be bypassed — it layers multiple independent controls, so a failure in any single layer doesn't immediately result in a full compromise. Concretely: even if application-level authorization has a bug, a properly configured database user with minimally scoped permissions limits what an attacker who found that bug could actually do; even if a web server has a vulnerability, network-level firewall rules limiting what can reach it in the first place provide an additional, independent layer.

**Least privilege: every component should have exactly the access it needs, no more.** A backend service's database credentials shouldn't have `DROP TABLE` permission if the service only ever needs `SELECT`/`INSERT`/`UPDATE` on specific tables — not because the service's own code is expected to misbehave, but because if that service is ever compromised (via any vulnerability, known or unknown), the *attacker* inherits exactly whatever privileges the compromised component had — minimizing that blast radius in advance, before any specific vulnerability is even known, is the whole point of least privilege as a proactive design discipline rather than a reactive patch.

**Why this is an architecture-level discipline, not just "write secure code."** The security-engineering track (Lessons 17.1-17.6) covers specific vulnerability classes (memory safety, injection, authentication flaws) in depth — genuinely important, code-level knowledge. This lesson is about the layer above that: even with perfectly secure code at every individual layer (an unrealistic assumption to design around, which is exactly the point), a system with no trust boundary awareness, no defense in depth, and components running with excessive privilege is architecturally fragile — one single flaw anywhere becomes a full compromise, rather than being contained by the surrounding architecture.

## Attempt

1. For a real system (TARDOC or Mahall), draw an explicit trust-boundary diagram: every point where data enters from outside your control (public API requests, webhook payloads from external services, user uploads) and every point where it crosses into a "trusted" internal context (database writes, internal service calls). Mark, for each boundary, what validation/sanitization currently happens there.

2. Audit your database credentials/connection configuration for at least one real service against the least-privilege principle: what permissions does the actual database user your application connects with actually have, and are any of them broader than what the application's real, current functionality requires (e.g. does it have `DROP`/`ALTER` permissions it never actually uses)? If you find excess permissions, describe what a correctly-scoped set would look like.

3. Design (and, where practical, implement) one concrete defense-in-depth improvement for a real trust boundary from step 1 — e.g., if you currently rely solely on application-level input validation for a specific field, add a database-level `CHECK` constraint enforcing the same rule at a second, independent layer, so a bug in the application-level validation wouldn't alone be sufficient to insert invalid data.

4. Simulate a single-layer failure and confirm your defense-in-depth improvement from step 3 actually provides the additional protection intended: deliberately bypass or disable your application-level validation for a test (simulating a bug in that layer) and confirm the database-level constraint from step 3 still catches the invalid data, preventing it from actually being persisted despite the application-layer control having failed.

## Verify

Present your step 1 trust-boundary diagram, your step 2 least-privilege audit findings (including any excess permissions identified), and show the actual test results from step 4 — the application-level validation deliberately bypassed, and the database-level constraint still correctly rejecting the invalid data, demonstrating the defense-in-depth layer functioning independently.

## Failure drill

Take your step 2 least-privilege audit and, if you found excess permissions, construct a scenario demonstrating the actual risk: with the current, overly broad permissions, show (in a safe test environment, never against real data) that the application's database credentials *could* execute a destructive operation (like dropping a table) that the application's own code never intentionally does but that the credentials technically permit — then apply a correctly-scoped, minimal-privilege credential set and confirm the same destructive operation now fails with a permissions error at the database level, even if somehow triggered (e.g. via a hypothetical SQL injection or a bug). Explain why this demonstrates least privilege's real value: it doesn't prevent the hypothetical bug or injection from occurring, but it bounds the damage that specific class of failure could cause, independent of whatever caused the failure in the first place.

## Transfer

If TARDOC's webhook-receiving endpoints (if any) or user-facing upload features represent a trust boundary you haven't explicitly reasoned about using this lesson's framework, describe what's currently validated at that boundary, and identify at least one additional defense-in-depth layer (following this lesson's database-constraint pattern, or an analogous independent control) that would reduce the impact if the primary, application-level control at that boundary ever had an undiscovered flaw.

## Done when

You've produced a real trust-boundary diagram for an actual system and identified what's validated at each boundary, you've audited real database credentials against least privilege and identified any actual excess permissions, and you've implemented and directly tested a defense-in-depth control that continues protecting the system even when you deliberately simulate the primary control failing — demonstrating the layered-defense principle with real, working code rather than as an abstract claim.
