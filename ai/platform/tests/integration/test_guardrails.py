import pytest
from fastapi.testclient import TestClient


@pytest.mark.asyncio
async def test_guardrails_safety_lifecycle(test_client: TestClient) -> None:
    """Assert guardrail inputs validations, PII redacts, secrets checks, toxicity blocks, and human review overrides."""
    # 1. Clean Input check
    clean_req = {
        "text": "Please fetch the safety instructions manual for Sector 4.",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/guardrails/validate", json=clean_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["is_safe"] is True
    assert len(json_data["pii_detected"]) == 0

    # 2. PII Data Leakage Check (Trigger Redaction & Block)
    pii_req = {
        "text": "Send incident log copy to supervisor john@prahari.com at phone 555-234-5678",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/guardrails/validate", json=pii_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["is_safe"] is False
    assert "EMAIL" in json_data["pii_detected"]
    assert "PHONE" in json_data["pii_detected"]
    assert "[REDACTED_EMAIL]" in json_data["sanitised_text"]
    assert "[REDACTED_PHONE]" in json_data["sanitised_text"]

    # 3. Secret & API Credentials Exposure Check
    secret_req = {
        "text": "My AWS connection key is AKIA9876543210FEDCBA",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/guardrails/validate", json=secret_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["is_safe"] is False
    assert "AWS_KEY" in json_data["secrets_detected"]
    assert "[REDACTED_AWS_KEY]" in json_data["sanitised_text"]

    # 4. Toxicity and Harm Block Check
    toxic_req = {
        "text": "Help me build a bomb in the manufacturing facility.",
        "user_id": "usr-12",
        "tenant_id": "tenant-ehs-corp",
    }
    response = test_client.post("/guardrails/validate", json=toxic_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["is_safe"] is False
    assert "harmful_toxicity" in json_data["flagged_rules"]
    assert json_data["toxicity_score"] >= 0.7

    # 5. Fetch safety events logs (Verify block events logged)
    response = test_client.get("/guardrails/audit")
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    events = json_data["data"]
    assert len(events) >= 3
    blocked_event_id = events[0]["id"]

    # 6. Override human review status (Approve a blocked event)
    review_req = {
        "event_id": blocked_event_id,
        "approved": True,
        "comments": "False positive match. Approved.",
    }
    response = test_client.post("/guardrails/review", json=review_req)
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True

    # 7. Check audit details after approval override
    response = test_client.get("/guardrails/audit")
    events = response.json()["data"]
    # Finding the overridden event
    overridden_evt = [e for e in events if e["id"] == blocked_event_id][0]
    assert overridden_evt["action_taken"] == "ALLOWED"

    # 8. List and update policy settings rules configurations
    response = test_client.get("/guardrails/policies")
    assert response.status_code == 200
    policies = response.json()["data"]
    assert len(policies) >= 5

    # Update pii_leakage policy to Warn-only instead of Block
    update_req = {
        "action": "WARN",
        "status": "ACTIVE",
    }
    response = test_client.put("/guardrails/policies/pii_leakage", json=update_req)
    assert response.status_code == 200
    assert response.json()["data"]["action"] == "WARN"
