package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type FilePickerModel struct {
	gizmos   int
	gremlins int
}

func NewModel() FilePickerModel {
	return FilePickerModel{gizmos: 1, gremlins: 0}
}

func (m FilePickerModel) Init() tea.Cmd {
	return nil
}

func (m FilePickerModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m FilePickerModel) View() string {
	return "hello world"
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
