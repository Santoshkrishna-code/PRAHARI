-- Migration 008: Documents references

CREATE TABLE IF NOT EXISTS asset_documents (
    id           VARCHAR(64)  PRIMARY KEY,
    asset_id     VARCHAR(64)  NOT NULL REFERENCES assets(id),
    title        VARCHAR(500) NOT NULL,
    doc_type     VARCHAR(100) NOT NULL,
    storage_path TEXT         NOT NULL
);

CREATE INDEX idx_documents_asset ON asset_documents (asset_id);
