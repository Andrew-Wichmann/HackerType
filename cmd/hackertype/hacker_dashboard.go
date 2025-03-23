package main

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HackerDashboard struct {
	status       string
	worldSpinner spinner.Model
	ellipsis     spinner.Model
	progressbar  progress.Model
	loadPercent  float64
}

type progressTick struct{}

func progressCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(time.Time) tea.Msg { return progressTick{} })
}

func NewHackerDashboard() HackerDashboard {
	worldSpinner := spinner.New(spinner.WithSpinner(spinner.Globe))
	ellipsis := spinner.New(spinner.WithSpinner(spinner.Ellipsis))
	progress := progress.New()
	return HackerDashboard{
		status:       "Hacking in progress",
		worldSpinner: worldSpinner,
		ellipsis:     ellipsis,
		progressbar:  progress,
	}
}

func (hd HackerDashboard) Update(msg tea.Msg) (HackerDashboard, tea.Cmd) {
	log.Printf("Update %#v\n", msg)
	cmds := []tea.Cmd{}
	worldSpinner, cmd := hd.worldSpinner.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	hd.worldSpinner = worldSpinner
	ellipsis, cmd := hd.ellipsis.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	hd.ellipsis = ellipsis
	switch msg.(type) {
	case progressTick:
		hd.loadPercent += 0.01
		if hd.loadPercent >= 1 {
			hd.loadPercent = 0
		}
		cmd := hd.progressbar.SetPercent(hd.loadPercent)
		cmds = append(cmds, cmd, progressCmd())
	case progress.FrameMsg:
		progressbar, cmd := hd.progressbar.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		hd.progressbar = progressbar.(progress.Model)
	case tea.KeyMsg:
		hd.loadPercent += 0.001

	}
	return hd, tea.Batch(cmds...)
}

func (hd HackerDashboard) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Top, hd.worldSpinner.View(), hd.status, hd.ellipsis.View()), hd.progressbar.View())
}

func (hd HackerDashboard) Init() tea.Cmd {
	return tea.Batch(hd.worldSpinner.Tick, hd.ellipsis.Tick, hd.progressbar.Init(), progressCmd())
}
