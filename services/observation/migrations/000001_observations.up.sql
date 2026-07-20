-- Migration 001: Core observations table

CREATE TABLE IF NOT EXISTS observations (
    id                  VARCHAR(64)   PRIMARY KEY,
    observation_number  VARCHAR(64)   NOT NULL UNIQUE,
    asset_id            VARCHAR(64),
    contractor_id       VARCHAR(64),
    observation_type    VARCHAR(64)   NOT NULL,
    status_code         VARCHAR(32)   NOT NULL DEFAULT 'RECORDED',
    department_id       VARCHAR(64)   NOT NULL,
    title               VARCHAR(200)  NOT NULL,
    description        TEXT          NOT NULL,
    created_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted          BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector       TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, ''))
    ) STORED
);

CREATE INDEX idx_observations_status ON observations (status_code) WHERE is_deleted = false;
CREATE INDEX idx_observations_search ON observations USING GIN (search_vector);
