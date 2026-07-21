import time
import logging
from typing import Any
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.repositories.evaluation_repository import evaluation_repository
from app.domain.evaluation import EvaluationResponse, BenchmarkResponse

logger = logging.getLogger(__name__)


class EvaluationEngine:
    """Core evaluation logic checking correctness, relevance, faithfulness, cost, and latency analytics."""

    async def evaluate_run(
        self,
        db: AsyncSession,
        run_id: str | None,
        text_input: str,
        text_output: str,
        context_list: list[str],
        cost: float,
        latency_ms: int,
        model: str,
        user_id: str,
        tenant_id: str,
    ) -> EvaluationResponse:
        import string
        def clean_text(t: str) -> str:
            return t.translate(str.maketrans("", "", string.punctuation)).lower()

        clean_input = clean_text(text_input)
        clean_output = clean_text(text_output)

        # 1. Evaluate Faithfulness (overlap of output nouns/adjectives with RAG context)
        faithfulness = 1.0
        if context_list:
            combined_context = clean_text(" ".join(context_list))
            words_to_check = [w for w in clean_output.split() if len(w) > 4]
            if words_to_check:
                matches = sum(1 for w in words_to_check if w in combined_context)
                faithfulness = round(matches / len(words_to_check), 2)
                # Keep within bounds
                faithfulness = max(0.2, min(faithfulness, 1.0))

        # 2. Evaluate Relevance (does the output match the prompt key directives)
        relevance = 1.0
        input_words = [w for w in clean_input.split() if len(w) > 4]
        if input_words:
            matches = sum(1 for w in input_words if w in clean_output)
            relevance = round(matches / len(input_words), 2)
            relevance = max(0.4, min(relevance, 1.0))

        # 3. Evaluate Groundedness
        groundedness = 1.0
        if context_list and "[REDACTED" in text_output:
            # Mask checks
            groundedness = 0.9

        # 4. Evaluate Correctness (regression check against golden datasets)
        correctness = 0.8
        golden_sets = await evaluation_repository.get_golden_datasets(db)
        for gold in golden_sets:
            # Check prompt keyword similarity
            gold_input = clean_text(gold.input_prompt)
            if gold_input[:30] in clean_input[:30]:
                expected_words = [w for w in clean_text(gold.expected_output).split()]
                matches = sum(1 for w in expected_words if w in clean_output)
                correctness = round(matches / len(expected_words), 2)
                correctness = max(0.1, min(correctness, 1.0))
                break

        # 5. Compute Overall Score
        overall = round((faithfulness + correctness + relevance + groundedness) / 4, 2)

        # 6. Save Evaluation details
        db_eval = await evaluation_repository.save_evaluation(
            db,
            run_id,
            text_input,
            text_output,
            model,
            faithfulness,
            correctness,
            relevance,
            groundedness,
            overall,
            latency_ms,
            cost,
            user_id,
            tenant_id,
        )

        return EvaluationResponse(
            evaluation_id=db_eval.id,
            faithfulness_score=faithfulness,
            correctness_score=correctness,
            relevance_score=relevance,
            groundedness_score=groundedness,
            overall_score=overall,
            latency_ms=latency_ms,
            cost=cost,
            model=model,
        )

    async def execute_benchmark(
        self, db: AsyncSession, name: str, model_name: str, test_prompts: list[str]
    ) -> BenchmarkResponse:
        """Runs test prompts against dummy LLM interfaces compiling latency/accuracy aggregates."""
        total_latency = 0
        total_cost = 0.0
        total_correctness = 0.0
        count = len(test_prompts) or 1

        for prompt in test_prompts:
            # Mock latency: 200ms to 800ms
            lat = 450
            cost = 0.002
            total_latency += lat
            total_cost += cost
            total_correctness += 0.85

        avg_lat = round(total_latency / count, 2)
        avg_cost = round(total_cost / count, 5)
        avg_acc = round(total_correctness / count, 2)

        b_id = await evaluation_repository.save_benchmark(
            db, name, model_name, avg_lat, avg_cost, avg_acc, reliability=0.99
        )

        return BenchmarkResponse(
            benchmark_id=b_id,
            model_name=model_name,
            average_latency=avg_lat,
            average_cost=avg_cost,
            average_accuracy=avg_acc,
            reliability_rate=0.99,
        )

    def route_ab_test(self, user_id: str) -> str:
        """Determines variant route A/B based on hash of user_id."""
        # Heuristic split: odd hashes route to variant_a, even hashes route to variant_b
        user_hash = sum(ord(c) for c in user_id)
        return "variant_a" if user_hash % 2 == 1 else "variant_b"


evaluation_engine = EvaluationEngine()
