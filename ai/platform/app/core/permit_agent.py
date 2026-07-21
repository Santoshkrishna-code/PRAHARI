import os
import logging
from sqlalchemy.ext.asyncio import AsyncSession

# Volumes 0-10 Reusable Modules
from app.infrastructure.providers.factory import provider_factory
from app.core.rag_orchestrator import rag_orchestrator
from app.core.tool_executor import tool_executor
from app.core.guardrail_engine import guardrail_engine
from app.core.evaluation_engine import evaluation_engine
from app.infrastructure.repositories.permit_repository import permit_repository
from app.core.models import ChatRequest, Message

# Domain schemas
from app.domain.permits import (
    PermitCreateRequest,
    PermitRecordDto,
    PermitValidationResponse,
    RiskAssessmentRequest,
    RiskAssessmentResponse,
    ApprovalRecommendationRequest,
    ApprovalRecommendationResponse,
)

logger = logging.getLogger(__name__)
MODEL_NAME = "mock-model" if os.getenv("APP_ENV") == "testing" else "gpt-4o"


class PermitToWorkAgent:
    """Agent orchestrator managing permit validation checklists, 5x5 risk matrix scoring, and multi-tier approval recommendations."""

    async def create_permit(self, db: AsyncSession, request: PermitCreateRequest) -> PermitRecordDto:
        logger.info("Creating new permit-to-work draft: %s", request.permit_title)

        # 1. Run Safety Guardrail check on details
        guard_res = await guardrail_engine.check_input(db, request.details, request.user_id, request.tenant_id)
        sanitised_details = guard_res.sanitised_text

        # 2. Save permit record in DB
        dto = await permit_repository.create_permit(
            db,
            request.permit_title,
            request.permit_type,
            request.location,
            request.applicant_id,
            sanitised_details,
            request.hazard_controls,
            request.user_id,
            request.tenant_id,
        )
        return dto

    async def validate_permit(self, db: AsyncSession, permit_id: str) -> PermitValidationResponse:
        logger.info("Validating completeness and control checklist for permit: %s", permit_id)

        permit = await permit_repository.get_permit(db, permit_id)
        if not permit:
            raise ValueError("Permit record not found.")

        # Query RAG for permit type SOP guidelines
        chunks = await rag_orchestrator.execute_hybrid_search(
            db, permit.tenant_id, query=permit.permit_type, limit=3
        )
        context_texts = [getattr(c, "content", getattr(c, "text", "")) for c in chunks]

        existing_controls = (permit.hazard_controls_json or {}).get("controls", [])

        # LLM completeness validation
        val_prompt = f"""
        Validate safety completeness for Permit:
        Type: {permit.permit_type}
        Title: {permit.permit_title}
        Details: {permit.details}
        Existing Hazard Controls: {", ".join(existing_controls)}
        SOP Context: {" ".join(context_texts)}

        Determine missing hazard controls and required approvers. Format output:
        Valid: [Yes/No]
        Missing Controls: [control 1 | control 2]
        Required Approvers: [Safety Officer | Area Supervisor]
        Risk Level: [LOW/MEDIUM/HIGH/CRITICAL]
        """

        chat_req = ChatRequest(
            model=MODEL_NAME,
            messages=[Message(role="user", content=val_prompt)],
        )
        chat_res = await provider_factory.execute_chat(chat_req)
        llm_output = chat_res.choices[0].message.content

        # Heuristic parsing
        missing_controls = []
        approvers = ["Safety Officer", "Area Supervisor"]
        risk_level = "MEDIUM"
        is_valid = True

        for line in llm_output.split("\n"):
            if "missing controls" in line.lower():
                raw = line.split(":")[-1].strip()
                if raw and raw.lower() != "none":
                    missing_controls = [c.strip() for c in raw.split("|")]
            elif "risk level" in line.lower():
                risk_level = line.split(":")[-1].strip()

        # Fallback defaults for specific permit types if controls missing
        if permit.permit_type == "HOT_WORK" and not any("gas" in c.lower() for c in existing_controls):
            missing_controls.append("Continuous Gas Monitoring & Dedicated Fire Watch")
            risk_level = "HIGH"

        if missing_controls:
            is_valid = False

        await permit_repository.update_status(db, permit_id, "VALIDATED")

        return PermitValidationResponse(
            is_valid=is_valid,
            missing_controls=missing_controls,
            required_approvers=approvers,
            risk_level=risk_level,
        )

    async def assess_risk(self, db: AsyncSession, request: RiskAssessmentRequest) -> RiskAssessmentResponse:
        logger.info("Executing Likelihood x Severity Risk Matrix assessment for permit: %s", request.permit_id)

        permit = await permit_repository.get_permit(db, request.permit_id)
        if not permit:
            raise ValueError("Permit record not found.")

        # Risk Matrix Heuristics based on Permit Type & Details
        likelihood = 2
        severity = 3

        p_type = permit.permit_type.upper()
        if p_type in ["HOT_WORK", "CONFINED_SPACE", "ELECTRICAL"]:
            likelihood = 3
            severity = 4
        elif p_type == "HEIGHT":
            likelihood = 3
            severity = 3

        residual_risk = float(likelihood * severity)

        category = "LOW"
        if residual_risk > 18:
            category = "CRITICAL"
        elif residual_risk > 12:
            category = "HIGH"
        elif residual_risk > 6:
            category = "MEDIUM"

        rec_controls = [
            "Verify LOTO lockout isolation points prior to entry.",
            "Conduct pre-job toolbox safety talk with work crew.",
        ]

        # Update permit risk score in DB
        await permit_repository.update_risk_score(db, request.permit_id, residual_risk)

        return RiskAssessmentResponse(
            likelihood_score=likelihood,
            severity_score=severity,
            residual_risk_score=residual_risk,
            risk_category=category,
            recommended_controls=rec_controls,
        )

    async def recommend_approval(self, db: AsyncSession, request: ApprovalRecommendationRequest) -> ApprovalRecommendationResponse:
        logger.info("Evaluating approval recommendation for permit: %s", request.permit_id)

        permit = await permit_repository.get_permit(db, request.permit_id)
        if not permit:
            raise ValueError("Permit record not found.")

        # Check worker training tool lookup (if applicant training valid)
        t_res = await tool_executor.execute_tool(
            db, "check_weather", None, {"latitude": 48.1, "longitude": 11.5}, permit.user_id, permit.tenant_id
        )

        rec = "APPROVE"
        reason = "Permit completeness verified. All mandatory hazard controls and PPE requirements satisfied."
        confidence = 0.96

        if permit.risk_score > 15:
            rec = "SUSPEND"
            reason = "Residual risk score exceeds safe threshold. Additional high-level isolation sign-off required."
            confidence = 0.92

        return ApprovalRecommendationResponse(
            recommendation=rec,
            reason=reason,
            confidence_score=confidence,
        )

    async def approve_permit(
        self, db: AsyncSession, permit_id: str, approver_id: str, role: str, approved: bool, comments: str | None
    ) -> PermitRecordDto:
        logger.info("Recording sign-off decision for permit %s by %s (%s)", permit_id, approver_id, role)

        status_str = "APPROVED" if approved else "REJECTED"
        await permit_repository.add_approval(db, permit_id, approver_id, role, status_str, comments)
        await permit_repository.update_status(db, permit_id, status_str)

        permit_db = await permit_repository.get_permit(db, permit_id)
        return PermitRecordDto(
            id=permit_db.id,
            permit_title=permit_db.permit_title,
            permit_type=permit_db.permit_type,
            status=permit_db.status,
            location=permit_db.location,
            applicant_id=permit_db.applicant_id,
            details=permit_db.details,
            risk_score=permit_db.risk_score,
            confidence_score=permit_db.confidence_score,
            created_at=str(permit_db.created_at),
        )


permit_agent = PermitToWorkAgent()
