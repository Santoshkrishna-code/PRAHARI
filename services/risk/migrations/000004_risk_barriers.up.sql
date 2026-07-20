-- Migration 004: Bow-tie analysis barrier blocks

CREATE TABLE IF NOT EXISTS risk_barriers (
    id           VARCHAR(64) PRIMARY KEY,
    risk_id      VARCHAR(64) NOT NULL REFERENCES risk_assessments(id),
    barrier_type VARCHAR(64) NOT NULL,
    description  TEXT        NOT NULL,
    is_assured   BOOLEAN     NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_barriers_risk ON risk_barriers (risk_id);
