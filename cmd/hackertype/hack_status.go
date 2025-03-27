package main

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const HACK_IN_PROGRESS = "Hack in progress"
const HACK_COMPLETE = "HACK COMPLETE!!"

type HackStatus struct {
	status       string
	worldSpinner spinner.Model
	ellipsis     spinner.Model
}

type clearHackComplete struct{}

func resetStatus(delay time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(delay)
		return clearHackComplete{}
	}
}

func NewHackStatus() HackStatus {
	worldSpinner := spinner.New(spinner.WithSpinner(spinner.Globe))
	ellipsis := spinner.New(spinner.WithSpinner(spinner.Ellipsis))
	return HackStatus{
		status:       HACK_IN_PROGRESS,
		worldSpinner: worldSpinner,
		ellipsis:     ellipsis,
	}
}

func (hs HackStatus) Update(msg tea.Msg) (HackStatus, tea.Cmd) {
	cmds := []tea.Cmd{}
	worldSpinner, cmd := hs.worldSpinner.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	hs.worldSpinner = worldSpinner
	ellipsis, cmd := hs.ellipsis.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	hs.ellipsis = ellipsis
	switch msg.(type) {
	case HackFinished:
		hs.status = HACK_COMPLETE
		cmds = append(cmds, resetStatus(time.Duration(2*time.Second)))
	case clearHackComplete:
		hs.status = HACK_IN_PROGRESS
	}
	return hs, tea.Batch(cmds...)
}

func (hs HackStatus) View() string {
	if hs.status == HACK_COMPLETE {
		return HACK_COMPLETE
	} else {
		return lipgloss.JoinHorizontal(lipgloss.Top, hs.worldSpinner.View(), hs.status, hs.ellipsis.View())
	}
}

func (hs HackStatus) Init() tea.Cmd {
	return tea.Batch(hs.worldSpinner.Tick, hs.ellipsis.Tick)
}
