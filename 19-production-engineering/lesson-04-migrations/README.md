# Lesson 4: Safe Migrations

## Objective

Design and execute an expand/contract schema migration that supports a mixed-version production deployment — old and new application code running simultaneously against the same database during a rollout — without either version breaking, directly applying system-engineering Lesson 13's incremental-migration principles to the specific, high-stakes case of database schema changes.

## Prerequisites

System-engineering Lesson 13 (system evolution/strangler fig — this lesson is that lesson's backward-compatible-intermediate-states principle applied specifically to schema migrations), database-internals Lesson 1 (page/record storage — useful background for understanding what a schema migration actually does at the storage level), system-engineering Lesson 11 (safe deployments — rolling/canary deployments mean old and new code versions coexist briefly, which is exactly the constraint this lesson's migrations must support).

## Learn

**Why a naive schema migration breaks rolling deployments.** System-engineering Lesson 11 established that rolling/canary deployments mean old and new application code run simultaneously against the same infrastructure for some period. If a migration renames a column in one atomic step, the moment that migration runs, *every* still-running instance of the old application code (expecting the old column name) breaks immediately — even though the deployment of new code hasn't finished rolling out yet. This is a genuine, common production incident pattern: the migration itself, not the code deployment, is what caused the outage, precisely because it wasn't designed to tolerate the old code's continued existence during the transition window.

**Expand/contract, the standard pattern for migrations compatible with rolling deployment.** *Expand*: add the new schema element (a new column, a new table) without removing or renaming anything old — this step is purely additive and, by construction, cannot break code that doesn't yet know about the new element. *Migrate*: deploy application code that writes to *both* old and new representations simultaneously (a dual-write period, matching system-engineering Lesson 13's exact pattern), and backfill existing data into the new representation. *Verify*: confirm data consistency between old and new representations before proceeding — this is the step system-engineering Lesson 13's failure drill demonstrated the real cost of skipping. *Switch reads*: deploy code that reads from the new representation. *Contract*: only once every instance is confirmed running new code (no old instances remain that might still read/write the old representation) do you finally remove the old schema element — the truly destructive step, deliberately deferred until it's provably safe.

**Why "verify rollback works" needs to be tested, not assumed.** A migration needs a genuine rollback path — if the new code version has a critical bug discovered after deployment, you need to be able to roll back to the *old* code version while the database is in whatever intermediate expand/contract state it's currently in. This means testing, explicitly, that the *old* application code still functions correctly against the post-expand (but pre-contract) schema state — not just that the new code works, but that reverting doesn't itself become a second, unplanned outage because the old code can no longer function against the modified schema.

**Simulating mixed-version deployment for real, not just trusting the design on paper.** Because this class of bug (old code breaking against a schema change mid-migration) is specifically about the *interaction* between two code versions and one shared database state, the only way to be genuinely confident the migration is safe is to actually run both old and new code simultaneously against the migrated schema in a test environment and confirm both function correctly — a design that looks correct on paper can still have a subtle gap only actual mixed-version testing would reveal.

## Attempt

1. For a real or realistic schema change (e.g. splitting a single `name` column into `first_name`/`last_name`, or changing a column's data representation), design the full expand/contract sequence explicitly: what each step adds/changes, and crucially, confirm explicitly that the *old* application code continues to function correctly at every single intermediate step, not just the final state.

2. Implement the "expand" step (add the new column(s), nullable, no changes to existing behavior) and confirm your *unmodified, old* application code runs against the migrated schema with zero changes required — direct proof the expand step alone is non-breaking.

3. Implement the dual-write "migrate" step (new application code writing to both old and new representations) and the backfill (populating the new representation for existing rows), then implement the verification step explicitly: write a script comparing old and new representations for every row and confirming consistency, following system-engineering Lesson 13's exact discipline.

4. Simulate mixed-version deployment directly: run one instance of your *old* application code and one instance of your *new* (dual-write) application code simultaneously against the same, post-expand database, with both instances actively performing operations concurrently. Confirm both function correctly throughout — the old code oblivious to the new column, the new code correctly maintaining both representations — with no errors or data inconsistency introduced by either.

## Verify

Present your full expand/contract sequence design, your step 2 confirmation that old code runs unmodified against the expanded schema, your step 3 consistency-verification script's actual output confirming old/new representation agreement, and your step 4 mixed-version simulation's results confirming both code versions functioned correctly simultaneously.

## Failure drill

Skip the dual-write step and instead migrate directly from "old representation only" to "new representation only, with a one-time backfill but no ongoing dual-write" — then simulate a realistic rolling-deployment window where old code (still writing only to the old representation, unaware of the migration) continues running for some period after the backfill completed. Confirm that any write performed by the old code during this window is *not* reflected in the new representation (since old code never learned to dual-write), creating a real, silent data inconsistency between the two representations for exactly the rows old code touched during that window. Explain why this demonstrates precisely why the dual-write period isn't optional scaffolding to skip for a "simpler" migration — without it, any write from not-yet-updated old code during the rollout window is permanently invisible to the new representation, a genuine, silent data-loss bug that a migration lacking dual-write cannot avoid whenever any deployment takes longer than an instant to complete everywhere.

## Transfer

If TARDOC's repo restructuring or any of Mahall's schema evolution (per your project history) involved changing a table structure while the system remained live, describe, using this lesson's expand/contract framework, whether that migration was actually performed with genuine backward compatibility for old code during the transition, or whether it required (or got lucky with) a brief downtime window instead — and if you were to redo a similar change today with zero-downtime rolling deployment as a requirement, what the expand/contract sequence would need to look like.

## Done when

You've designed and executed a complete expand/contract migration sequence for a real schema change, confirmed old code runs correctly at every intermediate step (not just the final state), implemented and run genuine consistency verification between old and new representations, and directly simulated mixed-version deployment with both code versions running concurrently against the migrated schema without error — plus demonstrated, via the failure drill, the specific silent data-loss risk of skipping the dual-write step.
