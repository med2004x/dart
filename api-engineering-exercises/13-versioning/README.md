# Project 13 - API Evolution And Compatibility

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Add task due dates without breaking existing clients.

## Compatibility Questions

- Is adding an optional response field safe for all clients?
- Can a field become required?
- Can enum values expand?
- Can status codes change?
- Can pagination defaults change?
- Can field meaning change while name stays?

## Checkpoints

1. complete `compatibility-matrix.md`.
2. add optional `dueAt`.
3. keep old requests valid.
4. make new clients tolerate missing dueAt.
5. write old/new contract tests.
6. add deprecation documentation.
7. define usage measurement.
8. define removal criteria and date.

## Versioning Options

- compatible evolution without new version
- path version
- media type/version header
- separate operation/resource

Do not create a new major version for every additive field. Do not hide breaking
changes under the old contract.

## Failure Drills

1. make an optional request field required.
2. remove an existing response field.
3. change ID from number to string.
4. add an enum value to a client using exhaustive matching.
5. change 404 to 200 with null body.

## Done Means

Old and new clients have executable compatibility tests and a clear migration
window.

