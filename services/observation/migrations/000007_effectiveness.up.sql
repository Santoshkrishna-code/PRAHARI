-- Migration 007: Coaching effectiveness evaluations

CREATE TABLE IF NOT EXISTS effectiveness (
    id              VARCHAR(64) PRIMARY KEY,
    observation_id  VARCHAR(64) NOT NULL REFERENCES observations(id),
    evaluator_id    VARCHAR(64) NOT NULL,
    evaluation_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    improvement_rate INT        NOT NULL DEFAULT 1,
    notes           TEXT
);

CREATE INDEX idx_effectiveness_observation ON effectiveness (observation_id);
