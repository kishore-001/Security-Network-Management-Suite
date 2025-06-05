schema/server/server_logs.sql

CREATE TABLE IF NOT EXISTS server_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    time_stamp TIMESTAMP NOT NULL DEFAULT NOW(),
    log_ip VARCHAR(45),
    log_level VARCHAR(20),
    message TEXT,
    device_id UUID REFERENCES server_devices(id) ON DELETE CASCADE
);