-- Migration Phase 2: Dynamic localized messaging templates

CREATE TABLE IF NOT EXISTS templates (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    body TEXT NOT NULL
);
