# Lesson 8: Capstone — Build a Toy CPU

## Objective

Integrate lessons 1-7 by implementing a small CPU simulator with a real instruction set, memory hierarchy simulation, and measurable performance characteristics — in software, since building physical hardware is out of scope.

## Prerequisites

Lessons 1-7, completed. This lesson does not introduce new theory; it forces you to use all of it together.

## Learn

There is no new material here. The point of a capstone is combining separately-learned pieces into one working system, which is where gaps in understanding actually surface — you can know what a cache is and what a pipeline is in isolation and still fail to reason correctly about how they interact (e.g., a cache miss stalling a pipeline stage). This lesson is deliberately closer to a systems-engineering task than a theory lesson.

If nand2tetris is available to you, its projects 1-6 (gates through a full computer, culminating in an assembler) are the canonical version of this exercise and are strongly preferable to building from scratch if you have the time — use it as the reference and adapt the attempt below to whichever pieces it doesn't already cover.

## Attempt

Build a simulator (Go or C) for a small instruction set. Minimum required instructions:

```
LOAD  rd, addr       ; rd = mem[addr]
STORE rs, addr       ; mem[addr] = rs
ADD   rd, rs1, rs2   ; rd = rs1 + rs2
SUB   rd, rs1, rs2   ; rd = rs1 - rs2
BEQ   rs1, rs2, off  ; if rs1 == rs2, PC += off
JMP   addr           ; PC = addr
HALT                 ; stop execution
```

Required components:

1. A register file (array of at least 8 integer registers).
2. A flat memory array (at least 4KB, addressed by byte or word — your choice, document it).
3. A fetch-decode-execute loop (Lesson 3) that reads an instruction, decodes its fields, executes it, and updates PC — correctly handling the PC increment vs. branch/jump override case from Lesson 3's failure drill.
4. A simulated memory access cost: track a cycle counter, and charge different costs for register operations (1 cycle) versus memory operations (model a simple 2-level hierarchy: a small direct-mapped cache, e.g. 16 lines of 4 words each, with a hit costing ~4 cycles and a miss costing ~100 cycles, per Lesson 4's real numbers scaled down).
5. A simple branch predictor: predict "not taken" always (simplest baseline), track prediction accuracy across a test program, and charge a flush penalty (e.g. 3 cycles) on misprediction, per Lesson 5.

Write at least two test programs by hand in your instruction set:
- One that sums an array (LOAD in a loop, ADD, BEQ to loop) — exercises memory access patterns and branching.
- One that has no memory access at all, pure register arithmetic — should run with minimal cycle cost and zero cache activity.

## Verify

Run both test programs and report: final register/memory state (correctness — does it compute the right sum?), total cycle count, cache hit/miss count, and branch prediction accuracy. The register-only program should show zero cache accesses and the minimum possible cycle count; the array-sum program should show cache activity that you can reason about (does it hit mostly, given sequential access, per Lesson 4's spatial locality argument?).

## Failure drill

Take your array-summing test program and change its memory access pattern from sequential (stride 1) to a large stride that defeats your simulated cache's spatial locality, matching Lesson 4's stride experiment. Confirm your simulator's cache hit rate drops and total cycle count rises, and that this matches the direction (if not exact magnitude) of your real-hardware measurement from Lesson 4.

## Transfer

Extend the simulator with one instruction type not in the minimum list (e.g. a multiply, a conditional move, or an indirect jump through a register) and integrate it correctly into fetch-decode-execute, including deciding its cycle cost and whether it touches the cache model.

## Done when

Your simulator correctly executes both hand-written test programs with correct final state, reports believable cycle/cache/branch statistics that respond in the right direction to the failure drill's changed access pattern, and you can explain, using your own simulator as the concrete example, how all of lessons 1-5 combine into one running machine.
