import pytest
from fastapi.testclient import TestClient
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import AsyncSessionLocal
from app.infrastructure.repositories.knowledge_repository import knowledge_repository


@pytest.mark.asyncio
async def test_rag_ingestion_and_grounded_queries(test_client: TestClient) -> None:
    """Assert document ingestion, text segment chunking, permissions filters, and grounded citation answering."""
    tenant_id = "tenant-ehs-corp"
    doc_payload = {
        "filename": "warehouse_safety_sop.txt",
        "content_type": "text/plain",
        "file_size": 250,
        "tenant_id": tenant_id,
        "department": "Logistics",
        "project": "Safety-SOP",
    }

    # 1. Upload Document Metadata
    response = test_client.post(
        f"/documents/upload?filename={doc_payload['filename']}&content_type={doc_payload['content_type']}&file_size={doc_payload['file_size']}&tenant_id={doc_payload['tenant_id']}&department={doc_payload['department']}&project={doc_payload['project']}"
    )
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    doc_id = json_data["data"]["id"]
    assert json_data["data"]["status"] == "PENDING"

    # 2. Ingest and Chunk File Content
    # Raw document content containing explicit safety guidelines
    sop_content = (
        "Warehouse safety SOP guidelines require storing lithium battery cells in fireproof vaults. "
        "PPE standard EHS-30 requires using heat-resistant gloves when handling chemicals. "
        "Lockout-tagout procedure LOTO-04 requires cutting pressure lines before servicing main valve controls."
    )
    
    response = test_client.post(
        "/documents/index",
        json={"document_id": doc_id, "text_content": sop_content},
    )
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True

    # Confirm status updated to INDEXED in DB
    async with AsyncSessionLocal() as session:
        doc = await knowledge_repository.get_document(session, doc_id)
        assert doc is not None
        assert doc.status == "INDEXED"

    # 3. Test Search with matching terms
    search_payload = {
        "query": "lithium battery storage",
        "limit": 3,
        "department": "Logistics",
    }
    response = test_client.post(
        f"/search?tenant_id={tenant_id}",
        json=search_payload,
    )
    assert response.status_code == 200
    json_data = response.json()
    assert json_data["success"] is True
    chunks = json_data["data"]
    assert len(chunks) > 0
    assert "battery" in chunks[0]["content"]

    # 4. Test Permissions isolation filter
    # Retrieve using a mismatching department: should return empty chunks list
    search_payload["department"] = "Finance"
    response = test_client.post(
        f"/search?tenant_id={tenant_id}",
        json=search_payload,
    )
    assert response.status_code == 200
    json_data = response.json()
    assert len(json_data["data"]) == 0

    # 5. Execute Grounded Query answering with Citations
    query_payload = {
        "query": "Where do we store lithium battery cells?",
        "limit": 3,
        "department": "Logistics",
    }
    response = test_client.post(
        f"/query?tenant_id={tenant_id}",
        json=query_payload,
    )
    assert response.status_code == 200
    json_data = response.json()
    
    # Grounded answer must contain text response and citations mapping
    assert json_data["answer"] is not None
    assert "warehouse_safety_sop.txt" in json_data["citations"][0]["document_name"]
    assert json_data["citations"][0]["page_number"] is not None
