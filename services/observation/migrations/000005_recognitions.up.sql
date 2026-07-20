-- Migration 005: Positive rewards recognitions logs

CREATE TABLE IF NOT EXISTS recognitions (
    id                 VARCHAR(64) PRIMARY KEY,
    observation_id      VARCHAR(64) NOT NULL REFERENCES observations(id),
    recognized_person_id VARCHAR(64) NOT NULL,
    granted_by_id       VARCHAR(64) NOT NULL,
    reason             TEXT        NOT NULL,
    granted_at         TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_recognitions_observation ON recognitions (observation_id);
