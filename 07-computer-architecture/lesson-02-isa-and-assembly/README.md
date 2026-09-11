# Lesson 2: Instruction Set Architecture and Assembly

## Objective

Read and write basic x86-64 (or ARM64) assembly, and explain what an ISA is: the contract between hardware and software.

## Prerequisites

Lesson 1 (digital logic) — an ALU is the hardware that executes the arithmetic instructions you'll read here.

## Learn

The ISA is the interface, not the implementation. It defines the instructions, registers, memory model, and calling convention a CPU exposes — different CPUs (an Intel chip, an AMD chip, an Apple M-series chip) implement the same x86-64 or ARM64 ISA completely differently in hardware while running the same compiled binary correctly. This is the same idea as an API contract in software: the caller doesn't need to know the implementation, only the interface.

**Registers.** A modern x86-64 CPU has a small number of general-purpose registers (`rax`, `rbx`, `rcx`, `rdx`, `rsi`, `rdi`, `rsp`, `rbp`, `r8`–`r15`) — tens of bytes total, versus gigabytes of RAM. Registers are fast because they're physically inside the CPU; every trip to memory costs orders of magnitude more cycles (Lesson 4 covers why). Good compiled code keeps hot values in registers as long as possible.

**Instruction categories.**
- Data movement: `mov dst, src`
- Arithmetic/logic: `add`, `sub`, `imul`, `and`, `or`, `xor`, `shl`/`shr`
- Control flow: `cmp` followed by conditional jumps (`je`, `jne`, `jl`, `jg`), unconditional `jmp`
- Stack: `push`, `pop`, `call`, `ret`

**The stack and calling convention.** `call` pushes the return address and jumps; `ret` pops it and jumps back. Function arguments in the System V AMD64 ABI go in `rdi`, `rsi`, `rdx`, `rcx`, `r8`, `r9` for the first six integer/pointer arguments, and the return value comes back in `rax`. This is not arbitrary — it's a convention every compiler agrees on so that code compiled separately (your program and libc) can call each other correctly. This is the hardware-level version of an API contract.

**Why a Go/C developer needs this:** when you read a stack trace, a segfault address, or `objdump` output while debugging, or reason about why a tight loop is slow, you're reading at this level. `go tool compile -S` and `gcc -S` both let you see the assembly your own code compiles to.

## Attempt

1. Write this C function, compile it with `gcc -O0 -S add.c` and read the generated assembly:
   ```c
   int add(int a, int b) {
       return a + b;
   }
   ```
   Identify which registers hold `a`, `b`, and the return value.

2. Compile the same function with `gcc -O2 -S` and diff against the `-O0` output. Note what the optimizer removed or changed (likely: no stack frame, direct register arithmetic).

3. Write a small function with a loop (sum 1 to N) in C, compile with `-O0 -S`, and manually trace the assembly for N=3 — write down the value of every register at each instruction until it returns.

4. If you have Go installed: `go build -gcflags="-S" main.go` on a small function, and identify the calling convention differences from C's ABI (Go historically used its own convention; note whether your Go version uses register-based ABI, which became default in Go 1.17+).

## Verify

Produce a hand-traced register table for step 3 (columns: instruction, `rax`, `rcx`, or whichever registers your compiler chose; one row per instruction) that matches the actual execution — confirm by running the compiled binary and checking the final printed result matches your traced value.

## Failure drill

Take the `-O0` assembly from step 1, manually flip the `add` instruction's operand order or change the immediate value, reassemble it by hand-editing (or note the change if you can't reassemble), and predict the wrong output before checking. This confirms you're reading the semantics of the instruction, not pattern-matching the syntax.

## Transfer

Take one hot function from a real Go or C project you've written (TARDOC, Mahall, or Lead Sourcer — anything with a tight loop or hot path), compile it with `-S`, and identify one place where the compiler either kept a value in a register across many instructions or spilled it to the stack. State why, based on how many live values were needed at once.

## Done when

You can read `-O0` assembly for a simple function and trace every register's value by hand, and you can explain the calling convention (which registers hold arguments, which holds the return value, what `call`/`ret` do to the stack) without looking it up.
