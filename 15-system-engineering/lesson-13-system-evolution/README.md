# Lesson 13: System Evolution

## Objective

Plan and execute an incremental migration of a real system component — changing something significant (a schema, a core dependency, an architectural pattern) without a disruptive rewrite or a risky big-bang cutover, using the strangler fig pattern and backward-compatible intermediate states.

## Prerequisites

System-engineering Lesson 11 (safe deployments — migration is deployment applied to a structural change, not just a code release), API engineering Lesson 13 (versioning — a specific instance of the general "change without breaking existing consumers" problem this lesson generalizes).

## Learn

**Why "rewrite it from scratch" is usually the wrong instinct, even when the current system has real problems.** A full rewrite defers all value delivery until the rewrite is complete (which reliably takes longer than initially estimated — a well-documented pattern across the industry, not a personal failing), during which the old system continues accumulating the very problems motivating the rewrite, and the eventual cutover is a single, high-risk, big-bang event exactly like the "deploy and hope" pattern system-engineering Lesson 11 warned against, just at a much larger scale. Incremental migration — changing the system in small, individually low-risk, individually valuable steps — avoids all of these problems, at the cost of being conceptually harder to plan than "just rewrite it," since it requires the old and new to coexist correctly during the transition.

**The strangler fig pattern, the standard approach for incremental system replacement.** Named after the strangler fig plant, which grows around an existing tree, gradually taking over its structural role until the original tree can be removed with the fig now fully self-supporting. Applied to software: build the new component alongside the old one, route an increasing share of traffic/functionality to the new component incrementally (directly reusing system-engineering Lesson 11's canary/rolling patterns, now applied to a structural migration rather than a routine deployment), and only remove the old component once the new one has fully taken over and been validated under real production load — at every intermediate step, the system remains fully functional, just with a shifting internal split between old and new.

**Backward-compatible intermediate states: the discipline that makes incremental migration actually possible.** Each individual step in a migration needs to leave the system in a fully working state — this often means temporarily supporting *both* the old and new forms of something simultaneously (e.g. a database migration that adds a new column, backfills it, and only removes the old column in a later, separate step, rather than one migration doing all three at once) — API engineering Lesson 13's non-breaking-change discipline, applied internally to your own system's evolution, not just to external API consumers.

**Feature flags: the mechanism that typically controls the traffic split during a strangler-fig migration.** A runtime-configurable flag that determines whether a given request/operation uses the old or new code path, letting you shift the split (0% new → 5% → 50% → 100%) without a code deployment at each step, and — critically — letting you instantly revert to the old path (flip the flag back) if the new path shows a problem, without needing system-engineering Lesson 11's full redeploy-based rollback for what's fundamentally a routing decision, not a code change.

## Attempt

1. Identify a real component in TARDOC or Mahall that could benefit from an incremental migration (e.g. a database schema change, a core library swap, or an internal API's shape change) — describe the current state and the desired end state.

2. Design the migration as a sequence of individually safe, backward-compatible steps (not "swap it in one go") — for a schema-change example specifically: add the new column (nullable, so existing writes aren't affected), backfill existing data, update application code to write to *both* old and new columns simultaneously (a dual-write period), verify data consistency between old and new, switch reads to the new column, stop writing to the old column, and only then drop the old column — write out this full sequence explicitly for your real chosen component.

3. Implement a feature flag controlling at least one step of your migration (e.g. whether a specific operation reads from the old or new data representation), and demonstrate flipping it both directions — confirm the system behaves correctly with the flag in either state, and confirm flipping it back (simulating an emergency rollback) is fast and doesn't require a code deployment.

4. Execute (or simulate, if a full real migration isn't practical for this exercise) at least the first 2-3 steps of your step 2 sequence against a real or realistic test dataset, confirming the system remains fully functional at each intermediate step — not just at the very end.

## Verify

Present your full step 2 migration sequence, your step 3 feature flag implementation with demonstrated bidirectional flipping, and evidence (test results, or a description of your executed/simulated steps) that the system remained functional at each intermediate stage of step 4, not just before and after the full migration.

## Failure drill

Deliberately skip the "dual-write, verify consistency" step from your migration sequence — switch reads directly to the new representation immediately after adding it, without first confirming the backfill actually produced correct, consistent data. Construct a test case where the backfill has a subtle bug (e.g. a transformation that's correct for most records but wrong for some edge case), and confirm that skipping the verification step means this bug goes undetected until it's already affecting real reads — directly demonstrating why the verification step, though it adds time to the migration, is what actually catches this class of problem before it causes user-visible incorrect behavior, rather than the migration merely *appearing* to succeed while quietly serving wrong data.

## Transfer

If TARDOC's repo restructuring (mentioned in your project history as a completed piece of work) or Mahall's Store Studio architectural work (also mentioned, involving two parallel rendering systems coexisting) involved anything resembling this lesson's incremental-migration pattern, describe, using what you actually know about how that work was approached, whether it followed a strangler-fig-style incremental approach or a more direct replacement — and if you were to redo a similar structural change today, what specific step from this lesson's migration sequence (dual-write, verification, feature-flag-controlled cutover) you'd most want to make sure was explicitly included, based on what you now understand about why each step exists.

## Done when

You've designed a complete, genuinely incremental migration sequence for a real system component with each step individually leaving the system in a working state, you've implemented and tested a feature flag enabling fast, code-deployment-free rollback during the migration, and you've demonstrated — via the failure drill — a concrete case where skipping the verification step would have let a real data-correctness bug through undetected, understanding specifically why that step exists rather than treating it as unnecessary caution.
