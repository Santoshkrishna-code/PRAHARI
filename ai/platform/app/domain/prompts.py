from typing import Any, Literal
from pydantic import BaseModel, Field


class PromptCategoryDto(BaseModel):
    id: str
    name: str
    code: str
    description: str | None = None


class PromptTagDto(BaseModel):
    id: str
    name: str
    code: str


class PromptVersionDto(BaseModel):
    id: str
    prompt_id: str
    version_string: str = Field(..., description="Semantic version string e.g. 1.0.0")
    system_template: str
    user_template: str
    few_shots: list[dict[str, Any]] | None = Field(default_factory=list)
    response_format: dict[str, Any] | None = None
    status: Literal["DRAFT", "REVIEW", "APPROVED", "DEPRECATED"] = "DRAFT"
    created_by: str
    approved_by: str | None = None
    metadata: dict[str, Any] | None = Field(default_factory=dict)


class PromptDto(BaseModel):
    id: str
    name: str = Field(..., description="Unique slug name e.g. safety-audit-incident")
    description: str | None = None
    category_id: str | None = None
    tags: list[str] = Field(default_factory=list)
    active_version_string: str | None = None
    versions: list[PromptVersionDto] = Field(default_factory=list)


# Requests and Response Schemas
class RenderRequest(BaseModel):
    prompt_name: str
    version: str | None = None  # If None, resolve active approved version
    variables: dict[str, Any] = Field(default_factory=dict)


class RenderResponse(BaseModel):
    system_prompt: str
    user_prompt: str
    variables_resolved: list[str]
    variables_missing: list[str]
    estimated_tokens: int
    estimated_cost: float


class ValidateRequest(BaseModel):
    system_template: str
    user_template: str
    variables: dict[str, Any] = Field(default_factory=dict)


class ValidateResponse(BaseModel):
    is_valid: bool
    errors: list[str] = Field(default_factory=list)
    warnings: list[str] = Field(default_factory=list)
    prompt_injection_detected: bool = False
    unsafe_instructions_detected: bool = False


class CompareRequest(BaseModel):
    prompt_name: str
    version_a: str
    version_b: str


class CompareResponse(BaseModel):
    prompt_name: str
    version_a: str
    version_b: str
    diff_system: str
    diff_user: str


class PromptTestRequest(BaseModel):
    prompt_name: str
    version: str
    variables: dict[str, Any] = Field(default_factory=dict)
    expected_substrings: list[str] = Field(default_factory=list)


class PromptTestResponse(BaseModel):
    passed: bool
    rendered_system: str
    rendered_user: str
    matched_substrings: list[str] = Field(default_factory=list)
    failed_substrings: list[str] = Field(default_factory=list)
