# Lesson 3: CPU Datapath

## Objective

Explain the fetch-decode-execute cycle and trace how a single instruction moves through a CPU's datapath.

## Prerequisites

Lesson 1 (gates, adders, muxes — the building blocks), Lesson 2 (instructions and registers — what's being executed).

## Learn

A CPU is a state machine that repeats one loop: **fetch, decode, execute, (memory), write-back.**

1. **Fetch.** The Program Counter (PC) register holds the address of the next instruction. The CPU reads memory at that address, gets the instruction bits, and increments the PC (usually by 4 bytes for fixed-width ISAs like ARM, or a variable amount for x86's variable-length encoding).

2. **Decode.** The instruction bits are split into fields: opcode (what operation), source registers, destination register, immediate value if any. Control logic (built from the muxes and gates of Lesson 1) reads the opcode and sets signals that configure the rest of the datapath for this specific instruction.

3. **Execute.** The ALU (the 1-bit ALU from Lesson 1's transfer task, replicated N times) performs the arithmetic or logic operation on the register values read in decode.

4. **Memory.** If the instruction is a load or store, the ALU result (an address) is used to read or write memory. Non-memory instructions skip this stage.

5. **Write-back.** The result (from the ALU or from memory) is written into the destination register.

**Single-cycle vs multi-cycle.** A naive design does all five stages in one long clock cycle, sized to fit the slowest instruction (a memory access). This wastes time on every fast instruction (like a register-to-register add) because the clock still has to wait for the slowest path. A multi-cycle design breaks this into separate shorter cycles per stage, but only one instruction is "in flight" at a time. Lesson 5 covers pipelining, which overlaps multiple instructions' stages to get both speed benefits at once.

**Why this matters:** when you reason about "how many cycles does this loop take," or read a comment about branch misprediction cost, you're reasoning about this datapath. It's also the mental model behind why some operations (division, memory access) are inherently slower than others (register add) — they involve more or slower stages.

## Attempt

1. Draw (on paper or in a text diagram) the datapath for a single `add rd, rs1, rs2` instruction (RISC-style: register-register add), labeling each of the 5 stages and which registers/wires carry data between them.

2. Trace the same diagram for a `load rd, offset(rs1)` instruction (memory load). Identify the extra step this instruction needs that the register-add did not (the ALU computes an address, not a result).

3. Using a MIPS or RISC-V single-cycle datapath diagram (search for "RISC-V single cycle datapath" — this is standard course material, e.g. from Patterson & Hennessy or nand2tetris's Hack CPU), hand-trace the control signals for one `add` and one `beq` (branch-if-equal) instruction. Write down which control signals differ between the two.

4. Implement the fetch-decode-execute loop as pseudocode or actual code (Go or C) for a toy 4-instruction ISA of your own design (e.g. `LOAD`, `ADD`, `STORE`, `JMP`), operating on a simulated register file (an array) and simulated memory (another array). This is the software version of the hardware loop — it should produce correct results for a 5-10 instruction test program you write by hand.

## Verify

Run your toy CPU simulator (from step 4) on a hand-written program that computes something checkable — e.g., sum the first 5 integers stored in memory into a register — and confirm the register file's final state matches a hand-computed expected result.

## Failure drill

In your toy simulator, forget to increment the PC after a non-jump instruction (leave it pointing at the same instruction). Run it and observe the infinite loop. Explain in one sentence why every real CPU's fetch stage must update PC unconditionally except when a control-flow instruction explicitly overrides it.

## Transfer

Take nand2tetris's Hack CPU specification (project 5) and identify the equivalent of each of the 5 stages above in its design — Hack collapses some stages compared to a textbook RISC datapath; identify which ones and why that's possible for a simplified teaching ISA.

## Done when

You can draw the fetch-decode-execute-memory-writeback datapath from memory, explain what data moves on each stage for both an arithmetic and a memory instruction, and your toy CPU simulator correctly executes a hand-written 5+ instruction program.
