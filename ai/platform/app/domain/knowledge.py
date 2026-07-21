import datetime
from typing import Any
from pydantic import BaseModel, Field


class DocumentDto(BaseModel):
    id: str
    filename: str
    content_type: str
    file_size: int
    department: str | None = None
    project: str | None = None
    status: str
    tenant_id: str
    created_at: datetime.datetime


class ChunkDto(BaseModel):
    id: str
    document_id: str
    chunk_index: int
    content: str
    page_number: int | None = None
    section_title: str | None = None
    score: float = 0.0


class CitationDto(BaseModel):
    document_name: str
    page_number: int | None = None
    section_title: str | None = None
    chunk_id: str
    score: float


# Request/Response payloads
class SearchRequest(BaseModel):
    query: str
    limit: int = Field(5, ge=1, le=50)
    department: str | None = None
    project: str | None = None


class RetrieveResponse(BaseModel):
    chunks: list[ChunkDto] = Field(default_factory=list)


class QueryRequest(BaseModel):
    query: str
    session_id: str | None = None
    limit: int = Field(5, ge=1, le=50)
    department: str | None = None
    project: str | None = None


class QueryResponse(BaseModel):
    answer: str
    citations: list[CitationDto] = Field(default_factory=list)


class IndexRequest(BaseModel):
    document_id: str
    text_content: str

