import math
import time
import json
import logging
from typing import Any
from sqlalchemy.ext.asyncio import AsyncSession

from app.core.config import settings
from app.core.models import ChatRequest, Message
from app.domain.memory import MessageDto, SessionDto, MemoryEntryDto
from app.infrastructure.cache import cache_provider
from app.infrastructure.repositories.memory_repository import memory_repository
from app.infrastructure.providers.factory import provider_factory

logger = logging.getLogger(__name__)


class MemoryOrchestrator:
    """Orchestrates Session Memory operations, caching, and Long-Term Memory scoring."""

    async def get_session_messages(self, db: AsyncSession, session_id: str) -> list[MessageDto]:
        """Fetches sliding window messages from Redis cache or falls back to DB."""
        cache_key = f"ai:messages:{session_id}"
        
        try:
            cached_val = await cache_provider.get(cache_key)
            if cached_val:
                logger.debug("Cache hit for session messages: %s", cache_key)
                msgs_data = json.loads(cached_val)
                return [MessageDto.model_validate(m) for m in msgs_data]
        except Exception as e:
            logger.warning("Session cache retrieval failed: %s", str(e))

        # Fallback to Database read
        logger.info("Cache miss. Reading message history from database for: %s", session_id)
        messages = await memory_repository.get_history(db, session_id, limit=20)
        
        # Populate Redis cache
        try:
            serialized = json.dumps([m.model_dump(mode="json") for m in messages])
            await cache_provider.set(cache_key, serialized, ttl=7200)
        except Exception as e:
            logger.warning("Failed to write messages back to cache: %s", str(e))

        return messages

    async def add_message_and_sync_cache(
        self, db: AsyncSession, session_id: str, role: str, content: str
    ) -> MessageDto:
        """Appends message to DB, invalidates/updates cache, and trigger async summary update."""
        # 1. Write message to Database
        msg = await memory_repository.add_message(db, session_id, role, content)

        # 2. Append message to Redis cache array
        cache_key = f"ai:messages:{session_id}"
        try:
            cached_val = await cache_provider.get(cache_key)
            messages = []
            if cached_val:
                messages = [MessageDto.model_validate(m) for m in json.loads(cached_val)]
            else:
                # Cache was cold, fetch current history from DB
                messages = await memory_repository.get_history(db, session_id, limit=20)
            
            # Check sliding window size limit
            if len(messages) >= 20:
                messages.pop(0)

            # Add new message and update Redis
            messages.append(msg)
            serialized = json.dumps([m.model_dump(mode="json") for m in messages])
            await cache_provider.set(cache_key, serialized, ttl=7200)
        except Exception as e:
            logger.warning("Cache write failure on message append: %s", str(e))

        return msg

    async def compile_running_summary(self, db: AsyncSession, session_id: str) -> str:
        """Invokes provider layer to generate a running summary of the session history."""
        # Load history messages
        messages = await memory_repository.get_history(db, session_id, limit=30)
        if not messages:
            return "No conversation history exists to summarize."

        history_text = "\n".join(f"{m.role}: {m.content}" for m in messages)
        prompt = f"Write a brief, concise running summary of the following EHS conversation history:\n\n{history_text}"

        # Resolve LLM via provider factory to compile summary
        logger.info("Executing summary generation for session: %s", session_id)
        chat_req = ChatRequest(
            messages=[Message(role="user", content=prompt)],
            model=settings.DEFAULT_MODEL,
            temperature=0.3,
            use_cache=False,
        )
        
        try:
            chat_res = await provider_factory.execute_chat(chat_req)
            summary_text = chat_res.choices[0].message.content

            # Persist summary to DB
            await memory_repository.update_summary(db, session_id, summary_text)
            
            # Cache summary in Redis
            sum_cache_key = f"ai:summary:{session_id}"
            await cache_provider.set(sum_cache_key, summary_text, ttl=86400)

            return summary_text
        except Exception as e:
            logger.error("Failed to generate running summary: %s", str(e))
            return f"Summary compilation failed: {e}"

    def calculate_memory_score(
        self,
        entry: Any,  # MemoryEntryDb
        query_text: str,
        query_embedding: list[float] | None = None,
    ) -> tuple[float, float, float, float]:
        """Calculates decay recency, frequency, importance, and cosine relevance scores."""
        # 1. Recency: logarithmic decay
        delta_seconds = (datetime.datetime.utcnow() - entry.timestamp).total_seconds()
        delta_hours = max(0.1, delta_seconds / 3600.0)
        recency = 1.0 / (1.0 + math.log(1.0 + delta_hours))

        # Eager load score metrics record
        scores_record = entry.scores
        frequency_factor = 0.0
        importance_factor = 0.1
        if scores_record:
            # 2. Frequency factor (cap at 5 retrievals)
            frequency_factor = min(5.0, scores_record.frequency) / 5.0
            # 3. Importance factor (normalize scale of 1-10 to 0-1)
            importance_factor = scores_record.importance / 10.0

        # 4. Relevance: Cosine similarity or keyword overlapping index
        relevance = 0.0
        if query_embedding and entry.embedding_json:
            relevance = self._cosine_similarity(query_embedding, entry.embedding_json)
        else:
            # Keyword overlapping index fallback
            query_words = set(query_text.lower().split())
            content_words = set(entry.content.lower().split())
            if query_words:
                relevance = len(query_words.intersection(content_words)) / len(query_words)

        # 5. Composite Score weighting
        composite_score = (
            (0.15 * recency)
            + (0.15 * frequency_factor)
            + (0.30 * importance_factor)
            + (0.40 * relevance)
        )

        return composite_score, recency, relevance, frequency_factor

    def _cosine_similarity(self, vec_a: list[float], vec_b: list[float]) -> float:
        """Math helper computing cosine vector similarities."""
        if not vec_a or not vec_b or len(vec_a) != len(vec_b):
            return 0.0
        dot_product = sum(a * b for a, b in zip(vec_a, vec_b))
        norm_a = math.sqrt(sum(a * a for a in vec_a))
        norm_b = math.sqrt(sum(b * b for b in vec_b))
        if norm_a == 0.0 or norm_b == 0.0:
            return 0.0
        return dot_product / (norm_a * norm_b)

    async def retrieve_relevant_memories(
        self, db: AsyncSession, user_id: str, query: str, limit: int = 5
    ) -> list[MemoryEntryDto]:
        """Scores user declarative memory database blocks and returns top matches."""
        # Check if vector credentials exist to execute embedding similarity
        query_embedding = None
        try:
            # Try to compile embedding via Provider Layer
            from app.core.models import EmbeddingRequest
            emb_req = EmbeddingRequest(input=query, model="text-embedding-3-small")
            emb_res = await provider_factory.execute_embeddings(emb_req)
            if emb_res and emb_res.data:
                query_embedding = emb_res.data[0].embedding
        except Exception:
            logger.warning("Could not calculate query embedding. Falling back to keyword search.")

        # Load all candidate memories for user
        candidates = await memory_repository.get_memories_for_user(db, user_id)
        scored_entries = []

        for entry in candidates:
            score, recency, relevance, freq = self.calculate_memory_score(
                entry, query, query_embedding
            )
            
            # Add to list mapping to DTO fields
            dto = MemoryEntryDto(
                id=entry.id,
                user_id=entry.user_id,
                category=entry.category,
                content=entry.content,
                importance=entry.importance,
                relevance=round(relevance, 4),
                recency=round(recency, 4),
                score=round(score, 4),
                timestamp=entry.timestamp,
            )
            scored_entries.append((score, dto))

            # Record frequency increment inside Repository
            await memory_repository.increment_memory_frequency(db, entry.id)

        # Sort entries descending by composite score
        scored_entries.sort(key=lambda x: x[0], reverse=True)
        return [item[1] for item in scored_entries[:limit]]


memory_orchestrator = MemoryOrchestrator()
import datetime
