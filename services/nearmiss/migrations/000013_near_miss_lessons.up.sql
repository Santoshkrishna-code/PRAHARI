-- Migration 013: Safety lessons learned logs

CREATE TABLE IF NOT EXISTS near_miss_lessons (
    id          VARCHAR(64) PRIMARY KEY,
    near_miss_id VARCHAR(64) NOT NULL REFERENCES near_misses(id),
    summary     TEXT        NOT NULL
);
