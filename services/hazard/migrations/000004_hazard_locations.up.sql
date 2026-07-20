-- Migration 004: Location mappings references

CREATE TABLE IF NOT EXISTS hazard_locations (
    id          VARCHAR(64)  PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    facility_id VARCHAR(64)  NOT NULL
);
