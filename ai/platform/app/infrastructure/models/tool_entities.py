import datetime
from sqlalchemy import Column, String, Text, DateTime, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class ToolDb(Base):
    __tablename__ = "tools"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False, unique=True)
    description = Column(Text, nullable=False)
    category = Column(String(50), nullable=False)
    type = Column(String(20), nullable=False)  # REST, SQL, INTERNAL
    status = Column(String(20), nullable=False, default="ACTIVE")  # ACTIVE, DEPRECATED
    timeout_seconds = Column(Integer, default=30)
    retry_count = Column(Integer, default=3)
    active_version_string = Column(String(20), nullable=True)

    versions = relationship("ToolVersionDb", back_populates="tool", cascade="all, delete-orphan")
    executions = relationship("ToolExecutionDb", back_populates="tool", cascade="all, delete-orphan")


class ToolVersionDb(Base):
    __tablename__ = "tool_versions"

    id = Column(String(36), primary_key=True)
    tool_id = Column(String(36), ForeignKey("tools.id", ondelete="CASCADE"), nullable=False)
    version_string = Column(String(20), nullable=False)
    input_schema_json = Column(JSON, nullable=False)
    execution_target = Column(Text, nullable=False)
    status = Column(String(20), nullable=False, default="ACTIVE")  # ACTIVE, DEPRECATED

    tool = relationship("ToolDb", back_populates="versions")
    executions = relationship("ToolExecutionDb", back_populates="version", cascade="all, delete-orphan")


class ToolExecutionDb(Base):
    __tablename__ = "tool_executions"

    id = Column(String(36), primary_key=True)
    tool_id = Column(String(36), ForeignKey("tools.id", ondelete="CASCADE"), nullable=False)
    version_id = Column(String(36), ForeignKey("tool_versions.id", ondelete="CASCADE"), nullable=False)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    input_payload_json = Column(JSON, nullable=True)
    output_payload_json = Column(JSON, nullable=True)
    status = Column(String(20), nullable=False)  # SUCCESS, FAILED
    duration_ms = Column(Integer, nullable=False)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)

    tool = relationship("ToolDb", back_populates="executions")
    version = relationship("ToolVersionDb", back_populates="executions")


class ToolPermissionDb(Base):
    __tablename__ = "tool_permissions"

    id = Column(String(36), primary_key=True)
    tool_id = Column(String(36), ForeignKey("tools.id", ondelete="CASCADE"), nullable=False)
    role = Column(String(50), nullable=False)
    department = Column(String(50), nullable=True)
    tenant_id = Column(String(36), nullable=False)


class ToolAuditDb(Base):
    __tablename__ = "tool_audits"

    id = Column(String(36), primary_key=True)
    tool_id = Column(String(36), nullable=True)
    execution_id = Column(String(36), nullable=True)
    action = Column(String(50), nullable=False)  # REGISTER, EXECUTE, DEPRECATE
    user_id = Column(String(100), nullable=False)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)
    details = Column(Text, nullable=True)
