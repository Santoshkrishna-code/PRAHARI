import logging
from typing import Any, Callable

logger = logging.getLogger(__name__)


class ProviderRegistry:
    """Central Dependency Injection registry container class."""

    def __init__(self) -> None:
        self._providers: dict[type[Any] | str, Callable[..., Any]] = {}
        self._instances: dict[type[Any] | str, Any] = {}

    def register(
        self, key: type[Any] | str, provider_factory: Callable[..., Any]
    ) -> None:
        """Register provider constructor callback."""
        logger.debug("Registering provider factory for: %s", str(key))
        self._providers[key] = provider_factory

    def get(self, key: type[Any] | str) -> Any:
        """Fetch instantiated provider (Singleton implementation)."""
        if key in self._instances:
            return self._instances[key]

        if key not in self._providers:
            raise KeyError(f"No provider registered for type or key: {key}")

        # Instantiate provider using registered factory
        factory = self._providers[key]
        instance = factory()
        self._instances[key] = instance
        return instance

    def clear(self) -> None:
        """Clear all instances and registrations (useful for testing mock overrides)."""
        logger.debug("Clearing provider registry dependencies...")
        self._providers.clear()
        self._instances.clear()


registry = ProviderRegistry()
