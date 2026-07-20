-- Migration 001: Inspections core table

CREATE TABLE IF NOT EXISTS inspections (
    id                 VARCHAR(64)   PRIMARY KEY,
    inspection_number  VARCHAR(64)   NOT NULL UNIQUE,
    title              VARCHAR(200)  NOT NULL,
    description        TEXT          NOT NULL,
    inspection_type    VARCHAR(64)   NOT NULL,
    status_code        VARCHAR(32)   NOT NULL DEFAULT 'DRAFT',
    schedule_id        VARCHAR(64),
    inspector_id       VARCHAR(64)   NOT NULL,
    department_id      VARCHAR(64)   NOT NULL,
    asset_id           VARCHAR(64),
    linked_permit_id   VARCHAR(64),
    linked_incident_id VARCHAR(64),
    compliance_score   FLOAT         NOT NULL DEFAULT 0.0,
    started_at         TIMESTAMPTZ,
    completed_at       TIMESTAMPTZ,
    created_at         TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted         BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector      TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, ''))
    ) STORED
);

CREATE INDEX idx_inspections_status      ON inspections (status_code) WHERE is_deleted = false;
CREATE INDEX idx_inspections_type        ON inspections (inspection_type) WHERE is_deleted = false;
CREATE INDEX idx_inspections_inspector   ON inspections (inspector_id) WHERE is_deleted = false;
CREATE INDEX idx_inspections_department  ON inspections (department_id) WHERE is_deleted = false;
CREATE INDEX idx_inspections_asset       ON inspections (asset_id) WHERE is_deleted = false;
CREATE INDEX idx_inspections_search      ON inspections USING GIN (search_vector);
