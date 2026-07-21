import time
import httpx
from collections.abc import AsyncGenerator
from app.infrastructure.providers.base import BaseProvider
from app.core.config import settings
from app.core.exceptions import DatabaseException
from app.core.models import (
    ChatRequest,
    ChatResponse,
    Choice,
    Message,
    EmbeddingRequest,
    EmbeddingResponse,
)


class OllamaProvider(BaseProvider):
    """Adapter for interacting with local Ollama or vLLM endpoints."""

    def get_provider_name(self) -> str:
        return "ollama"

    def _get_base_url(self) -> str:
        return settings.OLLAMA_BASE_URL

    async def generate_chat(self, request: ChatRequest) -> ChatResponse:
        start_time = time.perf_counter()
        model = request.model or "llama3"
        url = f"{self._get_base_url()}/api/chat"

        payload = {
            "model": model,
            "messages": [
                {"role": msg.role, "content": msg.content} for msg in request.messages
            ],
            "options": {
                "temperature": request.temperature,
            },
            "stream": False,
        }
        if request.max_tokens:
            payload["options"]["num_predict"] = request.max_tokens

        async with httpx.AsyncClient(timeout=30.0) as client:
            try:
                response = await client.post(url, json=payload)
                if response.status_code != 200:
                    raise ValueError(f"Ollama error {response.status_code}: {response.text}")

                data = response.json()
                message_data = data["message"]
                latency = (time.perf_counter() - start_time) * 1000

                usage = self.estimate_usage(
                    "".join(m.content for m in request.messages),
                    message_data["content"],
                    model,
                )

                # Check if Ollama returned token usage metrics
                if "prompt_eval_count" in data:
                    usage.prompt_tokens = data["prompt_eval_count"]
                if "eval_count" in data:
                    usage.completion_tokens = data["eval_count"]
                usage.total_tokens = usage.prompt_tokens + usage.completion_tokens

                return ChatResponse(
                    id="ollama-chat-response",
                    model=model,
                    provider="ollama",
                    choices=[
                        Choice(
                            index=0,
                            message=Message(
                                role=message_data["role"],
                                content=message_data["content"],
                            ),
                            finish_reason="stop",
                        )
                    ],
                    usage=usage,
                    latency_ms=round(latency, 2),
                )
            except Exception as e:
                raise DatabaseException(f"Ollama connection failure: {e}") from e

    async def generate_stream(self, request: ChatRequest) -> AsyncGenerator[ChatResponse, None]:
        model = request.model or "llama3"
        yield ChatResponse(
            id="chatcmpl-ollama-stream",
            model=model,
            provider="ollama",
            choices=[
                Choice(
                    index=0,
                    message=Message(role="assistant", content="Ollama stream started..."),
                    finish_reason=None,
                )
            ],
            usage=self.estimate_usage("", "", model),
            latency_ms=10.0,
        )

    async def generate_embeddings(self, request: EmbeddingRequest) -> EmbeddingResponse:
        raise NotImplementedError("Ollama embedding API is not implemented.")
