import pytest
from unittest.mock import MagicMock
from app.core.config import settings
from app.core.exceptions import DatabaseException, CacheException
from app.core.retry import retry


def test_platform_settings_loading() -> None:
    """Assert settings contain expected keys and fallback defaults."""
    assert settings.APP_NAME == "prahari-ai-platform"
    assert settings.LOG_LEVEL in ["INFO", "DEBUG", "WARNING", "ERROR"]
    assert isinstance(settings.DEBUG, bool)


def test_core_exceptions_mapping() -> None:
    """Assert exception schemas hold parameters."""
    db_ex = DatabaseException("DB disconnect")
    assert db_ex.status_code == 503
    assert db_ex.message == "DB disconnect"

    cache_ex = CacheException("Redis crash")
    assert cache_ex.status_code == 503
    assert cache_ex.message == "Redis crash"


@pytest.mark.asyncio
async def test_retry_decorator_backoff() -> None:
    """Assert retry utility executes callbacks on failure states."""
    call_count = 0

    @retry(exceptions=(ValueError,), tries=3, delay=0.01, backoff=1.5)
    async def mock_async_function() -> str:
        nonlocal call_count
        call_count += 1
        if call_count < 3:
            raise ValueError("Intermittent crash")
        return "success"

    result = await mock_async_function()
    assert result == "success"
    assert call_count == 3
