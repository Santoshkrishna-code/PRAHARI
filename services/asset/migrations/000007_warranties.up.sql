-- Migration 007: Warranties contracts

CREATE TABLE IF NOT EXISTS asset_warranties (
    id              VARCHAR(64)  PRIMARY KEY,
    asset_id        VARCHAR(64)  NOT NULL REFERENCES assets(id),
    start_date      TIMESTAMPTZ  NOT NULL,
    end_date        TIMESTAMPTZ  NOT NULL,
    coverage_detail TEXT         NOT NULL,
    contact_person  VARCHAR(200) NOT NULL,
    contact_email   VARCHAR(200) NOT NULL
);

CREATE INDEX idx_warranties_asset ON asset_warranties (asset_id);
