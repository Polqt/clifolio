package ui

import (
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
)

type menuModel struct {
	open	bool
	items	[]state.Screen
	idx		int
}

func MenuModel() tea.Model {
	return &menuModel{
		items: []state.Screen{
			state.ScreenProjects,
			state.ScreenSkills,
			state.ScreenExperience,
			state.ScreenContact,
		},
	}
}

func (m *menuModel) Init() tea.Cmd {
	return nil
}

func (m *menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()
	switch msg := msg.(type) {
	
	case tea.KeyMsg:
		switch msg.String() {
		case km.Toggle:
			m.open = !m.open
		case km.Confirm, "enter":
			if m.open {
				return m, func() tea.Msg {
					return m.items[m.idx]
				}
			}
		case km.Down, "down", "j":
			if m.idx < len(m.items)-1 {
				m.idx++
			} else {
				m.idx = 0
			}
		case km.Up, "up", "k":
			if m.idx > 0 {
				m.idx--
			} else {
				m.idx = len(m.items) - 1
			}
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *menuModel) View() string {
	s := "What would you like to know about me?\n\n"
	for i, it := range m.items {
		mark := " "
		if i == m.idx {
			mark = ">"
		}
		s += mark + " " + it.String() + "\n"
	}

	return s
}
