package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type SaveCompleteMsg struct {
    FilePath string
    ContentLength int
}

type SaveErrorMsg struct {
    Err error
}

func SaveFileCmd(filename string, content string) tea.Cmd {
    return func() tea.Msg {
        err := os.WriteFile(filename, []byte(content), 0644)
        
        if err != nil {
            return SaveErrorMsg{
                Err: fmt.Errorf("failed to write to file '%s': %w", filename, err),
            }
        }
        
        return SaveCompleteMsg{
            FilePath: filename,
            ContentLength: len(content),
        }
    }
}