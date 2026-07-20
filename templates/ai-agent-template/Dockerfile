# ==============================================================================
# Base Image for Python Agent Service
# ==============================================================================
FROM python:3.10-slim AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Install Poetry
RUN curl -sSL https://install.python-poetry.org | python3 -
ENV PATH="/root/.local/bin:$PATH"

COPY pyproject.toml poetry.lock* ./

# Configure poetry to write virtualenvs to local directory
RUN poetry config virtualenvs.in-project true && \
    poetry install --no-interaction --no-ansi --no-root --only main

# ==============================================================================
# Production Image
# ==============================================================================
FROM python:3.10-slim AS runner

WORKDIR /app

# Run as non-root user for security compliance
RUN groupadd -g 999 praharigroup && \
    useradd -r -u 999 -g praharigroup prahariuser
    
COPY --from=builder /app/.venv /app/.venv
COPY . /app

ENV PATH="/app/.venv/bin:$PATH"
USER prahariuser

EXPOSE 8000

CMD ["python", "agent.py"]
