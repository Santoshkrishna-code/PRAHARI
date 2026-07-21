import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class ComplianceAssessmentDb(Base):
    __tablename__ = "compliance_assessments"

    id = Column(String(36), primary_key=True)
    target_subject = Column(String(100), nullable=False)
    subject_type = Column(String(50), nullable=False)  # PERMIT, INSPECTION, TRAINING, PPE
    compliance_score = Column(Float, nullable=False, default=100.0)
    status = Column(String(50), nullable=False)  # COMPLIANT, NON_COMPLIANT, FLAGGED
    details = Column(Text, nullable=False)
    capa_plan_json = Column(JSON, nullable=True)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    violations = relationship("ComplianceViolationDb", back_populates="assessment", cascade="all, delete-orphan")


class ComplianceViolationDb(Base):
    __tablename__ = "compliance_violations"

    id = Column(String(36), primary_key=True)
    assessment_id = Column(String(36), ForeignKey("compliance_assessments.id", ondelete="CASCADE"), nullable=False)
    regulation_violated = Column(String(50), nullable=False)  # OSHA, ISO_45001, NFPA, etc.
    severity = Column(String(20), nullable=False)  # LOW, MEDIUM, HIGH
    details = Column(Text, nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    assessment = relationship("ComplianceAssessmentDb", back_populates="violations")
