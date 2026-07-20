-- Migration 013: 5x5 Matrix mapping scoring values

CREATE TABLE IF NOT EXISTS hazard_risk_matrix (
    id          VARCHAR(64) PRIMARY KEY,
    likelihood  INT         NOT NULL,
    consequence INT         NOT NULL
);
