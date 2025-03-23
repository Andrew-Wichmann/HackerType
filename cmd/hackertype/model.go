package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{}

func (Model) View() string {
	return "Hello world"
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (Model) Init() tea.Cmd {
	return nil
}
