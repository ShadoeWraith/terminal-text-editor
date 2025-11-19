package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd
    var cmd tea.Cmd

    oldValue := m.textarea.Value()

    if keyMsg, ok := msg.(tea.KeyMsg); ok {
        action := HandleKeypress(keyMsg, m.keys) 

        switch action {
        case "quit":
            return m, tea.Quit 

        case "save":
            contentToSave := m.textarea.Value()
            return m, SaveFileCmd(m.filename, contentToSave)
            
        default:
            if !m.textarea.Focused() {
                cmd = m.textarea.Focus()
                cmds = append(cmds, cmd)
            }
        }
    }

    m.textarea, cmd = m.textarea.Update(msg)
    cmds = append(cmds, cmd)
    
    if m.textarea.Value() != oldValue {
        m.isDirty = true
    }
    
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.textarea.SetWidth(msg.Width - 4)
        m.textarea.SetHeight(msg.Height - 5)
        
    case SaveCompleteMsg:
        m.isDirty = false 
        m.status = fmt.Sprintf("Saved successfully to %s! (%d chars)", msg.FilePath, msg.ContentLength)
        return m, tea.Batch(cmds...)
    
    case SaveErrorMsg:
        m.err = msg.Err 
        m.status = "SAVE FAILED"
        return m, tea.Batch(cmds...)
    
    case ErrMsg:
        m.err = msg
        return m, tea.Batch(cmds...)
    }
    
    return m, tea.Batch(cmds...)
}