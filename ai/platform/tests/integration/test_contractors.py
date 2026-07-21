import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_contractor_safety_workflow_lifecycle(test_client: TestClient) -> None:
    """Assert contractor registration, safety pre-qualification checks, document verification, Risk Agent integration, scorecards, and history."""
    # 1. Register contractor profile
    reg_req = {
        "company_name": "Apex Scaffolding & Erectors Ltd",
        "contractor_type": "SCAFFOLDING",
        "contact_email": "safety@apexscaffolding.com",
        "license_number": "LIC-SCAF-9921",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/contractor/register", json=reg_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["company_name"] == "Apex Scaffolding & Erectors Ltd"
    assert json_data["status"] == "REGISTERED"
    assert json_data["safety_rating_score"] == 100.0
    contractor_id = json_data["id"]

    # 2. Pre-qualification evaluation (invokes Safety Compliance and Risk agents)
    qual_req = {"contractor_id": contractor_id}
    response = test_client.post("/contractor/qualify", json=qual_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["is_qualified"] is True
    assert json_data["status"] == "QUALIFIED"
    assert len(json_data["missing_trainings"]) >= 1

    # 3. Document verification
    ver_req = {
        "contractor_id": contractor_id,
        "document_type": "INSURANCE",
    }
    response = test_client.post("/contractor/verify", json=ver_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["is_valid"] is True

    # 4. Assess contractor task risk (calls Enterprise Risk Agent)
    response = test_client.post(f"/contractor/risk?contractor_id={contractor_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert "Contractor Task Risk" in json_data["title"]

    # 5. Generate contractor report
    response = test_client.post(f"/contractor/report?contractor_id={contractor_id}")
    assert response.status_code == 200
    json_data = response.json()["data"]
    assert json_data["contractor_id"] == contractor_id

    # 6. Retrieve contractor scorecard
    response = test_client.get(f"/contractor/scorecard?contractor_id={contractor_id}")
    assert response.status_code == 200
    json_data = response.json()["data"]
    assert json_data["safety_score"] == 100.0
    assert json_data["compliance_status"] == "COMPLIANT"

    # 7. Fetch contractor history logs
    response = test_client.get("/contractor/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["id"] == contractor_id
