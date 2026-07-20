-- Migration 005: Manufacturers

CREATE TABLE IF NOT EXISTS asset_manufacturers (
    id              VARCHAR(64)  PRIMARY KEY,
    name            VARCHAR(200) NOT NULL,
    contact_email   VARCHAR(200),
    contact_phone   VARCHAR(50),
    support_website TEXT
);

INSERT INTO asset_manufacturers (id, name, contact_email, contact_phone, support_website) VALUES
    ('mfg-siemens', 'Siemens AG', 'support@siemens.com', '+49-89-636-0', 'https://siemens.com'),
    ('mfg-flowserve', 'Flowserve Corp', 'support@flowserve.com', '+1-972-443-6500', 'https://flowserve.com');
