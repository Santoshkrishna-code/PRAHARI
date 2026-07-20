-- Migration 008: Verification checks logs

CREATE TABLE IF NOT EXISTS hazard_verifications (
    id                  VARCHAR(64) PRIMARY KEY,
    hazard_id           VARCHAR(64) NOT NULL REFERENCES hazards(id),
    verifier_id         VARCHAR(64) NOT NULL,
    verified_date       TIMESTAMPTZ NOT NULL,
    residual_risk_score INT         NOT NULL,
    comments            TEXT
);

CREATE INDEX idx_verifications_hazard ON hazard_verifications (hazard_id);
