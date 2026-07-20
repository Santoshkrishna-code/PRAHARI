-- Migration 003: Permit approvals

CREATE TABLE IF NOT EXISTS permit_approvals (
    id             VARCHAR(64)  PRIMARY KEY,
    permit_id      VARCHAR(64)  NOT NULL REFERENCES permits(id),
    approver_id    VARCHAR(64)  NOT NULL,
    approver_role  VARCHAR(64)  NOT NULL,
    decision       VARCHAR(32)  NOT NULL DEFAULT 'PENDING',
    comments       TEXT,
    signature_hash VARCHAR(256),
    decided_at     TIMESTAMPTZ,
    sequence_order INT          NOT NULL
);

CREATE INDEX idx_approvals_permit ON permit_approvals (permit_id);
