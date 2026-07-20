-- Migration 005: Findings tracking

CREATE TABLE IF NOT EXISTS inspection_findings (
    id                VARCHAR(64)  PRIMARY KEY,
    inspection_id     VARCHAR(64)  NOT NULL REFERENCES inspections(id),
    checklist_item_id VARCHAR(64)  NOT NULL,
    description       TEXT         NOT NULL,
    severity          VARCHAR(32)  NOT NULL,
    priority          VARCHAR(32)  NOT NULL,
    status            VARCHAR(32)  NOT NULL DEFAULT 'OPEN',
    identified_by     VARCHAR(64)  NOT NULL,
    identified_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_findings_inspection ON inspection_findings (inspection_id);
CREATE INDEX idx_findings_status     ON inspection_findings (status);
