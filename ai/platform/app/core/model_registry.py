from pydantic import BaseModel


class ModelMetadata(BaseModel):
    provider: str
    name: str
    version: str
    context_window: int
    supports_streaming: bool
    supports_vision: bool
    supports_function_calling: bool
    supports_embeddings: bool
    supports_json_mode: bool
    supports_audio: bool
    supports_images: bool
    max_tokens: int
    input_cost_per_million: float
    output_cost_per_million: float
    latency_tier: str  # "high", "medium", "low"


MODEL_REGISTRY: dict[str, ModelMetadata] = {
    "gpt-4o": ModelMetadata(
        provider="openai",
        name="gpt-4o",
        version="2024-05-13",
        context_window=128000,
        supports_streaming=True,
        supports_vision=True,
        supports_function_calling=True,
        supports_embeddings=False,
        supports_json_mode=True,
        supports_audio=False,
        supports_images=True,
        max_tokens=4096,
        input_cost_per_million=5.0,
        output_cost_per_million=15.0,
        latency_tier="medium",
    ),
    "gpt-3.5-turbo": ModelMetadata(
        provider="openai",
        name="gpt-3.5-turbo",
        version="0125",
        context_window=16385,
        supports_streaming=True,
        supports_vision=False,
        supports_function_calling=True,
        supports_embeddings=False,
        supports_json_mode=True,
        supports_audio=False,
        supports_images=False,
        max_tokens=4096,
        input_cost_per_million=0.5,
        output_cost_per_million=1.5,
        latency_tier="low",
    ),
    "claude-3-5-sonnet": ModelMetadata(
        provider="anthropic",
        name="claude-3-5-sonnet",
        version="20240620",
        context_window=200000,
        supports_streaming=True,
        supports_vision=True,
        supports_function_calling=True,
        supports_embeddings=False,
        supports_json_mode=True,
        supports_audio=False,
        supports_images=True,
        max_tokens=8192,
        input_cost_per_million=3.0,
        output_cost_per_million=15.0,
        latency_tier="medium",
    ),
    "gemini-1.5-pro": ModelMetadata(
        provider="gemini",
        name="gemini-1.5-pro",
        version="001",
        context_window=1000000,
        supports_streaming=True,
        supports_vision=True,
        supports_function_calling=True,
        supports_embeddings=False,
        supports_json_mode=True,
        supports_audio=True,
        supports_images=True,
        max_tokens=8192,
        input_cost_per_million=1.25,
        output_cost_per_million=5.0,
        latency_tier="high",
    ),
    "llama3": ModelMetadata(
        provider="ollama",
        name="llama3",
        version="latest",
        context_window=8192,
        supports_streaming=True,
        supports_vision=False,
        supports_function_calling=False,
        supports_embeddings=False,
        supports_json_mode=False,
        supports_audio=False,
        supports_images=False,
        max_tokens=2048,
        input_cost_per_million=0.0,
        output_cost_per_million=0.0,
        latency_tier="low",
    ),
    "text-embedding-3-small": ModelMetadata(
        provider="openai",
        name="text-embedding-3-small",
        version="1",
        context_window=8191,
        supports_streaming=False,
        supports_vision=False,
        supports_function_calling=False,
        supports_embeddings=True,
        supports_json_mode=False,
        supports_audio=False,
        supports_images=False,
        max_tokens=0,
        input_cost_per_million=0.02,
        output_cost_per_million=0.0,
        latency_tier="low",
    ),
    "mock-model": ModelMetadata(
        provider="mock",
        name="mock-model",
        version="1.0",
        context_window=8192,
        supports_streaming=True,
        supports_vision=True,
        supports_function_calling=True,
        supports_embeddings=True,
        supports_json_mode=True,
        supports_audio=True,
        supports_images=True,
        max_tokens=4096,
        input_cost_per_million=0.0,
        output_cost_per_million=0.0,
        latency_tier="low",
    ),
}
