from typing import Any, Literal
from pydantic import BaseModel, Field


class IncidentIntakeRequest(BaseModel):
    title: str
    description: str
    classification: str  # Lost Time Injury, Fatality, Fire, Chemical Spill, Gas Leak, etc.
    location: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class IncidentInvestigationDto(BaseModel):
    id: str
    title: str
    severity: str
    status: str
    details: str
    confidence_score: float
    reported_at: str


class RootCauseRequest(BaseModel):
    investigation_id: str
    methodology: Literal["5_WHYS", "BOWTIE", "FISHBONE"] = "5_WHYS"


class RootCauseResponse(BaseModel):
    methodology: str
    root_causes: list[str] = Field(default_factory=list)
    causative_chain: list[str] = Field(default_factory=list)


class ReportResponse(BaseModel):
    report_markdown: str
    overall_confidence: float
    grounding_score: float
    cost: float


class RecommendationRequest(BaseModel):
    investigation_id: str


class RecommendationResponse(BaseModel):
    corrective_actions: list[str] = Field(default_factory=list)
    preventive_actions: list[str] = Field(default_factory=list)
