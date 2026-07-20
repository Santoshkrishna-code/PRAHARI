-- Migration 013: Risk register snapshots history backup

CREATE TABLE IF NOT EXISTS risk_history (
    id          VARCHAR(64)  PRIMARY KEY,
    risk_number VARCHAR(64)  NOT NULL,
    title       VARCHAR(200) NOT NULL,
    description TEXT,
    backed_up_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
