-- Migration 006: Safety observations

CREATE TABLE IF NOT EXISTS inspection_observations (
    id            VARCHAR(64) PRIMARY KEY,
    inspection_id VARCHAR(64) NOT NULL REFERENCES inspections(id),
    description   TEXT        NOT NULL,
    observer_id   VARCHAR(64) NOT NULL,
    observed_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_observations_inspection ON inspection_observations (inspection_id);
