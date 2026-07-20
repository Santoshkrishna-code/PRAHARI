-- Migration 010: Comments threads

CREATE TABLE IF NOT EXISTS asset_comments (
    id                VARCHAR(64)  PRIMARY KEY,
    asset_id          VARCHAR(64)  NOT NULL REFERENCES assets(id),
    author_id         VARCHAR(64)  NOT NULL,
    parent_comment_id VARCHAR(64)  REFERENCES asset_comments(id),
    body              TEXT         NOT NULL,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_deleted        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_comments_asset ON asset_comments (asset_id);
