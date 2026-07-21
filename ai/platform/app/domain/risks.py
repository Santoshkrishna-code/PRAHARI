from typing import Any, Literal
from pydantic import BaseModel, Field


class RiskAssessmentRequest(BaseModel):
    title: str
    risk_type: Literal["OCCUPATIONAL", "FIRE", "CHEMICAL", "ELECTRICAL", "ENVIRONMENTAL", "MECHANICAL"]
    methodology: Literal["HIRA", "JSA", "BOWTIE", "FMEA", "HAZOP"] = "HIRA"
    details: str
    location: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class RiskAssessmentDto(BaseModel):
    id: str
    title: str
    risk_type: str
    methodology: str
    initial_risk_score: float
    residual_risk_score: float
    risk_category: str  # LOW, MEDIUM, HIGH, CRITICAL
    created_at: str


class RiskPredictRequest(BaseModel):
    subject: str
    historical_incidents_count: int = 0
    weather_condition: str | None = None


class RiskPredictResponse(BaseModel):
    predicted_risk_score: float
    trend_status: str  # STABLE, INCREASING, DECREASING
    forecast_summary: str
    risk_confidence: float


class MitigationPlanRequest(BaseModel):
    assessment_id: str


class MitigationPlanResponse(BaseModel):
    engineering_controls: list[str] = Field(default_factory=list)
    administrative_controls: list[str] = Field(default_factory=list)
    ppe_recommendations: list[str] = Field(default_factory=list)
    residual_risk: float


class RiskRegisterEntryDto(BaseModel):
    id: str
    assessment_id: str
    hazard_name: str
    severity_score: int
    likelihood_score: int
    trend_status: str
