-- Migration 006: Collected evidence references

CREATE TABLE IF NOT EXISTS evidence (
    id             VARCHAR(64) PRIMARY KEY,
    obligation_id  VARCHAR(64) NOT NULL REFERENCES obligations(id),
    uploaded_by_id VARCHAR(64) NOT NULL,
    storage_path   TEXT        NOT NULL,
    collected_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_evidence_obligation ON evidence (obligation_id);
