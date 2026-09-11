# Lesson 3: CPU Scheduling

## Objective

Understand the tradeoffs between CPU scheduling policies by implementing and measuring several of them directly, connecting the abstract turnaround/response-time metrics to concrete, felt differences in workload behavior.

## Prerequisites

Lesson 1 (process states — scheduling is specifically about deciding which *ready* process runs next).

## Learn

**The scheduling problem.** With more ready processes than CPU cores, the OS must decide which runs next and for how long. Different policies optimize for different, sometimes conflicting goals: overall throughput, fairness between processes, minimizing time until a specific job finishes, or minimizing time until a process starts getting *any* CPU time at all (responsiveness) — a scheduler tuned for one of these goals typically performs worse on another.

**First-Come-First-Served (FCFS).** Processes run in arrival order, to completion, no preemption. Simple, but suffers from the **convoy effect**: one long process arriving first delays every short process behind it, even though servicing the short ones first would have finished more total work sooner (this is directly analogous to a slow customer at the front of a checkout line delaying everyone behind them, regardless of how quick their own transactions would be).

**Shortest Job First (SJF).** Always runs whichever ready process has the shortest remaining execution time. Provably optimal for minimizing *average* turnaround time (a proof-by-exchange-argument result, structurally similar to greedy-algorithm optimality proofs in the algorithms track), but requires knowing job lengths in advance — rarely true in real systems — and can starve long jobs indefinitely if short jobs keep arriving.

**Round Robin (RR).** Each ready process gets a fixed time slice (quantum); if it doesn't finish within that slice, it's preempted and moved to the back of the ready queue. Much better response time than FCFS (nothing waits arbitrarily long for its first slice of CPU), but turnaround time for any individual job can be worse than FCFS if the quantum is small relative to context-switch overhead — pure time-slicing has a real cost.

**Turnaround time vs. response time, the two metrics that actually matter.** Turnaround time = completion time − arrival time (how long a job took start-to-finish, including all waiting). Response time = time of first CPU access − arrival time (how long until the job *starts* getting attention at all, regardless of when it finishes). A batch job (e.g. an overnight report) cares about turnaround; an interactive terminal session cares about response time — this distinction is exactly why real OS schedulers (Linux's CFS, for instance) blend ideas from multiple policies rather than picking one pure strategy.

## Attempt

1. Given this workload (arrival time, burst/execution time in ms): `P1(0, 10), P2(1, 3), P3(2, 7), P4(3, 2)`, simulate FCFS by hand: compute the completion time, turnaround time, and waiting time for each process, and the average turnaround time across all four.

2. Simulate SJF (non-preemptive — once a process starts, it runs to completion, but among *ready* processes at each decision point, pick the shortest) for the same workload, and compute the same metrics. Compare the average turnaround time to FCFS's — SJF's should be lower or equal, per the optimality claim in Learn.

3. Simulate Round Robin with quantum = 2ms for the same workload (processes not finished within their quantum go to the back of the ready queue; new arrivals join the ready queue as they arrive). Compute completion, turnaround, and — critically — response time (first CPU access) for each process, and compare RR's response times to FCFS's and SJF's, both of which can make a late-arriving short process wait a long time before even starting.

4. Implement one of these three policies (your choice) as an actual small simulator in Go or Python: input a list of `(arrival, burst)` pairs, run the simulation loop, and output completion/turnaround/waiting/response times per process plus the averages. Verify your code's output against your step 1/2/3 hand calculation for the same policy.

## Verify

Produce a single table comparing all three policies on the same workload: average turnaround time and average response time for FCFS, SJF, and RR side by side. State explicitly which policy wins on which metric, and confirm this matches the theoretical tradeoff described in Learn (SJF should minimize average turnaround; RR should minimize the *worst-case* response time, i.e. no process waits arbitrarily long before its first slice).

## Failure drill

Construct a workload specifically designed to make SJF starve a process: have one long job arrive first, and then a continuous stream of short jobs arriving one after another for a long time, each shorter than the long job's remaining time. Simulate SJF on this workload (extend your step 4 simulator, or reason through it by hand for a representative case) and confirm the long job's turnaround time grows unboundedly worse the longer the stream of short jobs continues, even though the long job arrived first and, under FCFS, would have finished at a fixed, predictable time. Explain why this is not a simulator bug but SJF's genuine, provable weakness — optimal *average* turnaround time can come at the cost of individual-job starvation, which is exactly why real systems that use SJF-like priority ideas (e.g. Linux's CFS uses virtual runtime accounting partly to avoid exactly this failure mode) build in explicit fairness mechanisms rather than using pure SJF.

## Transfer

Linux's actual scheduler (Completely Fair Scheduler, CFS, in kernels before the more recent EEVDF scheduler — note explicitly which your platform actually runs if you want to check, since this has changed in recent kernel versions) doesn't implement any of FCFS/SJF/RR directly, but blends ideas: it tracks each process's "virtual runtime" (accumulated CPU time, weighted by priority) and always favors the process with the least virtual runtime so far — approximating fairness over time in a way structurally closer to RR's fairness goal than SJF's throughput goal, while avoiding RR's fixed-quantum context-switch overhead by using adaptive time slices. State, using your own workload simulation results, why a batch-processing server (like a Celery worker pool processing TARDOC transcription jobs) and an interactive terminal session would reasonably want *different* scheduling behavior, and which of this lesson's three metrics (turnaround vs. response time) each workload type should be optimized for.

## Done when

You've hand-computed and then code-verified turnaround/response times for all three policies on the same workload, you've directly demonstrated SJF's starvation failure mode with a constructed adversarial workload, and you can explain — using your own comparison table — why no single one of these three simple policies dominates the others on both metrics simultaneously.
