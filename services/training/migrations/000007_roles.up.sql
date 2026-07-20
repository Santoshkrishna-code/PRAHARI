-- Migration 007: Workforce roles taxonomy

CREATE TABLE IF NOT EXISTS roles (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
