# Project 09 - Authentication And Authorization

If a syntax item is unfamiliar, use the [track quick reference](../QUICK-REFERENCE.md). It contains a generic example and official documentation without solving this project.

## Goal

Separate identity verification from permission checks on each resource.

## Definitions

Authentication:

```text
Who is calling?
```

Authorization:

```text
May this caller perform this operation on this stored resource?
```

## Checkpoints

1. validate API key or signed token.
2. place trusted actor identity in context.
3. load task and project ownership from storage.
4. check project membership and role.
5. define owner/member/viewer capabilities.
6. apply checks before returning protected data.
7. test cross-tenant IDs.
8. avoid logging credentials.

## Required Matrix

| Role | Read | Create | Update | Delete | Invite |
|---|---|---|---|---|---|
| owner | | | | | |
| member | | | | | |
| viewer | | | | | |

## Failure Drills

1. authorize from `ownerId` in request body.
2. check authentication but no ownership.
3. return resource details before permission failure.
4. trust unsigned token payload.

## Done Means

Changing any client-controlled ID cannot cross tenant boundaries.

