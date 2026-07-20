-- Migration 011: Audit Trail

CREATE TABLE IF NOT EXISTS inspection_audit (
    id          VARCHAR(64)  PRIMARY KEY,
    entity_type VARCHAR(64)  NOT NULL,
    entity_id   VARCHAR(64)  NOT NULL,
    action      VARCHAR(32)  NOT NULL,
    actor_id    VARCHAR(64)  NOT NULL,
    old_value   JSONB,
    new_value   JSONB,
    ip_address  VARCHAR(45),
    user_agent  TEXT,
    timestamp   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_entity ON inspection_audit (entity_type, entity_id);
