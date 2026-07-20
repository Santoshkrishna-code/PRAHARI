-- Migration 012: Recurrent schedules

CREATE TABLE IF NOT EXISTS inspection_schedule (
    id                  VARCHAR(64)  PRIMARY KEY,
    template_id         VARCHAR(64)  NOT NULL REFERENCES inspection_templates(id),
    frequency           VARCHAR(64)  NOT NULL,
    inspector_id        VARCHAR(64)  NOT NULL,
    department_id       VARCHAR(64)  NOT NULL,
    last_execution_date TIMESTAMPTZ  NOT NULL,
    next_execution_date TIMESTAMPTZ  NOT NULL,
    is_active           BOOLEAN      NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_schedule_active ON inspection_schedule (is_active);
