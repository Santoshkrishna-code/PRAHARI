-- Migration 005: Competency matrix mapping scopes

CREATE TABLE IF NOT EXISTS competency_matrix (
    id            VARCHAR(64) PRIMARY KEY,
    role_id       VARCHAR(64) NOT NULL,
    competency_id VARCHAR(64) NOT NULL REFERENCES competencies(id)
);

CREATE INDEX idx_matrix_role ON competency_matrix (role_id);
