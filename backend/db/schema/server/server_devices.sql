CREATE TABLE IF NOT EXISTS server_devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ip VARCHAR(45) NOT NULL UNIQUE,
    tag VARCHAR(100),
    os VARCHAR(100),
    access_token VARCHAR(255) NOT NULL,  
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

