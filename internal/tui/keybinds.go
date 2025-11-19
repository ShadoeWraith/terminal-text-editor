package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func DefaultKeyMap() CustomKeyMap {
	return CustomKeyMap{
		Quit: tea.Key{Type: tea.KeyCtrlC, Runes: []rune("c"), Alt: false},
        Save: tea.Key{Type: tea.KeyCtrlS, Runes: []rune("s"), Alt: false},
	}
}

func HandleKeypress(msg tea.KeyMsg, k CustomKeyMap) Msg {
    switch {
    case msg.String() == k.Quit.String():
        return "quit"
    case msg.String() == k.Save.String():
        return "save"

    default:
        return Msg(msg.String())
    }
}