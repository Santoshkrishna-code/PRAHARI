import os
import pytest

# Force environment overrides before any application modules parse configurations
os.environ["DATABASE_URL"] = "sqlite+aiosqlite:///test_ai_platform.db"
os.environ["APP_ENV"] = "testing"

from unittest.mock import AsyncMock, MagicMock
from collections.abc import Generator
from fastapi.testclient import TestClient
from app.main import app
from app.infrastructure.cache import cache_provider


@pytest.fixture(autouse=True)
def mock_redis_provider() -> Generator[MagicMock, None, None]:
    """Mock cache provider behaviors to avoid connection failures in tests."""
    mock_instance = MagicMock()
    
    # Store original methods
    orig_initialize = cache_provider.initialize
    orig_close = cache_provider.close
    orig_get = cache_provider.get
    orig_set = cache_provider.set
    orig_check_health = cache_provider.check_health
    
    # Overwrite with mocks
    cache_provider.initialize = MagicMock()
    cache_provider.close = AsyncMock()
    cache_provider.get = AsyncMock(return_value=None)
    cache_provider.set = AsyncMock(return_value=True)
    cache_provider.check_health = AsyncMock(return_value=True)
    
    yield mock_instance
    
    # Restore original methods
    cache_provider.initialize = orig_initialize
    cache_provider.close = orig_close
    cache_provider.get = orig_get
    cache_provider.set = orig_set
    cache_provider.check_health = orig_check_health


@pytest.fixture
def test_client() -> Generator[TestClient, None, None]:
    """Test client fixture wrapper for sync endpoints assertions."""
    with TestClient(app) as client:
        yield client


@pytest.fixture(scope="session", autouse=True)
def cleanup_test_db() -> Generator[None, None, None]:
    yield
    # Delete the test SQLite database file if it exists
    if os.path.exists("test_ai_platform.db"):
        try:
            os.remove("test_ai_platform.db")
        except Exception:
            pass
