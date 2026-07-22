import uuid
import time
import asyncio
from collections.abc import AsyncGenerator
from app.infrastructure.providers.base import BaseProvider
from app.core.models import (
    ChatRequest,
    ChatResponse,
    Choice,
    Message,
    EmbeddingRequest,
    EmbeddingResponse,
    EmbeddingData,
)


class MockProvider(BaseProvider):
    """Context-aware intelligent EHS provider for offline/demo execution."""

    def get_provider_name(self) -> str:
        return "mock"

    def _generate_dynamic_ehs_reasoning(self, prompt: str) -> str:
        query_lower = prompt.lower().strip()

        # Greetings & basic conversation
        if query_lower in ["hi", "hello", "hey", "greetings", "help"]:
          return (
              "Hello! I am the PRAHARI Autonomous AI Safety Supervisor."
              " I am connected to live plant telemetry, the digital twin,"
              " knowledge graph, and the enterprise event bus. How can I assist"
              " your safety, risk, or asset investigation today?"
          )

        # Asset specific: P-102
        if "p-102" in query_lower or "pump" in query_lower:
          return (
              "AI Supervisor Asset Evaluation [PUMP-P102]:\n"
              "• Current Telemetry: Vibration velocity = 11.8 mm/s, Bearing"
              " Temp = 94.1°C.\n"
              "• Physics Twin Model: PINN neural network predicts bearing race"
              " wear at 72% probability.\n"
              "• RUL Estimate: 18 days remaining.\n"
              "• Root Cause: Work order WO-7821 was 14 days overdue,"
              " causing lubrication breakdown under thermal load.\n"
              "• Recommendation: Approve work order WO-7821 for bearing race"
              " replacement within 24 hours."
          )

        # Incidents & 5-Whys RCA
        if (
            "5-whys" in query_lower
            "rca" in query_lower
            or "root cause" in query_lower
            or "incident" in query_lower
        ):
          return (
              "AI Supervisor Incident Analysis [INC-2026-0447]:\n"
              "1. Why did Pump P-102 trigger an alarm? → Vibration velocity"
              " reached 11.8 mm/s.\n"
              "2. Why was vibration elevated? → Misalignment of the outer"
              " bearing race.\n"
              "3. Why did misalignment occur? → Progressive friction wear from"
              " thermal overheating.\n"
              "4. Why did thermal overheating happen? → Lubrication oil pressure"
              " dropped due to extended service interval.\n"
              "5. Why was service extended? → Automated work order escalation"
              " was disabled in CMMS configuration for WO-7821."
          )

        # Permits & LOTO isolation
        if (
            "permit" in query_lower
            or "ptw" in query_lower
            or "loto" in query_lower
            or "isolation" in query_lower
        ):
          return (
              "AI Supervisor Permit Verification [PTW-8902]:\n"
              "• Permit Status: Approved Hot Work on Tank T-204.\n"
              "• Isolation Status: LOTO lock active on Valve V-88 (Gate"
              " Valve).\n"
              "• Gas Test Verification: Oxygen 20.9%, H2S 0ppm, LEL 0%."
              " Verified 18 minutes ago.\n"
              "• Isolation Conflict Check: ZERO conflicts detected with"
              " adjacent Recirculation Loop DC-101."
          )

        # Vision & PPE Detections
        if (
            "vision" in query_lower
            or "camera" in query_lower
            or "ppe" in query_lower
            or "helmet" in query_lower
        ):
          return (
              "AI Supervisor Vision Intelligence [CAM-002 / AGX-04]:\n"
              "• Location: Restricted Zone B (Reactor Complex North).\n"
              "• Frame Rate & Latency: 29.8 FPS • 14ms latency.\n"
              "• Detection Event: Hardhat violation detected (TRK-9904, 96.4%"
              " confidence).\n"
              "• Security Action: Contractor badge C-4412 expired 3 days ago."
              " Gate B access automatically revoked."
          )

        # Risk & Compliance
        if (
            "risk" in query_lower
            or "compliance" in query_lower
            or "iso" in query_lower
            or "osha" in query_lower
        ):
          return (
              "AI Supervisor Risk & Compliance Assessment:\n"
              "• Plant Risk Score: 18/25 (HIGH) centered on Recirculation Loop"
              " DC-101.\n"
              "• ISO 45001 Status: Audit trail fully logged across Event Bus"
              " (47 trace events).\n"
              "• Leading Indicators: Plant Safety Index = 94.2/100, TRIR ="
              " 0.42, MTBF = 2,140 hrs."
          )

        # Default dynamic response echoing user prompt context
        return (
            f"AI Supervisor Platform Analysis for query: '{prompt}'\n"
            "• Connected Pipeline: Query executed across PostgreSQL, Redis"
            " event cache, and Knowledge Graph.\n"
            "• Operational Context: Plant Alpha (Gulf Coast) — 1,284 signals"
            " normal, 2 active risks, 0 open incidents.\n"
            "• Recommendation: All critical parameters verified against ISO"
            " 45001 safety guidelines."
        )

    async def generate_chat(self, request: ChatRequest) -> ChatResponse:
      start_time = time.perf_counter()
      await asyncio.sleep(0.02)

      model = request.model or "gpt-4o"
      prompt_content = request.messages[-1].content if request.messages else ""
      completion_content = self._generate_dynamic_ehs_reasoning(prompt_content)

      usage = self.estimate_usage(prompt_content, completion_content, model)
      latency = (time.perf_counter() - start_time) * 1000

      return ChatResponse(
          id=f"chatcmpl-{uuid.uuid4()}",
          model=model,
          provider="mock",
          choices=[
              Choice(
                  index=0,
                  message=Message(
                      role="assistant", content=completion_content
                  ),
                  finish_reason="stop",
              )
          ],
          usage=usage,
          latency_ms=round(latency, 2),
      )

    async def generate_stream(
        self, request: ChatRequest
    ) -> AsyncGenerator[ChatResponse, None]:
      model = request.model or "gpt-4o"
      prompt_content = request.messages[-1].content if request.messages else ""
      full_response = self._generate_dynamic_ehs_reasoning(prompt_content)

      chunks = full_response.split(" ")
      accumulated = ""
      for idx, chunk in enumerate(chunks):
        word = (" " if idx > 0 else "") + chunk
        accumulated += word
        await asyncio.sleep(0.01)
        usage = self.estimate_usage(prompt_content, accumulated, model)

        yield ChatResponse(
            id=f"chatcmpl-{uuid.uuid4()}",
            model=model,
            provider="mock",
            choices=[
                Choice(
                    index=0,
                    message=Message(role="assistant", content=word),
                    finish_reason=(
                        "stop" if idx == len(chunks) - 1 else "length"
                    ),
                )
            ],
            usage=usage,
            latency_ms=12.0,
        )
