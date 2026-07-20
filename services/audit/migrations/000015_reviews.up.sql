-- Migration 015: Audit reviews checklists checks

CREATE TABLE IF NOT EXISTS reviews (
    id               VARCHAR(64) PRIMARY KEY,
    audit_id         VARCHAR(64) NOT NULL REFERENCES audits(id),
    reviewer_id      VARCHAR(64) NOT NULL,
    review_date      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    next_review_date TIMESTAMPTZ NOT NULL,
    notes            TEXT
);

CREATE INDEX idx_reviews_audit ON reviews (audit_id);
