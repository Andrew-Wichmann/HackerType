package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HackerDashboard struct {
	status  string
	spinner spinner.Model
}

func NewHackerDashboard() HackerDashboard {
	spinner := spinner.New(spinner.WithSpinner(spinner.Globe))
	return HackerDashboard{status: "Hacking in progress...", spinner: spinner}
}

func (hd HackerDashboard) Update(msg tea.Msg) (HackerDashboard, tea.Cmd) {
	spinner, cmd := hd.spinner.Update(msg)
	hd.spinner = spinner
	return hd, cmd
}

func (hd HackerDashboard) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, hd.spinner.View(), hd.status)
}

func (hd HackerDashboard) Init() tea.Cmd {
	return hd.spinner.Tick
}
