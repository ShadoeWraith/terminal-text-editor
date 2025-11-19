package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd
    var cmd tea.Cmd

    switch msg := msg.(type) {
    
    case tea.WindowSizeMsg:
        m.textarea.SetWidth(msg.Width - 4)
        m.textarea.SetHeight(msg.Height - 5)
        
    case tea.KeyMsg:
        action := HandleKeypress(msg) 

        switch action {
        case "quit":
            return m, tea.Quit
            
        default:
            if !m.textarea.Focused() {
                cmd = m.textarea.Focus()
                cmds = append(cmds, cmd)
            }
        }

    case ErrMsg:
        m.err = msg
        return m, nil
    }

    m.textarea, cmd = m.textarea.Update(msg)
    cmds = append(cmds, cmd)
    
    return m, tea.Batch(cmds...)
}