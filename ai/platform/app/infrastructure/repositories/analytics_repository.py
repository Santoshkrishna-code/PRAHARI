import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import engine
from app.infrastructure.models.analytics_entities import Base, ExecutiveReportDb, AnalyticsKpiSnapshotDb
from app.domain.analytics import ReportGenerateResponse, KpiMetricsResponse


class AnalyticsRepository:
    """Repository handling DB operations for Executive Reports and KPI Snapshots."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def save_report(
        self,
        session: AsyncSession,
        title: str,
        report_type: str,
        markdown: str,
        kpis: dict,
        user_id: str,
        tenant_id: str,
    ) -> ReportGenerateResponse:
        r_id = str(uuid.uuid4())
        db_r = ExecutiveReportDb(
            id=r_id,
            report_title=title,
            report_type=report_type,
            executive_summary_markdown=markdown,
            kpis_json=kpis,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_r)
        await session.commit()

        return ReportGenerateResponse(
            report_id=r_id,
            title=title,
            report_type=report_type,
            report_markdown=markdown,
            created_at=str(db_r.created_at),
        )

    async def get_reports_list(self, session: AsyncSession, tenant_id: str) -> list[ReportGenerateResponse]:
        stmt = (
            select(ExecutiveReportDb)
            .where(ExecutiveReportDb.tenant_id == tenant_id)
            .order_by(ExecutiveReportDb.created_at.desc())
        )
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            ReportGenerateResponse(
                report_id=r.id,
                title=r.report_title,
                report_type=r.report_type,
                report_markdown=r.executive_summary_markdown,
                created_at=str(r.created_at),
            )
            for r in rows
        ]

    async def get_latest_kpis(self, session: AsyncSession, tenant_id: str) -> KpiMetricsResponse:
        stmt = (
            select(AnalyticsKpiSnapshotDb)
            .where(AnalyticsKpiSnapshotDb.tenant_id == tenant_id)
            .order_by(AnalyticsKpiSnapshotDb.created_at.desc())
        )
        res = await session.execute(stmt)
        snapshot = res.scalars().first()

        if snapshot:
            return KpiMetricsResponse(
                trir=snapshot.trir,
                ltifr=snapshot.ltifr,
                near_miss_count=8,
                permit_compliance_pct=snapshot.permit_compliance_rate,
                inspection_completion_pct=snapshot.inspection_completion_rate,
                asset_availability_pct=snapshot.asset_availability_rate,
                overall_compliance_score=snapshot.compliance_score,
            )

        # Default platform baseline KPIs
        return KpiMetricsResponse(
            trir=1.1,
            ltifr=0.3,
            near_miss_count=12,
            permit_compliance_pct=97.2,
            inspection_completion_pct=95.5,
            asset_availability_pct=98.6,
            overall_compliance_score=94.0,
        )


analytics_repository = AnalyticsRepository()
