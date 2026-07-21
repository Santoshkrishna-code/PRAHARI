import uuid
from starlette.middleware.base import BaseHTTPMiddleware, RequestResponseEndpoint
from fastapi import Request, Response
import structlog


class RequestIdMiddleware(BaseHTTPMiddleware):
    """Middleware to inject or generate unique request-id trace headers."""

    async def dispatch(self, request: Request, call_next: RequestResponseEndpoint) -> Response:
        # Check if incoming request has correlation header
        request_id = request.headers.get("X-Request-ID")
        if not request_id:
            request_id = str(uuid.uuid4())

        # Seed correlation context into structured logger context variables
        structlog.contextvars.clear_contextvars()
        structlog.contextvars.bind_contextvars(request_id=request_id)

        # Set request-id inside request state for handler scopes
        request.state.request_id = request_id

        # Proceed with request pipeline execution
        response = await call_next(request)

        # Return request ID in outgoing headers
        response.headers["X-Request-ID"] = request_id
        return response
