-- Migration 008: Verifications checking signatures

CREATE TABLE IF NOT EXISTS near_miss_verifications (
    id            VARCHAR(64) PRIMARY KEY,
    near_miss_id  VARCHAR(64) NOT NULL REFERENCES near_misses(id),
    verifier_id   VARCHAR(64) NOT NULL,
    verified_date TIMESTAMPTZ NOT NULL,
    is_passed     BOOLEAN     NOT NULL DEFAULT FALSE,
    comments      TEXT
);

CREATE INDEX idx_verifications_nearmiss ON near_miss_verifications (near_miss_id);
