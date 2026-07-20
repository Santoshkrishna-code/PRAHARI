-- Migration 004: Risk Assessments

CREATE TABLE IF NOT EXISTS permit_risk_assessments (
    id                VARCHAR(64)  PRIMARY KEY,
    permit_id         VARCHAR(64)  NOT NULL REFERENCES permits(id),
    assessor_id       VARCHAR(64)  NOT NULL,
    likelihood_score  INT          NOT NULL,
    consequence_score INT          NOT NULL,
    risk_score        INT          NOT NULL,
    risk_level        VARCHAR(32)  NOT NULL,
    control_measures  JSONB,
    residual_risk     VARCHAR(32)  NOT NULL,
    ppe_required      JSONB,
    assessed_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_risk_assessments_permit ON permit_risk_assessments (permit_id);
