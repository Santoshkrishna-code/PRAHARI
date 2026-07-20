-- Migration 002: Hazard Types classifications

CREATE TABLE IF NOT EXISTS hazard_types (
    id          VARCHAR(64)  PRIMARY KEY,
    code        VARCHAR(64)  NOT NULL UNIQUE,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
