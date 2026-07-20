-- Migration 007: Gas Tests

CREATE TABLE IF NOT EXISTS permit_gas_tests (
    id                         VARCHAR(64)  PRIMARY KEY,
    permit_id                  VARCHAR(64)  NOT NULL REFERENCES permits(id),
    gas_type                   VARCHAR(32)  NOT NULL,
    reading_value              FLOAT        NOT NULL,
    unit                       VARCHAR(32)  NOT NULL,
    acceptable_min             FLOAT        NOT NULL,
    acceptable_max             FLOAT        NOT NULL,
    is_passed                  BOOLEAN      NOT NULL,
    tested_by                  VARCHAR(64)  NOT NULL,
    tested_at                  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    equipment_calibration_date TIMESTAMPTZ  NOT NULL,
    next_test_due              TIMESTAMPTZ  NOT NULL
);

CREATE INDEX idx_gas_tests_permit ON permit_gas_tests (permit_id);
