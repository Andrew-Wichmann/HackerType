package main

import (
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	hacker_code string
	current_pos int
}

func NewModel() Model {
	m := Model{}
	hacker_code, err := os.ReadFile("hacker_codes/main.c")
	if err != nil {
		panic(err)
	}
	m.hacker_code = string(hacker_code)
	return m
}

func (m Model) View() string {
	return m.hacker_code[:m.current_pos]
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
		m.current_pos += rand.Intn(5)
	}
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}
