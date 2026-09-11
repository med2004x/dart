# Lesson 2: Sets, Relations, and Functions

## Objective

Formalize sets, relations, and functions rigorously, and connect equivalence relations and partial orders directly to database normalization and dependency graphs you already work with.

## Prerequisites

Lesson 1 (logic and proofs — set identities are proven using the same techniques).

## Learn

**Sets.** A set is an unordered collection of distinct elements. Standard operations: union `A ∪ B`, intersection `A ∩ B`, difference `A - B`, complement `Aᶜ`. Set identities (De Morgan's laws: `(A∪B)ᶜ = Aᶜ∩Bᶜ`, and `(A∩B)ᶜ = Aᶜ∪Bᶜ`) mirror boolean logic identities exactly — this is not a coincidence, sets under union/intersection/complement form the same algebraic structure as booleans under OR/AND/NOT, which is why SQL's `WHERE` clause logic and set-based query reasoning use the same rules.

**Relations.** A relation `R` on a set `A` is a subset of `A × A` (the set of ordered pairs). `aRb` means `(a,b) ∈ R`. Relations can have properties:
- *Reflexive*: `aRa` for all `a`.
- *Symmetric*: `aRb ⟹ bRa`.
- *Transitive*: `aRb ∧ bRc ⟹ aRc`.
- *Antisymmetric*: `aRb ∧ bRa ⟹ a = b`.

**Equivalence relations** are reflexive, symmetric, and transitive — they partition a set into disjoint equivalence classes (e.g. "same remainder mod 3" partitions integers into 3 classes). This is the formal structure behind grouping/deduplication logic: if you've ever grouped records by "considered the same" under some criterion, you were implicitly defining an equivalence relation, and it's worth checking your "sameness" criterion is actually transitive — a common real bug is a similarity check that's reflexive and symmetric but not transitive (A "matches" B, B "matches" C, but A doesn't "match" C), which breaks any code assuming equivalence classes behave cleanly.

**Partial orders** are reflexive, antisymmetric, and transitive (like `≤`, or "is a subtype of," or "must run before" in a dependency graph) — but unlike a total order, not every pair of elements needs to be comparable. This is exactly the structure of a task dependency graph: task A "must happen before" task B is a partial order, and topological sort (covered in the algorithms track) is the algorithm that produces a valid linear ordering consistent with a partial order.

**Functions.** A function `f: A → B` maps every element of `A` to exactly one element of `B`. *Injective* (one-to-one): distinct inputs map to distinct outputs. *Surjective* (onto): every element of `B` is hit by some input. *Bijective*: both — this is the formal notion behind "this mapping is reversible/invertible," directly relevant to hash function design (you want collisions to be rare, i.e. close to injective, even though a hash function from an infinite domain to a finite range can never be truly injective).

## Attempt

1. Prove De Morgan's law `(A∪B)ᶜ = Aᶜ∩Bᶜ` using the definition of set membership (show `x ∈ (A∪B)ᶜ ⟺ x ∈ Aᶜ∩Bᶜ` by unpacking what each side means logically, connecting back to Lesson 1's propositional logic).

2. Define the relation `R` on integers: `aRb` iff `a ≡ b (mod 3)` (same remainder when divided by 3). Prove `R` is an equivalence relation (check all three properties explicitly) and list the equivalence classes.

3. Take the relation "task A must run before task B" on a small set of 5 tasks with some dependency pairs of your choosing (e.g. a build pipeline: compile before test, test before package, etc.). Confirm it's a valid partial order by checking antisymmetry and transitivity on your specific example, and identify any pairs of tasks that are *incomparable* (neither must run before the other) — this incomparability is exactly what allows parallel execution in a real build system.

4. For a hash function `h(x) = x mod 10` mapping integers to `{0,...,9}`: determine whether it's injective (it isn't — find two inputs that collide) and whether it's surjective onto `{0,...,9}` (it is, given enough distinct inputs). Explain why a real hash function used for a hash table intentionally is not injective (its domain is far larger than its range) but should distribute collisions roughly uniformly.

## Verify

For step 2, explicitly write out which integers fall into each of the 3 equivalence classes (mod 3), for at least the range -5 to 10, and confirm no integer appears in two classes.

## Failure drill

Construct a relation that is reflexive and symmetric but *not* transitive — for example, "person A and person B have met in person" on a set of people, where A met B, B met C, but A never met C. Confirm explicitly that this relation fails the transitivity check with your constructed example, and explain why treating this kind of relation as if it defined clean equivalence classes (e.g. grouping "people who've met" into isolated clusters) would silently produce wrong groupings — someone in cluster {A,B} and someone in cluster {B,C} might get incorrectly treated as unrelated even though B connects them.

## Transfer

Look at any deduplication, grouping, or "considered equal" logic in Mahall or Lead Sourcer (e.g. deduplicating leads by fuzzy-matched company name, or grouping products). State explicitly whether the equality/similarity check used is actually transitive, and if you're not sure, construct a concrete 3-item example (like the failure drill) to test it.

## Done when

You can prove a set identity from definitions, correctly verify whether a given relation is an equivalence relation or partial order by checking each required property explicitly, and you can identify a real (or plausible) case in your own systems where a "similarity" relation might silently fail to be transitive.
