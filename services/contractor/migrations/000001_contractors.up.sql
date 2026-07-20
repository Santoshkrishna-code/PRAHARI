-- Migration 001: Core contractor table

CREATE TABLE IF NOT EXISTS contractor (
    id                VARCHAR(64)   PRIMARY KEY,
    contractor_number VARCHAR(64)   NOT NULL UNIQUE,
    company_name      VARCHAR(200)  NOT NULL,
    tax_id            VARCHAR(64)   NOT NULL,
    status_code       VARCHAR(32)   NOT NULL DEFAULT 'REGISTERED',
    department_id     VARCHAR(64)   NOT NULL,
    registration_date TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    insurance_expiry  TIMESTAMPTZ   NOT NULL,
    created_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector     TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(company_name, ''))
    ) STORED
);

CREATE INDEX idx_contractor_status  ON contractor (status_code) WHERE is_deleted = false;
CREATE INDEX idx_contractor_search  ON contractor USING GIN (search_vector);
