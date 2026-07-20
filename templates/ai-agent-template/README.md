# AI Agent Template

This template establishes the architectural design pattern and folder layout for all LangGraph-based AI agents in the **PRAHARI** platform.

## File Breakdown

* **`agent.py`**: Invocation coordinator representing the primary API.
* **`graph.py`**: Compiles nodes and defines paths/edges.
* **`nodes.py`**: Concrete functions representing graph transitions.
* **`tools.py`**: Integrations (e.g. searching manual indexes or checking permit stores).
* **`prompts.py`**: Isolated system instructions.
* **`schemas.py`**: Pydantic input/output schemas and TypedDict state tracking.
* **`memory.py`**: Short/Long term checkpoint state savers.
* **`config.py`**: Local/Global Pydantic configs.

## How to Initialize a New AI Agent

1. Copy the contents of this folder into `ai/agents/<new-agent-name>`.
2. Update the system prompt inside `prompts.py` to match the target context (e.g. `permit-agent`, `compliance-agent`).
3. Add specialized Pydantic schema validation structures in `schemas.py`.
4. Modify the execution flow and inject specific tools in `graph.py` and `tools.py`.
5. Execute the local testing script:
   ```bash
   python agent.py
   ```
