import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_enterprise_risk_assessment_workflow(test_client: TestClient) -> None:
    """Assert ISO 31000 risk evaluations, predictive forecasting, mitigation hierarchy plans, risk registers, and trend metrics."""
    # 1. Risk Assessment submission (Chemical hazard HAZOP)
    assess_req = {
        "title": "Chemical vapor flange leak in Sector 2",
        "risk_type": "CHEMICAL",
        "methodology": "HAZOP",
        "details": "Potential benzene concentration above OSHA PEL limit.",
        "location": "Sector 2 Storage Tank",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/risk/assess", json=assess_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["title"] == "Chemical vapor flange leak in Sector 2"
    assert json_data["risk_type"] == "CHEMICAL"
    assert json_data["initial_risk_score"] == 20.0
    assert json_data["risk_category"] == "CRITICAL"
    assessment_id = json_data["id"]

    # 2. Predictive Risk forecasting
    predict_req = {
        "subject": "Sector 2 Chemical Plant operations",
        "historical_incidents_count": 4,
        "weather_condition": "STORM",
    }
    response = test_client.post("/risk/predict", json=predict_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["predicted_risk_score"] == 16.5
    assert json_data["trend_status"] == "INCREASING"

    # 3. Generate Control Mitigation Plan
    mitigate_req = {"assessment_id": assessment_id}
    response = test_client.post("/risk/mitigate", json=mitigate_req)
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["engineering_controls"]) >= 1
    assert len(json_data["administrative_controls"]) >= 1
    assert len(json_data["ppe_recommendations"]) >= 1
    assert json_data["residual_risk"] == 6.0

    # 4. Fetch 5x5 Risk Matrix configurations
    response = test_client.post("/risk/matrix")
    assert response.status_code == 200
    json_data = response.json()["data"]
    assert json_data["dimensions"] == "5x5"
    assert "CRITICAL" in json_data["categories"]

    # 5. Retrieve System-wide Risk Register
    response = test_client.get("/risk/register")
    assert response.status_code == 200
    register = response.json()["data"]
    assert len(register) >= 1
    assert register[0]["assessment_id"] == assessment_id

    # 6. Retrieve Enterprise Risk Trends
    response = test_client.get("/risk/trends")
    assert response.status_code == 200
    trends = response.json()["data"]
    assert trends["total_hazards_logged"] >= 1

    # 7. Fetch Risk History
    response = test_client.get("/risk/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["id"] == assessment_id
