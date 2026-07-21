import datetime
from sqlalchemy import Column, String, Text, DateTime, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class DocumentDb(Base):
    __tablename__ = "rag_documents"

    id = Column(String(36), primary_key=True)
    filename = Column(String(255), nullable=False)
    content_type = Column(String(100), nullable=False)
    file_size = Column(Integer, nullable=False)
    department = Column(String(100), nullable=True)
    project = Column(String(100), nullable=True)
    tenant_id = Column(String(36), nullable=False)
    status = Column(String(20), nullable=False, default="PENDING")  # PENDING, INDEXED, FAILED
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    versions = relationship("DocumentVersionDb", back_populates="document", cascade="all, delete-orphan")
    chunks = relationship("ChunkDb", back_populates="document", cascade="all, delete-orphan")


class DocumentVersionDb(Base):
    __tablename__ = "rag_document_versions"

    id = Column(String(36), primary_key=True)
    document_id = Column(String(36), ForeignKey("rag_documents.id", ondelete="CASCADE"), nullable=False)
    version_number = Column(Integer, nullable=False, default=1)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    document = relationship("DocumentDb", back_populates="versions")


class ChunkDb(Base):
    __tablename__ = "rag_chunks"

    id = Column(String(36), primary_key=True)
    document_id = Column(String(36), ForeignKey("rag_documents.id", ondelete="CASCADE"), nullable=False)
    chunk_index = Column(Integer, nullable=False)
    content = Column(Text, nullable=False)
    page_number = Column(Integer, nullable=True)
    section_title = Column(String(255), nullable=True)
    token_count = Column(Integer, nullable=False, default=0)
    embedding_json = Column(JSON, nullable=True)  # Cosine comparison float array

    document = relationship("DocumentDb", back_populates="chunks")


class RetrievalLogDb(Base):
    __tablename__ = "rag_retrieval_logs"

    id = Column(String(36), primary_key=True)
    query = Column(String(500), nullable=False)
    duration_ms = Column(Integer, nullable=False)
    source_count = Column(Integer, nullable=False)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)
