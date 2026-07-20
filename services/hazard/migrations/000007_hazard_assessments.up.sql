-- Migration 007: Assessment parameters

CREATE TABLE IF NOT EXISTS hazard_assessments (
    id          VARCHAR(64) PRIMARY KEY,
    likelihood  INT         NOT NULL,
    consequence INT         NOT NULL
);
