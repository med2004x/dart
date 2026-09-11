# Lesson 1: Read an Unfamiliar Repository

## Objective

Build an accurate mental model of a large, unfamiliar codebase before making any changes to it — a distinct, learnable skill from writing new code, and one this entire curriculum has been implicitly building toward through every "read before modifying" instruction in earlier tracks (OS Lesson 8's xv6 labs, database-internals' BusTub capstone, and others).

## Prerequisites

Linux tools Lesson 3 (Git as an engineering tool — `git log`/`git blame` are core tools for this lesson's investigation), discrete math Lesson 4 (graphs/DAGs — a codebase's module dependency structure is exactly this kind of graph, and mapping it is part of this lesson's task).

## Learn

**Why reading a large codebase is a genuinely different skill from writing one.** When you write code, you build your mental model incrementally, as you go — by the time the code exists, you already understand it, because you're the one who made every decision. Reading someone else's large, mature codebase requires reconstructing that understanding *after the fact*, from evidence (the code itself, tests, commit history, documentation) rather than lived decision-making — a skill that doesn't automatically transfer from being good at writing code, and one that's genuinely necessary for almost any real engineering job, where you'll spend far more time working in existing code than starting fresh.

**Entry points: where to actually start reading.** A large codebase has no single obvious starting point, but it does have identifiable entry points — the `main()` function or equivalent, the top-level HTTP route registrations, the primary public API surface. Starting from an entry point and tracing forward (what does this call, what does *that* call) is generally more productive than starting from a random file and trying to understand it in isolation, since an entry point gives you the actual, real execution context a piece of code operates within, rather than an abstract, disconnected reading.

**Tests as documentation, often more reliable than actual documentation.** A mature project's test suite frequently encodes the *intended* behavior more precisely and more currently than any prose documentation, which tends to drift out of sync with the actual code over time (a real, common problem, not a cynical assumption) — reading a module's tests before or alongside its implementation often clarifies intent (what inputs are expected, what edge cases were deliberately considered) faster than reading the implementation alone.

**Using `git log`/`git blame` as investigative tools, directly extending Linux tools Lesson 3.** For an unfamiliar piece of code that seems confusing or oddly specific, the commit history often explains *why* it's shaped that way — a bug fix, a specific edge case, a deliberate tradeoff — information invisible from reading the current code state alone, exactly as Linux tools Lesson 3 established.

**Producing a written architecture summary: the actual test of whether you've built a real mental model.** The discipline of writing down, concisely, what you've understood (module boundaries, data flow, key abstractions) forces you to notice gaps in your own understanding that feel like understanding until you try to articulate them precisely — a genuinely different and more rigorous test than simply feeling like you've "gotten the gist" of a codebase after skimming it.

## Attempt

1. Choose a mature, real, moderately-sized open-source Go, Rust, or C project (something genuinely substantial — thousands of lines, real history, real tests — not a toy example) that you haven't worked in before. A project related to something covered elsewhere in this curriculum (a database engine, an HTTP framework, a CLI tool you use) is a reasonable, motivated choice.

2. Identify and trace the project's entry point(s): find `main()` (or equivalent), and trace the first few layers of what it calls, building an initial map of the major subsystems/modules it wires together.

3. Read the test suite for one specific, moderately complex module or feature within the project, before reading that module's full implementation — note what you learn about intended behavior and edge cases purely from the tests, then read the implementation and confirm (or correct) your test-derived understanding.

4. Use `git log`/`git blame` (Linux tools Lesson 3) to investigate at least one piece of code that seems non-obvious or specifically shaped for a reason you don't immediately understand — find the commit(s) that explain why it's built that way, and confirm your investigation actually surfaced a real, specific reason (a bug fix, an edge case, a deliberate tradeoff documented in a commit message or linked issue).

## Verify

Write a genuine one-page architecture summary of the project: its major modules/subsystems, how they depend on each other (a dependency graph, per discrete math Lesson 4, if the structure is complex enough to warrant one), the entry point(s) and primary execution flow, and at least one specific piece of historical context you uncovered via step 4's investigation — a document specific enough that someone else could use it as a genuine starting point for their own exploration of the same codebase.

## Failure drill

Before doing any of the above steps, write a *quick, first-impressions* summary of the project based purely on skimming its README and file/directory names for 10 minutes — no deep reading. Then complete steps 2-4 properly, and compare your final, evidence-based architecture summary against your quick first-impressions one. Identify at least one place where your initial, surface-level understanding was meaningfully wrong or incomplete compared to what the actual investigation revealed. Explain why this comparison is valuable: it's a concrete, personal demonstration of why skimming alone (however tempting when facing a large, unfamiliar codebase under time pressure) produces a real, and easy to overestimate, gap in actual understanding compared to the entry-point-tracing, test-reading, history-investigating process this lesson actually requires.

## Transfer

Apply this lesson's methodology to a part of TARDOC, Mahall, or Lead Sourcer's own codebase that you didn't write yourself, or wrote long enough ago that it feels genuinely unfamiliar now — trace its entry point, read its tests before its implementation, and use `git log`/`git blame` to investigate at least one piece of code whose reasoning you don't currently remember. Even in your own project, this systematic approach may surface something you'd forgotten or never fully understood about your own past decisions.

## Done when

You've built and written down a genuine, evidence-based architecture summary of a real, unfamiliar codebase, you've used test-reading and git-history investigation as deliberate techniques (not just skimmed source files), and you've directly compared your rigorous, evidence-based understanding against a quick first-impressions guess, identifying a concrete gap between the two that demonstrates why the deeper process matters.
