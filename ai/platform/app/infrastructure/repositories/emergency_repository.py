import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import engine
from app.infrastructure.models.emergency_entities import Base, EmergencyRecordDb, EmergencyResourceDb
from app.domain.emergencies import EmergencyRecordDto, EmergencyResourceDto


class EmergencyRepository:
    """Repository handling DB operations for Emergency events and response equipment inventory."""

    async def initialize_tables(self) -> None:
        """Create tables and seed initial emergency resources on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

        # Seed initial emergency resources if table empty
        async with AsyncSession(engine) as session:
            stmt = select(EmergencyResourceDb)
            res = await session.execute(stmt)
            if not res.scalars().all():
                session.add_all(
                    [
                        EmergencyResourceDb(
                            id=str(uuid.uuid4()),
                            resource_name="Industrial HazMat Spill Kit Unit A",
                            resource_type="SPILL_KIT",
                            quantity=4,
                            location="Sector 2 Bay",
                            status="AVAILABLE",
                        ),
                        EmergencyResourceDb(
                            id=str(uuid.uuid4()),
                            resource_name="Self-Contained Breathing Apparatus (SCBA) Pack",
                            resource_type="SCBA",
                            quantity=10,
                            location="Main Rescue Locker",
                            status="AVAILABLE",
                        ),
                        EmergencyResourceDb(
                            id=str(uuid.uuid4()),
                            resource_name="Emergency Response Ambulance Unit 1",
                            resource_type="AMBULANCE",
                            quantity=2,
                            location="Medical Station North",
                            status="AVAILABLE",
                        ),
                    ]
                )
                await session.commit()

    async def create_emergency(
        self,
        session: AsyncSession,
        title: str,
        emergency_type: str,
        severity_level: str,
        location: str,
        affected_zone: str,
        evacuation_route: str,
        details: str,
        user_id: str,
        tenant_id: str,
    ) -> EmergencyRecordDto:
        e_id = str(uuid.uuid4())
        db_e = EmergencyRecordDb(
            id=e_id,
            title=title,
            emergency_type=emergency_type,
            severity_level=severity_level,
            status="ACTIVE",
            location=location,
            affected_zone=affected_zone,
            evacuation_route=evacuation_route,
            details=details,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_e)
        await session.commit()

        return EmergencyRecordDto(
            id=e_id,
            title=title,
            emergency_type=emergency_type,
            severity_level=severity_level,
            status="ACTIVE",
            location=location,
            affected_zone=affected_zone,
            evacuation_route=evacuation_route,
            created_at=str(db_e.created_at),
        )

    async def get_emergency(self, session: AsyncSession, id: str) -> EmergencyRecordDb | None:
        stmt = select(EmergencyRecordDb).where(EmergencyRecordDb.id == id)
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_emergencies_list(self, session: AsyncSession, tenant_id: str) -> list[EmergencyRecordDto]:
        stmt = (
            select(EmergencyRecordDb)
            .where(EmergencyRecordDb.tenant_id == tenant_id)
            .order_by(EmergencyRecordDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            EmergencyRecordDto(
                id=r.id,
                title=r.title,
                emergency_type=r.emergency_type,
                severity_level=r.severity_level,
                status=r.status,
                location=r.location,
                affected_zone=r.affected_zone,
                evacuation_route=r.evacuation_route,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def get_resources_list(self, session: AsyncSession) -> list[EmergencyResourceDto]:
        stmt = select(EmergencyResourceDb).order_by(EmergencyResourceDb.created_at.desc())
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            EmergencyResourceDto(
                id=r.id,
                resource_name=r.resource_name,
                resource_type=r.resource_type,
                quantity=r.quantity,
                location=r.location,
                status=r.status,
            )
            for r in rows
        ]

    async def update_emergency_status(self, session: AsyncSession, id: str, status: str) -> None:
        stmt = (
            update(EmergencyRecordDb)
            .where(EmergencyRecordDb.id == id)
            .values(status=status)
        )
        await session.execute(stmt)
        await session.commit()


emergency_repository = EmergencyRepository()
