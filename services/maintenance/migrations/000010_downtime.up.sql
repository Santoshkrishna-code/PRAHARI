-- Migration 010: Downtime records logs

CREATE TABLE IF NOT EXISTS maintenance_downtime (
    id             VARCHAR(64) PRIMARY KEY,
    maintenance_id VARCHAR(64) NOT NULL REFERENCES maintenance(id),
    asset_id       VARCHAR(64) NOT NULL,
    start_date     TIMESTAMPTZ NOT NULL,
    end_date       TIMESTAMPTZ NOT NULL,
    reason_code    VARCHAR(64) NOT NULL
);

CREATE INDEX idx_downtime_asset ON maintenance_downtime (asset_id);
