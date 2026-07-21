import uuid
import datetime
from typing import Any
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.tool_entities import (
    Base,
    ToolDb,
    ToolVersionDb,
    ToolExecutionDb,
    ToolPermissionDb,
    ToolAuditDb,
)
from app.domain.tools import ToolDto, ToolVersionDto


class ToolRepository:
    """Repository handling DB operations for tool registrations, versions, executions and permissions."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def register_tool(
        self,
        session: AsyncSession,
        name: str,
        description: str,
        category: str,
        tool_type: str,
        version_string: str,
        input_schema: dict,
        execution_target: str,
        timeout_seconds: int,
        retry_count: int,
        user_id: str,
    ) -> ToolDto:
        # Check if tool already exists
        stmt = select(ToolDb).where(ToolDb.name == name)
        res = await session.execute(stmt)
        db_tool = res.scalar_one_or_none()

        if not db_tool:
            db_tool = ToolDb(
                id=str(uuid.uuid4()),
                name=name,
                description=description,
                category=category,
                type=tool_type,
                timeout_seconds=timeout_seconds,
                retry_count=retry_count,
                active_version_string=version_string,
            )
            session.add(db_tool)
        else:
            db_tool.active_version_string = version_string
            db_tool.description = description
            db_tool.category = category
            db_tool.timeout_seconds = timeout_seconds
            db_tool.retry_count = retry_count

        # Check if version already exists
        ver_stmt = (
            select(ToolVersionDb)
            .where(ToolVersionDb.tool_id == db_tool.id)
            .where(ToolVersionDb.version_string == version_string)
        )
        ver_res = await session.execute(ver_stmt)
        db_ver = ver_res.scalar_one_or_none()

        if not db_ver:
            version_id = str(uuid.uuid4())
            db_ver = ToolVersionDb(
                id=version_id,
                tool_id=db_tool.id,
                version_string=version_string,
                input_schema_json=input_schema,
                execution_target=execution_target,
            )
            session.add(db_ver)
        else:
            db_ver.input_schema_json = input_schema
            db_ver.execution_target = execution_target

        # Log Audit
        db_audit = ToolAuditDb(
            id=str(uuid.uuid4()),
            tool_id=db_tool.id,
            action="REGISTER",
            user_id=user_id,
            details=f"Tool registered with version {version_string}",
        )
        session.add(db_audit)
        await session.commit()

        return ToolDto(
            id=db_tool.id,
            name=name,
            description=description,
            category=category,
            type=tool_type,
            status="ACTIVE",
            timeout_seconds=timeout_seconds,
            retry_count=retry_count,
            active_version_string=version_string,
        )

    async def get_tools_list(self, session: AsyncSession) -> list[ToolDto]:
        stmt = select(ToolDb)
        res = await session.execute(stmt)
        return [
            ToolDto(
                id=t.id,
                name=t.name,
                description=t.description,
                category=t.category,
                type=t.type,
                status=t.status,
                timeout_seconds=t.timeout_seconds,
                retry_count=t.retry_count,
                active_version_string=t.active_version_string,
            )
            for t in res.scalars().all()
        ]

    async def get_tool_by_name(self, session: AsyncSession, name: str) -> ToolDb | None:
        stmt = select(ToolDb).where(ToolDb.name == name).options(selectinload(ToolDb.versions))
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def get_version_by_string(
        self, session: AsyncSession, tool_id: str, version_str: str
    ) -> ToolVersionDb | None:
        stmt = (
            select(ToolVersionDb)
            .where(ToolVersionDb.tool_id == tool_id)
            .where(ToolVersionDb.version_string == version_str)
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def log_execution(
        self,
        session: AsyncSession,
        tool_id: str,
        version_id: str,
        user_id: str,
        tenant_id: str,
        input_payload: dict,
        output_payload: Any,
        status: str,
        duration_ms: int,
    ) -> str:
        exec_id = str(uuid.uuid4())
        db_exec = ToolExecutionDb(
            id=exec_id,
            tool_id=tool_id,
            version_id=version_id,
            user_id=user_id,
            tenant_id=tenant_id,
            input_payload_json=input_payload,
            output_payload_json=output_payload,
            status=status,
            duration_ms=duration_ms,
        )
        session.add(db_exec)
        
        # Audit log
        db_audit = ToolAuditDb(
            id=str(uuid.uuid4()),
            tool_id=tool_id,
            execution_id=exec_id,
            action="EXECUTE",
            user_id=user_id,
            details=f"Executed with status {status}",
        )
        session.add(db_audit)
        await session.commit()
        return exec_id

    async def add_permission(
        self, session: AsyncSession, tool_id: str, role: str, tenant_id: str, dept: str | None = None
    ) -> None:
        perm = ToolPermissionDb(
            id=str(uuid.uuid4()),
            tool_id=tool_id,
            role=role,
            tenant_id=tenant_id,
            department=dept,
        )
        session.add(perm)
        await session.commit()

    async def check_permission(
        self, session: AsyncSession, tool_id: str, role: str, tenant_id: str, dept: str | None = None
    ) -> bool:
        # If no permissions exist in DB, default to permit access (open config model)
        perm_stmt = select(ToolPermissionDb).where(ToolPermissionDb.tool_id == tool_id)
        perm_res = await session.execute(perm_stmt)
        records = perm_res.scalars().all()
        if not records:
            return True

        # Check matching
        for rec in records:
            if rec.tenant_id == tenant_id and rec.role == role:
                if not rec.department or rec.department == dept:
                    return True
        return False


tool_repository = ToolRepository()
