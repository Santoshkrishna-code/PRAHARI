import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import engine
from app.infrastructure.models.maintenance_entities import Base, MaintenancePredictionDb, MaintenanceWorkOrderDb
from app.domain.maintenance import MaintenancePredictionDto, MaintenancePlanResponse, ReliabilityDashboardResponse


class MaintenanceRepository:
    """Repository handling DB operations for Maintenance Predictions and Work Orders."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_prediction(
        self,
        session: AsyncSession,
        asset_id: str,
        failure_prob: float,
        rul_hours: float,
        rec_date: str,
        strategy: str,
        spare_parts: list[str],
        user_id: str,
        tenant_id: str,
    ) -> MaintenancePredictionDto:
        p_id = str(uuid.uuid4())
        db_p = MaintenancePredictionDb(
            id=p_id,
            asset_id=asset_id,
            failure_probability=failure_prob,
            rul_hours=rul_hours,
            recommended_maintenance_date=rec_date,
            maintenance_strategy=strategy,
            status="RECOMMENDED",
            spare_parts_json={"parts": spare_parts},
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_p)
        await session.commit()

        return MaintenancePredictionDto(
            id=p_id,
            asset_id=asset_id,
            failure_probability=failure_prob,
            rul_hours=rul_hours,
            recommended_maintenance_date=rec_date,
            maintenance_strategy=strategy,
            status="RECOMMENDED",
            created_at=str(db_p.created_at),
        )

    async def get_predictions_list(self, session: AsyncSession, tenant_id: str) -> list[MaintenancePredictionDto]:
        stmt = (
            select(MaintenancePredictionDb)
            .where(MaintenancePredictionDb.tenant_id == tenant_id)
            .order_by(MaintenancePredictionDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            MaintenancePredictionDto(
                id=r.id,
                asset_id=r.asset_id,
                failure_probability=r.failure_probability,
                rul_hours=r.rul_hours,
                recommended_maintenance_date=r.recommended_maintenance_date,
                maintenance_strategy=r.maintenance_strategy,
                status=r.status,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def create_work_order(
        self, session: AsyncSession, asset_id: str, title: str, priority: str, downtime_est: float, spare_parts: list[str]
    ) -> MaintenancePlanResponse:
        wo_id = str(uuid.uuid4())
        db_wo = MaintenanceWorkOrderDb(
            id=wo_id,
            asset_id=asset_id,
            work_order_title=title,
            priority=priority,
            status="OPEN",
            downtime_hours_est=downtime_est,
        )
        session.add(db_wo)
        await session.commit()

        return MaintenancePlanResponse(
            work_order_id=wo_id,
            asset_id=asset_id,
            priority=priority,
            downtime_hours_est=downtime_est,
            required_spare_parts=spare_parts,
        )

    async def get_reliability_metrics(self, session: AsyncSession, tenant_id: str) -> ReliabilityDashboardResponse:
        stmt_p = select(MaintenancePredictionDb).where(MaintenancePredictionDb.tenant_id == tenant_id)
        res_p = await session.execute(stmt_p)
        preds = res_p.scalars().all()

        stmt_wo = select(MaintenanceWorkOrderDb).where(MaintenanceWorkOrderDb.status == "OPEN")
        res_wo = await session.execute(stmt_wo)
        wos = res_wo.scalars().all()

        total = len(preds)
        high_risk = len([p for p in preds if p.failure_probability > 0.6])
        avg_rul = (sum(p.rul_hours for p in preds) / total) if total > 0 else 720.0

        return ReliabilityDashboardResponse(
            total_assets_monitored=max(total, 5),
            high_risk_assets_count=high_risk,
            average_rul_hours=round(avg_rul, 1),
            pending_work_orders_count=len(wos),
        )


maintenance_repository = MaintenanceRepository()
