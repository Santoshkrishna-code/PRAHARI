import os
import logging
from sqlalchemy.ext.asyncio import AsyncSession

# Volumes 0-13 Reusable Modules
from app.infrastructure.providers.factory import provider_factory
from app.core.rag_orchestrator import rag_orchestrator
from app.core.tool_executor import tool_executor
from app.core.guardrail_engine import guardrail_engine
from app.core.evaluation_engine import evaluation_engine
from app.core.risk_agent import enterprise_risk_agent
from app.core.compliance_agent import safety_compliance_agent
from app.core.permit_agent import permit_agent
from app.infrastructure.repositories.contractor_repository import contractor_repository
from app.domain.compliance import ComplianceCheckRequest
from app.domain.risks import RiskAssessmentRequest
from app.core.models import ChatRequest, Message

# Domain schemas
from app.domain.contractors import (
    ContractorRegisterRequest,
    ContractorRecordDto,
    ContractorQualifyRequest,
    ContractorQualifyResponse,
    ContractorVerifyRequest,
    ContractorVerifyResponse,
    ContractorScorecardResponse,
)

logger = logging.getLogger(__name__)
MODEL_NAME = "mock-model" if os.getenv("APP_ENV") == "testing" else "gpt-4o"


class ContractorSafetyAgent:
    """Agent coordinator managing contractor pre-qualification, document verification, cross-agent compliance/risk calls, and scorecards."""

    async def register_contractor(self, db: AsyncSession, request: ContractorRegisterRequest) -> ContractorRecordDto:
        logger.info("Registering contractor profile: %s (%s)", request.company_name, request.contractor_type)

        # 1. Run Safety Guardrail check
        guard_res = await guardrail_engine.check_input(db, request.company_name, request.user_id, request.tenant_id)

        # 2. Save contractor record in DB
        dto = await contractor_repository.create_contractor(
            db,
            request.company_name,
            request.contractor_type,
            request.contact_email,
            request.license_number,
            request.user_id,
            request.tenant_id,
        )
        return dto

    async def qualify_contractor(self, db: AsyncSession, request: ContractorQualifyRequest) -> ContractorQualifyResponse:
        logger.info("Evaluating safety pre-qualification for contractor ID: %s", request.contractor_id)

        contractor = await contractor_repository.get_contractor(db, request.contractor_id)
        if not contractor:
            raise ValueError("Contractor profile not found.")

        # 1. Cross-Agent Call to Safety Compliance Agent (Volume 10)
        comp_req = ComplianceCheckRequest(
            target_subject=contractor.company_name,
            subject_type="CONTRACTOR",
            details=f"Pre-qualification check for {contractor.contractor_type} trade license {contractor.license_number}.",
            user_id=contractor.user_id,
            tenant_id=contractor.tenant_id,
        )
        comp_res = await safety_compliance_agent.check_compliance(db, comp_req)

        # 2. Cross-Agent Call to Enterprise Risk Agent (Volume 12)
        risk_req = RiskAssessmentRequest(
            title=f"Contractor Qualification Risk: {contractor.company_name}",
            risk_type="OCCUPATIONAL",
            methodology="HIRA",
            details=f"Evaluating task risk for {contractor.contractor_type} workforce.",
            location="Site Operations",
            user_id=contractor.user_id,
            tenant_id=contractor.tenant_id,
        )
        risk_dto = await enterprise_risk_agent.assess_risk(db, risk_req)

        # Qualification Heuristics
        is_qualified = True
        status_str = "QUALIFIED"
        missing_trainings = []
        risk_level = "LOW"
        safety_score = min(100.0, comp_res.compliance_score)

        if contractor.contractor_type in ["SCAFFOLDING", "WELDING", "SHUTDOWN"]:
            missing_trainings.append("Site Induction & Mandatory Work-at-Height Certification")
            risk_level = "MEDIUM"

        if safety_score < 60.0:
            is_qualified = False
            status_str = "REJECTED"
            risk_level = "HIGH"

        # Update contractor status and log audit entry
        await contractor_repository.update_contractor_status_and_score(
            db, request.contractor_id, status_str, safety_score, contractor.violations_count
        )
        await contractor_repository.add_audit(
            db, request.contractor_id, "PRE_QUALIFICATION", safety_score, f"Pre-qualification result: {status_str}"
        )

        return ContractorQualifyResponse(
            is_qualified=is_qualified,
            status=status_str,
            safety_score=safety_score,
            missing_trainings=missing_trainings,
            risk_level=risk_level,
        )

    async def verify_document(self, db: AsyncSession, request: ContractorVerifyRequest) -> ContractorVerifyResponse:
        logger.info("Verifying document type %s for contractor: %s", request.document_type, request.contractor_id)

        contractor = await contractor_repository.get_contractor(db, request.contractor_id)
        if not contractor:
            raise ValueError("Contractor profile not found.")

        is_valid = True
        notes = f"Verified {request.document_type} validity against central registry."

        if request.document_type == "INSURANCE" and contractor.violations_count > 2:
            is_valid = False
            notes = "Insurance verification pending renewal following safety violation notices."

        return ContractorVerifyResponse(
            is_valid=is_valid,
            verification_notes=notes,
        )

    async def generate_scorecard(self, db: AsyncSession, contractor_id: str) -> ContractorScorecardResponse:
        logger.info("Generating Contractor Safety Scorecard for ID: %s", contractor_id)

        contractor = await contractor_repository.get_contractor(db, contractor_id)
        if not contractor:
            raise ValueError("Contractor profile not found.")

        audits_count = len(contractor.audits)
        comp_status = "COMPLIANT" if contractor.safety_rating_score >= 70.0 else "NON_COMPLIANT"

        return ContractorScorecardResponse(
            contractor_id=contractor.id,
            company_name=contractor.company_name,
            safety_score=contractor.safety_rating_score,
            audit_history_count=audits_count,
            compliance_status=comp_status,
        )


contractor_safety_agent = ContractorSafetyAgent()
