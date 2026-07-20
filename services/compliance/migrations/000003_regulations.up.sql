-- Migration 003: Regulations catalog codes

CREATE TABLE IF NOT EXISTS regulations (
    id          VARCHAR(64)  PRIMARY KEY,
    code        VARCHAR(64)  NOT NULL UNIQUE,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
