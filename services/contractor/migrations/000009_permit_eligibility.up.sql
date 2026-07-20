-- Migration 009: Permit eligibility validations checks

CREATE TABLE IF NOT EXISTS contractor_permit_eligibility (
    id                   VARCHAR(64) PRIMARY KEY,
    worker_id            VARCHAR(64) NOT NULL REFERENCES contractor_workers(id),
    is_eligible          BOOLEAN     NOT NULL DEFAULT TRUE,
    ineligibility_reason TEXT
);

CREATE INDEX idx_eligibility_worker ON contractor_permit_eligibility (worker_id);
