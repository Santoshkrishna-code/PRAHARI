-- Migration 003: Types specifications templates

CREATE TABLE IF NOT EXISTS asset_types (
    id          VARCHAR(64)   PRIMARY KEY,
    category_id VARCHAR(64)   NOT NULL REFERENCES asset_categories(id),
    name        VARCHAR(200)  NOT NULL,
    description TEXT,
    is_active   BOOLEAN       NOT NULL DEFAULT TRUE
);

INSERT INTO asset_types (id, category_id, name, description, is_active) VALUES
    ('type-centrifugal-pump', 'cat-rotating', 'Centrifugal Pump', 'Standard water/oil fluid pumps', true),
    ('type-pressure-vessel',  'cat-fixed',    'Pressure Vessel',  'Storage boiler tanks', true);
