# Multi-Stage Production Dockerfile for PRAHARI AI Runtime Platform
FROM python:3.13-slim AS builder

WORKDIR /app

# Install system compilation headers
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
COPY pyproject.toml .
RUN pip install --no-cache-dir --upgrade pip setuptools wheel
RUN pip install --no-cache-dir \
    fastapi>=0.115.0 \
    uvicorn>=0.30.0 \
    pydantic>=2.9.0 \
    pydantic-settings>=2.5.0 \
    sqlalchemy>=2.0.35 \
    asyncpg>=0.29.0 \
    aiosqlite>=0.20.0 \
    redis>=5.0.0 \
    httpx>=0.27.0 \
    pytest>=8.0.0 \
    structlog>=24.0.0

FROM python:3.13-slim AS runner

WORKDIR /app

# Copy python dependencies from builder
COPY --from=builder /usr/local/lib/python3.13/site-packages /usr/local/lib/python3.13/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin

# Install runtime utilities
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy application source code
COPY app/ app/
COPY tests/ tests/
COPY pyproject.toml .

# Create non-root runtime execution user
RUN useradd -m -u 1000 appuser && chown -R appuser:appuser /app
USER appuser

EXPOSE 8000

HEALTHCHECK --interval=15s --timeout=5s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:8000/health || exit 1

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
