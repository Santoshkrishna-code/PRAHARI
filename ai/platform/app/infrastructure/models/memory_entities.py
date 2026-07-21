import datetime
from sqlalchemy import Column, String, Text, DateTime, ForeignKey, Integer, Float, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class SessionDb(Base):
    __tablename__ = "memory_sessions"

    id = Column(String(36), primary_key=True)
    user_id = Column(String(36), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    status = Column(String(20), nullable=False, default="ACTIVE")  # ACTIVE, EXPIRED
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    expires_at = Column(DateTime, nullable=True)

    messages = relationship("MessageDb", back_populates="session", cascade="all, delete-orphan")
    summary = relationship("ConversationSummaryDb", back_populates="session", uselist=False, cascade="all, delete-orphan")


class MessageDb(Base):
    __tablename__ = "memory_messages"

    id = Column(String(36), primary_key=True)
    session_id = Column(String(36), ForeignKey("memory_sessions.id", ondelete="CASCADE"), nullable=False)
    role = Column(String(20), nullable=False)  # user, assistant
    content = Column(Text, nullable=False)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)
    metadata_json = Column(JSON, nullable=True)

    session = relationship("SessionDb", back_populates="messages")


class ConversationSummaryDb(Base):
    __tablename__ = "memory_conversation_summaries"

    id = Column(String(36), primary_key=True)
    session_id = Column(String(36), ForeignKey("memory_sessions.id", ondelete="CASCADE"), nullable=False, unique=True)
    summary_text = Column(Text, nullable=False)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow, onupdate=datetime.datetime.utcnow)

    session = relationship("SessionDb", back_populates="summary")


class MemoryEntryDb(Base):
    __tablename__ = "memory_entries"

    id = Column(String(36), primary_key=True)
    user_id = Column(String(36), nullable=False)
    category = Column(String(50), nullable=False)  # preference, profile, incident, task
    content = Column(Text, nullable=False)
    importance = Column(Integer, nullable=False, default=1)  # 1 to 10
    embedding_json = Column(JSON, nullable=True)  # Mapped array of floats
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)

    scores = relationship("MemoryScoreDb", back_populates="entry", uselist=False, cascade="all, delete-orphan")


class MemoryScoreDb(Base):
    __tablename__ = "memory_scores"

    id = Column(String(36), primary_key=True)
    entry_id = Column(String(36), ForeignKey("memory_entries.id", ondelete="CASCADE"), nullable=False, unique=True)
    recency = Column(Float, default=1.0)
    frequency = Column(Integer, default=1)
    importance = Column(Float, default=1.0)
    confidence = Column(Float, default=1.0)

    entry = relationship("MemoryEntryDb", back_populates="scores")


class UserProfileDb(Base):
    __tablename__ = "memory_user_profiles"

    id = Column(String(36), primary_key=True)
    email = Column(String(100), nullable=False, unique=True)
    role = Column(String(50), nullable=False)
    department = Column(String(50), nullable=False)
    timezone = Column(String(50), nullable=False, default="UTC")

    preferences = relationship("PreferenceDb", back_populates="user", cascade="all, delete-orphan")


class PreferenceDb(Base):
    __tablename__ = "memory_preferences"

    id = Column(String(36), primary_key=True)
    user_id = Column(String(36), ForeignKey("memory_user_profiles.id", ondelete="CASCADE"), nullable=False)
    key = Column(String(100), nullable=False)
    value = Column(Text, nullable=False)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow, onupdate=datetime.datetime.utcnow)

    user = relationship("UserProfileDb", back_populates="preferences")


class RetentionPolicyDb(Base):
    __tablename__ = "memory_retention_policies"

    id = Column(String(36), primary_key=True)
    tenant_id = Column(String(36), nullable=False, unique=True)
    policy_type = Column(String(50), nullable=False, default="STANDARD")  # STANDARD, STRICT, LEGAL_HOLD
    retention_days = Column(Integer, default=365)


class MemoryAuditDb(Base):
    __tablename__ = "memory_audits"

    id = Column(String(36), primary_key=True)
    session_id = Column(String(36), nullable=True)
    action = Column(String(50), nullable=False)  # CREATE_SESSION, ADD_MESSAGE, PURGE
    user_id = Column(String(100), nullable=False)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)
    details = Column(Text, nullable=True)
