import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.incident_entities import Base, IncidentInvestigationDb, IncidentEvidenceDb
from app.domain.incidents import IncidentInvestigationDto


class IncidentRepository:
    """Repository handling DB operations for Incident Investigations and evidence collection data."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_investigation(
        self,
        session: AsyncSession,
        title: str,
        details: str,
        classification: str,
        location: str,
        user_id: str,
        tenant_id: str,
    ) -> IncidentInvestigationDto:
        # Heuristic severity mapper based on classification tags
        sev = "LOW"
        c_upper = classification.upper()
        if "FATALITY" in c_upper or "EXPLOSION" in c_upper or "GAS LEAK" in c_upper:
            sev = "CRITICAL"
        elif "FIRE" in c_upper or "CHEMICAL SPILL" in c_upper or "LOST TIME INJURY" in c_upper:
            sev = "HIGH"
        elif "MEDICAL" in c_upper or "FIRST AID" in c_upper:
            sev = "MEDIUM"

        inv_id = str(uuid.uuid4())
        db_inv = IncidentInvestigationDb(
            id=inv_id,
            title=title,
            severity=sev,
            status="INTAKE",
            details=details,
            confidence_score=0.95,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_inv)
        await session.commit()

        return IncidentInvestigationDto(
            id=inv_id,
            title=title,
            severity=sev,
            status="INTAKE",
            details=details,
            confidence_score=0.95,
            reported_at=str(db_inv.reported_at),
        )

    async def get_investigation(self, session: AsyncSession, id: str) -> IncidentInvestigationDb | None:
        stmt = select(IncidentInvestigationDb).where(IncidentInvestigationDb.id == id).options(selectinload(IncidentInvestigationDb.evidence))
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_investigations_list(self, session: AsyncSession, tenant_id: str) -> list[IncidentInvestigationDto]:
        stmt = select(IncidentInvestigationDb).where(IncidentInvestigationDb.tenant_id == tenant_id).order_by(IncidentInvestigationDb.reported_at.desc())
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            IncidentInvestigationDto(
                id=r.id,
                title=r.title,
                severity=r.severity,
                status=r.status,
                details=r.details,
                confidence_score=r.confidence_score,
                reported_at=str(r.reported_at),
            )
            for r in rows
        ]

    async def add_evidence(
        self, session: AsyncSession, investigation_id: str, evidence_type: str, source_path: str, findings: str
    ) -> None:
        db_ev = IncidentEvidenceDb(
            id=str(uuid.uuid4()),
            investigation_id=investigation_id,
            evidence_type=evidence_type,
            source_path=source_path,
            findings=findings,
        )
        session.add(db_ev)
        await session.commit()

    async def update_root_causes(self, session: AsyncSession, id: str, root_causes: dict) -> None:
        stmt = (
            update(IncidentInvestigationDb)
            .where(IncidentInvestigationDb.id == id)
            .values(root_causes_json=root_causes)
        )
        await session.execute(stmt)
        await session.commit()

    async def update_recommendations(self, session: AsyncSession, id: str, recommendations: dict) -> None:
        stmt = (
            update(IncidentInvestigationDb)
            .where(IncidentInvestigationDb.id == id)
            .values(recommendations_json=recommendations)
        )
        await session.execute(stmt)
        await session.commit()

    async def update_status(self, session: AsyncSession, id: str, status: str) -> None:
        stmt = (
            update(IncidentInvestigationDb)
            .where(IncidentInvestigationDb.id == id)
            .values(status=status)
        )
        await session.execute(stmt)
        await session.commit()


incident_repository = IncidentRepository()
