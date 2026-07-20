-- Migration Phase 3: User notification channels preferences opt-out rules

CREATE TABLE IF NOT EXISTS preferences (
    user_id VARCHAR(64) NOT NULL,
    channel VARCHAR(64) NOT NULL,
    enabled BOOLEAN DEFAULT TRUE,
    PRIMARY KEY (user_id, channel)
);
