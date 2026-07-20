from graph import create_agent_graph
from schemas import AgentInput, SafetyReport

class SafetyAgent:
    def __init__(self):
        self.graph = create_agent_graph()
        
    def evaluate_safety_entry(
        self, worker_id: str, zone_id: str, role: str, plant: str = "Main Plant"
    ) -> SafetyReport:
        """Executes the state graph safety audit workflow."""
        inputs = {
            "input": AgentInput(
                worker_id=worker_id,
                zone_id=zone_id,
                worker_role=role,
                plant_name=plant
            ),
            "messages": []
        }
        
        # Invoke compiled graph
        result = self.graph.invoke(inputs)
        return result["final_report"]

if __name__ == "__main__":
    agent = SafetyAgent()
    report = agent.evaluate_safety_entry(
        worker_id="usr-99201",
        zone_id="zone-haz-01",
        role="Welder"
    )
    print("Agent Evaluation Complete:")
    print(f"Allowed Entry: {report.allowed_entry}")
    print(f"Mitigations: {report.mitigation_steps}")
