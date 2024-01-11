package main

// see https://github.com/charmbracelet/bubbletea/blob/master/examples/result/main.go

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"Hey", "Hello"}

type model struct {
	cursor int
	choice string
}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "enter":
			m.choice = choices[m.cursor]
			return m, tea.Quit
		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}
	return m, nil
}
func (m model) View() string {
	// builder
	s := strings.Builder{}
	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("[•] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(choices[i] + "\n")
	}
	return s.String()
}

func main() {
	p := tea.NewProgram(model{})
	m, err := p.Run()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if m, ok := m.(model); ok && m.choice != "" {
		fmt.Printf("You chose %s!\n", m.choice)
	}
}