-- Migration 004: Scheduling bookings

CREATE TABLE IF NOT EXISTS maintenance_schedules (
    id                     VARCHAR(64) PRIMARY KEY,
    maintenance_id          VARCHAR(64) NOT NULL REFERENCES maintenance(id),
    scheduled_start_date   TIMESTAMPTZ NOT NULL,
    scheduled_end_date     TIMESTAMPTZ NOT NULL,
    estimated_downtime_min INT         NOT NULL DEFAULT 0
);

CREATE INDEX idx_schedules_maintenance ON maintenance_schedules (maintenance_id);
