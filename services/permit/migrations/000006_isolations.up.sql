-- Migration 006: Isolations

CREATE TABLE IF NOT EXISTS permit_isolations (
    id                    VARCHAR(64)  PRIMARY KEY,
    permit_id             VARCHAR(64)  NOT NULL REFERENCES permits(id),
    isolation_type        VARCHAR(64)  NOT NULL,
    equipment_id          VARCHAR(64)  NOT NULL,
    equipment_description TEXT,
    isolation_point       VARCHAR(200) NOT NULL,
    lock_number           VARCHAR(64)  NOT NULL,
    tag_number            VARCHAR(64)  NOT NULL,
    isolated_by           VARCHAR(64)  NOT NULL,
    isolated_at           TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    verified_by           VARCHAR(64),
    verified_at           TIMESTAMPTZ,
    removed_by            VARCHAR(64),
    removed_at            TIMESTAMPTZ,
    status                VARCHAR(32)  NOT NULL DEFAULT 'APPLIED'
);

CREATE INDEX idx_isolations_permit ON permit_isolations (permit_id);
