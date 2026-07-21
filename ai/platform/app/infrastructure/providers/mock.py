import uuid
import time
import asyncio
from collections.abc import AsyncGenerator
from app.infrastructure.providers.base import BaseProvider
from app.core.models import (
    ChatRequest,
    ChatResponse,
    Choice,
    Message,
    EmbeddingRequest,
    EmbeddingResponse,
    EmbeddingData,
)


class MockProvider(BaseProvider):
    """Mock implementation of the provider interface for testing and offline modes."""

    def get_provider_name(self) -> str:
        return "mock"

    async def generate_chat(self, request: ChatRequest) -> ChatResponse:
        start_time = time.perf_counter()
        
        # Simulate short execution network latency
        await asyncio.sleep(0.01)

        model = request.model or "mock-model"
        prompt_content = "".join(msg.content for msg in request.messages)
        completion_content = f"Mock EHS response matching query of length {len(prompt_content)}."

        usage = self.estimate_usage(prompt_content, completion_content, model)
        latency = (time.perf_counter() - start_time) * 1000

        return ChatResponse(
            id=f"chatcmpl-{uuid.uuid4()}",
            model=model,
            provider="mock",
            choices=[
                Choice(
                    index=0,
                    message=Message(role="assistant", content=completion_content),
                    finish_reason="stop",
                )
            ],
            usage=usage,
            latency_ms=round(latency, 2),
        )

    async def generate_stream(self, request: ChatRequest) -> AsyncGenerator[ChatResponse, None]:
        model = request.model or "mock-model"
        prompt_content = "".join(msg.content for msg in request.messages)
        words = ["This", " is", " a", " mock", " stream", " response."]

        accumulated = ""
        for idx, word in enumerate(words):
            accumulated += word
            usage = self.estimate_usage(prompt_content, accumulated, model)
            
            yield ChatResponse(
                id="chatcmpl-mockstream",
                model=model,
                provider="mock",
                choices=[
                    Choice(
                        index=0,
                        message=Message(role="assistant", content=word),
                        finish_reason="stop" if idx == len(words) - 1 else None,
                    )
                ],
                usage=usage,
                latency_ms=5.0 * (idx + 1),
            )
            await asyncio.sleep(0.01)

    async def generate_embeddings(self, request: EmbeddingRequest) -> EmbeddingResponse:
        model = request.model or "mock-model"
        inputs = [request.input] if isinstance(request.input, str) else request.input

        embeddings_data = []
        for idx, text in enumerate(inputs):
            # Return static unit vector representing embedding
            embeddings_data.append(
                EmbeddingData(
                    index=idx,
                    embedding=[0.1] * 1536,
                )
            )

        usage = self.estimate_usage("".join(inputs), "", model)

        return EmbeddingResponse(
            model=model,
            provider="mock",
            data=embeddings_data,
            usage=usage,
        )
