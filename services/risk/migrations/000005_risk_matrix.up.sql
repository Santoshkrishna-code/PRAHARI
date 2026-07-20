-- Migration 005: 5x5 dynamic configurations matrix

CREATE TABLE IF NOT EXISTS risk_matrix (
    id          VARCHAR(64) PRIMARY KEY,
    likelihood  INT         NOT NULL DEFAULT 1,
    consequence INT         NOT NULL DEFAULT 1
);
