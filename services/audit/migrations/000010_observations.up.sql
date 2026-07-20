-- Migration 010: Observations details

CREATE TABLE IF NOT EXISTS observations (
    id          VARCHAR(64) PRIMARY KEY,
    finding_id  VARCHAR(64) NOT NULL REFERENCES findings(id),
    description TEXT        NOT NULL
);

CREATE INDEX idx_observations_finding ON observations (finding_id);
