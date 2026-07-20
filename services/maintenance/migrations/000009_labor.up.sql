-- Migration 009: Labor tracking

CREATE TABLE IF NOT EXISTS maintenance_labor (
    id             VARCHAR(64) PRIMARY KEY,
    maintenance_id VARCHAR(64) NOT NULL REFERENCES maintenance(id),
    technician_id  VARCHAR(64) NOT NULL,
    hours_worked   FLOAT       NOT NULL DEFAULT 0.0,
    hourly_rate    FLOAT       NOT NULL DEFAULT 0.0
);

CREATE INDEX idx_labor_maint ON maintenance_labor (maintenance_id);
