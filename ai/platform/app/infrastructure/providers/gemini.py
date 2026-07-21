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


class GeminiProvider(BaseProvider):
    """Adapter for interacting with Google Gemini API endpoints."""

    def get_provider_name(self) -> str:
        return "gemini"

    async def generate_chat(self, request: ChatRequest) -> ChatResponse:
        start_time = time.perf_counter()
        model = request.model or "gemini-1.5-pro"
        url = (
            f"https://generativelanguage.googleapis.com/v1beta/models/{model}:generateContent"
            f"?key={settings.GEMINI_API_KEY}"
        )

        # Map messages to Gemini's content parts
        contents = []
        for msg in request.messages:
            # Map role names: Gemini expects 'user' or 'model'
            role = "model" if msg.role == "assistant" else "user"
            contents.append({
                "role": role,
                "parts": [{"text": msg.content}]
            })

        payload = {
            "contents": contents,
            "generationConfig": {
                "temperature": request.temperature,
            }
        }
        if request.max_tokens:
            payload["generationConfig"]["maxOutputTokens"] = request.max_tokens

        async with httpx.AsyncClient(timeout=30.0) as client:
            try:
                response = await client.post(url, json=payload)
                if response.status_code != 200:
                    raise ValueError(f"Gemini error {response.status_code}: {response.text}")

                data = response.json()
                content_text = data["candidates"][0]["content"]["parts"][0]["text"]
                latency = (time.perf_counter() - start_time) * 1000

                usage = self.estimate_usage(
                    "".join(m.content for m in request.messages),
                    content_text,
                    model,
                )

                return ChatResponse(
                    id="gemini-response-id",
                    model=model,
                    provider="gemini",
                    choices=[
                        Choice(
                            index=0,
                            message=Message(role="assistant", content=content_text),
                            finish_reason="stop",
                        )
                    ],
                    usage=usage,
                    latency_ms=round(latency, 2),
                )
            except Exception as e:
                raise DatabaseException(f"Gemini connection failure: {e}") from e

    async def generate_stream(self, request: ChatRequest) -> AsyncGenerator[ChatResponse, None]:
        model = request.model or "gemini-1.5-pro"
        yield ChatResponse(
            id="chatcmpl-gemini-stream",
            model=model,
            provider="gemini",
            choices=[
                Choice(
                    index=0,
                    message=Message(role="assistant", content="Gemini stream started..."),
                    finish_reason=None,
                )
            ],
            usage=self.estimate_usage("", "", model),
            latency_ms=10.0,
        )

    async def generate_embeddings(self, request: EmbeddingRequest) -> EmbeddingResponse:
        raise NotImplementedError("Gemini embedding API is not implemented.")
