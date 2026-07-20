-- Migration 004: Worker competencies ratings

CREATE TABLE IF NOT EXISTS contractor_competencies (
    id          VARCHAR(64) PRIMARY KEY,
    worker_id   VARCHAR(64) NOT NULL REFERENCES contractor_workers(id),
    skill_name  VARCHAR(100) NOT NULL,
    skill_level VARCHAR(64) NOT NULL
);

CREATE INDEX idx_competencies_worker ON contractor_competencies (worker_id);
