package main

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HackerDashboard struct {
	hackProgress progress.Model
	hackStatus   HackStatus
	loadPercent  float64
}

type progressTick struct{}

func progressCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(time.Time) tea.Msg { return progressTick{} })
}

func NewHackerDashboard() HackerDashboard {
	progress := progress.New()
	hackStatus := NewHackStatus()
	return HackerDashboard{
		hackStatus:   hackStatus,
		hackProgress: progress,
	}
}

type HackFinished struct{}

func finishHack() tea.Msg {
	return HackFinished{}
}

func (hd HackerDashboard) Update(msg tea.Msg) (HackerDashboard, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg.(type) {
	case progressTick:
		hd.loadPercent += 0.01
		if hd.loadPercent >= 1 {
			hd.loadPercent = 0
			cmds = append(cmds, finishHack)
		}
		cmd := hd.hackProgress.SetPercent(hd.loadPercent)
		cmds = append(cmds, cmd, progressCmd())
	case progress.FrameMsg:
		progressbar, cmd := hd.hackProgress.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		hd.hackProgress = progressbar.(progress.Model)
	case tea.KeyMsg:
		hd.loadPercent += 0.004
	}
	hs, cmd := hd.hackStatus.Update(msg)
	hd.hackStatus = hs
	cmds = append(cmds, cmd)
	return hd, tea.Batch(cmds...)
}

func (hd HackerDashboard) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, hd.hackStatus.View(), hd.hackProgress.View())
}

func (hd HackerDashboard) Init() tea.Cmd {
	return tea.Batch(hd.hackProgress.Init(), hd.hackStatus.Init(), progressCmd())
}
