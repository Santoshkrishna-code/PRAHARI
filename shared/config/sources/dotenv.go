package sources

import (
	"bufio"
	"os"
	"strings"
)

// LoadDotEnv reads a local configuration file (e.g. .env) and injects parameters into os environment variables.
func LoadDotEnv(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Optional file: skip silently if missing in production
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines or comment lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		// Strip quotes if present
		val = strings.Trim(val, `"'`)

		if key != "" {
			_ = os.Setenv(key, val)
		}
	}

	return scanner.Err()
}
