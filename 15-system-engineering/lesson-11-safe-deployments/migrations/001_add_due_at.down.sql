-- This loses due-date data. Use only after proving rollback requirements.
ALTER TABLE app.tasks
DROP COLUMN IF EXISTS due_at;
