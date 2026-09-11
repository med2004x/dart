# Lesson 3: Patch and Review

## Objective

Implement the smallest correct fix for the bug reproduced in Lesson 2, add proper regression coverage, and review your own diff with the same critical rigor system-engineering Lesson 15's architecture review applied to design documents — now applied to a concrete code change.

## Prerequisites

Lesson 2 (bug reproduction — this lesson fixes exactly the bug reproduced there), system-engineering Lesson 15 (architecture review — the critical-self-review discipline this lesson applies at the code-diff level).

## Learn

**Why "the smallest correct patch" is a deliberate constraint, not just an efficiency preference.** A patch that fixes the reported bug *and* refactors surrounding code *and* adds an unrelated improvement is harder to review (a reviewer has to evaluate multiple, entangled changes at once), harder to reason about if something goes wrong later (which part of the combined change caused a regression — directly echoing Linux tools Lesson 3's bisect-ability concern about mixed commits), and more likely to be rejected or delayed by a maintainer who doesn't want to accept the unrelated scope creep alongside the actual fix they asked for. The smallest correct patch is specifically the change that fixes the bug and nothing else — resisting the temptation to "improve things while I'm in here" is a real, disciplined skill.

**Regression coverage: turning Lesson 2's reproduction test from "currently failing" to "will alert us if this ever breaks again."** Lesson 2 produced a failing test demonstrating the bug. The patch itself should make that exact test pass — and the test then remains in the codebase permanently, running on every future change, specifically to catch the bug if it's ever accidentally reintroduced (a real, common occurrence — a later, unrelated refactor can reintroduce a previously-fixed bug if nothing is actively checking for it).

**Reviewing your own diff as a maintainer would — a genuinely different, more critical stance than "does this look right to me."** A maintainer reviewing an external contribution asks specific, somewhat adversarial questions: does this actually fix the root cause, or just paper over the specific symptom in the specific reported case (a real, common distinction — a fix that only handles the exact reported scenario, without addressing the underlying issue, will likely resurface as a "similar but not identical" bug report later)? Does the fix introduce any new edge cases or regressions of its own? Is the change scoped appropriately (per the previous point)? Does it follow the project's existing conventions (Lesson 1's investigation)? Deliberately adopting this more skeptical stance toward your own change — rather than the more forgiving "I wrote this, it probably works" instinct — is the actual skill this lesson builds, directly extending system-engineering Lesson 15's critical-review discipline to a concrete code artifact.

**Root cause versus symptom: a distinction worth making explicit in your own review.** If your fix works by adding a special-case check for the exact condition in the bug report, without addressing why that condition produced incorrect behavior in the first place, you've likely fixed the symptom, not the root cause — and a maintainer's review (or your own self-review, done with appropriate rigor) should specifically probe for this, since a symptom-only fix is a real, common way a bug appears fixed in review but resurfaces in a slightly different form later.

## Attempt

1. Implement a fix for Lesson 2's reproduced bug, and confirm your Lesson 2 failing test now passes — the most basic verification that the fix addresses what was reproduced.

2. Explicitly evaluate whether your fix addresses the root cause or just the specific reported symptom (per Learn's final point) — trace back from the failing test to understand *why* the original code produced incorrect behavior, and confirm your fix addresses that underlying reason, not just the specific input case your test happens to check. If you find your first attempt was symptom-only, revise it to address the actual root cause.

3. Review your own diff using the maintainer's-critical-stance discipline from Learn: is the change minimal and focused (no unrelated scope creep), does it follow the project's existing code conventions (Lesson 1's investigation), and can you think of any related edge case your fix might not correctly handle (test at least one such edge case explicitly, beyond just the original reported scenario).

4. Write a clear commit message / patch description explaining the bug, its root cause, and the fix — written for a reviewer who has *not* just spent hours reproducing and understanding the bug the way you have, meaning it needs to concisely convey context a fresh reader would actually need.

## Verify

Present your diff (the actual code change), your passing regression test (previously failing per Lesson 2, now passing), your step 2 root-cause reasoning, and your step 3 edge-case test confirming your fix handles at least one related scenario beyond the original report.

## Failure drill

Take your first-draft fix (before step 2's root-cause revision, if you had one, or deliberately construct a symptom-only "fix" now for this exercise if your first attempt happened to get the root cause right immediately) and construct a *related but distinct* bug scenario — one that a root-cause fix would handle correctly but a symptom-only fix would not. Test your symptom-only version against this related scenario and confirm it fails, then test your proper, root-cause-addressing fix against the same scenario and confirm it passes. Explain, using this concrete before/after comparison, why root-cause analysis matters beyond just "being more thorough" — a symptom-only fix can pass code review and the original reproduction test while still harboring the actual underlying defect, which will very plausibly resurface as a fresh, seemingly-unrelated bug report later, costing more total effort than getting it right the first time would have.

## Transfer

Apply this lesson's patch-and-self-review discipline to a real bug you've fixed (or are currently fixing) in TARDOC, Mahall, or Lead Sourcer — review your actual diff using the maintainer's-critical-stance checklist from Learn, and honestly assess whether your original fix (if already merged) addressed the root cause or just the reported symptom, using this lesson's distinction as the specific test.

## Done when

You've implemented a fix for Lesson 2's reproduced bug that passes the regression test, explicitly verified (and if necessary revised toward) a root-cause fix rather than a symptom-only patch, reviewed your own diff with genuine maintainer-level critical scrutiny including testing at least one related edge case, and demonstrated — via the failure drill's constructed related scenario — the concrete difference in outcome between a symptom-only fix and a root-cause fix.
