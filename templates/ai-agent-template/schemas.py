from typing import List, Dict, Any, TypedDict, Annotated
from pydantic import BaseModel, Field
import operator

class AgentInput(BaseModel):
    worker_id: str = Field(..., description="Unique ID of the worker requesting check")
    zone_id: str = Field(..., description="Target safety zone ID")
    worker_role: str = Field(..., description="Role of the worker")
    plant_name: str = Field("Main Plant", description="Plant name context")

class RiskFactor(BaseModel):
    hazard_type: str = Field(..., description="Type of hazard identified (e.g. PPE, GAS, HEIGHTS)")
    description: str = Field(..., description="Detailed description of the hazard risk")
    severity: str = Field(..., description="Severity level: LOW, MEDIUM, HIGH, CRITICAL")

class SafetyReport(BaseModel):
    allowed_entry: bool = Field(..., description="Whether worker is approved to enter zone")
    mitigation_steps: List[str] = Field(..., description="Actionable safety mitigations to apply")
    hazards_detected: List[RiskFactor] = Field(..., description="List of detected risks")

# LangGraph state definition
class AgentState(TypedDict):
    input: AgentInput
    safety_manual_context: str
    permit_status: str
    vitals_status: str
    hazards: List[RiskFactor]
    messages: Annotated[list, operator.add]
    final_report: SafetyReport
