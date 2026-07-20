-- Migration 008: Certifications validity status trackers

CREATE TABLE IF NOT EXISTS certifications (
    id          VARCHAR(64)  PRIMARY KEY,
    issuer      VARCHAR(200) NOT NULL,
    valid_until TIMESTAMPTZ  NOT NULL,
    status      VARCHAR(32)  NOT NULL DEFAULT 'ACTIVE'
);
