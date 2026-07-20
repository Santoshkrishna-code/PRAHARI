-- Migration 009: Workforce licenses permissions

CREATE TABLE IF NOT EXISTS licenses (
    id         VARCHAR(64)  PRIMARY KEY,
    trainee_id VARCHAR(64)  NOT NULL,
    issuer     VARCHAR(200) NOT NULL,
    valid_until TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_licenses_trainee ON licenses (trainee_id);
