-- Migration 004: Competencies list criteria

CREATE TABLE IF NOT EXISTS competencies (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
