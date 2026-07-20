-- Migration 014: Milestones timeline event records

CREATE TABLE IF NOT EXISTS contractor_timeline (
    id            VARCHAR(64)  PRIMARY KEY,
    contractor_id VARCHAR(64)  NOT NULL REFERENCES contractor(id),
    event_type    VARCHAR(64)  NOT NULL,
    actor_id      VARCHAR(64)  NOT NULL,
    description   TEXT         NOT NULL,
    metadata      JSONB,
    occurred_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_contractor ON contractor_timeline (contractor_id);
