# Lesson 3: OpenAPI Contract

## Objective

Write a real OpenAPI specification for an API before implementing it, and use it as the actual source of truth — generating documentation and, ideally, validation from it — rather than treating the spec as documentation written after the fact to describe whatever the code already does.

## Prerequisites

Lessons 1-2 (requirements and resource modeling — OpenAPI is the formal notation for expressing exactly what those lessons had you reason through informally).

## Learn

**What OpenAPI actually is.** A structured (YAML or JSON), machine-readable specification format for describing an HTTP API: its paths, methods, request/response schemas, status codes, and more. Being machine-readable — not just human documentation — is the entire point: real tooling can generate interactive documentation (Swagger UI), generate client SDKs in multiple languages, generate server-side request validation, and generate contract tests, all from one source of truth, rather than each of these being separately maintained (and inevitably drifting out of sync with each other and with the actual implementation) documents.

**Why writing the spec first (not after) changes the outcome.** If you write the OpenAPI spec after implementing the endpoint, you're describing whatever the implementation happened to produce — including any accidental inconsistencies, undocumented edge cases, or implementation-detail leakage Lesson 1 warned about. Writing the spec first forces the same "what does the client actually need" discipline Lesson 1 introduced, made concrete and precise enough that ambiguity becomes visible immediately (a vague requirement like "return an error" has to become a specific status code and response schema the moment you try to write it in OpenAPI's structured format).

**Schemas, precisely.** OpenAPI describes request/response bodies using JSON Schema — specifying required vs. optional fields, types, formats (e.g. a string that should be an email, a date-time), and constraints (min/max length, numeric ranges, enum values). This precision is what enables auto-generated validation: a request body that violates the schema (missing a required field, wrong type) can be rejected automatically, before your handler code even runs, by tooling that reads the same spec you wrote to communicate the contract to human readers.

**Reusable components.** OpenAPI supports defining reusable schema components (e.g. a standard `Error` response shape used across every endpoint's failure responses) referenced via `$ref` rather than duplicated inline everywhere — this matters because a duplicated schema, copy-pasted into 15 different endpoint definitions, will inevitably drift out of sync when one copy gets updated and the others don't, exactly the kind of maintenance problem structured, referenced components are designed to prevent.

## Attempt

1. Write a complete OpenAPI 3.x specification (YAML) for one resource from your Lesson 2 resource model, covering at least the list and get-one operations: paths, methods, response schemas (including field types and which fields are required), and at least two distinct error responses (e.g. 404 for not found, 400 for a malformed request) with their own schemas.

2. Define a reusable `Error` schema component (with at minimum a machine-readable error code and a human-readable message field) and reference it via `$ref` from every error response in your spec, rather than duplicating the error shape inline for each endpoint.

3. Validate your spec using a real tool (the `swagger-cli validate` command, or an online OpenAPI validator, or your editor's OpenAPI linting if available) and fix any structural errors it reports — confirming your YAML is not just plausible-looking but actually spec-compliant.

4. Generate interactive documentation from your spec using a real tool (Swagger UI, Redoc, or an equivalent — many can run locally against a spec file with minimal setup) and actually view the rendered result, confirming it accurately reflects what you intended and is something you'd be comfortable handing to an external API consumer as documentation.

## Verify

Show your validated OpenAPI YAML for the resource you chose, confirm it passed the validator with no errors, and show a screenshot or description of the generated documentation output, confirming it correctly renders your defined paths, schemas, and error responses.

## Failure drill

Deliberately introduce an inconsistency into your spec: define a field as `required` in the request schema, but then write example request payloads (if your spec includes examples) that omit it, or define a response field's type as `string` in the schema but write an example value that's actually a number. Run your validator again and observe whether it catches this specific class of inconsistency (schema-vs-example mismatch) — many validators check spec structure but not example-vs-schema consistency, which is worth discovering directly rather than assuming your tooling catches every possible contract inconsistency. Explain, based on what you actually observed the validator catch versus not catch, why "the spec validates successfully" is a necessary but not sufficient condition for "the spec is actually correct and internally consistent" — validation catches structural/syntactic problems, not necessarily semantic ones like this.

## Transfer

If TARDOC's or Mahall's API doesn't currently have an OpenAPI spec, pick one existing, already-implemented endpoint and write the spec for it *after the fact* (the reverse of this lesson's recommended workflow, done deliberately here for a practical reason) — and explicitly note every place where writing the spec for existing code revealed something you hadn't consciously decided (an inconsistent field naming convention, an undocumented error case your code actually handles, a field that's technically optional in your implementation but you'd always assumed was required). This exercise demonstrates concretely why writing the spec first, as this lesson recommends, avoids discovering this kind of unintentional inconsistency after the API is already in use by real clients.

## Done when

You have a validated, structurally correct OpenAPI spec for a real resource with properly reused error-schema components, you've generated and reviewed real interactive documentation from it, and you've identified — through the failure drill — a specific category of inconsistency your validation tooling does not catch, so you understand precisely what "validates successfully" does and doesn't guarantee about your spec's actual correctness.
