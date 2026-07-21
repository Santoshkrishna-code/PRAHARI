from typing import Any, Literal
from pydantic import BaseModel, Field


class ToolDto(BaseModel):
    id: str
    name: str = Field(..., description="Unique slug name of tool e.g. create_incident")
    description: str
    category: str
    type: str  # REST, SQL, INTERNAL
    status: str
    timeout_seconds: int = 30
    retry_count: int = 3
    active_version_string: str | None = None


class ToolVersionDto(BaseModel):
    id: str
    tool_id: str
    version_string: str
    input_schema: dict[str, Any] = Field(default_factory=dict)
    execution_target: str
    status: Literal["DRAFT", "ACTIVE", "DEPRECATED"] = "ACTIVE"


# Request/Response definitions
class ToolRegisterRequest(BaseModel):
    name: str
    description: str
    category: str
    type: Literal["REST", "SQL", "INTERNAL"]
    version_string: str = "1.0.0"
    input_schema: dict[str, Any] = Field(default_factory=dict)
    execution_target: str = Field(..., description="API URL, SQL template, or internal service trigger name")
    timeout_seconds: int = 30
    retry_count: int = 3


class ToolExecuteRequest(BaseModel):
    tool_name: str
    version: str | None = None
    payload: dict[str, Any] = Field(default_factory=dict)
    user_id: str = "user-id"
    tenant_id: str = "tenant-id"


class ToolExecuteResponse(BaseModel):
    execution_id: str
    success: bool
    output: Any = None
    duration_ms: int
    error: str | None = None


# MCP Protocol Interfaces
class McpInitializeRequest(BaseModel):
    client_name: str
    client_version: str
    protocol_version: str = "1.0"


class McpInitializeResponse(BaseModel):
    protocol_version: str = "1.0"
    server_name: str = "PRAHARI-MCP-Server"
    server_version: str = "1.0.0"
    capabilities: dict[str, Any] = Field(default_factory=dict)


class McpToolDto(BaseModel):
    name: str
    description: str
    inputSchema: dict[str, Any] = Field(default_factory=dict)


class McpToolsListResponse(BaseModel):
    tools: list[McpToolDto] = Field(default_factory=list)
