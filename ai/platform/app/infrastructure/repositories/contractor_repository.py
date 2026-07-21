import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.contractor_entities import Base, ContractorRecordDb, ContractorAuditDb
from app.domain.contractors import ContractorRecordDto


class ContractorRepository:
    """Repository handling DB operations for Contractor profiles and audit records."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_contractor(
        self,
        session: AsyncSession,
        company_name: str,
        contractor_type: str,
        contact_email: str,
        license_number: str,
        user_id: str,
        tenant_id: str,
    ) -> ContractorRecordDto:
        c_id = str(uuid.uuid4())
        db_c = ContractorRecordDb(
            id=c_id,
            company_name=company_name,
            contractor_type=contractor_type,
            status="REGISTERED",
            license_number=license_number,
            contact_email=contact_email,
            safety_rating_score=100.0,
            violations_count=0,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_c)
        await session.commit()

        return ContractorRecordDto(
            id=c_id,
            company_name=company_name,
            contractor_type=contractor_type,
            status="REGISTERED",
            safety_rating_score=100.0,
            violations_count=0,
            created_at=str(db_c.created_at),
        )

    async def get_contractor(self, session: AsyncSession, id: str) -> ContractorRecordDb | None:
        stmt = (
            select(ContractorRecordDb)
            .where(ContractorRecordDb.id == id)
            .options(selectinload(ContractorRecordDb.audits))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_contractors_list(self, session: AsyncSession, tenant_id: str) -> list[ContractorRecordDto]:
        stmt = (
            select(ContractorRecordDb)
            .where(ContractorRecordDb.tenant_id == tenant_id)
            .order_by(ContractorRecordDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            ContractorRecordDto(
                id=r.id,
                company_name=r.company_name,
                contractor_type=r.contractor_type,
                status=r.status,
                safety_rating_score=r.safety_rating_score,
                violations_count=r.violations_count,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def add_audit(
        self, session: AsyncSession, contractor_id: str, audit_type: str, score: float, comments: str | None
    ) -> None:
        db_audit = ContractorAuditDb(
            id=str(uuid.uuid4()),
            contractor_id=contractor_id,
            audit_type=audit_type,
            score=score,
            comments=comments,
        )
        session.add(db_audit)
        await session.commit()

    async def update_contractor_status_and_score(
        self, session: AsyncSession, id: str, status: str, safety_score: float, violations_count: int
    ) -> None:
        stmt = (
            update(ContractorRecordDb)
            .where(ContractorRecordDb.id == id)
            .values(status=status, safety_rating_score=safety_score, violations_count=violations_count)
        )
        await session.execute(stmt)
        await session.commit()


contractor_repository = ContractorRepository()
