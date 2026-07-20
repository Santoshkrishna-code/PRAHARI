-- Migration 009: Comment thread

CREATE TABLE IF NOT EXISTS near_miss_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    near_miss_id      VARCHAR(64)  NOT NULL REFERENCES near_misses(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES near_miss_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_nearmiss ON near_miss_comments (near_miss_id);
