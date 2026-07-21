import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_executive_reporting_and_analytics_workflow(test_client: TestClient) -> None:
    """Assert CEO/Plant executive dashboards, board report generation, trend forecasting, enterprise KPIs, benchmarking, insights, and history."""
    # 1. Generate CEO Executive Dashboard
    dash_req = {
        "dashboard_type": "CEO",
        "timeframe_days": 30,
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/analytics/dashboard", json=dash_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["dashboard_type"] == "CEO"
    assert json_data["safety_score"] == 94.5
    assert "# CEO EXECUTIVE EHS DASHBOARD" in json_data["summary_markdown"]

    # 2. Compile Board Report
    report_req = {
        "report_type": "BOARD_REPORT",
        "title": "Q3 Enterprise EHS & AI Performance Review",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/analytics/report", json=report_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["title"] == "Q3 Enterprise EHS & AI Performance Review"
    assert "# Q3 Enterprise EHS & AI Performance Review" in json_data["report_markdown"]
    report_id = json_data["report_id"]

    # 3. Predict Trend Forecast
    forecast_req = {
        "metric_name": "Overall Safety Incident Risk",
        "forecast_months": 6,
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/analytics/forecast", json=forecast_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["metric_name"] == "Overall Safety Incident Risk"
    assert json_data["trend_direction"] == "IMPROVING"

    # 4. Fetch Enterprise KPI Metrics
    response = test_client.post("/analytics/kpi?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    json_data = response.json()["data"]
    assert json_data["trir"] > 0
    assert json_data["permit_compliance_pct"] > 90.0

    # 5. Benchmark Operating Units
    bench_req = {
        "target_unit": "Plant East (Houston)",
        "compare_unit": "Plant West (Dallas)",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/analytics/benchmark", json=bench_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["leader"] == "Plant East (Houston)"

    # 6. Retrieve AI-Generated Strategic Insights
    response = test_client.get("/analytics/insights?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    insights = response.json()["data"]
    assert len(insights) >= 2

    # 7. Fetch Executive Report History Logs
    response = test_client.get("/analytics/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["report_id"] == report_id
