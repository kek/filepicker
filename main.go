package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type FilePickerModel struct {
	files    []string
	cursor   int
	contents string
}

func NewModel() FilePickerModel {
	return FilePickerModel{files: ListFiles(), cursor: 0}
}

func (m FilePickerModel) Init() tea.Cmd {
	return readContents(m)
}

func (m FilePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			cmd = tea.Quit
		case "j":
			fallthrough
		case "down":
			m.cursor++
		case "k":
			fallthrough
		case "up":
			m.cursor--
		}
	case string:
		m.contents = msg
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor >= len(m.files) {
		m.cursor = len(m.files) - 1
	}
	return m, cmd
}

func (m FilePickerModel) View() string {
	s := ""
	for i, f := range m.files {
		if i == m.cursor {
			s += "> "
		} else {
			s += "  "
		}
		s += f + "\n"
	}
	s += m.contents
	return s
}

func readContents(m FilePickerModel) tea.Cmd {
	return func() tea.Msg {
		return "hello contents"
	}
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
