-- Migration 001: Core risk assessments table

CREATE TABLE IF NOT EXISTS risk_assessments (
    id                  VARCHAR(64)   PRIMARY KEY,
    risk_number         VARCHAR(64)   NOT NULL UNIQUE,
    asset_id            VARCHAR(64),
    department_id       VARCHAR(64)   NOT NULL,
    inherent_risk_score  INT           NOT NULL DEFAULT 0,
    residual_risk_score  INT           NOT NULL DEFAULT 0,
    status_code         VARCHAR(32)   NOT NULL DEFAULT 'DRAFT',
    title               VARCHAR(200)  NOT NULL,
    description        TEXT          NOT NULL,
    created_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted          BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector       TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, ''))
    ) STORED
);

CREATE INDEX idx_risk_status ON risk_assessments (status_code) WHERE is_deleted = false;
CREATE INDEX idx_risk_search ON risk_assessments USING GIN (search_vector);
