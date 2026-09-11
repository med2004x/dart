# Lesson 2: Reproduce a Real Bug

## Objective

Take a real, reported bug from an open-source project's issue tracker and turn it into a minimal, deterministic reproduction — and ideally, a failing automated test — the essential first step of any real bug fix, and a genuinely different skill from writing new features.

## Prerequisites

Lesson 1 (repository reading — you need a working mental model of the project before you can reliably reproduce a bug within it), Linux tools Lesson 4 (debuggers — often needed to narrow down a reproduction), algorithms track (general debugging methodology, if a dedicated "systematic debugging" lesson exists earlier in this curriculum).

## Learn

**Why "I can't reproduce it" is one of the most common reasons a real bug report goes unfixed.** A bug report often describes a symptom observed under specific, sometimes not-fully-documented conditions (a particular input, a particular environment, a particular sequence of operations) — without a reliable way to make the bug happen on demand, it's essentially impossible to confirm a fix actually addresses it (you could make an unrelated change, believe it fixed the bug because the original symptom happened not to recur, and be wrong). Reproduction isn't a formality before the "real" work of fixing — it's the foundation that makes verifying any fix possible at all.

**Minimizing a reproduction: finding the smallest case that still exhibits the bug.** A bug initially reported against a large, complex real-world scenario (a specific user's large dataset, a specific complex sequence of application usage) is hard to reason about directly. The discipline of minimizing — repeatedly removing or simplifying elements of the reproduction while confirming the bug still occurs — produces a much smaller, more tractable case that's easier to reason about, easier to debug (Linux tools Lesson 4's tools work better on a smaller, faster-running case), and easier for a reviewer or maintainer to quickly understand and confirm.

**Determinism: making the bug reproduce reliably, not just "sometimes."** Some bugs are inherently non-deterministic in their surface symptom (a race condition, OS Lesson 2) even though their underlying cause is fully deterministic — for these, the goal shifts from "make it fail every single time" (which may be genuinely difficult or impossible for a timing-dependent bug) to "make it fail reliably enough, or under a specific enough condition, to reason about and eventually fix" — sometimes this means finding a way to make the race condition more likely to manifest (e.g. adding an artificial delay at a specific point, or running under a stress-testing tool) rather than eliminating the underlying non-determinism entirely.

**Capturing the reproduction as an automated test: the artifact that actually matters going forward.** A reproduction that only exists as "run these manual steps" is fragile — it depends on someone correctly repeating the exact steps, and provides no ongoing protection against the bug recurring later (a **regression** — the same bug reappearing after being fixed once, often because nothing was left in place to catch it). A reproduction captured as an automated, failing test is what actually gets checked into the codebase, runs on every future change, and prevents the specific bug from silently coming back — this is precisely why Lesson 3's patch process explicitly requires adding regression coverage as part of a complete fix, not the reproduction alone.

## Attempt

1. Choose an open, unresolved bug report from a real project's issue tracker (ideally the same project from Lesson 1, so you already have a working mental model of its structure) — pick one with enough detail that reproduction seems plausible, but that you haven't already seen the fix for.

2. Attempt to reproduce the bug exactly as reported first (following the reporter's described steps/environment as closely as possible), and confirm you can actually observe the reported symptom before doing anything else — if you cannot reproduce it as reported, document specifically what you tried and what happened instead, since "attempted reproduction, could not confirm" is itself a legitimate and useful finding to report back (real open-source maintainers value this over silence).

3. Once reproduced, work to minimize the reproduction: systematically simplify the input/scenario while re-confirming the bug still occurs after each simplification, until you reach a case that's clearly smaller/simpler than the original report but still reliably exhibits the same underlying issue.

4. Write an automated, failing test capturing your minimized reproduction, following the project's existing test conventions (per Lesson 1's investigation of its test suite) — confirm the test fails in a way that clearly demonstrates the bug (not just "test failed" but specifically showing the incorrect actual behavior versus the expected correct behavior).

## Verify

Present your minimized reproduction (showing the progression from the original report's complexity down to your simplified case), and your failing automated test with its actual failure output clearly showing the bug's symptom.

## Failure drill

Take your step 3 minimization process and deliberately over-simplify at one point — remove or simplify something that turns out to actually be necessary for the bug to manifest (a realistic risk during minimization: it's easy to simplify away the specific condition that was actually triggering the issue, ending up with a "reproduction" that no longer reproduces anything). Confirm this over-simplified version no longer exhibits the bug, and explicitly back up one step to the last version that still worked, adding back only the specific element that turned out to be necessary. Explain why this back-and-forth, confirm-after-every-step discipline (rather than removing multiple things at once and checking only at the end) is what actually makes minimization reliable — removing several things simultaneously risks exactly this failure mode, where you can't tell which specific removal broke the reproduction.

## Transfer

If TARDOC, Mahall, or Lead Sourcer has any known, currently-unreproduced or intermittently-occurring bug (a realistic possibility for any real, actively developed system), apply this lesson's methodology to it: attempt reproduction, minimize if successful, and write a failing automated test — using this real, personally-relevant case as a genuine test of the skill, rather than only ever practicing it on someone else's open-source project.

## Done when

You've reproduced a real, previously-unreproduced-by-you bug from an actual open-source project's issue tracker, minimized it to a smaller, clearer case through a careful, step-by-step process, and captured it as a failing automated test following the project's own conventions — plus directly experienced, via the failure drill, why confirming reproduction after every single simplification step (not just at the end) is what makes the minimization process reliable rather than accidentally destructive.
