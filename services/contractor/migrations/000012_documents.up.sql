-- Migration 012: Contracts document storage references

CREATE TABLE IF NOT EXISTS contractor_documents (
    id            VARCHAR(64)  PRIMARY KEY,
    contractor_id VARCHAR(64)  NOT NULL REFERENCES contractor(id),
    title         VARCHAR(200) NOT NULL,
    storage_path  TEXT         NOT NULL
);

CREATE INDEX idx_docs_contractor ON contractor_documents (contractor_id);
