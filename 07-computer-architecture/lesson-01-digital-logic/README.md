# Lesson 1: Digital Logic

## Objective

Explain how logic gates combine into circuits that compute, store, and control state. Build the primitives from NAND upward: gates, half/full adders, multiplexers, and a 1-bit latch.

## Prerequisites

None. This is the base of the track.

## Learn

Every digital circuit reduces to switches. A transistor acts as a gate: current flows or it doesn't. From that single fact, all logic follows.

**Gates.** NOT, AND, OR are the primitives you're taught first, but NAND (NOT-AND) is the practical primitive: every other gate can be built from NAND alone. This matters because real chips are fabricated from a small standard cell library, and NAND is cheap to fabricate.

```
NAND truth table:
A B | Out
0 0 | 1
0 1 | 1
1 0 | 1
1 1 | 0
```

From NAND: `NOT(A) = NAND(A, A)`. `AND(A,B) = NOT(NAND(A,B))`. `OR(A,B) = NAND(NOT(A), NOT(B))`.

**Adders.** A half adder takes two 1-bit inputs and produces a sum bit and a carry bit: `Sum = A XOR B`, `Carry = A AND B`. A full adder adds a carry-in as well, which is what lets you chain N of them into an N-bit ripple-carry adder — the circuit that computes `A + B` in a real CPU's ALU.

**Multiplexers.** A 2-to-1 mux selects one of two inputs based on a select line: `Out = (S AND A) OR (NOT(S) AND B)`. Muxes are how a CPU picks between "use the ALU result" and "use the memory value" when writing back to a register — control logic is built almost entirely from muxes and gates.

**Latches and flip-flops.** Combinational logic (gates, adders, muxes) has no memory — output depends only on current input. To store a bit, you feed a gate's output back into its own input. An SR latch built from two cross-coupled NOR gates holds state until told to change. A D flip-flop samples its input on a clock edge, which is the basis for registers: a CPU's registers are just banks of flip-flops.

**Why this matters for you as a backend engineer:** you will never design a chip. The reason to know this is that every abstraction above it — instructions, cache lines, memory ordering — is built on this substrate and inherits its constraints. Understanding "a flip-flop only changes state on a clock edge" is what makes clock speed, pipelining, and race conditions in hardware make sense later.

## Attempt

Using [nand2tetris](https://www.nand2tetris.org/) project 1 (or a logic simulator of your choice — a truth table on paper is enough if you have no simulator):

1. Build NOT, AND, OR, XOR from NAND only. Write out each truth table and confirm by hand.
2. Build a half adder and a full adder from your gates above. Verify with a truth table for all 8 input combinations of the full adder (A, B, Cin).
3. Chain four full adders into a 4-bit ripple-carry adder. Compute `0110 + 0101` by hand tracing every carry, then confirm against your circuit.
4. Build a 2-to-1 mux and a 4-to-1 mux (the 4-to-1 uses two select lines and three 2-to-1 muxes, or the equivalent gate expression).
5. Describe in your own words (no circuit needed) why an SR latch made of two NOR gates can hold a value with no clock — trace what happens when S and R are both 0 after S was pulsed to 1.

## Verify

For the 4-bit adder, produce a table of at least 5 addition pairs including one that overflows 4 bits (e.g. `1111 + 0001`) and show the resulting carry-out. State explicitly what the carry-out bit means for a real CPU (it becomes the overflow/carry flag).

## Failure drill

Chain the full adders without connecting Cin from one stage to Cout of the previous stage — leave each Cin tied to 0. Compute `0110 + 0101` again with this broken wiring. Explain exactly which additions now produce a wrong 4-bit result and why, and connect this to what "ripple carry" actually means — the propagation delay per stage is the source of the ripple-carry adder's speed limit, which is why real CPUs use carry-lookahead adders instead.

## Transfer

Design (on paper) a 1-bit ALU that takes A, B, and a 2-bit operation select, and outputs AND, OR, or SUM based on the select — this is the basic cell that gets replicated N times to build an N-bit ALU, which is where Lesson 3 picks up.

## Done when

You can explain, without notes, why NAND alone is sufficient to build any logic function, trace a 4-bit ripple-carry addition including every intermediate carry bit, and state why an unclocked latch can still hold state.
