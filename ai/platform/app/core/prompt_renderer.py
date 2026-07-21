from typing import Any
from jinja2 import Environment, meta, Template, BaseLoader
from app.core.exceptions import ValidationException


class PromptRenderer:
    """Prompt template compiler using Jinja2 for layouts compiling and variables resolving."""

    def __init__(self) -> None:
        # Strict undefined variable raising configuration
        self._env = Environment(loader=BaseLoader())

    def get_template_variables(self, template_str: str) -> set[str]:
        """Parses the template string and extracts all undeclared variables."""
        try:
            parsed_content = self._env.parse(template_str)
            return meta.find_undeclared_variables(parsed_content)
        except Exception as e:
            raise ValidationException(f"Failed to parse template variables: {e}") from e

    def render(self, template_str: str, variables: dict[str, Any]) -> str:
        """Renders the template with the provided variables dictionary."""
        try:
            template = self._env.from_string(template_str)
            return template.render(**variables)
        except Exception as e:
            raise ValidationException(f"Template rendering compilation failure: {e}") from e

    def resolve_variables(
        self,
        system_template: str,
        user_template: str,
        input_variables: dict[str, Any],
    ) -> tuple[str, str, list[str], list[str]]:
        """Compiles system and user prompts, tracking used and missing variables."""
        # 1. Discover all variables required by the templates
        system_vars = self.get_template_variables(system_template)
        user_vars = self.get_template_variables(user_template)
        required_vars = system_vars.union(user_vars)

        # 2. Identify missing variables (where no default or value is supplied)
        missing_vars = [var for var in required_vars if var not in input_variables]
        resolved_vars = [var for var in required_vars if var in input_variables]

        # 3. Supply empty defaults for missing variables to prevent rendering crashes
        full_variables = input_variables.copy()
        for var in missing_vars:
            full_variables[var] = ""

        # 4. Render system and user templates
        rendered_system = self.render(system_template, full_variables)
        rendered_user = self.render(user_template, full_variables)

        return rendered_system, rendered_user, resolved_vars, missing_vars


prompt_renderer = PromptRenderer()
