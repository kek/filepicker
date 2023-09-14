package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
			if m.cursor < len(m.files)-1 {
				m.cursor++
				cmd = readContents(m)
			}
		case "k":
			fallthrough
		case "up":
			if m.cursor > 0 {
				m.cursor--
				cmd = readContents(m)
			}
		}
	case content:
		m.contents = string(msg)
	}
	return m, cmd
}

func (m FilePickerModel) View() string {
	sidebar := ""
	for i, f := range m.files {
		if i == m.cursor {
			sidebar += "> "
		} else {
			sidebar += "  "
		}
		sidebar += f + "\n"
	}

	style := lipgloss.NewStyle().Height(20).BorderStyle(lipgloss.NormalBorder())

	sidebar = style.Render(sidebar)

	crop := lipgloss.NewStyle().MaxWidth(40).MaxHeight(20)
	contents := crop.Render(m.contents)
	contents = style.Render(contents)

	view := lipgloss.JoinHorizontal(0, sidebar, contents)
	return view
}

type content string

func readContents(m FilePickerModel) tea.Cmd {
	return func() tea.Msg {
		result, err := os.ReadFile(m.files[m.cursor])
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(1)
		}
		return content(result)
	}
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
