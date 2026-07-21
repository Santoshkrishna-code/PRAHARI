import time
import json
import asyncio
import uuid
from typing import Any
from collections.abc import AsyncGenerator
import boto3
from botocore.config import Config
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


class BedrockProvider(BaseProvider):
    """Adapter for AWS Bedrock converse API models integration."""

    def get_provider_name(self) -> str:
        return "bedrock"

    def _get_client(self) -> Any:
        """Retrieves Amazon Bedrock runtime client."""
        config = Config(read_timeout=30, connect_timeout=10, retries={"max_attempts": 2})
        return boto3.client(
            service_name="bedrock-runtime",
            region_name=settings.AWS_REGION,
            aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
            aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
            config=config,
        )

    async def generate_chat(self, request: ChatRequest) -> ChatResponse:
        start_time = time.perf_counter()
        model = request.model or "meta.llama3-8b-instruct-v1:0"

        # Separate system messages for Bedrock Converse API format
        system_prompts = []
        messages_payload = []
        for msg in request.messages:
            if msg.role == "system":
                system_prompts.append({"text": msg.content})
            else:
                messages_payload.append({
                    "role": "assistant" if msg.role == "assistant" else "user",
                    "content": [{"text": msg.content}],
                })

        try:
            client = self._get_client()
            
            # Map request fields into Converse payload
            converse_params = {
                "modelId": model,
                "messages": messages_payload,
                "inferenceConfig": {
                    "temperature": request.temperature,
                }
            }
            if request.max_tokens:
                converse_params["inferenceConfig"]["maxTokens"] = request.max_tokens
            if system_prompts:
                converse_params["system"] = system_prompts

            # Run in executor thread since boto3 client calls are synchronous
            # This prevents blocking the async main execution thread
            loop = asyncio.get_running_loop()
            response = await loop.run_in_executor(
                None,
                lambda: client.converse(**converse_params)
            )

            content_text = response["output"]["message"]["content"][0]["text"]
            latency = (time.perf_counter() - start_time) * 1000

            # Estimate token parameters
            usage_data = response.get("usage", {})
            usage = self.estimate_usage(
                "".join(m.content for m in request.messages),
                content_text,
                model,
            )
            if usage_data:
                usage.prompt_tokens = usage_data.get("inputTokens", usage.prompt_tokens)
                usage.completion_tokens = usage_data.get("outputTokens", usage.completion_tokens)
                usage.total_tokens = usage.prompt_tokens + usage.completion_tokens

            return ChatResponse(
                id=f"bedrock-{uuid.uuid4()}",
                model=model,
                provider="bedrock",
                choices=[
                    Choice(
                        index=0,
                        message=Message(role="assistant", content=content_text),
                        finish_reason=response.get("stopReason"),
                    )
                ],
                usage=usage,
                latency_ms=round(latency, 2),
            )
        except Exception as e:
            # If boto3/credentials are missing in local dev, fallback gracefully or raise
            raise DatabaseException(f"AWS Bedrock Converse API failure: {e}") from e

    async def generate_stream(self, request: ChatRequest) -> AsyncGenerator[ChatResponse, None]:
        model = request.model or "meta.llama3-8b-instruct-v1:0"
        yield ChatResponse(
            id="chatcmpl-bedrock-stream",
            model=model,
            provider="bedrock",
            choices=[
                Choice(
                    index=0,
                    message=Message(role="assistant", content="Bedrock stream started..."),
                    finish_reason=None,
                )
            ],
            usage=self.estimate_usage("", "", model),
            latency_ms=10.0,
        )

    async def generate_embeddings(self, request: EmbeddingRequest) -> EmbeddingResponse:
        raise NotImplementedError("Bedrock embeddings API is not implemented.")
