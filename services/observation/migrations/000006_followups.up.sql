-- Migration 006: Follow-up checks verifications

CREATE TABLE IF NOT EXISTS followups (
    id             VARCHAR(64) PRIMARY KEY,
    observation_id VARCHAR(64) NOT NULL REFERENCES observations(id),
    follower_id   VARCHAR(64) NOT NULL,
    follow_up_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    notes          TEXT        NOT NULL,
    is_passed      BOOLEAN     NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_followups_observation ON followups (observation_id);
