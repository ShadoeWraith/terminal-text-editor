package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Msg string 

type CustomKeyMap struct {
	Quit tea.Key
}

func DefaultKeyMap() CustomKeyMap {
	return CustomKeyMap{
		Quit: tea.Key{Type: tea.KeyCtrlC, Runes: []rune("c"), Alt: false},
	}
}

func HandleKeypress(msg tea.KeyMsg) Msg {
    k := DefaultKeyMap() 

    switch {
    case msg.String() == k.Quit.String():
        return "quit" 

    default:
        return Msg(msg.String())
    }
}