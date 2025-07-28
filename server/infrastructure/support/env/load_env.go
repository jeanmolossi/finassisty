package env

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func LoadEnv() {
	const envFile = ".env"

	var envpath string

	err := filepath.WalkDir("..", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == envFile {
			envpath = path
			return filepath.SkipAll
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	if envpath == "" {
		panic("no envpath")
	}

	data, err := os.ReadFile(envpath)
	if err != nil {
		panic(err)
	}

	parseEnvContent(data)
}

func newLine(r rune) bool {
	return r == '\n'
}

func parseEnvContent(content []byte) {
	strContent := string(content)
	lines := strings.FieldsFunc(strContent, newLine)

	const kvSize = 2

	for _, line := range lines {
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		parts := strings.SplitN(line, "=", kvSize)

		// has just the key set
		if len(parts) == 1 {
			//nolint:errcheck // just set some environment
			os.Setenv(parts[0], "")
			continue
		}

		// has key and value
		if len(parts) == kvSize {
			//nolint:errcheck // just set some environment
			os.Setenv(parts[0], unquoteString(parts[1]))
		}
	}
}

func unquoteString(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}

	return s
}
