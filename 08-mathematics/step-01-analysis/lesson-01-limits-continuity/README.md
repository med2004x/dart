# Lesson 1: Limits and Continuity

## Objective

Build a rigorous (epsilon-delta) understanding of limits and continuity, not just calculator intuition, and connect it to why numerical algorithms behave the way they do.

## Prerequisites

None. This is the start of the analysis track.

## Learn

**Informal limit.** `lim(x→a) f(x) = L` means: as `x` gets arbitrarily close to `a` (but not equal to `a`), `f(x)` gets arbitrarily close to `L`.

**Formal (epsilon-delta) definition.** `lim(x→a) f(x) = L` if: for every `ε > 0`, there exists a `δ > 0` such that whenever `0 < |x - a| < δ`, it follows that `|f(x) - L| < δ`... more precisely `|f(x) - L| < ε`. In words: no matter how tight a tolerance `ε` you demand around `L`, you can find a small enough neighborhood `δ` around `a` (excluding `a` itself) that guarantees `f(x)` lands within that tolerance.

Worked example: prove `lim(x→2) (3x + 1) = 7` using epsilon-delta.

Given `ε > 0`, we need `δ > 0` such that `0 < |x - 2| < δ` implies `|(3x+1) - 7| < ε`.

`|(3x+1) - 7| = |3x - 6| = 3|x - 2|`.

We want `3|x - 2| < ε`, i.e. `|x - 2| < ε/3`. So choose `δ = ε/3`.

Check: if `0 < |x - 2| < δ = ε/3`, then `|(3x+1) - 7| = 3|x-2| < 3·(ε/3) = ε`. Done — for any ε, δ = ε/3 works.

**Continuity.** `f` is continuous at `a` if `lim(x→a) f(x) = f(a)` — the limit exists, `f(a)` is defined, and they're equal. Intuitively: no jump, hole, or asymptote at that point. A function continuous everywhere on an interval can be drawn without lifting the pen.

**Why this is not just formalism for its own sake.** Floating-point arithmetic is not continuous the way real-number arithmetic is — it has removable jumps from rounding. Numerical algorithms (root-finding, optimization, gradient descent) rely on continuity assumptions about the underlying mathematical function even while operating on a discrete floating-point approximation of it. When an algorithm fails to converge or behaves erratically near a boundary, understanding whether the underlying function is actually continuous there (or has a genuine discontinuity, like division by zero or a conditional branch in your model) is often the actual diagnosis.

## Attempt

1. Prove `lim(x→3) (2x - 1) = 5` using the epsilon-delta definition, following the worked example's structure exactly (isolate `|f(x) - L|` in terms of `|x - a|`, solve for the required `δ`).

2. Prove `lim(x→0) x² = 0` using epsilon-delta. This one is harder because the algebra is not linear — you'll need to bound `|x|` first (e.g. assume `δ ≤ 1` so that `|x| < 1`, which lets you bound `x²= |x|·|x| < |x|·1`).

3. Determine, with justification, whether each of these is continuous at the indicated point:
   - `f(x) = (x² - 1)/(x - 1)` at `x = 1` (note this is undefined at x=1 as written — is it removable?)
   - `f(x) = 1/x` at `x = 0`
   - `f(x) = |x|` at `x = 0`

4. In Go or Python, plot or numerically evaluate `f(x) = sin(x)/x` for `x` approaching 0 from both sides (e.g. x = 0.1, 0.01, 0.001, -0.1, -0.01, -0.001) and compare against the known limit `lim(x→0) sin(x)/x = 1`. Note that `f(0)` itself is undefined (0/0) even though the limit exists — this is the standard example of a removable discontinuity.

## Verify

For problems 1 and 2, your final answer should be an explicit formula for `δ` in terms of `ε` (e.g. `δ = ε/3`, or for problem 2, something like `δ = min(1, ε)`), and you should be able to verify it algebraically holds for at least one concrete numeric example (pick `ε = 0.01`, compute your `δ`, and check the inequality actually holds at the boundary).

## Failure drill

For `f(x) = 1/x`, attempt the epsilon-delta proof that `lim(x→0) f(x) = L` for some finite `L` and observe it cannot be made to work for any `L` — no `δ` can force `|1/x - L| < ε` for all `x` in a punctured neighborhood of 0, because `1/x` grows unboundedly as `x → 0`. Write one sentence explaining why this failure is the formal meaning of "the limit does not exist" (or diverges to infinity) rather than just an intuition.

## Transfer

Take a numerical function you might implement in code (e.g. a sigmoid `1/(1+e^-x)`, or a piecewise pricing function like TARDOC's tax bracket logic) and identify by inspection whether it is continuous everywhere in its domain, and if not, exactly where the discontinuity is and whether it's removable or a genuine jump.

## Done when

You can produce an epsilon-delta proof for a new linear or simple quadratic limit without referencing the worked example, and you can identify and classify discontinuities (removable vs. jump vs. infinite) in a function you haven't seen before.
