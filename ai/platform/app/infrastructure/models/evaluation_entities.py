import datetime
from sqlalchemy import Column, String, Text, DateTime, ForeignKey, Integer, Float, JSON
from sqlalchemy.orm import DeclarativeBase, relationship


class Base(DeclarativeBase):
    pass


class EvaluationDb(Base):
    __tablename__ = "ai_evaluations"

    id = Column(String(36), primary_key=True)
    run_id = Column(String(36), nullable=True)
    text_input = Column(Text, nullable=False)
    text_output = Column(Text, nullable=False)
    model = Column(String(50), nullable=False)
    faithfulness_score = Column(Float, nullable=False)
    correctness_score = Column(Float, nullable=False)
    relevance_score = Column(Float, nullable=False)
    groundedness_score = Column(Float, nullable=False)
    overall_score = Column(Float, nullable=False)
    latency_ms = Column(Integer, nullable=False)
    cost = Column(Float, nullable=False)
    user_id = Column(String(100), nullable=False)
    tenant_id = Column(String(36), nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)

    feedbacks = relationship("FeedbackDb", back_populates="evaluation", cascade="all, delete-orphan")


class BenchmarkDb(Base):
    __tablename__ = "ai_benchmarks"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False)
    model_name = Column(String(50), nullable=False)
    average_latency = Column(Float, nullable=False)
    average_cost = Column(Float, nullable=False)
    average_accuracy = Column(Float, nullable=False)
    reliability_rate = Column(Float, nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)


class ExperimentDb(Base):
    __tablename__ = "ai_experiments"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False)
    description = Column(Text, nullable=True)
    status = Column(String(20), nullable=False, default="ACTIVE")  # ACTIVE, COMPLETED
    variant_a_config_json = Column(JSON, nullable=False)
    variant_b_config_json = Column(JSON, nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)


class GoldenDatasetDb(Base):
    __tablename__ = "ai_golden_datasets"

    id = Column(String(36), primary_key=True)
    name = Column(String(100), nullable=False)
    input_prompt = Column(Text, nullable=False)
    expected_output = Column(Text, nullable=False)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)


class FeedbackDb(Base):
    __tablename__ = "ai_human_feedbacks"

    id = Column(String(36), primary_key=True)
    evaluation_id = Column(String(36), ForeignKey("ai_evaluations.id", ondelete="CASCADE"), nullable=False)
    rating = Column(Integer, nullable=False)  # 1 to 5 scale
    user_comments = Column(Text, nullable=True)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)

    evaluation = relationship("EvaluationDb", back_populates="feedbacks")
