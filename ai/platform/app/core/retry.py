import asyncio
import functools
import logging
from collections.abc import Callable
from typing import Any, TypeVar, ParamSpec

logger = logging.getLogger(__name__)

T = TypeVar("T")
P = ParamSpec("P")


def retry(
    exceptions: tuple[type[Exception], ...] = (Exception,),
    tries: int = 3,
    delay: float = 1.0,
    backoff: float = 2.0,
) -> Callable[[Callable[P, Any]], Callable[P, Any]]:
    """Decorator to retry an async function with exponential backoff."""

    def decorator(func: Callable[P, Any]) -> Callable[P, Any]:
        @functools.wraps(func)
        async def wrapper(*args: P.args, **kwargs: P.kwargs) -> Any:
            attempt_delay = delay
            for attempt in range(1, tries + 1):
                try:
                    return await func(*args, **kwargs)
                except exceptions as e:
                    if attempt == tries:
                        logger.error(
                            "Failed to execute %s after %d attempts. Exception: %s",
                            func.__name__,
                            tries,
                            str(e),
                        )
                        raise e
                    logger.warning(
                        "Attempt %d/%d for %s failed with exception: %s. Retrying in %ss...",
                        attempt,
                        tries,
                        func.__name__,
                        str(e),
                        attempt_delay,
                    )
                    await asyncio.sleep(attempt_delay)
                    attempt_delay *= backoff

        return wrapper

    return decorator
