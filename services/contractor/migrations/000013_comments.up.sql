-- Migration 013: Comment threads

CREATE TABLE IF NOT EXISTS contractor_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    contractor_id     VARCHAR(64)  NOT NULL REFERENCES contractor(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES contractor_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_contractor ON contractor_comments (contractor_id);
