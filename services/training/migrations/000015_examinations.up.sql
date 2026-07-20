-- Migration 015: Written examinations topics

CREATE TABLE IF NOT EXISTS examinations (
    id          VARCHAR(64) PRIMARY KEY,
    training_id VARCHAR(64) NOT NULL REFERENCES training_programs(id),
    topic       VARCHAR(200) NOT NULL
);

CREATE INDEX idx_exam_training ON examinations (training_id);
