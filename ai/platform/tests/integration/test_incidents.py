import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_incident_investigation_workflow_lifecycle(test_client: TestClient) -> None:
    """Assert incident intake creation, RAG-guided search context, 5 Whys RCA engine, corrective planning, and final report."""
    # 1. Intake creation
    intake_req = {
        "title": "Acid leak from pipeline Sector 4",
        "description": "Safety valve failed, releasing 50L sulfuric acid onto Sector 4 concrete containment bund.",
        "classification": "Chemical Spill",
        "location": "Sector 4-B",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/incident/investigate", json=intake_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["title"] == "Acid leak from pipeline Sector 4"
    assert json_data["severity"] == "HIGH"
    assert json_data["status"] == "INVESTIGATION"
    investigation_id = json_data["id"]

    # 2. Run Root Cause Analysis (5 Whys)
    rca_req = {
        "investigation_id": investigation_id,
        "methodology": "5_WHYS",
    }
    response = test_client.post("/incident/root-cause", json=rca_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["methodology"] == "5_WHYS"
    assert len(json_data["root_causes"]) >= 1
    assert len(json_data["causative_chain"]) >= 5

    # 3. Recommend Corrective & Preventive actions
    rec_req = {
        "investigation_id": investigation_id,
    }
    response = test_client.post("/incident/recommend", json=rec_req)
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["corrective_actions"]) >= 1
    assert len(json_data["preventive_actions"]) >= 1

    # 4. Compile final PDF-ready Markdown report
    response = test_client.post("/incident/report", json=rec_req)
    assert response.status_code == 200
    json_data = response.json()
    assert "# INCIDENT INVESTIGATION REPORT" in json_data["report_markdown"]
    assert json_data["overall_confidence"] >= 0.9

    # 5. Fetch incident history logs
    response = test_client.get("/incident/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["id"] == investigation_id
