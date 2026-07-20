-- Migration 002: Observation Types classifications

CREATE TABLE IF NOT EXISTS observation_types (
    id          VARCHAR(64)  PRIMARY KEY,
    code        VARCHAR(64)  NOT NULL UNIQUE,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
