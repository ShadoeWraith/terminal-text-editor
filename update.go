package main

import tea "github.com/charmbracelet/bubbletea"

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
