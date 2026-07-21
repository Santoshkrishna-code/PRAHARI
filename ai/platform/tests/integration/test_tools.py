import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_tool_calling_and_mcp_handshake_lifecycle(test_client: TestClient) -> None:
    """Assert tool registrations, validations checks, SQL/REST stubs, discovery syntax, and MCP initialization."""
    # 1. Register a REST Tool
    rest_tool = {
        "name": "check_weather",
        "description": "Fetch current weather metrics for plant coordinates.",
        "category": "EHS-Environment",
        "type": "REST",
        "version_string": "1.0.0",
        "input_schema": {
            "required": ["latitude", "longitude"],
            "properties": {
                "latitude": {"type": "number"},
                "longitude": {"type": "number"},
            },
        },
        "execution_target": "https://mock-api.prahari/weather",
    }
    response = test_client.post("/tools/register", json=rest_tool)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    assert json_data["data"]["name"] == "check_weather"

    # 2. Register an Internal Service Tool
    internal_tool = {
        "name": "create_incident_report",
        "description": "Logs safety incident reports directly to microservices.",
        "category": "EHS-Management",
        "type": "INTERNAL",
        "version_string": "1.0.0",
        "input_schema": {
            "required": ["description", "location"],
            "properties": {
                "description": {"type": "string"},
                "location": {"type": "string"},
            },
        },
        "execution_target": "create_incident",
    }
    response = test_client.post("/tools/register", json=internal_tool)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True

    # 3. List Registered Tools
    response = test_client.get("/tools")
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["data"]) >= 2

    # 4. Discover Tool schemas (OpenAI Function Calling format)
    response = test_client.post("/tools/discover")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    openai_tools = json_data["data"]
    assert len(openai_tools) >= 2
    assert openai_tools[0]["type"] == "function"

    # 5. Execute REST Tool Sandbox
    execute_rest = {
        "tool_name": "check_weather",
        "payload": {"latitude": 48.13, "longitude": 11.58},
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/tools/execute", json=execute_rest)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    assert json_data["output"]["status"] == "MOCKED_REST_RESPONSE"

    # 6. Execute Internal Tool Sandbox
    execute_internal = {
        "tool_name": "create_incident_report",
        "payload": {
            "description": "Hazardous gas alarm sounded on tank flow sensor C2.",
            "location": "Sector 4-B",
        },
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/tools/execute", json=execute_internal)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    assert "INC-" in json_data["output"]["incident_id"]
    assert json_data["output"]["status"] == "OPEN"

    # 7. Execute Validation error scan (missing parameters check)
    execute_faulty = {
        "tool_name": "create_incident_report",
        "payload": {"description": "no location supplied"},
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/tools/execute", json=execute_faulty)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is False
    assert "Missing required parameter" in json_data["error"]

    # 8. MCP Initialization handshake and tools listings
    mcp_handshake = {
        "client_name": "MockAgentClient",
        "client_version": "1.0.0",
        "protocol_version": "1.0",
    }
    response = test_client.post("/mcp/initialize", json=mcp_handshake)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["protocol_version"] == "1.0"
    assert "tools" in json_data["capabilities"]

    # 9. MCP List tools check
    response = test_client.get("/mcp/tools")
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["tools"]) >= 2
