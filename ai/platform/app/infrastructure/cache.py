import logging
from typing import Any
import redis.asyncio as redis
from app.core.config import settings
from app.core.exceptions import CacheException

logger = logging.getLogger(__name__)


class CacheProvider:
    """Async Redis operations provider wrapper class."""

    def __init__(self) -> None:
        self._client: redis.Redis[Any] | None = None

    def initialize(self) -> None:
        """Create redis client connection pool."""
        logger.info("Initializing Redis connection pool...")
        self._client = redis.from_url(
            settings.REDIS_URL,
            encoding="utf-8",
            decode_responses=True,
        )

    async def close(self) -> None:
        """Close redis client session."""
        if self._client:
            logger.info("Closing Redis connection pool...")
            await self._client.close()

    async def get(self, key: str) -> str | None:
        """Fetch string value from Redis."""
        if not self._client:
            raise CacheException("Redis client is not initialized")
        try:
            value: str | None = await self._client.get(key)
            return value
        except Exception as e:
            logger.error("Redis GET failed for key %s: %s", key, str(e))
            raise CacheException(f"Redis GET failed: {e}") from e

    async def set(self, key: str, value: str, ttl: int | None = None) -> bool:
        """Set string value in Redis with optional TTL."""
        if not self._client:
            raise CacheException("Redis client is not initialized")
        try:
            res: bool = await self._client.set(key, value, ex=ttl)
            return res
        except Exception as e:
            logger.error("Redis SET failed for key %s: %s", key, str(e))
            raise CacheException(f"Redis SET failed: {e}") from e

    async def check_health(self) -> bool:
        """Executes ping command to verify Redis connectivity."""
        if not self._client:
            return False
        try:
            return await self._client.ping()
        except Exception as e:
            logger.error("Redis ping check failed: %s", str(e))
            return False


cache_provider = CacheProvider()
