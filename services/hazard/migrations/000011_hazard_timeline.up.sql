-- Migration 011: Timeline milestone event records

CREATE TABLE IF NOT EXISTS hazard_timeline (
    id          VARCHAR(64)  PRIMARY KEY,
    hazard_id   VARCHAR(64)  NOT NULL REFERENCES hazards(id),
    event_type  VARCHAR(64)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    description TEXT         NOT NULL,
    metadata    JSONB,
    occurred_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_timeline_hazard ON hazard_timeline (hazard_id);
