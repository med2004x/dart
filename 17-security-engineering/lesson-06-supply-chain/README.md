# Lesson 6: Supply Chain Security

## Objective

Treat your dependencies and build pipeline as part of your actual attack surface — pin dependencies to specific, verified versions, generate a real dependency inventory, and review a CI/build workflow for trust-boundary gaps, applying system-engineering Lesson 10's trust-boundary framework to code you didn't write yourself.

## Prerequisites

System-engineering Lesson 10 (trust boundaries and least privilege — this lesson applies both to a boundary many developers overlook: the code pulled in from external dependencies and the systems that build and deploy your software).

## Learn

**Why "it's just a dependency" undersells the actual risk.** Every third-party package you depend on runs with the same privileges as your own code — it can read your environment variables (including secrets), make network requests, read and write files, exactly like code you wrote yourself. A compromised dependency (via a hijacked maintainer account, a malicious package update, or a typosquatted package name similar to a popular legitimate one) is a real, documented attack vector, not a theoretical one — supply-chain attacks against widely-used package registries (npm, PyPI, and others) have caused real, significant incidents. Treating "I imported a library" as fundamentally different from "I wrote this code myself, from a trust perspective" is the actual mistake this lesson corrects.

**Pinning dependencies: eliminating "it changed underneath me" as an attack vector.** An unpinned or loosely-pinned dependency (`^1.2.0`, allowing any compatible minor/patch update) means your build can pull in a *different* version of a dependency than the one you actually reviewed or tested against, without you explicitly deciding to update — if that new version has been compromised (or simply has a new, unreviewed bug), you've inherited it automatically. Pinning to an exact version (and, more robustly, to a specific cryptographic hash of that version's contents, which most modern package managers support via lockfiles) means updates only happen when you explicitly choose to update and can review what changed.

**Dependency inventory: knowing what's actually in your build, transitively.** Most real projects have far more *transitive* dependencies (dependencies of your direct dependencies, and so on) than direct ones — a project with 20 direct dependencies might easily have hundreds of transitive ones, most never directly reviewed by you. A generated dependency inventory (via `go list -m all`, `npm ls --all`, or equivalent) makes this actual attack surface visible and auditable, rather than an unknown, unexamined quantity — and tools that check this inventory against known-vulnerability databases (e.g. `govulncheck` for Go, `npm audit`) let you proactively find and address known issues rather than discovering them reactively.

**CI/build pipeline as its own trust boundary, often under-scrutinized.** Your CI system typically has real, significant privileges — access to secrets (deployment credentials, signing keys), and the ability to produce and publish the actual artifact that gets deployed to production. A CI workflow that runs untrusted code (e.g. a pull request from an external contributor, if your project accepts them) with the same privileges as your trusted, internal builds is a real trust-boundary violation — many CI systems have specific, deliberate features (restricted secret access for PR-triggered builds, requiring maintainer approval before running CI on external contributions) precisely to prevent this, and not using them is a real, often-overlooked gap.

## Attempt

1. For a real project (TARDOC, Mahall, or Lead Sourcer), generate a complete dependency inventory (`go list -m all` for Go, or the equivalent for whatever language/tooling that project uses) and report the total count of direct versus transitive dependencies — confirm you're actually surprised (or not) by how many transitive dependencies exist relative to what you directly chose to add.

2. Run a vulnerability scanner against that same dependency tree (`govulncheck ./...` for Go, or `npm audit`/equivalent) and report any actual findings — for any real vulnerability found, note its severity and whether a fixed version is available, and if so, whether upgrading to it is a reasonable, low-risk change or something requiring more careful testing.

3. Audit your project's dependency-pinning configuration (its lockfile — `go.sum`, `package-lock.json`, or equivalent) and confirm it's actually being used correctly: are dependency versions pinned to exact, hash-verified versions, or is there any place in your build configuration that could pull an unpinned, "latest compatible" version despite the lockfile's presence (a real, easy-to-introduce misconfiguration in some build setups).

4. Review your project's actual CI/build configuration (a GitHub Actions workflow file, or equivalent) for trust-boundary issues per Learn: does it run any external/untrusted code (e.g. a PR from a fork) with access to secrets it shouldn't have? Are deployment credentials scoped to only what the build actually needs (least privilege, system-engineering Lesson 10), or broader than necessary? Report your actual findings.

## Verify

Present your step 1 dependency count (direct vs. transitive), your step 2 vulnerability scan results with any real findings, your step 3 pinning-configuration audit, and your step 4 CI trust-boundary review — real findings from a real project's actual configuration, not a hypothetical exercise.

## Failure drill

Deliberately introduce an unpinned or loosely-pinned dependency into a test project (or reason through your real project's lockfile mechanics if introducing an actual change isn't practical) and demonstrate that a build performed at two different times (or after the loosely-pinned dependency's upstream publishes a new version) can pull in different actual code without any explicit change to your own project's source — confirm this by checking the resolved version/hash before and after, showing it genuinely changed underneath you. Explain why this specific, demonstrated behavior is exactly the risk pinning is meant to eliminate — an attacker who compromises a popular, loosely-pinned dependency doesn't need to compromise *your* project at all; they only need to publish a malicious update to a dependency many projects loosely trust, and it propagates automatically into every build that doesn't pin tightly enough to prevent it.

## Transfer

If TARDOC, Mahall, or Lead Sourcer's CI pipeline currently has broader deployment credential access than strictly necessary for what each specific workflow actually does (a common, easy-to-accumulate gap as projects grow), describe what a properly scoped, least-privilege version of that access would look like, and what concrete change (splitting a broad credential into narrower, workflow-specific ones, or restricting which branches/triggers can access deployment secrets at all) would close the gap.

## Done when

You've generated and reviewed a real dependency inventory for an actual project, run a real vulnerability scan and reported genuine findings (or their absence), audited your actual lockfile/pinning configuration for real gaps, and reviewed your actual CI configuration for trust-boundary issues — all against real project configuration, not a hypothetical, plus demonstrated concretely (via the failure drill) how an unpinned dependency's resolved version can genuinely change underneath a project without any explicit source change.
