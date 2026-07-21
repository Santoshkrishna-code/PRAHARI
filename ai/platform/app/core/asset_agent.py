import os
import logging
from sqlalchemy.ext.asyncio import AsyncSession

# Volumes 0-14 Reusable Modules
from app.infrastructure.providers.factory import provider_factory
from app.core.rag_orchestrator import rag_orchestrator
from app.core.tool_executor import tool_executor
from app.core.guardrail_engine import guardrail_engine
from app.core.evaluation_engine import evaluation_engine
from app.core.risk_agent import enterprise_risk_agent
from app.core.compliance_agent import safety_compliance_agent
from app.core.inspection_agent import intelligent_inspection_agent
from app.infrastructure.repositories.asset_repository import asset_repository
from app.domain.inspections import ChecklistGenerationRequest, ChecklistGenerationResponse
from app.domain.compliance import ComplianceCheckRequest
from app.domain.risks import RiskAssessmentRequest
from app.core.models import ChatRequest, Message

# Domain schemas
from app.domain.assets import (
    AssetHealthCheckRequest,
    AssetRecordDto,
    AssetShutdownRequest,
    AssetShutdownResponse,
    AssetScorecardResponse,
)

logger = logging.getLogger(__name__)
MODEL_NAME = "mock-model" if os.getenv("APP_ENV") == "testing" else "gpt-4o"


class AssetEquipmentSafetyAgent:
    """Agent coordinator managing sensor health metrics, calibration checks, shutdown decisions, and cross-agent safety integrations."""

    async def evaluate_health(self, db: AsyncSession, request: AssetHealthCheckRequest) -> AssetRecordDto:
        logger.info("Evaluating operational health for asset: %s (%s)", request.asset_name, request.asset_category)

        # 1. Run Safety Guardrail check
        guard_res = await guardrail_engine.check_input(db, request.asset_name, request.user_id, request.tenant_id)

        # 2. Sensor Health Score Calculation
        base_score = 100.0
        if request.vibration_mm_s > 4.5:
            base_score -= 25.0
        if request.temperature_c > 95.0:
            base_score -= 20.0
        if request.pressure_bar > 12.0:
            base_score -= 20.0

        health_score = max(0.0, base_score)
        status_str = "OPERATIONAL"
        if health_score < 60.0:
            status_str = "MAINTENANCE_REQUIRED"

        # 3. Save asset record in DB
        dto = await asset_repository.create_asset(
            db,
            request.asset_name,
            request.asset_category,
            request.location,
            request.operating_hours,
            health_score,
            request.user_id,
            request.tenant_id,
            status=status_str,
        )

        # 4. Log health check entry
        await asset_repository.add_log(
            db, dto.id, "HEALTH_CHECK", health_score, f"Health evaluation score {health_score}/100."
        )

        return dto

    async def trigger_asset_inspection(self, db: AsyncSession, asset_id: str) -> ChecklistGenerationResponse:
        logger.info("Triggering inspection checklist for asset ID: %s", asset_id)

        asset = await asset_repository.get_asset(db, asset_id)
        if not asset:
            raise ValueError("Asset record not found.")

        # Cross-Agent Call to Intelligent Inspection Agent (Volume 13)
        chk_req = ChecklistGenerationRequest(
            inspection_type="EQUIPMENT",
            location=asset.location,
        )
        chk_res = await intelligent_inspection_agent.generate_checklist(db, chk_req)
        return chk_res

    async def recommend_shutdown(self, db: AsyncSession, request: AssetShutdownRequest) -> AssetShutdownResponse:
        logger.info("Evaluating emergency shutdown request for asset ID: %s", request.asset_id)

        asset = await asset_repository.get_asset(db, request.asset_id)
        if not asset:
            raise ValueError("Asset record not found.")

        # Update status to SHUTDOWN in DB
        await asset_repository.update_asset_status_and_health(db, request.asset_id, "SHUTDOWN", 0.0)
        await asset_repository.add_log(db, request.asset_id, "SHUTDOWN_NOTICE", 0.0, request.reason)

        return AssetShutdownResponse(
            asset_id=asset.id,
            status="SHUTDOWN",
            action_taken=f"Emergency shutdown triggered: {request.reason}",
            safety_lockout_issued=True,
        )

    async def generate_scorecard(self, db: AsyncSession, asset_id: str) -> AssetScorecardResponse:
        logger.info("Generating Asset Integrity Scorecard for ID: %s", asset_id)

        asset = await asset_repository.get_asset(db, asset_id)
        if not asset:
            raise ValueError("Asset record not found.")

        # Cross-Agent Call to Safety Compliance Agent (Volume 10)
        comp_req = ComplianceCheckRequest(
            target_subject=asset.asset_name,
            subject_type="PPE",
            details=f"ISO 55001 asset integrity check for {asset.asset_category}.",
            user_id=asset.user_id,
            tenant_id=asset.tenant_id,
        )
        comp_res = await safety_compliance_agent.check_compliance(db, comp_req)

        return AssetScorecardResponse(
            asset_id=asset.id,
            asset_name=asset.asset_name,
            health_score=asset.health_score,
            operational_status=asset.status,
            calibration_valid=(comp_res.compliance_score >= 70.0),
        )


asset_equipment_safety_agent = AssetEquipmentSafetyAgent()
