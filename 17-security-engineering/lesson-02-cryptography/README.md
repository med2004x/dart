# Lesson 2: Cryptography Fundamentals

## Objective

Use standard cryptographic primitives (hashing, HMAC, signatures) correctly in real code, and understand precisely why "rolling your own crypto" is dangerous — not as a platitude, but by directly experiencing a specific, common way naive cryptographic code fails.

## Prerequisites

Networking Lesson 5 (TLS — this lesson covers the underlying primitives that TLS's handshake and certificate mechanisms are built from), API engineering Lesson 11 (webhooks — HMAC signing was already introduced there practically; this lesson covers the underlying theory more completely).

## Learn

**Hashing: one-way, deterministic, fixed-output transformation.** A cryptographic hash function (SHA-256, for instance) takes arbitrary input and produces a fixed-size output, with the property that finding two different inputs producing the same output (a collision) should be computationally infeasible, and that the output reveals nothing about the input beyond its hash value (you cannot practically reverse a hash to recover the original input). Used for: verifying data integrity (does a downloaded file match its published hash), and — critically, discussed further below — never for storing passwords directly, since a plain hash is fast to compute, meaning an attacker with a stolen hash database can try enormous numbers of password guesses per second.

**HMAC: a hash function combined with a secret key, for authenticity.** API engineering Lesson 11 already used this practically for webhook signing. The underlying principle: a plain hash proves data integrity (it wasn't corrupted) but not authenticity (anyone could compute the same hash for altered data, since hashing requires no secret) — HMAC combines the data with a secret key before hashing, so only someone possessing the secret key could have produced a valid HMAC for given data, providing both integrity and authenticity together.

**Digital signatures: asymmetric authenticity, without a shared secret.** Unlike HMAC (which requires both parties to share the same secret key), a digital signature uses a key pair — a private key (kept secret, used to sign) and a public key (freely shared, used to verify). This solves a problem HMAC can't: many different parties can verify a signature (using the public key) without any of them being able to *forge* a new signature (which requires the private key) — exactly the mechanism behind TLS certificates (networking Lesson 5), where a CA's private key signs certificates and anyone can verify that signature using the CA's well-known public key.

**Password storage, specifically: why a plain hash (even a cryptographically strong one) is the wrong tool.** Passwords need a hash function specifically designed to be *slow* and *resource-intensive* to compute (bcrypt, scrypt, or Argon2 — not general-purpose fast hashes like SHA-256), because the threat model is different from data-integrity checking: an attacker with a stolen password-hash database will try to guess passwords by hashing many candidates and comparing — a fast hash lets them try billions of guesses per second on modern hardware, while a deliberately slow, memory-hard function like bcrypt limits guessing throughput dramatically, making large-scale offline cracking attempts far less practical even with a stolen database.

**Why "rolling your own crypto" is genuinely, specifically dangerous — not just a vague warning.** Cryptographic primitives are deceptively easy to implement in a way that *looks* correct (produces plausible-looking output, passes casual testing) while containing a subtle flaw that completely undermines the security property the primitive is supposed to provide — timing side-channels (API engineering Lesson 11's constant-time-comparison point), improper randomness (a genuinely common, catastrophic mistake — using a non-cryptographic random number generator for a key or nonce), or subtly incorrect padding/mode-of-operation choices, each of which can be invisible to functional testing while being fully exploitable by someone who understands the specific flaw.

## Attempt

1. Compute SHA-256 hashes of several inputs (using your language's standard crypto library, never a hand-rolled implementation) and confirm the hash changes completely for even a single-character input change (the "avalanche effect" — a defining property of a good cryptographic hash), demonstrating this isn't a simple checksum where small input changes produce small output changes.

2. Implement HMAC-based message signing and verification (extending API engineering Lesson 11's practical webhook use directly), and confirm that altering either the message or the signature (but not both consistently) causes verification to correctly fail — test both cases explicitly, not just the happy path.

3. Generate a public/private key pair (e.g. using Ed25519 or RSA via your standard library), sign a message with the private key, and verify it with the public key — then confirm that a different, unrelated key pair's public key correctly fails to verify a signature made with the first pair's private key, demonstrating the public key genuinely corresponds to a specific private key, not just "any signature looks valid to any public key."

4. Implement password hashing correctly using bcrypt (via a standard library, e.g. `golang.org/x/crypto/bcrypt`), and separately, incorrectly, using plain SHA-256 with no salt. Benchmark (performance track Lesson 1's methodology) how many hash computations per second each approach allows on your machine, and use that number to estimate how long an attacker with a stolen hash database could brute-force a modest password space (e.g. all 6-digit numeric PINs) under each approach — report the concrete time difference, making the "why does slowness matter for password hashing" point numerically real rather than abstract.

## Verify

Present your step 2 HMAC verification results (both the correct-pass and correct-fail cases), your step 3 signature verification results (correct key pair passes, wrong key pair fails), and your step 4 brute-force time estimates for bcrypt versus plain SHA-256, showing the concrete, large difference in attacker cost.

## Failure drill

Implement a deliberately naive "custom" password verification scheme: instead of using bcrypt's built-in, constant-time comparison (which handles the timing-safety concern automatically), implement your own comparison of a computed hash against a stored one using a standard, non-constant-time string comparison (`==` or equivalent). This directly reuses API engineering Lesson 11's constant-time-comparison concern, now applied to password verification specifically rather than webhook signatures. You don't need to demonstrate a successful timing attack (genuinely hard to do reliably in a small exercise), but research and explain, in your own words, why a naive comparison here reintroduces exactly the timing side-channel risk that lesson identified, and why using your standard library's provided, already-hardened comparison function (rather than writing your own) is the correct default specifically because these subtle correctness requirements are easy to miss when re-implementing something that already exists in a battle-tested library.

## Transfer

If TARDOC or Mahall currently stores user passwords, confirm explicitly (by checking the actual code, not assuming) whether it uses a proper slow, memory-hard hash function (bcrypt, scrypt, Argon2) or something weaker (plain SHA-256, MD5, or worse). If it's using something weaker, describe what migrating existing stored password hashes to bcrypt would require (typically: hash with the new algorithm at the user's *next successful login*, since you cannot retroactively re-hash a password you don't have in plaintext — a real, common migration pattern worth knowing rather than assuming a simple database update would suffice).

## Done when

You've correctly used hashing, HMAC, and digital signatures in real, working code with both success and failure cases tested explicitly, you've quantitatively demonstrated — with real benchmark numbers — why slow password hashing matters against brute-force attacks, and you can explain, using the constant-time-comparison example, a specific, concrete way that reimplementing a cryptographic operation yourself (rather than using a vetted library function) can silently reintroduce a real vulnerability despite passing all functional tests.
