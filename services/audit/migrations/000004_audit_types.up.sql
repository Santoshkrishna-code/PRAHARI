-- Migration 004: Audit types classifications catalog

CREATE TABLE IF NOT EXISTS audit_types (
    id          VARCHAR(64)  PRIMARY KEY,
    code        VARCHAR(64)  NOT NULL UNIQUE,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
