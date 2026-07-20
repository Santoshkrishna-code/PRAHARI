CREATE TABLE IF NOT EXISTS clinics (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    address TEXT,
    contact_no VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS physicians (
    id VARCHAR(50) PRIMARY KEY,
    license_number VARCHAR(100) UNIQUE NOT NULL,
    full_name VARCHAR(200) NOT NULL,
    specialty VARCHAR(100),
    contact_email VARCHAR(100),
    clinic_id VARCHAR(50) REFERENCES clinics(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS laboratories (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    license_number VARCHAR(100) UNIQUE NOT NULL,
    contact_email VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS health_profiles (
    id VARCHAR(50) PRIMARY KEY,
    worker_id VARCHAR(50) UNIQUE NOT NULL,
    worker_type VARCHAR(20) NOT NULL,
    department_id VARCHAR(50) NOT NULL,
    clearance_status VARCHAR(50) NOT NULL,
    medical_status VARCHAR(50) NOT NULL,
    blood_type VARCHAR(10),
    date_of_birth DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN DEFAULT false
);

CREATE TABLE IF NOT EXISTS medical_records (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    record_date TIMESTAMP WITH TIME ZONE NOT NULL,
    record_type VARCHAR(50) NOT NULL,
    physician_id VARCHAR(50) REFERENCES physicians(id),
    diagnosis_code VARCHAR(50),
    clinical_notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS medical_examinations (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    exam_type VARCHAR(50) NOT NULL,
    exam_date TIMESTAMP WITH TIME ZONE NOT NULL,
    physician_id VARCHAR(50) REFERENCES physicians(id),
    clinic_id VARCHAR(50) REFERENCES clinics(id),
    vitals_bp VARCHAR(20),
    vitals_pulse INT,
    weight_kg NUMERIC(5, 2),
    height_cm NUMERIC(5, 2),
    findings TEXT,
    recommendations TEXT,
    outcome_status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS fitness_assessments (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    assessment_date TIMESTAMP WITH TIME ZONE NOT NULL,
    evaluator_id VARCHAR(50) REFERENCES physicians(id),
    result_code VARCHAR(50) NOT NULL,
    notes TEXT,
    next_review_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS medical_clearances (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    clearance_date TIMESTAMP WITH TIME ZONE NOT NULL,
    expiry_date TIMESTAMP WITH TIME ZONE NOT NULL,
    is_approved BOOLEAN DEFAULT false,
    approved_by_id VARCHAR(50),
    scope_of_work TEXT,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS medical_restrictions (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    restriction_code VARCHAR(50) NOT NULL,
    description TEXT,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    is_permanent BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS occupational_illnesses (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    incident_id VARCHAR(50),
    illness_name VARCHAR(200) NOT NULL,
    icd10_code VARCHAR(50),
    diagnosis_date DATE NOT NULL,
    severity_code VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS health_surveillance (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    program_type VARCHAR(100) NOT NULL,
    start_date DATE NOT NULL,
    next_due_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS exposure_records (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    agent_name VARCHAR(100) NOT NULL,
    exposure_level NUMERIC(10, 4) NOT NULL,
    unit_of_measure VARCHAR(20) NOT NULL,
    limit_threshold NUMERIC(10, 4) NOT NULL,
    monitoring_date TIMESTAMP WITH TIME ZONE NOT NULL,
    is_over_limit BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS vaccinations (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    vaccine_name VARCHAR(100) NOT NULL,
    dose_number INT NOT NULL,
    administered_date DATE NOT NULL,
    expiry_date DATE,
    batch_number VARCHAR(100),
    provider_name VARCHAR(200),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS rehabilitation_programs (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    program_name VARCHAR(200) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    status VARCHAR(50) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS return_to_work (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    target_return_date DATE NOT NULL,
    actual_return_date DATE,
    is_phased BOOLEAN DEFAULT false,
    weekly_target_hours INT,
    status VARCHAR(50) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS laboratory_results (
    id VARCHAR(50) PRIMARY KEY,
    exam_id VARCHAR(50) REFERENCES medical_examinations(id),
    laboratory_id VARCHAR(50) REFERENCES laboratories(id),
    test_name VARCHAR(100) NOT NULL,
    test_value VARCHAR(100) NOT NULL,
    reference_range VARCHAR(100),
    is_abnormal BOOLEAN DEFAULT false,
    test_date TIMESTAMP WITH TIME ZONE NOT NULL,
    physician_notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS appointments (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    clinic_id VARCHAR(50) REFERENCES clinics(id),
    physician_id VARCHAR(50) REFERENCES physicians(id),
    scheduled_time TIMESTAMP WITH TIME ZONE NOT NULL,
    purpose VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments (
    id VARCHAR(50) PRIMARY KEY,
    target_type VARCHAR(50) NOT NULL,
    target_id VARCHAR(50) NOT NULL,
    author_id VARCHAR(50) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS attachments (
    id VARCHAR(50) PRIMARY KEY,
    target_type VARCHAR(50) NOT NULL,
    target_id VARCHAR(50) NOT NULL,
    file_name VARCHAR(200) NOT NULL,
    file_url TEXT NOT NULL,
    uploaded_by VARCHAR(50) NOT NULL,
    uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS timeline (
    id VARCHAR(50) PRIMARY KEY,
    health_profile_id VARCHAR(50) REFERENCES health_profiles(id),
    event_type VARCHAR(100) NOT NULL,
    event_date TIMESTAMP WITH TIME ZONE NOT NULL,
    actor_id VARCHAR(50) NOT NULL,
    description TEXT,
    metadata TEXT
);

CREATE TABLE IF NOT EXISTS audit_trail (
    id VARCHAR(50) PRIMARY KEY,
    action VARCHAR(50) NOT NULL,
    resource VARCHAR(100) NOT NULL,
    resource_id VARCHAR(50) NOT NULL,
    actor_id VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    old_state TEXT,
    new_state TEXT
);

CREATE TABLE IF NOT EXISTS metrics (
    metric_key VARCHAR(100) PRIMARY KEY,
    metric_value NUMERIC(10, 4) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
