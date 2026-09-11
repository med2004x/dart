SELECT
    current_database() AS database_name,
    current_user AS role_name,
    current_schema() AS active_schema,
    inet_server_port() AS server_port;
