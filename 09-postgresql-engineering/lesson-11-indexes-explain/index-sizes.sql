SELECT
    indexrelid::regclass AS index_name,
    pg_size_pretty(pg_relation_size(indexrelid)) AS size
FROM pg_index
WHERE indrelid = 'app.tasks'::regclass
ORDER BY pg_relation_size(indexrelid) DESC;
