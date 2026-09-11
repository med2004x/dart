# Lesson 3: Modular Monolith

## Objective

Design a single-deployable system with genuine internal module boundaries — clear interfaces, no circular dependencies, enforced separation — understanding why "modular monolith" is a deliberate middle ground between an unstructured single codebase and premature microservices, not a compromise to be embarrassed about.

## Prerequisites

Discrete math Lesson 4 (graphs/DAGs — module dependencies should form a DAG, directly reusing that lesson's cycle-detection concepts), software-engineering-master skill (architecture patterns generally, if available — this lesson's specific focus is the monolith-with-real-boundaries pattern specifically).

## Learn

**Why "just split everything into microservices" is often the wrong early default.** Microservices solve real problems (independent deployment, independent scaling, team autonomy at large organizational scale) at a real cost (network calls where function calls used to suffice, per distributed Lesson 1's failure-mode complexity; operational overhead of running and monitoring many services; harder-to-trace cross-service bugs). For a solo developer or small team building a system that doesn't yet have the scale or organizational reasons driving those benefits, that cost is frequently paid with none of the corresponding benefit realized — a genuinely common, well-documented mistake of premature architectural complexity.

**What a modular monolith actually is.** A single deployable application, internally organized into modules with clear, enforced boundaries — each module owns its own data access and exposes a defined interface to other modules, rather than every part of the codebase freely reaching into every other part's internals. This gets you most of microservices' organizational benefit (clear ownership boundaries, the ability to reason about one module without understanding the whole system) without the network-call and operational costs, and — importantly — without foreclosing a *later* migration to actual microservices, since well-defined internal module boundaries are precisely what makes extracting a module into its own service later a tractable, incremental change rather than a rewrite.

**Enforcing boundaries, concretely, not just as a documentation convention.** A boundary that exists only as a comment or a folder structure, with nothing preventing code in module A from directly importing and calling module B's internal, unexported functions or directly querying module B's database tables, isn't a real boundary — it's a suggestion that will erode under time pressure. Real enforcement in Go typically means: unexported (lowercase) types/functions for anything not meant to be part of a module's public interface, a clear package structure where cross-module calls only happen through exported interfaces, and ideally, each module owning its own database tables with no other module directly querying them (all cross-module data access going through the owning module's own interface, even internally).

**No circular dependencies — a real, checkable structural requirement, not just good taste.** If module A depends on module B, and module B depends on module A, you've created a circular coupling that makes it genuinely difficult to reason about, test, or later extract either module independently — this is directly the DAG requirement from discrete math Lesson 4's dependency-graph discussion, and Go's own compiler actually enforces the *package-level* version of this (circular package imports are a compile error), though nothing stops circular *conceptual* dependencies at a coarser module level if a module is spread across multiple packages without discipline.

## Attempt

1. For a real or plausible extension of TARDOC or Mahall, identify at least 4 natural module boundaries (e.g. for TARDOC: clinic management, subscription/billing, transcription processing, notifications) and, for each, state what data it owns and what operations it exposes to other modules.

2. Draw the dependency graph between your 4+ modules (which modules call into which others) and confirm, using discrete math Lesson 4's cycle-detection reasoning, that it's genuinely acyclic — if you find a cycle in your initial design, identify which dependency should be inverted or which shared concern should be extracted into a lower-level, shared module both can depend on without depending on each other.

3. Implement (or restructure existing code) so that cross-module calls happen only through explicitly exported interfaces — for at least one pair of modules, show the actual Go package structure with unexported internals and a clear, minimal exported interface, and confirm (by attempting to reach into the internals from outside the module and observing a compile error) that Go's own visibility rules are actually enforcing the boundary, not just your own discipline.

4. Identify one place in an existing codebase (yours or a hypothetical one matching this pattern) where a module boundary is currently violated — code in one conceptual area directly querying another module's database tables, or directly calling another module's unexported logic through some workaround. Describe specifically what would need to change to correct it, and what risk the current violation creates (e.g. changing module B's internal schema silently breaking module A, since nothing enforces that A only depends on B's stated interface).

## Verify

Present your step 2 dependency graph (confirmed acyclic) and your step 3 Go package structure demonstrating a real, compiler-enforced boundary between at least one module pair, including the actual compile error you triggered when attempting to violate it.

## Failure drill

Deliberately introduce a circular dependency between two of your modules (e.g. module A's exported interface needs to call into module B, and you also add a call from module B back into module A for a seemingly reasonable reason). Attempt to compile and observe Go's own circular-import error if the modules are separate packages — or, if you structured them within the same package (which wouldn't trigger Go's compiler check), reason explicitly through why this specific structural mistake is now much harder to detect, since nothing automatically flags a conceptual circular dependency within a single package the way Go's compiler flags one between packages. Explain why enforcing true package-level separation (not just folder organization within one package) is what actually gives you compiler-verified protection against this specific structural problem, rather than relying on manual discipline that erodes over time.

## Transfer

If TARDOC's actual codebase (restructured per your project history — "restructured the entire repo," mentioned among recent TARDOC work) currently has clear module boundaries matching this lesson's pattern, describe which modules you'd identify in the actual repo, and whether any of them currently violate the "no direct cross-module database access" principle from Learn — if you're not certain, this is worth actually checking in the real codebase rather than assuming, since boundary violations are exactly the kind of thing that accumulates silently under time pressure without a deliberate check.

## Done when

You've identified real, natural module boundaries for a genuine system and confirmed their dependency graph is acyclic, you've implemented at least one pair of modules with a compiler-enforced (not just documented) boundary and demonstrated the enforcement actually triggers a compile error on violation, and you've identified and reasoned about the risk of an actual or plausible boundary violation in a real or realistic codebase.
