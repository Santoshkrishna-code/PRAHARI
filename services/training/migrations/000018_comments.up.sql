-- Migration 018: Comments threads

CREATE TABLE IF NOT EXISTS comments (
    id                VARCHAR(64)  PRIMARY KEY,
    training_id       VARCHAR(64)  NOT NULL REFERENCES training_programs(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_training ON comments (training_id);
