import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.compliance_entities import Base, ComplianceAssessmentDb, ComplianceViolationDb
from app.domain.compliance import ComplianceAssessmentDto, ViolationDto


class ComplianceRepository:
    """Repository handling DB operations for Compliance Assessments and violation tracking data."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_assessment(
        self,
        session: AsyncSession,
        target_subject: str,
        subject_type: str,
        details: str,
        score: float,
        status: str,
        user_id: str,
        tenant_id: str,
    ) -> ComplianceAssessmentDto:
        ass_id = str(uuid.uuid4())
        db_ass = ComplianceAssessmentDb(
            id=ass_id,
            target_subject=target_subject,
            subject_type=subject_type,
            compliance_score=score,
            status=status,
            details=details,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_ass)
        await session.commit()

        return ComplianceAssessmentDto(
            id=ass_id,
            target_subject=target_subject,
            subject_type=subject_type,
            compliance_score=score,
            status=status,
            details=details,
            created_at=str(db_ass.created_at),
        )

    async def get_assessment(self, session: AsyncSession, id: str) -> ComplianceAssessmentDb | None:
        stmt = (
            select(ComplianceAssessmentDb)
            .where(ComplianceAssessmentDb.id == id)
            .options(selectinload(ComplianceAssessmentDb.violations))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_assessments_list(self, session: AsyncSession, tenant_id: str) -> list[ComplianceAssessmentDto]:
        stmt = (
            select(ComplianceAssessmentDb)
            .where(ComplianceAssessmentDb.tenant_id == tenant_id)
            .order_by(ComplianceAssessmentDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            ComplianceAssessmentDto(
                id=r.id,
                target_subject=r.target_subject,
                subject_type=r.subject_type,
                compliance_score=r.compliance_score,
                status=r.status,
                details=r.details,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def add_violation(
        self, session: AsyncSession, assessment_id: str, regulation_violated: str, severity: str, details: str
    ) -> ViolationDto:
        v_id = str(uuid.uuid4())
        db_viol = ComplianceViolationDb(
            id=v_id,
            assessment_id=assessment_id,
            regulation_violated=regulation_violated,
            severity=severity,
            details=details,
        )
        session.add(db_viol)
        await session.commit()
        return ViolationDto(
            id=v_id,
            assessment_id=assessment_id,
            regulation_violated=regulation_violated,
            severity=severity,
            details=details,
        )

    async def get_violations_by_assessment(self, session: AsyncSession, assessment_id: str) -> list[ViolationDto]:
        stmt = select(ComplianceViolationDb).where(ComplianceViolationDb.assessment_id == assessment_id)
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            ViolationDto(
                id=r.id,
                assessment_id=r.assessment_id,
                regulation_violated=r.regulation_violated,
                severity=r.severity,
                details=r.details,
            )
            for r in rows
        ]

    async def update_capa_plan(self, session: AsyncSession, id: str, capa_plan: dict) -> None:
        stmt = (
            update(ComplianceAssessmentDb)
            .where(ComplianceAssessmentDb.id == id)
            .values(capa_plan_json=capa_plan)
        )
        await session.execute(stmt)
        await session.commit()

    async def update_score_and_status(self, session: AsyncSession, id: str, score: float, status: str) -> None:
        stmt = (
            update(ComplianceAssessmentDb)
            .where(ComplianceAssessmentDb.id == id)
            .values(compliance_score=score, status=status)
        )
        await session.execute(stmt)
        await session.commit()


compliance_repository = ComplianceRepository()
