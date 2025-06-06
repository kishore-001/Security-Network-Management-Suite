CREATE TABLE IF NOT EXISTS server_log_volume (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    log_date DATE NOT NULL,
    volume_of_requests INT NOT NULL,
    device_id UUID REFERENCES server_devices(id) ON DELETE CASCADE
);
