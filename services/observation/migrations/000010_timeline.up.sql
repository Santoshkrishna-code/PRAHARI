-- Migration 010: Timeline milestone event records

CREATE TABLE IF NOT EXISTS timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    observation_id VARCHAR(64) NOT NULL REFERENCES observations(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_observation ON timeline (observation_id);
