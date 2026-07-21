import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class ExecutiveReportDb(Base):
    __tablename__ = "executive_reports"

    id = Column(String(36), primary_key=True)
    report_title = Column(String(100), nullable=False)
    report_type = Column(String(50), nullable=False)  # BOARD_REPORT, EXECUTIVE_SUMMARY, etc.
    executive_summary_markdown = Column(Text, nullable=False)
    kpis_json = Column(JSON, nullable=True)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)


class AnalyticsKpiSnapshotDb(Base):
    __tablename__ = "analytics_kpi_snapshots"

    id = Column(String(36), primary_key=True)
    tenant_id = Column(String(36), nullable=False)
    trir = Column(Float, nullable=False, default=1.2)
    ltifr = Column(Float, nullable=False, default=0.4)
    permit_compliance_rate = Column(Float, nullable=False, default=96.5)
    inspection_completion_rate = Column(Float, nullable=False, default=94.0)
    asset_availability_rate = Column(Float, nullable=False, default=98.2)
    compliance_score = Column(Float, nullable=False, default=92.0)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
