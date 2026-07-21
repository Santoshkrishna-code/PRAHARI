import datetime
from typing import Any
from sqlalchemy import Column, String, Text, DateTime, ForeignKey, Table, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


# Association table for prompts and tags many-to-many relationship
prompt_tags_association = Table(
    "prompt_tags_association",
    Base.metadata,
    Column("prompt_id", String(36), ForeignKey("prompts.id", ondelete="CASCADE"), primary_key=True),
    Column("tag_id", String(36), ForeignKey("prompt_tags.id", ondelete="CASCADE"), primary_key=True),
)


class PromptCategoryDb(Base):
    __tablename__ = "prompt_categories"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False)
    code = Column(String(50), nullable=False, unique=True)
    description = Column(String(255), nullable=True)

    prompts = relationship("PromptDb", back_populates="category")


class PromptTagDb(Base):
    __tablename__ = "prompt_tags"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False)
    code = Column(String(50), nullable=False, unique=True)


class PromptDb(Base):
    __tablename__ = "prompts"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False, unique=True)
    description = Column(Text, nullable=True)
    category_id = Column(String(36), ForeignKey("prompt_categories.id", ondelete="SET NULL"), nullable=True)
    active_version_string = Column(String(20), nullable=True)

    category = relationship("PromptCategoryDb", back_populates="prompts")
    versions = relationship("PromptVersionDb", back_populates="prompt", cascade="all, delete-orphan")
    tags = relationship("PromptTagDb", secondary=prompt_tags_association)


class PromptVersionDb(Base):
    __tablename__ = "prompt_versions"

    id = Column(String(36), primary_key=True)
    prompt_id = Column(String(36), ForeignKey("prompts.id", ondelete="CASCADE"), nullable=False)
    version_string = Column(String(20), nullable=False)  # e.g. 1.0.0
    system_template = Column(Text, nullable=False)
    user_template = Column(Text, nullable=False)
    few_shots = Column(JSON, nullable=True)  # List of message structures
    response_format = Column(JSON, nullable=True)  # Schema constraints JSON
    status = Column(String(20), nullable=False, default="DRAFT")  # DRAFT, REVIEW, APPROVED, DEPRECATED
    created_by = Column(String(100), nullable=False)
    approved_by = Column(String(100), nullable=True)
    metadata_json = Column(JSON, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    prompt = relationship("PromptDb", back_populates="versions")
    approvals = relationship("PromptApprovalDb", back_populates="version", cascade="all, delete-orphan")


class PromptAuditDb(Base):
    __tablename__ = "prompt_audits"

    id = Column(String(36), primary_key=True)
    prompt_id = Column(String(36), ForeignKey("prompts.id", ondelete="CASCADE"), nullable=False)
    version_id = Column(String(36), nullable=True)
    action = Column(String(50), nullable=False)  # CREATE, UPDATE_VERSION, ACTIVATE, DEPRECATE, ROLLBACK
    user_id = Column(String(100), nullable=False)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)
    details = Column(Text, nullable=True)


class PromptApprovalDb(Base):
    __tablename__ = "prompt_approvals"

    id = Column(String(36), primary_key=True)
    version_id = Column(String(36), ForeignKey("prompt_versions.id", ondelete="CASCADE"), nullable=False)
    reviewer_id = Column(String(100), nullable=False)
    decision = Column(String(20), nullable=False)  # APPROVED, REJECTED
    comments = Column(Text, nullable=True)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)

    version = relationship("PromptVersionDb", back_populates="approvals")
