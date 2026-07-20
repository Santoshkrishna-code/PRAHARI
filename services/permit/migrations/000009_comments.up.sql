-- Migration 009: Threaded Comments

CREATE TABLE IF NOT EXISTS permit_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    permit_id         VARCHAR(64)  NOT NULL REFERENCES permits(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES permit_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_permit ON permit_comments (permit_id);
