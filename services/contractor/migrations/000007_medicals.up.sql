-- Migration 007: Worker physical clearances medical checks

CREATE TABLE IF NOT EXISTS contractor_medicals (
    id           VARCHAR(64) PRIMARY KEY,
    worker_id    VARCHAR(64) NOT NULL REFERENCES contractor_workers(id),
    evaluated_at TIMESTAMPTZ NOT NULL,
    expiry_date  TIMESTAMPTZ NOT NULL,
    is_fit       BOOLEAN     NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_medicals_worker ON contractor_medicals (worker_id);
