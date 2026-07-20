-- Migration 010: Timeline

CREATE TABLE IF NOT EXISTS inspection_timeline (
    id            VARCHAR(64)  PRIMARY KEY,
    inspection_id VARCHAR(64)  NOT NULL REFERENCES inspections(id),
    event_type    VARCHAR(64)  NOT NULL,
    actor_id      VARCHAR(64)  NOT NULL,
    description   TEXT         NOT NULL,
    metadata      JSONB,
    occurred_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_inspection ON inspection_timeline (inspection_id);
