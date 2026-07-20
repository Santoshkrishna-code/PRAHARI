-- Migration Phase 3: Operational workflow tasks assignments

CREATE TABLE IF NOT EXISTS workflow_tasks (
    id VARCHAR(64) PRIMARY KEY,
    instance_id VARCHAR(64) REFERENCES workflow_instances(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(64) NOT NULL,
    assigned_to VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS workflow_approvals (
    task_id VARCHAR(64) PRIMARY KEY REFERENCES workflow_tasks(id) ON DELETE CASCADE,
    approvers TEXT[] NOT NULL,
    strategy VARCHAR(64) NOT NULL,
    decisions JSONB NOT NULL
);
