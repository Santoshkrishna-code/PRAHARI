import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_compliance_workflow_lifecycle(test_client: TestClient) -> None:
    """Assert compliance check, regulatory violations logging, CAPA recommendations, reports, and site compliance scores."""
    # 1. Compliance check with a clean compliant subject
    clean_req = {
        "target_subject": "Harness inspection Sector 2",
        "subject_type": "INSPECTION",
        "details": "Harness inspection completed successfully. Log shows zero wears or tears.",
        "user_id": "usr-12",
        "tenant_id": "tenant-compliance-test",
    }
    response = test_client.post("/compliance/check", json=clean_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "COMPLIANT"
    assert json_data["compliance_score"] == 100.0

    # 2. Compliance check triggering a high severity violation
    violation_req = {
        "target_subject": "PPE Checklist Sector 4 flange work",
        "subject_type": "PPE",
        "details": "Missing standard eye protection shield and mandatory heat glove cert.",
        "user_id": "usr-12",
        "tenant_id": "tenant-compliance-test",
    }
    response = test_client.post("/compliance/check", json=violation_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "NON_COMPLIANT"
    assert json_data["compliance_score"] == 55.0
    assessment_id = json_data["id"]

    # 3. Generate CAPA recommendations
    capa_req = {
        "assessment_id": assessment_id,
    }
    response = test_client.post("/compliance/capa", json=capa_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["compliance_risk"] == "HIGH"
    assert len(json_data["corrective_actions"]) >= 1
    assert len(json_data["preventive_actions"]) >= 1

    # 4. Compile final PDF-ready Markdown report
    response = test_client.post("/compliance/report", json=capa_req)
    assert response.status_code == 200
    json_data = response.json()
    assert "# SAFETY COMPLIANCE REPORT" in json_data["report_markdown"]
    assert json_data["score"] == 55.0

    # 5. Fetch site overall compliance score
    response = test_client.get("/compliance/score?tenant_id=tenant-compliance-test")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    # Average of 100.0 and 55.0 = 77.5
    assert json_data["data"]["average_score"] == 77.5

    # 6. Fetch compliance check history logs
    response = test_client.get("/compliance/history?tenant_id=tenant-compliance-test")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 2
