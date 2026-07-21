from typing import Any, Literal
from pydantic import BaseModel, Field


class EmergencyAssessRequest(BaseModel):
    title: str
    emergency_type: Literal["FIRE", "EXPLOSION", "CHEMICAL_SPILL", "GAS_LEAK", "MEDICAL", "ACTIVE_THREAT"]
    location: str
    details: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class EmergencyRecordDto(BaseModel):
    id: str
    title: str
    emergency_type: str
    severity_level: str  # LEVEL_1, LEVEL_2, LEVEL_3_CRITICAL
    status: str  # ACTIVE, STABILIZED, RECOVERED, CLOSED
    location: str
    affected_zone: str
    evacuation_route: str
    created_at: str


class EvacuationPlanRequest(BaseModel):
    emergency_id: str


class EvacuationPlanResponse(BaseModel):
    emergency_id: str
    primary_route: str
    alternative_route: str
    assembly_point: str
    restricted_zone: str


class EmergencyNotifyRequest(BaseModel):
    emergency_id: str
    message: str
    channels: list[str] = Field(default_factory=lambda: ["SMS", "PA_SYSTEM", "MOBILE_APP"])


class EmergencyNotifyResponse(BaseModel):
    emergency_id: str
    notifications_sent: int
    status: str


class SitrepResponse(BaseModel):
    emergency_id: str
    situation_summary: str
    severity_level: str
    active_responders_count: int
    sitrep_markdown: str


class EmergencyResourceDto(BaseModel):
    id: str
    resource_name: str
    resource_type: str
    quantity: int
    location: str
    status: str  # AVAILABLE, DISPATCHED
