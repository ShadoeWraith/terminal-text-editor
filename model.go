package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// model represents the entire state of our TUI application.

// Init implements tea.Model.
func (m model) Init() tea.Cmd {
	return textarea.Blink
}

// --- Initialization Logic ---

// initialModel is the function that initializes the model, following the tea.Model interface.
func initialModel(filename string) model {
	// 1. Initialize the textarea component
	ti := textarea.New()
	ti.Placeholder = "Once upon a time..."
	ti.Focus()

	// Set initial dimensions
	ti.SetWidth(80)
	ti.SetHeight(20)

	// Apply custom styling (matching your previous styles)
	ti.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.Color("236")).
		Foreground(lipgloss.Color("252"))
	ti.FocusedStyle.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	ti.FocusedStyle.Text = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))

	m := model{
		textarea: ti, // Attach the newly created textarea
		filename: filename,
		err:      nil,
	}

	// 2. File Reading Logic
	if filename != "" {
		content, err := os.ReadFile(filename)
		if err == nil {
			// File read successfully, load content into the editor
			contentStr := string(content)
			m.textarea.SetValue(contentStr) // Set the value on the attached textarea
			m.loadedContentLength = len(contentStr)

			// NEW: Set the cursor to the beginning of the document (index 0)
			// Check for empty file
			if len(contentStr) == 0 {
				m.err = fmt.Errorf("file '%s' loaded successfully, but is empty (0 characters)", filepath.Base(filename))
			}
		} else if os.IsNotExist(err) {
			m.err = fmt.Errorf("file not found: %s. Starting with empty buffer", filepath.Base(filename))
		} else {
			m.err = fmt.Errorf("error reading file %s: %v", filepath.Base(filename), err)
		}
	}

	return m
}
