from typing import Any, Literal
from pydantic import BaseModel, Field


class SupervisorChatRequest(BaseModel):
    user_query: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class SupervisorChatResponse(BaseModel):
    session_id: str
    plan_status: str  # PLANNED, EXECUTING, COMPLETED, FAILED
    execution_dag: list[dict[str, Any]] = Field(default_factory=list)
    consensus_summary: str
    details_by_agent: dict[str, Any] = Field(default_factory=dict)


class SupervisorPlanRequest(BaseModel):
    user_query: str


class SupervisorPlanResponse(BaseModel):
    planned_steps: list[dict[str, str]] = Field(default_factory=list)
    estimated_latency_sec: float


class SupervisorExecuteRequest(BaseModel):
    session_id: str


class SupervisorStatusResponse(BaseModel):
    session_id: str
    status: str
    completed_subtasks_count: int
    total_subtasks_count: int


class SupervisorDagResponse(BaseModel):
    session_id: str
    nodes: list[dict[str, str]] = Field(default_factory=list)
    edges: list[dict[str, str]] = Field(default_factory=list)
