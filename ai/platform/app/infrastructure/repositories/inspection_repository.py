import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.inspection_entities import Base, InspectionRecordDb, InspectionFindingDb
from app.domain.inspections import InspectionRecordDto


class InspectionRepository:
    """Repository handling DB operations for Inspection records and finding entries."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_inspection(
        self,
        session: AsyncSession,
        title: str,
        inspection_type: str,
        location: str,
        inspector_id: str,
        scope_notes: str | None,
        user_id: str,
        tenant_id: str,
    ) -> InspectionRecordDto:
        i_id = str(uuid.uuid4())
        db_insp = InspectionRecordDb(
            id=i_id,
            title=title,
            inspection_type=inspection_type,
            status="PLANNED",
            location=location,
            inspector_id=inspector_id,
            score=100.0,
            findings_count=0,
            summary_notes=scope_notes,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_insp)
        await session.commit()

        return InspectionRecordDto(
            id=i_id,
            title=title,
            inspection_type=inspection_type,
            status="PLANNED",
            location=location,
            inspector_id=inspector_id,
            score=100.0,
            findings_count=0,
            summary_notes=scope_notes,
            created_at=str(db_insp.created_at),
        )

    async def get_inspection(self, session: AsyncSession, id: str) -> InspectionRecordDb | None:
        stmt = (
            select(InspectionRecordDb)
            .where(InspectionRecordDb.id == id)
            .options(selectinload(InspectionRecordDb.findings))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_inspections_list(self, session: AsyncSession, tenant_id: str) -> list[InspectionRecordDto]:
        stmt = (
            select(InspectionRecordDb)
            .where(InspectionRecordDb.tenant_id == tenant_id)
            .order_by(InspectionRecordDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            InspectionRecordDto(
                id=r.id,
                title=r.title,
                inspection_type=r.inspection_type,
                status=r.status,
                location=r.location,
                inspector_id=r.inspector_id,
                score=r.score,
                findings_count=r.findings_count,
                summary_notes=r.summary_notes,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def add_finding(
        self,
        session: AsyncSession,
        inspection_id: str,
        finding_type: str,
        severity: str,
        details: str,
        corrective_action: str | None,
    ) -> None:
        db_f = InspectionFindingDb(
            id=str(uuid.uuid4()),
            inspection_id=inspection_id,
            finding_type=finding_type,
            severity=severity,
            details=details,
            corrective_action=corrective_action,
        )
        session.add(db_f)
        await session.commit()

    async def update_inspection_status_and_score(
        self, session: AsyncSession, id: str, status: str, score: float, findings_count: int
    ) -> None:
        stmt = (
            update(InspectionRecordDb)
            .where(InspectionRecordDb.id == id)
            .values(status=status, score=score, findings_count=findings_count)
        )
        await session.execute(stmt)
        await session.commit()


inspection_repository = InspectionRepository()
