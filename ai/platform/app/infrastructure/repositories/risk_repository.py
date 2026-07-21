import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.risk_entities import Base, RiskAssessmentRecordDb, RiskRegisterDb
from app.domain.risks import RiskAssessmentDto, RiskRegisterEntryDto


class RiskRepository:
    """Repository handling DB operations for Risk Assessments and System-wide Risk Registers."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_assessment(
        self,
        session: AsyncSession,
        title: str,
        risk_type: str,
        methodology: str,
        initial_score: float,
        residual_score: float,
        category: str,
        hazards: list[str],
        controls: dict,
        user_id: str,
        tenant_id: str,
    ) -> RiskAssessmentDto:
        r_id = str(uuid.uuid4())
        db_risk = RiskAssessmentRecordDb(
            id=r_id,
            title=title,
            risk_type=risk_type,
            methodology=methodology,
            initial_risk_score=initial_score,
            residual_risk_score=residual_score,
            risk_category=category,
            hazards_json={"hazards": hazards},
            recommended_controls_json=controls,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_risk)
        await session.commit()

        return RiskAssessmentDto(
            id=r_id,
            title=title,
            risk_type=risk_type,
            methodology=methodology,
            initial_risk_score=initial_score,
            residual_risk_score=residual_score,
            risk_category=category,
            created_at=str(db_risk.created_at),
        )

    async def get_assessment(self, session: AsyncSession, id: str) -> RiskAssessmentRecordDb | None:
        stmt = (
            select(RiskAssessmentRecordDb)
            .where(RiskAssessmentRecordDb.id == id)
            .options(selectinload(RiskAssessmentRecordDb.register_entries))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_assessments_list(self, session: AsyncSession, tenant_id: str) -> list[RiskAssessmentDto]:
        stmt = (
            select(RiskAssessmentRecordDb)
            .where(RiskAssessmentRecordDb.tenant_id == tenant_id)
            .order_by(RiskAssessmentRecordDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            RiskAssessmentDto(
                id=r.id,
                title=r.title,
                risk_type=r.risk_type,
                methodology=r.methodology,
                initial_risk_score=r.initial_risk_score,
                residual_risk_score=r.residual_risk_score,
                risk_category=r.risk_category,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def add_register_entry(
        self,
        session: AsyncSession,
        assessment_id: str,
        hazard_name: str,
        severity_score: int,
        likelihood_score: int,
        trend_status: str,
    ) -> RiskRegisterEntryDto:
        reg_id = str(uuid.uuid4())
        db_reg = RiskRegisterDb(
            id=reg_id,
            assessment_id=assessment_id,
            hazard_name=hazard_name,
            severity_score=severity_score,
            likelihood_score=likelihood_score,
            trend_status=trend_status,
        )
        session.add(db_reg)
        await session.commit()

        return RiskRegisterEntryDto(
            id=reg_id,
            assessment_id=assessment_id,
            hazard_name=hazard_name,
            severity_score=severity_score,
            likelihood_score=likelihood_score,
            trend_status=trend_status,
        )

    async def get_register_entries(self, session: AsyncSession) -> list[RiskRegisterEntryDto]:
        stmt = select(RiskRegisterDb).order_by(RiskRegisterDb.created_at.desc())
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            RiskRegisterEntryDto(
                id=r.id,
                assessment_id=r.assessment_id,
                hazard_name=r.hazard_name,
                severity_score=r.severity_score,
                likelihood_score=r.likelihood_score,
                trend_status=r.trend_status,
            )
            for r in rows
        ]

    async def update_residual_risk(self, session: AsyncSession, id: str, residual_score: float, controls: dict) -> None:
        stmt = (
            update(RiskAssessmentRecordDb)
            .where(RiskAssessmentRecordDb.id == id)
            .values(residual_risk_score=residual_score, recommended_controls_json=controls)
        )
        await session.execute(stmt)
        await session.commit()


risk_repository = RiskRepository()
