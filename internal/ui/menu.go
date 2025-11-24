package ui

import (
	"clifolio/internal/ui/state"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type menuModel struct {
	choices 	[]string
	cursor 		int
	selected 	map[int]struct{}
}

func MenuModel() menuModel {
	return menuModel{
		choices: []string{"Projects", "Skills", "Experience", "Contact"},
		selected: make(map[int]struct{}),
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	case tea.KeyMsg:
		
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		
		case "enter", " ":
			switch m.cursor {
			case 0:
				return m, func() tea.Msg { return state.Projects }
			case 1:
				return m, func() tea.Msg { return state.Skills }
			case 2:
				return m, func() tea.Msg { return state.Experience }
			case 3:
				return m, func() tea.Msg { return state.Contact }
			}
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	s := "\n\n What would you like to know about me?\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += cursor + " [" + checked + "] " + choice + "\n"
	}

	s += "\nPress q to quit.\n"
	return s
}

func RunMenu() {
	p := tea.NewProgram(MenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}