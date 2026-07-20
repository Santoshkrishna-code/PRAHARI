-- Migration 017: Certification renewals schedules

CREATE TABLE IF NOT EXISTS renewals (
    id               VARCHAR(64) PRIMARY KEY,
    certification_id VARCHAR(64) NOT NULL REFERENCES certifications(id),
    scheduled_date   TIMESTAMPTZ NOT NULL,
    is_completed     BOOLEAN     NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_renew_cert ON renewals (certification_id);
