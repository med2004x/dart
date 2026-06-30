SELECT current_user, current_database();

SELECT id, title
FROM app.tasks
ORDER BY id
LIMIT 1;

-- Run separately and expect denial:
-- DROP TABLE app.tasks;
-- CREATE ROLE should_fail;
-- CREATE TABLE app.should_fail (id integer);
