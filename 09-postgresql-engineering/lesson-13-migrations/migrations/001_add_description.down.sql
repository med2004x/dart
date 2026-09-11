-- This destroys description data. Application rollback should normally keep
-- the additive schema instead of immediately running this file.
ALTER TABLE app.tasks
DROP COLUMN IF EXISTS description;
