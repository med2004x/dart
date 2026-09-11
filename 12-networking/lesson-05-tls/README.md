# Lesson 5: TLS and Secure Transport

## Objective

Understand what TLS actually provides — encryption, authentication, integrity — and specifically what a certificate proves (and doesn't prove), by inspecting a real handshake rather than treating "HTTPS" as an opaque padlock icon.

## Prerequisites

Lesson 4 (HTTP — TLS is most commonly encountered wrapping HTTP into HTTPS, though it's a general-purpose transport security layer, not HTTP-specific), probability Lesson 1 of the mathematics track is a useful (not required) background for the asymmetric-cryptography intuition, though this lesson doesn't require deriving any of the underlying math.

## Learn

**Three separate guarantees, commonly conflated into "it's secure."** TLS provides: **confidentiality** (an eavesdropper on the network path can't read the content — the connection is encrypted), **integrity** (the content can't be silently modified in transit without detection — any tampering breaks a cryptographic check), and **authentication** (the client can verify it's actually talking to the server it intended to, not an impostor intercepting the connection — this is specifically what the certificate mechanism provides). These are genuinely separable properties; understanding them separately matters because a system can have some without others (e.g. an encrypted-but-unauthenticated connection is still vulnerable to a man-in-the-middle who simply presents their own valid encryption keys, since nothing was checking *whose* keys they were).

**What a certificate actually proves.** A certificate is a signed statement, from a Certificate Authority (CA) that the client already trusts (via a pre-installed list of trusted root CAs), asserting "this public key belongs to this domain name." When your browser connects to `example.com` and receives a certificate, verification means: checking the certificate's signature was made by a trusted CA, checking the certificate hasn't expired, and checking the certificate's domain name matches the one you intended to connect to. This proves the server presenting the certificate possesses the corresponding private key *and* that a CA was willing to vouch for the domain-to-key binding at issuance time — it does **not** prove the server is trustworthy, well-configured, or free of vulnerabilities; it only proves you're talking to the entity that legitimately controls that domain (as far as the CA's issuance process verified), not an on-path impersonator.

**The handshake, at a level of detail useful for actually reading one.** TLS 1.3 (the current standard) roughly: client sends supported cipher suites and a key-exchange value (`ClientHello`); server responds with its chosen cipher suite, its certificate, and its own key-exchange value (`ServerHello` plus certificate); both sides independently derive the same symmetric session key from the exchanged key-exchange values (via Diffie-Hellman-family key exchange, without ever transmitting the actual shared secret over the wire — the mathematical property that makes this possible is a genuinely elegant result, out of scope to derive here but worth knowing exists); subsequent application data is encrypted with that derived symmetric key. Certificate verification (checking the CA signature and domain match) happens as part of processing the server's certificate message.

**Why self-signed certificates trigger warnings, and when that's actually fine.** A self-signed certificate has no CA vouching for the domain-to-key binding — it's the server saying "trust me" with nothing external to verify it against. Browsers correctly warn about this by default, because in the open internet, an attacker can just as easily self-sign a certificate claiming to be `your-bank.com`. For local development or testing (this lesson's practical exercise), a self-signed certificate is a reasonable, common practice — the warning is expected and acceptable specifically because you control both endpoints and aren't relying on the CA trust chain to establish that this is really your own test server.

## Attempt

1. Generate a local, self-signed TLS certificate (e.g. `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 1 -nodes`) and use it to run a minimal HTTPS server (Go's `net/http` with `ListenAndServeTLS`, pointing at your generated cert/key, is straightforward for this). Connect to it with `curl -k` (the `-k` flag skips certificate verification, appropriate here since you know it's self-signed and you're testing locally) and confirm the connection succeeds and is encrypted.

2. Connect to the same self-signed server *without* `-k` (i.e. `curl` with normal certificate verification enabled) and observe it reject the connection, reporting the certificate isn't trusted. Read the actual error message and confirm it references the missing CA trust chain, not a problem with the encryption itself — directly demonstrating the separation between "encrypted" and "authenticated" from Learn: the encryption would work fine, but verification correctly refuses to proceed without a trust basis.

3. Use `openssl s_client -connect example.com:443` (substituting a real domain) to manually inspect a real TLS handshake against a production server. In the output, identify the certificate chain presented (often more than one certificate — the server's own certificate plus one or more intermediate CA certificates linking up to a root CA your system trusts), and the negotiated TLS version and cipher suite.

4. Deliberately connect to a server using its IP address directly (bypassing DNS) but request a certificate check against a hostname that doesn't match what the certificate actually covers (e.g. connect to a real server's IP but present a `Host` header, or use `openssl s_client -connect <ip>:443 -servername wrong-hostname.example`), and observe the resulting hostname-mismatch verification failure — demonstrating the domain-matching part of certificate verification explicitly and separately from the CA-trust-chain part exercised in step 2.

## Verify

For step 3, report the actual certificate chain length (how many certificates were presented) and the negotiated cipher suite for a real production server, and state in your own words what each certificate in the chain is vouching for relative to the next one up the chain.

## Failure drill

Take your step 1 self-signed server and, using a packet capture tool (Lesson 7 preview) or a network proxy tool, attempt to read the actual HTTP request/response content passing over your `curl -k` connection at the raw TCP byte level (not through the TLS-aware tool itself, but genuinely at the wire). Confirm the payload is unreadable ciphertext, not plaintext HTTP — directly demonstrating the confidentiality guarantee is real and independent of whether authentication was properly verified in this specific test (since you used `-k` to skip that check). Explain why this is the concrete, hands-on version of the abstract "separable guarantees" point from Learn: your step 1/2 test showed authentication can fail while encryption still works, and this drill shows encryption genuinely works (not just claimed) even in a scenario where you deliberately bypassed authentication.

## Transfer

If TARDOC's clinic-facing API or Mahall's storefronts serve traffic over HTTPS (a near-certain requirement given they handle client data and payment-adjacent flows), describe, using this lesson's certificate-chain and expiration concepts, what operational risk an expired certificate poses (connections start failing verification, exactly like your step 2 self-signed rejection, except now unexpectedly for real users) and why automated certificate renewal (e.g. via Let's Encrypt's ACME protocol, which issues free, short-lived certificates specifically designed to be renewed automatically before expiry) is standard practice rather than manually renewing a certificate on a calendar reminder.

## Done when

You've run a real TLS handshake against your own self-signed server and observed both the encrypted-but-unverified case and the properly-rejected-for-lack-of-trust case, you've inspected a real production certificate chain and can explain what each link in it vouches for, and you can state clearly, using your own failure-drill observation of genuinely unreadable ciphertext, why confidentiality and authentication are separate guarantees that TLS provides together but that can be reasoned about — and can fail — independently.
