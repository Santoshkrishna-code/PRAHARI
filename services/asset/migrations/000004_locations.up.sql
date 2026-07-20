-- Migration 004: Hierarchical plant locations

CREATE TABLE IF NOT EXISTS asset_locations (
    id          VARCHAR(64)   PRIMARY KEY,
    parent_id   VARCHAR(64)   REFERENCES asset_locations(id),
    name        VARCHAR(200)  NOT NULL,
    description TEXT,
    is_active   BOOLEAN       NOT NULL DEFAULT TRUE
);

INSERT INTO asset_locations (id, parent_id, name, description, is_active) VALUES
    ('loc-plant-1', NULL,          'Refinery Plant 1', 'Main production plant', true),
    ('loc-bay-a',   'loc-plant-1', 'Maintenance Bay A', 'Pump repairs deck bay', true);
