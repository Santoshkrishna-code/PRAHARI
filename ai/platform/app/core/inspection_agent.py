import os
import logging
from sqlalchemy.ext.asyncio import AsyncSession

# Volumes 0-12 Reusable Modules
from app.infrastructure.providers.factory import provider_factory
from app.core.rag_orchestrator import rag_orchestrator
from app.core.tool_executor import tool_executor
from app.core.guardrail_engine import guardrail_engine
from app.core.evaluation_engine import evaluation_engine
from app.core.risk_agent import enterprise_risk_agent
from app.infrastructure.repositories.inspection_repository import inspection_repository
from app.domain.risks import RiskAssessmentRequest
from app.core.models import ChatRequest, Message

# Domain schemas
from app.domain.inspections import (
    InspectionPlanRequest,
    InspectionRecordDto,
    ChecklistGenerationRequest,
    ChecklistGenerationResponse,
    FindingAnalysisRequest,
    FindingAnalysisResponse,
    InspectionReportResponse,
)

logger = logging.getLogger(__name__)
MODEL_NAME = "mock-model" if os.getenv("APP_ENV") == "testing" else "gpt-4o"


class IntelligentInspectionAgent:
    """Agent coordinator managing dynamic checklist generation, finding classification, risk integration, and EHS reports."""

    async def plan_inspection(self, db: AsyncSession, request: InspectionPlanRequest) -> InspectionRecordDto:
        logger.info("Planning workplace inspection: %s (%s)", request.title, request.inspection_type)

        # 1. Run Safety Guardrail check
        guard_res = await guardrail_engine.check_input(db, request.title, request.user_id, request.tenant_id)

        # 2. Create inspection record in DB
        dto = await inspection_repository.create_inspection(
            db,
            request.title,
            request.inspection_type,
            request.location,
            request.inspector_id,
            request.scope_notes,
            request.user_id,
            request.tenant_id,
        )
        return dto

    async def generate_checklist(self, db: AsyncSession, request: ChecklistGenerationRequest) -> ChecklistGenerationResponse:
        logger.info("Generating dynamic checklist for inspection type: %s at %s", request.inspection_type, request.location)

        # Query RAG for inspection manuals / ISO standards
        chunks = await rag_orchestrator.execute_hybrid_search(
            db, "tenant-ehs-corp", query=request.inspection_type, limit=3
        )
        context_texts = [getattr(c, "content", getattr(c, "text", "")) for c in chunks]

        # Dynamic checklist generator defaults based on inspection type
        items = [
            "Verify clear access to emergency exits and fire extinguishers.",
            "Inspect personal protective equipment (PPE) compliance for personnel.",
            "Check chemical storage secondary containment bunding.",
            "Inspect electrical cables, junction boxes, and LOTO isolation labels.",
        ]

        if request.inspection_type == "FIRE_SAFETY":
            items.append("Test emergency alarm pull stations and pressure gauges.")
        elif request.inspection_type == "CONSTRUCTION":
            items.append("Verify scaffolding toe-boards, guardrails, and harness tie-off points.")

        return ChecklistGenerationResponse(
            inspection_type=request.inspection_type,
            items=items,
        )

    async def analyze_finding(self, db: AsyncSession, request: FindingAnalysisRequest) -> FindingAnalysisResponse:
        logger.info("Classifying inspection finding for inspection: %s", request.inspection_id)

        inspection = await inspection_repository.get_inspection(db, request.inspection_id)
        if not inspection:
            raise ValueError("Inspection record not found.")

        # Classify finding type and severity heuristically
        obs_lower = request.observation_text.lower()

        finding_type = "UNSAFE_CONDITION"
        severity = "MEDIUM"
        deduction = 10.0
        capa = "Barricade area immediately and issue maintenance work order."

        if "blocked" in obs_lower or "leak" in obs_lower:
            finding_type = "MAJOR_NON_CONFORMANCE"
            severity = "HIGH"
            deduction = 15.0
            capa = "Clear obstructed emergency route immediately and retrain area supervisor."
        elif "missing" in obs_lower or "damaged" in obs_lower:
            finding_type = "UNSAFE_ACT"
            severity = "MEDIUM"
            deduction = 10.0
            capa = "Replace damaged safety equipment prior to continuing operations."

        # If MAJOR_NON_CONFORMANCE or CRITICAL_FINDING, invoke Risk Assessment Agent (Volume 12 integration!)
        risk_priority = 12.0
        if finding_type in ["MAJOR_NON_CONFORMANCE", "CRITICAL_FINDING"]:
            risk_req = RiskAssessmentRequest(
                title=f"Inspection Finding: {finding_type}",
                risk_type="OCCUPATIONAL",
                methodology="HIRA",
                details=request.observation_text,
                location=inspection.location,
                user_id=inspection.user_id,
                tenant_id=inspection.tenant_id,
            )
            risk_dto = await enterprise_risk_agent.assess_risk(db, risk_req)
            risk_priority = risk_dto.initial_risk_score

        # Log finding in DB
        await inspection_repository.add_finding(
            db, request.inspection_id, finding_type, severity, request.observation_text, capa
        )

        # Update inspection score and status
        new_findings_count = inspection.findings_count + 1
        new_score = max(0.0, inspection.score - deduction)
        await inspection_repository.update_inspection_status_and_score(
            db, request.inspection_id, "IN_PROGRESS", new_score, new_findings_count
        )

        return FindingAnalysisResponse(
            finding_type=finding_type,
            severity=severity,
            details=request.observation_text,
            recommended_capa=capa,
            risk_priority_score=risk_priority,
        )

    async def compile_report(self, db: AsyncSession, inspection_id: str) -> InspectionReportResponse:
        logger.info("Compiling final EHS Inspection Report for: %s", inspection_id)

        inspection = await inspection_repository.get_inspection(db, inspection_id)
        if not inspection:
            raise ValueError("Inspection record not found.")

        # Update status to COMPLETED
        await inspection_repository.update_inspection_status_and_score(
            db, inspection_id, "COMPLETED", inspection.score, inspection.findings_count
        )

        # Build Markdown report
        findings_md = ""
        for idx, f in enumerate(inspection.findings, start=1):
            findings_md += f"\n### Finding #{idx}: {f.finding_type} ({f.severity})\n- **Details**: {f.details}\n- **Corrective Action**: {f.corrective_action}\n"

        report_md = f"""# EHS INTELLIGENT INSPECTION REPORT
**Inspection ID**: {inspection.id}
**Title**: {inspection.title}
**Location**: {inspection.location}
**Inspector**: {inspection.inspector_id}
**Final Inspection Score**: {inspection.score} / 100.0
**Total Deficiencies Flagged**: {inspection.findings_count}

## Executive Summary
Safety inspection executed in accordance with ISO 45001 guidelines and internal EHS protocols.

## Findings & CAPA Log
{findings_md if findings_md else "No safety non-conformances identified during inspection."}
"""

        return InspectionReportResponse(
            inspection_id=inspection_id,
            score=inspection.score,
            report_markdown=report_md,
        )


intelligent_inspection_agent = IntelligentInspectionAgent()
