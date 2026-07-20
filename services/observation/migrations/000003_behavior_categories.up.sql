-- Migration 003: BBS Categories taxonomy

CREATE TABLE IF NOT EXISTS behavior_categories (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
