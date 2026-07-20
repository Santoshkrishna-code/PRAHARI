-- Migration 006: Periodic reassessment reviews checks

CREATE TABLE IF NOT EXISTS risk_reviews (
    id               VARCHAR(64) PRIMARY KEY,
    risk_id          VARCHAR(64) NOT NULL REFERENCES risk_assessments(id),
    reviewer_id      VARCHAR(64) NOT NULL,
    review_date      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    next_review_date TIMESTAMPTZ NOT NULL,
    notes            TEXT
);

CREATE INDEX idx_reviews_risk ON risk_reviews (risk_id);
