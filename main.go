package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// --- Model and Error Types ---

type errMsg error

// model represents the entire state of our TUI application.
type model struct {
	textarea            textarea.Model
	filename            string
	loadedContentLength int
	err                 error
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

// --- Lifecycle Methods (Model, Init, Update, View) ---

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// Handle window resizing messages
	case tea.WindowSizeMsg:
		// Calculate available space for the text area:
		// We subtract 4 lines for the header (2 lines) and the footer/status (2 lines)
		// and a few characters (4) for padding/borders on the sides.
		m.textarea.SetWidth(msg.Width - 4)
		m.textarea.SetHeight(msg.Height - 5)
		
	case tea.KeyMsg:
		switch msg.String() { // Use msg.String() for complex shortcuts like Ctrl+A
		case "esc":
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case "ctrl+c":
			return m, tea.Quit

		case " ":
			m.textarea.SetCursor(0)
			
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// Define Styles
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Padding(0, 1).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("63")).
		Width(m.textarea.Width() + 2).
		Align(lipgloss.Center)

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Padding(0, 1).
		MarginTop(1)

	// Build the UI parts
	headerText := "GopherEdit - Minimal TUI Editor"
	if m.filename != "" {
		// If a file is loaded, show the base filename in the header
		headerText = fmt.Sprintf("GopherEdit - %s", filepath.Base(m.filename))

		if m.loadedContentLength > 0 {
			headerText = fmt.Sprintf("%s (%d chars)", headerText, m.loadedContentLength)
		}
	}

	header := headerStyle.Render(headerText)
	status := ""
	if m.err != nil {
		status = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(m.err.Error()) // Red error message
	}

	// Normal view: Header, Textarea, Footer
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		m.textarea.View(), // Renders the actual text editor
		status,
		// UPDATED FOOTER MESSAGE
		footerStyle.Render("Keys: ESC to blur, any key to focus, Ctrl+C to quit, Ctrl+A to select all"),
	)
}

// --- Main Function ---

func main() {
	var filename string

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	// Now calls the standalone function `initialModel`
	p := tea.NewProgram(initialModel(filename), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}