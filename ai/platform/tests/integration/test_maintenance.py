import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_predictive_maintenance_workflow(test_client: TestClient) -> None:
    """Assert failure probability scoring, Remaining Useful Life (RUL) estimation, condition-based work order creation, reliability dashboard, and history logs."""
    # 1. Register asset health to get valid asset ID (Volume 15)
    asset_req = {
        "asset_name": "Main Gas Compressor Unit 1",
        "asset_category": "COMPRESSOR",
        "location": "Compressor Station 3",
        "operating_hours": 2100.0,
        "vibration_mm_s": 4.2,
        "temperature_c": 82.0,
        "pressure_bar": 8.5,
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/asset/health", json=asset_req)
    assert response.status_code == 200
    asset_id = response.json()["id"]

    # 2. Predict Equipment Failure
    predict_req = {
        "asset_id": asset_id,
        "operating_hours": 2100.0,
        "vibration_trend_mm_s": 4.2,
        "temperature_trend_c": 82.0,
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/maintenance/predict", json=predict_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["asset_id"] == asset_id
    assert json_data["failure_probability"] > 0.4
    assert json_data["status"] == "RECOMMENDED"

    # 3. Estimate Remaining Useful Life (RUL)
    rul_req = {"asset_id": asset_id}
    response = test_client.post("/maintenance/rul", json=rul_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["remaining_useful_life_hours"] > 0
    assert json_data["confidence_score"] == 0.91

    # 4. Create Condition-Based Work Order Plan
    plan_req = {
        "asset_id": asset_id,
        "work_order_title": "Vibration damper replacement and main shaft overhaul",
    }
    response = test_client.post("/maintenance/plan", json=plan_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["priority"] in ["HIGH", "CRITICAL"]
    assert len(json_data["required_spare_parts"]) >= 1

    # 5. Generate Maintenance Executive Report
    response = test_client.post(f"/maintenance/report?asset_id={asset_id}")
    assert response.status_code == 200
    json_data = response.json()["data"]
    assert json_data["asset_id"] == asset_id

    # 6. Retrieve Plant-Wide Reliability Dashboard
    response = test_client.get("/maintenance/dashboard?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    metrics = response.json()["data"]
    assert metrics["total_assets_monitored"] >= 1

    # 7. Fetch Maintenance Prediction History Logs
    response = test_client.get("/maintenance/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["asset_id"] == asset_id
