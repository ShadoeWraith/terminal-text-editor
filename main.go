package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

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