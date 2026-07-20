-- Migration 011: Compliance safety violations performance evaluations

CREATE TABLE IF NOT EXISTS contractor_performance (
    id               VARCHAR(64) PRIMARY KEY,
    contractor_id    VARCHAR(64) NOT NULL REFERENCES contractor(id),
    compliance_score FLOAT       NOT NULL DEFAULT 100.0,
    safety_viol      INT         NOT NULL DEFAULT 0
);

CREATE INDEX idx_performance_contractor ON contractor_performance (contractor_id);
