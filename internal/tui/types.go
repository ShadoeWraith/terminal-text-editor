package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textarea            textarea.Model

	keys 				CustomKeyMap

	filename            string
	loadedContentLength int
	isDirty 			bool

	err                 error
	status 				string
}

type Msg string 

type CustomKeyMap struct {
	Quit tea.Key
    Save tea.Key
}

type SaveCompleteMsg struct {
    FilePath string
    ContentLength int
}

type ErrMsg struct {
	Err error
}

