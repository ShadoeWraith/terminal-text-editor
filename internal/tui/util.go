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

func SaveFileContents(filename string, content string) error {
    if filename == "" {
        return fmt.Errorf("cannot save: filename is empty")
    }
 
    err := os.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        return fmt.Errorf("failed to write to file '%s': %w", filename, err)
    }
    
    return nil
}