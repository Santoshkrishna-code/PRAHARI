-- Migration 013: Behavioral Safety culture index logs

CREATE TABLE IF NOT EXISTS observation_metrics (
    id         SERIAL PRIMARY KEY,
    metric_key VARCHAR(100) NOT NULL,
    score      NUMERIC      NOT NULL,
    logged_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
