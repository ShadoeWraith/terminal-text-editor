package tui

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
	)
}

func InitialModel(filename string) Model {
	ta := textarea.New()
	ta.Focus()

	ta.SetWidth(80)
	ta.SetHeight(20)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.Color("236")).
		Foreground(lipgloss.Color("252"))
	ta.FocusedStyle.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	ta.FocusedStyle.Text = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	m := Model{
		textarea: ta, 
		keys: 	  DefaultKeyMap(),
		filename: filename,
		isDirty:  false,
		err:      nil,
	}


	content, err := readFileContents(filename)
	if err == nil {
		contentStr := string(content)
		m.textarea.SetValue(contentStr)
		m.loadedContentLength = len(contentStr)

		if len(contentStr) == 0 {
			m.err = fmt.Errorf("file '%s' loaded successfully, but is empty (0 characters)", filepath.Base(filename))
		}
	}

	return m
}
