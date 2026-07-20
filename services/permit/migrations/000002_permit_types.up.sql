-- Migration 002: Permit templates and seeded items

CREATE TABLE IF NOT EXISTS permit_types (
    id                     VARCHAR(64)   PRIMARY KEY,
    code                   VARCHAR(64)   NOT NULL UNIQUE,
    name                   VARCHAR(200)  NOT NULL,
    description            TEXT,
    default_duration_hours INT           NOT NULL DEFAULT 8,
    preconditions          JSONB,
    is_active              BOOLEAN       NOT NULL DEFAULT TRUE
);

INSERT INTO permit_types (id, code, name, description, default_duration_hours, preconditions, is_active) VALUES
    ('pt-hot',   'HOT_WORK',          'Hot Work Permit',          'Work involving open flames or sparks', 8, '{"gas_test_required": true, "loto_required": false}', true),
    ('pt-cold',  'COLD_WORK',         'Cold Work Permit',         'General mechanical maintenance',       8, '{"gas_test_required": false, "loto_required": false}', true),
    ('pt-conf',  'CONFINED_SPACE',    'Confined Space Permit',    'Enclosed vessel entry',               8, '{"gas_test_required": true, "loto_required": true}', true),
    ('pt-elec',  'ELECTRICAL',        'Electrical Permit',        'High-voltage circuit switching',      8, '{"gas_test_required": false, "loto_required": true}', true),
    ('pt-mech',  'MECHANICAL',        'Mechanical Permit',        'Rotating equipment overhaul',         8, '{"gas_test_required": false, "loto_required": true}', true),
    ('pt-exca',  'EXCAVATION',        'Excavation Permit',        'Trenching and ground breaking',       8, '{"gas_test_required": true, "loto_required": false}', true),
    ('pt-height','WORKING_AT_HEIGHT', 'Working at Height Permit', 'Scaffolding or boom lift work',        8, '{"gas_test_required": false, "loto_required": false}', true),
    ('pt-rad',   'RADIATION',         'Radiation Permit',         'X-Ray radiography testing',            8, '{"gas_test_required": false, "loto_required": false}', true),
    ('pt-chem',  'CHEMICAL',          'Chemical Permit',          'Hazardous substance handling',        8, '{"gas_test_required": true, "loto_required": false}', true),
    ('pt-line',  'LINE_BREAKING',     'Line Breaking Permit',     'Opening hydrocarbon piping',          8, '{"gas_test_required": true, "loto_required": true}', true),
    ('pt-isol',  'ISOLATION',         'Isolation Permit',         'Pre-work energy lockout checks',      8, '{"gas_test_required": false, "loto_required": true}', true),
    ('pt-maint', 'MAINTENANCE',       'Maintenance Permit',       'Routine non-hazardous servicing',      8, '{"gas_test_required": false, "loto_required": false}', true);
