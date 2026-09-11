# Lesson 1: Vectors

## Objective

Understand vectors as both geometric objects and data structures, master dot product and norms, and connect them to concrete uses (similarity search, embeddings) relevant to real backend systems.

## Prerequisites

None. This is the start of the linear algebra track.

## Learn

**A vector** is an ordered list of numbers, `v = (v₁, v₂, ..., vₙ)`, interpreted either geometrically (a point or arrow in n-dimensional space) or as data (a row of features, a word embedding, a database record's numeric columns). Both interpretations are simultaneously valid and useful — this dual reading is exactly why linear algebra underlies both computer graphics and machine learning.

**Vector addition and scalar multiplication:** `u + v = (u₁+v₁, ..., uₙ+vₙ)`, `c·v = (c·v₁, ..., c·vₙ)`. Geometrically, addition is "tip to tail" combination; scalar multiplication stretches or shrinks (and flips direction if `c < 0`).

**Dot product:** `u · v = u₁v₁ + u₂v₂ + ... + uₙvₙ` (a scalar, not a vector). Two equivalent formulas connect it to geometry: `u · v = |u||v|cos(θ)`, where `θ` is the angle between the vectors. This means the dot product measures *alignment*: it's large and positive when vectors point the same direction, zero when perpendicular (orthogonal), negative when pointing opposite directions.

**Norm (length):** `|v| = √(v₁² + v₂² + ... + vₙ²)` — this is the Euclidean/L2 norm, the direct generalization of the Pythagorean theorem to n dimensions. Note `|v| = √(v · v)`.

**Why this matters directly for you:** pgvector, which TARDOC already uses, stores embeddings as vectors and computes similarity using cosine similarity — `cos(θ) = (u · v)/(|u||v|)`, exactly the rearranged dot-product formula above. When you write a similarity search query, you are literally invoking this formula. Understanding it means you can reason about *why* cosine similarity ignores vector magnitude and only cares about direction (angle), which is a deliberate and important property for comparing embeddings of different lengths/scales fairly.

## Attempt

1. For `u = (1, 2, 3)` and `v = (4, -1, 2)`, compute `u + v`, `2u - v`, `u · v`, `|u|`, and `|v|`.

2. Determine whether `u = (1, 2)` and `v = (-4, 2)` are orthogonal (perpendicular) using the dot product. Then determine whether `u = (1, 2)` and `v = (2, 4)` point in the same direction, using either the dot product formula or by inspection.

3. Compute the cosine similarity between `a = (1, 0, 1, 0)` and `b = (0, 1, 0, 1)` (should be 0 — completely dissimilar/orthogonal), and between `a = (1, 0, 1, 0)` and `c = (2, 0, 2, 0)` (should be 1 — same direction, different magnitude, demonstrating that cosine similarity ignores scale).

4. Implement a `cosineSimilarity(a, b []float64) float64` function in Go (or Python) from the raw formula (dot product, then divide by product of norms — do not use a library), and verify it against your hand-computed answers from step 3.

## Verify

Your `cosineSimilarity` function's output for step 3's two pairs should match your hand-computed values (0 and 1 respectively) to within floating-point precision (say, 1e-9). Test at least one more pair with a known angle (e.g. two identical vectors should give cosine similarity exactly 1).

## Failure drill

Call your `cosineSimilarity` function with one vector being the zero vector `(0, 0, 0, 0)`. Observe what happens (likely division by zero, producing `NaN` or `Inf` in Go, or a runtime error in Python). Explain why this is not a bug in your formula but a genuine mathematical undefined case — the angle between the zero vector and anything is undefined, since the zero vector has no direction — and state what defensive check a production implementation (like the kind pgvector or an embeddings pipeline would need) should include.

## Transfer

If TARDOC's pgvector-based knowledge base (mentioned in your project history) does similarity search, describe — without necessarily reading the actual query code right now — what you'd expect the underlying SQL/pgvector operator to be computing, in terms of the dot product and norm formulas from this lesson, and why cosine similarity (rather than raw Euclidean distance) is often preferred for comparing text embeddings specifically.

## Done when

You can compute dot product, norm, and cosine similarity by hand for small vectors, your from-scratch implementation matches your hand calculations, and you can explain in one sentence why cosine similarity is undefined for a zero vector.
