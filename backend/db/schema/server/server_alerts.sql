CREATE TABLE alerts (
    id SERIAL PRIMARY KEY,
    host VARCHAR(45) NOT NULL,
    severity TEXT NOT NULL CHECK (severity IN ('info', 'warning', 'critical')),
    content TEXT NOT NULL,
    time TIMESTAMPTZ DEFAULT now()
);
