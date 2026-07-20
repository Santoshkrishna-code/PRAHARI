-- Migration 011: Verification checklists reviews

CREATE TABLE IF NOT EXISTS reviews (
    id               VARCHAR(64) PRIMARY KEY,
    compliance_id    VARCHAR(64) NOT NULL REFERENCES compliance_register(id),
    reviewer_id      VARCHAR(64) NOT NULL,
    review_date      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    next_review_date TIMESTAMPTZ NOT NULL,
    notes            TEXT
);

CREATE INDEX idx_reviews_compliance ON reviews (compliance_id);
