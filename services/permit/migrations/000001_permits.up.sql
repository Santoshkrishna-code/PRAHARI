-- Migration 001: Permits core table

CREATE TABLE IF NOT EXISTS permits (
    id                 VARCHAR(64)   PRIMARY KEY,
    permit_number      VARCHAR(64)   NOT NULL UNIQUE,
    title              VARCHAR(200)  NOT NULL,
    description        TEXT          NOT NULL,
    permit_type_id     VARCHAR(64)   NOT NULL,
    status_code        VARCHAR(32)   NOT NULL DEFAULT 'DRAFT',
    risk_level         VARCHAR(32)   NOT NULL,
    applicant_id       VARCHAR(64)   NOT NULL,
    supervisor_id      VARCHAR(64)   NOT NULL,
    issuer_id          VARCHAR(64),
    receiver_id        VARCHAR(64),
    department_id      VARCHAR(64)   NOT NULL,
    contractor_id      VARCHAR(64),
    work_area_id       VARCHAR(64)   NOT NULL,
    work_description   TEXT          NOT NULL,
    planned_start_at   TIMESTAMPTZ   NOT NULL,
    planned_end_at     TIMESTAMPTZ   NOT NULL,
    actual_start_at    TIMESTAMPTZ,
    actual_end_at      TIMESTAMPTZ,
    valid_until        TIMESTAMPTZ,
    extension_count    INT           NOT NULL DEFAULT 0,
    linked_incident_id VARCHAR(64),
    created_at         TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    is_deleted         BOOLEAN       NOT NULL DEFAULT FALSE,

    search_vector      TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', coalesce(title, '') || ' ' || coalesce(description, '') || ' ' || coalesce(work_description, ''))
    ) STORED
);

CREATE INDEX idx_permits_status      ON permits (status_code) WHERE is_deleted = false;
CREATE INDEX idx_permits_type        ON permits (permit_type_id) WHERE is_deleted = false;
CREATE INDEX idx_permits_applicant   ON permits (applicant_id) WHERE is_deleted = false;
CREATE INDEX idx_permits_supervisor  ON permits (supervisor_id) WHERE is_deleted = false;
CREATE INDEX idx_permits_department  ON permits (department_id) WHERE is_deleted = false;
CREATE INDEX idx_permits_work_area   ON permits (work_area_id) WHERE is_deleted = false;
CREATE INDEX idx_permits_valid_until ON permits (valid_until) WHERE is_deleted = false;
CREATE INDEX idx_permits_search      ON permits USING GIN (search_vector);
