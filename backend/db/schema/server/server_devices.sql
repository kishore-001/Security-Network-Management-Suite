schema/server/server_devices.sql

CREATE TABLE IF NOT EXISTS server_devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mac VARCHAR(17) NOT NULL,
    ip VARCHAR(45) NOT NULL,
    tag VARCHAR(100),
    os VARCHAR(100),
    token TEXT NOT NULL UNIQUE,
    added_at TIMESTAMPTZ DEFAULT NOW()
);
