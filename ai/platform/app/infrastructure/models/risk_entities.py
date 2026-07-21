import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class RiskAssessmentRecordDb(Base):
    __tablename__ = "risk_assessments"

    id = Column(String(36), primary_key=True)
    title = Column(String(100), nullable=False)
    risk_type = Column(String(50), nullable=False)  # OCCUPATIONAL, FIRE, CHEMICAL, ELECTRICAL, etc.
    methodology = Column(String(50), nullable=False)  # HIRA, JSA, BOWTIE, FMEA, HAZOP
    initial_risk_score = Column(Float, nullable=False, default=15.0)
    residual_risk_score = Column(Float, nullable=False, default=5.0)
    risk_category = Column(String(20), nullable=False)  # LOW, MEDIUM, HIGH, CRITICAL
    hazards_json = Column(JSON, nullable=True)
    recommended_controls_json = Column(JSON, nullable=True)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    register_entries = relationship("RiskRegisterDb", back_populates="assessment", cascade="all, delete-orphan")


class RiskRegisterDb(Base):
    __tablename__ = "risk_register"

    id = Column(String(36), primary_key=True)
    assessment_id = Column(String(36), ForeignKey("risk_assessments.id", ondelete="CASCADE"), nullable=False)
    hazard_name = Column(String(100), nullable=False)
    severity_score = Column(Integer, nullable=False)
    likelihood_score = Column(Integer, nullable=False)
    trend_status = Column(String(20), nullable=False, default="STABLE")  # STABLE, INCREASING, DECREASING
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    assessment = relationship("RiskAssessmentRecordDb", back_populates="register_entries")
