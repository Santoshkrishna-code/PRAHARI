import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import engine
from app.infrastructure.models.prompt_entities import (
    Base,
    PromptDb,
    PromptVersionDb,
    PromptCategoryDb,
    PromptTagDb,
    PromptAuditDb,
    PromptApprovalDb,
)
from app.domain.prompts import PromptDto, PromptVersionDto, PromptCategoryDto


class PromptRepository:
    """Repository implementation managing DB actions for prompts and version histories."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_category(self, session: AsyncSession, name: str, code: str, desc: str | None = None) -> PromptCategoryDto:
        cat_id = str(uuid.uuid4())
        db_cat = PromptCategoryDb(id=cat_id, name=name, code=code, description=desc)
        session.add(db_cat)
        await session.commit()
        return PromptCategoryDto(id=cat_id, name=name, code=code, description=desc)

    async def create_prompt(
        self, session: AsyncSession, name: str, description: str | None, category_id: str | None, user_id: str
    ) -> PromptDto:
        prompt_id = str(uuid.uuid4())
        db_prompt = PromptDb(id=prompt_id, name=name, description=description, category_id=category_id)
        session.add(db_prompt)
        
        # Log Audit
        db_audit = PromptAuditDb(
            id=str(uuid.uuid4()),
            prompt_id=prompt_id,
            action="CREATE",
            user_id=user_id,
            details=f"Prompt '{name}' registered.",
        )
        session.add(db_audit)
        await session.commit()
        return PromptDto(id=prompt_id, name=name, description=description, category_id=category_id)

    async def add_version(
        self,
        session: AsyncSession,
        prompt_name: str,
        version_str: str,
        system_template: str,
        user_template: str,
        user_id: str,
        few_shots: list[dict] | None = None,
        response_format: dict | None = None,
    ) -> PromptVersionDto:
        # Resolve prompt parent ID
        stmt = select(PromptDb).where(PromptDb.name == prompt_name)
        res = await session.execute(stmt)
        prompt = res.scalar_one_or_none()
        if not prompt:
            raise ValueError(f"Prompt '{prompt_name}' not found.")

        version_id = str(uuid.uuid4())
        db_version = PromptVersionDb(
            id=version_id,
            prompt_id=prompt.id,
            version_string=version_str,
            system_template=system_template,
            user_template=user_template,
            few_shots=few_shots,
            response_format=response_format,
            status="DRAFT",
            created_by=user_id,
        )
        session.add(db_version)

        # Log Audit
        db_audit = PromptAuditDb(
            id=str(uuid.uuid4()),
            prompt_id=prompt.id,
            version_id=version_id,
            action="ADD_VERSION",
            user_id=user_id,
            details=f"Draft version '{version_str}' registered.",
        )
        session.add(db_audit)
        await session.commit()

        return self._map_version(db_version)

    async def approve_version(self, session: AsyncSession, version_id: str, reviewer_id: str) -> None:
        stmt = select(PromptVersionDb).where(PromptVersionDb.id == version_id)
        res = await session.execute(stmt)
        version = res.scalar_one_or_none()
        if not version:
            raise ValueError("Prompt version not found.")

        # Update version status to approved
        version.status = "APPROVED"
        version.approved_by = reviewer_id

        # Insert approval log
        db_approval = PromptApprovalDb(
            id=str(uuid.uuid4()),
            version_id=version_id,
            reviewer_id=reviewer_id,
            decision="APPROVED",
            comments="Approved for production run.",
        )
        session.add(db_approval)

        # Update prompt parent to make this active version
        prompt_stmt = select(PromptDb).where(PromptDb.id == version.prompt_id)
        prompt_res = await session.execute(prompt_stmt)
        prompt = prompt_res.scalar_one()
        
        # Deprecate previous active version if any
        if prompt.active_version_string:
            dep_stmt = (
                update(PromptVersionDb)
                .where(PromptVersionDb.prompt_id == prompt.id)
                .where(PromptVersionDb.version_string == prompt.active_version_string)
                .values(status="DEPRECATED")
            )
            await session.execute(dep_stmt)

        prompt.active_version_string = version.version_string

        # Log Audit
        db_audit = PromptAuditDb(
            id=str(uuid.uuid4()),
            prompt_id=prompt.id,
            version_id=version_id,
            action="ACTIVATE",
            user_id=reviewer_id,
            details=f"Approved and activated version '{version.version_string}' as active.",
        )
        session.add(db_audit)
        await session.commit()

    async def get_active_version(self, session: AsyncSession, prompt_name: str) -> PromptVersionDto | None:
        """Fetches active approved template version for a prompt."""
        stmt = (
            select(PromptVersionDb)
            .join(PromptDb)
            .where(PromptDb.name == prompt_name)
            .where(PromptVersionDb.version_string == PromptDb.active_version_string)
            .where(PromptVersionDb.status == "APPROVED")
        )
        res = await session.execute(stmt)
        db_version = res.scalar_one_or_none()
        if not db_version:
            return None
        return self._map_version(db_version)

    async def get_version_by_number(self, session: AsyncSession, prompt_name: str, version_str: str) -> PromptVersionDto | None:
        """Fetches version details for requested semantic string."""
        stmt = (
            select(PromptVersionDb)
            .join(PromptDb)
            .where(PromptDb.name == prompt_name)
            .where(PromptVersionDb.version_string == version_str)
        )
        res = await session.execute(stmt)
        db_version = res.scalar_one_or_none()
        if not db_version:
            return None
        return self._map_version(db_version)

    async def get_versions_history(self, session: AsyncSession, prompt_name: str) -> list[PromptVersionDto]:
        """Lists semantic versioning history for a prompt."""
        stmt = (
            select(PromptVersionDb)
            .join(PromptDb)
            .where(PromptDb.name == prompt_name)
            .order_by(PromptVersionDb.created_at.desc())
        )
        res = await session.execute(stmt)
        return [self._map_version(v) for v in res.scalars().all()]

    async def get_prompts_list(self, session: AsyncSession) -> list[PromptDto]:
        """Lists all registered prompt definitions."""
        stmt = select(PromptDb)
        res = await session.execute(stmt)
        return [
            PromptDto(
                id=p.id,
                name=p.name,
                description=p.description,
                category_id=p.category_id,
                active_version_string=p.active_version_string,
            )
            for p in res.scalars().all()
        ]

    def _map_version(self, db_v: PromptVersionDb) -> PromptVersionDto:
        return PromptVersionDto(
            id=db_v.id,
            prompt_id=db_v.prompt_id,
            version_string=db_v.version_string,
            system_template=db_v.system_template,
            user_template=db_v.user_template,
            few_shots=db_v.few_shots or [],
            response_format=db_v.response_format,
            status=db_v.status,
            created_by=db_v.created_by,
            approved_by=db_v.approved_by,
            metadata=db_v.metadata_json or {},
        )


prompt_repository = PromptRepository()
