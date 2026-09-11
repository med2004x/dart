SET lock_timeout = '2s';
SET statement_timeout = '30s';

ALTER TABLE app.tasks
ADD COLUMN IF NOT EXISTS due_at timestamptz;
