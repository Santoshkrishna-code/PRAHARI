-- Migration 007: Assigned Technicians mappings

CREATE TABLE IF NOT EXISTS maintenance_technicians (
    id             VARCHAR(64) PRIMARY KEY,
    maintenance_id VARCHAR(64) NOT NULL REFERENCES maintenance(id),
    user_id        VARCHAR(64) NOT NULL,
    lead_tech      BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_technicians_maint ON maintenance_technicians (maintenance_id);
