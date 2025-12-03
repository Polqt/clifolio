package ui

import (
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Experience struct {
	Position    string
	Company     string
	Date        string
	Description string
}
type experienceModel struct {
	cursor      int
	experiences []Experience
}

func ExperienceModel() *experienceModel {
	return &experienceModel{
		experiences: []Experience{
			{
				Position:    "Part Time Mobile Developer",
				Company:     "K92 Paints",
				Date:        "December 2024 - April 2025",
				Description: "Developed and maintained a mobile application using Flutter, enhancing user engagement and streamlining paint selection processes for customers.",
			},
		},
		cursor: 0,
	}
}

func (m *experienceModel) Init() tea.Cmd {
	return nil
}

func (m *experienceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		case km.Down, "down", "j":
			if m.cursor < len(m.experiences)-1 {
				m.cursor++
			}
		case km.Up, "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case km.Back, "esc":
			return m, func() tea.Msg { return state.ScreenMenu }
		}
	}
	return m, nil
}

func (m *experienceModel) View() string {
	s := "\n\nExperience\n\n"
	for i, exp := range m.experiences {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s at %s (%s)\n%s\n\n", cursor, exp.Position, exp.Company, exp.Date, exp.Description)
	}
	s += "\nUse j/k to navigate, q to quit.\n"
	return s
}
