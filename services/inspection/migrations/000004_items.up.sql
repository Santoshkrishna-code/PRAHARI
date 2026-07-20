-- Migration 004: Individual question items

CREATE TABLE IF NOT EXISTS inspection_items (
    id             VARCHAR(64)  PRIMARY KEY,
    checklist_id   VARCHAR(64)  NOT NULL REFERENCES inspection_checklists(id),
    question       TEXT         NOT NULL,
    description    TEXT,
    category_name  VARCHAR(200) NOT NULL,
    response_type  VARCHAR(64)  NOT NULL,
    response_value VARCHAR(500),
    is_passed      BOOLEAN      NOT NULL DEFAULT FALSE,
    comments       TEXT
);

CREATE INDEX idx_items_checklist ON inspection_items (checklist_id);
