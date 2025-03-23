package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HackerDashboard struct {
	status       string
	worldSpinner spinner.Model
	ellipsis     spinner.Model
}

func NewHackerDashboard() HackerDashboard {
	worldSpinner := spinner.New(spinner.WithSpinner(spinner.Globe))
	ellipsis := spinner.New(spinner.WithSpinner(spinner.Ellipsis))
	return HackerDashboard{status: "Hacking in progress", worldSpinner: worldSpinner, ellipsis: ellipsis}
}

func (hd HackerDashboard) Update(msg tea.Msg) (HackerDashboard, tea.Cmd) {
	worldSpinner, cmd1 := hd.worldSpinner.Update(msg)
	hd.worldSpinner = worldSpinner
	ellipsis, cmd2 := hd.ellipsis.Update(msg)
	hd.ellipsis = ellipsis
	return hd, tea.Batch(cmd1, cmd2)
}

func (hd HackerDashboard) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, hd.worldSpinner.View(), hd.status, hd.ellipsis.View())
}

func (hd HackerDashboard) Init() tea.Cmd {
	return tea.Batch(hd.worldSpinner.Tick, hd.ellipsis.Tick)
}
