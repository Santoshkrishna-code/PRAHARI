-- Migration 011: Non-Conformities (NCR) safety violations

CREATE TABLE IF NOT EXISTS nonconformities (
    id          VARCHAR(64) PRIMARY KEY,
    finding_id  VARCHAR(64) NOT NULL REFERENCES findings(id),
    severity    VARCHAR(32) NOT NULL DEFAULT 'MINOR',
    description TEXT        NOT NULL
);

CREATE INDEX idx_noncon_finding ON nonconformities (finding_id);
