from starlette.middleware.base import BaseHTTPMiddleware, RequestResponseEndpoint
from fastapi import Request, Response
from opentelemetry import trace
import structlog


class TelemetryContextMiddleware(BaseHTTPMiddleware):
    """Middleware to bind active OpenTelemetry trace context to structlog variables."""

    async def dispatch(self, request: Request, call_next: RequestResponseEndpoint) -> Response:
        current_span = trace.get_current_span()
        if current_span:
            span_context = current_span.get_span_context()
            if span_context and span_context.is_valid:
                # Bind trace parameters so they map to structured JSON log keys
                structlog.contextvars.bind_contextvars(
                    trace_id=format(span_context.trace_id, "032x"),
                    span_id=format(span_context.span_id, "016x"),
                )

        return await call_next(request)
