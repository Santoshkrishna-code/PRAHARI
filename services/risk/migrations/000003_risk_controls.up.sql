-- Migration 003: Control barriers checklists

CREATE TABLE IF NOT EXISTS risk_controls (
    id           VARCHAR(64) PRIMARY KEY,
    risk_id      VARCHAR(64) NOT NULL REFERENCES risk_assessments(id),
    control_type VARCHAR(64) NOT NULL,
    description  TEXT        NOT NULL,
    effect_value INT         NOT NULL DEFAULT 1
);

CREATE INDEX idx_controls_risk ON risk_controls (risk_id);
