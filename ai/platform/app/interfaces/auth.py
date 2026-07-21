from jose import jwt, JWTError
from starlette.middleware.base import BaseHTTPMiddleware, RequestResponseEndpoint
from fastapi import Request, Response
from fastapi.responses import JSONResponse
import structlog

from app.core.config import settings

logger = structlog.get_logger(__name__)


class AuthenticationMiddleware(BaseHTTPMiddleware):
    """Middleware to parse JWT tokens and validate Tenant parameters on protected API scopes."""

    async def dispatch(self, request: Request, call_next: RequestResponseEndpoint) -> Response:
        path = request.url.path

        # Exempt infrastructure health checks and OpenAPI docs from authentication
        exempt_paths = ["/health", "/metrics", "/docs", "/openapi.json", "/redoc"]
        if any(path.startswith(p) for p in exempt_paths):
            return await call_next(request)

        # Require authentication for protected paths starting with /api/
        if path.startswith("/api/"):
            auth_header = request.headers.get("Authorization")
            tenant_header = request.headers.get("X-Tenant-ID")

            if not auth_header or not auth_header.startswith("Bearer "):
                return self._unauthorized_response("Missing or invalid Authorization header", path)

            if not tenant_header:
                return self._unauthorized_response("Missing X-Tenant-ID identification header", path)

            token = auth_header.split(" ")[1]
            try:
                # Decode and verify JWT signature
                payload = jwt.decode(
                    token,
                    settings.JWT_SECRET_KEY,
                    algorithms=[settings.JWT_ALGORITHM],
                )
                
                # Bind user context details to request state
                request.state.user = payload
                request.state.tenant_id = tenant_header
                structlog.contextvars.bind_contextvars(
                    user_id=payload.get("sub"), tenant_id=tenant_header
                )

            except JWTError as e:
                logger.warning("Token signature decoding failed", error=str(e))
                return self._unauthorized_response("Token decoding signature check failed", path)

        return await call_next(request)

    def _unauthorized_response(self, detail: str, path: str) -> JSONResponse:
        """Returns standard RFC 7807 error format."""
        return JSONResponse(
            status_code=401,
            content={
                "type": "https://prahari.io/errors/unauthorized",
                "title": "Unauthorized Access Validation Failure",
                "status": 401,
                "detail": detail,
                "instance": path,
            },
            headers={"WWW-Authenticate": "Bearer"},
        )
