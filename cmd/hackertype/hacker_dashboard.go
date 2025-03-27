package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HackerDashboard struct {
	hackProgress HackProgress
	hackStatus   HackStatus
}

type progressTick struct{}

func progressCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(time.Time) tea.Msg { return progressTick{} })
}

func NewHackerDashboard() HackerDashboard {
	progress := NewHackProgress(HackSpeed(1))
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
	hs, cmd := hd.hackStatus.Update(msg)
	hd.hackStatus = hs
	cmds = append(cmds, cmd)
	hp, cmd := hd.hackProgress.Update(msg)
	hd.hackProgress = hp
	cmds = append(cmds, cmd)
	return hd, tea.Batch(cmds...)
}

func (hd HackerDashboard) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, hd.hackStatus.View(), hd.hackProgress.View())
}

func (hd HackerDashboard) Init() tea.Cmd {
	return tea.Batch(hd.hackProgress.Init(), hd.hackStatus.Init(), progressCmd())
}
