-- Migration Phase 1: Outbound message status tracking tables

CREATE TABLE IF NOT EXISTS notifications (
    id VARCHAR(64) PRIMARY KEY,
    recipient VARCHAR(255) NOT NULL,
    channel VARCHAR(64) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(64) NOT NULL
);
