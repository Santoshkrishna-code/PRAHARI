import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import selectinload
from app.infrastructure.database import engine
from app.infrastructure.models.agent_entities import (
    Base,
    AgentGraphDb,
    AgentNodeDb,
    AgentEdgeDb,
    AgentStateCheckpointDb,
    HumanApprovalDb,
)
from app.domain.agents import AgentGraphDto, AgentNodeDto, AgentEdgeDto


class AgentRepository:
    """Repository handling DB operations for Agent Graph configurations, State Checkpoints, and Human approvals."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

    async def register_graph(
        self,
        session: AsyncSession,
        name: str,
        description: str,
        start_node: str,
        nodes: list[dict],
        edges: list[dict],
    ) -> AgentGraphDto:
        # Check if graph already exists
        stmt = select(AgentGraphDb).where(AgentGraphDb.name == name)
        res = await session.execute(stmt)
        db_graph = res.scalar_one_or_none()

        if db_graph:
            # Delete old graph to clean cascades
            await session.delete(db_graph)
            await session.commit()

        graph_id = str(uuid.uuid4())
        db_graph = AgentGraphDb(
            id=graph_id,
            name=name,
            description=description,
            start_node=start_node,
        )
        session.add(db_graph)

        # Save Nodes
        node_dtos = []
        for n in nodes:
            n_id = str(uuid.uuid4())
            db_node = AgentNodeDb(
                id=n_id,
                graph_id=graph_id,
                name=n["name"],
                prompt_template_name=n.get("prompt_template_name"),
                tools_allowed_json=n.get("tools_allowed", []),
                model=n.get("model", "gpt-4o"),
            )
            session.add(db_node)
            node_dtos.append(
                AgentNodeDto(
                    id=n_id,
                    name=n["name"],
                    prompt_template_name=n.get("prompt_template_name"),
                    tools_allowed=n.get("tools_allowed", []),
                    model=n.get("model", "gpt-4o"),
                )
            )

        # Save Edges
        edge_dtos = []
        for e in edges:
            e_id = str(uuid.uuid4())
            db_edge = AgentEdgeDb(
                id=e_id,
                graph_id=graph_id,
                source_node=e["source_node"],
                target_node=e["target_node"],
                conditional_expr=e.get("conditional_expr"),
            )
            session.add(db_edge)
            edge_dtos.append(
                AgentEdgeDto(
                    id=e_id,
                    source_node=e["source_node"],
                    target_node=e["target_node"],
                    conditional_expr=e.get("conditional_expr"),
                )
            )

        await session.commit()
        return AgentGraphDto(
            id=graph_id,
            name=name,
            description=description,
            start_node=start_node,
            nodes=node_dtos,
            edges=edge_dtos,
        )

    async def get_graph(self, session: AsyncSession, graph_id: str) -> AgentGraphDb | None:
        stmt = (
            select(AgentGraphDb)
            .where(AgentGraphDb.id == graph_id)
            .options(selectinload(AgentGraphDb.nodes), selectinload(AgentGraphDb.edges))
        )
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def save_checkpoint(
        self,
        session: AsyncSession,
        execution_id: str,
        graph_id: str,
        current_node: str,
        state: dict,
        status: str,
        user_id: str,
        tenant_id: str,
    ) -> None:
        stmt = select(AgentStateCheckpointDb).where(AgentStateCheckpointDb.execution_id == execution_id)
        res = await session.execute(stmt)
        db_check = res.scalar_one_or_none()

        if db_check:
            db_check.current_node = current_node
            db_check.state_json = state
            db_check.status = status
        else:
            db_check = AgentStateCheckpointDb(
                id=str(uuid.uuid4()),
                execution_id=execution_id,
                graph_id=graph_id,
                current_node=current_node,
                state_json=state,
                status=status,
                user_id=user_id,
                tenant_id=tenant_id,
            )
            session.add(db_check)
        await session.commit()

    async def get_checkpoint(self, session: AsyncSession, execution_id: str) -> AgentStateCheckpointDb | None:
        stmt = select(AgentStateCheckpointDb).where(AgentStateCheckpointDb.execution_id == execution_id)
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def create_human_approval(self, session: AsyncSession, execution_id: str, node_name: str) -> str:
        app_id = str(uuid.uuid4())
        db_app = HumanApprovalDb(
            id=app_id,
            execution_id=execution_id,
            node_name=node_name,
            status="PENDING",
        )
        session.add(db_app)
        await session.commit()
        return app_id

    async def resolve_approval(self, session: AsyncSession, execution_id: str, approved: bool, comments: str | None) -> None:
        stmt = (
            select(HumanApprovalDb)
            .where(HumanApprovalDb.execution_id == execution_id)
            .where(HumanApprovalDb.status == "PENDING")
        )
        res = await session.execute(stmt)
        db_app = res.scalar_one_or_none()
        if db_app:
            db_app.status = "APPROVED" if approved else "REJECTED"
            db_app.comments = comments
            await session.commit()


agent_repository = AgentRepository()
