# Lesson 5: Async Runtime Fundamentals

## Objective

Understand Rust's async/await model — futures, executors, and wakeups — at a mechanistic level by building a minimal executor, and measure the concrete throughput difference between async and blocking I/O handling under concurrent load.

## Prerequisites

OS Lesson 2 (threads — async is an alternative concurrency model, best understood by contrast), Linux tools Lesson 5 / performance track Lesson 6 (I/O-bound vs. CPU-bound diagnosis — async specifically targets I/O-bound workloads).

## Learn

**Why async exists: the cost of one-OS-thread-per-connection at scale.** A traditional blocking-I/O server spawns one OS thread per concurrent connection (or uses a thread pool); each thread that's waiting on I/O (a slow network read, per computer architecture Lesson 7's latency numbers) is blocked, consuming a full OS thread's resources (stack memory, kernel scheduling overhead, OS Lesson 3) while doing nothing but waiting. At large connection counts (thousands or more), this becomes genuinely expensive — thousands of OS threads, most idle, waiting on I/O, is a real resource cost. Async I/O lets a single OS thread efficiently handle many concurrent, I/O-waiting tasks by *not* blocking that thread while waiting — instead, the task yields control back to a runtime, which can run other ready tasks, and resumes the original task only once its I/O is actually ready.

**A `Future` is a state machine, not a running computation.** Unlike a Go goroutine (which starts running immediately when spawned), a Rust `Future` does nothing on its own — it's a value representing a computation that *can* be driven forward by repeatedly calling `.poll()`, which either returns `Poll::Ready(value)` (done) or `Poll::Pending` (not done yet, will notify when progress is possible). This is a fundamentally different model from Go's goroutines: Rust's async is "cooperative" — nothing happens until something (an executor) actively polls the future forward.

**The executor: what actually drives futures to completion.** An executor (Tokio is the most common production one; this lesson has you build a minimal toy version) maintains a set of pending futures and repeatedly polls each one, using the wakeup mechanism (below) to know *when* re-polling a specific future is worthwhile, rather than busy-polling every future constantly (which would defeat the entire efficiency purpose).

**Wakeups: how a pending future signals "poll me again, I might be ready now."** When a future's `.poll()` returns `Pending` because it's waiting on something (a socket becoming readable, a timer firing), it registers a `Waker` with whatever it's waiting on — when that underlying event actually occurs (the OS reports the socket is now readable, via a mechanism like `epoll` on Linux), the `Waker` is invoked, which tells the executor "this specific future is worth polling again now," rather than the executor needing to blindly re-poll every pending future on every cycle.

**Backpressure in an async context, connecting to system-engineering Lesson 6.** An async system can still be overwhelmed if tasks are spawned faster than the executor can make progress on them — async doesn't eliminate the need for the bounded-queue backpressure principle from system-engineering Lesson 6, it just changes the specific resource being managed (pending futures/tasks rather than OS threads).

## Attempt

1. Write a simple async function using Tokio (`async fn fetch(url: &str) -> Result<...>`) that performs an HTTP request (networking Lesson 1-4's territory, now via an async HTTP client library), and run several of these concurrently using `tokio::join!` or spawning multiple tasks — confirm they genuinely execute concurrently (overlapping wall-clock time), not sequentially, by measuring total elapsed time against what sequential execution would take.

2. Build a minimal toy executor: implement a simple `Future` trait manually (not using `async`/`await` sugar, to force understanding the underlying `.poll()` mechanism directly) for something simple like a future that becomes ready after N calls to `.poll()` (simulating "not ready yet" without real I/O), and write a basic executor loop that polls a set of these futures repeatedly until all complete, tracking how many total poll calls were needed.

3. Instrument your toy executor from step 2 to log every poll call with a timestamp and the future's current state, and observe the actual poll pattern — confirm you can see the executor moving between futures (interleaving progress) rather than driving one to completion before starting the next, directly observing the cooperative multitasking model from Learn.

4. Compare async and blocking I/O handling under concurrent load directly: implement a small server handling many concurrent, artificially slow "requests" (simulated with an async sleep, standing in for slow I/O) once using Tokio's async model and once using a traditional one-thread-per-connection blocking model (OS Lesson 2's approach). Measure and compare actual resource usage (thread/task count, memory) and total throughput at a high concurrency level (e.g. 1000 simultaneous simulated slow requests) between the two approaches.

## Verify

Report your step 1 concurrent-vs-sequential timing comparison confirming genuine overlap, your step 2-3 toy executor's actual poll-count/timing log demonstrating cooperative interleaving, and your step 4 async-vs-blocking resource usage comparison at high concurrency, with real measured numbers for both approaches.

## Failure drill

Take your step 1 async code and deliberately perform a long-running, CPU-bound computation (not I/O) inside an `async fn` without yielding (no `.await` points during the computation) — run it alongside other concurrent async tasks and observe that the other tasks' progress stalls while your CPU-bound task runs, since it never yields control back to the executor. Explain why this demonstrates a genuine limitation of the cooperative model from Learn: async is specifically designed for I/O-bound concurrency (many tasks waiting, not computing), and a task that doesn't yield (either because it's genuinely CPU-bound with no I/O, or because it was written without inserting appropriate yield points) can starve every other task sharing the same executor thread — a fundamentally different failure mode from OS-thread-based concurrency, where the OS scheduler (OS Lesson 3) would preemptively interrupt a CPU-bound thread regardless of whether it "yields" voluntarily.

## Transfer

If TARDOC's transcription pipeline were rewritten to use async I/O for its external API calls (the Groq-hosted Whisper endpoint, per your project history) instead of blocking calls within Celery's process/thread-based worker model, describe, using this lesson's resource-usage comparison, what benefit you'd expect at your actual current concurrency scale — and honestly assess whether TARDOC's actual concurrent-request volume is large enough that this lesson's efficiency gains would be meaningfully noticeable, or whether the added complexity of an async rewrite wouldn't be justified at your current scale, connecting back to system-engineering Lesson 2's capacity-planning discipline of reasoning from real numbers rather than assuming a technique is beneficial regardless of actual scale.

## Done when

You've run genuinely concurrent async I/O and measured its overlap directly, you've built a minimal toy executor and observed its actual cooperative polling/interleaving behavior through your own instrumentation, you've measured a real resource-usage difference between async and thread-per-connection blocking models at high concurrency, and you've directly demonstrated — via the failure drill — async's specific limitation with non-yielding CPU-bound tasks, understanding this as a real scope boundary rather than a general flaw.
