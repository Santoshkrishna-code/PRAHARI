import pytest
from fastapi.testclient import TestClient
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import AsyncSessionLocal
from app.infrastructure.repositories.memory_repository import memory_repository
from app.core.memory_orchestrator import memory_orchestrator


@pytest.mark.asyncio
async def test_session_creation_and_history_caching(test_client: TestClient) -> None:
    """Assert session initialization, message appends, and cache retrieval works."""
    # 1. Create Session
    sess_payload = {"user_id": "usr-ehs-45", "tenant_id": "tenant-corp"}
    response = test_client.post("/memory/session", json=sess_payload)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    session_id = json_data["data"]["id"]

    # 2. Append User Message
    msg_payload = {
        "session_id": session_id,
        "role": "user",
        "content": "Gas leakage reported at Sector 4.",
    }
    response = test_client.post("/memory/message", json=msg_payload)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    assert json_data["data"]["content"] == "Gas leakage reported at Sector 4."

    # 3. Retrieve History Window
    response = test_client.get(f"/memory/history?session_id={session_id}")
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["data"]) == 1
    assert json_data["data"][0]["content"] == "Gas leakage reported at Sector 4."


@pytest.mark.asyncio
async def test_memory_scoring_and_retrieval(test_client: TestClient) -> None:
    """Assert declarative memories persistence, and score rankings are computed."""
    user_id = "usr-worker-101"
    async with AsyncSessionLocal() as session:
        # Create user profile preferences
        await memory_repository.save_preference(session, user_id, "lang", "en")
        
        # Seed two long term memories with different importance ranks
        m1 = await memory_repository.create_memory_entry(
            session, user_id, "preference", "User prefers metric temperature readings.", importance=5
        )
        m2 = await memory_repository.create_memory_entry(
            session, user_id, "incident", "Sector 3 experienced high pressure valve rupture last winter.", importance=9
        )
        assert m1.id is not None
        assert m2.id is not None

    # Retrieve relevant memories for incident search query
    response = test_client.get(f"/memory/relevant?user_id={user_id}&query=valve+leakage+pressure")
    assert response.status_code == 200
    json_data = response.json()
    
    # Assert return payload
    memories = json_data["memories"]
    assert len(memories) >= 2
    
    # The incident memory (importance 9) should rank higher than preference (importance 5)
    assert memories[0]["importance"] == 9
    assert "pressure" in memories[0]["content"]
    assert memories[0]["score"] > memories[1]["score"]


@pytest.mark.asyncio
async def test_purge_user_data(test_client: TestClient) -> None:
    """Assert compliance purge deletes user profiles, preferences, and sessions."""
    user_id = "usr-to-purge"
    
    async with AsyncSessionLocal() as session:
        # Seed preference data
        await memory_repository.save_preference(session, user_id, "mode", "dark")
        prefs = await memory_repository.get_preferences(session, user_id)
        assert prefs["mode"] == "dark"

    # Send DELETE request
    response = test_client.delete(f"/memory?user_id={user_id}")
    assert response.status_code == 200
    
    # Verify preferences database row is deleted
    async with AsyncSessionLocal() as session:
        prefs_after = await memory_repository.get_preferences(session, user_id)
        assert len(prefs_after) == 0
