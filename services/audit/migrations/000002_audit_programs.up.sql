-- Migration 002: Audit programs mapping

CREATE TABLE IF NOT EXISTS audit_programs (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
