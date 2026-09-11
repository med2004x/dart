# Lesson 5: Recurrences and Asymptotic Analysis

## Objective

Solve recurrence relations (the mathematical structure behind recursive algorithm running time) and master Big-O notation rigorously, not just informally — this is the direct mathematical foundation for algorithm complexity analysis you'll use constantly in the algorithms track.

## Prerequisites

Lesson 1 (induction — recurrence solutions are typically verified by induction), Lesson 1 of the analysis track (limits — Big-O is defined using a limiting/bounding argument).

## Learn

**Recurrence relations.** A recurrence defines a sequence (or a function's running time) in terms of smaller instances of itself. Example: `T(n) = 2T(n/2) + n` — this is literally the recurrence for merge sort's running time: splitting into two halves (`2T(n/2)`) plus linear-time merging (`n`).

**Solving by the Master Theorem.** For recurrences of the form `T(n) = aT(n/b) + f(n)` (a subproblems, each of size n/b, plus f(n) combining work):

- If `f(n) = O(n^(log_b(a) - ε))` for some ε>0, then `T(n) = Θ(n^log_b(a))` (the recursive splitting dominates).
- If `f(n) = Θ(n^log_b(a))`, then `T(n) = Θ(n^log_b(a) · log n)` (balanced — this is merge sort's case).
- If `f(n) = Ω(n^(log_b(a) + ε))` for some ε>0 (and a regularity condition holds), then `T(n) = Θ(f(n))` (the combining work dominates).

Worked example: solve `T(n) = 2T(n/2) + n` (merge sort).

Here `a=2, b=2, f(n)=n`. `log_b(a) = log₂(2) = 1`. `f(n) = n = n¹ = Θ(n^log_b(a))` — this is case 2 (balanced). So `T(n) = Θ(n log n)` — matching what you already know as merge sort's complexity, now derived rather than memorized.

**Big-O, formally.** `f(n) = O(g(n))` means there exist positive constants `c` and `n₀` such that `f(n) ≤ c·g(n)` for all `n ≥ n₀`. This is an upper bound "eventually" (for large enough n), up to a constant factor — it deliberately ignores constant factors and small-n behavior, which is exactly why Big-O tells you about scaling trends, not which of two O(n) algorithms is actually faster on your real input sizes (a lower-order-term-heavy O(n) algorithm can easily be slower than an O(n log n) one for realistic n — this is a genuine, common practical trap in reading Big-O too literally).

`Ω(g(n))` is the corresponding lower bound, and `Θ(g(n))` means both — a tight bound.

## Attempt

1. Solve `T(n) = T(n/2) + 1` using the Master Theorem (this is binary search's recurrence — halve the problem, constant work per step). Identify `a, b, f(n)`, determine which case applies, and state the resulting complexity class.

2. Solve `T(n) = 2T(n/2) + n²` using the Master Theorem (case 3 — the combining work dominates). State the resulting complexity and explain in one sentence why the `n²` combining step, not the recursive splitting, determines the overall growth here.

3. Prove `3n² + 5n + 2 = O(n²)` using the formal definition (find explicit constants `c` and `n₀` such that the inequality holds for all `n ≥ n₀` — e.g. try `c = 4` and find the smallest `n₀` that works by solving `3n²+5n+2 ≤ 4n²` for n).

4. Prove `n = O(n²)` is true but `n² = O(n)` is false, using the formal definition for the true case, and a proof-by-contradiction sketch for why no valid `c, n₀` can exist for the false case (as `n → ∞`, `n²/n = n` grows without bound, so no constant `c` can keep `n² ≤ cn` for all large n).

## Verify

For step 3, plug your chosen `c` and `n₀` back into the inequality with at least two specific values of `n ≥ n₀` and confirm numerically that `3n²+5n+2 ≤ c·n²` actually holds — don't just trust the algebra, check concrete numbers.

## Failure drill

Attempt to apply the Master Theorem to `T(n) = 2T(n/2) + n log n` and notice that `f(n) = n log n` doesn't cleanly fit any of the three cases as stated (it sits in the gap between case 2 and case 3 — technically case 2 requires `f(n) = Θ(n^log_b(a))` exactly, and `n log n` isn't polynomially equal to `n` in the required sense, though it's close). This is a genuine limitation of the basic Master Theorem, not a mistake on your part — record which case you tried to force it into and why the fit isn't clean, since recognizing when a tool doesn't apply cleanly is as important as applying it when it does (the actual answer requires the more general Akra-Bazzi method, which is out of scope here — the point of this drill is noticing the gap, not resolving it).

## Transfer

Take one recursive function you've implemented in the algorithms track (Lesson 09 recursion, Lesson 14 dynamic programming, or Lesson 12 trees) and write down its recurrence relation explicitly (how does the size of the subproblem(s) relate to the original, and how much non-recursive work happens per call), then solve it with the Master Theorem if it fits the form, or explain why it doesn't fit if it's not a simple `aT(n/b)+f(n)` shape (e.g., unequal subproblem sizes, as in some DP recurrences).

## Done when

You can identify `a`, `b`, and `f(n)` from a recurrence and correctly apply the Master Theorem's three cases, and you can produce a formal Big-O proof with explicit constants for a new polynomial, rather than asserting the bound by inspection alone.
