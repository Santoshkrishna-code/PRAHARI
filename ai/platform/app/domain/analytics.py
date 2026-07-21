from typing import Any, Literal
from pydantic import BaseModel, Field


class ExecutiveDashboardRequest(BaseModel):
    dashboard_type: Literal["CEO", "PLANT", "SAFETY", "MAINTENANCE", "COMPLIANCE"]
    timeframe_days: int = 30
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class ExecutiveDashboardResponse(BaseModel):
    dashboard_type: str
    safety_score: float
    trir: float
    ltifr: float
    permit_compliance_rate: float
    asset_availability_rate: float
    summary_markdown: str


class ReportGenerateRequest(BaseModel):
    report_type: Literal["EXECUTIVE_SUMMARY", "BOARD_REPORT", "SAFETY_PERFORMANCE", "PREDICTIVE_FORECAST"]
    title: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class ReportGenerateResponse(BaseModel):
    report_id: str
    title: str
    report_type: str
    report_markdown: str
    created_at: str


class ForecastRequest(BaseModel):
    metric_name: str
    forecast_months: int = 6
    tenant_id: str = "tenant-ehs-corp"


class ForecastResponse(BaseModel):
    metric_name: str
    current_value: float
    projected_value: float
    trend_direction: str  # IMPROVING, STABLE, DEGRADING
    confidence: float


class KpiMetricsResponse(BaseModel):
    trir: float
    ltifr: float
    near_miss_count: int
    permit_compliance_pct: float
    inspection_completion_pct: float
    asset_availability_pct: float
    overall_compliance_score: float


class UnitBenchmarkRequest(BaseModel):
    target_unit: str
    compare_unit: str
    tenant_id: str = "tenant-ehs-corp"


class UnitBenchmarkResponse(BaseModel):
    target_unit: str
    compare_unit: str
    target_safety_score: float
    compare_safety_score: float
    performance_delta: float
    leader: str
