-- Migration Phase 2: Workflow Instance trackers

CREATE TABLE IF NOT EXISTS workflow_instances (
    id VARCHAR(64) PRIMARY KEY,
    definition_id VARCHAR(64) REFERENCES workflow_definitions(id) ON DELETE CASCADE,
    version INT NOT NULL,
    current_step_id VARCHAR(64),
    state VARCHAR(64) NOT NULL,
    variables JSONB NOT NULL
);
