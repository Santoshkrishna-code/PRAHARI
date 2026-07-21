import uuid
import datetime
from sqlalchemy import select, update, delete
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.memory_entities import (
    Base,
    SessionDb,
    MessageDb,
    ConversationSummaryDb,
    MemoryEntryDb,
    MemoryScoreDb,
    UserProfileDb,
    PreferenceDb,
)
from app.domain.memory import MessageDto, SessionDto, MemoryEntryDto


class MemoryRepository:
    """Repository handling DB operations for Session, Conversation History, and Declarative Memory."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def create_session(self, session: AsyncSession, user_id: str, tenant_id: str) -> SessionDto:
        session_id = str(uuid.uuid4())
        db_session = SessionDb(id=session_id, user_id=user_id, tenant_id=tenant_id)
        session.add(db_session)
        await session.commit()
        return SessionDto(
            id=session_id,
            user_id=user_id,
            tenant_id=tenant_id,
            created_at=datetime.datetime.utcnow(),
            status="ACTIVE",
        )

    async def get_session(self, session: AsyncSession, session_id: str) -> SessionDto | None:
        stmt = select(SessionDb).where(SessionDb.id == session_id)
        res = await session.execute(stmt)
        db_s = res.scalar_one_or_none()
        if not db_s:
            return None
        return SessionDto(
            id=db_s.id,
            user_id=db_s.user_id,
            tenant_id=db_s.tenant_id,
            created_at=db_s.created_at,
            status=db_s.status,
        )

    async def add_message(
        self, session: AsyncSession, session_id: str, role: str, content: str, metadata: dict | None = None
    ) -> MessageDto:
        message_id = str(uuid.uuid4())
        db_msg = MessageDb(
            id=message_id,
            session_id=session_id,
            role=role,
            content=content,
            metadata_json=metadata or {},
        )
        session.add(db_msg)
        await session.commit()
        return MessageDto(
            id=message_id,
            role=role,
            content=content,
            timestamp=datetime.datetime.utcnow(),
            metadata=metadata or {},
        )

    async def get_history(self, session: AsyncSession, session_id: str, limit: int = 50) -> list[MessageDto]:
        stmt = (
            select(MessageDb)
            .where(MessageDb.session_id == session_id)
            .order_by(MessageDb.timestamp.asc())
            .limit(limit)
        )
        res = await session.execute(stmt)
        return [
            MessageDto(
                id=m.id,
                role=m.role,
                content=m.content,
                timestamp=m.timestamp,
                metadata=m.metadata_json or {},
            )
            for m in res.scalars().all()
        ]

    async def update_summary(self, session: AsyncSession, session_id: str, summary_text: str) -> None:
        # Check if summary already exists
        stmt = select(ConversationSummaryDb).where(ConversationSummaryDb.session_id == session_id)
        res = await session.execute(stmt)
        db_sum = res.scalar_one_or_none()

        if db_sum:
            db_sum.summary_text = summary_text
        else:
            db_sum = ConversationSummaryDb(
                id=str(uuid.uuid4()),
                session_id=session_id,
                summary_text=summary_text,
            )
            session.add(db_sum)
        await session.commit()

    async def get_summary(self, session: AsyncSession, session_id: str) -> str | None:
        stmt = select(ConversationSummaryDb).where(ConversationSummaryDb.session_id == session_id)
        res = await session.execute(stmt)
        db_sum = res.scalar_one_or_none()
        return db_sum.summary_text if db_sum else None

    # Declarative Memories
    async def create_memory_entry(
        self,
        session: AsyncSession,
        user_id: str,
        category: str,
        content: str,
        importance: int,
        embedding: list[float] | None = None,
    ) -> MemoryEntryDto:
        entry_id = str(uuid.uuid4())
        db_entry = MemoryEntryDb(
            id=entry_id,
            user_id=user_id,
            category=category,
            content=content,
            importance=importance,
            embedding_json=embedding,
        )
        session.add(db_entry)

        # Create base score entry
        db_score = MemoryScoreDb(
            id=str(uuid.uuid4()),
            entry_id=entry_id,
            importance=float(importance),
            recency=1.0,
            frequency=1,
        )
        session.add(db_score)
        await session.commit()

        return MemoryEntryDto(
            id=entry_id,
            user_id=user_id,
            category=category,
            content=content,
            importance=importance,
            timestamp=datetime.datetime.utcnow(),
        )

    async def get_memories_for_user(self, session: AsyncSession, user_id: str) -> list[MemoryEntryDb]:
        """Fetches memory records eager loading score metrics."""
        stmt = (
            select(MemoryEntryDb)
            .where(MemoryEntryDb.user_id == user_id)
            .options(selectinload(MemoryEntryDb.scores))
        )
        res = await session.execute(stmt)
        return list(res.scalars().all())

    async def increment_memory_frequency(self, session: AsyncSession, entry_id: str) -> None:
        stmt = select(MemoryScoreDb).where(MemoryScoreDb.entry_id == entry_id)
        res = await session.execute(stmt)
        db_score = res.scalar_one_or_none()
        if db_score:
            db_score.frequency += 1
            await session.commit()

    async def save_profile(
        self, session: AsyncSession, user_id: str, email: str, role: str, dept: str, tz: str
    ) -> None:
        stmt = select(UserProfileDb).where(UserProfileDb.id == user_id)
        res = await session.execute(stmt)
        db_profile = res.scalar_one_or_none()

        if db_profile:
            db_profile.email = email
            db_profile.role = role
            db_profile.department = dept
            db_profile.timezone = tz
        else:
            db_profile = UserProfileDb(
                id=user_id,
                email=email,
                role=role,
                department=dept,
                timezone=tz,
            )
            session.add(db_profile)
        await session.commit()

    async def save_preference(self, session: AsyncSession, user_id: str, key: str, value: str) -> None:
        # Check if user profile exists to satisfy foreign key constraints
        stmt = select(UserProfileDb).where(UserProfileDb.id == user_id)
        res = await session.execute(stmt)
        db_profile = res.scalar_one_or_none()
        if not db_profile:
            # Create a stub profile to satisfy FK
            db_profile = UserProfileDb(
                id=user_id,
                email=f"{user_id}@prahari.internal",
                role="user",
                department="general",
                timezone="UTC",
            )
            session.add(db_profile)

        pref_stmt = select(PreferenceDb).where(PreferenceDb.user_id == user_id).where(PreferenceDb.key == key)
        pref_res = await session.execute(pref_stmt)
        db_pref = pref_res.scalar_one_or_none()

        if db_pref:
            db_pref.value = value
        else:
            db_pref = PreferenceDb(
                id=str(uuid.uuid4()),
                user_id=user_id,
                key=key,
                value=value,
            )
            session.add(db_pref)
        await session.commit()

    async def get_preferences(self, session: AsyncSession, user_id: str) -> dict[str, str]:
        stmt = select(PreferenceDb).where(PreferenceDb.user_id == user_id)
        res = await session.execute(stmt)
        return {p.key: p.value for p in res.scalars().all()}

    async def purge_user_memory(self, session: AsyncSession, user_id: str) -> None:
        """Completely purges all data related to a user ID to support compliance deletion."""
        # Delete entries and summaries linked to the user sessions
        sess_stmt = select(SessionDb.id).where(SessionDb.user_id == user_id)
        sess_res = await session.execute(sess_stmt)
        session_ids = list(sess_res.scalars().all())

        if session_ids:
            await session.execute(delete(SessionDb).where(SessionDb.id.in_(session_ids)))

        await session.execute(delete(MemoryEntryDb).where(MemoryEntryDb.user_id == user_id))
        await session.execute(delete(PreferenceDb).where(PreferenceDb.user_id == user_id))
        await session.execute(delete(UserProfileDb).where(UserProfileDb.id == user_id))
        await session.commit()


memory_repository = MemoryRepository()
