from abc import ABC, abstractmethod
from collections.abc import AsyncGenerator
from app.core.models import (
    ChatRequest,
    ChatResponse,
    CompletionRequest,
    CompletionResponse,
    EmbeddingRequest,
    EmbeddingResponse,
    Usage,
)
from app.core.model_registry import MODEL_REGISTRY


class BaseProvider(ABC):
    """Abstract interface defining the requirements for all EHS AI LLM providers."""

    @abstractmethod
    def get_provider_name(self) -> str:
        """Returns the registered code string representing the vendor name."""
        pass

    @abstractmethod
    async def generate_chat(self, request: ChatRequest) -> ChatResponse:
        """Sends chat query and returns complete choice block."""
        pass

    @abstractmethod
    async def generate_stream(self, request: ChatRequest) -> AsyncGenerator[ChatResponse, None]:
        """Sends chat query and yields streaming chunks."""
        pass

    @abstractmethod
    async def generate_embeddings(self, request: EmbeddingRequest) -> EmbeddingResponse:
        """Requests vector arrays for inputs."""
        pass

    async def generate_completion(self, request: CompletionRequest) -> CompletionResponse:
        """Wrapper compiling a legacy prompt query into a single message chat generation."""
        chat_req = ChatRequest(
            messages=[{"role": "user", "content": request.prompt}],
            model=request.model,
            temperature=request.temperature,
            max_tokens=request.max_tokens,
            stream=request.stream,
            use_cache=request.use_cache,
        )
        chat_res = await self.generate_chat(chat_req)
        return CompletionResponse(
            text=chat_res.choices[0].message.content,
            model=chat_res.model,
            provider=chat_res.provider,
            usage=chat_res.usage,
            latency_ms=chat_res.latency_ms,
            cached=chat_res.cached,
        )

    def estimate_usage(self, prompt: str, completion: str, model_name: str) -> Usage:
        """Heuristic token usage calculator based on whitespace divisions."""
        metadata = MODEL_REGISTRY.get(model_name)
        
        # Approximate 4 characters per token as standard LLM rule of thumb
        prompt_tokens = max(1, len(prompt) // 4)
        completion_tokens = max(1, len(completion) // 4)
        total_tokens = prompt_tokens + completion_tokens

        cost = 0.0
        if metadata:
            input_cost = (prompt_tokens / 1_000_000) * metadata.input_cost_per_million
            output_cost = (completion_tokens / 1_000_000) * metadata.output_cost_per_million
            cost = input_cost + output_cost

        return Usage(
            prompt_tokens=prompt_tokens,
            completion_tokens=completion_tokens,
            total_tokens=total_tokens,
            cost=round(cost, 6),
        )
