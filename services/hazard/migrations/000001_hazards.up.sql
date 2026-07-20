-- Migration 001: Core hazards table

CREATE TABLE IF NOT EXISTS hazards (
    id                  VARCHAR(64)   PRIMARY KEY,
    hazard_number       VARCHAR(64)   NOT NULL UNIQUE,
    asset_id            VARCHAR(64),
    contractor_id       VARCHAR(64),
    hazard_type         VARCHAR(64)   NOT NULL,
    initial_risk_score  INT           NOT NULL DEFAULT 0,
    residual_risk_score INT           NOT NULL DEFAULT 0,
    status_code         VARCHAR(32)   NOT NULL DEFAULT 'IDENTIFIED',
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

CREATE INDEX idx_hazards_status ON hazards (status_code) WHERE is_deleted = false;
CREATE INDEX idx_hazards_search ON hazards USING GIN (search_vector);
