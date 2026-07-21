from typing import Literal
from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        extra="ignore",
    )

    APP_NAME: str = "prahari-ai-platform"
    APP_ENV: Literal["development", "production", "testing"] = "development"
    PORT: int = 8000
    DEBUG: bool = True

    # Security
    JWT_SECRET_KEY: str = "secure_development_jwt_secret_key_change_in_production_32_bytes"
    JWT_ALGORITHM: str = "HS256"

    # Database & Redis
    DATABASE_URL: str = (
        "postgresql+asyncpg://postgres:postgres_secure_pass@localhost:5432/prahari_db"
    )
    REDIS_URL: str = "redis://:redis_secure_pass@localhost:6379/0"

    # LLM Provider Keys and Configurations
    OPENAI_API_KEY: str = "mock-openai-key"
    OPENAI_API_BASE: str | None = None
    
    ANTHROPIC_API_KEY: str = "mock-anthropic-key"
    
    GEMINI_API_KEY: str = "mock-gemini-key"
    
    AZURE_OPENAI_API_KEY: str | None = None
    AZURE_OPENAI_ENDPOINT: str | None = None
    
    AWS_ACCESS_KEY_ID: str | None = None
    AWS_SECRET_ACCESS_KEY: str | None = None
    AWS_REGION: str = "us-east-1"
    
    OLLAMA_BASE_URL: str = "http://localhost:11434"
    VLLM_BASE_URL: str = "http://localhost:8000/v1"
    
    DEFAULT_PROVIDER: str = "openai"
    DEFAULT_MODEL: str = "gpt-4o"
    FALLBACK_MODELS: list[str] = ["gpt-4o", "claude-3-5-sonnet", "gemini-1.5-pro", "llama3"]

    # Telemetry & Logging
    OTEL_EXPORTER_OTLP_ENDPOINT: str = "http://localhost:4317"
    OTEL_SERVICE_NAME: str = "prahari-ai-platform"
    ENABLE_TELEMETRY: bool = False
    LOG_LEVEL: str = "INFO"
    JSON_LOGS: bool = True


settings = Settings()
