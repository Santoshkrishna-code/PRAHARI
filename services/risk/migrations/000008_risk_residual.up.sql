-- Migration 008: Residual post-mitigation risk scores

CREATE TABLE IF NOT EXISTS risk_residual (
    id                   VARCHAR(64) PRIMARY KEY,
    risk_id              VARCHAR(64) NOT NULL REFERENCES risk_assessments(id),
    residual_likelihood  INT         NOT NULL DEFAULT 1,
    residual_consequence INT         NOT NULL DEFAULT 1
);

CREATE INDEX idx_residual_risk ON risk_residual (risk_id);
