-- Migration 006: Checklists checks status

CREATE TABLE IF NOT EXISTS maintenance_checklists (
    id             VARCHAR(64)  PRIMARY KEY,
    maintenance_id VARCHAR(64)  NOT NULL REFERENCES maintenance(id),
    name           VARCHAR(200) NOT NULL,
    is_passed      BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_checklists_maint ON maintenance_checklists (maintenance_id);
