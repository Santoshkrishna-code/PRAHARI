from typing import Any, Literal
from pydantic import BaseModel, Field


class AssetHealthCheckRequest(BaseModel):
    asset_name: str
    asset_category: Literal["PUMP", "COMPRESSOR", "BOILER", "PRESSURE_VESSEL", "PIPELINE", "CRANE", "ROBOT"]
    location: str
    operating_hours: float = 0.0
    vibration_mm_s: float = 1.2
    temperature_c: float = 65.0
    pressure_bar: float = 5.0
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class AssetRecordDto(BaseModel):
    id: str
    asset_name: str
    asset_category: str
    status: str  # OPERATIONAL, MAINTENANCE_REQUIRED, SHUTDOWN, DECOMMISSIONED
    health_score: float
    operating_hours: float
    created_at: str


class AssetShutdownRequest(BaseModel):
    asset_id: str
    reason: str


class AssetShutdownResponse(BaseModel):
    asset_id: str
    status: str
    action_taken: str
    safety_lockout_issued: bool


class AssetScorecardResponse(BaseModel):
    asset_id: str
    asset_name: str
    health_score: float
    operational_status: str
    calibration_valid: bool
