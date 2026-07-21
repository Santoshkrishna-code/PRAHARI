import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_intelligent_inspection_workflow(test_client: TestClient) -> None:
    """Assert inspection planning, dynamic risk-based checklists, finding classification, Risk Agent integration, reports, and templates."""
    # 1. Plan workplace inspection
    plan_req = {
        "title": "Weekly Chemical Storage Area Safety Audit",
        "inspection_type": "CHEMICAL_STORAGE",
        "location": "Chemical Store Unit 3",
        "inspector_id": "insp-88",
        "scope_notes": "Routine weekly environmental containment audit",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/inspection/plan", json=plan_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["title"] == "Weekly Chemical Storage Area Safety Audit"
    assert json_data["status"] == "PLANNED"
    assert json_data["score"] == 100.0
    inspection_id = json_data["id"]

    # 2. Generate dynamic checklist
    chk_req = {
        "inspection_type": "CHEMICAL_STORAGE",
        "location": "Chemical Store Unit 3",
    }
    response = test_client.post("/inspection/checklist", json=chk_req)
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["items"]) >= 3

    # 3. Analyze finding (triggers Risk Agent call on major non-conformance)
    find_req = {
        "inspection_id": inspection_id,
        "observation_text": "Secondary containment bunding drain valve left open with minor chemical leak.",
    }
    response = test_client.post("/inspection/analyze", json=find_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["finding_type"] == "MAJOR_NON_CONFORMANCE"
    assert json_data["severity"] == "HIGH"
    assert json_data["risk_priority_score"] > 0

    # 4. Compile final PDF-ready Markdown EHS Inspection Report
    response = test_client.post(f"/inspection/report?inspection_id={inspection_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert "# EHS INTELLIGENT INSPECTION REPORT" in json_data["report_markdown"]
    assert json_data["score"] == 85.0

    # 5. Fetch supported inspection templates
    response = test_client.get("/inspection/templates")
    assert response.status_code == 200
    templates = response.json()["data"]
    assert "CHEMICAL_STORAGE" in templates
    assert "FIRE_SAFETY" in templates

    # 6. Fetch inspection history logs
    response = test_client.get("/inspection/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["id"] == inspection_id
