# Lesson 2: TCP Fundamentals

## Objective

Understand precisely what TCP guarantees (reliable, ordered, byte-stream delivery) and the mechanisms — the three-way handshake, retransmission, windowing — that provide those guarantees over an unreliable network layer beneath it.

## Prerequisites

Lesson 1 (sockets — this lesson explains the mechanism behind the byte-stream abstraction that lesson used practically).

## Learn

**What TCP actually promises, precisely.** Below TCP sits IP, which provides no reliability guarantees at all — packets ("datagrams") can be lost, duplicated, or arrive out of order, with no notification to either endpoint. TCP is built on top of this unreliable layer to provide: **reliable delivery** (lost data is detected and retransmitted), **ordered delivery** (even if IP packets arrive out of order, TCP reassembles them into the correct byte order before handing data to the application), and **flow control** (matching the sender's rate to what the receiver can actually process, distinct from — though related in spirit to — the pipe backpressure from OS Lesson 6).

**The three-way handshake.** Establishing a TCP connection requires: client sends `SYN` (synchronize — "I want to start a connection, here's my initial sequence number"), server responds `SYN-ACK` (acknowledging the client's SYN and sending its own), client responds `ACK` (acknowledging the server's SYN). Only after this exchange completes do both sides consider the connection established. This is why establishing a new TCP connection has an inherent minimum latency cost of roughly one round-trip time (RTT) before any application data can even be sent — a real, measurable cost that's part of why connection reuse (Lesson 4's HTTP keep-alive) matters for performance.

**Sequence numbers and acknowledgment.** Every byte sent over TCP is logically numbered; the receiver acknowledges (`ACK`s) the sequence number up to which it has successfully received data, in order. If the sender doesn't receive an ACK for data within some timeout, it assumes loss and retransmits — this is the actual mechanism providing the reliability guarantee, built entirely from "number everything, acknowledge what arrived, resend what wasn't acknowledged in time."

**The sliding window.** Rather than sending one segment and waiting for its ACK before sending the next (which would badly under-utilize available bandwidth on any connection with meaningful RTT), TCP allows multiple unacknowledged segments to be "in flight" at once, up to a window size — this window is influenced by both the receiver's advertised capacity (flow control — how much the receiver can currently buffer) and congestion control (a separate mechanism estimating how much the *network path* can currently sustain without excessive loss, which adjusts dynamically based on observed loss/delay). This is the direct network-layer analog of Lesson 1 (math-for-engineering)'s queueing theory: too much in flight relative to what the path can sustain leads to the same nonlinear congestion/latency blowup that queueing theory predicts for any system pushed near saturation.

**Connection teardown.** A four-way close (`FIN`/`ACK` from each side, since each direction of the bidirectional stream must be closed independently) — this is why a TCP connection can be in a "half-closed" state where one side has finished sending but can still receive, a detail that matters for correctly implementing protocols that need to signal "I'm done sending" without immediately terminating the whole connection.

## Attempt

1. Using a packet capture tool (Lesson 7 covers this properly; `tcpdump` or Wireshark is fine for this preview use) capture the traffic from a simple `curl` request to a real server (or your own Lesson 1 TCP server), and identify the three-way handshake's three packets (`SYN`, `SYN-ACK`, `ACK`) by their TCP flags in the capture, before any application data appears.

2. In the same capture, identify the connection teardown — the `FIN`/`ACK` exchange — and note whether it's a clean four-way close or whether your specific test used a `RST` (reset — an abrupt, non-graceful termination) instead, which can happen if one side simply closes without a proper `FIN` sequence.

3. Implement a drastically simplified reliable-delivery mechanism over raw UDP (deliberately unreliable, to make the point): send a sequence of numbered "packets" over UDP, have the receiver track which sequence numbers it has received and send back ACKs, and have the sender retransmit any packet not ACKed within a timeout. Test it by deliberately dropping some packets on the sending side (e.g. skip sending every 5th packet, simulating loss) and confirm your retransmission logic still results in the receiver eventually getting every packet, in order.

4. Measure the actual connection-establishment cost from Learn: time how long it takes to establish a new TCP connection to a real remote server (not localhost, since localhost RTT is near-zero and won't demonstrate the effect) versus reusing an already-open connection for a second request. Report the actual measured difference, connecting it concretely to the one-RTT handshake cost.

## Verify

For step 3, report how many packets you deliberately dropped, how many retransmissions your logic performed, and confirm the final received sequence at the receiver is complete and correctly ordered despite the simulated loss — this is your own minimal, working proof of the reliability mechanism TCP provides for real.

## Failure drill

Modify step 3's simplified protocol to remove the *ordering* guarantee — deliver packets to the receiver's processing logic in whatever order they happen to arrive/get retransmitted, without reordering by sequence number first. Send a message where ordering matters (e.g. numbered chunks of text that only make sense in the correct sequence) with some deliberate reordering introduced (e.g. process out-of-order UDP arrival, or explicitly shuffle processing order for the test), and observe the resulting reassembled message is garbled even though every individual packet was correctly and reliably delivered. Explain why reliable delivery and ordered delivery are two genuinely separate guarantees — your step 3 implementation added reliability (nothing gets lost) without necessarily adding ordering (this drill), and real TCP has to provide *both*, which is precisely why it needs sequence numbers for reassembly, not just an ACK/retransmit loop.

## Transfer

Explain, using this lesson's window/congestion-control discussion and the math-for-engineering queueing lesson's `ρ→1` blowup, why a network path experiencing packet loss due to congestion (too much traffic relative to available capacity) creates a feedback loop: loss triggers retransmission, which — if not managed by proper congestion control — could add even more traffic to an already-overloaded path, worsening the very congestion that caused the original loss. State in your own words why TCP's actual congestion-control algorithms (you don't need to implement one, just understand the goal) deliberately reduce the sending rate in response to detected loss, rather than simply retransmitting immediately at the same rate, which naive reliability alone (like your step 3 implementation) would do.

## Done when

You've identified a real three-way handshake and connection teardown in an actual packet capture, you've built a working (if simplified) reliable-delivery mechanism over UDP and confirmed it survives deliberate packet loss, and you can explain — using your own failure-drill result — why reliable delivery and ordered delivery are distinct guarantees that both need to be explicitly provided, not automatic consequences of each other.
