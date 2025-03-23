package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type HackerDashboard struct {
	status string
}

func NewHackerDashboard() HackerDashboard {
	return HackerDashboard{status: "Hacking in progress..."}
}

func (hd HackerDashboard) Update(msg tea.Msg) (HackerDashboard, tea.Cmd) {
	return HackerDashboard{}, nil
}

func (hd HackerDashboard) View() string {
	return hd.status
}
