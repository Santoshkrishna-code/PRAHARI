-- Migration 004: Root causes mappings

CREATE TABLE IF NOT EXISTS near_miss_causes (
    id          VARCHAR(64) PRIMARY KEY,
    near_miss_id VARCHAR(64) NOT NULL REFERENCES near_misses(id),
    root_cause  VARCHAR(64) NOT NULL,
    description TEXT
);

CREATE INDEX idx_causes_nearmiss ON near_miss_causes (near_miss_id);
