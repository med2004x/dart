# Lesson 4: Orthogonality and Least Squares

## Objective

Understand orthogonal projections and derive least-squares regression from a linear-algebra perspective, connecting it forward to the statistics track's regression lesson.

## Prerequisites

Lesson 1 (dot product — orthogonality is dot product zero), Lesson 2 (matrices, solving `Ax = b`).

## Learn

**Orthogonal vectors.** From Lesson 1, `u` and `v` are orthogonal if `u · v = 0`. An **orthogonal basis** is a set of mutually orthogonal vectors that span a space; an **orthonormal basis** additionally has each vector normalized to length 1. Orthonormal bases are computationally convenient because projections onto them decompose cleanly (no cross-terms to worry about).

**Projection.** The projection of vector `b` onto vector `a` is `proj_a(b) = (a·b / a·a)·a` — the "shadow" `b` casts on the line through `a`. This is the geometric core of least-squares regression.

**The overdetermined system problem.** Real data almost never satisfies `Ax = b` exactly — you have more equations (data points) than unknowns (parameters), and noise means no `x` fits perfectly. Instead of solving `Ax = b` exactly (usually impossible), least squares finds `x` that minimizes the squared error `|Ax - b|²` — the closest `Ax` can get to `b`.

**The normal equations.** The minimizing `x` satisfies `AᵀAx = Aᵀb` (derived by setting the gradient of `|Ax-b|²` to zero — a direct application of Lesson 5 of the analysis track's multivariable calculus). If `AᵀA` is invertible, `x = (AᵀA)⁻¹Aᵀb`. This is exactly what "linear regression" computes: `A` is your design matrix (each row a data point's features, with a column of 1s for the intercept), `b` is the observed outcomes, and `x` is the fitted coefficients.

Worked example: fit a line `y = mx + c` to points `(1,2), (2,3), (3,5)` using least squares.

```
A = [[1,1],[2,1],[3,1]]   (columns: x, then 1 for intercept)
b = [2,3,5]

AᵀA = [[1²+2²+3², 1+2+3],[1+2+3, 3]] = [[14,6],[6,3]]
Aᵀb = [1·2+2·3+3·5, 2+3+5] = [23,10]

Solve [[14,6],[6,3]] x = [23,10]:
From row 2: 6m + 3c = 10 → 2m + c = 10/3
From row 1: 14m + 6c = 23
Substitute c = 10/3 - 2m into row 1: 14m + 6(10/3 - 2m) = 23 → 14m + 20 - 12m = 23 → 2m = 3 → m = 1.5
c = 10/3 - 3 = 1/3
```

So the best-fit line is approximately `y = 1.5x + 0.33`.

## Attempt

1. Compute the projection of `b = (3, 4)` onto `a = (1, 0)`, and separately onto `a = (0, 1)`. Confirm these two projections sum to `b` exactly (this works here because `(1,0)` and `(0,1)` are orthonormal — state why this wouldn't generally work for a non-orthogonal pair).

2. Set up and solve the normal equations `AᵀA x = Aᵀb` by hand for fitting a line to the points `(0,1), (1,3), (2,4), (3,6)`, following the worked example's structure exactly.

3. Implement least-squares line fitting in Go or Python: build the design matrix `A` and vector `b` from a list of (x,y) points, compute `AᵀA` and `Aᵀb` (matrix operations from Lesson 2), and solve the resulting 2×2 system either with your Lesson 2 matrix-inversion code or by direct substitution. Test it against the points from step 2 and confirm it matches your hand-solved coefficients.

4. Run your implementation on a noisy synthetic dataset: generate points along `y = 2x + 1` with small random noise added to each `y` (e.g. `y = 2x + 1 + noise`, noise drawn from a small uniform or normal range) for 20+ points, and confirm the fitted `m` and `c` come out close to 2 and 1 respectively despite the noise.

## Verify

For step 3, your computed `(m, c)` must exactly match your hand-solved values from step 2 (within floating-point rounding). For step 4, report the fitted `m` and `c` and confirm they're within a reasonable tolerance of the true values (2 and 1) — the exact tolerance depends on how much noise you added, so state what noise range you used alongside your result.

## Failure drill

Fit a line to data that's *not* well-approximated by a line at all — e.g. points sampled from `y = x²` over a range like `x ∈ {-3,...,3}`. Run your least-squares fitter anyway (it will produce *some* line, since it always finds the best linear fit even to nonlinear data) and plot or print both the data and the fitted line. Explain why the residuals (differences between actual `y` and fitted `y`) are clearly not small and random here, unlike the noisy-linear case in step 4 — this is the geometric seed of why residual analysis (covered properly in the statistics track's regression lesson) matters for validating whether a linear model was appropriate at all.

## Transfer

If you've ever built a simple pricing formula, scoring function, or trend estimate in TARDOC, Mahall, or Lead Sourcer that was tuned by eye (manually adjusted coefficients to "look right" against some data), describe how you'd instead set it up as a least-squares problem: what would the design matrix's columns be (the features), and what would the target vector `b` represent.

## Done when

You can set up and solve the normal equations for a small dataset by hand, your least-squares implementation matches the hand-solved result and correctly recovers known coefficients from noisy synthetic data, and you can explain why least squares still produces an answer even when the true relationship isn't linear, and why that's a limitation to watch for rather than a feature.
