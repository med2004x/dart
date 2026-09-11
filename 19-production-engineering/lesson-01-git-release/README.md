# Lesson 1: Release Engineering

## Objective

Build a real, reproducible release pipeline that produces versioned, traceable artifacts — and a genuine rollback capability — turning "deploy" from an ad hoc manual process into a repeatable, auditable one.

## Prerequisites

Linux tools Lesson 3 (Git as an engineering tool — releases are fundamentally about a specific, identifiable point in Git history), system-engineering Lesson 11 (safe deployments — this lesson covers the artifact-production side that lesson's deployment patterns consume).

## Learn

**Reproducible builds: why "it built successfully" isn't the same as "it built the same thing every time."** A build is reproducible if building from the same source at the same commit produces a bit-for-bit (or at least functionally) identical artifact every time, regardless of when or where it's built. Non-reproducibility usually stems from unpinned dependencies (security-engineering Lesson 6's exact concern, now applied to build-time rather than runtime risk), embedded timestamps, or non-deterministic build steps — and it matters because a non-reproducible build makes it genuinely hard to verify "the artifact currently running in production actually corresponds to this specific, reviewed source commit," a real auditability and debugging concern, not just a theoretical purity goal.

**Version pinning and build metadata: making every artifact traceable back to its exact origin.** A release artifact should embed enough metadata (the exact Git commit hash, the build timestamp, the version number) that, given a running artifact, you can determine precisely what source code produced it — this directly enables system-engineering Lesson 9's observability discipline (correlating a production issue back to a specific code version) and Lesson 12's ADR practice (knowing which version a given architectural decision actually shipped in).

**A rollback artifact: the previous release, kept ready, not rebuilt on demand.** System-engineering Lesson 11 established that fast rollback requires the previous known-good state to be immediately available, not reconstructed under pressure during an incident. This means your release pipeline should retain previous release artifacts (not just source code — the actual built, deployable artifact) for some retention window, so a rollback is "deploy this already-built, already-verified artifact" rather than "rebuild from an old commit and hope the build environment hasn't changed in the meantime" — the latter reintroduces exactly the reproducibility risk this lesson's first point warned about, at the worst possible time (during an active incident).

**Why manual, ad hoc deployment is a real risk, not just an inconvenience.** A deployment process that involves manually running commands, copying files, or remembering a specific sequence of steps is: error-prone (a step forgotten or done out of order under time pressure), non-auditable (no record of exactly what was done, by whom, when), and slow to execute correctly during an incident when speed matters most. A scripted, automated pipeline eliminates all three risks by making the process itself the documentation, executed identically every time regardless of who triggers it or how much pressure they're under.

## Attempt

1. For a real project (TARDOC, Mahall, or a representative subset), build a release pipeline script (using your CI system, e.g. GitHub Actions, or a local script if CI isn't set up yet) that: checks out a specific Git commit, builds the artifact (a Go binary, a Docker image, or equivalent), and embeds the commit hash and build timestamp into the artifact itself (e.g. via a version endpoint the running application exposes, or embedded build-time constants).

2. Confirm reproducibility directly: build the same commit twice (in two separate pipeline runs, ideally with some time between them) and compare the resulting artifacts — if using Docker images, compare layer hashes; if a Go binary, compare via checksums after normalizing any legitimately non-deterministic elements like embedded build timestamps (which you'd expect to differ) versus everything else (which shouldn't).

3. Implement artifact retention: configure your pipeline to keep at least the current and previous N release artifacts accessible and ready to deploy (not just source code tags), and demonstrate a rollback by actually deploying a previous, retained artifact — confirm the deployed artifact's embedded version metadata (from step 1) correctly identifies it as the older, retained build, not a freshly rebuilt one.

4. Query your running deployed artifact's version endpoint (or equivalent) and confirm it correctly reports the exact commit hash and build timestamp matching what you know was actually deployed — a real, working traceability chain from running production code back to its exact source origin.

## Verify

Show your pipeline's actual configuration/script, your step 2 reproducibility comparison (with any legitimate, expected differences explicitly noted and explained), and your step 3-4 rollback demonstration with the version endpoint correctly confirming which specific build is actually running after the rollback.

## Failure drill

Deliberately introduce a non-reproducibility source into your build (e.g. a dependency version left unpinned, allowing it to resolve to whatever's "latest" at build time rather than a fixed version) and rebuild the same commit twice with some time gap, ideally after the unpinned dependency has had an opportunity to actually change upstream (or simulate this by deliberately changing what "latest" would resolve to between builds). Confirm the two builds now genuinely differ in a way not explained by expected, legitimate variation (like build timestamps) — directly reproducing security-engineering Lesson 6's "dependency resolved differently underneath me" finding, but now observed as a build-reproducibility problem specifically. Explain why this makes your earlier traceability claim (step 4) weaker than it appeared: if the same commit can produce genuinely different artifacts depending on when it's built, "this artifact came from commit X" is a necessary but not sufficient statement — you also need "and this artifact was built with exactly these pinned dependency versions" for the traceability chain to be complete and trustworthy.

## Transfer

If TARDOC or Mahall's current deployment process involves any manual, undocumented steps (a realistic possibility given the single-developer, move-fast context described in your project history), describe specifically what those steps are, and what the first, most valuable piece of this lesson's pipeline (build automation, version embedding, or artifact retention) would be to implement first, given your actual current risk profile and time constraints.

## Done when

You've built a real, working release pipeline that produces versioned, traceable artifacts, you've confirmed build reproducibility (or found and understood a real gap in it via the failure drill), and you've demonstrated an actual rollback to a previously-built, retained artifact with the running system's version metadata correctly confirming exactly what's deployed.
