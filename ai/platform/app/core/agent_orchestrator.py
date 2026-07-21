import uuid
import logging
from typing import Any
from sqlalchemy.ext.asyncio import AsyncSession

from app.domain.agents import AgentRunResponse
from app.infrastructure.repositories.agent_repository import agent_repository
from app.core.tool_executor import tool_executor

logger = logging.getLogger(__name__)


class AgentOrchestrator:
    """Orchestrates multi-agent execution graphs, checkpoint state stores, and human-in-the-loop pauses."""

    async def execute_graph(
        self,
        db: AsyncSession,
        graph_id: str,
        execution_id: str | None,
        input_state: dict[str, Any],
        user_id: str,
        tenant_id: str,
        start_node_override: str | None = None,
    ) -> AgentRunResponse:
        # Resolve graph configuration
        graph = await agent_repository.get_graph(db, graph_id)
        if not graph:
            return AgentRunResponse(
                execution_id="none",
                status="FAILED",
                current_node="none",
                message=f"Graph '{graph_id}' not found.",
            )

        exec_id = execution_id or str(uuid.uuid4())
        state = input_state.copy()

        # Resolve initial node
        current_node_name = start_node_override or graph.start_node
        nodes_map = {n.name: n for n in graph.nodes}

        while current_node_name in nodes_map:
            node = nodes_map[current_node_name]
            logger.info("Executing Graph Node: %s in Execution Session: %s", node.name, exec_id)

            # Heuristic check for human-in-the-loop sign-off check
            if ("approve" in node.name.lower() or "approval" in node.name.lower()) and not state.get("reviewer_approved"):
                logger.info("Sign-off check required. Halting execution on node: %s", node.name)
                # Create pending human approval task in DB
                await agent_repository.create_human_approval(db, exec_id, node.name)
                
                # Save checkpoint state in DB
                await agent_repository.save_checkpoint(
                    db, exec_id, graph_id, node.name, state, "INTERRUPTED", user_id, tenant_id
                )
                return AgentRunResponse(
                    execution_id=exec_id,
                    status="INTERRUPTED",
                    current_node=node.name,
                    state=state,
                    message=f"Sign-off check required on node '{node.name}'. Save checkpoint and paused.",
                )

            # Run Node Logic (Stubs representing specialized workers)
            # Worker Agent Node implementations modifying graph state variables
            if "incident_classifier" in node.name.lower():
                # Heuristic: analyze incident text and determine risk
                desc = state.get("incident_desc", "")
                if "toxic" in desc.lower() or "explosion" in desc.lower():
                    state["risk_level"] = "CRITICAL"
                    state["reviewer_required"] = True
                else:
                    state["risk_level"] = "LOW"
                    state["reviewer_required"] = False
                state["incident_processed"] = True

            elif "tool_runner" in node.name.lower() and node.tools_allowed_json:
                # Triggers tools execution
                tool_to_run = node.tools_allowed_json[0]
                if tool_to_run == "check_weather":
                    res = await tool_executor.execute_tool(
                        db,
                        tool_to_run,
                        None,
                        {"latitude": 48.1, "longitude": 11.5},
                        user_id,
                        tenant_id,
                    )
                    state["weather_data"] = res.output

            # Determine next target node transition
            target_node = None
            out_edges = [e for e in graph.edges if e.source_node == node.name]

            if len(out_edges) == 1:
                target_node = out_edges[0].target_node
            elif len(out_edges) > 1:
                # Conditional Routing evaluation
                for edge in out_edges:
                    if edge.conditional_expr == "risk_is_critical":
                        if state.get("risk_level") == "CRITICAL":
                            target_node = edge.target_node
                            break
                    elif edge.conditional_expr == "risk_is_low":
                        if state.get("risk_level") != "CRITICAL":
                            target_node = edge.target_node
                            break
            
            # Transition to target node or halt completed
            if not target_node:
                break
            current_node_name = target_node

        # Save completed execution state checkpoint
        await agent_repository.save_checkpoint(
            db, exec_id, graph_id, current_node_name, state, "COMPLETED", user_id, tenant_id
        )
        return AgentRunResponse(
            execution_id=exec_id,
            status="COMPLETED",
            current_node=current_node_name,
            state=state,
            message="Graph executed to completion successfully.",
        )

    async def resume_execution(
        self, db: AsyncSession, execution_id: str, approved: bool, comments: str | None
    ) -> AgentRunResponse:
        # Load state checkpoint
        checkpoint = await agent_repository.get_checkpoint(db, execution_id)
        if not checkpoint:
            return AgentRunResponse(
                execution_id=execution_id,
                status="FAILED",
                current_node="none",
                message="Target state checkpoint session not found.",
            )

        # Resolve human approval task status
        await agent_repository.resolve_approval(db, execution_id, approved, comments)

        if not approved:
            # Audit state as failed/rejected
            checkpoint.status = "FAILED"
            await db.commit()
            return AgentRunResponse(
                execution_id=execution_id,
                status="FAILED",
                current_node=checkpoint.current_node,
                state=checkpoint.state_json,
                message="Review Sign-off rejected by supervisor. Halted execution.",
            )

        # Resume executing graph on approved node
        state = checkpoint.state_json
        state["reviewer_approved"] = True
        state["reviewer_comments"] = comments

        return await self.execute_graph(
            db,
            checkpoint.graph_id,
            execution_id,
            state,
            checkpoint.user_id,
            checkpoint.tenant_id,
            start_node_override=checkpoint.current_node,
        )


agent_orchestrator = AgentOrchestrator()
