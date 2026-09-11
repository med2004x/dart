# Lesson 5: Pipelines and Branch Prediction

## Objective

Explain instruction pipelining, identify the three hazard types that break it, and explain why branch misprediction is expensive and how prediction mitigates it.

## Prerequisites

Lesson 3 (the fetch-decode-execute-memory-writeback stages being pipelined).

## Learn

**The idea.** In the single-cycle or multi-cycle datapath from Lesson 3, one instruction fully completes before the next one starts any stage. Pipelining overlaps them: while instruction 1 is in its execute stage, instruction 2 can already be in decode, and instruction 3 in fetch. With 5 stages, up to 5 instructions can be in flight simultaneously. Throughput approaches one instruction completed per cycle, even though each individual instruction still takes 5 cycles start-to-finish (latency is unchanged; throughput improves).

**Hazards** are situations where this overlap produces wrong results if not handled.

- *Structural hazard*: two instructions need the same hardware resource in the same cycle (e.g. both want to access memory at once). Fixed by duplicating hardware or stalling one instruction.
- *Data hazard*: an instruction needs a value that a previous, still-in-flight instruction hasn't produced yet. Example: `add r1, r2, r3` immediately followed by `sub r4, r1, r5` — the `sub` needs `r1` before the `add` has finished writing it back. Fixed by **forwarding** (routing the ALU result directly to the next stage that needs it, bypassing the register file) or, when forwarding isn't enough, stalling.
- *Control hazard*: a branch instruction's outcome (taken or not taken, and to where) isn't known until it's evaluated, but the pipeline has already fetched the next 1-4 instructions assuming a particular outcome. If the guess was wrong, all that fetched work must be discarded — a **pipeline flush**.

**Branch prediction** exists to reduce control hazard cost. Instead of stalling until every branch resolves, the CPU predicts the outcome (based on history — "this branch was taken the last 9 times, predict taken") and speculatively continues fetching down the predicted path. If correct, no cost. If wrong, flush and restart — the deeper the pipeline, the more cycles wasted per misprediction. Modern CPUs achieve 95%+ prediction accuracy on typical code, but the cost of the remaining mispredictions is still measurable and matters for hot loops with unpredictable branches (e.g. branching on data that's effectively random).

**Why this matters:** this is the mechanism behind advice like "avoid unpredictable branches in hot loops" or "branchless code can be faster." It's also the underlying reason `sort`-then-`filter` can beat `filter` directly on unsorted data in some benchmarks — sorting first makes the filter's branch predictable.

## Attempt

1. Draw a pipeline diagram (rows = instructions, columns = clock cycles, cells = stage) for these three instructions issued back-to-back in a 5-stage pipeline, with no hazard handling:
   ```
   add r1, r2, r3
   sub r4, r1, r5
   or  r6, r1, r7
   ```
   Mark exactly which cycle the data hazard on `r1` occurs (when `sub` needs `r1` in its execute stage, has `add` written it back yet?).

2. Redraw the same diagram assuming full forwarding (ALU-to-ALU forwarding available) and mark the cycle where the hazard is now resolved without stalling.

3. Write a benchmark (Go or C) with a branch inside a hot loop over an array: sort the array first, run the loop, time it; then run the identical loop on the unsorted array; time it. The comparison should isolate branch predictability — e.g., `if arr[i] > threshold { sum += arr[i] }` over sorted vs. unsorted data of the same values.

4. Record the timing difference and connect it explicitly to prediction accuracy: sorted data makes the branch outcome highly predictable (long runs of taken or not-taken), unsorted data makes it closer to random.

## Verify

Report actual wall-clock numbers from step 3/4 for both sorted and unsorted input over the same data (same values, just reordered) and the same array size, large enough that the effect isn't noise (typically need arrays in the hundreds of thousands to millions of elements, repeated over enough iterations to get a stable measurement).

## Failure drill

Replace the branch in your benchmark with a branchless equivalent (e.g., using a mask/multiply instead of `if`: `sum += arr[i] * (arr[i] > threshold ? 1 : 0)`, or the bitwise equivalent). Time it on the unsorted data. Explain why the branchless version's time should now be close to the sorted-branching version's time — it eliminated the misprediction cost entirely by eliminating the branch.

## Transfer

Find one real branch in a hot path of your own code (TARDOC's per-clinic processing loop, or Lead Sourcer's scoring logic are plausible candidates) and assess, without necessarily changing it, whether its branch condition is likely predictable (correlated with sorted/grouped data) or unpredictable (effectively random per iteration).

## Done when

You can draw a pipeline diagram showing a data hazard and its resolution via forwarding, and you have measured a real timing difference caused by branch predictability on your own machine, and you can explain in one sentence why a mispredicted branch costs more the deeper the pipeline.
