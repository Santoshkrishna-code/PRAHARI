import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_permit_to_work_workflow_lifecycle(test_client: TestClient) -> None:
    """Assert permit creation, SOP validation checks, 5x5 risk matrix scoring, approval recommendations, formal sign-offs, and templates."""
    # 1. Permit Draft creation
    create_req = {
        "permit_title": "Hot work welding on pipeline flange",
        "permit_type": "HOT_WORK",
        "location": "Sector 4-B",
        "applicant_id": "emp-104",
        "details": "Welding containment flange joint in Sector 4-B.",
        "hazard_controls": ["Fire Extinguisher on site"],
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/permit/create", json=create_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["permit_title"] == "Hot work welding on pipeline flange"
    assert json_data["status"] == "DRAFT"
    permit_id = json_data["id"]

    # 2. Validate permit completeness against SOPs
    response = test_client.post(f"/permit/validate?permit_id={permit_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["is_valid"] is False
    assert len(json_data["missing_controls"]) >= 1
    assert json_data["risk_level"] == "HIGH"

    # 3. Assess 5x5 Risk Matrix
    risk_req = {"permit_id": permit_id}
    response = test_client.post("/permit/risk", json=risk_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["residual_risk_score"] == 12.0
    assert json_data["risk_category"] == "MEDIUM"

    # 4. Evaluate Approval recommendation
    rec_req = {"permit_id": permit_id}
    response = test_client.post("/permit/recommend", json=rec_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["recommendation"] in ["APPROVE", "SUSPEND"]

    # 5. Formal sign-off approval
    response = test_client.post(f"/permit/approve?permit_id={permit_id}&approver_id=officer-01&role=SAFETY_OFFICER&approved=true")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["status"] == "APPROVED"

    # 6. Fetch supported permit templates
    response = test_client.get("/permit/templates")
    assert response.status_code == 200
    templates = response.json()["data"]
    assert "HOT_WORK" in templates
    assert "CONFINED_SPACE" in templates

    # 7. Fetch permit history logs
    response = test_client.get("/permit/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["id"] == permit_id
