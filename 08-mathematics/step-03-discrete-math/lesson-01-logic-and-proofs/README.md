# Lesson 1: Logic and Proof Techniques

## Objective

Master propositional logic notation and the standard proof techniques (direct, contrapositive, contradiction, induction) that underlie every correctness argument you'll make about algorithms in this track.

## Prerequisites

None. This is the start of the discrete math track.

## Learn

**Propositional logic.** Statements combine via AND (`∧`), OR (`∨`), NOT (`¬`), implication (`→`), and biconditional (`↔`). `P → Q` is false only when `P` is true and `Q` is false — it's *vacuously true* whenever `P` is false, which trips people up (e.g. "if pigs fly, then 2+2=5" is true as a logical statement, because the premise is false). This matters directly for reading code: an `if` guard that's never triggered doesn't make the code inside it "true" or "tested," it makes the implication vacuously satisfied without ever being checked.

**Contrapositive.** `P → Q` is logically equivalent to `¬Q → ¬P`. Sometimes the contrapositive is much easier to prove directly than the original statement.

**Proof by direct argument.** Assume the hypothesis, chain logical steps, arrive at the conclusion.

**Proof by contradiction.** Assume the negation of what you want to prove, derive a logical impossibility, conclude the original statement must be true.

Worked example (contradiction): prove `√2` is irrational.

Assume, for contradiction, that `√2 = p/q` in lowest terms (p, q integers, no common factor). Then `2 = p²/q²`, so `p² = 2q²`, so `p²` is even, so `p` is even (odd² is odd), so `p = 2k` for some integer k. Substituting: `4k² = 2q²`, so `q² = 2k²`, so `q` is also even. But then `p` and `q` share a factor of 2, contradicting "lowest terms." Contradiction — so `√2` cannot be rational.

**Proof by mathematical induction.** To prove a statement `P(n)` holds for all natural numbers `n ≥ n₀`: (1) *base case* — prove `P(n₀)`; (2) *inductive step* — assume `P(k)` holds (inductive hypothesis) and prove `P(k+1)` follows from it. This is directly analogous to recursion in code: the base case is the recursion's base case, and the inductive step is the recursive case assuming the smaller subproblem is already correctly solved.

Worked example (induction): prove `1 + 2 + ... + n = n(n+1)/2` for all `n ≥ 1`.

*Base case* (n=1): LHS = 1, RHS = 1·2/2 = 1. Equal. ✓

*Inductive step:* assume `1+2+...+k = k(k+1)/2` (hypothesis). Show `1+2+...+k+(k+1) = (k+1)(k+2)/2`.

`1+2+...+k+(k+1) = k(k+1)/2 + (k+1)` [by hypothesis] `= (k+1)[k/2 + 1] = (k+1)(k+2)/2`. Matches the target. ✓

By induction, the formula holds for all `n ≥ 1`.

## Attempt

1. Determine, with a truth table, whether `(P → Q) ∧ (Q → R)` implies `P → R` (this is the transitivity of implication — confirm it's a tautology, true under every assignment of P, Q, R).

2. Prove by contradiction: there are infinitely many prime numbers. (Standard proof: assume finitely many primes `p₁,...,pₙ`, consider `N = p₁·p₂·...·pₙ + 1`, show N is either itself prime or has a prime factor not in the list — either way, contradiction.)

3. Prove by induction: `n² ≥ n` for all `n ≥ 1`. Then prove by induction (a harder one): `2ⁿ > n` for all `n ≥ 1`.

4. Prove by induction that a recursive function computing the sum `1+2+...+n` (`func sum(n int) int { if n == 0 { return 0 }; return n + sum(n-1) }`) is correct — i.e., prove the recursive definition actually computes `n(n+1)/2`. This connects the induction proof directly to verifying recursive code, not just abstract formulas.

## Verify

For each induction proof, explicitly label your base case and inductive step, and state the inductive hypothesis in words before using it — a proof that skips stating the hypothesis explicitly is a common way to accidentally beg the question.

## Failure drill

Attempt to "prove" a false statement by induction where the inductive step secretly relies on the base case being a specific unstated value that doesn't actually work — for example, try to "prove" `n² = n` for all `n ≥ 1` (false for n=2) and see where the argument breaks. Identify precisely which step fails (it should be the base case itself, since n=1 gives `1=1` which appears to check out, so instead check n=2 explicitly against the claim and confirm the claim itself is false — the exercise is to notice you must always sanity-check even a syntactically valid-looking induction against a concrete counterexample).

## Transfer

Take one recursive function from your own code (Lesson 09 of the algorithms track covers recursion, or use any recursive function you've already written) and write an informal induction-style justification for why it's correct: state the base case, state what you assume is true for a smaller input, and state why the recursive case combines that assumption correctly to handle the current input.

## Done when

You can write a direct, contradiction, and induction proof for a new statement without prompting, and you can explain in your own words why induction is the formal justification for why recursive code with a correct base case and correct recursive step actually terminates with the right answer.
