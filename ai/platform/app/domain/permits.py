from typing import Any, Literal
from pydantic import BaseModel, Field


class PermitCreateRequest(BaseModel):
    permit_title: str
    permit_type: Literal["HOT_WORK", "COLD_WORK", "CONFINED_SPACE", "HEIGHT", "LOTO", "ELECTRICAL"]
    location: str
    applicant_id: str
    details: str
    hazard_controls: list[str] = Field(default_factory=list)
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class PermitRecordDto(BaseModel):
    id: str
    permit_title: str
    permit_type: str
    status: str  # DRAFT, VALIDATED, APPROVED, REJECTED, SUSPENDED, CLOSED
    location: str
    applicant_id: str
    details: str
    risk_score: float
    confidence_score: float
    created_at: str


class PermitValidationResponse(BaseModel):
    is_valid: bool
    missing_controls: list[str] = Field(default_factory=list)
    required_approvers: list[str] = Field(default_factory=list)
    risk_level: str  # LOW, MEDIUM, HIGH, CRITICAL


class RiskAssessmentRequest(BaseModel):
    permit_id: str


class RiskAssessmentResponse(BaseModel):
    likelihood_score: int
    severity_score: int
    residual_risk_score: float
    risk_category: str  # LOW, MEDIUM, HIGH, CRITICAL
    recommended_controls: list[str] = Field(default_factory=list)


class ApprovalRecommendationRequest(BaseModel):
    permit_id: str


class ApprovalRecommendationResponse(BaseModel):
    recommendation: Literal["APPROVE", "REJECT", "SUSPEND"]
    reason: str
    confidence_score: float
