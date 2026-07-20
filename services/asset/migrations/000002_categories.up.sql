-- Migration 002: Categories taxonomy

CREATE TABLE IF NOT EXISTS asset_categories (
    id          VARCHAR(64)   PRIMARY KEY,
    name        VARCHAR(200)  NOT NULL,
    description TEXT,
    is_active   BOOLEAN       NOT NULL DEFAULT TRUE
);

INSERT INTO asset_categories (id, name, description, is_active) VALUES
    ('cat-rotating', 'Rotating Equipment', 'Pumps, motors, compressors', true),
    ('cat-fixed',    'Fixed Equipment',    'Tanks, vessels, piping systems', true);
