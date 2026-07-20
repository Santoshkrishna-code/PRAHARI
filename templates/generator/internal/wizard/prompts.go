package wizard

import (
	"fmt"
	"os"
	"path/filepath"
)

// Answers holds CLI wizard inputs.
type Answers struct {
	ServiceName string
	ModuleName  string
	Port        string
	Database    string
}

// DefaultAnswers returns standard bootstrap variables.
func DefaultAnswers(name string) Answers {
	return Answers{
		ServiceName: name + "-service",
		ModuleName:  "prahari/services/" + name,
		Port:        "8080",
		Database:    "postgres",
	}
}

// PromptInteractive triggers standard prompts, fallback to default configurations.
func PromptInteractive() error {
	fmt.Println("=== PRAHARI Interactive Service Scaffold Wizard ===")
	fmt.Println("Using default parameters to scaffold service 'incident'...")
	return Scaffold("service", "incident", DefaultAnswers("incident"))
}

// Scaffold creates the folder directory tree and writes files evaluated with variables.
func Scaffold(subType, name string, answers Answers) error {
	targetDir := filepath.Join("services", name)
	fmt.Printf("[WIZARD] Scaffolding workspace paths under: %s/\n", targetDir)

	err := os.MkdirAll(filepath.Join(targetDir, "cmd/server"), 0755)
	if err != nil {
		return fmt.Errorf("failed to create cmd path: %w", err)
	}

	err = os.MkdirAll(filepath.Join(targetDir, "internal/bootstrap"), 0755)
	if err != nil {
		return fmt.Errorf("failed to create bootstrap path: %w", err)
	}

	// Write mock main entry point to verify scaffold compilation
	mainCode := fmt.Sprintf(`package main

import "fmt"

func main() {
	fmt.Println("Service %s successfully bootstrapped on port %s!")
}
`, answers.ServiceName, answers.Port)

	err = os.WriteFile(filepath.Join(targetDir, "cmd/server/main.go"), []byte(mainCode), 0644)
	if err != nil {
		return fmt.Errorf("failed to write main server entrypoint: %w", err)
	}

	fmt.Printf("[WIZARD] %s scaffold completed successfully!\n", subType)
	return nil
}
