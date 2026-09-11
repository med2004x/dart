# Lesson 5: Consensus with Raft

## Objective

Understand and implement core components of the Raft consensus algorithm — leader election and log replication — the rigorous protocol that solves Lesson 4's "how do you correctly elect a new leader after the old one fails" problem, which that lesson deliberately left unaddressed.

## Prerequisites

Lesson 4 (replication — Raft is specifically the algorithm providing correct, safe leader election and log replication for exactly that lesson's replicated state machine pattern), Lesson 1 (failure models — Raft is designed to tolerate exactly the failure modes covered there).

## Learn

**Why leader election is genuinely hard, not just "pick whoever notices the leader is gone first."** Multiple followers might simultaneously notice the leader seems unresponsive and each try to become the new leader — if this isn't handled carefully, you can end up with two nodes both believing they're the leader simultaneously (a "split brain" scenario), each accepting writes independently, leading to genuinely divergent, conflicting state — precisely the kind of correctness failure Lesson 4's single-leader design was meant to avoid in the first place. Raft's entire election mechanism exists to guarantee at most one leader can be legitimately elected per **term** (a logical, incrementing "epoch" number), even under network partitions and message loss.

**Terms and elections, mechanically.** Time is divided into terms, each with at most one leader. A follower that hasn't heard from a leader within a timeout becomes a candidate, increments the term number, and requests votes from other nodes. A node grants its vote to at most one candidate per term (first-come, first-served within that term) — this "at most one vote per term per node" rule is the key safety mechanism: a candidate needs votes from a **majority** of nodes to become leader, and since any two majorities out of the same total node set must overlap by at least one node, it's structurally impossible for two different candidates to both win a majority in the *same* term — at least one node's single vote would have to go to both, which the rule forbids.

**Log replication, and why "majority" reappears here too.** The leader appends new operations to its own log and replicates them to followers (Lesson 4's mechanism); it only considers an entry "committed" (safe to apply and report success to the client) once a *majority* of nodes (including itself) have durably stored it — this majority requirement is what guarantees a committed entry survives even if the leader immediately crashes afterward: any newly-elected leader (which itself required a majority vote to win) must, by the pigeonhole-style overlap argument, share at least one node with any majority that stored a given committed entry, and Raft's election rules specifically ensure a candidate can only win if its log is at least as up-to-date as that majority — guaranteeing the new leader definitely has every previously committed entry.

**Why this connects directly to discrete math's combinatorics/pigeonhole lesson.** The "any two majorities of the same set overlap" property is a direct pigeonhole argument (discrete math track, combinatorics lesson) — with N total nodes, two subsets each larger than N/2 cannot be disjoint, since their combined size would exceed N. This single combinatorial fact is the mathematical foundation underneath Raft's entire safety guarantee — worth recognizing the connection explicitly rather than treating Raft's correctness as separate, unrelated magic.

## Attempt

MIT's 6.5840 (referenced throughout this track) provides a rigorous, well-tested lab sequence for implementing Raft incrementally, with a real test harness that specifically injects network partitions, delays, and crashes — using it directly is strongly preferred over reconstructing an equivalent test harness from scratch, since correctly testing a consensus implementation under adversarial conditions is itself a nontrivial engineering task the course has already solved well.

1. Implement Raft leader election (the course's first Raft lab): nodes start as followers, time out and become candidates, request votes, and a majority-vote winner becomes leader for that term. Test it under the course's provided test harness, including scenarios where the current leader is forcibly disconnected and a new election must correctly occur.

2. Implement log replication on top of your working leader election: the leader accepts client operations, appends to its log, replicates to followers, and only reports an operation as committed once a majority of nodes have it durably stored (per Learn). Test correctness under the course's provided tests, including scenarios with simulated network partitions during replication.

3. Trace, by hand, a concrete scenario using a 5-node cluster: node A is leader in term 3, has replicated log entries to nodes B and C (a majority: A, B, C = 3 out of 5) but not yet to D and E, then A crashes. Reason through which nodes are eligible to win the subsequent election (per the "candidate's log must be at least as up-to-date" rule) and confirm that whichever node wins, it necessarily has the entry that was already committed via the A/B/C majority — write out explicitly why D or E winning the election without that entry would be prevented by Raft's voting rules.

4. Deliberately construct (using the test harness's partition-injection capability) a scenario with a network partition splitting a 5-node cluster into a group of 3 and a group of 2. Confirm the group of 3 (still a majority) can elect a leader and continue making progress, while the group of 2 (a minority) cannot — directly observing the majority requirement's real consequence: a minority partition is correctly unable to elect a leader or commit new entries, preventing a split-brain scenario even though those 2 nodes can still communicate with each other.

## Verify

For step 4, report the actual test harness output confirming the majority partition successfully elected a leader and committed new entries while the minority partition did not, and reconnect the partition afterward, confirming the previously-minority nodes correctly catch up to the majority's state once communication is restored.

## Failure drill

Attempt to deliberately break the "one vote per term per node" rule in your implementation (allow a node to vote for two different candidates within the same term) and rerun your step 1 election tests under network conditions likely to produce contested elections (e.g. multiple nodes becoming candidates near-simultaneously). Attempt to reproduce a split-brain scenario — two nodes each believing they won the same term's election. Explain, using the majority-overlap pigeonhole argument from Learn, precisely which specific guarantee your broken version violated, and why restoring the one-vote-per-term rule is what makes the overlap argument (and therefore the whole safety guarantee) hold again.

## Transfer

Consensus systems like etcd, Consul, and CockroachDB all use Raft (or a close variant) internally for exactly this leader-election/log-replication purpose. If you were designing a coordination mechanism for a hypothetical multi-instance deployment of TARDOC or Mahall (e.g. multiple API server instances needing to agree on which one currently "owns" a specific background job, to avoid the exact overlapping-Celery-task risk mentioned in distributed Lesson 1's transfer task), describe at a conceptual level why using an existing, battle-tested Raft-based system (etcd, for instance) for that specific coordination need would be strongly preferable to attempting to hand-roll an equivalent leader-election mechanism yourself, given everything this lesson demonstrated about how subtle correctly implementing consensus actually is.

## Done when

You've implemented and passed the test harness's leader election and log replication tests, including under simulated partitions and crashes, you've hand-traced a concrete 5-node scenario confirming why a newly-elected leader necessarily has all previously-committed entries, and you've directly observed a minority partition correctly failing to make progress while a majority partition succeeds — connecting this observed behavior explicitly to the pigeonhole-based majority-overlap argument underlying Raft's safety guarantee.
