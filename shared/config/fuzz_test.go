package config_test

import (
	"os"
	"testing"

	"prahari/shared/config/sources"
)

// FuzzLoadDotEnv runs native fuzzing testing on the raw dotenv file parser.
func FuzzLoadDotEnv(f *testing.F) {
	// Add seed corpus for structural base checks
	f.Add("DB_PORT=5432\nDB_HOST=localhost\n# comment line\nKEY=\"quoted-val\"")
	f.Add("")
	f.Add("INVALID_LINE_NO_EQUALS")
	f.Add("=EMPTY_KEY")

	f.Fuzz(func(t *testing.T, data string) {
		tmpFile, err := os.CreateTemp("", "fuzz-dotenv-*")
		if err != nil {
			return
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		if _, err := tmpFile.WriteString(data); err != nil {
			return
		}

		// Execute parser validation
		_ = sources.LoadDotEnv(tmpFile.Name())
	})
}
