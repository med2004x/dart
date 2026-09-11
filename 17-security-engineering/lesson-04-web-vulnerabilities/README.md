# Lesson 4: Web Security Failure Modes

## Objective

Reproduce, understand, and patch the core web vulnerability classes — injection, CSRF, SSRF, XSS — in local, controlled examples, understanding each as a specific instance of untrusted input crossing a trust boundary (system-engineering Lesson 10) without adequate handling.

## Prerequisites

System-engineering Lesson 10 (trust boundaries — every vulnerability in this lesson is a concrete instance of that lesson's general principle), API engineering Lesson 5 (validation — the general defense this lesson's specific vulnerabilities all relate back to).

## Learn

**SQL injection: untrusted input crossing into a trusted query-construction context.** If user input is directly concatenated into a SQL query string (`"SELECT * FROM users WHERE name = '" + userInput + "'"`) rather than passed as a parameterized value, an attacker can craft input that changes the query's actual structure (e.g. input containing `' OR '1'='1`, altering the WHERE clause's logic entirely). The fix — parameterized queries/prepared statements — isn't just "escape the input carefully" (a fragile, error-prone approach), it's structurally separating the query's fixed logic from the data values, so user input can never be interpreted as SQL syntax regardless of its content, exactly the "validate/sanitize at the trust boundary" principle from system-engineering Lesson 10, applied specifically to database queries.

**Cross-Site Scripting (XSS): untrusted input crossing into a trusted HTML/JavaScript rendering context.** If user-supplied content is rendered directly into an HTML page without escaping (`<div>{{userComment}}</div>` where `userComment` might contain `<script>...</script>`), an attacker can inject and execute arbitrary JavaScript in other users' browsers when they view the page — the same category of bug as SQL injection, just with HTML/JS as the target syntax instead of SQL. The fix is the same principle: proper output escaping (converting special characters to their safe, literal representations) at the point where untrusted data crosses into a rendering context, ensuring it's always treated as data, never as executable markup/script.

**Cross-Site Request Forgery (CSRF): exploiting a browser's automatic credential inclusion.** Browsers automatically attach session cookies to requests to a given domain, regardless of what page initiated the request. If a malicious site can trick a logged-in user's browser into making a request to your site (e.g. a hidden form auto-submitting to `your-site.com/delete-account`), the browser will include the user's real session cookie, and your server — seeing a validly authenticated request — may process it, even though the *user* never intended to make that request. The standard defense: CSRF tokens (a random, unpredictable value embedded in your own forms, checked on submission) that an attacker's foreign site has no way to know or include, since they can't read your page's content, only trigger a request to it.

**Server-Side Request Forgery (SSRF): tricking your server into making requests on an attacker's behalf.** If your server accepts a URL from user input and fetches it (e.g. "fetch this image URL to generate a thumbnail"), an attacker can supply an internal URL (e.g. `http://localhost:6379` for a Redis instance, or a cloud metadata endpoint like `http://169.254.169.254`) that your server — trusted to make internal network requests — will happily fetch, potentially exposing internal services or sensitive metadata the attacker couldn't reach directly from outside your network. This is again the same underlying pattern: untrusted input (a URL) crossing into a trusted-context operation (your server's own network access privileges) without adequate validation of what that input is allowed to specify.

## Attempt

1. In a local, sandboxed test application, construct a SQL injection vulnerability (a query built via string concatenation with user input) and demonstrate a successful injection (e.g. bypassing an intended WHERE clause condition using crafted input). Fix it using parameterized queries and confirm the same malicious input now has no effect beyond being treated as a literal, harmless string value.

2. Construct an XSS vulnerability (unescaped user content rendered into HTML) and demonstrate it by injecting a harmless but observable script (e.g. one that triggers a visible `alert()` or writes to the page in an obviously detectable way) — confirm it actually executes when the page is viewed. Fix it using proper output escaping (your template engine's default escaping, or an explicit escaping function) and confirm the same input now renders as inert, literal text rather than executing.

3. Construct a CSRF vulnerability (a state-changing endpoint with no CSRF token check, relying solely on session cookies for authentication) and demonstrate the attack using a separate, local "malicious" HTML page that auto-submits a form to your vulnerable endpoint — confirm, while logged into your test app in a browser, that visiting the malicious page triggers the unintended action. Fix it by adding CSRF token generation and validation, and confirm the same attack page can no longer successfully trigger the action.

4. Construct an SSRF vulnerability (an endpoint that fetches a user-supplied URL server-side with no restriction) and demonstrate it by supplying a URL pointing to a local service only your server should be able to reach (e.g. a simple local HTTP server running on `localhost` at a port not exposed externally), confirming your vulnerable endpoint successfully fetches content it shouldn't be able to reach from outside. Fix it with an allowlist restricting fetchable URLs to expected, legitimate external domains, and confirm the same internal-URL attack is now rejected.

## Verify

For each of the four vulnerabilities, show both the successful attack (before the fix) and the same attack correctly failing (after the fix) — real, working before/after demonstrations for all four, not just a description of the general mechanism.

## Failure drill

Take your step 2 XSS fix (proper output escaping) and construct a *second*, different injection point in the same application that you didn't originally fix (e.g. if you fixed escaping in a comment-rendering template, try injecting via a different field, like a username, rendered in a different part of the page that uses a different, unescaped rendering path). Confirm this second injection point is still vulnerable despite your first fix, demonstrating that fixing one specific instance of a vulnerability class doesn't automatically fix every instance — each place untrusted data crosses into a rendering (or query, or fetch) context needs its own correct handling, and a security review needs to systematically check *every* trust boundary crossing, not stop after finding and fixing the first example.

## Transfer

Audit one real endpoint in TARDOC or Mahall for each of these four vulnerability classes: does it use parameterized queries throughout (not just in the one place you might think to check), does user-generated content get properly escaped everywhere it's rendered, do state-changing endpoints have CSRF protection if they rely on cookie-based sessions, and does anything fetch user-supplied URLs server-side without restriction. Report your actual findings — if everything checks out, that's a legitimate, useful finding; if you find a real gap, describe specifically what fixing it would require using this lesson's patterns.

## Done when

You've successfully demonstrated and then correctly fixed all four vulnerability classes in local, controlled examples, with real before/after evidence for each, and you've demonstrated — via the failure drill — that fixing one instance of a vulnerability class doesn't protect every instance, understanding why a systematic audit across all trust-boundary crossings, not a single spot-fix, is what real security review requires.
