-- Migration 003: Recurring Plans

CREATE TABLE IF NOT EXISTS maintenance_plans (
    id            VARCHAR(64)  PRIMARY KEY,
    asset_id      VARCHAR(64)  NOT NULL,
    title         VARCHAR(200) NOT NULL,
    interval_code VARCHAR(64)  NOT NULL,
    is_active     BOOLEAN      NOT NULL DEFAULT TRUE,
    last_run_date TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    next_run_date TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_plans_asset ON maintenance_plans (asset_id);
