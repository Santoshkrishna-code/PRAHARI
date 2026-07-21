import difflib
import datetime
from typing import Any
from fastapi import APIRouter, Response, HTTPException, Depends, Query
from fastapi.responses import StreamingResponse
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload

# Core Models
from app.core.models import (
    HealthStatus,
    StandardResponse,
    ChatRequest,
    ChatResponse,
    Choice,
    Message,
    EmbeddingRequest,
    EmbeddingResponse,
    CompletionRequest,
    CompletionResponse,
)
from app.core.model_registry import MODEL_REGISTRY, ModelMetadata

# Prompt Domain Models
from app.domain.prompts import (
    PromptDto,
    PromptVersionDto,
    RenderRequest,
    RenderResponse,
    ValidateRequest,
    ValidateResponse,
    CompareRequest,
    CompareResponse,
    PromptTestRequest,
    PromptTestResponse,
)

# Memory Domain Models
from app.domain.memory import (
    SessionDto,
    MessageDto,
    MemoryEntryDto,
    SessionCreateRequest,
    MessageAddRequest,
    MemorySearchRequest,
    RelevantMemoryResponse,
)

# RAG Domain Models
from app.domain.knowledge import (
    DocumentDto,
    ChunkDto,
    CitationDto,
    SearchRequest,
    RetrieveResponse,
    QueryRequest,
    QueryResponse,
    IndexRequest,
)

# Tools & MCP Domain Models
from app.domain.tools import (
    ToolDto,
    ToolVersionDto,
    ToolRegisterRequest,
    ToolExecuteRequest,
    ToolExecuteResponse,
    McpInitializeRequest,
    McpInitializeResponse,
    McpToolDto,
    McpToolsListResponse,
)

# Multi-Agent Workflow Domain Models
from app.domain.agents import (
    AgentGraphDto,
    AgentNodeDto,
    AgentEdgeDto,
    AgentRunRequest,
    AgentRunResponse,
    AgentApprovalRequest,
)

from app.domain.guardrails import (
    PolicyDto,
    PolicyCreateRequest,
    PolicyUpdateRequest,
    GuardrailCheckRequest,
    GuardrailCheckResponse,
    HumanReviewRequest,
    SafetyEventDto,
)

# Safety Compliance Agent Domain Models
from app.domain.compliance import (
    ComplianceCheckRequest,
    ComplianceAssessmentDto,
    CapaPlanRequest,
    CapaPlanResponse,
    ComplianceReportResponse,
)

# Incident Agent Domain Models
from app.domain.incidents import (
    IncidentIntakeRequest,
    IncidentInvestigationDto,
    RootCauseRequest,
    RootCauseResponse,
    RecommendationRequest,
    RecommendationResponse,
    ReportResponse,
)

# Permit-to-Work Agent Domain Models
from app.domain.permits import (
    PermitCreateRequest,
    PermitRecordDto,
    PermitValidationResponse,
    RiskAssessmentRequest,
    RiskAssessmentResponse,
    ApprovalRecommendationRequest,
    ApprovalRecommendationResponse,
)

# Enterprise Risk Agent Domain Models
from app.domain.risks import (
    RiskAssessmentRequest as EnterpriseRiskAssessmentRequest,
    RiskAssessmentDto,
    RiskPredictRequest,
    RiskPredictResponse,
    MitigationPlanRequest,
    MitigationPlanResponse,
    RiskRegisterEntryDto,
)

# Intelligent Inspection Agent Domain Models
from app.domain.inspections import (
    InspectionPlanRequest,
    InspectionRecordDto,
    ChecklistGenerationRequest,
    ChecklistGenerationResponse,
    FindingAnalysisRequest,
    FindingAnalysisResponse,
    InspectionReportResponse,
)

# Contractor Safety Agent Domain Models
from app.domain.contractors import (
    ContractorRegisterRequest,
    ContractorRecordDto,
    ContractorQualifyRequest,
    ContractorQualifyResponse,
    ContractorVerifyRequest,
    ContractorVerifyResponse,
    ContractorScorecardResponse,
)

# Asset & Equipment Safety Agent Domain Models
from app.domain.assets import (
    AssetHealthCheckRequest,
    AssetRecordDto,
    AssetShutdownRequest,
    AssetShutdownResponse,
    AssetScorecardResponse,
)

# Emergency Response Agent Domain Models
from app.domain.emergencies import (
    EmergencyAssessRequest,
    EmergencyRecordDto,
    EvacuationPlanRequest,
    EvacuationPlanResponse,
    EmergencyNotifyRequest,
    EmergencyNotifyResponse,
    SitrepResponse,
    EmergencyResourceDto,
)

# Predictive Maintenance Agent Domain Models
from app.domain.maintenance import (
    MaintenancePredictRequest,
    MaintenancePredictionDto,
    RulEstimateRequest,
    RulEstimateResponse,
    MaintenancePlanRequest,
    MaintenancePlanResponse,
    ReliabilityDashboardResponse,
)

# Executive Reporting & Analytics Agent Domain Models
from app.domain.analytics import (
    ExecutiveDashboardRequest,
    ExecutiveDashboardResponse,
    ReportGenerateRequest,
    ReportGenerateResponse,
    ForecastRequest,
    ForecastResponse,
    KpiMetricsResponse,
    UnitBenchmarkRequest,
    UnitBenchmarkResponse,
)

# Multi-Agent Supervisor & Collaboration Agent Domain Models
from app.domain.supervisor import (
    SupervisorChatRequest,
    SupervisorChatResponse,
    SupervisorPlanRequest,
    SupervisorPlanResponse,
    SupervisorExecuteRequest,
    SupervisorStatusResponse,
    SupervisorDagResponse,
)

# Infrastructure
from app.infrastructure.database import check_db_health, get_db_session
from app.infrastructure.cache import cache_provider
from app.infrastructure.telemetry import get_metrics_payload
from app.infrastructure.providers.factory import provider_factory
from app.infrastructure.repositories.prompt_repository import prompt_repository
from app.infrastructure.repositories.memory_repository import memory_repository
from app.infrastructure.repositories.knowledge_repository import knowledge_repository
from app.infrastructure.repositories.tool_repository import tool_repository
from app.infrastructure.repositories.agent_repository import agent_repository
from app.infrastructure.repositories.evaluation_repository import evaluation_repository
from app.infrastructure.repositories.guardrail_repository import guardrail_repository
from app.infrastructure.repositories.incident_repository import incident_repository
from app.core.incident_agent import incident_agent
from app.infrastructure.repositories.compliance_repository import compliance_repository
from app.core.compliance_agent import safety_compliance_agent
from app.infrastructure.repositories.permit_repository import permit_repository
from app.core.permit_agent import permit_agent
from app.infrastructure.repositories.risk_repository import risk_repository
from app.core.risk_agent import enterprise_risk_agent
from app.infrastructure.repositories.inspection_repository import inspection_repository
from app.core.inspection_agent import intelligent_inspection_agent
from app.infrastructure.repositories.contractor_repository import contractor_repository
from app.core.contractor_agent import contractor_safety_agent
from app.infrastructure.repositories.asset_repository import asset_repository
from app.core.asset_agent import asset_equipment_safety_agent
from app.infrastructure.repositories.emergency_repository import emergency_repository
from app.core.emergency_agent import emergency_response_agent
from app.infrastructure.repositories.maintenance_repository import maintenance_repository
from app.core.maintenance_agent import predictive_maintenance_agent
from app.infrastructure.repositories.analytics_repository import analytics_repository
from app.core.analytics_agent import executive_reporting_agent
from app.infrastructure.repositories.supervisor_repository import supervisor_repository
from app.core.supervisor_agent import multi_agent_supervisor
from app.infrastructure.models.tool_entities import ToolDb
from app.core.prompt_renderer import prompt_renderer
from app.core.prompt_validator import prompt_validator
from app.core.memory_orchestrator import memory_orchestrator
from app.core.rag_orchestrator import rag_orchestrator
from app.core.tool_executor import tool_executor
from app.core.agent_orchestrator import agent_orchestrator
from app.core.guardrail_engine import guardrail_engine
from app.core.evaluation_engine import evaluation_engine

router = APIRouter()


@router.get(
    "/health",
    response_model=HealthStatus,
    summary="Retrieve platform components health status",
)
async def health_check() -> HealthStatus:
    """Executes ping commands across core DB and Cache backends."""
    db_alive = await check_db_health()
    redis_alive = await cache_provider.check_health()

    status = "OK" if db_alive and redis_alive else "DEGRADED"
    return HealthStatus(
        status=status,
        database="HEALTHY" if db_alive else "UNHEALTHY",
        cache="HEALTHY" if redis_alive else "UNHEALTHY",
    )


@router.get(
    "/metrics",
    summary="Retrieve scraped Prometheus metrics metrics",
)
async def metrics_endpoint() -> Response:
    """Returns the formatted Prometheus registry metrics payload."""
    payload = get_metrics_payload()
    return Response(content=payload, media_type="text/plain; version=0.0.4")


