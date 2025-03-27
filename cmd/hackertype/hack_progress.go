package main

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type HackSpeed int

type HackProgress struct {
	hackSpeed   HackSpeed
	progress    progress.Model
	loadPercent float64
}

func NewHackProgress(speed HackSpeed) HackProgress {
	progress := progress.New()
	return HackProgress{hackSpeed: speed, progress: progress}
}

func (hp HackProgress) Update(msg tea.Msg) (HackProgress, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg.(type) {
	case progressTick:
		hp.loadPercent += 0.01
		if hp.loadPercent >= 1 {
			hp.loadPercent = 0
			cmds = append(cmds, finishHack)
		}
		cmd := hp.progress.SetPercent(hp.loadPercent)
		cmds = append(cmds, cmd, progressCmd())
	case progress.FrameMsg:
		progressbar, cmd := hp.progress.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		hp.progress = progressbar.(progress.Model)
	case tea.KeyMsg:
		hp.loadPercent += 0.004
	}
	return hp, tea.Batch(cmds...)
}

func (hp HackProgress) View() string {
	return hp.progress.View()
}

func (hp HackProgress) Init() tea.Cmd {
	return hp.progress.Init()
}
