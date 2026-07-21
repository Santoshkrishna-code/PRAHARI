# PRAHARI Platform: Requirements Specification

This document details the functional and non-functional requirements governing the PRAHARI Platform, supporting compliance standards for heavy industrial operations.

---

## 1. Functional Requirements

### 1.1 Process Safety Management (PSM)
- **PHA Study Management**: The platform must support creation, edition, and closure of Process Hazard Analysis (PHA) studies including HAZOP, LOPA, BowTie, FMEA, and What-If.
- **Node Segmentation**: Users must be able to decompose plant systems into nodes, define process design intents, deviations, causes, consequences, and safeguards.
- **LOPA Scenario Mapping**: For high-risk consequences, the platform must map Layer of Protection Analysis (LOPA) scenarios, calculating initiating event frequency, Independent Protection Layers (IPLs) probability of failure on demand (PFD), and final mitigated risk.
- **Pre-Start Safety Review (PSSR)**: Enable checklist-based PSSR workflows triggered by Management of Change (MOC) procedures prior to introducing hazards.

### 1.2 Chemical Lifecycle & Inventory
- **SDS Parsing**: The AI Agent must automatically parse uploaded PDF Safety Data Sheets (SDS), extracting CAS numbers, IUPAC names, GHS classifications, pictograms, exposure limits, and fire/spill handling rules.
- **Compatibility Auditing**: When placing a chemical container in a storage area, the system must cross-reference compatibility matrices and raise block/warn events if incompatible chemicals are co-located.
- **Threshold Enforcement**: The system must enforce Maximum Allowable Quantities (MAQ) for hazard categories inside fire control areas, alerting supervisors when thresholds are exceeded.

### 1.3 Real-Time Computer Vision Alerting
- **PPE Compliance**: Monitor live RTSP streams to detect lack of helmets, safety vests, boots, or protective eyewear on personnel in designated zones.
- **Intrusion Detection**: Define geofenced exclusion zones (e.g., crane swing areas, high-voltage rooms) and trigger immediate alarms upon unauthorized human presence.
- **Environmental Hazard Detection**: Detect smoke, open flame, and liquid spills on floors using camera feeds.

---

## 2. Non-Functional Requirements

### 2.1 Performance & Latency
- **API Response**: P99 latency of API Gateway must be < 200ms under standard operational load.
- **Video Ingestion**: Camera streams must be processed with an end-to-end ingestion-to-inference latency of < 100ms.
- **Alert Dispatch**: Critical life-safety alerts must reach edge mobile apps or horn triggers in < 1.5 seconds.
- **Audit Logging**: Write operations must generate immutable audit events to the logs database asynchronously within 50ms of transaction commit.

### 2.2 Availability & Resilience
- **Uptime**: Core services must achieve 99.99% annual uptime.
- **Recovery Point Objective (RPO)**: Primary database RPO < 1 minute via streaming replicas.
- **Recovery Time Objective (RTO)**: Disaster Recovery (DR) site promotion RTO < 15 minutes.
- **Offline Edge Capability**: Computer vision inferencing at the edge must operate and buffer event notifications locally during network isolation.

### 2.3 Security & Cryptography
- **Identity & SSO**: Support OAuth 2.0 and OIDC authentication protocols.
- **Data at Rest Encryption**: Store all database storage blocks, S3 objects, and backups using envelope encryption via customer-managed KMS keys (AES-256-GCM).
- **Data in Transit**: Force TLS 1.3 for all external traffic and TLS 1.2 minimum with PFS cipher suites for internal service-to-service communication.
- **Access Control**: Support Attribute-Based Access Control (ABAC) to enforce multi-tenant separation and data access limits (e.g., site-specific access).

### 2.4 Compliance & Certifications
- **OSHA 1910.119 (PSM)**: Strict adherence to Process Safety Information, Process Hazard Analysis, and Management of Change directives.
- **EU REACH & CLP**: Maintain updated chemical registrations mapped to European Chemicals Agency (ECHA) databases.
- **Seveso III Directive (2012/18/EU)**: Enforce chemical quantity thresholds to identify lower-tier and upper-tier hazardous establishments.
- **ISO 45001**: Align workflow structures to support Occupational Health & Safety Management System audits.

### 2.5 Scalability
- **Horizontal Scaling**: All microservices must run as stateless Kubernetes pods supporting Horizontal Pod Autoscaler (HPA) triggers based on CPU/Memory and custom Prometheus metrics.
- **Event Streaming**: Kafka cluster must scale to support up to 50,000 messages/sec.
- **Database Partitioning**: Partition historical logs, inventory transactions, and CV alert tables by time intervals (e.g., monthly).

### 2.6 Accessibility & Localization
- **Accessibility**: User interfaces must conform to WCAG 2.1 Level AA standards.
- **Localization**: Support multi-language interfaces starting with English (en-US), German (de-DE), Spanish (es-ES), and French (fr-FR). Time zones and metric/imperial systems must be fully localized per plant site configuration.

---

## 3. Assumptions & Constraints
- **Infrastructure**: The primary hosting platform is AWS (EKS, RDS, MSK).
- **Camera Standards**: Visual feeds are delivered via standard RTSP over H.264/H.265 compression profiles.
- **Edge Hardware**: Edge CV deployments assume availability of NVIDIA Jetson Orin or compatible GPU-accelerated gateway devices.
- **Connectivity**: Remote sites may experience occasional network degradation; telemetry data must queue locally and sync once connection resumes.

---

## 4. Acceptance Criteria

### 4.1 Incident Logging and Intake
```
Scenario: Incident logging triggers automated notifications and action tracking.
  Given a supervisor is logged in with valid site credentials
  When they submit an incident report with severity level "CRITICAL"
  Then the system must write the event to PostgreSQL
  And publish an "IncidentCreated" event to Kafka topic "prahari.incidents.events"
  And dispatch SMS/Push alerts to all registered site safety officers in < 5 seconds
  And automatically spawn an active Action Item entry with status "Assigned".
```

### 4.2 Chemical Intake and SDS Verification
```
Scenario: PDF Safety Data Sheet upload auto-registers hazard classes.
  Given a warehouse operator uploads a standard SDS PDF
  When the AI agent extracts the CAS registry number "64-17-5" (Ethanol)
  Then the system must cross-reference GHS hazard codes (H225, H319)
  And register the container with status "Awaiting Physical Intake Validation"
  And run a compatibility scan against neighboring containers in Storage Section C.
```
