import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_emergency_response_workflow(test_client: TestClient) -> None:
    """Assert real-time emergency classification, evacuation route planning, responder alert notifications, SITREP briefs, emergency resources, and history."""
    # 1. Real-time Emergency Event Assessment (Chemical Spill LEVEL_3_CRITICAL)
    assess_req = {
        "title": "Industrial Chemical Spill in Tank Farm 2",
        "emergency_type": "CHEMICAL_SPILL",
        "location": "Tank Farm Sector 2",
        "details": "Solvent line containment breach leaking toxic vapors.",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/emergency/assess", json=assess_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["title"] == "Industrial Chemical Spill in Tank Farm 2"
    assert json_data["severity_level"] == "LEVEL_3_CRITICAL"
    assert json_data["status"] == "ACTIVE"
    emergency_id = json_data["id"]

    # 2. Plan Evacuation Routes
    evac_req = {"emergency_id": emergency_id}
    response = test_client.post("/emergency/evacuate", json=evac_req)
    assert response.status_code == 200
    json_data = response.json()
    assert "Emergency Stairwell East" in json_data["primary_route"] or "Corridor" in json_data["primary_route"]
    assert json_data["assembly_point"] is not None

    # 3. Broadcast Responder Alert Notifications
    notif_req = {
        "emergency_id": emergency_id,
        "message": "LEVEL 3 CRITICAL: Chemical Spill in Tank Farm 2. Evacuate immediately.",
        "channels": ["SMS", "PA_SYSTEM"],
    }
    response = test_client.post("/emergency/notify", json=notif_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["notifications_sent"] > 0
    assert json_data["status"] == "ALERT_BROADCAST_COMPLETE"

    # 4. Generate Situation Report (SITREP) Briefing
    response = test_client.post(f"/emergency/report?emergency_id={emergency_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert "# EMERGENCY SITUATION REPORT (SITREP)" in json_data["sitrep_markdown"]
    assert json_data["severity_level"] == "LEVEL_3_CRITICAL"

    # 5. Fetch Emergency Equipment Resources Inventory
    response = test_client.get("/emergency/resources")
    assert response.status_code == 200
    resources = response.json()["data"]
    assert len(resources) >= 3

    # 6. Fetch Emergency History Logs
    response = test_client.get("/emergency/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["id"] == emergency_id
