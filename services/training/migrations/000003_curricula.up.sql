-- Migration 003: Curricula syllabus layouts

CREATE TABLE IF NOT EXISTS curricula (
    id          VARCHAR(64) PRIMARY KEY,
    course_id   VARCHAR(64) NOT NULL REFERENCES courses(id),
    description TEXT        NOT NULL
);

CREATE INDEX idx_curricula_course ON curricula (course_id);
