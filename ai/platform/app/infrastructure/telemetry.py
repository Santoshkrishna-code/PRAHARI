import logging
from fastapi import FastAPI
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.instrumentation.fastapi import FastAPIInstrumentor
from prometheus_client import CollectorRegistry, Counter, Histogram, generate_latest

from app.core.config import settings

logger = logging.getLogger(__name__)

# Prometheus Metrics Registry Setup
prometheus_registry = CollectorRegistry()

# Default Platform Metrics
HTTP_REQUEST_COUNT = Counter(
    name="platform_http_requests_total",
    documentation="Total count of HTTP requests processed by the platform",
    labelnames=["method", "endpoint", "status_code"],
    registry=prometheus_registry,
)

HTTP_REQUEST_LATENCY = Histogram(
    name="platform_http_request_latency_seconds",
    documentation="HTTP request processing latency in seconds",
    labelnames=["method", "endpoint"],
    registry=prometheus_registry,
)


def initialize_telemetry(app: FastAPI) -> None:
    """Setup OpenTelemetry Tracing and Instrument FastAPI."""
    if not settings.ENABLE_TELEMETRY:
        logger.info("OpenTelemetry tracing is disabled.")
        return

    logger.info("Initializing OpenTelemetry Tracer Provider...")
    resource = Resource.create(
        attributes={
            "service.name": settings.OTEL_SERVICE_NAME,
            "deployment.environment": settings.APP_ENV,
        }
    )

    provider = TracerProvider(resource=resource)
    
    # Configure OTLP Exporter pointing to collector endpoint
    otlp_exporter = OTLPSpanExporter(endpoint=settings.OTEL_EXPORTER_OTLP_ENDPOINT)
    span_processor = BatchSpanProcessor(otlp_exporter)
    provider.add_span_processor(span_processor)
    
    trace.set_tracer_provider(provider)

    # FastAPI Automatic Instrumentation
    FastAPIInstrumentor.instrument_app(app)
    logger.info("OpenTelemetry tracer instrumentation completed successfully.")


def get_metrics_payload() -> bytes:
    """Formats metrics in Prometheus format."""
    return generate_latest(prometheus_registry)
