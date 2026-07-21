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


class AnthropicProvider(BaseProvider):
    """Adapter for interacting with Anthropic Claude API endpoints."""

    def get_provider_name(self) -> str:
        return "anthropic"

    def _get_headers(self) -> dict[str, str]:
        return {
            "x-api-key": settings.ANTHROPIC_API_KEY,
            "anthropic-version": "2023-06-01",
            "content-type": "application/json",
        }

    async def generate_chat(self, request: ChatRequest) -> ChatResponse:
        start_time = time.perf_counter()
        model = request.model or "claude-3-5-sonnet"
        url = "https://api.anthropic.com/v1/messages"

        # Separate system messages for Claude's payload structure
        system_content = ""
        claude_messages = []
        for msg in request.messages:
            if msg.role == "system":
                system_content += msg.content + "\n"
            else:
                claude_messages.append({"role": msg.role, "content": msg.content})

        payload = {
            "model": model,
            "messages": claude_messages,
            "max_tokens": request.max_tokens or 4096,
            "temperature": request.temperature,
        }
        if system_content:
            payload["system"] = system_content.strip()

        async with httpx.AsyncClient(timeout=30.0) as client:
            try:
                response = await client.post(url, headers=self._get_headers(), json=payload)
                if response.status_code != 200:
                    raise ValueError(f"Anthropic error {response.status_code}: {response.text}")

                data = response.json()
                content_text = data["content"][0]["text"]
                
                # Retrieve token metrics from responses
                usage_data = data.get("usage", {})
                latency = (time.perf_counter() - start_time) * 1000

                usage = self.estimate_usage(
                    "".join(m.content for m in request.messages),
                    content_text,
                    model,
                )
                if usage_data:
                    usage.prompt_tokens = usage_data.get("input_tokens", usage.prompt_tokens)
                    usage.completion_tokens = usage_data.get("output_tokens", usage.completion_tokens)
                    usage.total_tokens = usage.prompt_tokens + usage.completion_tokens

                return ChatResponse(
                    id=data["id"],
                    model=model,
                    provider="anthropic",
                    choices=[
                        Choice(
                            index=0,
                            message=Message(role="assistant", content=content_text),
                            finish_reason=data.get("stop_reason"),
                        )
                    ],
                    usage=usage,
                    latency_ms=round(latency, 2),
                )
            except Exception as e:
                raise DatabaseException(f"Anthropic connection error: {e}") from e

    async def generate_stream(self, request: ChatRequest) -> AsyncGenerator[ChatResponse, None]:
        # Yield single placeholder to confirm stream initialization (similar to OpenAI)
        model = request.model or "claude-3-5-sonnet"
        yield ChatResponse(
            id="chatcmpl-anthropic-stream",
            model=model,
            provider="anthropic",
            choices=[
                Choice(
                    index=0,
                    message=Message(role="assistant", content="Claude stream started..."),
                    finish_reason=None,
                )
            ],
            usage=self.estimate_usage("", "", model),
            latency_ms=10.0,
        )

    async def generate_embeddings(self, request: EmbeddingRequest) -> EmbeddingResponse:
        # Anthropic does not support native embeddings; we raise exception
        raise NotImplementedError("Anthropic provider does not support embeddings.")
