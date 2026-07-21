import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class IncidentInvestigationDb(Base):
    __tablename__ = "incident_investigations"

    id = Column(String(36), primary_key=True)
    title = Column(String(100), nullable=False)
    severity = Column(String(50), nullable=False)  # LOW, MEDIUM, HIGH, CRITICAL
    status = Column(String(50), nullable=False)  # INTAKE, INVESTIGATION, ANALYSIS, RECOMMENDATION, CLOSED
    details = Column(Text, nullable=False)
    root_causes_json = Column(JSON, nullable=True)
    recommendations_json = Column(JSON, nullable=True)
    confidence_score = Column(Float, nullable=False, default=1.0)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    reported_at = Column(DateTime, default=datetime.datetime.utcnow)

    evidence = relationship("IncidentEvidenceDb", back_populates="investigation", cascade="all, delete-orphan")


class IncidentEvidenceDb(Base):
    __tablename__ = "incident_evidence"

    id = Column(String(36), primary_key=True)
    investigation_id = Column(String(36), ForeignKey("incident_investigations.id", ondelete="CASCADE"), nullable=False)
    evidence_type = Column(String(50), nullable=False)  # DOCUMENT, PHOTO, INTERVIEW
    source_path = Column(String(255), nullable=False)
    findings = Column(Text, nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    investigation = relationship("IncidentInvestigationDb", back_populates="evidence")
