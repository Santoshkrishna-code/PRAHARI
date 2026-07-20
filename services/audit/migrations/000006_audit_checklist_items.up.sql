-- Migration 006: Audit checklist item questions

CREATE TABLE IF NOT EXISTS audit_checklist_items (
    id           VARCHAR(64) PRIMARY KEY,
    checklist_id VARCHAR(64) NOT NULL REFERENCES audit_checklists(id),
    question     TEXT        NOT NULL
);

CREATE INDEX idx_items_checklist ON audit_checklist_items (checklist_id);
