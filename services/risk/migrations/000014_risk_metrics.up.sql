-- Migration 014: Process safety index metrics dashboard logs

CREATE TABLE IF NOT EXISTS risk_metrics (
    id         SERIAL PRIMARY KEY,
    metric_key VARCHAR(100) NOT NULL,
    score      NUMERIC      NOT NULL,
    logged_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
