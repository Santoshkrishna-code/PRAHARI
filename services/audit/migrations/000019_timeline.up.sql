-- Migration 019: Timeline milestone events history log

CREATE TABLE IF NOT EXISTS timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    audit_id    VARCHAR(64)  NOT NULL REFERENCES audits(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_audit ON timeline (audit_id);
