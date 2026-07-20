-- Migration 011: Timeline

CREATE TABLE IF NOT EXISTS permit_timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    permit_id   VARCHAR(64)  NOT NULL REFERENCES permits(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_permit ON permit_timeline (permit_id);
