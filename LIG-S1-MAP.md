# LIG S1 — Subject Mapping to DART Tracks

FSEGS Sfax's Licence Informatique de Gestion (LIG), semestre 1, covers the following subjects. Rather than duplicating content in a separate, standalone folder, each subject is mapped to its existing (and, in most cases, considerably deeper) DART track. Study these tracks directly — this file is just the index.

| LIG S1 Subject | DART Track | Notes |
|---|---|---|
| Système d'exploitation | [`10-operating-systems`](10-operating-systems) | 8 lessons — processes, threads, scheduling, virtual memory, filesystems, IPC, synchronization, xv6 labs. Covers a real, working Linux/Unix model, not just OS theory. |
| Système logique et architecture des ordinateurs | [`07-computer-architecture`](07-computer-architecture) | 8 lessons — digital logic, ISA/assembly, CPU datapath, caches, pipelining, virtual memory, storage, capstone CPU simulator. |
| CCN (C) | [`04-c-programming-and-memory`](04-c-programming-and-memory) | 6 lessons — compilation toolchain, memory model, pointers/arrays/strings, allocation, undefined behavior, capstone library. |
| Algorithme et structure des données | [`05-problem-solving-and-algorithms`](05-problem-solving-and-algorithms) | 15 lessons — search, sorting, two pointers, sliding window, stacks/queues, recursion, trees, graphs, dynamic programming, capstone. |
| Math (analyse, algèbre, probabilités, statistiques) | [`08-mathematics`](08-mathematics) | 27 lessons across 6 sub-tracks — analysis, linear algebra, discrete math, probability, statistics, math-for-engineering. University depth, not a summary. |
| Base de données | [`09-postgresql-engineering`](09-postgresql-engineering) and [`13-database-internals`](13-database-internals) | Practical SQL/PostgreSQL engineering, plus how a DBMS actually works underneath (pages, buffer pools, B+ trees, WAL) — deeper than most S1 DB courses require, which is fine; nothing stops you from stopping at the practical-SQL level for exam purposes and treating `13-database-internals` as bonus depth. |
| Business communication | [`22-business-communication`](22-business-communication) | New track, 4 lessons — written messages, professional email, oral presentation, meetings/collaboration. |

## Not covered here

**Compta et principe de gestion** — explicitly out of scope per your instruction; not covered anywhere in this repository.

**Apprentissage et raisonnement** — the exact syllabus for this course could not be reliably confirmed. If you can get the actual course outline, exam, or lecture slides (even a photo), bring them and a dedicated track can be built against the real content rather than a guess.

## How to actually use this for exam prep

The DART tracks are built for depth and hands-on practice, not for exam-format review. For each subject above, the mapped track will get you well past what an exam requires — if you're short on time before a specific exam, prioritize:
- Reading the `Learn` sections of each lesson (the actual explanatory content) over doing every `Attempt`/`Failure drill` step.
- The earlier lessons in each track (foundational material more likely to be directly examined) over later capstone-style lessons (more likely to be practical/project-based, less likely to be exam content verbatim).
- Cross-checking against your actual course slides/TD sheets where you have them, since DART's scope and depth won't line up exactly with what a specific professor chooses to emphasize or test.
