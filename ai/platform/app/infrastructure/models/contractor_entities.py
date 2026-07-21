import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class ContractorRecordDb(Base):
    __tablename__ = "contractor_records"

    id = Column(String(36), primary_key=True)
    company_name = Column(String(100), nullable=False)
    contractor_type = Column(String(50), nullable=False)  # CIVIL, MECHANICAL, SCAFFOLDING, etc.
    status = Column(String(50), nullable=False, default="REGISTERED")  # REGISTERED, QUALIFIED, REJECTED, SUSPENDED
    license_number = Column(String(100), nullable=False)
    contact_email = Column(String(100), nullable=False)
    safety_rating_score = Column(Float, nullable=False, default=100.0)
    violations_count = Column(Integer, nullable=False, default=0)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    audits = relationship("ContractorAuditDb", back_populates="contractor", cascade="all, delete-orphan")


class ContractorAuditDb(Base):
    __tablename__ = "contractor_audits"

    id = Column(String(36), primary_key=True)
    contractor_id = Column(String(36), ForeignKey("contractor_records.id", ondelete="CASCADE"), nullable=False)
    audit_type = Column(String(50), nullable=False)  # PRE_QUALIFICATION, PERIODIC_SAFETY_AUDIT
    score = Column(Float, nullable=False, default=100.0)
    comments = Column(Text, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    contractor = relationship("ContractorRecordDb", back_populates="audits")
