package ui

import (
	"clifolio/internal/ui/state"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Category struct {
	Name  string
	Items []string
}

type skillsModel struct {
	categories []Category
	catIndex   int
	cursor     int
}

func SkillsModel() *skillsModel {
	return &skillsModel{
		categories: []Category{
			{
				Name:  "Languages",
				Items: []string{"Go", "Python", "JavaScript", "TypeScript", "SQL"},
			},
			{
				Name:  "Frameworks",
				Items: []string{"React", "Vue", "Next", "Svelte", "MUI", "TailwindCSS", "React Native", "Flutter", "tRPC"},
			},
			{
				Name:  "Tools",
				Items: []string{"Docker", "Git", "GitHub"},
			},
			{
				Name:  "Other",
				Items: []string{"REST", "GraphQL", "gRPC", "Microservices", "CI/CD", "TDD", "Agile"},
			},		
		},
		catIndex: 0,
		cursor:   0,
	}
}

func (m *skillsModel) Init() tea.Cmd {
	return nil
}

func (m *skillsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			if m.catIndex > 0 {
				m.catIndex--
				m.cursor = 0
			}
		case "right", "l":
			if m.catIndex < len(m.categories)-1 {
				m.catIndex++
				m.cursor = 0
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.categories[m.catIndex].Items)-1 {
				m.cursor++
			}
		case "b", "esc":
			return m, func() tea.Msg  {
				return state.Menu
			}
		}
	}
	return m, nil
}

func (m *skillsModel) View() string {
	s := "\n\nSkills\n\n"

	for i, c := range m.categories {
		if i == m.catIndex {
			s += fmt.Sprintf("[ %s ] ", c.Name)
		} else {
			s += fmt.Sprintf("  %s   ", c.Name)
		}
	}
	s += "\n\n"

	items := m.categories[m.catIndex].Items
	if len(items) == 0 {
		s += "(no titems in this category)\n\n"
		s += "Press b to go back\n"
		return s
	}

	for i, it := range items {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += fmt.Sprintf(" %s %s\n", cursor, it)
	}

	s += "\nUse ←/→ to switch category, j/k to navigate, b to go back, q to quit.\n"
	return s
}