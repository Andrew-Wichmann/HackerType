package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("bubbletea.log", "program")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	program := tea.NewProgram(NewModel(), tea.WithAltScreen())
	_, err = program.Run()
	if err != nil {
		panic(err)
	}
}
