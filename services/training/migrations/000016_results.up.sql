-- Migration 016: Pass/fail results scores

CREATE TABLE IF NOT EXISTS results (
    id             VARCHAR(64) PRIMARY KEY,
    examination_id VARCHAR(64) NOT NULL REFERENCES examinations(id),
    trainee_id     VARCHAR(64) NOT NULL,
    score          INT         NOT NULL DEFAULT 0,
    is_passed      BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_results_exam ON results (examination_id);
