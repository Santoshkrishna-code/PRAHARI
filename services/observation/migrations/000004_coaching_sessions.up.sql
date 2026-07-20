-- Migration 004: Coaching feedback dialogues sessions

CREATE TABLE IF NOT EXISTS coaching_sessions (
    id             VARCHAR(64) PRIMARY KEY,
    observation_id VARCHAR(64) NOT NULL REFERENCES observations(id),
    coach_id       VARCHAR(64) NOT NULL,
    session_date   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    topics         TEXT,
    feedback       TEXT
);

CREATE INDEX idx_coaching_observation ON coaching_sessions (observation_id);
