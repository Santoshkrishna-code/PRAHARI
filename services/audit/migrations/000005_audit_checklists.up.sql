-- Migration 005: Audit checklist frameworks catalog

CREATE TABLE IF NOT EXISTS audit_checklists (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
