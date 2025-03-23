package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	hackerTextArea  HackerTextArea
	hackerDashboard HackerDashboard
}

func NewModel() Model {
	hackerTextArea := NewHackerTextArea()
	hackerDashboard := NewHackerDashboard()
	m := Model{hackerTextArea: hackerTextArea, hackerDashboard: hackerDashboard}
	return m
}

func (m Model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, m.hackerTextArea.View(), m.hackerDashboard.View())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	hackerTextArea, cmd := m.hackerTextArea.Update(msg)
	m.hackerTextArea = hackerTextArea
	return m, cmd
}

func (m Model) Init() tea.Cmd {
	return nil
}
