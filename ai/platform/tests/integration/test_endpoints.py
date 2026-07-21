from fastapi.testclient import TestClient
from jose import jwt
from app.core.config import settings


def test_health_endpoint_response(test_client: TestClient) -> None:
    """Assert health endpoint returns correct parameters."""
    response = test_client.get("/health")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "OK"
    assert json_data["database"] == "HEALTHY"
    assert json_data["cache"] == "HEALTHY"


def test_metrics_endpoint_response(test_client: TestClient) -> None:
    """Assert metrics endpoint returns text data."""
    response = test_client.get("/metrics")
    assert response.status_code == 200
    assert "text/plain" in response.headers["content-type"]


def test_unauthorized_ping_endpoint(test_client: TestClient) -> None:
    """Assert protected path yields 401 on missing auth details."""
    response = test_client.get("/api/v1/ping")
    assert response.status_code == 401
    assert "WWW-Authenticate" in response.headers


def test_authorized_ping_endpoint(test_client: TestClient) -> None:
    """Assert protected endpoint returns success when valid JWT is supplied."""
    payload = {"sub": "usr-test-12", "role": "engineer"}
    token = jwt.encode(payload, settings.JWT_SECRET_KEY, algorithm=settings.JWT_ALGORITHM)

    headers = {
        "Authorization": f"Bearer {token}",
        "X-Tenant-ID": "tenant-99",
    }
    response = test_client.get("/api/v1/ping", headers=headers)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    assert json_data["data"] == "pong"
