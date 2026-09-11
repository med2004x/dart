-- Wrap exploratory mutations in BEGIN/ROLLBACK.

-- 1. Preview and rename one task. Use RETURNING.

-- 2. Mark one task done only when its current status is open.

-- 3. Postpone open tasks in one project by interval '1 day'.

-- 4. Delete one test task and return its ID/title.

-- 5. Run a bulk update, inspect it, and ROLLBACK.