@router.get(
    "/api/v1/ping",
    response_model=StandardResponse[str],
    summary="Simple API Ping check",
)
async def api_ping() -> StandardResponse[str]:
    """Base validation endpoint checking auth integration."""
    return StandardResponse(data="pong", message="AI Platform base responds successfully")


# LLM Execution Routers
@router.post(
    "/chat",
    response_model=ChatResponse,
    summary="Generate chat response across models",
)
async def chat_endpoint(request: ChatRequest) -> ChatResponse:
    """Routes query request to resolved provider."""
    return await provider_factory.execute_chat(request)


@router.post(
    "/complete",
    response_model=CompletionResponse,
    summary="Generate single-turn text completion",
)
async def complete_endpoint(request: CompletionRequest) -> CompletionResponse:
    """Converts prompt to chat payload and executes."""
    model = request.model or "gpt-4o"
    provider = provider_factory.resolve_provider_for_model(model)
    return await provider.generate_completion(request)


@router.post(
    "/stream",
    summary="Generate streaming chat response chunks",
)
async def stream_endpoint(request: ChatRequest) -> StreamingResponse:
    """Establishes Server-Sent Events (SSE) connection yielding JSON response blocks."""
    model = request.model or "gpt-4o"
    provider = provider_factory.resolve_provider_for_model(model)

    async def event_generator():
        async for chunk in provider.generate_stream(request):
            yield f"data: {chunk.model_dump_json()}\n\n"

    return StreamingResponse(event_generator(), media_type="text/event-stream")


@router.post(
    "/embedding",
    response_model=EmbeddingResponse,
    summary="Request text embedding vectors",
)
async def embedding_endpoint(request: EmbeddingRequest) -> EmbeddingResponse:
    """Calculates embedding coordinates for inputs."""
    return await provider_factory.execute_embeddings(request)


@router.post(
    "/vision",
    response_model=ChatResponse,
    summary="Vision analysis abstraction endpoint",
)
async def vision_endpoint(request: ChatRequest) -> ChatResponse:
    """Stub pipeline abstraction routing visual payload files checks."""
    model = request.model or "gpt-4o"
    metadata = MODEL_REGISTRY.get(model)
    if not metadata or not metadata.supports_vision:
        raise HTTPException(
            status_code=400,
            detail=f"Requested model '{model}' does not support vision processing.",
        )
    return await provider_factory.execute_chat(request)


