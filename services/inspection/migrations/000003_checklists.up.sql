-- Migration 003: Checklists instances

CREATE TABLE IF NOT EXISTS inspection_checklists (
    id                    VARCHAR(64)  PRIMARY KEY,
    inspection_id         VARCHAR(64)  NOT NULL REFERENCES inspections(id),
    checklist_template_id VARCHAR(64)  NOT NULL REFERENCES inspection_templates(id),
    name                  VARCHAR(200) NOT NULL
);

CREATE INDEX idx_checklists_inspection ON inspection_checklists (inspection_id);
