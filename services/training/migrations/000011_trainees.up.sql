-- Migration 011: Trainees rosters mappings

CREATE TABLE IF NOT EXISTS trainees (
    id          VARCHAR(64) PRIMARY KEY,
    training_id VARCHAR(64) NOT NULL REFERENCES training_programs(id),
    user_id     VARCHAR(64) NOT NULL
);

CREATE INDEX idx_trainees_training ON trainees (training_id);
