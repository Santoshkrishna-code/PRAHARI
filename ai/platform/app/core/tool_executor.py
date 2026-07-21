import time
import httpx
import logging
from typing import Any
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.repositories.tool_repository import tool_repository
from app.domain.tools import ToolExecuteResponse
from app.core.exceptions import ValidationException

logger = logging.getLogger(__name__)


class ToolExecutor:
    """Invoker executing target calls (REST / Database / Internal services) in sandbox."""

    async def validate_payload(self, schema: dict[str, Any], payload: dict[str, Any]) -> None:
        """Heuristic validator checking input keys against json-schema requirement blocks."""
        required = schema.get("required", [])
        for req in required:
            if req not in payload:
                raise ValidationException(f"Missing required parameter '{req}' in execution payload.")

        properties = schema.get("properties", {})
        for key, val in payload.items():
            if key not in properties:
                logger.warning("Payload key '%s' not defined in tool properties schema.", key)

    async def execute_tool(
        self,
        db: AsyncSession,
        tool_name: str,
        version_str: str | None,
        payload: dict[str, Any],
        user_id: str,
        tenant_id: str,
    ) -> ToolExecuteResponse:
        start_time = time.time()
        
        # 1. Resolve Tool metadata record
        tool = await tool_repository.get_tool_by_name(db, tool_name)
        if not tool:
            return ToolExecuteResponse(
                execution_id="none",
                success=False,
                duration_ms=0,
                error=f"Tool '{tool_name}' is not registered in the database.",
            )

        # 2. Resolve version details
        version_string = version_str or tool.active_version_string
        version = await tool_repository.get_version_by_string(db, tool.id, version_string)
        if not version:
            return ToolExecuteResponse(
                execution_id="none",
                success=False,
                duration_ms=0,
                error=f"Tool version '{version_string}' is not registered.",
            )

        # 3. Validate Inputs against schema
        try:
            await self.validate_payload(version.input_schema_json, payload)
        except Exception as e:
            return ToolExecuteResponse(
                execution_id="none",
                success=False,
                duration_ms=0,
                error=f"Validation failed: {e}",
            )

        # 4. Route execution by tool type
        output = None
        success = True
        error_msg = None

        try:
            if tool.type == "REST":
                output = await self._execute_rest_tool(version.execution_target, payload, tool.timeout_seconds)
            elif tool.type == "INTERNAL":
                output = await self._execute_internal_tool(version.execution_target, payload)
            elif tool.type == "SQL":
                output = await self._execute_sql_tool(version.execution_target, payload)
            else:
                raise ValueError(f"Unsupported tool execution type: {tool.type}")
        except Exception as e:
            success = False
            error_msg = str(e)
            logger.error("Execution of tool '%s' failed: %s", tool_name, error_msg)

        duration = int((time.time() - start_time) * 1000)

        # 5. Log execution details to Database and Audits
        status = "SUCCESS" if success else "FAILED"
        exec_id = await tool_repository.log_execution(
            db, tool.id, version.id, user_id, tenant_id, payload, output, status, duration
        )

        return ToolExecuteResponse(
            execution_id=exec_id,
            success=success,
            output=output,
            duration_ms=duration,
            error=error_msg,
        )

    async def _execute_rest_tool(self, target_url: str, payload: dict[str, Any], timeout: int) -> Any:
        """Triggers REST endpoint using HTTP POST."""
        async with httpx.AsyncClient(timeout=timeout) as client:
            # Check if mock environment
            if "mock-api" in target_url:
                return {"status": "MOCKED_REST_RESPONSE", "input_received": payload}
            
            res = await client.post(target_url, json=payload)
            res.raise_for_status()
            return res.json()

    async def _execute_internal_tool(self, service_target: str, payload: dict[str, Any]) -> Any:
        """Executes stub microservice actions for digital twins, incident reporting and permits."""
        if service_target == "create_incident":
            return {
                "incident_id": f"INC-{int(time.time())}",
                "status": "OPEN",
                "reporter_notes": payload.get("description", "Safety breach"),
                "site_location": payload.get("location", "Sector 1"),
                "created_at": str(time.time()),
            }
        elif service_target == "approve_permit":
            return {
                "permit_id": payload.get("permit_id", f"PRM-{int(time.time())}"),
                "approved": True,
                "approver": "mcp-service-agent",
                "expiry": "24h",
            }
        elif service_target == "query_digital_twin":
            return {
                "equipment_id": payload.get("equipment_id", "EQ-001"),
                "current_pressure": "4.2 bar",
                "historical_drift": "none",
                "coordinates": {"lat": 48.1351, "lon": 11.582},
            }
        return {"status": "INTERNAL_TRIGGERED", "target": service_target, "payload": payload}

    async def _execute_sql_tool(self, sql_template: str, payload: dict[str, Any]) -> Any:
        """Simulates read-only database query parameters mappings."""
        # Clean query: block unsafe keywords
        query_upper = sql_template.upper()
        for unsafe in ["DELETE", "DROP", "UPDATE", "INSERT", "ALTER", "TRUNCATE"]:
            if unsafe in query_upper:
                raise PermissionError(f"SQL Sandbox Violation: Query uses forbidden keyword '{unsafe}'")

        return {
            "query": sql_template,
            "params": payload,
            "columns": ["id", "status", "risk_level"],
            "rows": [
                [1, "OPEN", "HIGH"],
                [2, "RESOLVED", "LOW"],
            ],
            "count": 2,
        }


tool_executor = ToolExecutor()
