from typing import Any, Literal
from pydantic import BaseModel, Field
from datetime import datetime


class PolicyDto(BaseModel):
    id: str
    name: str
    code: str
    description: str | None = None
    action: str  # BLOCK, WARN, LOG
    status: str  # ACTIVE, DISABLED


class PolicyCreateRequest(BaseModel):
    name: str
    code: str
    description: str | None = None
    action: Literal["BLOCK", "WARN", "LOG"] = "BLOCK"


class PolicyUpdateRequest(BaseModel):
    action: Literal["BLOCK", "WARN", "LOG"]
    status: Literal["ACTIVE", "DISABLED"]


class GuardrailCheckRequest(BaseModel):
    text: str
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class GuardrailCheckResponse(BaseModel):
    is_safe: bool
    flagged_rules: list[str] = Field(default_factory=list)
    sanitised_text: str
    pii_detected: list[str] = Field(default_factory=list)
    secrets_detected: list[str] = Field(default_factory=list)
    toxicity_score: float = 0.0
    hallucination_detected: bool = False


class HumanReviewRequest(BaseModel):
    event_id: str
    approved: bool
    comments: str | None = None


class SafetyEventDto(BaseModel):
    id: str
    user_id: str
    tenant_id: str
    input_text: str
    output_text: str | None = None
    action_taken: str
    timestamp: str
