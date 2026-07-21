import datetime
from sqlalchemy import Column, String, Text, DateTime, ForeignKey, Boolean, Float
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class PolicyDb(Base):
    __tablename__ = "guardrail_policies"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False)
    code = Column(String(50), nullable=False, unique=True)
    description = Column(Text, nullable=True)
    action = Column(String(20), nullable=False, default="BLOCK")  # BLOCK, WARN, LOG
    status = Column(String(20), nullable=False, default="ACTIVE")  # ACTIVE, DISABLED
    created_at = Column(DateTime, default=datetime.datetime.utcnow)


class SafetyEventDb(Base):
    __tablename__ = "guardrail_safety_events"

    id = Column(String(36), primary_key=True)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    input_text = Column(Text, nullable=False)
    output_text = Column(Text, nullable=True)
    action_taken = Column(String(20), nullable=False)  # ALLOWED, BLOCKED, FLAGGED
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)

    violations = relationship("ViolationDb", back_populates="event", cascade="all, delete-orphan")
    pii_events = relationship("PiiEventDb", back_populates="event", cascade="all, delete-orphan")
    security_events = relationship("SecurityEventDb", back_populates="event", cascade="all, delete-orphan")
    reviews = relationship("HumanReviewDb", back_populates="event", cascade="all, delete-orphan")


class ViolationDb(Base):
    __tablename__ = "guardrail_violations"

    id = Column(String(36), primary_key=True)
    event_id = Column(String(36), ForeignKey("guardrail_safety_events.id", ondelete="CASCADE"), nullable=False)
    policy_id = Column(String(36), ForeignKey("guardrail_policies.id", ondelete="SET NULL"), nullable=True)
    rule_triggered = Column(String(100), nullable=False)
    severity = Column(String(20), nullable=False, default="MEDIUM")  # LOW, MEDIUM, HIGH
    details = Column(Text, nullable=True)

    event = relationship("SafetyEventDb", back_populates="violations")


class PiiEventDb(Base):
    __tablename__ = "guardrail_pii_events"

    id = Column(String(36), primary_key=True)
    event_id = Column(String(36), ForeignKey("guardrail_safety_events.id", ondelete="CASCADE"), nullable=False)
    pii_type = Column(String(50), nullable=False)  # EMAIL, PHONE, PAN, etc.
    masked_value = Column(String(255), nullable=False)

    event = relationship("SafetyEventDb", back_populates="pii_events")


class SecurityEventDb(Base):
    __tablename__ = "guardrail_security_events"

    id = Column(String(36), primary_key=True)
    event_id = Column(String(36), ForeignKey("guardrail_safety_events.id", ondelete="CASCADE"), nullable=False)
    security_type = Column(String(50), nullable=False)  # API_KEY, PASSWORD, JWT
    secret_detected = Column(String(255), nullable=False)

    event = relationship("SafetyEventDb", back_populates="security_events")


class HumanReviewDb(Base):
    __tablename__ = "guardrail_human_reviews"

    id = Column(String(36), primary_key=True)
    event_id = Column(String(36), ForeignKey("guardrail_safety_events.id", ondelete="CASCADE"), nullable=False)
    status = Column(String(20), nullable=False, default="PENDING")  # PENDING, APPROVED, REJECTED
    reviewer_id = Column(String(100), nullable=True)
    override_approved = Column(Boolean, nullable=True)
    notes = Column(Text, nullable=True)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)

    event = relationship("SafetyEventDb", back_populates="reviews")
