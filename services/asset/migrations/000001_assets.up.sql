-- Migration 001: Assets core table

CREATE TABLE IF NOT EXISTS assets (
    id                    VARCHAR(64)   PRIMARY KEY,
    asset_number          VARCHAR(64)   NOT NULL UNIQUE,
    name                  VARCHAR(200)  NOT NULL,
    description          TEXT          NOT NULL,
    serial_number        VARCHAR(100)  NOT NULL,
    lifecycle_status      VARCHAR(64)   NOT NULL DEFAULT 'REGISTERED',
    operational_status    VARCHAR(64)   NOT NULL DEFAULT 'STOPPED',
    criticality_code      VARCHAR(32)   NOT NULL DEFAULT 'MEDIUM',
    department_id         VARCHAR(64)   NOT NULL,
    location_id           VARCHAR(64)   NOT NULL,
    category_id           VARCHAR(64)   NOT NULL,
    type_id               VARCHAR(64)   NOT NULL,
    manufacturer_id       VARCHAR(64)   NOT NULL,
    model_number          VARCHAR(100),
    purchase_date         TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    installation_date     TIMESTAMPTZ,
    last_maintenance_date TIMESTAMPTZ,
    health_score          FLOAT         NOT NULL DEFAULT 100.0,
    condition_score       FLOAT         NOT NULL DEFAULT 100.0,
    remaining_useful_life FLOAT         NOT NULL DEFAULT 87600.0,
    failure_probability   FLOAT         NOT NULL DEFAULT 0.0,
    created_at            TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted            BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector         TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(name, '') || ' ' || coalesce(description, '') || ' ' || coalesce(serial_number, ''))
    ) STORED
);

CREATE INDEX idx_assets_lifecycle      ON assets (lifecycle_status) WHERE is_deleted = false;
CREATE INDEX idx_assets_operational    ON assets (operational_status) WHERE is_deleted = false;
CREATE INDEX idx_assets_criticality    ON assets (criticality_code) WHERE is_deleted = false;
CREATE INDEX idx_assets_location       ON assets (location_id) WHERE is_deleted = false;
CREATE INDEX idx_assets_search         ON assets USING GIN (search_vector);
