-- Migration 003: Severity levels scales

CREATE TABLE IF NOT EXISTS near_miss_severity (
    id          VARCHAR(64)  PRIMARY KEY,
    level       VARCHAR(64)  NOT NULL UNIQUE,
    score       INT          NOT NULL DEFAULT 0,
    description TEXT
);
