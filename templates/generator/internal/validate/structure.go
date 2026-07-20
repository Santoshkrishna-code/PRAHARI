package validate

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CheckCleanArchitecture Rules checks that Go files in domain/ and application/ do not import adapters.
func CheckCleanArchitectureRules(rootPath string) []error {
	var errorsList []error

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			rulesErrors := auditImports(path)
			errorsList = append(errorsList, rulesErrors...)
		}
		return nil
	})

	if err != nil {
		errorsList = append(errorsList, fmt.Errorf("failed to traverse directory trees: %w", err))
	}

	return errorsList
}

func auditImports(path string) []error {
	var fileErrors []error

	file, err := os.Open(path)
	if err != nil {
		return []error{fmt.Errorf("failed to open file %s: %w", path, err)}
	}
	defer file.Close()

	isDomain := strings.Contains(path, "/domain/")
	isApp := strings.Contains(path, "/application/")

	scanner := bufio.NewScanner(file)
	inImports := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "import (") {
			inImports = true
			continue
		}
		if inImports && line == ")" {
			inImports = false
			continue
		}

		// Single line imports
		if strings.HasPrefix(line, "import ") {
			targetImport := strings.Trim(strings.TrimPrefix(line, "import "), `"`)
			if err := checkImportRules(path, targetImport, isDomain, isApp); err != nil {
				fileErrors = append(fileErrors, err)
			}
		}

		// Multiline imports
		if inImports && line != "" {
			targetImport := strings.Trim(line, `"`)
			if err := checkImportRules(path, targetImport, isDomain, isApp); err != nil {
				fileErrors = append(fileErrors, err)
			}
		}
	}

	return fileErrors
}

func checkImportRules(filePath, importedPackage string, isDomain, isApp bool) error {
	// Rule 1: domain layers cannot import infrastructure or interfaces
	if isDomain {
		if strings.Contains(importedPackage, "/infrastructure") || strings.Contains(importedPackage, "/interfaces") {
			return fmt.Errorf("architecture violation in %s: domain model cannot import adapter '%s'", filePath, importedPackage)
		}
	}

	// Rule 2: application service layers cannot import interface presentation handlers
	if isApp {
		if strings.Contains(importedPackage, "/interfaces") {
			return fmt.Errorf("architecture violation in %s: application services cannot import presentation controllers '%s'", filePath, importedPackage)
		}
	}

	return nil
}
