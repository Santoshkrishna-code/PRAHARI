from typing import Any, Literal
from pydantic import BaseModel, Field


class InspectionPlanRequest(BaseModel):
    title: str
    inspection_type: Literal["DAILY_SAFETY", "EQUIPMENT", "FIRE_SAFETY", "CHEMICAL_STORAGE", "CONSTRUCTION"]
    location: str
    inspector_id: str
    scope_notes: str | None = None
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class InspectionRecordDto(BaseModel):
    id: str
    title: str
    inspection_type: str
    status: str  # PLANNED, IN_PROGRESS, COMPLETED
    location: str
    inspector_id: str
    score: float
    findings_count: int
    summary_notes: str | None = None
    created_at: str


class ChecklistGenerationRequest(BaseModel):
    inspection_type: str
    location: str


class ChecklistGenerationResponse(BaseModel):
    inspection_type: str
    items: list[str] = Field(default_factory=list)


class FindingAnalysisRequest(BaseModel):
    inspection_id: str
    observation_text: str


class FindingAnalysisResponse(BaseModel):
    finding_type: str  # UNSAFE_ACT, UNSAFE_CONDITION, MINOR_NON_CONFORMANCE, MAJOR_NON_CONFORMANCE, CRITICAL_FINDING
    severity: str  # LOW, MEDIUM, HIGH, CRITICAL
    details: str
    recommended_capa: str
    risk_priority_score: float


class InspectionReportResponse(BaseModel):
    inspection_id: str
    score: float
    report_markdown: str
