-- Migration 005: Requirements checklist clauses

CREATE TABLE IF NOT EXISTS requirements (
    id            VARCHAR(64)  PRIMARY KEY,
    obligation_id VARCHAR(64)  NOT NULL REFERENCES obligations(id),
    clause        VARCHAR(200) NOT NULL,
    description   TEXT
);

CREATE INDEX idx_requirements_obligation ON requirements (obligation_id);
