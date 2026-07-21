import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_multi_agent_supervisor_workflow(test_client: TestClient) -> None:
    """Assert AI Operating System multi-agent task decomposition, subtask dispatch, consensus/conflict resolution, execution DAG graphs, status, and history."""
    # 1. Unified Supervisor Chat Entrypoint (multi-agent prompt dispatching to Risk & Asset agents)
    chat_req = {
        "user_query": "Assess operational risk and compressor asset health for boiler house bay.",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/supervisor/chat", json=chat_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["plan_status"] == "COMPLETED"
    assert len(json_data["execution_dag"]) >= 2
    assert "MULTI-AGENT CONSENSUS APPROVED" in json_data["consensus_summary"] or "CONFLICT RESOLVED" in json_data["consensus_summary"]
    session_id = json_data["session_id"]

    # 2. Decompose Complex Request into Task Execution Plan DAG
    plan_req = {"user_query": "Prepare full EHS inspection and risk report"}
    response = test_client.post("/supervisor/plan", json=plan_req)
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["planned_steps"]) >= 2

    # 3. Query Active Supervisor Session Status
    response = test_client.get(f"/supervisor/status?session_id={session_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["session_id"] == session_id
    assert json_data["completed_subtasks_count"] >= 1

    # 4. Fetch Execution DAG Graph Visualization (nodes & edges)
    response = test_client.get(f"/supervisor/graph?session_id={session_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["nodes"]) >= 3
    assert len(json_data["edges"]) >= 2

    # 5. Fetch Supervisor Session History Logs
    response = test_client.get("/supervisor/history?tenant_id=tenant-ehs-corp")
    assert response.status_code == 200
    history = response.json()["data"]
    assert len(history) >= 1
    assert history[0]["session_id"] == session_id
