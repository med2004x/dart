# Lesson 3: UDP and DNS

## Objective

Understand UDP as TCP's deliberately unreliable counterpart and why that's sometimes the right tradeoff, and implement a minimal DNS client to see name resolution as an actual protocol exchange rather than a black-box library call.

## Prerequisites

Lesson 2 (TCP — UDP is best understood by contrast with everything TCP adds on top of it).

## Learn

**UDP: everything TCP adds, deliberately removed.** UDP provides no handshake, no retransmission, no ordering guarantee, no flow/congestion control — just "send this datagram to this address, best effort." This sounds strictly worse than TCP until you consider what those guarantees cost: connection setup latency (Lesson 2's one-RTT handshake), retransmission delay (waiting for a lost packet to be resent, which can be worse than just accepting the loss for some workloads), and head-of-line blocking (TCP's strict ordering means a single lost packet can delay delivery of *later*, already-arrived data to the application, since TCP won't hand data to the application out of order). For applications where a late or lost individual packet is less harmful than the delay of waiting for a guaranteed, ordered retransmission — real-time audio/video, or (this lesson's focus) DNS queries, where a lost query can simply be retried from scratch rather than needing in-protocol retransmission machinery — UDP's simplicity and lower latency win.

**DNS: the protocol that resolves names to addresses.** Every DNS query is (in the common case) a single UDP datagram sent to a resolver, and a single UDP datagram back — matching UDP's connectionless, request-response pattern well, and avoiding TCP's handshake cost for what's meant to be a fast, frequent operation (DNS falls back to TCP for responses too large for a single UDP datagram, or in some other specific cases, but UDP is the default path).

**Recursive vs. authoritative resolution.** When you query a **recursive resolver** (e.g. your ISP's DNS server, or a public one like 8.8.8.8), you're asking it to fully resolve the name on your behalf — it will, if it doesn't already have the answer cached, query the DNS hierarchy itself (starting from root servers, then TLD servers for `.com`/`.ch`/etc., then the domain's own **authoritative server**, which holds the actual, definitive record for that specific domain) and return you the final answer. An authoritative server, by contrast, only answers queries about domains it's specifically responsible for — it doesn't recurse on your behalf. This distinction matters operationally: a misconfigured authoritative server for your own domain (e.g. `tardocsolver.ch`, mentioned in your project history) breaks resolution for everyone querying it, while a slow or overloaded recursive resolver mainly affects lookup latency, not correctness.

**DNS caching and TTL.** Each DNS record has a Time-To-Live (TTL) — a duration for which resolvers are permitted to cache the answer before re-querying. This is why DNS changes (e.g. pointing a domain at a new server) don't take effect instantly everywhere — every resolver that cached the old answer will continue serving it until that specific cached entry's TTL expires, which is exactly why lowering a record's TTL in advance of a planned change is a standard practice for minimizing propagation delay.

## Attempt

1. Use `dig` (or `nslookup`) to query a real domain's A record (its IPv4 address) and observe the raw output, including the TTL value returned. Run the same query again immediately and note whether the TTL you observe has decreased (most resolvers report the *remaining* TTL of a cached answer, decreasing in real time, which is itself direct evidence the answer came from cache rather than a fresh upstream query).

2. Use `dig +trace` (which performs and displays the full recursive resolution process, from root servers down to the authoritative answer) on a real domain, and identify, in the output, the transition from root server responses to TLD server responses to the final authoritative answer — write down, in your own words, which server in the trace was the *authoritative* one for the domain you queried.

3. Implement a minimal DNS client from scratch (Go or C, using raw UDP sockets from Lesson 1 — not a library DNS resolver call): construct a DNS query packet by hand for an A record lookup (following the DNS wire format — a header plus a question section; you'll need to reference the DNS packet format, e.g. RFC 1035 or a concise summary), send it via UDP to a known public resolver (e.g. 8.8.8.8 port 53), and parse the raw response bytes to extract the returned IP address.

4. Capture your own DNS client's query and response with a packet capture tool (a preview of Lesson 7) and compare the raw bytes you constructed by hand against what a real system resolver (e.g. what's generated when you run `dig`) sends for an equivalent query — confirm your hand-built query is structurally valid (the destination server responded correctly, rather than rejecting or ignoring a malformed packet).

## Verify

For step 3, report the actual IP address your from-scratch client extracted for a well-known domain, and independently confirm it matches what `dig` reports for the same domain at roughly the same time (allowing for the possibility of legitimately different answers if the domain uses geo-based or load-balanced DNS responses, which is worth noting as a real possibility if your two lookups happen to hit different resolver infrastructure).

## Failure drill

Modify your step 3 client to query a domain that doesn't exist (e.g. a deliberately made-up subdomain), and observe the response's flags/response-code field (an `NXDOMAIN` response, encoded in the DNS response header) rather than a valid IP address. Parse this response code explicitly in your client and confirm you can distinguish "this domain doesn't exist" from "no response arrived" (a timeout) and from "a valid IP was returned" — three genuinely different outcomes a real DNS client must handle differently, and a common source of bugs in hand-rolled or poorly-tested DNS handling code that only tests the success path.

## Transfer

If TARDOC's `tardocsolver.ch` domain (noted in your project history as blocked by a SWITCH registry restriction, with a `.ink` domain used as a workaround) required a DNS change once the restriction lifted, describe, using this lesson's TTL/caching discussion, what you'd need to plan for in terms of propagation delay — specifically, what you'd want to check about the *current* TTL on any existing records before making a change, and why lowering the TTL well in advance of a planned cutover (if you controlled the records early enough) would reduce how long some users continue resolving to the old answer after the change.

## Done when

You've traced a real domain's full recursive resolution path using `dig +trace` and correctly identified the authoritative answer within it, your from-scratch DNS client correctly parses a real response to extract an IP address, and you've handled the `NXDOMAIN` case explicitly and can explain why distinguishing "doesn't exist" from "timed out" from "resolved successfully" matters for writing a DNS client that behaves correctly under real, not just happy-path, conditions.
