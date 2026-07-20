-- Migration 014: Course assessments evaluations

CREATE TABLE IF NOT EXISTS assessments (
    id          VARCHAR(64) PRIMARY KEY,
    training_id VARCHAR(64) NOT NULL REFERENCES training_programs(id),
    trainee_id  VARCHAR(64) NOT NULL,
    score       INT         NOT NULL DEFAULT 0,
    is_passed   BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_assess_training ON assessments (training_id);
