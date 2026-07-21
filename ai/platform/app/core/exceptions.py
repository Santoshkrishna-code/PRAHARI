class BasePlatformException(Exception):
    """Base class for all platform exceptions."""

    def __init__(self, message: str, status_code: int = 500) -> None:
        self.message = message
        self.status_code = status_code
        super().__init__(message)


class DatabaseException(BasePlatformException):
    """Exception raised for database operations failures."""

    def __init__(self, message: str) -> None:
        super().__init__(message, status_code=503)


class CacheException(BasePlatformException):
    """Exception raised for Redis cache operations failures."""

    def __init__(self, message: str) -> None:
        super().__init__(message, status_code=503)


class UnauthorizedException(BasePlatformException):
    """Exception raised for token validation or authentication failures."""

    def __init__(self, message: str = "Invalid credentials or unauthorized access") -> None:
        super().__init__(message, status_code=401)


class NotFoundException(BasePlatformException):
    """Exception raised when a requested resource is not found."""

    def __init__(self, message: str = "Resource not found") -> None:
        super().__init__(message, status_code=404)


class ValidationException(BasePlatformException):
    """Exception raised for validation errors on request inputs."""

    def __init__(self, message: str = "Request validation failed") -> None:
        super().__init__(message, status_code=422)
