import datetime
from sqlalchemy import Column, String, Text, DateTime, ForeignKey, JSON, Integer
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class AgentGraphDb(Base):
    __tablename__ = "agent_graphs"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False, unique=True)
    description = Column(Text, nullable=True)
    start_node = Column(String(100), nullable=False)

    nodes = relationship("AgentNodeDb", back_populates="graph", cascade="all, delete-orphan")
    edges = relationship("AgentEdgeDb", back_populates="graph", cascade="all, delete-orphan")


class AgentNodeDb(Base):
    __tablename__ = "agent_nodes"

    id = Column(String(36), primary_key=True)
    graph_id = Column(String(36), ForeignKey("agent_graphs.id", ondelete="CASCADE"), nullable=False)
    name = Column(String(100), nullable=False)
    prompt_template_name = Column(String(100), nullable=True)
    tools_allowed_json = Column(JSON, nullable=True)  # List of allowed tool names
    model = Column(String(50), nullable=False, default="gpt-4o")

    graph = relationship("AgentGraphDb", back_populates="nodes")


class AgentEdgeDb(Base):
    __tablename__ = "agent_edges"

    id = Column(String(36), primary_key=True)
    graph_id = Column(String(36), ForeignKey("agent_graphs.id", ondelete="CASCADE"), nullable=False)
    source_node = Column(String(100), nullable=False)
    target_node = Column(String(100), nullable=False)
    conditional_expr = Column(String(255), nullable=True)

    graph = relationship("AgentGraphDb", back_populates="edges")


class AgentStateCheckpointDb(Base):
    __tablename__ = "agent_state_checkpoints"

    id = Column(String(36), primary_key=True)
    execution_id = Column(String(36), nullable=False, index=True)
    graph_id = Column(String(36), ForeignKey("agent_graphs.id", ondelete="CASCADE"), nullable=False)
    current_node = Column(String(100), nullable=False)
    state_json = Column(JSON, nullable=False)  # Serialized dictionary of graph execution variables
    status = Column(String(20), nullable=False)  # RUNNING, COMPLETED, INTERRUPTED, FAILED
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow, onupdate=datetime.datetime.utcnow)


class HumanApprovalDb(Base):
    __tablename__ = "agent_human_approvals"

    id = Column(String(36), primary_key=True)
    execution_id = Column(String(36), nullable=False, index=True)
    node_name = Column(String(100), nullable=False)
    status = Column(String(20), nullable=False, default="PENDING")  # PENDING, APPROVED, REJECTED
    comments = Column(Text, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
