# Lesson 1: C Syntax, Compilation, and Toolchain

## Objective

Understand what actually happens between writing C source and running a binary — preprocessing, compilation, assembly, linking — and be able to diagnose which stage a build failure came from instead of treating "compile error" as one undifferentiated thing.

## Prerequisites

None. This is the start of the C track. Lesson 2 of computer architecture (ISA and assembly) is useful background but not required — this lesson generates the assembly you'd read there.

## Learn

**The four stages, concretely.** `gcc -o prog main.c` looks like one step but is actually four, each independently inspectable:

1. **Preprocessing** (`gcc -E main.c`): expands `#include`, `#define` macros, and conditional `#ifdef` blocks into plain C text. No syntax checking happens here — a macro that expands to garbage won't error until the next stage.
2. **Compilation** (`gcc -S main.c`, producing `main.s`): translates preprocessed C into assembly for the target ISA. This is where syntax errors, type errors, and most warnings surface.
3. **Assembly** (`gcc -c main.c`, producing `main.o`): assembles the `.s` file into machine code, packaged as an object file — machine instructions plus a symbol table (names of functions/globals defined and referenced, not yet resolved to addresses).
4. **Linking** (`gcc main.o -o main`): combines one or more object files (and libraries) into a final executable, resolving symbol references between them — this is where "undefined reference to `foo`" errors come from, distinct from a compile error, because linking happens after each file compiled successfully on its own.

**Why separating these matters practically.** "It doesn't compile" is imprecise — a missing semicolon fails at compilation, a missing library fails at linking, and these need completely different fixes. Once you can tell which stage failed from the error message alone (compiler errors reference line numbers in your source; linker errors reference symbol names with no line number, because the linker has no concept of source lines, only object code), debugging build failures stops being guesswork.

**Warnings are not optional noise.** `-Wall -Wextra` catches real bugs at compile time that would otherwise become undefined behavior at runtime (Lesson 5) — an uninitialized variable, a signed/unsigned comparison, a mismatched printf format specifier. Treating warnings as advisory rather than as bugs-not-yet-triggered is one of the most common ways C programs ship with latent memory bugs.

## Attempt

1. Write a minimal C program that prints "hello" and includes one unused variable and one signed/unsigned comparison. Compile it first with no flags, then with `gcc -Wall -Wextra -o prog main.c`. Record exactly which warnings appear only with the flags enabled.

2. Run each stage separately and inspect the output at each step:
   ```
   gcc -E main.c -o main.i    # inspect main.i - see macro expansion, no code logic
   gcc -S main.i -o main.s    # inspect main.s - this is what CA Lesson 2 covered reading
   gcc -c main.s -o main.o    # inspect with: objdump -d main.o (or nm main.o for symbols)
   gcc main.o -o main         # final link
   ```
   For each stage, note in one sentence what changed from the previous stage's output.

3. Deliberately break the build at each stage and record the exact error message:
   - Preprocessing failure: `#include "nonexistent.h"`
   - Compilation failure: a syntax error (missing `}`) or type error (assigning a string literal to an `int`)
   - Linking failure: declare `void foo(void);` and call it, but never define `foo` anywhere

4. Write a two-file program (`main.c` calling a function defined in `helper.c`, with a shared `helper.h` header). Compile and link them as two separate object files (`gcc -c main.c`, `gcc -c helper.c`, `gcc main.o helper.o -o prog`), and confirm it links and runs correctly — this demonstrates the actual purpose of separate compilation: each file compiles independently, only linking needs all the pieces together.

## Verify

For step 3, produce a table: which stage failed, the exact error text, and which single word or phrase in the error message told you it was that stage (e.g. linker errors mention "undefined reference" with no source line number; compiler errors cite a specific file:line).

## Failure drill

Take the two-file program from step 4 and change `helper.h`'s function signature (e.g. change a parameter type) without recompiling `helper.c`, only recompiling and relinking against the stale `helper.o`. If your platform doesn't catch this at link time (C's linker generally does not check function signatures across translation units, only symbol names), run the resulting binary and observe either a crash, garbage output, or by-luck-correct behavior depending on ABI compatibility of the mismatched types. Explain in your own words why C's separate compilation model trusts the header declaration at compile time but cannot verify at link time that the actual object file honors it — this is a real, historically common class of bug in C projects with stale build artifacts.

## Transfer

Take a Go file you've already written (Go compiles to a single step from the outside, `go build`) and explain, based on this lesson, which of C's four stages Go's toolchain still performs internally even though it doesn't expose them as separate flags the way `gcc` does — specifically, does Go have an equivalent of preprocessing (no, no macro preprocessor), and does it have an equivalent of separate object-file linking (yes, internally, across packages).

## Done when

You can name which of the four stages a given build error came from purely from reading the error message, you've inspected the actual intermediate output of all four stages for one program, and you can explain why a stale, mismatched object file can link successfully but fail unpredictably at runtime.
