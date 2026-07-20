-- Migration 020: Timeline milestone events history log

CREATE TABLE IF NOT EXISTS timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    training_id VARCHAR(64)  NOT NULL REFERENCES training_programs(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_training ON timeline (training_id);
