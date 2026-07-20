-- Migration 011: Timeline milestone event records

CREATE TABLE IF NOT EXISTS risk_timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    risk_id     VARCHAR(64)  NOT NULL REFERENCES risk_assessments(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_risk_timeline_id ON risk_timeline (risk_id);
