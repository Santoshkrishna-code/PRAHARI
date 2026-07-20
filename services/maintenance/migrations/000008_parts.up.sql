-- Migration 008: Spare parts usage logs

CREATE TABLE IF NOT EXISTS maintenance_spare_parts (
    id             VARCHAR(64)  PRIMARY KEY,
    maintenance_id VARCHAR(64)  NOT NULL REFERENCES maintenance(id),
    part_number    VARCHAR(100) NOT NULL,
    quantity_used  INT          NOT NULL,
    unit_cost      FLOAT        NOT NULL DEFAULT 0.0
);

CREATE INDEX idx_spareparts_maint ON maintenance_spare_parts (maintenance_id);
