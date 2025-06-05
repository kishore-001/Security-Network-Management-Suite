schema/server/server_log_info.sql

CREATE TABLE IF NOT EXISTS server_log_info (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    table_name VARCHAR(255) NOT NULL
);