# Prompt Registry CRUD Routers
@router.post(
    "/prompts",
    response_model=StandardResponse[PromptDto],
    summary="Register a new prompt definition",
)
async def create_prompt(
    name: str,
    description: str | None = None,
    category_id: str | None = None,
    user_id: str = "sys-admin",
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[PromptDto]:
    try:
        dto = await prompt_repository.create_prompt(db, name, description, category_id, user_id)
        return StandardResponse(data=dto, message="Prompt master registered successfully.")
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@router.get(
    "/prompts",
    response_model=StandardResponse[list[PromptDto]],
    summary="List all registered prompt definitions",
)
async def list_prompts(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[list[PromptDto]]:
    dto_list = await prompt_repository.get_prompts_list(db)
    return StandardResponse(data=dto_list)


@router.post(
    "/prompts/{name}/versions",
    response_model=StandardResponse[PromptVersionDto],
    summary="Register a new semantic version draft for a prompt",
)
async def create_prompt_version(
    name: str,
    version_str: str,
    system_template: str,
    user_template: str,
    few_shots: list[dict[str, Any]] | None = None,
    response_format: dict[str, Any] | None = None,
    user_id: str = "sys-admin",
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[PromptVersionDto]:
    try:
        dto = await prompt_repository.add_version(
            db, name, version_str, system_template, user_template, user_id, few_shots, response_format
        )
        return StandardResponse(data=dto, message="New draft version registered.")
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@router.post(
    "/prompts/versions/{version_id}/approve",
    response_model=StandardResponse[str],
    summary="Approve draft and activate version as production current",
)
async def approve_prompt_version(
    version_id: str,
    reviewer_id: str = "reviewer-01",
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    try:
        await prompt_repository.approve_version(db, version_id, reviewer_id)
        return StandardResponse(data=version_id, message="Prompt version approved and activated.")
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


# Prompt Operations & Rendering
@router.post(
    "/render",
    response_model=RenderResponse,
    summary="Render templates using active prompt configurations and variables",
)
async def render_prompt(
    request: RenderRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RenderResponse:
    if request.version:
        version = await prompt_repository.get_version_by_number(db, request.prompt_name, request.version)
    else:
        version = await prompt_repository.get_active_version(db, request.prompt_name)

    if not version:
        raise HTTPException(
            status_code=404,
            detail=f"No approved prompt version found for: {request.prompt_name}",
        )

    sys_text, usr_text, resolved, missing = prompt_renderer.resolve_variables(
        version.system_template, version.user_template, request.variables
    )

    tokens, cost = prompt_validator.estimate_cost(sys_text, usr_text, "gpt-4o")

    return RenderResponse(
        system_prompt=sys_text,
        user_prompt=usr_text,
        variables_resolved=resolved,
        variables_missing=missing,
        estimated_tokens=tokens,
        estimated_cost=cost,
    )


@router.post(
    "/validate",
    response_model=ValidateResponse,
    summary="Scan prompt templates for syntactical formatting and injection flags",
)
async def validate_prompt(request: ValidateRequest) -> ValidateResponse:
    valid, errors, warnings, injection, unsafe = prompt_validator.validate_templates(
        request.system_template, request.user_template, request.variables
    )
    return ValidateResponse(
        is_valid=valid,
        errors=errors,
        warnings=warnings,
        prompt_injection_detected=injection,
        unsafe_instructions_detected=unsafe,
    )


@router.post(
    "/compare",
    response_model=CompareResponse,
    summary="Compare and compute diff between system/user prompts for two versions",
)
async def compare_prompt(
    request: CompareRequest,
    db: AsyncSession = Depends(get_db_session),
) -> CompareResponse:
    ver_a = await prompt_repository.get_version_by_number(db, request.prompt_name, request.version_a)
    ver_b = await prompt_repository.get_version_by_number(db, request.prompt_name, request.version_b)

    if not ver_a or not ver_b:
        raise HTTPException(status_code=404, detail="One or both versions not found.")

    diff_sys = "\n".join(
        difflib.unified_diff(
            ver_a.system_template.splitlines(),
            ver_b.system_template.splitlines(),
            fromfile=request.version_a,
            tofile=request.version_b,
        )
    )

    diff_usr = "\n".join(
        difflib.unified_diff(
            ver_a.user_template.splitlines(),
            ver_b.user_template.splitlines(),
            fromfile=request.version_a,
            tofile=request.version_b,
        )
    )

    return CompareResponse(
        prompt_name=request.prompt_name,
        version_a=request.version_a,
        version_b=request.version_b,
        diff_system=diff_sys,
        diff_user=diff_usr,
    )


@router.post(
    "/test",
    response_model=PromptTestResponse,
    summary="Execute test run against rendered templates validating expected substrings",
)
async def test_prompt(
    request: PromptTestRequest,
    db: AsyncSession = Depends(get_db_session),
) -> PromptTestResponse:
    version = await prompt_repository.get_version_by_number(db, request.prompt_name, request.version)
    if not version:
        raise HTTPException(status_code=404, detail="Target prompt version not found.")

    sys_text, usr_text, _, _ = prompt_renderer.resolve_variables(
        version.system_template, version.user_template, request.variables
    )

    combined_text = sys_text + "\n" + usr_text
    matched = []
    failed = []

    for substr in request.expected_substrings:
        if substr.lower() in combined_text.lower():
            matched.append(substr)
        else:
            failed.append(substr)

    passed = len(failed) == 0

    return PromptTestResponse(
        passed=passed,
        rendered_system=sys_text,
        rendered_user=usr_text,
        matched_substrings=matched,
        failed_substrings=failed,
    )


# Memory Engine Routers
@router.post(
    "/memory/session",
    response_model=StandardResponse[SessionDto],
    summary="Create a new conversation session",
)
async def create_session(
    request: SessionCreateRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[SessionDto]:
    dto = await memory_repository.create_session(db, request.user_id, request.tenant_id)
    return StandardResponse(data=dto, message="Session created.")


@router.post(
    "/memory/message",
    response_model=StandardResponse[MessageDto],
    summary="Append a new message to conversation history",
)
async def add_message(
    request: MessageAddRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[MessageDto]:
    dto = await memory_orchestrator.add_message_and_sync_cache(
        db, request.session_id, request.role, request.content
    )
    return StandardResponse(data=dto, message="Message appended successfully.")


@router.get(
    "/memory/session/{session_id}",
    response_model=StandardResponse[SessionDto],
    summary="Retrieve session metadata parameters",
)
async def get_session_info(
    session_id: str,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[SessionDto]:
    dto = await memory_repository.get_session(db, session_id)
    if not dto:
        raise HTTPException(status_code=404, detail="Session not found.")
    return StandardResponse(data=dto)


@router.get(
    "/memory/history",
    response_model=StandardResponse[list[MessageDto]],
    summary="Retrieve conversation history window for a session",
)
async def get_session_history(
    session_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[MessageDto]]:
    msgs = await memory_orchestrator.get_session_messages(db, session_id)
    return StandardResponse(data=msgs)


@router.get(
    "/memory/relevant",
    response_model=RelevantMemoryResponse,
    summary="Retrieve relevant scored long term memories and active summary",
)
async def get_relevant_memories(
    user_id: str = Query(...),
    query: str = Query(...),
    session_id: str | None = Query(None),
    limit: int = Query(5, ge=1, le=20),
    db: AsyncSession = Depends(get_db_session),
) -> RelevantMemoryResponse:
    memories = await memory_orchestrator.retrieve_relevant_memories(db, user_id, query, limit)
    
    summary = None
    if session_id:
        summary = await memory_repository.get_summary(db, session_id)

    return RelevantMemoryResponse(memories=memories, summary=summary)


@router.post(
    "/memory/summarize",
    response_model=StandardResponse[str],
    summary="Manually compile running summary of a session history",
)
async def summarize_session(
    session_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    summary = await memory_orchestrator.compile_running_summary(db, session_id)
    return StandardResponse(data=summary, message="Summary updated successfully.")


@router.post(
    "/memory/search",
    response_model=StandardResponse[list[MemoryEntryDto]],
    summary="Search relevant scored declarative memories",
)
async def search_memories(
    request: MemorySearchRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[MemoryEntryDto]]:
    memories = await memory_orchestrator.retrieve_relevant_memories(
        db, request.user_id, request.query, request.limit
    )
    return StandardResponse(data=memories)


@router.delete(
    "/memory",
    response_model=StandardResponse[str],
    summary="Purge all session histories and profile parameters for a user ID",
)
async def purge_user_memory(
    user_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    await memory_repository.purge_user_memory(db, user_id)
    return StandardResponse(data=user_id, message="User memory purged for compliance.")


# RAG Knowledge Platform Routers
@router.post(
    "/documents/upload",
    response_model=StandardResponse[DocumentDto],
    summary="Upload and register document metadata",
)
async def upload_document(
    filename: str,
    content_type: str,
    file_size: int,
    tenant_id: str,
    department: str | None = None,
    project: str | None = None,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[DocumentDto]:
    dto = await knowledge_repository.create_document(
        db, filename, content_type, file_size, tenant_id, department, project
    )
    return StandardResponse(data=dto, message="Document registered.")


@router.post(
    "/documents/index",
    response_model=StandardResponse[str],
    summary="Parses, chunks, and indexes registered document",
)
async def index_document(
    request: IndexRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    doc = await knowledge_repository.get_document(db, request.document_id)
    if not doc:
        raise HTTPException(status_code=404, detail="Document not found.")

    # Ingest chunks and embeddings synchronously in request lifecycle
    await rag_orchestrator.ingest_and_index_document(
        db, doc.id, doc.filename, request.text_content, doc.tenant_id, doc.department, doc.project
    )
    return StandardResponse(data=request.document_id, message="Document chunks indexed successfully.")


@router.post(
    "/documents/reindex",
    response_model=StandardResponse[str],
    summary="Reindex a document",
)
async def reindex_document(
    request: IndexRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    # Clear existing chunks
    await knowledge_repository.delete_document(db, request.document_id)
    # Reindex
    return await index_document(request, db)


@router.post(
    "/search",
    response_model=StandardResponse[list[ChunkDto]],
    summary="Search matching chunks using hybrid query rankings",
)
async def search_documents(
    request: SearchRequest,
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[ChunkDto]]:
    chunks = await rag_orchestrator.execute_hybrid_search(
        db, tenant_id, request.query, request.limit, request.department, request.project
    )
    return StandardResponse(data=chunks)


@router.post(
    "/retrieve",
    response_model=RetrieveResponse,
    summary="Retrieve matching segments without metadata filters",
)
async def retrieve_documents(
    request: SearchRequest,
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> RetrieveResponse:
    chunks = await rag_orchestrator.execute_hybrid_search(
        db, tenant_id, request.query, request.limit, request.department, request.project
    )
    return RetrieveResponse(chunks=chunks)


@router.post(
    "/query",
    response_model=QueryResponse,
    summary="Execute grounded query against RAG index returning answers with citations",
)
async def query_knowledge(
    request: QueryRequest,
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> QueryResponse:
    return await rag_orchestrator.execute_grounded_query(
        db, tenant_id, request.query, request.limit, request.department, request.project
    )


@router.get(
    "/documents",
    response_model=StandardResponse[list[DocumentDto]],
    summary="List all registered documents for tenant",
)
async def list_documents(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[DocumentDto]]:
    docs = await knowledge_repository.get_documents(db, tenant_id)
    return StandardResponse(data=docs)


@router.get(
    "/documents/{id}",
    response_model=StandardResponse[DocumentDto],
    summary="Retrieve document metadata",
)
async def get_document_info(
    id: str,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[DocumentDto]:
    d = await knowledge_repository.get_document(db, id)
    if not d:
        raise HTTPException(status_code=404, detail="Document not found.")
    dto = DocumentDto(
        id=d.id,
        filename=d.filename,
        content_type=d.content_type,
        file_size=d.file_size,
        department=d.department,
        project=d.project,
        status=d.status,
        tenant_id=d.tenant_id,
        created_at=d.created_at,
    )
    return StandardResponse(data=dto)


@router.delete(
    "/documents/{id}",
    response_model=StandardResponse[str],
    summary="Delete document record and associated chunk indexes",
)
async def delete_document(
    id: str,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    await knowledge_repository.delete_document(db, id)
    return StandardResponse(data=id, message="Document and associated index entries deleted.")


# Tool Calling & MCP Routers
@router.post(
    "/tools/register",
    response_model=StandardResponse[ToolDto],
    summary="Register a new action tool target and schema config",
)
async def register_tool(
    request: ToolRegisterRequest,
    user_id: str = "sys-admin",
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[ToolDto]:
    dto = await tool_repository.register_tool(
        db,
        request.name,
        request.description,
        request.category,
        request.type,
        request.version_string,
        request.input_schema,
        request.execution_target,
        request.timeout_seconds,
        request.retry_count,
        user_id,
    )
    return StandardResponse(data=dto, message="Tool registered successfully.")


@router.get(
    "/tools",
    response_model=StandardResponse[list[ToolDto]],
    summary="List all registered tools",
)
async def list_tools(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[list[ToolDto]]:
    tools = await tool_repository.get_tools_list(db)
    return StandardResponse(data=tools)


@router.get(
    "/tools/{id}",
    response_model=StandardResponse[ToolDto],
    summary="Get details of registered tool by ID",
)
async def get_tool_details(
    id: str,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[ToolDto]:
    # Search in list
    tools = await tool_repository.get_tools_list(db)
    for t in tools:
        if t.id == id:
            return StandardResponse(data=t)
    raise HTTPException(status_code=404, detail="Tool not found.")


@router.post(
    "/tools/execute",
    response_model=ToolExecuteResponse,
    summary="Execute registered tool parameters under validation limits",
)
async def execute_tool(
    request: ToolExecuteRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ToolExecuteResponse:
    # 1. Resolve tool ID to perform permission check
    tool = await tool_repository.get_tool_by_name(db, request.tool_name)
    if not tool:
        raise HTTPException(status_code=404, detail="Tool not found.")

    # 2. Verify permission limits (default roles mapping)
    permitted = await tool_repository.check_permission(
        db, tool.id, role="admin", tenant_id=request.tenant_id
    )
    if not permitted:
        raise HTTPException(
            status_code=403,
            detail=f"Access Denied: Permissions constraint blocked execution of: {request.tool_name}",
        )

    # 3. Invokes Executor
    return await tool_executor.execute_tool(
        db, request.tool_name, request.version, request.payload, request.user_id, request.tenant_id
    )


@router.post(
    "/tools/discover",
    response_model=StandardResponse[list[dict[str, Any]]],
    summary="Discover active tool schemas formatted for LLMs OpenAI tools syntax",
)
async def discover_tools(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[list[dict[str, Any]]]:
    stmt = select(ToolDb).options(selectinload(ToolDb.versions))
    res = await db.execute(stmt)
    tools_list = res.scalars().all()
    
    openai_tools = []
    for t in tools_list:
        # Find active version
        active_ver = None
        for v in t.versions:
            if v.version_string == t.active_version_string:
                active_ver = v
                break
        
        input_schema = active_ver.input_schema_json if active_ver else {}
        openai_tools.append(
            {
                "type": "function",
                "function": {
                    "name": t.name,
                    "description": t.description,
                    "parameters": input_schema,
                },
            }
        )

    return StandardResponse(data=openai_tools)


# Model Context Protocol (MCP) compatible endpoints
@router.post(
    "/mcp/initialize",
    response_model=McpInitializeResponse,
    summary="Initialize MCP client handshake session capability negotiation",
)
async def mcp_initialize(request: McpInitializeRequest) -> McpInitializeResponse:
    return McpInitializeResponse(
        protocol_version=request.protocol_version,
        capabilities={
            "tools": {"listChanged": True},
            "resources": {"subscribe": True},
            "prompts": {"listChanged": True},
        },
    )


@router.get(
    "/mcp/tools",
    response_model=McpToolsListResponse,
    summary="List all registered tools formatted as standard MCP protocol tools list",
)
async def mcp_list_tools(db: AsyncSession = Depends(get_db_session)) -> McpToolsListResponse:
    stmt = select(ToolDb).options(selectinload(ToolDb.versions))
    res = await db.execute(stmt)
    tools_list = res.scalars().all()

    mcp_tools = []
    for t in tools_list:
        active_ver = None
        for v in t.versions:
            if v.version_string == t.active_version_string:
                active_ver = v
                break
        
        input_schema = active_ver.input_schema_json if active_ver else {}
        mcp_tools.append(
            McpToolDto(
                name=t.name,
                description=t.description,
                inputSchema=input_schema,
            )
        )
    return McpToolsListResponse(tools=mcp_tools)


# Multi-Agent Workflow Routers
@router.post(
    "/agent/graphs",
    response_model=StandardResponse[AgentGraphDto],
    summary="Register a new multi-agent execution graph configuration",
)
async def register_agent_graph(
    name: str,
    description: str,
    start_node: str,
    nodes: list[dict[str, Any]],
    edges: list[dict[str, Any]],
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[AgentGraphDto]:
    try:
        dto = await agent_repository.register_graph(db, name, description, start_node, nodes, edges)
        return StandardResponse(data=dto, message="Agent workflow graph compiled and registered.")
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@router.post(
    "/agent/run",
    response_model=AgentRunResponse,
    summary="Start a new multi-agent execution graph run session",
)
async def run_agent_graph(
    request: AgentRunRequest,
    db: AsyncSession = Depends(get_db_session),
) -> AgentRunResponse:
    return await agent_orchestrator.execute_graph(
        db, request.graph_id, None, request.input_state, request.user_id, request.tenant_id
    )


@router.post(
    "/agent/approve",
    response_model=AgentRunResponse,
    summary="Resolve pending human-in-the-loop sign-off approving/rejecting next agent routing",
)
async def approve_agent_step(
    request: AgentApprovalRequest,
    db: AsyncSession = Depends(get_db_session),
) -> AgentRunResponse:
    return await agent_orchestrator.resume_execution(
        db, request.execution_id, request.approved, request.comments
    )


@router.get(
    "/agent/status/{execution_id}",
    response_model=StandardResponse[dict[str, Any]],
    summary="Retrieve current state checkpoint and status of execution session",
)
async def get_agent_execution_status(
    execution_id: str,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, Any]]:
    checkpoint = await agent_repository.get_checkpoint(db, execution_id)
    if not checkpoint:
        raise HTTPException(status_code=404, detail="Execution session not found.")
    
    return StandardResponse(
        data={
            "execution_id": execution_id,
            "status": checkpoint.status,
            "current_node": checkpoint.current_node,
            "state": checkpoint.state_json,
        }
    )


# Guardrails & Safety Routers
@router.post(
    "/guardrails/validate",
    response_model=GuardrailCheckResponse,
    summary="Verify user prompt inputs or outputs safety rules",
)
async def validate_guardrails(
    request: GuardrailCheckRequest,
    db: AsyncSession = Depends(get_db_session),
) -> GuardrailCheckResponse:
    return await guardrail_engine.check_input(db, request.text, request.user_id, request.tenant_id)


@router.post(
    "/guardrails/check",
    response_model=GuardrailCheckResponse,
    summary="Safety check alias target routing",
)
async def check_guardrails_alias(
    request: GuardrailCheckRequest,
    db: AsyncSession = Depends(get_db_session),
) -> GuardrailCheckResponse:
    return await validate_guardrails(request, db)


@router.post(
    "/guardrails/review",
    response_model=StandardResponse[str],
    summary="Resolve PENDING human override safety reviews status",
)
async def resolve_safety_review(
    request: HumanReviewRequest,
    reviewer_id: str = "security-officer-01",
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    try:
        await guardrail_repository.resolve_human_review(
            db, request.event_id, request.approved, reviewer_id, request.comments
        )
        return StandardResponse(data=request.event_id, message="Human safety review resolved successfully.")
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@router.get(
    "/guardrails/policies",
    response_model=StandardResponse[list[PolicyDto]],
    summary="List all safety policies configs",
)
async def get_safety_policies(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[list[PolicyDto]]:
    policies = await guardrail_repository.get_policies(db)
    return StandardResponse(data=policies)


@router.post(
    "/guardrails/policies",
    response_model=StandardResponse[PolicyDto],
    summary="Register a new safety governance policy",
)
async def create_safety_policy(
    request: PolicyCreateRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[PolicyDto]:
    try:
        dto = await guardrail_repository.create_policy(
            db, request.name, request.code, request.description, request.action
        )
        return StandardResponse(data=dto, message="Safety policy registered successfully.")
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))


@router.put(
    "/guardrails/policies/{code}",
    response_model=StandardResponse[PolicyDto],
    summary="Update safety rules action and status configurations",
)
async def update_safety_policy(
    code: str,
    request: PolicyUpdateRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[PolicyDto]:
    dto = await guardrail_repository.update_policy(db, code, request.action, request.status)
    if not dto:
        raise HTTPException(status_code=404, detail="Target policy not found.")
    return StandardResponse(data=dto, message="Safety policy updated successfully.")


@router.get(
    "/guardrails/audit",
    response_model=StandardResponse[list[SafetyEventDto]],
    summary="Retrieve safety violations events audits trail logs",
)
async def get_safety_audits(
    limit: int = Query(50, ge=1, le=100),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[SafetyEventDto]]:
    events = await guardrail_repository.get_safety_events(db, limit)
    return StandardResponse(data=events)


# AI Evaluation, Benchmarking & Observability Routers
from app.domain.evaluation import (
    EvaluationRequest,
    EvaluationResponse,
    BenchmarkRequest,
    BenchmarkResponse,
    ExperimentRequest,
    ExperimentResponse,
    LeaderboardEntryDto,
    HumanFeedbackRequest,
)

@router.post(
    "/evaluate",
    response_model=EvaluationResponse,
    summary="Evaluate a single LLM request/response run iteration",
)
async def evaluate_run(
    request: EvaluationRequest,
    db: AsyncSession = Depends(get_db_session),
) -> EvaluationResponse:
    return await evaluation_engine.evaluate_run(
        db,
        request.run_id,
        request.text_input,
        request.text_output,
        request.retrieved_context,
        request.cost,
        request.latency_ms,
        request.model,
        request.user_id,
        request.tenant_id,
    )


@router.post(
    "/benchmark",
    response_model=BenchmarkResponse,
    summary="Trigger automated comparative benchmarking against test datasets",
)
async def benchmark_model(
    request: BenchmarkRequest,
    db: AsyncSession = Depends(get_db_session),
) -> BenchmarkResponse:
    return await evaluation_engine.execute_benchmark(
        db, request.name, request.model_name, request.test_prompts
    )


@router.post(
    "/experiment",
    response_model=StandardResponse[str],
    summary="Register a new A/B prompt or configuration variant experiment",
)
async def create_ab_experiment(
    request: ExperimentRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    db_exp = await evaluation_repository.create_experiment(
        db, request.name, request.description, request.variant_a_config, request.variant_b_config
    )
    return StandardResponse(data=db_exp.id, message="A/B Experiment variant registered successfully.")


@router.post(
    "/feedback",
    response_model=StandardResponse[str],
    summary="Submit user thumbs/ratings feedback score for evaluation checks",
)
async def submit_human_feedback(
    request: HumanFeedbackRequest,
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    await evaluation_repository.save_feedback(db, request.evaluation_id, request.rating, request.comments)
    return StandardResponse(data=request.evaluation_id, message="Human evaluation feedback saved.")


@router.get(
    "/leaderboard",
    response_model=StandardResponse[list[LeaderboardEntryDto]],
    summary="Fetch comparative accuracy / latency leaderboard statistics per model",
)
async def get_model_leaderboard(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[list[LeaderboardEntryDto]]:
    entries = await evaluation_repository.get_leaderboard(db)
    return StandardResponse(data=entries)


@router.get(
    "/quality",
    response_model=StandardResponse[dict[str, float]],
    summary="Retrieve average faithfulness, correctness, and relevance metrics",
)
async def get_quality_analytics(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[dict[str, float]]:
    evals = await evaluation_repository.get_evaluations(db)
    if not evals:
        return StandardResponse(
            data={"faithfulness": 1.0, "correctness": 1.0, "relevance": 1.0, "groundedness": 1.0, "overall": 1.0}
        )
    
    avg_f = sum(e.faithfulness_score for e in evals) / len(evals)
    avg_c = sum(e.correctness_score for e in evals) / len(evals)
    avg_r = sum(e.relevance_score for e in evals) / len(evals)
    avg_g = sum(e.groundedness_score for e in evals) / len(evals)
    avg_o = sum(e.overall_score for e in evals) / len(evals)
    
    return StandardResponse(
        data={
            "faithfulness": round(avg_f, 2),
            "correctness": round(avg_c, 2),
            "relevance": round(avg_r, 2),
            "groundedness": round(avg_g, 2),
            "overall": round(avg_o, 2),
        }
    )


@router.get(
    "/cost",
    response_model=StandardResponse[dict[str, Any]],
    summary="Retrieve total model tokens billing costs parameters",
)
async def get_cost_analytics(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[dict[str, Any]]:
    evals = await evaluation_repository.get_evaluations(db)
    total_cost = sum(e.cost for e in evals)
    
    # Cost per model group mapping
    model_costs = {}
    for e in evals:
        model_costs[e.model] = model_costs.get(e.model, 0.0) + e.cost
        
    return StandardResponse(
        data={
            "total_accumulated_cost": round(total_cost, 5),
            "breakdown_by_model": {m: round(c, 5) for m, c in model_costs.items()},
        }
    )


@router.get(
    "/latency",
    response_model=StandardResponse[dict[str, Any]],
    summary="Retrieve latency distributions metrics",
)
async def get_latency_analytics(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[dict[str, Any]]:
    evals = await evaluation_repository.get_evaluations(db)
    if not evals:
        return StandardResponse(data={"average_latency_ms": 0, "max_latency_ms": 0})
    
    total_lat = sum(e.latency_ms for e in evals)
    max_lat = max(e.latency_ms for e in evals)
    avg_lat = int(total_lat / len(evals))
    
    return StandardResponse(
        data={
            "average_latency_ms": avg_lat,
            "max_latency_ms": max_lat,
            "requests_count": len(evals),
        }
    )


@router.get(
    "/hallucinations",
    response_model=StandardResponse[dict[str, int]],
    summary="Retrieve counts of ungrounded or unverified response segments",
)
async def get_hallucinations_analytics(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[dict[str, int]]:
    evals = await evaluation_repository.get_evaluations(db)
    hallucination_runs = sum(1 for e in evals if e.groundedness_score < 0.8)
    return StandardResponse(
        data={
            "hallucination_flagged_runs": hallucination_runs,
            "total_runs": len(evals),
        }
    )


@router.get(
    "/reports",
    response_model=StandardResponse[dict[str, Any]],
    summary="Yield detailed statistical aggregation report",
)
async def get_aggregated_reports(db: AsyncSession = Depends(get_db_session)) -> StandardResponse[dict[str, Any]]:
    # Combine latency, cost, and quality data
    q = await get_quality_analytics(db)
    c = await get_cost_analytics(db)
    l = await get_latency_analytics(db)
    
    return StandardResponse(
        data={
            "quality": q.data,
            "billing": c.data,
            "latency": l.data,
            "timestamp": str(datetime.datetime.utcnow()),
        }
    )


# Incident Investigation Agent Routers
@router.post(
    "/incident/investigate",
    response_model=IncidentInvestigationDto,
    summary="Register a new safety incident intake and run initial RAG investigation",
)
async def investigate_incident(
    request: IncidentIntakeRequest,
    db: AsyncSession = Depends(get_db_session),
) -> IncidentInvestigationDto:
    return await incident_agent.investigate_incident(db, request)


@router.post(
    "/incident/root-cause",
    response_model=RootCauseResponse,
    summary="Generate Root Cause Analysis (RCA) chain (e.g. 5 Whys) for incident",
)
async def analyze_root_cause(
    request: RootCauseRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RootCauseResponse:
    return await incident_agent.analyze_root_cause(db, request)


@router.post(
    "/incident/recommend",
    response_model=RecommendationResponse,
    summary="Determine corrective and preventive actions for safety root causes",
)
async def recommend_actions(
    request: RecommendationRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RecommendationResponse:
    return await incident_agent.recommend_corrective_actions(db, request.investigation_id)


@router.post(
    "/incident/report",
    response_model=ReportResponse,
    summary="Compile final PDF-ready Markdown EHS incident investigation report",
)
async def compile_incident_report(
    request: RecommendationRequest,  # Reuses same shape
    db: AsyncSession = Depends(get_db_session),
) -> ReportResponse:
    return await incident_agent.compile_report(db, request.investigation_id)


@router.post(
    "/incident/summarize",
    response_model=StandardResponse[str],
    summary="Heuristic incident text summarizer",
)
async def summarize_incident(
    description: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[str]:
    chat_req = ChatRequest(
        model="gpt-4o",
        messages=[Message(role="user", content=f"Summarize incident: {description}")],
    )
    chat_res = await provider_factory.execute_chat(chat_req)
    return StandardResponse(data=chat_res.choices[0].message.content, message="Summary generated.")


@router.get(
    "/incident/history",
    response_model=StandardResponse[list[IncidentInvestigationDto]],
    summary="Retrieve history log of safety investigations for tenant",
)
async def get_incident_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[IncidentInvestigationDto]]:
    history = await incident_repository.get_investigations_list(db, tenant_id)
    return StandardResponse(data=history)


# Safety Compliance Agent Routers
@router.post(
    "/compliance/check",
    response_model=ComplianceAssessmentDto,
    summary="Evaluate activity details compliance against regulatory frameworks",
)
async def check_compliance(
    request: ComplianceCheckRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ComplianceAssessmentDto:
    return await safety_compliance_agent.check_compliance(db, request)


@router.post(
    "/compliance/audit",
    response_model=ComplianceAssessmentDto,
    summary="Audit operational permits/PPE compliance checklists",
)
async def audit_compliance(
    request: ComplianceCheckRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ComplianceAssessmentDto:
    return await safety_compliance_agent.check_compliance(db, request)


@router.post(
    "/compliance/capa",
    response_model=CapaPlanResponse,
    summary="Compile CAPA recommendations for flagged compliance violations",
)
async def generate_compliance_capa(
    request: CapaPlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> CapaPlanResponse:
    return await safety_compliance_agent.generate_capa(db, request)


@router.post(
    "/compliance/recommend",
    response_model=CapaPlanResponse,
    summary="CAPA recommendations wrapper alias",
)
async def recommend_compliance_capa(
    request: CapaPlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> CapaPlanResponse:
    return await safety_compliance_agent.generate_capa(db, request)


@router.post(
    "/compliance/report",
    response_model=ComplianceReportResponse,
    summary="Compile final PDF-ready Markdown EHS compliance report",
)
async def compile_compliance_report(
    request: CapaPlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ComplianceReportResponse:
    return await safety_compliance_agent.compile_report(db, request.assessment_id)


@router.get(
    "/compliance/history",
    response_model=StandardResponse[list[ComplianceAssessmentDto]],
    summary="Retrieve history log of compliance checks for tenant",
)
async def get_compliance_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[ComplianceAssessmentDto]]:
    history = await compliance_repository.get_assessments_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/compliance/score",
    response_model=StandardResponse[dict[str, float]],
    summary="Retrieve average compliance score metrics across tenant assessments",
)
async def get_site_compliance_score(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, float]]:
    history = await compliance_repository.get_assessments_list(db, tenant_id)
    if not history:
        return StandardResponse(data={"average_score": 100.0})
    avg_score = sum(h.compliance_score for h in history) / len(history)
    return StandardResponse(data={"average_score": round(avg_score, 2)})


# Permit-to-Work Agent Routers
@router.post(
    "/permit/create",
    response_model=PermitRecordDto,
    summary="Register a new permit-to-work draft request",
)
async def create_permit(
    request: PermitCreateRequest,
    db: AsyncSession = Depends(get_db_session),
) -> PermitRecordDto:
    return await permit_agent.create_permit(db, request)


@router.post(
    "/permit/validate",
    response_model=PermitValidationResponse,
    summary="Validate permit completeness and hazard controls checklist",
)
async def validate_permit(
    permit_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> PermitValidationResponse:
    return await permit_agent.validate_permit(db, permit_id)


@router.post(
    "/permit/review",
    response_model=PermitValidationResponse,
    summary="Review permit completeness alias wrapper",
)
async def review_permit(
    permit_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> PermitValidationResponse:
    return await permit_agent.validate_permit(db, permit_id)


@router.post(
    "/permit/risk",
    response_model=RiskAssessmentResponse,
    summary="Calculate 5x5 Likelihood x Severity Risk Matrix for permit",
)
async def assess_permit_risk(
    request: RiskAssessmentRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RiskAssessmentResponse:
    return await permit_agent.assess_risk(db, request)


@router.post(
    "/permit/recommend",
    response_model=ApprovalRecommendationResponse,
    summary="Evaluate approval recommendation (APPROVE / REJECT / SUSPEND) for permit",
)
async def recommend_permit_approval(
    request: ApprovalRecommendationRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ApprovalRecommendationResponse:
    return await permit_agent.recommend_approval(db, request)


@router.post(
    "/permit/approve",
    response_model=PermitRecordDto,
    summary="Record formal supervisor / safety officer approval or rejection decision",
)
async def approve_permit(
    permit_id: str = Query(...),
    approver_id: str = Query("approver-01"),
    role: str = Query("SAFETY_OFFICER"),
    approved: bool = Query(True),
    comments: str | None = Query(None),
    db: AsyncSession = Depends(get_db_session),
) -> PermitRecordDto:
    return await permit_agent.approve_permit(db, permit_id, approver_id, role, approved, comments)


@router.get(
    "/permit/history",
    response_model=StandardResponse[list[PermitRecordDto]],
    summary="Retrieve history log of permit records for tenant",
)
async def get_permit_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[PermitRecordDto]]:
    history = await permit_repository.get_permits_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/permit/templates",
    response_model=StandardResponse[list[str]],
    summary="List supported high-risk permit templates",
)
async def list_permit_templates() -> StandardResponse[list[str]]:
    templates = [
        "HOT_WORK",
        "COLD_WORK",
        "CONFINED_SPACE",
        "HEIGHT",
        "LOTO",
        "ELECTRICAL",
        "EXCAVATION",
        "LIFTING",
    ]
    return StandardResponse(data=templates)


# Enterprise Risk Assessment Agent Routers
@router.post(
    "/risk/assess",
    response_model=RiskAssessmentDto,
    summary="Evaluate operational hazard risks using ISO 31000 & multi-methodology frameworks",
)
async def assess_risk(
    request: EnterpriseRiskAssessmentRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RiskAssessmentDto:
    return await enterprise_risk_agent.assess_risk(db, request)


@router.post(
    "/risk/predict",
    response_model=RiskPredictResponse,
    summary="Predict future risk trends based on weather, asset health, and historical incidents",
)
async def predict_risk(
    request: RiskPredictRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RiskPredictResponse:
    return await enterprise_risk_agent.predict_risk(db, request)


@router.post(
    "/risk/mitigate",
    response_model=MitigationPlanResponse,
    summary="Generate engineering, administrative, and PPE control mitigation plans",
)
async def generate_risk_mitigation(
    request: MitigationPlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> MitigationPlanResponse:
    return await enterprise_risk_agent.generate_mitigation_plan(db, request)


@router.post(
    "/risk/review",
    response_model=RiskAssessmentDto,
    summary="Risk assessment review alias wrapper",
)
async def review_risk(
    request: EnterpriseRiskAssessmentRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RiskAssessmentDto:
    return await enterprise_risk_agent.assess_risk(db, request)


@router.post(
    "/risk/matrix",
    response_model=StandardResponse[dict[str, Any]],
    summary="Retrieve 5x5 Likelihood x Severity Risk Matrix threshold configurations",
)
async def get_risk_matrix() -> StandardResponse[dict[str, Any]]:
    matrix_info = {
        "dimensions": "5x5",
        "categories": {
            "CRITICAL": "Score > 18",
            "HIGH": "Score 13-18",
            "MEDIUM": "Score 7-12",
            "LOW": "Score 1-6",
        },
    }
    return StandardResponse(data=matrix_info)


@router.get(
    "/risk/history",
    response_model=StandardResponse[list[RiskAssessmentDto]],
    summary="Retrieve history log of risk assessments for tenant",
)
async def get_risk_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[RiskAssessmentDto]]:
    history = await risk_repository.get_assessments_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/risk/register",
    response_model=StandardResponse[list[RiskRegisterEntryDto]],
    summary="Retrieve active system-wide Risk Register entries",
)
async def get_risk_register(
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[RiskRegisterEntryDto]]:
    register = await risk_repository.get_register_entries(db)
    return StandardResponse(data=register)


@router.get(
    "/risk/trends",
    response_model=StandardResponse[dict[str, Any]],
    summary="Retrieve enterprise risk trends summary across departments",
)
async def get_risk_trends(
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, Any]]:
    register = await risk_repository.get_register_entries(db)
    increasing = [r for r in register if r.trend_status == "INCREASING"]
    stable = [r for r in register if r.trend_status == "STABLE"]
    decreasing = [r for r in register if r.trend_status == "DECREASING"]
    trends_summary = {
        "total_hazards_logged": len(register),
        "increasing_count": len(increasing),
        "stable_count": len(stable),
        "decreasing_count": len(decreasing),
    }
    return StandardResponse(data=trends_summary)


# Intelligent Inspection Agent Routers
@router.post(
    "/inspection/plan",
    response_model=InspectionRecordDto,
    summary="Plan workplace safety inspection session",
)
async def plan_inspection(
    request: InspectionPlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> InspectionRecordDto:
    return await intelligent_inspection_agent.plan_inspection(db, request)


@router.post(
    "/inspection/checklist",
    response_model=ChecklistGenerationResponse,
    summary="Generate dynamic risk-based inspection checklist items",
)
async def generate_inspection_checklist(
    request: ChecklistGenerationRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ChecklistGenerationResponse:
    return await intelligent_inspection_agent.generate_checklist(db, request)


@router.post(
    "/inspection/analyze",
    response_model=FindingAnalysisResponse,
    summary="Classify inspection findings and invoke Risk Assessment Agent for major non-conformances",
)
async def analyze_inspection_finding(
    request: FindingAnalysisRequest,
    db: AsyncSession = Depends(get_db_session),
) -> FindingAnalysisResponse:
    return await intelligent_inspection_agent.analyze_finding(db, request)


@router.post(
    "/inspection/recommend",
    response_model=FindingAnalysisResponse,
    summary="Finding analysis alias wrapper",
)
async def recommend_inspection_finding(
    request: FindingAnalysisRequest,
    db: AsyncSession = Depends(get_db_session),
) -> FindingAnalysisResponse:
    return await intelligent_inspection_agent.analyze_finding(db, request)


@router.post(
    "/inspection/report",
    response_model=InspectionReportResponse,
    summary="Compile final PDF-ready Markdown EHS inspection report",
)
async def compile_inspection_report(
    inspection_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> InspectionReportResponse:
    return await intelligent_inspection_agent.compile_report(db, inspection_id)


@router.post(
    "/inspection/summarize",
    response_model=InspectionReportResponse,
    summary="Report compilation alias wrapper",
)
async def summarize_inspection(
    inspection_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> InspectionReportResponse:
    return await intelligent_inspection_agent.compile_report(db, inspection_id)


@router.get(
    "/inspection/history",
    response_model=StandardResponse[list[InspectionRecordDto]],
    summary="Retrieve history log of workplace inspections for tenant",
)
async def get_inspection_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[InspectionRecordDto]]:
    history = await inspection_repository.get_inspections_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/inspection/templates",
    response_model=StandardResponse[list[str]],
    summary="List supported inspection type templates",
)
async def list_inspection_templates() -> StandardResponse[list[str]]:
    templates = [
        "DAILY_SAFETY",
        "EQUIPMENT",
        "FIRE_SAFETY",
        "CHEMICAL_STORAGE",
        "CONSTRUCTION",
        "WAREHOUSE",
    ]
    return StandardResponse(data=templates)


# Contractor Safety Agent Routers
@router.post(
    "/contractor/register",
    response_model=ContractorRecordDto,
    summary="Register a new contractor profile for safety management",
)
async def register_contractor(
    request: ContractorRegisterRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ContractorRecordDto:
    return await contractor_safety_agent.register_contractor(db, request)


@router.post(
    "/contractor/qualify",
    response_model=ContractorQualifyResponse,
    summary="Evaluate contractor pre-qualification eligibility and missing training certifications",
)
async def qualify_contractor(
    request: ContractorQualifyRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ContractorQualifyResponse:
    return await contractor_safety_agent.qualify_contractor(db, request)


@router.post(
    "/contractor/verify",
    response_model=ContractorVerifyResponse,
    summary="Verify contractor licenses, insurance validity, or medical fitness certificates",
)
async def verify_contractor_document(
    request: ContractorVerifyRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ContractorVerifyResponse:
    return await contractor_safety_agent.verify_document(db, request)


@router.post(
    "/contractor/risk",
    response_model=RiskAssessmentDto,
    summary="Evaluate task risk for contractor workforce using Risk Assessment Agent",
)
async def assess_contractor_risk(
    contractor_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> RiskAssessmentDto:
    contractor = await contractor_repository.get_contractor(db, contractor_id)
    if not contractor:
        raise ValueError("Contractor profile not found.")
    risk_req = EnterpriseRiskAssessmentRequest(
        title=f"Contractor Task Risk: {contractor.company_name}",
        risk_type="OCCUPATIONAL",
        methodology="HIRA",
        details=f"Evaluating workforce safety risks for {contractor.contractor_type}.",
        location="Site Operations",
        user_id=contractor.user_id,
        tenant_id=contractor.tenant_id,
    )
    return await enterprise_risk_agent.assess_risk(db, risk_req)


@router.post(
    "/contractor/recommend",
    response_model=ContractorQualifyResponse,
    summary="Contractor qualification recommendation alias wrapper",
)
async def recommend_contractor(
    request: ContractorQualifyRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ContractorQualifyResponse:
    return await contractor_safety_agent.qualify_contractor(db, request)


@router.post(
    "/contractor/report",
    response_model=StandardResponse[dict[str, Any]],
    summary="Generate comprehensive contractor qualification & safety compliance report",
)
async def generate_contractor_report(
    contractor_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, Any]]:
    scorecard = await contractor_safety_agent.generate_scorecard(db, contractor_id)
    report_data = {
        "contractor_id": scorecard.contractor_id,
        "company_name": scorecard.company_name,
        "safety_score": scorecard.safety_score,
        "compliance_status": scorecard.compliance_status,
        "summary": f"Contractor {scorecard.company_name} safety evaluation score {scorecard.safety_score}/100.",
    }
    return StandardResponse(data=report_data)


@router.get(
    "/contractor/history",
    response_model=StandardResponse[list[ContractorRecordDto]],
    summary="Retrieve history log of registered contractors for tenant",
)
async def get_contractor_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[ContractorRecordDto]]:
    history = await contractor_repository.get_contractors_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/contractor/scorecard",
    response_model=StandardResponse[ContractorScorecardResponse],
    summary="Retrieve contractor performance scorecard",
)
async def get_contractor_scorecard(
    contractor_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[ContractorScorecardResponse]:
    scorecard = await contractor_safety_agent.generate_scorecard(db, contractor_id)
    return StandardResponse(data=scorecard)


# Asset & Equipment Safety Agent Routers
@router.post(
    "/asset/health",
    response_model=AssetRecordDto,
    summary="Evaluate operational health score for heavy machinery & industrial equipment",
)
async def evaluate_asset_health(
    request: AssetHealthCheckRequest,
    db: AsyncSession = Depends(get_db_session),
) -> AssetRecordDto:
    return await asset_equipment_safety_agent.evaluate_health(db, request)


@router.post(
    "/asset/inspect",
    response_model=ChecklistGenerationResponse,
    summary="Trigger dynamic asset inspection checklist via Intelligent Inspection Agent",
)
async def inspect_asset(
    asset_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> ChecklistGenerationResponse:
    return await asset_equipment_safety_agent.trigger_asset_inspection(db, asset_id)


@router.post(
    "/asset/recommend",
    response_model=StandardResponse[dict[str, Any]],
    summary="Generate preventive maintenance & overhaul recommendations for asset",
)
async def recommend_asset_maintenance(
    asset_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, Any]]:
    asset = await asset_repository.get_asset(db, asset_id)
    if not asset:
        raise ValueError("Asset record not found.")
    rec_data = {
        "asset_id": asset.id,
        "asset_name": asset.asset_name,
        "recommendation": "Perform vibration analysis and bearing lubrication overhaul.",
        "urgency": "HIGH" if asset.health_score < 70.0 else "NORMAL",
    }
    return StandardResponse(data=rec_data)


@router.post(
    "/asset/risk",
    response_model=RiskAssessmentDto,
    summary="Evaluate operational hazard risk for asset via Risk Assessment Agent",
)
async def assess_asset_risk(
    asset_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> RiskAssessmentDto:
    asset = await asset_repository.get_asset(db, asset_id)
    if not asset:
        raise ValueError("Asset record not found.")
    risk_req = EnterpriseRiskAssessmentRequest(
        title=f"Asset Operational Risk: {asset.asset_name}",
        risk_type="MECHANICAL",
        methodology="FMEA",
        details=f"Operational health check evaluation score {asset.health_score}/100.",
        location=asset.location,
        user_id=asset.user_id,
        tenant_id=asset.tenant_id,
    )
    return await enterprise_risk_agent.assess_risk(db, risk_req)


@router.post(
    "/asset/report",
    response_model=StandardResponse[dict[str, Any]],
    summary="Generate asset health & operational readiness report",
)
async def generate_asset_report(
    asset_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, Any]]:
    scorecard = await asset_equipment_safety_agent.generate_scorecard(db, asset_id)
    report_data = {
        "asset_id": scorecard.asset_id,
        "asset_name": scorecard.asset_name,
        "health_score": scorecard.health_score,
        "operational_status": scorecard.operational_status,
        "summary": f"Asset {scorecard.asset_name} health score is {scorecard.health_score}/100. Status: {scorecard.operational_status}.",
    }
    return StandardResponse(data=report_data)


@router.post(
    "/asset/shutdown",
    response_model=AssetShutdownResponse,
    summary="Trigger emergency safety shutdown decision for compromised asset",
)
async def shutdown_asset(
    request: AssetShutdownRequest,
    db: AsyncSession = Depends(get_db_session),
) -> AssetShutdownResponse:
    return await asset_equipment_safety_agent.recommend_shutdown(db, request)


@router.get(
    "/asset/history",
    response_model=StandardResponse[list[AssetRecordDto]],
    summary="Retrieve history log of equipment safety records for tenant",
)
async def get_asset_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[AssetRecordDto]]:
    history = await asset_repository.get_assets_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/asset/scorecard",
    response_model=StandardResponse[AssetScorecardResponse],
    summary="Retrieve asset integrity scorecard",
)
async def get_asset_scorecard(
    asset_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[AssetScorecardResponse]:
    scorecard = await asset_equipment_safety_agent.generate_scorecard(db, asset_id)
    return StandardResponse(data=scorecard)


# Emergency Response Agent Routers
@router.post(
    "/emergency/assess",
    response_model=EmergencyRecordDto,
    summary="Classify real-time workplace emergency type and severity level",
)
async def assess_emergency(
    request: EmergencyAssessRequest,
    db: AsyncSession = Depends(get_db_session),
) -> EmergencyRecordDto:
    return await emergency_response_agent.assess_emergency(db, request)


@router.post(
    "/emergency/respond",
    response_model=EmergencyRecordDto,
    summary="Emergency response assessment alias wrapper",
)
async def respond_emergency(
    request: EmergencyAssessRequest,
    db: AsyncSession = Depends(get_db_session),
) -> EmergencyRecordDto:
    return await emergency_response_agent.assess_emergency(db, request)


@router.post(
    "/emergency/evacuate",
    response_model=EvacuationPlanResponse,
    summary="Generate safe primary & alternative evacuation routes and assembly points",
)
async def plan_evacuation(
    request: EvacuationPlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> EvacuationPlanResponse:
    return await emergency_response_agent.plan_evacuation(db, request.emergency_id)


@router.post(
    "/emergency/notify",
    response_model=EmergencyNotifyResponse,
    summary="Broadcast emergency alerts across PA systems, SMS, and Mobile App channels",
)
async def notify_emergency_responders(
    request: EmergencyNotifyRequest,
    db: AsyncSession = Depends(get_db_session),
) -> EmergencyNotifyResponse:
    return await emergency_response_agent.notify_responders(db, request)


@router.post(
    "/emergency/report",
    response_model=SitrepResponse,
    summary="Compile Situation Report (SITREP) tactical briefing",
)
async def generate_emergency_report(
    emergency_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> SitrepResponse:
    return await emergency_response_agent.generate_sitrep(db, emergency_id)


@router.post(
    "/emergency/sitrep",
    response_model=SitrepResponse,
    summary="SITREP briefing report alias wrapper",
)
async def generate_emergency_sitrep(
    emergency_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> SitrepResponse:
    return await emergency_response_agent.generate_sitrep(db, emergency_id)


@router.get(
    "/emergency/history",
    response_model=StandardResponse[list[EmergencyRecordDto]],
    summary="Retrieve history log of emergency events for tenant",
)
async def get_emergency_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[EmergencyRecordDto]]:
    history = await emergency_repository.get_emergencies_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/emergency/resources",
    response_model=StandardResponse[list[EmergencyResourceDto]],
    summary="Retrieve active emergency response equipment inventory",
)
async def get_emergency_resources(
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[EmergencyResourceDto]]:
    resources = await emergency_repository.get_resources_list(db)
    return StandardResponse(data=resources)


# Predictive Maintenance Agent Routers
@router.post(
    "/maintenance/predict",
    response_model=MaintenancePredictionDto,
    summary="Predict failure probability and recommended maintenance strategy for equipment",
)
async def predict_maintenance_failure(
    request: MaintenancePredictRequest,
    db: AsyncSession = Depends(get_db_session),
) -> MaintenancePredictionDto:
    return await predictive_maintenance_agent.predict_failure(db, request)


@router.post(
    "/maintenance/recommend",
    response_model=MaintenancePredictionDto,
    summary="Predictive maintenance recommendation alias wrapper",
)
async def recommend_maintenance_strategy(
    request: MaintenancePredictRequest,
    db: AsyncSession = Depends(get_db_session),
) -> MaintenancePredictionDto:
    return await predictive_maintenance_agent.predict_failure(db, request)


@router.post(
    "/maintenance/rul",
    response_model=RulEstimateResponse,
    summary="Estimate Remaining Useful Life (RUL) in hours & days",
)
async def estimate_maintenance_rul(
    request: RulEstimateRequest,
    db: AsyncSession = Depends(get_db_session),
) -> RulEstimateResponse:
    return await predictive_maintenance_agent.estimate_rul(db, request)


@router.post(
    "/maintenance/plan",
    response_model=MaintenancePlanResponse,
    summary="Generate condition-based work order plan & spare parts allocation",
)
async def plan_maintenance_work_order(
    request: MaintenancePlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> MaintenancePlanResponse:
    return await predictive_maintenance_agent.create_maintenance_plan(db, request)


@router.post(
    "/maintenance/report",
    response_model=StandardResponse[dict[str, Any]],
    summary="Generate predictive maintenance & asset reliability executive report",
)
async def generate_maintenance_report(
    asset_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, Any]]:
    rul_res = await predictive_maintenance_agent.estimate_rul(db, RulEstimateRequest(asset_id=asset_id))
    report_data = {
        "asset_id": asset_id,
        "remaining_useful_life_hours": rul_res.remaining_useful_life_hours,
        "degradation_trend": rul_res.degradation_trend,
        "summary": f"Asset RUL is estimated at {rul_res.remaining_useful_life_hours} hours. Trend: {rul_res.degradation_trend}.",
    }
    return StandardResponse(data=report_data)


@router.get(
    "/maintenance/history",
    response_model=StandardResponse[list[MaintenancePredictionDto]],
    summary="Retrieve history log of maintenance predictions for tenant",
)
async def get_maintenance_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[MaintenancePredictionDto]]:
    history = await maintenance_repository.get_predictions_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/maintenance/dashboard",
    response_model=StandardResponse[ReliabilityDashboardResponse],
    summary="Retrieve plant-wide asset reliability & maintenance metrics dashboard",
)
async def get_maintenance_dashboard(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[ReliabilityDashboardResponse]:
    metrics = await predictive_maintenance_agent.get_dashboard_metrics(db, tenant_id)
    return StandardResponse(data=metrics)


# Executive Reporting & Analytics Agent Routers
@router.post(
    "/analytics/dashboard",
    response_model=ExecutiveDashboardResponse,
    summary="Generate CEO / Plant / Safety / Maintenance / Compliance executive dashboard",
)
async def generate_executive_dashboard(
    request: ExecutiveDashboardRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ExecutiveDashboardResponse:
    return await executive_reporting_agent.generate_dashboard(db, request)


@router.post(
    "/analytics/report",
    response_model=ReportGenerateResponse,
    summary="Compile board-level executive report and strategic recommendations",
)
async def generate_executive_report(
    request: ReportGenerateRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ReportGenerateResponse:
    return await executive_reporting_agent.generate_report(db, request)


@router.post(
    "/analytics/forecast",
    response_model=ForecastResponse,
    summary="Predict enterprise metric trends over custom month horizon",
)
async def forecast_analytics_metric(
    request: ForecastRequest,
    db: AsyncSession = Depends(get_db_session),
) -> ForecastResponse:
    return await executive_reporting_agent.forecast_metric(db, request)


@router.post(
    "/analytics/kpi",
    response_model=StandardResponse[KpiMetricsResponse],
    summary="Calculate enterprise safety, permit, inspection, and asset KPIs",
)
async def get_kpi_metrics(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[KpiMetricsResponse]:
    kpis = await analytics_repository.get_latest_kpis(db, tenant_id)
    return StandardResponse(data=kpis)


@router.post(
    "/analytics/benchmark",
    response_model=UnitBenchmarkResponse,
    summary="Benchmark performance across operating units, plants, or business departments",
)
async def benchmark_operating_units(
    request: UnitBenchmarkRequest,
    db: AsyncSession = Depends(get_db_session),
) -> UnitBenchmarkResponse:
    return await executive_reporting_agent.benchmark_units(db, request)


@router.post(
    "/analytics/summary",
    response_model=StandardResponse[dict[str, Any]],
    summary="Generate natural language narrative executive summary",
)
async def generate_executive_summary(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[dict[str, Any]]:
    kpis = await analytics_repository.get_latest_kpis(db, tenant_id)
    summary_data = {
        "tenant_id": tenant_id,
        "trir": kpis.trir,
        "overall_compliance_score": kpis.overall_compliance_score,
        "executive_narrative": f"Enterprise safety score is {kpis.overall_compliance_score}%. TRIR is maintained at {kpis.trir}.",
    }
    return StandardResponse(data=summary_data)


@router.get(
    "/analytics/history",
    response_model=StandardResponse[list[ReportGenerateResponse]],
    summary="Retrieve history log of generated executive board reports for tenant",
)
async def get_analytics_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[ReportGenerateResponse]]:
    history = await analytics_repository.get_reports_list(db, tenant_id)
    return StandardResponse(data=history)


@router.get(
    "/analytics/insights",
    response_model=StandardResponse[list[dict[str, Any]]],
    summary="Retrieve AI-generated strategic insights, emerging risks, and leading indicators",
)
async def get_analytics_insights(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[dict[str, Any]]]:
    insights = [
        {
            "category": "PREDICTIVE_MAINTENANCE",
            "title": "Vibration anomaly in Compressor Station 3",
            "impact": "HIGH",
            "recommendation": "Perform lubrication overhaul within 72 hours to prevent forced shutdown.",
        },
        {
            "category": "PERMIT_COMPLIANCE",
            "title": "Permit Sign-off Audit Excellence",
            "impact": "MEDIUM",
            "recommendation": "Maintain 100% supervisor verification logs across all high-temperature zones.",
        },
    ]
    return StandardResponse(data=insights)


# Multi-Agent Supervisor & Collaboration Agent Routers
@router.post(
    "/supervisor/chat",
    response_model=SupervisorChatResponse,
    summary="Unified AI Operating System chat entrypoint orchestrating all 10 domain agents",
)
async def supervisor_chat(
    request: SupervisorChatRequest,
    db: AsyncSession = Depends(get_db_session),
) -> SupervisorChatResponse:
    return await multi_agent_supervisor.process_query(db, request)


@router.post(
    "/supervisor/plan",
    response_model=SupervisorPlanResponse,
    summary="Decompose complex user prompt into multi-agent task execution DAG plan",
)
async def supervisor_plan(
    request: SupervisorPlanRequest,
    db: AsyncSession = Depends(get_db_session),
) -> SupervisorPlanResponse:
    return await multi_agent_supervisor.plan_query(db, request)


@router.post(
    "/supervisor/execute",
    response_model=SupervisorChatResponse,
    summary="Execute multi-agent subtask dispatch plan alias wrapper",
)
async def supervisor_execute(
    request: SupervisorChatRequest,
    db: AsyncSession = Depends(get_db_session),
) -> SupervisorChatResponse:
    return await multi_agent_supervisor.process_query(db, request)


@router.post(
    "/supervisor/workflow",
    response_model=SupervisorChatResponse,
    summary="Multi-agent workflow orchestration alias wrapper",
)
async def supervisor_workflow(
    request: SupervisorChatRequest,
    db: AsyncSession = Depends(get_db_session),
) -> SupervisorChatResponse:
    return await multi_agent_supervisor.process_query(db, request)


@router.get(
    "/supervisor/status",
    response_model=SupervisorStatusResponse,
    summary="Retrieve status of active supervisor session and subtask completion progress",
)
async def get_supervisor_status(
    session_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> SupervisorStatusResponse:
    sess = await supervisor_repository.get_session(db, session_id)
    if not sess:
        raise ValueError("Supervisor session not found.")
    tasks = sess.tasks
    return SupervisorStatusResponse(
        session_id=sess.id,
        status=sess.plan_status,
        completed_subtasks_count=len([t for t in tasks if t.status == "SUCCESS"]),
        total_subtasks_count=max(len(tasks), 1),
    )


@router.get(
    "/supervisor/history",
    response_model=StandardResponse[list[dict[str, Any]]],
    summary="Retrieve history log of supervisor sessions for tenant",
)
async def get_supervisor_history(
    tenant_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> StandardResponse[list[dict[str, Any]]]:
    sessions = await supervisor_repository.get_sessions_list(db, tenant_id)
    history = [
        {
            "session_id": s.id,
            "user_query": s.user_query,
            "status": s.plan_status,
            "consensus_summary": s.unified_response,
            "created_at": str(s.created_at),
        }
        for s in sessions
    ]
    return StandardResponse(data=history)


@router.get(
    "/supervisor/graph",
    response_model=SupervisorDagResponse,
    summary="Retrieve execution DAG graph nodes & edges for supervisor session",
)
async def get_supervisor_graph(
    session_id: str = Query(...),
    db: AsyncSession = Depends(get_db_session),
) -> SupervisorDagResponse:
    return await multi_agent_supervisor.get_dag(db, session_id)


# Metadata Registry Routers
@router.get(
    "/models",
    response_model=dict[str, ModelMetadata],
    summary="List metadata parameters for registered models",
)
async def list_models() -> dict[str, ModelMetadata]:
    """Exposes supported models matrix."""
    return MODEL_REGISTRY


@router.get(
    "/providers",
    response_model=list[str],
    summary="List active EHS platform LLM provider keys",
)
async def list_providers() -> list[str]:
    """Exposes active adapter options."""
    return ["openai", "anthropic", "gemini", "bedrock", "ollama", "mock"]
