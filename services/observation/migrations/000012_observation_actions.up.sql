-- Migration 012: Corrective Action plans tasks

CREATE TABLE IF NOT EXISTS observation_actions (
    id                   VARCHAR(64) PRIMARY KEY,
    observation_id       VARCHAR(64) NOT NULL REFERENCES observations(id),
    description          TEXT        NOT NULL,
    target_date          TIMESTAMPTZ NOT NULL,
    responsible_party_id VARCHAR(64) NOT NULL,
    is_implemented       BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_actions_observation ON observation_actions (observation_id);
