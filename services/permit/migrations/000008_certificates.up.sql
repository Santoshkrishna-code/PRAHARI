-- Migration 008: Digital Certificates

CREATE TABLE IF NOT EXISTS permit_certificates (
    id               VARCHAR(64)  PRIMARY KEY,
    permit_id        VARCHAR(64)  NOT NULL REFERENCES permits(id),
    type             VARCHAR(64)  NOT NULL,
    issued_to        VARCHAR(64)  NOT NULL,
    issued_by        VARCHAR(64)  NOT NULL,
    signature_hash   VARCHAR(256) NOT NULL,
    issued_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    expires_at       TIMESTAMPTZ,
    revoked_at       TIMESTAMPTZ,
    certificate_data JSONB        NOT NULL
);

CREATE INDEX idx_certificates_permit ON permit_certificates (permit_id);
