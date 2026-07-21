import time
import json
import logging
import hashlib
from collections.abc import AsyncGenerator
from typing import Any
from pydantic import BaseModel

from app.core.config import settings
from app.core.exceptions import DatabaseException, NotFoundException
from app.core.model_registry import MODEL_REGISTRY
from app.core.models import (
    ChatRequest,
    ChatResponse,
    Choice,
    Message,
    EmbeddingRequest,
    EmbeddingResponse,
    Usage,
)
from app.infrastructure.cache import cache_provider

# Import Adapters
from app.infrastructure.providers.base import BaseProvider
from app.infrastructure.providers.openai import OpenAIProvider
from app.infrastructure.providers.anthropic import AnthropicProvider
from app.infrastructure.providers.gemini import GeminiProvider
from app.infrastructure.providers.bedrock import BedrockProvider
from app.infrastructure.providers.ollama import OllamaProvider
from app.infrastructure.providers.mock import MockProvider

logger = logging.getLogger(__name__)


class CircuitBreaker:
    """Tracks provider failures and temporarily trips access to prevent timeouts."""

    def __init__(self, failure_threshold: int = 3, recovery_time_seconds: int = 60) -> None:
        self.failure_threshold = failure_threshold
        self.recovery_time_seconds = recovery_time_seconds
        self.failure_counts: dict[str, int] = {}
        self.tripped_until: dict[str, float] = {}

    def is_tripped(self, provider: str) -> bool:
        now = time.time()
        if provider in self.tripped_until:
            if now < self.tripped_until[provider]:
                return True
            # Cool-down expired, reset state
            self.failure_counts[provider] = 0
            del self.tripped_until[provider]
        return False

    def record_failure(self, provider: str) -> None:
        self.failure_counts[provider] = self.failure_counts.get(provider, 0) + 1
        if self.failure_counts[provider] >= self.failure_threshold:
            self.tripped_until[provider] = time.time() + self.recovery_time_seconds
            logger.warning(
                "Circuit breaker tripped for provider '%s' until %s.",
                provider,
                self.tripped_until[provider],
            )

    def record_success(self, provider: str) -> None:
        self.failure_counts[provider] = 0


class ProviderFactory:
    """Manages adapter registration, resolution, failovers, and caching."""

    def __init__(self) -> None:
        self._providers: dict[str, BaseProvider] = {}
        self._circuit_breaker = CircuitBreaker()
        self._register_default_providers()

    def _register_default_providers(self) -> None:
        self._providers["openai"] = OpenAIProvider()
        self._providers["anthropic"] = AnthropicProvider()
        self._providers["gemini"] = GeminiProvider()
        self._providers["bedrock"] = BedrockProvider()
        self._providers["ollama"] = OllamaProvider()
        self._providers["mock"] = MockProvider()

    def get_provider(self, name: str) -> BaseProvider:
        """Fetch matching provider instance."""
        if name not in self._providers:
            raise NotFoundException(f"Provider '{name}' is not registered.")
        return self._providers[name]

    def resolve_provider_for_model(self, model: str) -> BaseProvider:
        """Looks up the associated provider name for a model registry entry."""
        metadata = MODEL_REGISTRY.get(model)
        if not metadata:
            raise NotFoundException(f"Model '{model}' is not registered.")
        return self.get_provider(metadata.provider)

    def _generate_cache_key(self, request_type: str, request_data: BaseModel) -> str:
        """Constructs a deterministic sha256 cache key for Redis lookup."""
        dumped = request_data.model_dump_json(exclude={"use_cache", "stream"})
        payload_hash = hashlib.sha256(dumped.encode("utf-8")).hexdigest()
        return f"ai:cache:{request_type}:{payload_hash}"

    async def execute_chat(self, request: ChatRequest) -> ChatResponse:
        """Executes chat generation with auto-failover, caching, and circuit breakers."""
        # Define execution path chain
        model_chain = [request.model] if request.model else []
        if request.fallback_chain:
            model_chain.extend(request.fallback_chain)
        else:
            model_chain.extend(settings.FALLBACK_MODELS)

        # De-duplicate fallback chain preserves order
        unique_chain = []
        for m in model_chain:
            if m and m not in unique_chain:
                unique_chain.append(m)

        if not unique_chain:
            unique_chain = [settings.DEFAULT_MODEL]

        # 1. Caching Layer Check
        cache_key = None
        if request.use_cache and not request.stream:
            # Construct key based on first requested model
            request.model = unique_chain[0]
            cache_key = self._generate_cache_key("chat", request)
            try:
                cached_val = await cache_provider.get(cache_key)
                if cached_val:
                    logger.info("Cache hit for request key: %s", cache_key)
                    cached_data = json.loads(cached_val)
                    response = ChatResponse.model_validate(cached_data)
                    response.cached = True
                    return response
            except Exception as e:
                logger.warning("Cache fetch error bypassed: %s", str(e))

        # 2. Sequential Failover Execution Chain
        last_exception = None
        for model in unique_chain:
            try:
                metadata = MODEL_REGISTRY.get(model)
                if not metadata:
                    logger.warning("Model '%s' not registered. Skipping...", model)
                    continue

                provider_name = metadata.provider
                
                # Check Circuit Breaker state
                if self._circuit_breaker.is_tripped(provider_name):
                    logger.warning(
                        "Skipping provider '%s' due to active tripped circuit breaker.",
                        provider_name,
                    )
                    continue

                provider = self.get_provider(provider_name)
                logger.info(
                    "Executing chat generation on model '%s' via provider '%s'...",
                    model,
                    provider_name,
                )

                # Clone request modifying target model parameter
                request.model = model
                response = await provider.generate_chat(request)
                
                # Record success to cool down circuit breaker metrics
                self._circuit_breaker.record_success(provider_name)

                # 3. Write Back to Cache
                if cache_key:
                    try:
                        await cache_provider.set(
                            cache_key, response.model_dump_json(), ttl=86400
                        )
                    except Exception as e:
                        logger.warning("Cache write failure bypassed: %s", str(e))

                return response

            except Exception as e:
                logger.error(
                    "Execution failed on model '%s' with exception: %s. Initiating failover...",
                    model,
                    str(e),
                )
                self._circuit_breaker.record_failure(MODEL_REGISTRY[model].provider)
                last_exception = e

        # All models exhausted
        raise DatabaseException(
            f"All LLM fallback options exhausted. Last exception: {last_exception}"
        )

    async def execute_embeddings(self, request: EmbeddingRequest) -> EmbeddingResponse:
        """Direct embedding resolver (caching enabled)."""
        model = request.model or "text-embedding-3-small"
        request.model = model

        # Caching check
        cache_key = self._generate_cache_key("embeddings", request)
        try:
            cached_val = await cache_provider.get(cache_key)
            if cached_val:
                cached_data = json.loads(cached_val)
                return EmbeddingResponse.model_validate(cached_data)
        except Exception as e:
            logger.warning("Cache error bypassed: %s", str(e))

        provider = self.resolve_provider_for_model(model)
        response = await provider.generate_embeddings(request)

        # Write cache
        try:
            await cache_provider.set(cache_key, response.model_dump_json(), ttl=86400)
        except Exception as e:
            logger.warning("Cache write failure bypassed: %s", str(e))

        return response


provider_factory = ProviderFactory()
