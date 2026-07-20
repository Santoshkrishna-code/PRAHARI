from langgraph.graph import StateGraph, END
from schemas import AgentState
from nodes import (
    retrieve_permit_status_node,
    retrieve_manual_context_node,
    call_bedrock_llm_node
)

def create_agent_graph() -> StateGraph:
    # 1. Initialize State Graph
    workflow = StateGraph(AgentState)
    
    # 2. Add Nodes
    workflow.add_node("retrieve_permit", retrieve_permit_status_node)
    workflow.add_node("retrieve_manuals", retrieve_manual_context_node)
    workflow.add_node("assess_risks", call_bedrock_llm_node)
    
    # 3. Establish Edges
    workflow.set_entry_point("retrieve_permit")
    workflow.add_edge("retrieve_permit", "retrieve_manuals")
    workflow.add_edge("retrieve_manuals", "assess_risks")
    workflow.add_edge("assess_risks", END)
    
    # 4. Compile Graph
    return workflow.compile()
