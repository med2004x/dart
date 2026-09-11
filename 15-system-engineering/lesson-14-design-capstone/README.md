# Lesson 14: Design Capstone

## Objective

Produce a complete system design document for a substantial, realistic system — integrating requirements/capacity planning, module boundaries, consistency decisions, reliability targets, observability, deployment strategy, and documented tradeoffs — the system-engineering track's equivalent of the integration capstones throughout this curriculum, but focused on design artifacts rather than only running code.

## Prerequisites

Lessons 1-13, completed. This lesson is explicitly about producing a design document as a deliverable, not (only) working code — the discipline of writing a design down clearly enough that someone else could evaluate or implement it is itself the skill being tested here, distinct from (though informed by) the hands-on implementation work in earlier lessons.

## Learn

There is no new material. The specific gap this capstone tests: can you take everything reasoned through in isolation across Lessons 1-13 (request boundaries, capacity numbers, module structure, consistency choices, reliability targets, observability plan, deployment strategy) and present it as one coherent design that a critical reader — a technical collaborator, a future hire, or your own future self returning to the project after months away — could actually evaluate, critique, or implement from, without needing you present to fill in unstated assumptions.

## Attempt

Choose a real, substantial system to design — a genuine extension of TARDOC, Mahall, or Lead Sourcer that doesn't yet exist but is plausible given your actual business context (e.g. "a multi-clinic reporting dashboard for TARDOC," "a seller analytics feature for Mahall," or "a lead-quality feedback loop for Lead Sourcer") — and produce a complete design document covering:

1. **Requirements and capacity** (Lesson 2): concrete numbers for expected load, latency targets (with percentiles, not just averages), and data volume projections, derived from real or realistically estimated business context, not arbitrary round numbers.

2. **Resource/module boundaries** (Lesson 3, and API engineering Lesson 2): a clear module decomposition with a justified, acyclic dependency graph, and — if the system exposes an API — a resource model per API engineering's framework.

3. **Request boundary decisions** (Lesson 1): for the system's core operations, explicit reasoning about what's synchronous versus deferred to background/async processing.

4. **Consistency model** (Lesson 7): explicit classification of which data needs strong consistency versus where eventual consistency is an acceptable, deliberate tradeoff, with justification.

5. **Reliability targets and mechanisms** (Lesson 8): a stated availability target with business justification, and the specific mechanisms (circuit breakers, graceful degradation, retry/timeout policy per Lesson 5) that would actually achieve it.

6. **Observability plan** (Lesson 9): what logs, metrics, and traces the system would need to actually diagnose problems after deployment, not just "we'll add logging later."

7. **Deployment strategy** (Lesson 11): which pattern (rolling, blue-green, canary) fits this system's actual constraints, and why.

8. **At least 2 ADRs** (Lesson 12) for genuine design decisions within this system, with real alternatives considered and honest tradeoffs stated.

9. **A migration/rollout plan** (Lesson 13) if the system needs to integrate with or replace any existing functionality, using the strangler-fig/incremental approach rather than a big-bang cutover.

## Verify

The design document itself is the primary deliverable — it should be complete enough that a competent engineer unfamiliar with your specific reasoning could read it and understand not just *what* you're proposing, but *why*, including the tradeoffs you're knowingly accepting. As a concrete verification step, have someone else (or, working solo, revisit it after a meaningful gap with fresh eyes, echoing Lesson 12's critical-reader exercise) attempt to identify the single weakest part of the design — the decision they'd push back on hardest — and record what that was and whether you find the critique persuasive on reflection.

## Failure drill

Take one section of your design document — the consistency model (item 4) is a good candidate, given how easy it is to get subtly wrong — and deliberately construct a scenario where your stated design choice would produce a real, user-visible problem, similar to Lesson 7's failure drill but now applied to your own, larger system design rather than a small isolated exercise. Walk through the scenario in writing: what would a user actually observe, and does your design document's stated tradeoffs (item 4) actually account for this specific case, or does working through this scenario reveal a gap in your original reasoning that needs to be revised? Update the document if you find a genuine gap, rather than leaving an identified flaw undocumented.

## Transfer

Compare your design document's reliability-target section (item 5) against what you know or can reasonably infer about TARDOC or Mahall's actual current reliability characteristics, given their real infrastructure (a single Contabo VPS, per your project history) — state explicitly whether your proposed new system's target is realistic given that it would likely run on similar infrastructure, or whether achieving your stated target would require infrastructure changes beyond what currently exists, and if so, name them specifically rather than leaving the gap implicit.

## Done when

You've produced a complete design document covering all 9 required elements for a real, substantial system, grounded in actual or realistically estimated numbers rather than placeholders, you've subjected at least one section to genuine critical scrutiny (via the failure drill or a second reader) and revised it if a real gap was found, and the document is specific and complete enough that someone unfamiliar with your reasoning could evaluate it on its merits, not just take your conclusions on faith.
