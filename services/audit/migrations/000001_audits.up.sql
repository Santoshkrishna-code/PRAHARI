-- Migration 001: Core audits table

CREATE TABLE IF NOT EXISTS audits (
    id                VARCHAR(64)   PRIMARY KEY,
    audit_number      VARCHAR(64)   NOT NULL UNIQUE,
    asset_id          VARCHAR(64),
    department_id     VARCHAR(64)   NOT NULL,
    compliance_rating NUMERIC       NOT NULL DEFAULT 100.0,
    status_code       VARCHAR(32)   NOT NULL DEFAULT 'DRAFT',
    title             VARCHAR(200)  NOT NULL,
    description      TEXT          NOT NULL,
    created_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector     TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, ''))
    ) STORED
);

CREATE INDEX idx_audit_status ON audits (status_code) WHERE is_deleted = false;
CREATE INDEX idx_audit_search ON audits USING GIN (search_vector);
