BEGIN;

UPDATE app.accounts
SET balance = balance + 10
WHERE id = 1
RETURNING id, balance;

-- This waits while session A owns the conflicting row lock.
ROLLBACK;
