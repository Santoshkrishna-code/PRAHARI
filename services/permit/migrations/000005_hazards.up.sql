-- Migration 005: Hazards

CREATE TABLE IF NOT EXISTS permit_hazards (
    id              VARCHAR(64)  PRIMARY KEY,
    permit_id       VARCHAR(64)  NOT NULL REFERENCES permits(id),
    type            VARCHAR(64)  NOT NULL,
    description     TEXT         NOT NULL,
    control_measure TEXT         NOT NULL,
    identified_by   VARCHAR(64)  NOT NULL,
    identified_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_hazards_permit ON permit_hazards (permit_id);
