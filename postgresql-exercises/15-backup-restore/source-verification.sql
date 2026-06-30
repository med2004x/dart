SELECT count(*) AS user_count FROM app.users;
SELECT count(*) AS project_count FROM app.projects;
SELECT count(*) AS task_count FROM app.tasks;
SELECT count(*) AS constraint_count
FROM information_schema.table_constraints
WHERE table_schema = 'app';
