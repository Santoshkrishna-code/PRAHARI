from typing import Any, Literal
from pydantic import BaseModel, Field


class ContractorRegisterRequest(BaseModel):
    company_name: str
    contractor_type: Literal["CIVIL", "MECHANICAL", "ELECTRICAL", "SCAFFOLDING", "WELDING", "SHUTDOWN", "CONSTRUCTION"]
    contact_email: str
    license_number: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class ContractorRecordDto(BaseModel):
    id: str
    company_name: str
    contractor_type: str
    status: str  # REGISTERED, QUALIFIED, REJECTED, SUSPENDED
    safety_rating_score: float
    violations_count: int
    created_at: str


class ContractorQualifyRequest(BaseModel):
    contractor_id: str


class ContractorQualifyResponse(BaseModel):
    is_qualified: bool
    status: str
    safety_score: float
    missing_trainings: list[str] = Field(default_factory=list)
    risk_level: str  # LOW, MEDIUM, HIGH, CRITICAL


class ContractorVerifyRequest(BaseModel):
    contractor_id: str
    document_type: str  # INSURANCE, TRADE_LICENSE, COMPETENCY_CERT, MEDICAL_FITNESS


class ContractorVerifyResponse(BaseModel):
    is_valid: bool
    verification_notes: str


class ContractorScorecardResponse(BaseModel):
    contractor_id: str
    company_name: str
    safety_score: float
    audit_history_count: int
    compliance_status: str
