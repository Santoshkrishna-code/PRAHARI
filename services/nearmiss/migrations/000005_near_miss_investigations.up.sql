-- Migration 005: Root cause investigations

CREATE TABLE IF NOT EXISTS near_miss_investigations (
    id                 VARCHAR(64) PRIMARY KEY,
    near_miss_id       VARCHAR(64) NOT NULL REFERENCES near_misses(id),
    lead_investigator_id VARCHAR(64) NOT NULL,
    investigation_date TIMESTAMPTZ NOT NULL,
    findings           TEXT        NOT NULL,
    methodology        VARCHAR(64) NOT NULL
);

CREATE INDEX idx_investigations_nearmiss ON near_miss_investigations (near_miss_id);
