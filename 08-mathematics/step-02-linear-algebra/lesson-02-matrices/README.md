# Lesson 2: Matrices

## Objective

Understand matrices as linear transformations and as systems-of-equations solvers, master matrix multiplication and inversion, and connect them to solving real linear systems in code.

## Prerequisites

Lesson 1 (vectors — a matrix transforms vectors, and its columns/rows are themselves vectors).

## Learn

**A matrix** is a rectangular array of numbers, and it has two equally important interpretations:

1. **A linear transformation.** A matrix `A` (size m×n) takes an n-dimensional vector and produces an m-dimensional vector via `Av`. This transformation is *linear*: `A(u + v) = Au + Av` and `A(cv) = c(Av)`. Rotations, scaling, projections, and reflections in graphics are all matrices.

2. **A system of linear equations.** The system
   ```
   2x + 3y = 8
   x - y = 1
   ```
   is exactly `Ax = b` where `A = [[2,3],[1,-1]]`, `x = (x,y)`, `b = (8,1)`. Solving the system is finding the vector `x` that this matrix maps to `b`.

**Matrix multiplication.** `(AB)ᵢⱼ = Σₖ Aᵢₖ·Bₖⱼ` — row `i` of A dotted with column `j` of B. Matrix multiplication is *not commutative* in general (`AB ≠ BA`), which surprises people coming from ordinary number arithmetic — this reflects that transformations don't generally commute (rotate-then-scale differs from scale-then-rotate).

Worked example: multiply `A = [[1,2],[3,4]]` and `B = [[5,6],[7,8]]`.

```
AB = [[1·5+2·7, 1·6+2·8], [3·5+4·7, 3·6+4·8]]
   = [[5+14, 6+16], [15+28, 18+32]]
   = [[19, 22], [43, 50]]
```

**Solving Ax = b.** For a small system, Gaussian elimination (row-reduce the augmented matrix `[A | b]` to reduced row-echelon form) gives the solution directly. For larger systems in code, you don't hand-roll Gaussian elimination — you use a library (LAPACK-backed, e.g. `gonum` in Go, `numpy.linalg.solve` in Python) — but understanding what the library is doing underneath matters for reasoning about when a system has no solution, one solution, or infinitely many (determined by whether `A` is invertible).

**Invertibility.** A square matrix `A` is invertible if there exists `A⁻¹` such that `AA⁻¹ = I` (identity matrix). If invertible, `Ax = b` has the unique solution `x = A⁻¹b`. If not invertible (singular), the system either has no solution or infinitely many, depending on `b`. A matrix is singular exactly when its determinant is 0 (Lesson 3 covers this alongside eigenvalues) — geometrically, a singular matrix collapses n-dimensional space into a lower dimension, losing information irreversibly.

## Attempt

1. Compute `AB` and `BA` for `A = [[1,0],[2,1]]`, `B = [[1,1],[0,1]]`. Confirm they are different, demonstrating non-commutativity concretely.

2. Solve the system `2x + 3y = 8`, `x - y = 1` by hand using substitution or elimination. Then write it as `Ax = b` and confirm your solution satisfies the matrix equation by computing `Ax` and checking it equals `b`.

3. Find `A⁻¹` for `A = [[2,1],[1,1]]` by hand (for a 2×2 matrix `[[a,b],[c,d]]`, `A⁻¹ = (1/(ad-bc))·[[d,-b],[-c,a]]`). Verify by computing `AA⁻¹` and confirming it equals the identity matrix.

4. Implement 2×2 and 3×3 matrix multiplication from scratch in Go (plain nested loops over `[][]float64`, no library), and verify your implementation against the hand-computed result from step 1.

## Verify

Your matrix multiplication implementation's output for step 1 should exactly match your hand-computed `AB` and `BA`. For step 3, `AA⁻¹` should equal the identity matrix `[[1,0],[0,1]]` to within floating-point rounding.

## Failure drill

Attempt to invert a singular matrix, e.g. `A = [[1,2],[2,4]]` (note row 2 is exactly 2× row 1 — the rows are linearly dependent). Compute the determinant `ad-bc = 1·4 - 2·2 = 0` and observe that the inversion formula divides by zero. Explain geometrically what's happening: this matrix maps every input vector onto a single line (it collapses 2D space to 1D), so there's no way to uniquely "undo" the transformation — information about which input produced a given output on that line is lost.

## Transfer

If Mahall's or TARDOC's data pipeline ever computes a linear combination of features, applies a fixed transformation to a set of data points, or solves any kind of "best fit" problem (even informally, like a pricing formula that's a weighted sum of inputs), describe how that operation could be expressed as a matrix-vector product, even if the actual code doesn't use matrix libraries explicitly.

## Done when

You can multiply two matrices by hand and verify it in code, you can solve a small linear system and verify the solution by substitution, and you can explain — using the determinant-zero example — what it means geometrically for a matrix to be non-invertible.
