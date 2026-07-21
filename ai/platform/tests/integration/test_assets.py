import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_asset_equipment_safety_workflow(test_client: TestClient) -> None:
    """Assert asset health scoring, inspection triggers, maintenance recommendations, Risk Agent calls, emergency shutdowns, and scorecards."""
    # 1. Evaluate Asset Health (compromised metrics triggering maintenance required)
    eval_req = {
        "asset_name": "High Pressure Boiler Unit 4",
        "asset_category": "BOILER",
        "location": "Boiler House East",
        "operating_hours": 1420.5,
        "vibration_mm_s": 5.2,
        "temperature_c": 98.0,
        "pressure_bar": 13.5,
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/asset/health", json=eval_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["asset_name"] == "High Pressure Boiler Unit 4"
    assert json_data["status"] == "MAINTENANCE_REQUIRED"
    assert json_data["health_score"] == 35.0
    asset_id = json_data["id"]

    # 2. Trigger Asset Inspection (calls Intelligent Inspection Agent)
    response = test_client.post(f"/asset/inspect?asset_id={asset_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["items"]) >= 3

    # 3. Generate Preventive Maintenance recommendation
    response = test_client.post(f"/asset/recommend?asset_id={asset_id}")
    assert response.status_code == 200
    json_data = response.json()["data"]
    assert json_data["urgency"] == "HIGH"

    # 4. Assess Asset Operational Risk (calls Enterprise Risk Agent)
    response = test_client.post(f"/asset/risk?asset_id={asset_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert "Asset Operational Risk" in json_data["title"]

    # 5. Trigger Emergency Shutdown
    shut_req = {
        "asset_id": asset_id,
        "reason": "Abnormal vibration and temperature threshold breach.",
    }
    response = test_client.post("/asset/shutdown", json=shut_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "SHUTDOWN"
    assert json_data["safety_lockout_issued"] is True

    # 6. Retrieve Asset Integrity Scorecard
    response = test_client.get(f"/asset/scorecard?asset_id={asset_id}")
    assert response.status_code == 200
    json_data = response.json()["data"]
    assert json_data["asset_id"] == asset_id
    assert json_data["operational_status"] == "SHUTDOWN"

    # 7. Fetch Asset History logs
    response = test_client.get("/asset/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["id"] == asset_id
