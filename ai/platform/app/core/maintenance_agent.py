import os
import logging
import datetime
from sqlalchemy.ext.asyncio import AsyncSession

# Volumes 0-16 Reusable Modules
from app.infrastructure.providers.factory import provider_factory
from app.core.rag_orchestrator import rag_orchestrator
from app.core.tool_executor import tool_executor
from app.core.guardrail_engine import guardrail_engine
from app.core.evaluation_engine import evaluation_engine
from app.core.risk_agent import enterprise_risk_agent
from app.infrastructure.repositories.asset_repository import asset_repository
from app.infrastructure.repositories.maintenance_repository import maintenance_repository
from app.domain.risks import RiskAssessmentRequest
from app.core.models import ChatRequest, Message

# Domain schemas
from app.domain.maintenance import (
    MaintenancePredictRequest,
    MaintenancePredictionDto,
    RulEstimateRequest,
    RulEstimateResponse,
    MaintenancePlanRequest,
    MaintenancePlanResponse,
    ReliabilityDashboardResponse,
)

logger = logging.getLogger(__name__)
MODEL_NAME = "mock-model" if os.getenv("APP_ENV") == "testing" else "gpt-4o"


class PredictiveMaintenanceAgent:
    """Agent coordinator managing failure prediction models, RUL estimations, condition-based work orders, and reliability metrics."""

    async def predict_failure(self, db: AsyncSession, request: MaintenancePredictRequest) -> MaintenancePredictionDto:
        logger.info("Predicting failure trajectory for asset ID: %s", request.asset_id)

        # 1. Fetch Asset Record (Volume 15 integration)
        asset = await asset_repository.get_asset(db, request.asset_id)

        # 2. Predictive Failure Scoring
        vib = request.vibration_trend_mm_s
        hours = request.operating_hours

        prob = min(0.95, round((vib / 6.0) * 0.65 + (hours / 5000.0) * 0.35, 2))
        rul = max(24.0, round((1.0 - prob) * 1200.0, 1))

        strategy = "CONDITION_BASED" if prob > 0.5 else "PREVENTIVE"
        rec_date = (datetime.datetime.utcnow() + datetime.timedelta(days=int(rul / 24.0))).strftime("%Y-%m-%d")

        spare_parts = ["Vibration Damper Gasket", "Heavy-Duty Bearing Assembly", "High-Temp Lubricant Seal"]

        # 3. Save prediction in DB
        dto = await maintenance_repository.create_prediction(
            db,
            request.asset_id,
            prob,
            rul,
            rec_date,
            strategy,
            spare_parts,
            request.user_id,
            request.tenant_id,
        )
        return dto

    async def estimate_rul(self, db: AsyncSession, request: RulEstimateRequest) -> RulEstimateResponse:
        logger.info("Estimating Remaining Useful Life (RUL) for asset ID: %s", request.asset_id)

        asset = await asset_repository.get_asset(db, request.asset_id)

        rul_h = 480.0
        trend = "ACCELERATING"
        if asset:
            rul_h = max(48.0, round(asset.health_score * 8.0, 1))
            if asset.health_score > 75.0:
                trend = "STABLE"

        rul_d = round(rul_h / 24.0, 1)

        return RulEstimateResponse(
            asset_id=request.asset_id,
            remaining_useful_life_hours=rul_h,
            remaining_useful_life_days=rul_d,
            confidence_score=0.91,
            degradation_trend=trend,
        )

    async def create_maintenance_plan(self, db: AsyncSession, request: MaintenancePlanRequest) -> MaintenancePlanResponse:
        logger.info("Creating predictive maintenance work order plan for asset ID: %s", request.asset_id)

        asset = await asset_repository.get_asset(db, request.asset_id)

        priority = "HIGH"
        downtime = 6.0
        if asset and asset.health_score < 40.0:
            priority = "CRITICAL"
            downtime = 12.0

        spare_parts = ["Mechanical Seal Ring", "Overhaul Shaft Bearing", "Coupling Flange Kit"]

        res = await maintenance_repository.create_work_order(
            db, request.asset_id, request.work_order_title, priority, downtime, spare_parts
        )
        return res

    async def get_dashboard_metrics(self, db: AsyncSession, tenant_id: str) -> ReliabilityDashboardResponse:
        logger.info("Retrieving plant-wide reliability dashboard metrics for tenant: %s", tenant_id)
        return await maintenance_repository.get_reliability_metrics(db, tenant_id)


predictive_maintenance_agent = PredictiveMaintenanceAgent()
