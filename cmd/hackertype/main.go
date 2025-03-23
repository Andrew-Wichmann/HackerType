package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(Model{}, tea.WithAltScreen())
	_, err := program.Run()
	if err != nil {
		panic(err)
	}
}
