import os
import logging
from sqlalchemy.ext.asyncio import AsyncSession

# Volumes 0-11 Reusable Modules
from app.infrastructure.providers.factory import provider_factory
from app.core.rag_orchestrator import rag_orchestrator
from app.core.tool_executor import tool_executor
from app.core.guardrail_engine import guardrail_engine
from app.core.evaluation_engine import evaluation_engine
from app.infrastructure.repositories.risk_repository import risk_repository
from app.core.models import ChatRequest, Message

# Domain schemas
from app.domain.risks import (
    RiskAssessmentRequest,
    RiskAssessmentDto,
    RiskPredictRequest,
    RiskPredictResponse,
    MitigationPlanRequest,
    MitigationPlanResponse,
)

logger = logging.getLogger(__name__)
MODEL_NAME = "mock-model" if os.getenv("APP_ENV") == "testing" else "gpt-4o"


class EnterpriseRiskAgent:
    """Agent orchestrator managing hazard identification, ISO 31000 risk scoring, predictive trend analytics, and control mitigations."""

    async def assess_risk(self, db: AsyncSession, request: RiskAssessmentRequest) -> RiskAssessmentDto:
        logger.info("Executing risk assessment (%s) for: %s", request.methodology, request.title)

        # 1. Run Safety Guardrail check
        guard_res = await guardrail_engine.check_input(db, request.details, request.user_id, request.tenant_id)
        sanitised_details = guard_res.sanitised_text

        # 2. Retrieve ISO 31000 / OSHA standards via RAG
        chunks = await rag_orchestrator.execute_hybrid_search(
            db, request.tenant_id, query=request.risk_type, limit=3
        )
        context_texts = [getattr(c, "content", getattr(c, "text", "")) for c in chunks]

        # 3. Calculate Likelihood & Severity Scores
        likelihood = 3
        severity = 3
        r_type = request.risk_type.upper()

        if r_type in ["FIRE", "EXPLOSION", "CHEMICAL"]:
            likelihood = 4
            severity = 5
        elif r_type in ["ELECTRICAL", "MECHANICAL"]:
            likelihood = 3
            severity = 4

        initial_score = float(likelihood * severity)
        residual_score = round(initial_score * 0.35, 2)

        category = "LOW"
        if initial_score > 18:
            category = "CRITICAL"
        elif initial_score > 12:
            category = "HIGH"
        elif initial_score > 6:
            category = "MEDIUM"

        controls = {
            "engineering": ["Install automated emergency shut-off valves", "Enclose high-pressure lines"],
            "administrative": ["Mandatory permit sign-off and JSA briefing"],
            "ppe": ["Thermal resistant gloves", "Face shield visor"],
        }

        # 4. Save Risk Assessment Record in DB
        dto = await risk_repository.create_assessment(
            db,
            request.title,
            request.risk_type,
            request.methodology,
            initial_score,
            residual_score,
            category,
            [f"Primary hazard identified: {request.risk_type}"],
            controls,
            request.user_id,
            request.tenant_id,
        )

        # 5. Log entry in System-wide Risk Register
        await risk_repository.add_register_entry(
            db, dto.id, request.title, severity, likelihood, "STABLE"
        )

        return dto

    async def predict_risk(self, db: AsyncSession, request: RiskPredictRequest) -> RiskPredictResponse:
        logger.info("Executing predictive risk forecasting for: %s", request.subject)

        # Predictive analytics heuristic
        score = 8.0
        trend = "STABLE"
        summary = "Operational risk remains within normal statistical bounds."

        incidents = request.historical_incidents_count
        weather = (request.weather_condition or "").upper()

        if incidents > 3 or "STORM" in weather or "RAIN" in weather:
            score = 16.5
            trend = "INCREASING"
            summary = "Elevated risk predicted due to recent near-miss clusters and unfavorable weather parameters."
        elif incidents == 0:
            score = 4.2
            trend = "DECREASING"
            summary = "Favorable risk trajectory following recent preventive maintenance inspections."

        return RiskPredictResponse(
            predicted_risk_score=score,
            trend_status=trend,
            forecast_summary=summary,
            risk_confidence=0.92,
        )

    async def generate_mitigation_plan(self, db: AsyncSession, request: MitigationPlanRequest) -> MitigationPlanResponse:
        logger.info("Generating risk mitigation plan for assessment ID: %s", request.assessment_id)

        assessment = await risk_repository.get_assessment(db, request.assessment_id)
        if not assessment:
            raise ValueError("Risk assessment record not found.")

        controls_data = assessment.recommended_controls_json or {}
        eng = controls_data.get("engineering", ["Install isolation barriers", "Fit gas leak detectors"])
        admin = controls_data.get("administrative", ["Daily safety toolbox talks", "Permit-to-work sign-off"])
        ppe = controls_data.get("ppe", ["Standard hard hat", "Safety glasses", "Steel-toe boots"])

        residual = round(assessment.initial_risk_score * 0.3, 2)
        await risk_repository.update_residual_risk(db, request.assessment_id, residual, controls_data)

        return MitigationPlanResponse(
            engineering_controls=eng,
            administrative_controls=admin,
            ppe_recommendations=ppe,
            residual_risk=residual,
        )


enterprise_risk_agent = EnterpriseRiskAgent()
