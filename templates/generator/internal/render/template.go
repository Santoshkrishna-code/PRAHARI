package render

import (
	"strings"
)

// Evaluate replaces template double curly brackets placeholders with input variables maps.
func Evaluate(content string, vars map[string]string) string {
	for placeholder, replacement := range vars {
		content = strings.ReplaceAll(content, "{{"+placeholder+"}}", replacement)
	}
	return content
}
