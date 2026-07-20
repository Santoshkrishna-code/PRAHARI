-- Migration 010: Tags and incident-tag join table

CREATE TABLE IF NOT EXISTS tags (
    id    VARCHAR(64)  PRIMARY KEY,
    name  VARCHAR(100) NOT NULL UNIQUE,
    color VARCHAR(7)
);

CREATE TABLE IF NOT EXISTS incident_tags (
    incident_id VARCHAR(64) NOT NULL REFERENCES incidents(id),
    tag_id      VARCHAR(64) NOT NULL REFERENCES tags(id),
    PRIMARY KEY (incident_id, tag_id)
);

CREATE INDEX idx_incident_tags_incident ON incident_tags (incident_id);
CREATE INDEX idx_incident_tags_tag      ON incident_tags (tag_id);
