-- Migration 002: Obligations checklist entries

CREATE TABLE IF NOT EXISTS obligations (
    id             VARCHAR(64) PRIMARY KEY,
    compliance_id  VARCHAR(64) NOT NULL REFERENCES compliance_register(id),
    regulation_id  VARCHAR(64) NOT NULL,
    standard_id    VARCHAR(64) NOT NULL,
    due_date       TIMESTAMPTZ NOT NULL,
    expiration_date TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_obligations_register ON obligations (compliance_id);
