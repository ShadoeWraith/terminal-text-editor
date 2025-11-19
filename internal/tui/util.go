package tui

import (
	"fmt"
	"os"
)

func readFileContents(filename string) (string, error) {
	if filename == "" {
		return "", fmt.Errorf("filename cannot be empty")
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("could not read file '%s': %w", filename, err)
	}

	return string(content), nil
}