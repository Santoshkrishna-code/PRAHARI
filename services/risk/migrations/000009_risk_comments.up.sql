-- Migration 009: Risk Comments thread

CREATE TABLE IF NOT EXISTS risk_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    risk_id           VARCHAR(64)  NOT NULL REFERENCES risk_assessments(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES risk_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_risk_comments_id ON risk_comments (risk_id);
