-- Migration 005: Control measures hierarchies

CREATE TABLE IF NOT EXISTS hazard_controls (
    id           VARCHAR(64)  PRIMARY KEY,
    hazard_id    VARCHAR(64)  NOT NULL REFERENCES hazards(id),
    control_type VARCHAR(64)  NOT NULL,
    description  TEXT         NOT NULL,
    is_active    BOOLEAN      NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_controls_hazard ON hazard_controls (hazard_id);
