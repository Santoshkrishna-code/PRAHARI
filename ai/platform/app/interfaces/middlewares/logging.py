import time
from starlette.middleware.base import BaseHTTPMiddleware, RequestResponseEndpoint
from fastapi import Request, Response
import structlog

from app.infrastructure.telemetry import HTTP_REQUEST_COUNT, HTTP_REQUEST_LATENCY

logger = structlog.get_logger(__name__)


class LoggingMiddleware(BaseHTTPMiddleware):
    """Middleware to intercept request latencies and record Prometheus counters."""

    async def dispatch(self, request: Request, call_next: RequestResponseEndpoint) -> Response:
        path = request.url.path
        method = request.method
        start_time = time.perf_counter()

        logger.info(
            "Request started",
            path=path,
            method=method,
            client_ip=request.client.host if request.client else None,
        )

        try:
            response = await call_next(request)
            duration = time.perf_counter() - start_time
            status_code = response.status_code

            logger.info(
                "Request completed",
                path=path,
                method=method,
                status_code=status_code,
                duration_seconds=round(duration, 4),
            )

            # Record Prometheus metrics
            HTTP_REQUEST_COUNT.labels(
                method=method, endpoint=path, status_code=str(status_code)
            ).inc()
            HTTP_REQUEST_LATENCY.labels(method=method, endpoint=path).observe(duration)

            return response

        except Exception as e:
            duration = time.perf_counter() - start_time
            logger.error(
                "Request failed with unhandled exception",
                path=path,
                method=method,
                duration_seconds=round(duration, 4),
                error=str(e),
            )
            # Record failed request counts
            HTTP_REQUEST_COUNT.labels(
                method=method, endpoint=path, status_code="500"
            ).inc()
            raise e
