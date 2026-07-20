-- Migration 013: Parent-child asset structures & relationships maps

CREATE TABLE IF NOT EXISTS asset_hierarchies (
    id                VARCHAR(64) PRIMARY KEY,
    parent_asset_id   VARCHAR(64) NOT NULL REFERENCES assets(id),
    child_asset_id    VARCHAR(64) NOT NULL REFERENCES assets(id),
    relationship_type VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS asset_relationships (
    id                VARCHAR(64) PRIMARY KEY,
    source_asset_id   VARCHAR(64) NOT NULL REFERENCES assets(id),
    target_asset_id   VARCHAR(64) NOT NULL REFERENCES assets(id),
    dependency_type   VARCHAR(64) NOT NULL
);

CREATE INDEX idx_hierarchies_parent ON asset_hierarchies (parent_asset_id);
CREATE INDEX idx_relationships_src  ON asset_relationships (source_asset_id);
