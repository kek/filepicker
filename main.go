package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type FilePickerModel struct {
	files    []string
	cursor   int
	gizmos   int
	gremlins int
}

func NewModel() FilePickerModel {
	return FilePickerModel{gizmos: 1, gremlins: 0, files: ListFiles(), cursor: 0}
}

func (m FilePickerModel) Init() tea.Cmd {
	return nil
}

func (m FilePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m FilePickerModel) View() string {
	s := ""
	for _, f := range m.files {
		s += f + "\n"
	}
	return s
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
