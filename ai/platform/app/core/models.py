from typing import Any, Generic, TypeVar, Literal
from pydantic import BaseModel, Field

T = TypeVar("T")


class StandardResponse(BaseModel, Generic[T]):
    success: bool = True
    message: str = "Operation completed successfully"
    data: T | None = None


class PaginationMetadata(BaseModel):
    page: int = Field(..., description="Current page index")
    limit: int = Field(..., description="Items limit per page")
    total: int = Field(..., description="Total items count")
    total_pages: int = Field(..., description="Total pages count")


class PaginatedResponse(BaseModel, Generic[T]):
    success: bool = True
    message: str = "Records retrieved successfully"
    data: list[T] = Field(default_factory=list)
    pagination: PaginationMetadata


class ErrorDetails(BaseModel):
    type: str = Field(..., description="Link to the error type specification")
    title: str = Field(..., description="Brief error category description")
    status: int = Field(..., description="HTTP status code")
    detail: str = Field(..., description="Comprehensive explanation of the issue")
    instance: str = Field(..., description="Request path instance of the error")
    invalid_params: list[dict[str, Any]] | None = Field(
        None, description="Invalid parameters listing for validation failures"
    )

class HealthStatus(BaseModel):
    status: str
    database: str
    cache: str

# Message Definition
class Message(BaseModel):
    role: Literal["system", "user", "assistant"]
    content: str

# Chat & Completion Requests
class ChatRequest(BaseModel):
    messages: list[Message]
    model: str | None = None
    temperature: float = Field(0.7, ge=0.0, le=2.0)
    max_tokens: int | None = Field(None, gt=0)
    stream: bool = False
    use_cache: bool = True
    fallback_chain: list[str] | None = None

class CompletionRequest(BaseModel):
    prompt: str
    model: str | None = None
    temperature: float = Field(0.7, ge=0.0, le=2.0)
    max_tokens: int | None = Field(None, gt=0)
    stream: bool = False
    use_cache: bool = True

# Response Support Structures
class Choice(BaseModel):
    message: Message
    index: int
    finish_reason: str | None = None

class Usage(BaseModel):
    prompt_tokens: int = 0
    completion_tokens: int = 0
    total_tokens: int = 0
    cost: float = 0.0

class ChatResponse(BaseModel):
    id: str
    model: str
    provider: str
    choices: list[Choice]
    usage: Usage
    latency_ms: float
    cached: bool = False

class CompletionResponse(BaseModel):
    text: str
    model: str
    provider: str
    usage: Usage
    latency_ms: float
    cached: bool = False

# Embedding Schemas
class EmbeddingRequest(BaseModel):
    input: str | list[str]
    model: str | None = None

class EmbeddingData(BaseModel):
    embedding: list[float]
    index: int

class EmbeddingResponse(BaseModel):
    model: str
    provider: str
    data: list[EmbeddingData]
    usage: Usage

