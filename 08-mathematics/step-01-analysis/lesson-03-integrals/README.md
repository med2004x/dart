# Lesson 3: Integrals

## Objective

Understand the definite integral as a limit of Riemann sums, connect it to the derivative via the Fundamental Theorem of Calculus, and implement numerical integration.

## Prerequisites

Lesson 2 (derivatives — the Fundamental Theorem connects the two).

## Learn

**The area problem.** The definite integral `∫[a to b] f(x) dx` is defined as the limit of a Riemann sum: partition `[a, b]` into `n` subintervals of width `Δx = (b-a)/n`, sum `f(xᵢ)·Δx` over sample points `xᵢ` in each subinterval, and take the limit as `n → ∞` (equivalently, as `Δx → 0`). Geometrically this is "area under the curve," but the same construction — sum of many small contributions, taken to a limit — generalizes to anything accumulated continuously (total distance from a velocity function, total cost from a marginal cost function, total probability from a density function in Lesson 04 of the probability track).

**Fundamental Theorem of Calculus (FTC).** This is the result that makes integrals computable without summing infinitely many rectangles by hand. If `F' = f` (F is an antiderivative of f), then `∫[a to b] f(x) dx = F(b) - F(a)`. This says integration and differentiation are inverse operations — a genuinely deep and non-obvious fact, not a notational coincidence.

Worked example: compute `∫[0 to 2] x² dx` two ways.

*Via FTC:* an antiderivative of `x²` is `F(x) = x³/3`. So `∫[0 to 2] x² dx = F(2) - F(0) = 8/3 - 0 = 8/3`.

*Via Riemann sum (right-endpoint, to check):* with `n` subintervals of width `Δx = 2/n`, sample points `xᵢ = i·Δx`. Sum `Σ f(xᵢ)Δx = Σ (i·2/n)²·(2/n)`. As `n → ∞` this converges to `8/3` — the FTC gives you the answer directly instead of taking this limit by hand.

**Numerical integration** approximates the Riemann sum with a finite `n`, which is exactly what you do in code when there's no closed-form antiderivative (common for real-world functions, e.g. a probability density with no elementary antiderivative). The trapezoidal rule and Simpson's rule are standard finite approximations more accurate than naive rectangles for the same `n`.

## Attempt

1. Compute `∫[1 to 3] (2x + 1) dx` using the FTC (find an antiderivative, evaluate at both bounds, subtract). Show the antiderivative you used.

2. Compute `∫[0 to π] sin(x) dx` using the FTC (antiderivative of `sin(x)` is `-cos(x)`).

3. Implement numerical integration in Go or Python using the trapezoidal rule for `f(x) = x²` over `[0, 2]`, with `n = 10`, `n = 100`, and `n = 1000` subintervals. Compare each result against the exact value `8/3` from the Learn section's worked example, and record the error at each `n`.

4. Implement the same trapezoidal integration for a function with no simple closed-form antiderivative: `f(x) = e^(-x²)` over `[0, 1]` (this is proportional to the Gaussian/normal distribution's density, which reappears in the probability and statistics tracks). Use `n = 1000`.

## Verify

For step 3, your error should shrink as `n` increases (trapezoidal rule error is `O(1/n²)`, so going from n=10 to n=100 should reduce error by roughly 100x, not just 10x — check whether your numbers actually show this quadratic improvement or only linear, and if only linear, look for a bug in your implementation).

## Failure drill

Deliberately implement the sum using left-endpoints only, with a small `n` (e.g. `n = 4`), on a function that's strictly increasing on the interval (like `f(x) = x²` on `[0,2]`). Compare against right-endpoint and against the true value. Explain why left-endpoint systematically underestimates and right-endpoint systematically overestimates for an increasing function, and why this bias shrinks as `n` grows even though it doesn't disappear at any finite `n`.

## Transfer

If TARDOC or any of your systems compute something that's naturally a continuous accumulation approximated discretely (e.g., integrating a rate over time — total cost accrued from a per-minute billing rate, or total risk exposure over a time window), identify one such place, even informally, and state whether a discrete sum you already compute is effectively a Riemann sum approximation of a continuous integral.

## Done when

You can compute a definite integral via the FTC by finding an antiderivative, you understand why FTC is a nontrivial result (derivative and integral as inverse operations) rather than a definition, and your trapezoidal rule implementation shows measurably decreasing error as `n` increases on a function with a known exact answer.
