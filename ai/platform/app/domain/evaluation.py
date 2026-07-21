from typing import Any
from pydantic import BaseModel, Field


class EvaluationRequest(BaseModel):
    run_id: str | None = None
    text_input: str
    text_output: str
    retrieved_context: list[str] = Field(default_factory=list)
    cost: float = 0.0
    latency_ms: int = 0
    model: str = "gpt-4o"
    user_id: str = "sys-admin"
    tenant_id: str = "tenant-ehs-corp"


class EvaluationResponse(BaseModel):
    evaluation_id: str
    faithfulness_score: float
    correctness_score: float
    relevance_score: float
    groundedness_score: float
    overall_score: float
    latency_ms: int
    cost: float
    model: str


class BenchmarkRequest(BaseModel):
    name: str
    model_name: str
    test_prompts: list[str] = Field(default_factory=list)


class BenchmarkResponse(BaseModel):
    benchmark_id: str
    model_name: str
    average_latency: float
    average_cost: float
    average_accuracy: float
    reliability_rate: float


class ExperimentRequest(BaseModel):
    name: str
    description: str | None = None
    variant_a_config: dict[str, Any] = Field(default_factory=dict)
    variant_b_config: dict[str, Any] = Field(default_factory=dict)


class ExperimentResponse(BaseModel):
    experiment_id: str
    name: str
    status: str


class LeaderboardEntryDto(BaseModel):
    model_name: str
    average_latency: float
    average_cost: float
    average_accuracy: float
    reliability_rate: float
    score: float


class HumanFeedbackRequest(BaseModel):
    evaluation_id: str
    rating: int = Field(..., ge=1, le=5)
    comments: str | None = None
