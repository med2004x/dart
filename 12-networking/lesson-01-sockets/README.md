# Lesson 1: Sockets and Byte Streams

## Objective

Build a working TCP client/server, and understand why "sending a message" over TCP requires application-level framing — TCP itself has no concept of message boundaries, only a continuous byte stream.

## Prerequisites

OS Lesson 6 (IPC — sockets are conceptually the network-spanning generalization of the local IPC mechanisms covered there).

## Learn

**A socket is an endpoint, addressed by an IP and port.** Creating a TCP socket, binding it to an address, and calling `listen()`/`accept()` (server side) or `connect()` (client side) establishes a connection the OS represents, on each end, as a file-descriptor-like handle — `read()`/`write()` (or `send()`/`recv()`) work on it much like the pipe file descriptors from OS Lesson 6, which is a deliberate design consistency in Unix: "everything is a file descriptor" extends to network connections too.

**TCP is a byte stream, not a message protocol.** This is the single most important, most commonly misunderstood fact in this lesson. If you `write()` "hello" and then `write()` "world" on the sending side, the receiving side's `read()` calls are **not** guaranteed to return "hello" and "world" as two separate reads — TCP may deliver them as one read of "helloworld", or split unpredictably across multiple reads, or combine multiple sender-side writes into fewer receiver-side reads. There is no message boundary preserved by TCP itself; it only guarantees the *bytes* arrive in order, without duplication, eventually (or the connection reports an error) — nothing about how those bytes are grouped into individual `read()`/`write()` calls.

**Framing: the application's job of imposing message boundaries.** Since TCP won't do it for you, any protocol built on TCP needs its own convention for where one message ends and the next begins. Two standard approaches: **length-prefixing** (send a fixed-size length field before each message, so the receiver knows exactly how many more bytes to read to complete it) and **delimiter-based** (send messages separated by a special byte sequence, like HTTP's `\r\n` line endings, with the receiver scanning for the delimiter). Length-prefixing is simpler to implement correctly and doesn't require escaping if the delimiter byte could appear in the data itself; delimiter-based framing is more human-readable for text protocols but requires care around escaping.

**Partial reads and writes are normal, not exceptional.** `read()` can return fewer bytes than requested (even if more are coming, and even if the connection is healthy) — this is not an error, it's normal TCP behavior, and code that assumes a single `read()` call always returns a complete message (or the full requested buffer) will intermittently and unpredictably fail, precisely because it happens to work when the OS/network delivers everything in one chunk during testing and silently breaks whenever the actual delivery happens to fragment differently.

## Attempt

1. Write a minimal TCP server (bind, listen, accept) and client (connect) in Go or C that exchanges a single short message. Confirm it works for a small message.

2. Demonstrate the "no message boundaries" fact directly: modify the client to send two separate messages via two separate `write()`/`Write()` calls in quick succession (e.g. "hello" then "world", with no delay between them), and have the server do a single `read()` immediately after accepting the connection, printing exactly what it received in that one read. Run it several times if needed to try to observe the two messages arriving concatenated in a single read (this is timing-dependent and may not reproduce every single run, particularly on loopback connections, which is itself worth noting as a limitation of this specific test — but if you can force a large enough payload or introduce artificial delay variance, you should be able to observe non-1:1 read/write correspondence at least sometimes).

3. Implement length-prefixed framing: before sending each message, send a fixed-size (e.g. 4-byte) integer giving the message's length, then the message bytes themselves. On the receiving side, first read exactly 4 bytes (looping if a single `read()` returns fewer, since Learn established this is normal) to get the length, then read exactly that many more bytes (again looping as needed) to get the complete message. Test with messages of varying length, including ones larger than a typical single `read()` buffer size, to force the receiver's read-loop to actually execute more than one iteration.

4. Deliberately simulate a partial-read scenario without relying on network timing luck: send a large message (e.g. several hundred KB) over a loopback connection, but on the receiving side, call `read()` with a small buffer (e.g. 16 bytes) repeatedly in a loop, and count how many `read()` calls it actually took to receive the complete message. Confirm it took more than one call, and that your length-prefixed framing from step 3 correctly reassembled the complete message despite arriving across many small reads.

## Verify

For step 4, report the exact number of `read()` calls it took to receive the full message with your small buffer size, and confirm the reassembled message exactly matches the original bytes sent (e.g. via a checksum or direct byte comparison) — this is your concrete proof that your framing and read-loop logic correctly handles partial reads rather than just happening to work on short messages.

## Failure drill

Write a *naive* receiver that assumes a single `read()` call always returns a complete message (no length-prefix framing, no read-loop — just one `read()` call per expected message) and test it against step 4's large-message, small-buffer-size scenario. Confirm it receives only a fragment of the intended message (whatever fit in the first `read()` call) and either processes that truncated fragment as if it were complete, or errors out unexpectedly depending on what the fragment happens to contain. Explain why this naive version might have appeared to work correctly during earlier, smaller-message testing (steps 1-2), lulling a developer into believing the assumption was safe, when in fact it was only ever true by coincidence of small message size and favorable timing, not because TCP guarantees it.

## Transfer

HTTP (covered properly in Lesson 4) solves the exact same framing problem this lesson covers, using a mix of both strategies covered in Learn: headers are delimiter-based (terminated by `\r\n`, with the full header block terminated by an empty line), while the body is commonly length-prefixed via the `Content-Length` header (or uses chunked transfer encoding, a different but related framing mechanism for when the length isn't known in advance). Read the raw bytes of a real HTTP request (e.g. capture one with `nc -l` acting as a minimal server and a real browser or `curl` as the client, or use Lesson 7's packet capture tooling) and identify, in the raw byte stream, exactly where framing convention is doing the work of telling the receiver where the headers end and the body begins.

## Done when

You have a working TCP client/server with correct length-prefixed message framing, you've directly demonstrated (not just read about) that TCP doesn't preserve write-call boundaries, and you've shown your framing implementation correctly reassembles a message that arrived across many small, partial `read()` calls, while explaining why a naive single-read implementation would silently fail under exactly those conditions.
