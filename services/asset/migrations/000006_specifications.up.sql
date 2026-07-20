-- Migration 006: Specifications attributes

CREATE TABLE IF NOT EXISTS asset_specifications (
    id         VARCHAR(64) PRIMARY KEY,
    asset_id   VARCHAR(64) NOT NULL REFERENCES assets(id),
    attributes JSONB       NOT NULL
);

CREATE INDEX idx_specs_asset ON asset_specifications (asset_id);
