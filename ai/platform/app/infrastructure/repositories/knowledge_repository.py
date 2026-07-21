import uuid
import datetime
from sqlalchemy import select, delete
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.knowledge_entities import (
    Base,
    DocumentDb,
    DocumentVersionDb,
    ChunkDb,
    RetrievalLogDb,
)
from app.domain.knowledge import DocumentDto, ChunkDto


class KnowledgeRepository:
    """Repository handling DB operations for ingested files, indexes, and chunk collections."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_document(
        self,
        session: AsyncSession,
        filename: str,
        content_type: str,
        file_size: int,
        tenant_id: str,
        department: str | None = None,
        project: str | None = None,
    ) -> DocumentDto:
        doc_id = str(uuid.uuid4())
        db_doc = DocumentDb(
            id=doc_id,
            filename=filename,
            content_type=content_type,
            file_size=file_size,
            tenant_id=tenant_id,
            department=department,
            project=project,
            status="PENDING",
        )
        session.add(db_doc)

        db_version = DocumentVersionDb(
            id=str(uuid.uuid4()),
            document_id=doc_id,
            version_number=1,
        )
        session.add(db_version)
        await session.commit()

        return DocumentDto(
            id=doc_id,
            filename=filename,
            content_type=content_type,
            file_size=file_size,
            tenant_id=tenant_id,
            department=department,
            project=project,
            status="PENDING",
            created_at=datetime.datetime.utcnow(),
        )

    async def update_doc_status(self, session: AsyncSession, doc_id: str, status: str) -> None:
        stmt = select(DocumentDb).where(DocumentDb.id == doc_id)
        res = await session.execute(stmt)
        db_doc = res.scalar_one_or_none()
        if db_doc:
            db_doc.status = status
            await session.commit()

    async def save_chunks(self, session: AsyncSession, chunks: list[ChunkDb]) -> None:
        for chunk in chunks:
            session.add(chunk)
        await session.commit()

    async def get_documents(self, session: AsyncSession, tenant_id: str) -> list[DocumentDto]:
        stmt = select(DocumentDb).where(DocumentDb.tenant_id == tenant_id)
        res = await session.execute(stmt)
        return [
            DocumentDto(
                id=d.id,
                filename=d.filename,
                content_type=d.content_type,
                file_size=d.file_size,
                department=d.department,
                project=d.project,
                status=d.status,
                tenant_id=d.tenant_id,
                created_at=d.created_at,
            )
            for d in res.scalars().all()
        ]

    async def get_document(self, session: AsyncSession, doc_id: str) -> DocumentDb | None:
        stmt = select(DocumentDb).where(DocumentDb.id == doc_id)
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def delete_document(self, session: AsyncSession, doc_id: str) -> None:
        stmt = delete(DocumentDb).where(DocumentDb.id == doc_id)
        await session.execute(stmt)
        await session.commit()

    async def get_chunks_for_retrieval(
        self,
        session: AsyncSession,
        tenant_id: str,
        department: str | None = None,
        project: str | None = None,
    ) -> list[ChunkDb]:
        """Loads chunks for permissions-matched candidate documents."""
        stmt = (
            select(ChunkDb)
            .join(DocumentDb)
            .where(DocumentDb.tenant_id == tenant_id)
            .options(selectinload(ChunkDb.document))
        )
        if department:
            stmt = stmt.where(DocumentDb.department == department)
        if project:
            stmt = stmt.where(DocumentDb.project == project)

        res = await session.execute(stmt)
        return list(res.scalars().all())

    async def log_retrieval(self, session: AsyncSession, query: str, duration_ms: int, source_count: int) -> None:
        log = RetrievalLogDb(
            id=str(uuid.uuid4()),
            query=query,
            duration_ms=duration_ms,
            source_count=source_count,
        )
        session.add(log)
        await session.commit()


knowledge_repository = KnowledgeRepository()
