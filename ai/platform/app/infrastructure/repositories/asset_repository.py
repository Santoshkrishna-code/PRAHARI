import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.asset_entities import Base, AssetRecordDb, AssetLogDb
from app.domain.assets import AssetRecordDto


class AssetRepository:
    """Repository handling DB operations for Asset records and operational health logs."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_asset(
        self,
        session: AsyncSession,
        asset_name: str,
        asset_category: str,
        location: str,
        operating_hours: float,
        health_score: float,
        user_id: str,
        tenant_id: str,
        status: str = "OPERATIONAL",
    ) -> AssetRecordDto:
        a_id = str(uuid.uuid4())
        db_a = AssetRecordDb(
            id=a_id,
            asset_name=asset_name,
            asset_category=asset_category,
            status=status,
            location=location,
            health_score=health_score,
            operating_hours=operating_hours,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_a)
        await session.commit()

        return AssetRecordDto(
            id=a_id,
            asset_name=asset_name,
            asset_category=asset_category,
            status=status,
            health_score=health_score,
            operating_hours=operating_hours,
            created_at=str(db_a.created_at),
        )

    async def get_asset(self, session: AsyncSession, id: str) -> AssetRecordDb | None:
        stmt = (
            select(AssetRecordDb)
            .where(AssetRecordDb.id == id)
            .options(selectinload(AssetRecordDb.logs))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_assets_list(self, session: AsyncSession, tenant_id: str) -> list[AssetRecordDto]:
        stmt = (
            select(AssetRecordDb)
            .where(AssetRecordDb.tenant_id == tenant_id)
            .order_by(AssetRecordDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            AssetRecordDto(
                id=r.id,
                asset_name=r.asset_name,
                asset_category=r.asset_category,
                status=r.status,
                health_score=r.health_score,
                operating_hours=r.operating_hours,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def add_log(
        self, session: AsyncSession, asset_id: str, log_type: str, health_score: float, notes: str | None
    ) -> None:
        db_log = AssetLogDb(
            id=str(uuid.uuid4()),
            asset_id=asset_id,
            log_type=log_type,
            health_score=health_score,
            notes=notes,
        )
        session.add(db_log)
        await session.commit()

    async def update_asset_status_and_health(
        self, session: AsyncSession, id: str, status: str, health_score: float
    ) -> None:
        stmt = (
            update(AssetRecordDb)
            .where(AssetRecordDb.id == id)
            .values(status=status, health_score=health_score)
        )
        await session.execute(stmt)
        await session.commit()


asset_repository = AssetRepository()
