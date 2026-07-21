import time
import math
import uuid
import logging
from typing import Any
from sqlalchemy.ext.asyncio import AsyncSession

from app.core.config import settings
from app.core.models import EmbeddingRequest, ChatRequest, Message
from app.core.prompt_renderer import prompt_renderer
from app.core.prompt_validator import prompt_validator
from app.domain.knowledge import ChunkDto, CitationDto, QueryResponse
from app.infrastructure.models.knowledge_entities import ChunkDb
from app.infrastructure.repositories.knowledge_repository import knowledge_repository
from app.infrastructure.providers.factory import provider_factory
from app.core.chunker import RecursiveCharacterTextSplitter

logger = logging.getLogger(__name__)


class RagOrchestrator:
    """Manages text splitters ingestion, RRF hybrid retrieval rankings, and inline citations generation."""

    async def ingest_and_index_document(
        self,
        db: AsyncSession,
        doc_id: str,
        filename: str,
        text_content: str,
        tenant_id: str,
        department: str | None = None,
        project: str | None = None,
    ) -> None:
        """Splits raw file text, generates vector embeddings, and indexes segments."""
        # 1. Initialize text chunker splitter
        splitter = RecursiveCharacterTextSplitter(chunk_size=600, chunk_overlap=80)
        text_chunks = splitter.split_text(text_content)
        
        if not text_chunks:
            await knowledge_repository.update_doc_status(db, doc_id, "FAILED")
            return

        db_chunks = []
        for idx, text in enumerate(text_chunks):
            embedding = [0.0] * 1536  # Default fallback vector coordinates size
            try:
                # 2. Query provider layer for chunk text vector embedding
                emb_req = EmbeddingRequest(input=text, model="text-embedding-3-small")
                emb_res = await provider_factory.execute_embeddings(emb_req)
                if emb_res and emb_res.data:
                    embedding = emb_res.data[0].embedding
            except Exception as e:
                logger.warning("Failed to generate embedding for chunk %d: %s", idx, str(e))

            chunk_id = str(uuid.uuid4())
            db_chunks.append(
                ChunkDb(
                    id=chunk_id,
                    document_id=doc_id,
                    chunk_index=idx,
                    content=text,
                    page_number=(idx // 3) + 1,  # Mock heuristic page mapping
                    section_title=f"Section {idx + 1}",
                    token_count=len(text) // 4,
                    embedding_json=embedding,
                )
            )

        # 3. Save chunk records and update document status
        await knowledge_repository.save_chunks(db, db_chunks)
        await knowledge_repository.update_doc_status(db, doc_id, "INDEXED")

    async def execute_hybrid_search(
        self,
        db: AsyncSession,
        tenant_id: str,
        query: str,
        limit: int = 5,
        department: str | None = None,
        project: str | None = None,
    ) -> list[ChunkDto]:
        """Runs RRF (Reciprocal Rank Fusion) over semantic cosine rankings and BM25 matches."""
        # 1. Fetch query embedding coordinates
        query_embedding = None
        try:
            emb_req = EmbeddingRequest(input=query, model="text-embedding-3-small")
            emb_res = await provider_factory.execute_embeddings(emb_req)
            if emb_res and emb_res.data:
                query_embedding = emb_res.data[0].embedding
        except Exception as e:
            logger.warning("Query embedding generation failed: %s. Falling back to keyword search.", str(e))

        # 2. Get permission-filtered candidate chunks from DB
        candidates = await knowledge_repository.get_chunks_for_retrieval(
            db, tenant_id, department, project
        )
        if not candidates:
            return []

        # 3. Calculate separate similarity scores
        semantic_ranks = []
        keyword_ranks = []

        query_words = set(query.lower().split())

        for chunk in candidates:
            # Semantic search rank coordinate similarity
            similarity = 0.0
            if query_embedding and chunk.embedding_json:
                similarity = self._cosine_similarity(query_embedding, chunk.embedding_json)
            semantic_ranks.append((similarity, chunk))

            # BM25 keyword overlap fallback rank
            overlap = 0.0
            if query_words:
                content_words = set(chunk.content.lower().split())
                overlap = len(query_words.intersection(content_words)) / len(query_words)
            keyword_ranks.append((overlap, chunk))

        # Sort candidate lists to determine positions
        semantic_ranks.sort(key=lambda x: x[0], reverse=True)
        keyword_ranks.sort(key=lambda x: x[0], reverse=True)

        # Map ranks
        sem_rank_map = {item[1].id: idx + 1 for idx, item in enumerate(semantic_ranks)}
        key_rank_map = {item[1].id: idx + 1 for idx, item in enumerate(keyword_ranks)}

        # 4. Compute RRF Scores
        # Formula: RRF(d) = 1 / (60 + semantic_rank) + 1 / (60 + keyword_rank)
        rrf_results = []
        for chunk in candidates:
            r_sem = sem_rank_map[chunk.id]
            r_key = key_rank_map[chunk.id]
            
            rrf_score = (1.0 / (60.0 + r_sem)) + (1.0 / (60.0 + r_key))
            
            # Map chunk Db to DTO fields
            dto = ChunkDto(
                id=chunk.id,
                document_id=chunk.document_id,
                chunk_index=chunk.chunk_index,
                content=chunk.content,
                page_number=chunk.page_number,
                section_title=chunk.section_title,
                score=round(rrf_score, 6),
            )
            rrf_results.append((rrf_score, dto))

        # Sort descending by RRF score
        rrf_results.sort(key=lambda x: x[0], reverse=True)
        return [item[1] for item in rrf_results[:limit]]

    async def execute_grounded_query(
        self,
        db: AsyncSession,
        tenant_id: str,
        query: str,
        limit: int = 5,
        department: str | None = None,
        project: str | None = None,
    ) -> QueryResponse:
        """Retrieves contexts, injects into instructions prompts, and returns answers with citations."""
        start_time = time.time()
        
        # 1. Retrieve RRF ranked hybrid chunks
        chunks = await self.execute_hybrid_search(
            db, tenant_id, query, limit, department, project
        )

        if not chunks:
            # Log zero count execution and return empty
            duration = int((time.time() - start_time) * 1000)
            await knowledge_repository.log_retrieval(db, query, duration, 0)
            return QueryResponse(
                answer="I could not find any source documents in the knowledge repository to answer your query.",
                citations=[],
            )

        # 2. Build context blocks and citation schemas mappings
        context_blocks = []
        citations = []
        
        for idx, chunk in enumerate(chunks):
            # Load filename by doc ID
            doc = await knowledge_repository.get_document(db, chunk.document_id)
            doc_name = doc.filename if doc else f"Doc-{chunk.document_id[:6]}"

            # Add source chunk definition
            ref_label = f"[{doc_name}#Chunk-{chunk.id[:4]}]"
            context_blocks.append(
                f"SOURCE CHUNK ID: {ref_label}\n"
                f"DOCUMENT: {doc_name} (Page {chunk.page_number}, {chunk.section_title})\n"
                f"CONTENT:\n{chunk.content}\n"
                f"----------------------------------------"
            )

            citations.append(
                CitationDto(
                    document_name=doc_name,
                    page_number=chunk.page_number,
                    section_title=chunk.section_title,
                    chunk_id=chunk.id,
                    score=chunk.score,
                )
            )

        context_string = "\n".join(context_blocks)
        
        # 3. Assemble grounded prompt template instructions
        system_instructions = (
            "You are a helpful safety advisor. Answer the user query based ONLY on the provided source documents chunks.\n"
            "Ground your responses fully. Next to statements/facts you fetch, insert inline citation matching source chunk ID (e.g. [Filename#Chunk-xxxx]).\n"
            "If the provided context does not hold the answer to the query, respond that context information is insufficient."
        )
        
        user_content = (
            f"CONTEXT SOURCE CHUNKS:\n"
            f"{context_string}\n\n"
            f"USER QUERY: {query}\n"
            f"GROUNDED ANSWER:"
        )

        # 4. Invoke LLM provider
        chat_req = ChatRequest(
            messages=[
                Message(role="system", content=system_instructions),
                Message(role="user", content=user_content),
            ],
            model=settings.DEFAULT_MODEL,
            temperature=0.2,
        )

        try:
            chat_res = await provider_factory.execute_chat(chat_req)
            answer_text = chat_res.choices[0].message.content
        except Exception as e:
            logger.error("LLM RAG completion failure: %s", str(e))
            answer_text = f"Failed to execute LLM completion: {e}"

        # 5. Log metric stats to analytics table
        duration = int((time.time() - start_time) * 1000)
        await knowledge_repository.log_retrieval(db, query, duration, len(chunks))

        return QueryResponse(answer=answer_text, citations=citations)

    def _cosine_similarity(self, vec_a: list[float], vec_b: list[float]) -> float:
        """Math helper computing cosine similarity values."""
        if not vec_a or not vec_b or len(vec_a) != len(vec_b):
            return 0.0
        dot_product = sum(a * b for a, b in zip(vec_a, vec_b))
        norm_a = math.sqrt(sum(a * a for a in vec_a))
        norm_b = math.sqrt(sum(b * b for b in vec_b))
        if norm_a == 0.0 or norm_b == 0.0:
            return 0.0
        return dot_product / (norm_a * norm_b)


rag_orchestrator = RagOrchestrator()
