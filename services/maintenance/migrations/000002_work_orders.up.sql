-- Migration 002: Work orders cards

CREATE TABLE IF NOT EXISTS work_orders (
    id                VARCHAR(64)  PRIMARY KEY,
    maintenance_id    VARCHAR(64)  NOT NULL REFERENCES maintenance(id),
    work_order_number VARCHAR(64)  NOT NULL UNIQUE,
    scheduled_date    TIMESTAMPTZ  NOT NULL,
    actual_start_date TIMESTAMPTZ,
    actual_end_date   TIMESTAMPTZ,
    completed_by      VARCHAR(64),
    estimated_hours   FLOAT        NOT NULL DEFAULT 0.0,
    actual_hours      FLOAT        NOT NULL DEFAULT 0.0
);

CREATE INDEX idx_workorders_maintenance ON work_orders (maintenance_id);
