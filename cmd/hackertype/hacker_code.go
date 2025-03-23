package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type HackerTextArea struct {
	ta          textarea.Model
	hacker_code string
	current_pos int
}

func NewHackerTextArea() HackerTextArea {
	m := HackerTextArea{}
	files, err := os.ReadDir("hacker_codes/")
	if err != nil {
		panic(err)
	}
	file := files[rand.Intn(len(files))]
	hacker_code, err := os.ReadFile(fmt.Sprintf("hacker_codes/%s", file.Name()))
	if err != nil {
		panic(err)
	}
	m.hacker_code = string(hacker_code)
	m.ta = textarea.New()
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	m.ta.SetWidth(width / 2)
	m.ta.SetHeight(height)
	m.ta.MaxHeight = 0
	m.ta.MaxWidth = 0
	m.ta.CharLimit = 0
	return m
}

func (m HackerTextArea) View() string {
	return m.ta.View()
}

func (m HackerTextArea) Update(msg tea.Msg) (HackerTextArea, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		m.current_pos += rand.Intn(5)
		m.ta.SetValue(m.hacker_code[:m.current_pos])
		return m, nil
	}
	return m, nil
}
