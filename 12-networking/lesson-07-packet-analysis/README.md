# Lesson 7: Packet-Level Debugging

## Objective

Use packet captures as direct evidence when diagnosing network problems, rather than reasoning purely from application-level symptoms — and specifically learn to recognize retransmissions and handshake anomalies in raw capture output.

## Prerequisites

Lesson 2 (TCP — this lesson is about visually/programmatically identifying the mechanisms that lesson described), Linux tools Lesson 5 (the general "measure, don't guess" discipline, applied here specifically to the network layer).

## Learn

**Why packet captures are the ground truth for network problems.** Application-level logs tell you what your code observed (a timeout, an error, a slow response) but not *why* at the network level — was it DNS resolution taking too long (Lesson 3), a slow TCP handshake (Lesson 2), packet loss and retransmission, or the application server itself being slow to respond after the connection was already established. A packet capture shows exactly what went over the wire, with precise timestamps, letting you attribute delay to a specific phase rather than guessing from an aggregate "it was slow" symptom.

**Reading a capture: the essentials.** Tools like `tcpdump` (command-line) and Wireshark (GUI, generally easier for visual pattern recognition, though `tcpdump` is more often available on remote servers you're debugging) show each packet's source/destination, protocol, and (for TCP) flags (`SYN`, `ACK`, `FIN`, `RST`) and sequence/acknowledgment numbers. A normal, healthy TCP exchange shows a clean handshake (Lesson 2), a run of data packets each acknowledged in a timely manner, and a clean four-way close.

**Recognizing retransmissions.** A retransmitted packet is, at the raw protocol level, a packet with a sequence number that's already been sent before — Wireshark specifically flags these as "TCP Retransmission" in its analysis, and `tcpdump` requires you to notice the repeated sequence number yourself (part of why Wireshark's higher-level annotation is often more practical for this specific task). A capture showing retransmissions is direct, unambiguous evidence of packet loss on the path — not application slowness, not a slow server response, but the network layer failing to deliver on the first attempt, requiring TCP's own recovery mechanism from Lesson 2 to kick in.

**Diagnosing "slow" from a capture: attributing delay to a specific phase.** Given a capture of a slow request, you can measure precisely: time from the DNS query to its response (name resolution delay, Lesson 3), time from `SYN` to `SYN-ACK` (network path RTT to the server, plus the server's connection-accept latency), time from the final handshake `ACK` to the first byte of application response (this is "server think time" — the server received the request and took this long to actually respond, which is a symptom of the *application* being slow, not the network), and time spent on retransmissions if any appear. This decomposition is the entire value of packet-level debugging: it turns "the request was slow" into "the request was slow *specifically because* of X phase," which is what actually tells you where to look next.

## Attempt

1. Capture traffic for a real HTTP request to a server you control (your Lesson 4 server, or a public one) using `tcpdump -i any -w capture.pcap port 80` (adjust interface/port as needed), then open the resulting file in Wireshark (or use `tcpdump -r capture.pcap` for a text-based read if Wireshark isn't available) and identify the three-way handshake packets by their flags.

2. In the same capture, identify the first HTTP request packet (containing the actual `GET`/`POST` line) and the first response packet (containing `HTTP/1.1 ...`), and compute the time delta between them — this is your directly measured "server think time" for that specific request.

3. Deliberately introduce artificial packet loss to observe retransmission directly: using a tool like `tc` (Linux traffic control, e.g. `tc qdisc add dev lo root netem loss 20%` to add 20% loss on the loopback interface — remember to remove this afterward with `tc qdisc del dev lo root`, since it affects all loopback traffic while active) or, if unavailable, a network condition simulator appropriate to your platform, run a request against your Lesson 4 server through the lossy link and capture the traffic. Identify retransmitted packets in the capture (repeated sequence numbers, or Wireshark's explicit "TCP Retransmission" flag) and note how the overall request completion time compares to a capture with no artificial loss.

4. Given a capture from step 3 showing retransmissions, calculate the time cost specifically attributable to retransmission — measure elapsed time from the first (lost) attempt of a specific segment to its eventually-successful retransmission, and compare this against the total request time, quantifying what fraction of the observed slowness was retransmission-caused versus other phases.

## Verify

For step 2, report the actual measured server-think-time in milliseconds for a real request, and for step 4, report the actual retransmission-attributable delay in milliseconds, both as concrete numbers extracted from real capture timestamps — not estimates.

## Failure drill

Attempt to diagnose a "slow request" using *only* application-level logging (e.g. a log line printed when your Lesson 4 server starts and finishes handling a request, with no packet capture) for the same lossy-network scenario from step 3. Confirm that your application-level logs show the request as slow, but provide no way to distinguish "the network was lossy and retransmission ate the time" from "my server code itself was slow to process the request" — both would produce an identical symptom (a large gap between request-received and response-sent) if the retransmission delay is happening *before* your application code even sees the complete request, at the TCP layer beneath it. Explain why this specific ambiguity is exactly what motivates packet-level debugging as a distinct skill from application-log analysis: the two failure modes are indistinguishable from inside the application, but trivially distinguishable from a packet capture that shows the network layer independently of what your code was doing.

## Transfer

If you've ever suspected a slow API response in TARDOC or Mahall was "the network" versus "our server code" without being able to confirm which, describe how you'd now set up a packet capture (even briefly, in a controlled test rather than continuous production monitoring, given capture overhead and storage) to definitively attribute the delay to a specific phase using this lesson's decomposition (DNS, handshake, server-think-time, retransmission), rather than relying on intuition or elimination-by-guessing the way you may have before this lesson.

## Done when

You've captured and correctly identified a real three-way handshake and the server-think-time delta for an actual request, you've deliberately induced and observed real TCP retransmissions using artificial packet loss, and you can explain — using the failure drill's specific ambiguity — why application-level logging alone cannot distinguish network-layer delay from application-layer delay, while a packet capture can.
