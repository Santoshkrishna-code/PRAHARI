import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.supervisor_entities import Base, SupervisorSessionDb, SupervisorTaskExecutionDb


class SupervisorRepository:
    """Repository handling DB operations for Supervisor Sessions and Multi-Agent Task Executions."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_session(
        self,
        session: AsyncSession,
        user_query: str,
        status: str,
        dag_json: list,
        response_text: str,
        user_id: str,
        tenant_id: str,
    ) -> SupervisorSessionDb:
        s_id = str(uuid.uuid4())
        db_s = SupervisorSessionDb(
            id=s_id,
            user_query=user_query,
            plan_status=status,
            execution_dag_json={"dag": dag_json},
            unified_response=response_text,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_s)
        await session.commit()
        return db_s

    async def get_session(self, session: AsyncSession, id: str) -> SupervisorSessionDb | None:
        stmt = (
            select(SupervisorSessionDb)
            .where(SupervisorSessionDb.id == id)
            .options(selectinload(SupervisorSessionDb.tasks))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_sessions_list(self, session: AsyncSession, tenant_id: str) -> list[SupervisorSessionDb]:
        stmt = (
            select(SupervisorSessionDb)
            .where(SupervisorSessionDb.tenant_id == tenant_id)
            .order_by(SupervisorSessionDb.created_at.desc())
        )
        res = await session.execute(stmt)
        return list(res.scalars().all())

    async def add_task_execution(
        self, session: AsyncSession, session_id: str, step_number: int, assigned_agent: str, status: str, output: dict
    ) -> None:
        db_t = SupervisorTaskExecutionDb(
            id=str(uuid.uuid4()),
            session_id=session_id,
            step_number=step_number,
            assigned_agent=assigned_agent,
            status=status,
            output_json=output,
        )
        session.add(db_t)
        await session.commit()


supervisor_repository = SupervisorRepository()
