-- Migration 002: Hierarchical incident categories

CREATE TABLE IF NOT EXISTS incident_categories (
    id          VARCHAR(64)   PRIMARY KEY,
    name        VARCHAR(200)  NOT NULL,
    description TEXT,
    parent_id   VARCHAR(64)   REFERENCES incident_categories(id),
    sort_order  INT           NOT NULL DEFAULT 0,
    is_active   BOOLEAN       NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_categories_parent ON incident_categories (parent_id);
CREATE INDEX idx_categories_active ON incident_categories (is_active);

-- Seed default categories
INSERT INTO incident_categories (id, name, description, sort_order, is_active) VALUES
    ('cat-fire',       'Fire',              'Fire-related incidents including explosions',    1, true),
    ('cat-chemical',   'Chemical Spill',    'Hazardous chemical releases and contamination',  2, true),
    ('cat-electrical', 'Electrical',        'Electrical hazards including arc flash',          3, true),
    ('cat-fall',       'Fall',              'Falls from height or same-level slips',           4, true),
    ('cat-equipment',  'Equipment Failure', 'Mechanical or equipment malfunction events',     5, true),
    ('cat-enviro',     'Environmental',     'Environmental contamination events',              6, true),
    ('cat-structural', 'Structural',        'Structural integrity failures',                   7, true),
    ('cat-transport',  'Transportation',    'Vehicle and transportation incidents',             8, true),
    ('cat-biological', 'Biological',        'Biological hazard exposures',                     9, true),
    ('cat-ergonomic',  'Ergonomic',         'Repetitive strain and ergonomic injuries',       10, true),
    ('cat-other',      'Other',             'Incidents not covered by standard categories',   11, true);
