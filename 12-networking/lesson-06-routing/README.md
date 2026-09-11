# Lesson 6: Routing and Load Balancing

## Objective

Understand how packets actually find a path across networks (at a level useful for reasoning about real infrastructure, not full routing-protocol depth), and the concrete tradeoffs between L4 and L7 load balancing.

## Prerequisites

Lesson 2 (TCP — L4 balancing operates at exactly the transport layer covered there), Lesson 4 (HTTP — L7 balancing operates at the application layer covered there).

## Learn

**Routing, at the level that matters for this lesson.** Every IP packet carries a destination address; routers along the path each independently decide the "next hop" toward that destination, based on a routing table (built via routing protocols like BGP for the internet backbone, or simpler static/dynamic routing within a private network) — no single router knows the complete end-to-end path in advance, each just forwards toward what it believes is progress toward the destination, and the path emerges from the composition of many independent, local forwarding decisions. This is worth knowing exists even without implementing a router yourself, because it's the reason network paths can change mid-session (a route becomes unavailable, traffic reroutes) and why "traceroute" (showing the sequence of hops a packet actually took) is a meaningful diagnostic tool rather than a fixed, always-identical path.

**Load balancing: distributing traffic across multiple backend servers**, and the key distinction is *at which network layer the balancer operates*, because that determines what information it can see and act on.

**L4 (transport-layer) load balancing.** Operates on IP/TCP information only — source/destination IP and port. It doesn't parse or understand the application protocol (HTTP or otherwise) riding on top of the TCP connection at all; it just decides which backend gets a given *connection*, typically based on some combination of source IP, a hash, or round-robin, and then simply forwards packets for that connection's lifetime. Fast (minimal per-packet processing) and protocol-agnostic (works for any TCP/UDP traffic, not just HTTP), but blind to application-level information — it cannot route based on a URL path, a header, or a cookie, because it never parses the application data.

**L7 (application-layer) load balancing.** Terminates the connection itself (acting as an actual HTTP endpoint, per Lesson 4) and can inspect the full request — URL path, headers, cookies — before deciding which backend to forward to, and can even modify the request/response in transit (adding headers, rewriting paths). This enables routing decisions like "requests to `/api/v2` go to the new backend, everything else goes to the old one," or session affinity based on a cookie — genuinely impossible for a pure L4 balancer, which never sees that information. The cost: more per-request processing (parsing and potentially re-encoding HTTP), and it's HTTP-specific (or specific to whatever application protocol it understands), not general-purpose the way L4 balancing is.

**Connection affinity ("sticky sessions").** For stateful backends (a server holding session state in memory rather than in a shared store), a load balancer may need to ensure the *same client* always reaches the *same backend* — otherwise a request might hit a server with no knowledge of that client's prior session state. L4 balancers typically achieve rough affinity via a hash of the source IP (imperfect — many clients can share one IP behind NAT, and a client's IP itself can change); L7 balancers can use cookie-based affinity, which is more precise but requires the balancer to actively manage that cookie. The better long-term fix, worth naming even briefly, is usually making the backend itself stateless (session state in a shared store like Redis, not the server process's own memory) so affinity becomes unnecessary rather than something to engineer carefully around.

## Attempt

1. Run `traceroute` (or `mtr` for a continuously-updating view) to a real, distant server (not localhost or a same-network host) and record the sequence of hops shown. Identify at least one hop where the round-trip time jumps noticeably compared to the previous hop, and note that this doesn't necessarily mean that specific router is slow — some routers deprioritize or rate-limit the ICMP responses `traceroute` relies on, a real limitation of the tool worth knowing rather than over-interpreting single-hop latency spikes as ground truth.

2. Set up a minimal L4-style load-balancing simulation: run two instances of your Lesson 4 HTTP server on different ports (both otherwise identical), and write a small proxy (using raw TCP forwarding — accept a connection, open a matching connection to one of the two backends chosen by round-robin, and pipe bytes bidirectionally between them with no HTTP parsing at all) in front of them. Confirm requests alternate between the two backends, and confirm your proxy works correctly for arbitrary TCP traffic, not just HTTP specifically (test it by proxying something else simple, like a raw text echo, to prove the proxy itself doesn't care about the application protocol).

3. Extend the setup into an L7-style balancer: instead of blind byte-forwarding, actually parse the incoming HTTP request (Lesson 4's parsing logic) and route based on the request path — e.g. requests to `/a` go to backend 1, requests to `/b` go to backend 2, both determinable only by looking inside the HTTP request, which your L4 proxy from step 2 was structurally incapable of doing.

4. Add cookie-based session affinity to your L7 balancer from step 3: on a client's first request (no session cookie present), pick a backend and set a cookie identifying it; on subsequent requests with that cookie present, always route to the same backend regardless of path-based rules. Test with a sequence of requests simulating the same client (reusing the same cookie) and confirm they consistently reach the same backend.

## Verify

For step 2, send at least 10 requests through your L4 proxy and report the actual distribution across the two backends (should be close to 5/5 for simple round-robin), confirming the proxy is genuinely load-balancing rather than always hitting one backend by accident.

## Failure drill

Take your step 2 L4 proxy and attempt to add path-based routing to it *without* actually parsing the HTTP request — for example, try to guess the path from the TCP connection's timing, byte count, or some other signal visible without parsing. Confirm this is either impossible or produces unreliable, incorrect routing (you're not meant to succeed at this — the point is directly experiencing why L4 balancing structurally cannot make L7-level decisions, since the information genuinely isn't accessible at that layer without doing the parsing work that defines L7 balancing by definition). Explain in your own words, using this concrete failed attempt, why "L4 vs L7" isn't just a performance tradeoff but a genuine information-availability boundary.

## Transfer

If TARDOC's deployment (a single Contabo VPS, per your project history) or Mahall's storefront serving ever needed to scale beyond one backend instance, describe which load-balancing layer would be the right fit given their actual routing needs (e.g. does TARDOC's API need path-based routing to different backend versions, or would simple round-robin L4 balancing across identical backend instances suffice), and state explicitly whether either system currently holds server-side session state that would require sticky sessions (per Learn's discussion) or whether it's already effectively stateless (e.g. state in PostgreSQL/Redis rather than in-process), which would make load balancing significantly simpler to add later without needing session affinity at all.

## Done when

You've built both an L4 (blind TCP forwarding) and an L7 (HTTP-aware) load balancer and can point to specific routing decisions the L7 version can make that the L4 version structurally cannot, you've implemented and tested cookie-based session affinity, and you can explain — using your own failed attempt in the failure drill — exactly where the L4/L7 boundary sits in terms of what information is actually visible at each layer.
