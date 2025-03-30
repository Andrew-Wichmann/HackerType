package main

import (
	"math/rand/v2"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type idleProgressTick struct{}

func idleProgress() tea.Cmd {
	return tea.Tick(time.Second*1, func(time.Time) tea.Msg { return idleProgressTick{} })
}

type HackSpeed int

type HackProgress struct {
	hackSpeed         HackSpeed
	mainProgress      progress.Model
	keystrokeProgress progress.Model
	mainPercent       float64
	keystrokePercent  float64
}

func NewHackProgress(speed HackSpeed) HackProgress {
	mainProgress := progress.New()
	keystrokeProgress := progress.New()
	return HackProgress{hackSpeed: speed, mainProgress: mainProgress, keystrokeProgress: keystrokeProgress}
}

func (hp HackProgress) Update(msg tea.Msg) (HackProgress, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg.(type) {
	case idleProgressTick:
		hp.keystrokePercent += 0.01
		cmd := hp.keystrokeProgress.SetPercent(hp.keystrokePercent)
		cmds = append(cmds, cmd, idleProgress())
	case progress.FrameMsg:
		mainProgress, cmd := hp.mainProgress.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		hp.mainProgress = mainProgress.(progress.Model)
		keystrokeProgress, cmd := hp.keystrokeProgress.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		hp.keystrokeProgress = keystrokeProgress.(progress.Model)
	case tea.KeyMsg:
		hp.keystrokePercent += 0.001 * float64(hp.hackSpeed)
		cmd := hp.keystrokeProgress.SetPercent(hp.keystrokePercent)
		cmds = append(cmds, cmd)
	}
	if hp.keystrokePercent >= 1 {
		hp.mainPercent += 0.7 * rand.Float64()
		cmd := hp.mainProgress.SetPercent(hp.mainPercent)
		cmds = append(cmds, cmd)
		hp.keystrokePercent = 0
		cmd = hp.keystrokeProgress.SetPercent(hp.keystrokePercent)
		cmds = append(cmds, cmd)
	}
	if hp.mainPercent >= 1 {
		hp.mainPercent = 0
		cmds = append(cmds, finishHack)
	}
	return hp, tea.Batch(cmds...)
}

func (hp HackProgress) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, hp.mainProgress.View(), hp.keystrokeProgress.View())
}

func (hp HackProgress) Init() tea.Cmd {
	return tea.Batch(hp.mainProgress.Init(), hp.keystrokeProgress.Init())
}
