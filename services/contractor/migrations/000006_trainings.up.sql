-- Migration 006: Safety courses training compliance logs

CREATE TABLE IF NOT EXISTS contractor_trainings (
    id           VARCHAR(64)  PRIMARY KEY,
    worker_id    VARCHAR(64)  NOT NULL REFERENCES contractor_workers(id),
    course_name  VARCHAR(200) NOT NULL,
    completed_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    expiry_date  TIMESTAMPTZ  NOT NULL
);

CREATE INDEX idx_trainings_worker ON contractor_trainings (worker_id);
