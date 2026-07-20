-- Migration 002: Reusable templates

CREATE TABLE IF NOT EXISTS inspection_templates (
    id          VARCHAR(64)   PRIMARY KEY,
    name        VARCHAR(200)  NOT NULL,
    description TEXT,
    categories  JSONB         NOT NULL,
    items       JSONB         NOT NULL,
    is_active   BOOLEAN       NOT NULL DEFAULT TRUE
);

INSERT INTO inspection_templates (id, name, description, categories, items, is_active) VALUES
    ('temp-fire',   'Fire Extinguisher Walkthrough Audit', 'Monthly fire safety checks', '["extinguisher", "signage"]', '[]', true),
    ('temp-hazard', 'General Safety walkthrough check',   'Spill, exit blockage hazard audits', '["housekeeping", "exits"]', '[]', true);
