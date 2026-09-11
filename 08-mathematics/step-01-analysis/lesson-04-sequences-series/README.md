# Lesson 4: Sequences and Series

## Objective

Determine convergence or divergence of sequences and series rigorously, and connect series convergence to algorithm analysis (recurrences, amortized cost) and to floating-point summation error.

## Prerequisites

Lesson 1 (limits — sequence convergence is a limit as n → ∞).

## Learn

**Sequence convergence.** A sequence `(aₙ)` converges to `L` if for every `ε > 0`, there exists `N` such that for all `n > N`, `|aₙ - L| < ε` — the same epsilon-delta idea from Lesson 1, but with `n → ∞` (a natural number index) instead of `x → a` (a real variable).

Worked example: prove `aₙ = 1/n` converges to 0.

Given `ε > 0`, we need `N` such that `n > N` implies `|1/n - 0| < ε`, i.e. `1/n < ε`, i.e. `n > 1/ε`. So choose `N = 1/ε` (or `⌈1/ε⌉` to keep it an integer). For any `n > N`, `1/n < ε` holds. Done.

**Series.** A series `Σ aₙ` (sum to infinity) converges if the sequence of **partial sums** `Sₙ = a₁ + a₂ + ... + aₙ` converges to some finite limit. This is subtle: individual terms `aₙ` going to 0 is necessary but *not sufficient* for the series to converge — the harmonic series `Σ 1/n` has terms going to 0 but the sum diverges to infinity (this is a classic, genuinely surprising result worth internalizing, not just memorizing).

**Convergence tests** (used to determine convergence without computing the sum):
- Geometric series `Σ arⁿ` converges iff `|r| < 1`, with sum `a/(1-r)`.
- Ratio test: if `lim |aₙ₊₁/aₙ| = L`, the series converges if `L < 1`, diverges if `L > 1`, inconclusive if `L = 1`.
- Comparison test: if `0 ≤ aₙ ≤ bₙ` and `Σbₙ` converges, so does `Σaₙ`.

**Why this connects to your work as an engineer.** Recurrence relations for algorithm running time (e.g. `T(n) = 2T(n/2) + n` for merge sort) are analyzed using techniques directly related to series — the Master Theorem is essentially a convergence-test-style classification. Amortized analysis (why a dynamic array's `append` is O(1) amortized despite occasional O(n) resizes) is a geometric-series argument: the total cost of all resizes up to size `n` is a geometric series that sums to `O(n)`, spread over `n` appends. Floating-point summation of a long series in code accumulates rounding error, and understanding which series converge slowly (like harmonic) versus fast (like geometric with small ratio) tells you which sums are numerically risky to compute naively in a loop.

## Attempt

1. Prove `aₙ = (2n+1)/(n+3)` converges, and find its limit (hint: divide numerator and denominator by `n`, then take `n → ∞`; this limit is 2). Then write the epsilon-N proof for that limit, following the worked example's structure.

2. Determine convergence or divergence of `Σ (1/2)ⁿ` (geometric, find `r` and check `|r|<1`) and compute its sum using the geometric series formula.

3. Use the ratio test to determine whether `Σ n!/nⁿ` converges. Compute `lim |aₙ₊₁/aₙ|` and interpret the result.

4. In Go or Python, numerically sum the harmonic series `Σ 1/n` for `n = 1` to `100,000` and to `10,000,000`. Observe that the partial sum keeps growing (slowly) rather than settling — this is a live demonstration of divergence, since a genuinely convergent series' partial sums would stabilize as n grows, and this one visibly doesn't.

## Verify

For step 4, report both partial sums (n=100,000 and n=10,000,000) and confirm the second is measurably larger than the first (harmonic partial sums grow like `ln(n)`, so going from 10^5 to 10^7 terms should increase the sum by roughly `ln(10^7) - ln(10^5) = ln(100) ≈ 4.6`, growing without bound as n increases further — check your numbers are consistent with this logarithmic growth rate, not leveling off).

## Failure drill

Sum the harmonic series in *reverse order* (from `n = 10,000,000` down to `n = 1`) instead of forward, using 32-bit floating point (`float32` in Go) instead of 64-bit. Compare the final sum to the forward-order, 64-bit sum. Explain the discrepancy in terms of floating-point precision: adding a tiny number to a much larger accumulated sum loses precision (the tiny number can be partially or fully absorbed/rounded away), so summation order and precision both affect the actual computed result even though the true mathematical sum is order-independent.

## Transfer

Take the amortized-cost argument for dynamic array doubling (append cost, informally: most appends are O(1), occasional resizes are O(n), but averaged over n appends the resize costs sum to a geometric-like series bounded by O(n) total) and write it out as an explicit geometric series calculation: if resizes happen at sizes 1, 2, 4, 8, ..., n (doubling), what is `Σ (cost of each resize)`, and why is that sum O(n) rather than O(n log n)?

## Done when

You can prove sequence convergence via epsilon-N, correctly classify at least three different series as convergent or divergent using the appropriate test, and can explain in your own words why the harmonic series diverges despite its terms shrinking to zero.
