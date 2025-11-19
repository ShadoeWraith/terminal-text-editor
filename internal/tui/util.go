package tui

import (
	"fmt"
	"os"
)

func readFileContents(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("could not read file '%s': %w", filename, err)
	}

	return string(content), nil
}