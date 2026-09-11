# Lesson 4: HTTP Semantics

## Objective

Build a minimal HTTP server and client from raw sockets (not a framework) to understand exactly what HTTP adds on top of the TCP byte stream from Lessons 1-2, and correctly implement connection reuse (keep-alive) and status-code semantics.

## Prerequisites

Lesson 1 (sockets and framing — HTTP is, at its core, a specific framing convention layered on TCP), Lesson 2 (TCP — HTTP's connection reuse directly addresses the handshake cost covered there).

## Learn

**HTTP as a framing protocol, revisited.** Lesson 1 introduced framing generically; HTTP is a concrete, standardized instance of it. An HTTP/1.1 request is: a request line (`GET /path HTTP/1.1`), a series of `Header: value` lines each terminated by `\r\n`, a blank line (`\r\n` alone) marking the end of headers, and optionally a body — whose length is determined either by a `Content-Length` header (length-prefixing, in Lesson 1's terms) or `Transfer-Encoding: chunked` (a length-prefixing scheme applied per-chunk, used when the total length isn't known in advance, e.g. a streamed response). A response has the same shape, starting with a status line (`HTTP/1.1 200 OK`) instead of a request line.

**Status codes carry real semantic meaning, not just "success or failure."** The first digit groups them: 2xx (success), 3xx (redirection — the response tells the client to look elsewhere), 4xx (client error — the request itself was invalid or unauthorized), 5xx (server error — the request was valid but the server failed to fulfill it). This distinction matters practically: a client should generally retry a 5xx (the server might recover) but should not blindly retry a 4xx with the same request (it will fail identically, since the problem is the request itself) — conflating these categories in retry logic is a real, common source of wasted retries or, worse, silently swallowed persistent client errors.

**Connection reuse (keep-alive).** HTTP/1.1 defaults to keeping the underlying TCP connection open after a response, allowing subsequent requests to the same server to skip the handshake cost (Lesson 2) entirely. This is a direct, measurable performance optimization — Lesson 2's connection-establishment timing experiment demonstrated the cost keep-alive avoids. The `Connection: close` header (either side can send it) signals the connection should be closed after the current exchange instead of reused, and correctly handling this — not assuming every connection stays open forever, and not closing every connection unnecessarily — is part of implementing HTTP correctly rather than just "sort of working."

**Idempotency and methods, a preview relevant to API engineering later.** `GET` is defined as safe (should not have side effects) and idempotent (repeating it produces the same result as doing it once). `POST` is neither by convention. This isn't enforced by HTTP itself — nothing stops a server from implementing a `GET` handler with side effects — but violating the convention breaks assumptions clients, proxies, and caches make (e.g. a browser or proxy may prefetch or retry `GET` requests, assuming safety), which is exactly the kind of contract violation that produces confusing, hard-to-diagnose bugs far from the code that violated it.

## Attempt

1. Using raw TCP sockets (Lesson 1, not an HTTP library), implement a minimal HTTP server that parses an incoming request line and headers (up to the blank-line terminator), and responds with a hardcoded `200 OK` response and a small body, including a correct `Content-Length` header matching the actual body byte length.

2. Test your server with a real client (`curl -v` shows the raw request/response headers) and confirm the response is correctly interpreted — `curl` should display your status code and body without error, and `curl -v`'s verbose output should show the exact bytes exchanged, letting you confirm your hand-built response matches the format real clients expect.

3. Extend your server to support keep-alive: after responding, don't close the connection — loop back and parse another request on the same TCP connection, unless the client's request included `Connection: close`. Test with `curl` making two requests using the same connection (`curl` reuses connections automatically for sequential requests to the same host within one invocation when using certain flags, or more directly, use a small script that opens one TCP connection and sends two full requests over it manually) and confirm both are correctly handled without your server closing the socket between them.

4. Implement basic status-code handling: have your server return `404 Not Found` for unrecognized paths and `200 OK` for a specific known path, with each response's headers correctly reflecting its own body length. Test both paths and confirm the client-observed status code and body match expectations for each.

## Verify

For step 3, use a packet capture or `strace` on your server process to confirm only *one* TCP handshake (one `SYN`/`SYN-ACK`/`ACK` sequence, per Lesson 2) occurred across both requests, directly proving connection reuse is actually happening at the TCP level, not just that your application code returned two responses.

## Failure drill

Deliberately send a `Content-Length` header from your server that doesn't match the actual body's real byte length (e.g. claim 100 bytes but send only 50). Test with `curl -v` and observe the client either hang waiting for the remaining, never-arriving bytes (if it's a keep-alive connection, since the client believes more of the response is still coming), or produce a truncated/corrupted body. Explain, using Lesson 1's framing discussion directly, why this specific bug is a framing-boundary violation — the client is doing exactly what correct framing logic should do (trust the length prefix and read that many bytes), and the bug is entirely in your server providing an inaccurate length, not in the client's read logic.

## Transfer

Compare your hand-rolled server's request parsing to what Go's `net/http` package does automatically (correct header parsing, `Content-Length` handling, keep-alive management, all handled by the standard library). Write the equivalent minimal server using `net/http` instead of raw sockets, confirm it produces the same observable behavior as your hand-rolled version for steps 1-4, and state explicitly, using your own hand-rolled implementation as the point of comparison, what specific framing and connection-management complexity the standard library is quietly handling for you every time you write a normal Go HTTP handler without thinking about any of this.

## Done when

Your hand-rolled HTTP server correctly handles request parsing, accurate `Content-Length` responses, keep-alive connection reuse (verified at the TCP level, not just assumed), and distinct status codes for different paths, and you've directly demonstrated — via the failure drill — what breaks when the length-prefix framing contract is violated, connecting it back to Lesson 1's general framing principles rather than treating it as an HTTP-specific quirk.
