BEGIN;

SELECT id, balance
FROM app.accounts
WHERE id = 1
FOR UPDATE;

-- Wait while session B attempts its update.
-- COMMIT to release, or ROLLBACK to discard.
