-- Migration 002: Risk Register catalog details

CREATE TABLE IF NOT EXISTS risk_register (
    id          VARCHAR(64)  PRIMARY KEY,
    risk_number VARCHAR(64)  NOT NULL UNIQUE,
    title       VARCHAR(200) NOT NULL,
    description TEXT
);
