-- Migration 013: Timeline milestones event records

CREATE TABLE IF NOT EXISTS maintenance_timeline (
    id             VARCHAR(64)  PRIMARY KEY,
    maintenance_id VARCHAR(64)  NOT NULL REFERENCES maintenance(id),
    event_type     VARCHAR(64)  NOT NULL,
    actor_id       VARCHAR(64)  NOT NULL,
    description    TEXT         NOT NULL,
    metadata       JSONB,
    occurred_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_maint ON maintenance_timeline (maintenance_id);
