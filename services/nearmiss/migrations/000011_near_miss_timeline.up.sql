-- Migration 011: Timeline milestones events

CREATE TABLE IF NOT EXISTS near_miss_timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    near_miss_id VARCHAR(64) NOT NULL REFERENCES near_misses(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_nearmiss ON near_miss_timeline (near_miss_id);
