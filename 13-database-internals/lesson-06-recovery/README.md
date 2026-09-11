# Lesson 6: WAL and Crash Recovery

## Objective

Implement a minimal write-ahead log and crash-recovery replay mechanism, understanding precisely why "write the log first, the data second" is the rule that makes durability possible after an unexpected crash.

## Prerequisites

Lesson 2 (buffer pool — dirty pages and their flush timing are exactly what WAL ordering constrains), OS Lesson 5 (filesystem crash consistency — this lesson is the database-engine-level version of that same underlying problem, with stronger guarantees than a general filesystem journal alone provides).

## Learn

**The core problem, restated precisely.** A transaction's changes are first made to in-memory pages (in the buffer pool, Lesson 2). Those pages are only periodically flushed to disk — not after every single change, since that would be prohibitively slow (computer architecture Lesson 7's I/O latency numbers apply directly here). If the process crashes (or the machine loses power) before a dirty page is flushed, that in-memory change is gone — yet the transaction may have already told the client "committed," meaning the durability guarantee (a committed transaction's effects survive a crash) has been violated unless something else preserved the change.

**Write-ahead logging: the fix.** Before making a change to a page in the buffer pool, the database first writes a log record describing that change to a separate, append-only log file, and — critically — ensures that log record is durably on disk (an `fsync` or equivalent, forcing the OS to actually write it to physical storage, not just hand it to a buffer that might itself be lost on crash) *before* the transaction is allowed to report "committed" to the client. The actual data page can be flushed to its normal location later, whenever convenient — because if a crash happens before that flush, the durable log record is enough to reconstruct the change during recovery. This is the literal meaning of "write-ahead": the log entry must be durable *ahead of* both the commit acknowledgment and (though this ordering is slightly more nuanced across specific implementations) the corresponding data page's flush.

**Recovery: replaying the log after a crash.** On restart after a crash, the database reads the WAL from the last known-consistent point and replays every logged change, reconstructing whatever in-memory state was lost — this works because the log contains enough information (typically the specific bytes/values changed, or a description sufficient to redo the operation) to deterministically reproduce every committed change, even though the actual data pages on disk may have been left in an inconsistent, partially-updated state at crash time.

**Checkpointing.** Replaying the *entire* log from the very beginning of the database's history on every restart would become prohibitively slow as the log grows. A checkpoint periodically flushes all currently-dirty pages to disk and records a marker in the log noting "everything before this point is now safely reflected in the data files" — recovery after a crash only needs to replay from the most recent checkpoint forward, not from the beginning of time, bounding recovery time to roughly "how much happened since the last checkpoint" rather than "how much has ever happened."

## Attempt

1. Design and implement a minimal WAL format: an append-only log of records, each describing a single change (e.g. `{transactionID, pageID, offset, oldValue, newValue}` — enough to both redo the change during recovery, and, for a fuller implementation, undo an uncommitted transaction's changes, though this lesson's minimum scope only requires redo).

2. Implement a simple in-memory "database" (a map or small set of pages, standing in for your Lesson 1/2 page/buffer-pool work if you want to integrate directly, or a simpler in-memory structure if you're keeping this lesson self-contained) where every mutation first appends a WAL record (and, for this exercise, immediately fsyncs/flushes it to a real file on disk to make the durability property genuine, not just simulated in memory) before applying the change to the in-memory state.

3. Simulate a crash: run a sequence of transactions/changes against your system, then — without a clean shutdown — deliberately discard your in-memory state (simulate the process dying) while your WAL file remains on disk. Implement a `recover()` function that reads the WAL file from the start (or from the last checkpoint, if you implement step 4) and replays every committed change, reconstructing the in-memory state as it should have been at crash time.

4. Implement basic checkpointing: after some number of transactions (or on an explicit trigger), write a checkpoint marker to the WAL noting the current state has been fully reflected elsewhere (for this simplified exercise, you can simply also serialize the full in-memory state to a separate "checkpoint" file at that point). Modify `recover()` to start from the most recent checkpoint (loading that saved state) and replay only the WAL records that came after it, rather than replaying the entire log from the beginning every time.

## Verify

For step 3, run a test sequence with at least 20 changes, simulate the crash, and confirm your `recover()` function reconstructs a final state that exactly matches what the state would have been had no crash occurred (compare against a reference in-memory execution with no simulated crash, run separately, as your ground truth).

## Failure drill

Deliberately violate the write-ahead ordering rule: modify your system to apply a change to in-memory state *first*, and only append the WAL record *afterward* (reversing the correct order), and additionally simulate a crash occurring in the narrow window between the in-memory change and the WAL write actually completing (e.g. by having your test harness apply the in-memory change, then deliberately skip the WAL write for that specific change before simulating the crash). Confirm that `recover()` — which only knows about what's actually in the WAL — reconstructs a state missing that specific change, even though (in the pre-crash, still-running system) the change had already taken visible effect in memory and might even have been reported to a caller as complete. Explain, using this concrete broken scenario, exactly why "write-ahead" (log before commit acknowledgment, and log before/no-later-than the corresponding page flush) is not an arbitrary convention but the precise ordering requirement that recovery's correctness depends on — reversing it reopens exactly the durability gap WAL exists to close.

## Transfer

PostgreSQL's own WAL (which TARDOC's database relies on for durability) works on precisely this principle, and PostgreSQL's `synchronous_commit` setting controls exactly how aggressively it enforces the "WAL durable before commit acknowledged" ordering (weaker settings can improve throughput by acknowledging commits slightly before the WAL write is fully guaranteed durable, trading a small durability window for performance). Look up (via PostgreSQL documentation) what `synchronous_commit = off` specifically risks losing in the event of a crash, and state, using this lesson's WAL-ordering discussion, why that setting is precisely a controlled, deliberate relaxation of the write-ahead guarantee this lesson's failure drill demonstrated the danger of breaking accidentally.

## Done when

Your WAL-based system correctly recovers a matching final state after a simulated crash across a real, nontrivial sequence of changes, your checkpointing correctly bounds replay to only the log entries since the last checkpoint, and you've deliberately broken the write-ahead ordering and directly observed data loss as a result, connecting it explicitly to why the ordering rule — not just "have a log" — is what actually provides the durability guarantee.
