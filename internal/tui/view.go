package tui

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	indicator := ""

	if m.isDirty {
		indicator = "*"
	}

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

	headerText := "Go Editor - [untitled.txt]"
	if m.filename != "" {
		headerText = fmt.Sprintf("Go Edit - [%s%s]", filepath.Base(m.filename), indicator)	
	}

	header := headerStyle.Render(headerText)
	status := ""
	if m.err != nil {
		status = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(m.err.Error())
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		m.textarea.View(),
		status,
		footerStyle.Render("Keys: Ctrl+C to quit"),
	)
}
