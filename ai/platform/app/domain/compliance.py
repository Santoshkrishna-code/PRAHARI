from typing import Any, Literal
from pydantic import BaseModel, Field


class ComplianceCheckRequest(BaseModel):
    target_subject: str
    subject_type: Literal["PERMIT", "INSPECTION", "TRAINING", "PPE", "CONTRACTOR"]
    details: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class ComplianceAssessmentDto(BaseModel):
    id: str
    target_subject: str
    subject_type: str
    compliance_score: float
    status: str  # COMPLIANT, NON_COMPLIANT, FLAGGED
    details: str
    created_at: str


class ViolationDto(BaseModel):
    id: str
    assessment_id: str
    regulation_violated: str  # OSHA, ISO_45001, NFPA, etc.
    severity: str  # LOW, MEDIUM, HIGH
    details: str


class CapaPlanRequest(BaseModel):
    assessment_id: str


class CapaPlanResponse(BaseModel):
    corrective_actions: list[str] = Field(default_factory=list)
    preventive_actions: list[str] = Field(default_factory=list)
    compliance_risk: str  # LOW, MEDIUM, HIGH


class ComplianceReportResponse(BaseModel):
    report_markdown: str
    score: float
    confidence: float
