# Lesson 3: Eigenvalues and Eigenvectors

## Objective

Understand eigenvalues/eigenvectors as directions a transformation only scales (doesn't rotate), compute them by hand for small matrices, and connect them to PageRank-style ranking algorithms and PCA (dimensionality reduction).

## Prerequisites

Lesson 2 (matrices as linear transformations).

## Learn

**Definition.** For a square matrix `A`, a nonzero vector `v` is an eigenvector with eigenvalue `λ` if `Av = λv` — applying the transformation to `v` doesn't change its direction, only scales it by `λ` (or flips it, if `λ < 0`). Most vectors get both rotated and scaled by a general matrix; eigenvectors are the special directions where only scaling happens.

**Finding eigenvalues.** `Av = λv` rearranges to `(A - λI)v = 0`. For a nonzero `v` to satisfy this, `(A - λI)` must be singular (from Lesson 2: a matrix maps a nonzero vector to zero only if it's singular), so `det(A - λI) = 0`. This is the **characteristic equation**, and solving it for `λ` gives the eigenvalues.

Worked example: find eigenvalues of `A = [[2,1],[1,2]]`.

```
det(A - λI) = det([[2-λ, 1],[1, 2-λ]]) = (2-λ)² - 1 = 0
(2-λ)² = 1
2-λ = ±1
λ = 1 or λ = 3
```

For `λ = 3`: solve `(A - 3I)v = 0`, i.e. `[[-1,1],[1,-1]]v = 0`, giving `v = (1,1)` (or any scalar multiple).

For `λ = 1`: solve `(A - I)v = 0`, i.e. `[[1,1],[1,1]]v = 0`, giving `v = (1,-1)`.

**Why this matters practically.** PageRank (the algorithm that made Google's search ranking work) is literally finding the dominant eigenvector of a matrix representing the web's link structure — the eigenvector with the largest eigenvalue corresponds to the steady-state distribution of "importance" flowing through links. Principal Component Analysis (PCA), used for dimensionality reduction, finds the eigenvectors of a data covariance matrix — the eigenvector with the largest eigenvalue is the direction of maximum variance in the data, which is the axis you keep when compressing high-dimensional data to fewer dimensions.

## Attempt

1. Find the eigenvalues and corresponding eigenvectors of `A = [[4,0],[0,3]]` (a diagonal matrix — this should be immediate by inspection once you understand why; state why diagonal matrices have obvious eigenvalues).

2. Find the eigenvalues and eigenvectors of `A = [[3,1],[0,2]]` by hand, following the worked example's method (characteristic equation, then solve for eigenvectors at each eigenvalue).

3. Implement the **power iteration** method in Go or Python: starting from a random vector `v₀`, repeatedly compute `v_{k+1} = Av_k / |Av_k|` (normalize after each multiplication). Run this for the matrix from the worked example, `A = [[2,1],[1,2]]`, for at least 20 iterations, and observe `v_k` converging to the dominant eigenvector `(1,1)/√2` (normalized).

4. Track the ratio `|Av_k|/|v_k|` at each iteration of step 3 — this converges to the dominant eigenvalue (3, from the worked example) as a side effect of power iteration, without ever solving the characteristic equation directly.

## Verify

Report the vector `v_k` and the ratio from step 4 at iterations 5, 10, and 20 — the vector should be visibly converging toward `(1,1)` normalized (approximately `(0.707, 0.707)`), and the ratio should be converging toward 3.

## Failure drill

Run power iteration on a matrix with two eigenvalues of *equal* magnitude but opposite sign, e.g. `A = [[0,1],[1,0]]` (eigenvalues are +1 and -1). Observe that the iteration does not converge to a single vector — it oscillates between two vectors instead. Explain why: power iteration relies on the dominant eigenvalue being strictly larger in magnitude than all others, and this matrix has two eigenvalues tied in magnitude, breaking that assumption.

## Transfer

Describe, at a conceptual level (no implementation required), how you'd adapt power iteration into a toy version of PageRank: given a small graph of 4-5 web pages with links between them represented as a matrix (row/column i,j = 1 if page i links to page j, normalized so each row/column sums appropriately), state what the dominant eigenvector of that matrix would represent, and why the "importance" of a page is defined recursively (a page is important if important pages link to it) rather than by simple in-link counting.

## Done when

You can compute eigenvalues and eigenvectors for a 2×2 matrix by hand via the characteristic equation, your power iteration implementation converges to the correct dominant eigenvector and eigenvalue on a known example, and you can explain in your own words why power iteration fails when the top two eigenvalues are tied in magnitude.
