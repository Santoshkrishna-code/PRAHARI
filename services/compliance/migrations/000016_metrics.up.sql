-- Migration 016: Compliance scoring indexes dashboard logs

CREATE TABLE IF NOT EXISTS metrics (
    id         SERIAL PRIMARY KEY,
    metric_key VARCHAR(100) NOT NULL,
    score      NUMERIC      NOT NULL,
    logged_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
