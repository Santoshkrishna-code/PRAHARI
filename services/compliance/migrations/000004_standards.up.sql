-- Migration 004: Standards taxonomy framework

CREATE TABLE IF NOT EXISTS standards (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
