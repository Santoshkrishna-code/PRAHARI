-- Migration 003: Crew workers profiles

CREATE TABLE IF NOT EXISTS contractor_workers (
    id                VARCHAR(64)  PRIMARY KEY,
    contractor_id     VARCHAR(64)  NOT NULL REFERENCES contractor(id),
    first_name        VARCHAR(100) NOT NULL,
    last_name         VARCHAR(100) NOT NULL,
    passport_id       VARCHAR(100),
    onboarding_status VARCHAR(64)  NOT NULL DEFAULT 'Pending'
);

CREATE INDEX idx_workers_contractor ON contractor_workers (contractor_id);
