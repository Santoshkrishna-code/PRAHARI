-- Migration 003: Taxonomy categories

CREATE TABLE IF NOT EXISTS hazard_categories (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
