-- Migration 008: Comments

CREATE TABLE IF NOT EXISTS inspection_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    inspection_id     VARCHAR(64)  NOT NULL REFERENCES inspections(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES inspection_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_inspection ON inspection_comments (inspection_id);
