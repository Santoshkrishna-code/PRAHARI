-- Migration 002: Courses catalog registry

CREATE TABLE IF NOT EXISTS courses (
    id             VARCHAR(64)  PRIMARY KEY,
    course_code    VARCHAR(64)  NOT NULL UNIQUE,
    title          VARCHAR(200) NOT NULL,
    duration_hours INT          NOT NULL DEFAULT 1
);
