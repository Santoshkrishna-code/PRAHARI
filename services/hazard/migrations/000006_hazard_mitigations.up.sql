-- Migration 006: Mitigation Plans

CREATE TABLE IF NOT EXISTS hazard_mitigations (
    id                     VARCHAR(64) PRIMARY KEY,
    hazard_id              VARCHAR(64) NOT NULL REFERENCES hazards(id),
    description            TEXT        NOT NULL,
    target_completion_date TIMESTAMPTZ NOT NULL,
    responsible_party_id   VARCHAR(64) NOT NULL,
    is_implemented         BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_mitigations_hazard ON hazard_mitigations (hazard_id);
