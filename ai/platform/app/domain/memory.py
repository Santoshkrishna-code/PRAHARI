import datetime
from typing import Any
from pydantic import BaseModel, Field


class MessageDto(BaseModel):
    id: str
    role: str
    content: str
    timestamp: datetime.datetime
    metadata: dict[str, Any] = Field(default_factory=dict)


class SessionDto(BaseModel):
    id: str
    user_id: str
    tenant_id: str
    created_at: datetime.datetime
    status: str


class MemoryEntryDto(BaseModel):
    id: str
    user_id: str
    category: str
    content: str
    importance: int
    relevance: float = 0.0
    recency: float = 0.0
    score: float = 0.0
    timestamp: datetime.datetime


class UserProfileDto(BaseModel):
    id: str
    email: str
    role: str
    department: str
    timezone: str


class PreferenceDto(BaseModel):
    key: str
    value: str


# API Request/Response payload definitions
class SessionCreateRequest(BaseModel):
    user_id: str
    tenant_id: str


class MessageAddRequest(BaseModel):
    session_id: str
    role: str  # "user" | "assistant"
    content: str


class MemorySearchRequest(BaseModel):
    user_id: str
    query: str
    limit: int = Field(5, ge=1, le=50)


class RelevantMemoryResponse(BaseModel):
    memories: list[MemoryEntryDto] = Field(default_factory=list)
    summary: str | None = None
