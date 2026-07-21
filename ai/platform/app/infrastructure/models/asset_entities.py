import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class AssetRecordDb(Base):
    __tablename__ = "asset_records"

    id = Column(String(36), primary_key=True)
    asset_name = Column(String(100), nullable=False)
    asset_category = Column(String(50), nullable=False)  # PUMP, COMPRESSOR, BOILER, etc.
    status = Column(String(50), nullable=False, default="OPERATIONAL")  # OPERATIONAL, MAINTENANCE_REQUIRED, SHUTDOWN
    location = Column(String(100), nullable=False)
    health_score = Column(Float, nullable=False, default=100.0)
    operating_hours = Column(Float, nullable=False, default=0.0)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    logs = relationship("AssetLogDb", back_populates="asset", cascade="all, delete-orphan")


class AssetLogDb(Base):
    __tablename__ = "asset_logs"

    id = Column(String(36), primary_key=True)
    asset_id = Column(String(36), ForeignKey("asset_records.id", ondelete="CASCADE"), nullable=False)
    log_type = Column(String(50), nullable=False)  # HEALTH_CHECK, MAINTENANCE, SHUTDOWN_NOTICE
    health_score = Column(Float, nullable=False, default=100.0)
    notes = Column(Text, nullable=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    asset = relationship("AssetRecordDb", back_populates="logs")
