-- Migration 008: Trainees certifications trackers

CREATE TABLE IF NOT EXISTS certifications (
    id         VARCHAR(64)  PRIMARY KEY,
    trainee_id VARCHAR(64)  NOT NULL,
    course_id  VARCHAR(64)  NOT NULL REFERENCES courses(id),
    issuer     VARCHAR(200) NOT NULL,
    valid_until TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_certs_trainee ON certifications (trainee_id);
