import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class SupervisorSessionDb(Base):
    __tablename__ = "supervisor_sessions"

    id = Column(String(36), primary_key=True)
    user_query = Column(Text, nullable=False)
    plan_status = Column(String(50), nullable=False, default="COMPLETED")  # PLANNED, EXECUTING, COMPLETED, FAILED
    execution_dag_json = Column(JSON, nullable=True)
    unified_response = Column(Text, nullable=False)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    tasks = relationship("SupervisorTaskExecutionDb", back_populates="session", cascade="all, delete-orphan")


class SupervisorTaskExecutionDb(Base):
    __tablename__ = "supervisor_task_executions"

    id = Column(String(36), primary_key=True)
    session_id = Column(String(36), ForeignKey("supervisor_sessions.id", ondelete="CASCADE"), nullable=False)
    step_number = Column(Integer, nullable=False, default=1)
    assigned_agent = Column(String(50), nullable=False)
    status = Column(String(20), nullable=False, default="SUCCESS")
    output_json = Column(JSON, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    session = relationship("SupervisorSessionDb", back_populates="tasks")
