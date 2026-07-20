-- Migration 010: Company liability insurance validations

CREATE TABLE IF NOT EXISTS contractor_insurance (
    id            VARCHAR(64)  PRIMARY KEY,
    contractor_id VARCHAR(64)  NOT NULL REFERENCES contractor(id),
    policy_number VARCHAR(100) NOT NULL,
    expiry_date   TIMESTAMPTZ  NOT NULL,
    limit_amount  FLOAT        NOT NULL DEFAULT 0.0
);

CREATE INDEX idx_insurance_contractor ON contractor_insurance (contractor_id);
