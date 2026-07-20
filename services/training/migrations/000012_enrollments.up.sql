-- Migration 012: Trainees course enrollments

CREATE TABLE IF NOT EXISTS enrollments (
    id          VARCHAR(64) PRIMARY KEY,
    training_id VARCHAR(64) NOT NULL REFERENCES training_programs(id),
    trainee_id  VARCHAR(64) NOT NULL,
    status      VARCHAR(32) NOT NULL DEFAULT 'ENROLLED'
);

CREATE INDEX idx_enroll_training ON enrollments (training_id);
