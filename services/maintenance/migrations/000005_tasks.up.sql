-- Migration 005: Tasks walkthrough checkpoints

CREATE TABLE IF NOT EXISTS maintenance_tasks (
    id             VARCHAR(64)  PRIMARY KEY,
    maintenance_id VARCHAR(64)  NOT NULL REFERENCES maintenance(id),
    description    TEXT         NOT NULL,
    sequence_order INT          NOT NULL,
    is_completed   BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_tasks_maintenance ON maintenance_tasks (maintenance_id);
