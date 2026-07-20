-- Migration 005: Certifications credentials validities

CREATE TABLE IF NOT EXISTS contractor_certifications (
    id          VARCHAR(64)  PRIMARY KEY,
    worker_id   VARCHAR(64)  NOT NULL REFERENCES contractor_workers(id),
    cert_number VARCHAR(100) NOT NULL UNIQUE,
    title       VARCHAR(200) NOT NULL,
    expiry_date TIMESTAMPTZ  NOT NULL,
    issuer      VARCHAR(200) NOT NULL
);

CREATE INDEX idx_certs_worker ON contractor_certifications (worker_id);
