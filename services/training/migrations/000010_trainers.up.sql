-- Migration 010: Trainers instructors mappings

CREATE TABLE IF NOT EXISTS trainers (
    id          VARCHAR(64) PRIMARY KEY,
    training_id VARCHAR(64) NOT NULL REFERENCES training_programs(id),
    user_id     VARCHAR(64) NOT NULL
);

CREATE INDEX idx_trainers_training ON trainers (training_id);
