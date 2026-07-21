import logging
import sys
import structlog
from app.core.config import settings


def initialize_logging() -> None:
    """Configures structured JSON logging format based on settings."""
    log_level = getattr(logging, settings.LOG_LEVEL.upper(), logging.INFO)

    shared_processors: list[structlog.types.Processor] = [
        structlog.contextvars.merge_contextvars,
        structlog.processors.add_log_level,
        structlog.processors.TimeStamper(fmt="iso"),
        structlog.processors.StackInfoRenderer(),
        structlog.processors.format_exc_info,
    ]

    if settings.JSON_LOGS:
        # JSON logs configuration for Production
        processors = shared_processors + [
            structlog.processors.JSONRenderer(),
        ]
    else:
        # Colorized console logging for Local Development
        processors = shared_processors + [
            structlog.dev.ConsoleRenderer(),
        ]

    structlog.configure(
        processors=processors,
        logger_factory=structlog.PrintLoggerFactory(),
        wrapper_class=structlog.make_filtering_bound_logger(log_level),
        cache_logger_on_first_use=True,
    )

    # Bridge standard logging to use structlog formatting
    logging.basicConfig(
        format="%(message)s",
        stream=sys.stdout,
        level=log_level,
    )
    
    # Disable duplicate logs from underlying uvicorn components
    for uvicorn_logger in ["uvicorn", "uvicorn.error", "uvicorn.access"]:
        log = logging.getLogger(uvicorn_logger)
        log.handlers = []
        log.propagate = True
