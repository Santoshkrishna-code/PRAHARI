from typing import Any, Literal
from pydantic import BaseModel, Field


class AgentNodeDto(BaseModel):
    id: str
    name: str = Field(..., description="Unique node name e.g. safety_reviewer")
    prompt_template_name: str | None = None
    tools_allowed: list[str] = Field(default_factory=list)
    model: str = "gpt-4o"


class AgentEdgeDto(BaseModel):
    id: str
    source_node: str
    target_node: str
    conditional_expr: str | None = Field(None, description="Heuristic condition or LLM router target")


class AgentGraphDto(BaseModel):
    id: str
    name: str
    description: str
    start_node: str
    nodes: list[AgentNodeDto] = Field(default_factory=list)
    edges: list[AgentEdgeDto] = Field(default_factory=list)


# Requests/Responses
class AgentRunRequest(BaseModel):
    graph_id: str
    input_state: dict[str, Any] = Field(default_factory=dict)
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class AgentRunResponse(BaseModel):
    execution_id: str
    status: Literal["COMPLETED", "INTERRUPTED", "FAILED"]
    current_node: str
    state: dict[str, Any] = Field(default_factory=dict)
    message: str | None = None


class AgentApprovalRequest(BaseModel):
    execution_id: str
    approved: bool
    comments: str | None = None
