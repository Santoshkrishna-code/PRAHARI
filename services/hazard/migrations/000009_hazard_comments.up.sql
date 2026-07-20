-- Migration 009: Comment thread

CREATE TABLE IF NOT EXISTS hazard_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    hazard_id         VARCHAR(64)  NOT NULL REFERENCES hazards(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES hazard_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_hazard ON hazard_comments (hazard_id);
