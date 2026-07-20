-- Migration 006: Skills definitions

CREATE TABLE IF NOT EXISTS skills (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
