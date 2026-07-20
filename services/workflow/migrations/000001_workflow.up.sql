-- Migration Phase 1: Workflow Definition schemas

CREATE TABLE IF NOT EXISTS workflow_definitions (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    version INT NOT NULL,
    steps JSONB NOT NULL
);
