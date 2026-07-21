import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.permit_entities import Base, PermitRecordDb, PermitApprovalDb
from app.domain.permits import PermitRecordDto


class PermitRepository:
    """Repository handling DB operations for Permit-to-Work records and sign-off approvals."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_permit(
        self,
        session: AsyncSession,
        permit_title: str,
        permit_type: str,
        location: str,
        applicant_id: str,
        details: str,
        hazard_controls: list[str],
        user_id: str,
        tenant_id: str,
    ) -> PermitRecordDto:
        p_id = str(uuid.uuid4())
        db_permit = PermitRecordDb(
            id=p_id,
            permit_title=permit_title,
            permit_type=permit_type,
            status="DRAFT",
            location=location,
            applicant_id=applicant_id,
            details=details,
            hazard_controls_json={"controls": hazard_controls},
            risk_score=5.0,
            confidence_score=0.95,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_permit)
        await session.commit()

        return PermitRecordDto(
            id=p_id,
            permit_title=permit_title,
            permit_type=permit_type,
            status="DRAFT",
            location=location,
            applicant_id=applicant_id,
            details=details,
            risk_score=5.0,
            confidence_score=0.95,
            created_at=str(db_permit.created_at),
        )

    async def get_permit(self, session: AsyncSession, id: str) -> PermitRecordDb | None:
        stmt = (
            select(PermitRecordDb)
            .where(PermitRecordDb.id == id)
            .options(selectinload(PermitRecordDb.approvals))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_permits_list(self, session: AsyncSession, tenant_id: str) -> list[PermitRecordDto]:
        stmt = (
            select(PermitRecordDb)
            .where(PermitRecordDb.tenant_id == tenant_id)
            .order_by(PermitRecordDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            PermitRecordDto(
                id=r.id,
                permit_title=r.permit_title,
                permit_type=r.permit_type,
                status=r.status,
                location=r.location,
                applicant_id=r.applicant_id,
                details=r.details,
                risk_score=r.risk_score,
                confidence_score=r.confidence_score,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def add_approval(
        self, session: AsyncSession, permit_id: str, approver_id: str, role: str, status: str, comments: str | None
    ) -> None:
        db_app = PermitApprovalDb(
            id=str(uuid.uuid4()),
            permit_id=permit_id,
            approver_id=approver_id,
            role=role,
            status=status,
            comments=comments,
        )
        session.add(db_app)
        await session.commit()

    async def update_status(self, session: AsyncSession, id: str, status: str) -> None:
        stmt = (
            update(PermitRecordDb)
            .where(PermitRecordDb.id == id)
            .values(status=status)
        )
        await session.execute(stmt)
        await session.commit()

    async def update_risk_score(self, session: AsyncSession, id: str, risk_score: float) -> None:
        stmt = (
            update(PermitRecordDb)
            .where(PermitRecordDb.id == id)
            .values(risk_score=risk_score)
        )
        await session.execute(stmt)
        await session.commit()

    async def update_hazard_controls(self, session: AsyncSession, id: str, controls: list[str]) -> None:
        stmt = (
            update(PermitRecordDb)
            .where(PermitRecordDb.id == id)
            .values(hazard_controls_json={"controls": controls})
        )
        await session.execute(stmt)
        await session.commit()


permit_repository = PermitRepository()
