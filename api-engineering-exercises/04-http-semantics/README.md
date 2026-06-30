# Project 04 - HTTP Methods And Status Codes

## Goal

Implement method/path behavior where statuses carry accurate meaning.

## Starter

The server exposes `/tasks` but returns 200 for every method. Correct it.

## Required Behavior

| Method/path | Result |
|---|---|
| GET `/tasks` | 200 |
| POST `/tasks` | 201 |
| GET `/tasks/{id}` existing | 200 |
| GET `/tasks/{id}` missing | 404 |
| DELETE existing | 204 with no body |
| unsupported method | 405 with `Allow` |
| unknown route | 404 |

## Checkpoints

1. route collections and items separately.
2. distinguish malformed ID from missing ID.
3. set headers/status before body.
4. include `Location` on create.
5. ensure 204 has no response body.
6. write `httptest` cases for the table.

## Failure Drills

1. encode body before `WriteHeader(201)`.
2. return 404 for a known path with wrong method.
3. return 200 containing an error message.
4. send a body with 204.

## Done Means

A client can determine result category from method, status, and headers without
parsing English text.

