import re
from typing import Any
from app.core.exceptions import ValidationException
from app.core.model_registry import MODEL_REGISTRY
from app.core.prompt_renderer import prompt_renderer

# Prompt Injection Patterns (Common jailbreaks and bypass triggers)
PROMPT_INJECTION_PATTERNS = [
    r"ignore\s+(?:previous|all)\s+instructions",
    r"system\s+override",
    r"bypass\s+restrictions",
    r"you\s+are\s+now\s+in\s+developer\s+mode",
    r"ignore\s+system\s+directives",
    r"jailbreak",
    r"dan\s+mode",
    r"do\s+anything\s+now",
]

# Common API Key / Credential Regex
CREDENTIAL_PATTERns = [
    r"sk-[a-zA-Z0-9]{32,}",  # OpenAI Key format
    r"AIzaSy[a-zA-Z0-9_-]{33}",  # Google API key format
    r"secret[_-]?key",
    r"password\s*=\s*['\"][a-zA-Z0-9!@#$%^&*()_+]{6,}['\"]",
]


class PromptValidator:
    """Security guard and format validator analyzing prompt templates and variables."""

    def validate_templates(
        self,
        system_template: str,
        user_template: str,
        variables: dict[str, Any],
    ) -> tuple[bool, list[str], list[str], bool, bool]:
        """Runs compile, security, and injection checks on template layouts and inputs."""
        errors = []
        warnings = []
        injection_detected = False
        unsafe_detected = False

        # 1. Compilation Validation
        try:
            system_vars = prompt_renderer.get_template_variables(system_template)
            user_vars = prompt_renderer.get_template_variables(user_template)
            required_vars = system_vars.union(user_vars)
        except Exception as e:
            errors.append(f"Jinja template parsing failed: {e}")
            return False, errors, warnings, False, False

        # 2. Check for missing variables
        for var in required_vars:
            if var not in variables:
                warnings.append(f"Template variable '{{{{ {var} }}}}' has no runtime value supplied.")

        # 3. Check for unused variables
        for var in variables:
            if var not in required_vars:
                warnings.append(f"Supplied runtime variable '{var}' is not utilized by the templates.")

        # 4. Prompt Injection Security Check (on variables payload and user template)
        content_to_scan = [user_template] + [str(v) for v in variables.values()]
        for content in content_to_scan:
            for pattern in PROMPT_INJECTION_PATTERNS:
                if re.search(pattern, content, re.IGNORECASE):
                    injection_detected = True
                    errors.append(f"Security Warning: Potential prompt injection detected matching pattern: {pattern}")

        # 5. Secret and Credential Leaks Scan
        for content in content_to_scan:
            for pattern in CREDENTIAL_PATTERns:
                if re.search(pattern, content, re.IGNORECASE):
                    unsafe_detected = True
                    errors.append(f"Security Alert: Hardcoded API Key or credentials detected matching: {pattern}")

        is_valid = len(errors) == 0
        return is_valid, errors, warnings, injection_detected, unsafe_detected

    def estimate_cost(self, system_text: str, user_text: str, model_name: str) -> tuple[int, float]:
        """Calculates token counts and costs based on model registry coefficients."""
        model = model_name or "gpt-4o"
        metadata = MODEL_REGISTRY.get(model)
        
        total_chars = len(system_text) + len(user_text)
        # Approximate 4 characters per token
        estimated_tokens = max(1, total_chars // 4)

        cost = 0.0
        if metadata:
            cost = (estimated_tokens / 1_000_000) * metadata.input_cost_per_million

        return estimated_tokens, round(cost, 6)


prompt_validator = PromptValidator()
