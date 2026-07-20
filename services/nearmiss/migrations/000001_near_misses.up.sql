-- Migration 001: Core near_misses table

CREATE TABLE IF NOT EXISTS near_misses (
    id             VARCHAR(64)   PRIMARY KEY,
    near_miss_number VARCHAR(64) NOT NULL UNIQUE,
    asset_id       VARCHAR(64),
    contractor_id  VARCHAR(64),
    classification VARCHAR(64)   NOT NULL,
    severity_level VARCHAR(64)   NOT NULL,
    status_code    VARCHAR(32)   NOT NULL DEFAULT 'REPORTED',
    department_id  VARCHAR(64)   NOT NULL,
    title          VARCHAR(200)  NOT NULL,
    description    TEXT          NOT NULL,
    created_at     TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted     BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector  TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, ''))
    ) STORED
);

CREATE INDEX idx_nearmiss_status ON near_misses (status_code) WHERE is_deleted = false;
CREATE INDEX idx_nearmiss_search ON near_misses USING GIN (search_vector);
