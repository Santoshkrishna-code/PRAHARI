-- Migration 012: Audit recommendations

CREATE TABLE IF NOT EXISTS recommendations (
    id         VARCHAR(64) PRIMARY KEY,
    finding_id VARCHAR(64) NOT NULL REFERENCES findings(id),
    suggestion TEXT        NOT NULL
);

CREATE INDEX idx_recs_finding ON recommendations (finding_id);
