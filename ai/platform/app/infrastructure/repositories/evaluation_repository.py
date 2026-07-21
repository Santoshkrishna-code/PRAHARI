import uuid
import datetime
from sqlalchemy import select, func
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import engine
from app.infrastructure.models.evaluation_entities import (
    Base,
    EvaluationDb,
    BenchmarkDb,
    ExperimentDb,
    GoldenDatasetDb,
    FeedbackDb,
)
from app.domain.evaluation import LeaderboardEntryDto


class EvaluationRepository:
    """Repository handling Database operations for AI Quality Evaluations, comparative Benchmarks, A/B Testing, and feedback logs."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context and seed default Golden regression dataset on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

        # Seed default golden datasets
        async with AsyncSession(engine) as session:
            await self._seed_golden_datasets(session)

    async def _seed_golden_datasets(self, session: AsyncSession) -> None:
        defaults = [
            (
                "Hazardous Gas Safety Check",
                "What are the core EHS guidelines for hazardous gas storage?",
                "Install gas leak sensors, maintain 15-meter buffer zone, and ensure employee ventilation suits.",
            ),
            (
                "High Risk Permit Signoff",
                "What is the procedure for an high-risk work permit approval?",
                "Perform safety risk assessment, verify permit reviewer role sign-off, and log to audit database.",
            ),
        ]
        for name, prompt, expected in defaults:
            stmt = select(GoldenDatasetDb).where(GoldenDatasetDb.name == name)
            res = await session.execute(stmt)
            if not res.scalar_one_or_none():
                db_gold = GoldenDatasetDb(
                    id=str(uuid.uuid4()),
                    name=name,
                    input_prompt=prompt,
                    expected_output=expected,
                )
                session.add(db_gold)
        await session.commit()

    async def save_evaluation(
        self,
        session: AsyncSession,
        run_id: str | None,
        text_input: str,
        text_output: str,
        model: str,
        faithfulness: float,
        correctness: float,
        relevance: float,
        groundedness: float,
        overall: float,
        latency_ms: int,
        cost: float,
        user_id: str,
        tenant_id: str,
    ) -> EvaluationDb:
        db_eval = EvaluationDb(
            id=str(uuid.uuid4()),
            run_id=run_id or str(uuid.uuid4()),
            text_input=text_input,
            text_output=text_output,
            model=model,
            faithfulness_score=faithfulness,
            correctness_score=correctness,
            relevance_score=relevance,
            groundedness_score=groundedness,
            overall_score=overall,
            latency_ms=latency_ms,
            cost=cost,
            user_id=user_id,
            tenant_id=tenant_id,
        )
        session.add(db_eval)
        await session.commit()
        return db_eval

    async def get_evaluations(self, session: AsyncSession, limit: int = 50) -> list[EvaluationDb]:
        stmt = select(EvaluationDb).order_by(EvaluationDb.created_at.desc()).limit(limit)
        res = await session.execute(stmt)
        return list(res.scalars().all())

    async def save_benchmark(
        self,
        session: AsyncSession,
        name: str,
        model_name: str,
        avg_latency: float,
        avg_cost: float,
        avg_accuracy: float,
        reliability: float,
    ) -> str:
        b_id = str(uuid.uuid4())
        db_bench = BenchmarkDb(
            id=b_id,
            name=name,
            model_name=model_name,
            average_latency=avg_latency,
            average_cost=avg_cost,
            average_accuracy=avg_accuracy,
            reliability_rate=reliability,
        )
        session.add(db_bench)
        await session.commit()
        return b_id

    async def get_benchmarks(self, session: AsyncSession) -> list[BenchmarkDb]:
        stmt = select(BenchmarkDb).order_by(BenchmarkDb.created_at.desc())
        res = await session.execute(stmt)
        return list(res.scalars().all())

    async def create_experiment(
        self, session: AsyncSession, name: str, description: str | None, var_a: dict, var_b: dict
    ) -> ExperimentDb:
        e_id = str(uuid.uuid4())
        db_exp = ExperimentDb(
            id=e_id,
            name=name,
            description=description,
            status="ACTIVE",
            variant_a_config_json=var_a,
            variant_b_config_json=var_b,
        )
        session.add(db_exp)
        await session.commit()
        return db_exp

    async def get_experiments(self, session: AsyncSession) -> list[ExperimentDb]:
        stmt = select(ExperimentDb).order_by(ExperimentDb.created_at.desc())
        res = await session.execute(stmt)
        return list(res.scalars().all())

    async def save_feedback(self, session: AsyncSession, evaluation_id: str, rating: int, comments: str | None) -> None:
        f_id = str(uuid.uuid4())
        db_feed = FeedbackDb(
            id=f_id,
            evaluation_id=evaluation_id,
            rating=rating,
            user_comments=comments,
        )
        session.add(db_feed)
        await session.commit()

    async def get_golden_datasets(self, session: AsyncSession) -> list[GoldenDatasetDb]:
        stmt = select(GoldenDatasetDb)
        res = await session.execute(stmt)
        return list(res.scalars().all())

    async def get_leaderboard(self, session: AsyncSession) -> list[LeaderboardEntryDto]:
        """Aggregates scores from evaluations to compile model leaderboard rankings."""
        stmt = (
            select(
                EvaluationDb.model,
                func.avg(EvaluationDb.latency_ms).label("avg_latency"),
                func.avg(EvaluationDb.cost).label("avg_cost"),
                func.avg(EvaluationDb.overall_score).label("avg_score"),
            )
            .group_by(EvaluationDb.model)
        )
        res = await session.execute(stmt)
        rows = res.all()
        
        entries = []
        for r in rows:
            # Score formula: combination of correctness - latency scaling
            score_metric = float(r.avg_score or 0.0) * 100
            entries.append(
                LeaderboardEntryDto(
                    model_name=r.model,
                    average_latency=float(r.avg_latency or 0.0),
                    average_cost=float(r.avg_cost or 0.0),
                    average_accuracy=float(r.avg_score or 0.0),
                    reliability_rate=0.98,  # Default heuristic reliability rate
                    score=round(score_metric, 2),
                )
            )
        # Sort by overall score descending
        entries.sort(key=lambda x: x.score, reverse=True)
        return entries


evaluation_repository = EvaluationRepository()
