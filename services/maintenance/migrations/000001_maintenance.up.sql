-- Migration 001: Core maintenance table

CREATE TABLE IF NOT EXISTS maintenance (
    id                    VARCHAR(64)   PRIMARY KEY,
    maintenance_number    VARCHAR(64)   NOT NULL UNIQUE,
    asset_id              VARCHAR(64)   NOT NULL,
    maintenance_type      VARCHAR(64)   NOT NULL,
    priority              VARCHAR(32)   NOT NULL DEFAULT 'MEDIUM',
    department_id         VARCHAR(64)   NOT NULL,
    title                 VARCHAR(200)  NOT NULL,
    description          TEXT          NOT NULL,
    status_code           VARCHAR(32)   NOT NULL DEFAULT 'DRAFT',
    total_estimated_cost  FLOAT         NOT NULL DEFAULT 0.0,
    total_actual_cost     FLOAT         NOT NULL DEFAULT 0.0,
    created_at            TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted            BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector         TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, ''))
    ) STORED
);

CREATE INDEX idx_maintenance_status      ON maintenance (status_code) WHERE is_deleted = false;
CREATE INDEX idx_maintenance_type        ON maintenance (maintenance_type) WHERE is_deleted = false;
CREATE INDEX idx_maintenance_asset       ON maintenance (asset_id) WHERE is_deleted = false;
CREATE INDEX idx_maintenance_search      ON maintenance USING GIN (search_vector);
