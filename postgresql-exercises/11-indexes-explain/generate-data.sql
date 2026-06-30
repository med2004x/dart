INSERT INTO app.tasks (
    project_id,
    title,
    status,
    priority,
    created_at,
    updated_at
)
SELECT
    (SELECT min(id) FROM app.projects),
    'generated task ' || number,
    CASE WHEN number % 4 = 0 THEN 'done' ELSE 'open' END,
    (number % 5) + 1,
    now() - (number || ' minutes')::interval,
    now() - (number || ' minutes')::interval
FROM generate_series(1, 50000) AS number;

ANALYZE app.tasks;
