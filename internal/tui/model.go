package tui

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ErrMsg error

type Model struct {
	textarea            textarea.Model
	filename            string
	loadedContentLength int
	err                 error
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func InitialModel(filename string) Model {
	ti := textarea.New()
	ti.Focus()

	ti.SetWidth(80)
	ti.SetHeight(20)

	ti.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.Color("236")).
		Foreground(lipgloss.Color("252"))
	ti.FocusedStyle.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	ti.FocusedStyle.Text = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	m := Model{
		textarea: ti, 
		filename: filename,
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
