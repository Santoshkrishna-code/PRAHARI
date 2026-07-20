package renderer

import (
	"bytes"
	"text/template"
)

// Compiler compiles dynamic message schemas replacing variables.
type Compiler struct {
}

// NewCompiler constructs a Compiler.
func NewCompiler() *Compiler {
	return &Compiler{}
}

// Compile compiles string values evaluating placeholders keys.
func (c *Compiler) Compile(body string, vars map[string]string) (string, error) {
	tmpl, err := template.New("message").Parse(body)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
