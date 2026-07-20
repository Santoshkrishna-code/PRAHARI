from typing import Dict, Any
from schemas import AgentState, RiskFactor, SafetyReport
from tools import check_worker_permit_status, query_osha_safety_manuals
from config import settings

def retrieve_permit_status_node(state: AgentState) -> Dict[str, Any]:
    """Node to fetch the current active permit status of the target worker."""
    inp = state["input"]
    status = check_worker_permit_status.invoke({
        "worker_id": inp.worker_id,
        "zone_id": inp.zone_id
    })
    return {"permit_status": status}

def retrieve_manual_context_node(state: AgentState) -> Dict[str, Any]:
    """Node to query the Vector DB manuals for zone-related rules."""
    inp = state["input"]
    query = f"requirements for safety zone {inp.zone_id}"
    context = query_osha_safety_manuals.invoke({"query": query})
    return {"safety_manual_context": context}

def call_bedrock_llm_node(state: AgentState) -> Dict[str, Any]:
    """Node to invoke Amazon Bedrock/Claude to perform safety analysis."""
    # In production, initialize boto3 Bedrock client and invoke model.
    # For template mock representation:
    inp = state["input"]
    permit = state["permit_status"]
    context = state["safety_manual_context"]
    
    # Analyze state and form response structure
    allowed = "APPROVED" in permit
    hazards = []
    mitigations = ["Verify PPE on camera"]
    
    if not allowed:
        hazards.append(RiskFactor(
            hazard_type="PERMIT_VIOLATION",
            description="Worker has no active permit to enter hazardous area",
            severity="CRITICAL"
        ))
        mitigations.append("Block turnstile access")
        mitigations.append("Issue emergency alert to supervisor")
        
    report = SafetyReport(
        allowed_entry=allowed,
        mitigation_steps=mitigations,
        hazards_detected=hazards
    )
    
    return {"final_report": report}
