import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class PermitRecordDb(Base):
    __tablename__ = "permit_records"

    id = Column(String(36), primary_key=True)
    permit_title = Column(String(100), nullable=False)
    permit_type = Column(String(50), nullable=False)  # HOT_WORK, COLD_WORK, CONFINED_SPACE, HEIGHT, LOTO, ELECTRICAL
    status = Column(String(50), nullable=False)  # DRAFT, VALIDATED, APPROVED, REJECTED, SUSPENDED, CLOSED
    location = Column(String(100), nullable=False)
    applicant_id = Column(String(100), nullable=False)
    details = Column(Text, nullable=False)
    hazard_controls_json = Column(JSON, nullable=True)
    risk_score = Column(Float, nullable=False, default=5.0)
    confidence_score = Column(Float, nullable=False, default=0.95)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    approvals = relationship("PermitApprovalDb", back_populates="permit", cascade="all, delete-orphan")


class PermitApprovalDb(Base):
    __tablename__ = "permit_approvals"

    id = Column(String(36), primary_key=True)
    permit_id = Column(String(36), ForeignKey("permit_records.id", ondelete="CASCADE"), nullable=False)
    approver_id = Column(String(100), nullable=False)
    role = Column(String(50), nullable=False)  # SAFETY_OFFICER, AREA_SUPERVISOR, OPERATIONS_MANAGER
    status = Column(String(50), nullable=False, default="PENDING")  # APPROVED, REJECTED, PENDING
    comments = Column(Text, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    permit = relationship("PermitRecordDb", back_populates="approvals")
