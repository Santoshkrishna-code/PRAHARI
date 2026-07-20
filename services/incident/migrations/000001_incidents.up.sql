-- Migration 001: Core incidents table
-- The primary aggregate table for the Incident Management bounded context.

CREATE TABLE IF NOT EXISTS incidents (
    id              VARCHAR(64)   PRIMARY KEY,
    incident_number VARCHAR(64)   NOT NULL UNIQUE,
    title           VARCHAR(500)  NOT NULL,
    description     TEXT          NOT NULL,
    type            VARCHAR(32)   NOT NULL,
    category_id     VARCHAR(64)   NOT NULL,
    severity_level  VARCHAR(32)   NOT NULL,
    priority_level  VARCHAR(32)   NOT NULL,
    status_code     VARCHAR(32)   NOT NULL DEFAULT 'DRAFT',
    reporter_id     VARCHAR(64)   NOT NULL,
    assignee_id     VARCHAR(64),
    department_id   VARCHAR(64)   NOT NULL,
    location_id     VARCHAR(64)   NOT NULL,
    location_detail TEXT,
    occurred_at     TIMESTAMPTZ   NOT NULL,
    reported_at     TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    resolved_at     TIMESTAMPTZ,
    closed_at       TIMESTAMPTZ,
    created_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted      BOOLEAN       NOT NULL DEFAULT FALSE,

    -- Full-text search vector for title + description
    search_vector   TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, ''))
    ) STORED
);

-- Performance indexes
CREATE INDEX idx_incidents_status      ON incidents (status_code) WHERE is_deleted = false;
CREATE INDEX idx_incidents_severity    ON incidents (severity_level) WHERE is_deleted = false;
CREATE INDEX idx_incidents_category    ON incidents (category_id) WHERE is_deleted = false;
CREATE INDEX idx_incidents_reporter    ON incidents (reporter_id) WHERE is_deleted = false;
CREATE INDEX idx_incidents_assignee    ON incidents (assignee_id) WHERE is_deleted = false;
CREATE INDEX idx_incidents_department  ON incidents (department_id) WHERE is_deleted = false;
CREATE INDEX idx_incidents_occurred    ON incidents (occurred_at) WHERE is_deleted = false;
CREATE INDEX idx_incidents_number      ON incidents (incident_number);

-- GIN index for full-text search
CREATE INDEX idx_incidents_search      ON incidents USING GIN (search_vector);
