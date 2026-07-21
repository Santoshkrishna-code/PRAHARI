import pytest
from unittest.mock import AsyncMock, MagicMock
from app.core.models import ChatRequest, Message, EmbeddingRequest
from app.infrastructure.providers.factory import provider_factory
from app.core.exceptions import DatabaseException
from app.core.model_registry import MODEL_REGISTRY, ModelMetadata
from app.infrastructure.cache import cache_provider


@pytest.mark.asyncio
async def test_mock_provider_chat_generation() -> None:
    """Assert Mock Provider returns standard mocked messages."""
    provider = provider_factory.get_provider("mock")
    request = ChatRequest(
        messages=[Message(role="user", content="Test safety scan")],
        model="mock-model",
    )
    
    response = await provider.generate_chat(request)
    assert response.provider == "mock"
    assert "Mock EHS response" in response.choices[0].message.content
    assert response.usage.prompt_tokens > 0
    assert response.usage.cost == 0.0


@pytest.mark.asyncio
async def test_execute_chat_with_fallback() -> None:
    """Assert factory fails over to working models when prior providers crash."""
    # Register temporary failing model metadata
    MODEL_REGISTRY["failing-model"] = ModelMetadata(
        provider="openai",  # Will raise connection failure in mock test env
        name="failing-model",
        version="1.0",
        context_window=4096,
        supports_streaming=True,
        supports_vision=False,
        supports_function_calling=False,
        supports_embeddings=False,
        supports_json_mode=False,
        supports_audio=False,
        supports_images=False,
        max_tokens=2048,
        input_cost_per_million=1.0,
        output_cost_per_million=2.0,
        latency_tier="low",
    )

    request = ChatRequest(
        messages=[Message(role="user", content="Simulate database failure")],
        model="failing-model",
        fallback_chain=["mock-model"],
        use_cache=False,
    )

    # Execution should bypass failing-model and return success via fallback mock-model
    response = await provider_factory.execute_chat(request)
    assert response.model == "mock-model"
    assert response.provider == "mock"
    
    # Cleanup registry
    del MODEL_REGISTRY["failing-model"]


@pytest.mark.asyncio
async def test_redis_caching_hit(mock_redis_provider: MagicMock) -> None:
    """Assert second identical request yields cached response payload."""
    # First call - cache miss, query mock provider, serialize result to Redis
    request = ChatRequest(
        messages=[Message(role="user", content="Fetch safety code index")],
        model="mock-model",
        use_cache=True,
    )

    # Configure mock Redis to return None on first get, then stored string on subsequent get
    cached_payload = []
    
    async def mock_get(key: str) -> str | None:
        if not cached_payload:
            return None
        return cached_payload[0]
        
    async def mock_set(key: str, val: str, ttl: int | None = None) -> bool:
        cached_payload.append(val)
        return True

    provider_factory._circuit_breaker.record_success("mock")
    # Bind mocks to cache provider directly
    orig_get = cache_provider.get
    orig_set = cache_provider.set
    cache_provider.get = mock_get
    cache_provider.set = mock_set

    try:
        res1 = await provider_factory.execute_chat(request)
        assert res1.cached is False

        # Second call - mock Redis now holds payload, should return cached=True
        res2 = await provider_factory.execute_chat(request)
        assert res2.cached is True
        assert res2.choices[0].message.content == res1.choices[0].message.content
    finally:
        # Restore original cache provider methods
        cache_provider.get = orig_get
        cache_provider.set = orig_set
