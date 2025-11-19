package main

import "github.com/charmbracelet/bubbles/textarea"

type errMsg error

// model represents the entire state of our TUI application.
type model struct {
	textarea            textarea.Model
	filename            string
	loadedContentLength int
	err                 error
}