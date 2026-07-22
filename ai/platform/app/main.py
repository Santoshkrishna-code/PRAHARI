from contextlib import asynccontextmanager
from typing import AsyncGenerator
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse
import structlog

from app.core.config import settings
from app.core.exceptions import BasePlatformException
from app.infrastructure.logger import initialize_logging
from app.infrastructure.cache import cache_provider
from app.infrastructure.telemetry import initialize_telemetry
from app.infrastructure.registry import registry
from app.interfaces.api import router as api_router

# Import Middleware Classes
from app.interfaces.middlewares.request_id import RequestIdMiddleware
from app.interfaces.middlewares.logging import LoggingMiddleware
from app.interfaces.middlewares.telemetry import TelemetryContextMiddleware
from app.interfaces.middlewares.auth import AuthenticationMiddleware

# Configure logger
logger = structlog.get_logger(__name__)


@asynccontextmanager
async def app_lifespan(app: FastAPI) -> AsyncGenerator[None, None]:
    """Manages application startup and shutdown hook lifecycles."""
    # Startup Activities
    initialize_logging()
    logger.info("Starting PRAHARI AI Platform Foundation...")

    # Initialize cache connection pool
    cache_provider.initialize()
    await cache_provider.check_health()

    # Register providers in Dependency Injection registry
    registry.register("cache", lambda: cache_provider)

    # Initialize OpenTelemetry instrumentation
    initialize_telemetry(app)

    # Auto-initialize database tables schema (for prompt engineering registry)
    from app.infrastructure.repositories.prompt_repository import prompt_repository
    await prompt_repository.initialize_tables()

    # Auto-initialize database tables schema (for memory sessions and history)
    from app.infrastructure.repositories.memory_repository import memory_repository
    await memory_repository.initialize_tables()

    # Auto-initialize database tables schema (for RAG knowledge repository)
    from app.infrastructure.repositories.knowledge_repository import knowledge_repository
    await knowledge_repository.initialize_tables()

    # Auto-initialize database tables schema (for Tool calling platform)
    from app.infrastructure.repositories.tool_repository import tool_repository
    await tool_repository.initialize_tables()

    # Auto-initialize database tables schema (for Agent workflow graphs)
    from app.infrastructure.repositories.agent_repository import agent_repository
    await agent_repository.initialize_tables()

    # Auto-initialize database tables schema (for Safety Guardrails platform)
    from app.infrastructure.repositories.guardrail_repository import guardrail_repository
    await guardrail_repository.initialize_tables()

    # Auto-initialize database tables schema (for AI Evaluation platform)
    from app.infrastructure.repositories.evaluation_repository import evaluation_repository
    await evaluation_repository.initialize_tables()

    # Auto-initialize database tables schema (for Incident Investigation platform)
    from app.infrastructure.repositories.incident_repository import incident_repository
    await incident_repository.initialize_tables()

    # Auto-initialize database tables schema (for Safety Compliance platform)
    from app.infrastructure.repositories.compliance_repository import compliance_repository
    await compliance_repository.initialize_tables()

    # Auto-initialize database tables schema (for Permit-to-Work platform)
    from app.infrastructure.repositories.permit_repository import permit_repository
    await permit_repository.initialize_tables()

    # Auto-initialize database tables schema (for Enterprise Risk Assessment platform)
    from app.infrastructure.repositories.risk_repository import risk_repository
    await risk_repository.initialize_tables()

    # Auto-initialize database tables schema (for Intelligent Inspection platform)
    from app.infrastructure.repositories.inspection_repository import inspection_repository
    await inspection_repository.initialize_tables()

    # Auto-initialize database tables schema (for Contractor Safety platform)
    from app.infrastructure.repositories.contractor_repository import contractor_repository
    await contractor_repository.initialize_tables()

    # Auto-initialize database tables schema (for Asset & Equipment Safety platform)
    from app.infrastructure.repositories.asset_repository import asset_repository
    await asset_repository.initialize_tables()

    # Auto-initialize database tables schema (for Emergency Response platform)
    from app.infrastructure.repositories.emergency_repository import emergency_repository
    await emergency_repository.initialize_tables()

    # Auto-initialize database tables schema (for Predictive Maintenance platform)
    from app.infrastructure.repositories.maintenance_repository import maintenance_repository
    await maintenance_repository.initialize_tables()

    # Auto-initialize database tables schema (for Executive Reporting & Analytics platform)
    from app.infrastructure.repositories.analytics_repository import analytics_repository
    await analytics_repository.initialize_tables()

    # Auto-initialize database tables schema (for Multi-Agent Supervisor & Collaboration platform)
    from app.infrastructure.repositories.supervisor_repository import supervisor_repository
    await supervisor_repository.initialize_tables()

    logger.info("PRAHARI AI Platform Foundation started successfully.")
    yield

    # Shutdown Activities
    logger.info("Shutting down PRAHARI AI Platform Foundation...")
    await cache_provider.close()
    registry.clear()
    logger.info("PRAHARI AI Platform Foundation shutdown completed.")


# Create FastAPI App Instance
app = FastAPI(
    title=settings.APP_NAME,
    version="1.0.0",
    description="Volume 0: Foundation for PRAHARI AI Platform",
    lifespan=app_lifespan,
    docs_url="/docs",
    redoc_url="/redoc",
    openapi_url="/openapi.json",
)

# Register Middlewares (FIFO order)
app.add_middleware(AuthenticationMiddleware)
app.add_middleware(TelemetryContextMiddleware)
app.add_middleware(LoggingMiddleware)
app.add_middleware(RequestIdMiddleware)

# CORS middleware for S3-hosted frontend
from fastapi.middleware.cors import CORSMiddleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Register Exception Handlers mapping Domain exceptions to RFC 7807 Standard Responses
@app.exception_handler(BasePlatformException)
async def platform_exception_handler(request: Request, exc: BasePlatformException) -> JSONResponse:
    """Exception mapper for platform specific errors."""
    return JSONResponse(
        status_code=exc.status_code,
        content={
            "type": f"https://prahari.io/errors/{exc.__class__.__name__.lower()}",
            "title": exc.__class__.__name__,
            "status": exc.status_code,
            "detail": exc.message,
            "instance": request.url.path,
        },
    )


@app.exception_handler(Exception)
async def global_exception_handler(request: Request, exc: Exception) -> JSONResponse:
    """Fallback handler mapping generic Python failures."""
    logger.error("Unhandled platform system error occurred", error=str(exc))
    return JSONResponse(
        status_code=500,
        content={
            "type": "https://prahari.io/errors/internal-server-error",
            "title": "InternalServerError",
            "status": 500,
            "detail": "An unexpected error occurred. Please contact system support.",
            "instance": request.url.path,
        },
    )


# Attach Core API Endpoints Router
app.include_router(api_router)
