# PRAHARI Python Coding Standards
## AI & Computer Vision Programming Guidelines

This document outlines the coding standards, style rules, and structural guidelines for Python microservices, AI Agents, and Computer Vision workers in the PRAHARI platform.

---

## 1. Style & Code Formatting

1. **PEP 8 Compliance**: Code must strictly comply with PEP 8.
2. **Auto-Formatting**:
   - Enforce **`black`** for code formatting (88 character line limit).
   - Enforce **`isort`** for sorting imports.
3. **Naming Conventions**:
   - Variables, functions, and modules: `snake_case`.
   - Classes: `PascalCase`.
   - Constants: `UPPER_SNAKE_CASE`.

---

## 2. Type Hinting & Annotations

To ensure reliability in complex AI/CV execution lines, **type hints are mandatory** for all function signatures.
- Enforce **`mypy`** validation on pull requests.

```python
from typing import Dict, Any

def execute_risk_assessment(
    worker_id: str, 
    vitals: Dict[str, float]
) -> bool:
    # Business logic here
    return True
```

---

## 3. Dependency & Workspace Governance

- Use **Poetry** (`pyproject.toml`) for managing dependencies. Lock files (`poetry.lock`) must be committed.
- Use absolute imports (e.g., `from ai.orchestrator.config import settings`), not relative ones.
- Avoid hacking system path parameters (`sys.path.append`).

---

## 4. FastAPI & Pydantic Conventions

- Always utilize Pydantic v2 schemas for request and response body parsing.
- Always use dependency injection (`Depends`) for external services like OpenSearch and Bedrock clients.
- Handle exceptions using custom FastAPI Exception Handlers returning JSON errors matching our REST standard code structure.

---

## 5. AI Agent (LangGraph / LangChain) Guidelines

1. **Prompts Organization**: Prompts must NOT be inline strings inside graph nodes. They must reside in `prompts.py` or separate templates.
2. **State Management**: LangGraph states must inherit from TypedDict and define clear typing with comments.
3. **Model Integration**: Always access models through the configured AWS Bedrock LLM wrapper in the AI module.
4. **Guardrails**: All agent workflows must execute input/output verification checks (guardrails) before/after interacting with the model.
