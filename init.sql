CREATE SCHEMA IF NOT EXISTS devices_schema;

DROP TABLE IF EXISTS devices_schema.devices;

CREATE TABLE devices_schema.devices (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NULL,
    brand VARCHAR(255) NULL,
    state INT NULL,
    creation_time TIMESTAMP
);

-- Indexes for faster querying
CREATE INDEX IF NOT EXISTS idx_devices_brand ON devices_schema.devices(brand);
CREATE INDEX IF NOT EXISTS idx_devices_state ON devices_schema.devices(state);