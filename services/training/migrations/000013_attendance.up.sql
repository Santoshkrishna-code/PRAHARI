-- Migration 013: Attendance verifications logs

CREATE TABLE IF NOT EXISTS attendance (
    id            VARCHAR(64) PRIMARY KEY,
    training_id   VARCHAR(64) NOT NULL REFERENCES training_programs(id),
    trainee_id    VARCHAR(64) NOT NULL,
    attended_date TIMESTAMPTZ NOT NULL,
    is_present    BOOLEAN     NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_attend_training ON attendance (training_id);
