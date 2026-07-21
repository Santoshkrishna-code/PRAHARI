import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class InspectionRecordDb(Base):
    __tablename__ = "inspection_records"

    id = Column(String(36), primary_key=True)
    title = Column(String(100), nullable=False)
    inspection_type = Column(String(50), nullable=False)  # DAILY_SAFETY, EQUIPMENT, FIRE_SAFETY, etc.
    status = Column(String(50), nullable=False, default="PLANNED")  # PLANNED, IN_PROGRESS, COMPLETED
    location = Column(String(100), nullable=False)
    inspector_id = Column(String(100), nullable=False)
    score = Column(Float, nullable=False, default=100.0)
    findings_count = Column(Integer, nullable=False, default=0)
    summary_notes = Column(Text, nullable=True)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    findings = relationship("InspectionFindingDb", back_populates="inspection", cascade="all, delete-orphan")


class InspectionFindingDb(Base):
    __tablename__ = "inspection_findings"

    id = Column(String(36), primary_key=True)
    inspection_id = Column(String(36), ForeignKey("inspection_records.id", ondelete="CASCADE"), nullable=False)
    finding_type = Column(String(50), nullable=False)  # UNSAFE_ACT, UNSAFE_CONDITION, MAJOR_NON_CONFORMANCE, etc.
    severity = Column(String(20), nullable=False, default="MEDIUM")  # LOW, MEDIUM, HIGH, CRITICAL
    details = Column(Text, nullable=False)
    corrective_action = Column(Text, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    inspection = relationship("InspectionRecordDb", back_populates="findings")
