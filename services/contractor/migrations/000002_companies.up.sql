-- Migration 002: Contractor Companies

CREATE TABLE IF NOT EXISTS contractor_companies (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT,
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE
);
