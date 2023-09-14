package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	gizmos   int
	gremlins int
}

func initialModel() model {
	return model{gizmos: 1, gremlins: 0}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return "hello world"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
