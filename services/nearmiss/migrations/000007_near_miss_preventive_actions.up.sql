-- Migration 007: Preventive action plans

CREATE TABLE IF NOT EXISTS near_miss_preventive_actions (
    id                   VARCHAR(64) PRIMARY KEY,
    near_miss_id         VARCHAR(64) NOT NULL REFERENCES near_misses(id),
    description          TEXT        NOT NULL,
    target_date          TIMESTAMPTZ NOT NULL,
    responsible_party_id VARCHAR(64) NOT NULL,
    is_implemented       BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_preventive_nearmiss ON near_miss_preventive_actions (near_miss_id);
