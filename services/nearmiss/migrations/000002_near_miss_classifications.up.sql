-- Migration 002: Classifications taxonomy

CREATE TABLE IF NOT EXISTS near_miss_classifications (
    id          VARCHAR(64)  PRIMARY KEY,
    code        VARCHAR(64)  NOT NULL UNIQUE,
    name        VARCHAR(200) NOT NULL,
    description TEXT
);
