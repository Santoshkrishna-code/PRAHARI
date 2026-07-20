-- Migration 008: Site entries gate clearances

CREATE TABLE IF NOT EXISTS contractor_site_access (
    id                VARCHAR(64)  PRIMARY KEY,
    worker_id         VARCHAR(64)  NOT NULL REFERENCES contractor_workers(id),
    allowed_locations TEXT         NOT NULL,
    badge_number      VARCHAR(100) NOT NULL UNIQUE,
    access_start      TIMESTAMPTZ  NOT NULL,
    access_end        TIMESTAMPTZ  NOT NULL
);

CREATE INDEX idx_access_worker ON contractor_site_access (worker_id);
