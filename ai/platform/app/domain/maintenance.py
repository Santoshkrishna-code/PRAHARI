from typing import Any, Literal
from pydantic import BaseModel, Field


class MaintenancePredictRequest(BaseModel):
    asset_id: str
    operating_hours: float = 1200.0
    vibration_trend_mm_s: float = 3.8
    temperature_trend_c: float = 85.0
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class MaintenancePredictionDto(BaseModel):
    id: str
    asset_id: str
    failure_probability: float
    rul_hours: float
    recommended_maintenance_date: str
    maintenance_strategy: str  # PREDICTIVE, CONDITION_BASED, PREVENTIVE, RELIABILITY_CENTERED
    status: str  # RECOMMENDED, SCHEDULED, COMPLETED
    created_at: str


class RulEstimateRequest(BaseModel):
    asset_id: str


class RulEstimateResponse(BaseModel):
    asset_id: str
    remaining_useful_life_hours: float
    remaining_useful_life_days: float
    confidence_score: float
    degradation_trend: str  # STABLE, ACCELERATING, SEVERE


class MaintenancePlanRequest(BaseModel):
    asset_id: str
    work_order_title: str


class MaintenancePlanResponse(BaseModel):
    work_order_id: str
    asset_id: str
    priority: str  # LOW, MEDIUM, HIGH, CRITICAL
    downtime_hours_est: float
    required_spare_parts: list[str] = Field(default_factory=list)


class ReliabilityDashboardResponse(BaseModel):
    total_assets_monitored: int
    high_risk_assets_count: int
    average_rul_hours: float
    pending_work_orders_count: int
