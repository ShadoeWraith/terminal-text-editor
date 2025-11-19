package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func readFileContents(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("could not read file '%s': %w", filename, err)
	}

	return string(content), nil
}

func SaveFileCmd(filename string, content string) tea.Cmd {
    return func() tea.Msg {
        err := os.WriteFile(filename, []byte(content), 0644)
        
        if err != nil {
            return ErrMsg{
                Err: fmt.Errorf("failed to write to file '%s': %w", filename, err),
            }
        }
        
        return SaveCompleteMsg{
            FilePath: filename,
            ContentLength: len(content),
        }
    }
}