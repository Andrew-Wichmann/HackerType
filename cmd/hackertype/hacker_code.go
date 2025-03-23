package main

import (
	"fmt"
	"log"
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
	height      int // should maybe be fed from WindowRezie
}

func (hta *HackerTextArea) ResetHackerCode() {
	files, err := os.ReadDir("hacker_codes/")
	if err != nil {
		panic(err)
	}
	file := files[rand.Intn(len(files))]
	hacker_code, err := os.ReadFile(fmt.Sprintf("hacker_codes/%s", file.Name()))
	if err != nil {
		panic(err)
	}
	hta.hacker_code = string(hacker_code)
	hta.current_pos = 0
}

func NewHackerTextArea() HackerTextArea {
	m := HackerTextArea{}
	m.ResetHackerCode()
	m.ta = textarea.New()
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	m.height = height
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
		m.current_pos += rand.Intn(10)
		log.Println("%d", m.ta.Line())
		if m.ta.Line() >= m.height {
			m.ta.Reset()
			m.hacker_code = m.hacker_code[m.current_pos:]
			m.current_pos = 0
		} else {
			if m.current_pos > len(m.hacker_code) {
				m.ResetHackerCode()
			} else {
				m.ta.SetValue(m.hacker_code[:m.current_pos])
			}
		}
		return m, nil
	}
	return m, nil
}
