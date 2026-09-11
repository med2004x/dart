# Lesson 3: Combinatorics

## Objective

Master counting techniques (permutations, combinations, the pigeonhole principle) and connect them to hash collision probability and algorithmic complexity of brute-force search spaces.

## Prerequisites

Lesson 1 (proofs — combinatorial identities are often proven by induction or bijection).

## Learn

**The multiplication principle.** If task 1 can be done in `m` ways and task 2 in `n` ways (independently), both together can be done in `m·n` ways. This single idea underlies almost everything else in this lesson.

**Permutations.** The number of ways to arrange `n` distinct items in order is `n! = n·(n-1)·...·1`. The number of ways to choose and order `k` items from `n` is `P(n,k) = n!/(n-k)!`.

**Combinations.** The number of ways to choose `k` items from `n` *without* regard to order is `C(n,k) = n!/(k!(n-k)!)`, written `(n choose k)`. The relationship `P(n,k) = C(n,k)·k!` makes sense directly: choose the k items first (C(n,k) ways), then arrange them (k! ways).

Worked example: how many ways to choose a 3-person team from 8 people, and separately, how many ways to choose a president, VP, and secretary (ordered roles) from the same 8?

*Team (unordered):* `C(8,3) = 8!/(3!·5!) = (8·7·6)/(3·2·1) = 336/6 = 56`.

*Roles (ordered):* `P(8,3) = 8!/5! = 8·7·6 = 336`.

Note `P(8,3) = C(8,3)·3! = 56·6 = 336` — confirms the relationship above.

**The pigeonhole principle.** If you place `n` items into `m` containers and `n > m`, at least one container holds more than one item. This sounds trivial but proves surprisingly strong results.

Worked example (pigeonhole): in any group of 13 people, at least two share a birth month.

12 months = 12 "pigeonholes." 13 people = 13 "pigeons." Since 13 > 12, by pigeonhole at least one month contains 2+ people. Done — no need to know anything about the actual people.

**Why this matters concretely.** Hash collision probability (the "birthday problem") is a direct combinatorics application: with `n` items hashed into a table of size `m`, the probability of at least one collision grows surprisingly fast — for a 365-slot table, a 50% chance of collision occurs at only 23 items, far fewer than the naive "half of 365" intuition suggests. This is exactly why hash table load factor and collision handling matter even for tables that seem "mostly empty." Combinatorics also directly bounds brute-force search: if you're evaluating "try all permutations" or "try all subsets," `n!` or `2ⁿ` growth is why brute force becomes infeasible past very small `n`, motivating the algorithmic techniques (DP, pruning, heuristics) in the algorithms track.

## Attempt

1. A password must be exactly 8 characters, each character either a lowercase letter (26 options) or a digit (10 options), with repetition allowed. How many possible passwords are there? (Multiplication principle, not permutations — repetition is allowed here.)

2. From a deck of 52 cards, how many distinct 5-card hands are possible (order doesn't matter — this is poker-hand counting)? Use `C(52,5)`.

3. Use the pigeonhole principle to prove: in any list of 11 integers, at least two have the same remainder when divided by 10. State the pigeonholes and pigeons explicitly, following the birth-month worked example's structure.

4. Implement the birthday-problem probability calculation in Go or Python: for a hash table (or shared-birthday scenario) with `m` slots and `n` random items, compute the probability of at least one collision using `P(at least one collision) = 1 - P(no collisions) = 1 - [m·(m-1)·...·(m-n+1)] / mⁿ`. Compute this for `m = 365` at `n = 10, 23, 50, 100` and confirm the well-known result that `n=23` already gives roughly 50% collision probability.

## Verify

For step 4, your computed probability at `n=23, m=365` should be close to 0.507 (the standard birthday-problem result) — if your number is far off, check whether you're computing `P(no collision)` correctly as a *product* of decreasing fractions, not a sum, and whether you're subtracting from 1 correctly.

## Failure drill

Compute the birthday-problem probability the "intuitive but wrong" way: naively assume `P(collision) ≈ n/m` (linear guess — e.g. 23/365 ≈ 6%) and compare it to your correct calculation from step 4 (≈50%). The gap between 6% and 50% is large — explain in 2-3 sentences why the correct calculation grows so much faster than the naive linear guess (hint: it's driven by the number of *pairs* being compared, `C(n,2)`, which grows quadratically in `n`, not the number of items itself).

## Transfer

If Lead Sourcer or TARDOC ever hashes identifiers (emails, phone numbers, generated IDs) into a fixed-size structure (a hash map, a short generated code, a database shard key), use this lesson's birthday-problem formula to estimate, even roughly, how many items you could insert before collision probability crosses 1% or 50% for whatever key space size is actually in use (e.g. if using a 6-digit random ID, `m = 1,000,000` — at what `n` does collision probability become non-negligible?).

## Done when

You can correctly choose between permutation and combination formulas for a new counting problem and justify the choice, you can construct a pigeonhole argument for a new claim, and you can compute a real birthday-problem collision probability and explain why it grows faster than naive linear intuition suggests.
