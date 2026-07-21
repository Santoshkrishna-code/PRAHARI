import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_evaluation_and_observability_lifecycle(test_client: TestClient) -> None:
    """Assert quality scoring evaluations, comparative benchmarks, leaderboard rankings, and dashboard reports."""
    # 1. Run quality evaluation (matches seeder golden dataset)
    eval_req = {
        "text_input": "What are the core EHS guidelines for hazardous gas storage?",
        "text_output": "Install gas leak sensors, maintain 15-meter buffer zone, and ensure employee ventilation suits.",
        "retrieved_context": [
            "Install gas leak sensors and maintain 15-meter buffer zone inside plant Sector 4."
        ],
        "cost": 0.0035,
        "latency_ms": 450,
        "model": "gpt-4o",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/evaluate", json=eval_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["correctness_score"] == 1.0
    assert json_data["faithfulness_score"] >= 0.5
    assert json_data["overall_score"] >= 0.7
    evaluation_id = json_data["evaluation_id"]

    # 2. Register human override rating feedback
    feedback_req = {
        "evaluation_id": evaluation_id,
        "rating": 5,
        "comments": "Excellent grounding citation alignment accuracy.",
    }
    response = test_client.post("/feedback", json=feedback_req)
    assert response.status_code == 200
    assert response.json()["success"] is True

    # 3. Create A/B Testing Variant configurations
    exp_req = {
        "name": "RAG Chunk Size Experiment",
        "description": "Variant A checks 512 chunks, Variant B checks 1024 chunks.",
        "variant_a_config": {"chunk_size": 512},
        "variant_b_config": {"chunk_size": 1024},
    }
    response = test_client.post("/experiment", json=exp_req)
    assert response.status_code == 200
    assert response.json()["success"] is True

    # 4. Trigger model benchmarking metrics
    bench_req = {
        "name": "EHS Incidents Benchmark",
        "model_name": "claude-3-opus",
        "test_prompts": [
            "Describe safety guidelines.",
            "Explain permit sign-offs."
        ]
    }
    response = test_client.post("/benchmark", json=bench_req)
    assert response.status_code == 200
    assert response.json()["model_name"] == "claude-3-opus"

    # 5. Fetch model leaderboard
    response = test_client.get("/leaderboard")
    assert response.status_code == 200
    leaderboard = response.json()["data"]
    assert len(leaderboard) >= 1
    assert leaderboard[0]["model_name"] == "gpt-4o"

    # 6. Fetch latency, quality, cost dashboards analytics
    response = test_client.get("/quality")
    assert response.status_code == 200
    assert response.json()["data"]["faithfulness"] >= 0.5

    response = test_client.get("/cost")
    assert response.status_code == 200
    assert response.json()["data"]["total_accumulated_cost"] >= 0.0035

    response = test_client.get("/latency")
    assert response.status_code == 200
    assert response.json()["data"]["average_latency_ms"] == 450

    response = test_client.get("/reports")
    assert response.status_code == 200
    report = response.json()["data"]
    assert "quality" in report
    assert "billing" in report
    assert "latency" in report
