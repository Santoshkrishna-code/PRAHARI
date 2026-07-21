import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class MaintenancePredictionDb(Base):
    __tablename__ = "maintenance_predictions"

    id = Column(String(36), primary_key=True)
    asset_id = Column(String(36), nullable=False)
    failure_probability = Column(Float, nullable=False, default=0.15)
    rul_hours = Column(Float, nullable=False, default=720.0)
    recommended_maintenance_date = Column(String(50), nullable=False)
    maintenance_strategy = Column(String(50), nullable=False, default="CONDITION_BASED")
    status = Column(String(50), nullable=False, default="RECOMMENDED")  # RECOMMENDED, SCHEDULED, COMPLETED
    spare_parts_json = Column(JSON, nullable=True)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)


class MaintenanceWorkOrderDb(Base):
    __tablename__ = "maintenance_work_orders"

    id = Column(String(36), primary_key=True)
    asset_id = Column(String(36), nullable=False)
    work_order_title = Column(String(100), nullable=False)
    priority = Column(String(20), nullable=False, default="MEDIUM")  # LOW, MEDIUM, HIGH, CRITICAL
    status = Column(String(20), nullable=False, default="OPEN")  # OPEN, IN_PROGRESS, CLOSED
    downtime_hours_est = Column(Float, nullable=False, default=4.0)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
