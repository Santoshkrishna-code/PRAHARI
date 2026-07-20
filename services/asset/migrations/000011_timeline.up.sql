-- Migration 011: Timeline milestones

CREATE TABLE IF NOT EXISTS asset_timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    asset_id    VARCHAR(64)  NOT NULL REFERENCES assets(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_asset ON asset_timeline (asset_id);
