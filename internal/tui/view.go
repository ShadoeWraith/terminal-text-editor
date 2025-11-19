package tui

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)




func (m Model) View() string {
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
