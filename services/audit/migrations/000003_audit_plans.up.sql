-- Migration 003: Audit plans timeline schedules

CREATE TABLE IF NOT EXISTS audit_plans (
    id               VARCHAR(64) PRIMARY KEY,
    audit_program_id VARCHAR(64) NOT NULL REFERENCES audit_programs(id),
    scheduled_start  TIMESTAMPTZ NOT NULL,
    scheduled_end    TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_plans_program ON audit_plans (audit_program_id);
