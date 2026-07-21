import datetime
from sqlalchemy import Column, String, Text, DateTime, Float, ForeignKey, Integer, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class EmergencyRecordDb(Base):
    __tablename__ = "emergency_records"

    id = Column(String(36), primary_key=True)
    title = Column(String(100), nullable=False)
    emergency_type = Column(String(50), nullable=False)  # FIRE, EXPLOSION, CHEMICAL_SPILL, etc.
    severity_level = Column(String(50), nullable=False, default="LEVEL_1")  # LEVEL_1, LEVEL_2, LEVEL_3_CRITICAL
    status = Column(String(50), nullable=False, default="ACTIVE")  # ACTIVE, STABILIZED, RECOVERED, CLOSED
    location = Column(String(100), nullable=False)
    affected_zone = Column(String(100), nullable=False)
    evacuation_route = Column(String(200), nullable=False)
    details = Column(Text, nullable=False)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)


class EmergencyResourceDb(Base):
    __tablename__ = "emergency_resources"

    id = Column(String(36), primary_key=True)
    resource_name = Column(String(100), nullable=False)
    resource_type = Column(String(50), nullable=False)  # FIRE_EXTINGUISHER, AMBULANCE, SPILL_KIT, SCBA
    quantity = Column(Integer, nullable=False, default=1)
    location = Column(String(100), nullable=False)
    status = Column(String(20), nullable=False, default="AVAILABLE")  # AVAILABLE, DISPATCHED
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
