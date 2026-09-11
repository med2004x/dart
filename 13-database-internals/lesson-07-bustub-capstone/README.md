# Lesson 7: Capstone — Implement a DBMS Component

## Objective

Integrate Lessons 1-6 into one substantial, tested database component with documented invariants and measured performance — the database-internals track's equivalent of the integration capstones in earlier tracks.

## Prerequisites

Lessons 1-6, completed. No new theory — this lesson exists to force the individually-understood pieces (pages, buffer pool, B+ trees, query execution, transactions, WAL) to work together under real constraints, which is exactly where gaps in understanding tend to surface.

## Learn

There is no new material. CMU's BusTub (referenced throughout this track) is a real, substantial teaching database system used in CMU's actual database systems course, with a well-defined project structure and automated tests — using it directly, rather than reconstructing an equivalent from scratch, is a strong and encouraged option, consistent with this curriculum's general approach of preferring rigorous existing coursework where it exists.

The engineering discipline this capstone specifically tests: can you take components built and tested in isolation (your Lesson 1 page format, Lesson 2 buffer pool, Lesson 3 B+ tree, Lesson 4 execution engine, Lesson 6 WAL) and make them work correctly *together*, where each component's assumptions about the others must actually hold in practice, not just in each lesson's isolated test suite.

## Attempt

If using BusTub directly, complete at least the buffer pool manager project and the B+ tree index project from the official course materials (`https://github.com/cmu-db/bustub`), using its existing test infrastructure to verify correctness — this is the recommended default path, since BusTub's tests are considerably more rigorous than what's practical to construct from scratch in this lesson alone.

If building your own integrated system instead (a reasonable alternative if you want the components to be ones you fully wrote yourself, end to end), integrate your own Lessons 1-6 implementations into one coherent system with this minimum scope:

1. **Storage layer**: your Lesson 1 page format, with records actually read/written through your Lesson 2 buffer pool (not bypassing it) — every page access in the rest of the system should go through `fetchPage`/`unpinPage`, with correct pin discipline throughout.

2. **Indexing**: your Lesson 3 B+ tree, storing pointers to record locations (page ID + slot number, per Lesson 1's slotted-page design) rather than full record data directly in the tree — this is a meaningful integration detail: the index and the actual table storage are separate, connected via a pointer, exactly as in a real database.

3. **Query execution**: your Lesson 4 iterator-based operators (`SeqScan`, `IndexScan`, `Filter`, `Project`), with `IndexScan` genuinely using your integrated B+ tree from item 2, and `SeqScan` genuinely reading pages through your buffer pool from item 1 — not separate, disconnected implementations from each earlier lesson's isolated exercises.

4. **Durability**: every mutating operation (insert, update, delete) writes a WAL record (Lesson 6) before modifying the buffer pool's in-memory page state, with a working `recover()` path that can reconstruct correct state after a simulated crash of the *entire integrated system*, not just the isolated WAL exercise from Lesson 6.

Whichever path you choose (BusTub or your own integration), the requirements are:

- **Write focused tests for cross-component invariants** specifically — e.g. "after inserting N records and building an index on them, every indexed lookup returns a record whose actual page/slot location, when read through the buffer pool, matches what a full sequential scan would find for the same key" — a test that only passes if the storage, buffer pool, and index components are all correctly wired together, not testable by any single component's isolated unit tests.
- **Profile and document at least one real bottleneck.** Run a workload (e.g. inserting a large number of records, then performing many indexed lookups) and profile it (Linux tools Lesson 5 / computer architecture Lesson 4's tooling) to identify where time is actually going — report the actual measured bottleneck, which component it's in, and whether it matches or contradicts your intuition before measuring.

## Verify

Report your cross-component invariant test results (should pass consistently across multiple runs and multiple random test data seeds if your test generates random data), and your actual profiling output identifying a real bottleneck with supporting numbers, not a guess.

## Failure drill

Take your integrated system (or the BusTub project you completed) and deliberately break one cross-component assumption — for example, if you built your own system, bypass the buffer pool's pin discipline in exactly one code path (directly access a page without pinning it first) while the rest of the system correctly uses pins. Run your workload under concurrent access (multiple goroutines/threads performing operations simultaneously) and attempt to reproduce a corruption or crash caused specifically by this one broken assumption — e.g. a page being evicted while the unprotected code path is still reading it, per Lesson 2's pin/unpin discussion. If you cannot reliably reproduce a visible failure (this specific race can be timing-dependent and hard to force), document your reasoning for why the assumption violation is still a genuine bug even if your specific test run didn't happen to trigger a visible symptom — directly echoing C track Lesson 5's core point about undefined behavior not requiring a visible crash to be a real bug.

## Transfer

Compare one specific design decision in your integrated system (or in the BusTub project, if you used it) to PostgreSQL's actual corresponding mechanism — for instance, compare your B+ tree's specific node-splitting or branching-factor choice to PostgreSQL's actual B-tree index implementation, or your WAL's minimal record format to PostgreSQL's actual WAL record types — and identify at least one specific way PostgreSQL's real implementation is more sophisticated, and briefly explain what real-world requirement (concurrent access at much larger scale, more complex recovery scenarios, vacuum/vacuuming considerations) likely drove that additional complexity that your simplified educational version didn't need to handle.

## Done when

You've completed a genuinely integrated system (or the equivalent BusTub projects) where storage, buffer pool, indexing, execution, and durability components actually depend on and correctly use each other — not separate, disconnected pieces — your cross-component invariant tests pass consistently, and you've profiled a real bottleneck and connected at least one design decision in your system to the corresponding, more sophisticated real-world mechanism in PostgreSQL.
