import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_agent_graph_execution_and_human_signoff(test_client: TestClient) -> None:
    """Assert multi-agent workflow compilation, conditional routing, state checkpoints, and sign-off pauses."""
    # 1. Register tools required by tool nodes
    rest_tool = {
        "name": "check_weather",
        "description": "Fetch weather data.",
        "category": "EHS",
        "type": "REST",
        "version_string": "1.0.0",
        "input_schema": {"properties": {}},
        "execution_target": "https://mock-api.prahari/weather",
    }
    test_client.post("/tools/register", json=rest_tool)

    # 2. Register Agent Graph configuration
    graph_payload = {
        "name": "ehs_incident_investigation",
        "description": "Orchestrates safety audits and risk evaluations across agents.",
        "start_node": "incident_classifier",
        "nodes": [
            {"name": "incident_classifier", "model": "gpt-4o"},
            {"name": "tool_runner", "tools_allowed": ["check_weather"], "model": "gpt-4o"},
            {"name": "supervisor_approval", "model": "gpt-4o"},
            {"name": "auto_resolver", "model": "gpt-4o"},
        ],
        "edges": [
            {
                "source_node": "incident_classifier",
                "target_node": "supervisor_approval",
                "conditional_expr": "risk_is_critical",
            },
            {
                "source_node": "incident_classifier",
                "target_node": "tool_runner",
                "conditional_expr": "risk_is_low",
            },
            {"source_node": "tool_runner", "target_node": "auto_resolver"},
            {"source_node": "supervisor_approval", "target_node": "auto_resolver"},
        ],
    }
    
    response = test_client.post(
        f"/agent/graphs?name={graph_payload['name']}&description={graph_payload['description']}&start_node={graph_payload['start_node']}",
        json=graph_payload,
    )
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    graph_id = json_data["data"]["id"]

    # 3. Trigger LOW-risk execution run (routes incident -> tool_runner -> auto_resolver)
    low_risk_run = {
        "graph_id": graph_id,
        "input_state": {"incident_desc": "Minor water puddle on floor"},
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/agent/run", json=low_risk_run)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "COMPLETED"
    assert json_data["current_node"] == "auto_resolver"
    assert json_data["state"]["risk_level"] == "LOW"
    assert "MOCKED_REST_RESPONSE" in json_data["state"]["weather_data"]["status"]

    # 4. Trigger CRITICAL-risk execution run (routes incident -> supervisor_approval -> halts)
    critical_risk_run = {
        "graph_id": graph_id,
        "input_state": {"incident_desc": "Toxic gas release at tank C2"},
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/agent/run", json=critical_risk_run)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "INTERRUPTED"
    assert json_data["current_node"] == "supervisor_approval"
    assert json_data["state"]["risk_level"] == "CRITICAL"
    execution_id = json_data["execution_id"]

    # 5. Check checkpoint status retrieval
    response = test_client.get(f"/agent/status/{execution_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["data"]["status"] == "INTERRUPTED"
    assert json_data["data"]["current_node"] == "supervisor_approval"

    # 6. Resume/Approve human sign-off (resumes and transitions supervisor_approval -> auto_resolver)
    approval_payload = {
        "execution_id": execution_id,
        "approved": True,
        "comments": "Approved. Safe to resolve.",
    }
    response = test_client.post("/agent/approve", json=approval_payload)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "COMPLETED"
    assert json_data["current_node"] == "auto_resolver"
    assert json_data["state"]["reviewer_approved"] is True
    assert json_data["state"]["reviewer_comments"] == "Approved. Safe to resolve."
