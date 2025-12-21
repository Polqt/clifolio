package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	km := components.DefaultKeymap()

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
		case km.Up, "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case km.Down, "down":
			if m.cursor < len(m.categories[m.catIndex].Items)-1 {
				m.cursor++
			}
		case km.Back, "esc":
			return m, func() tea.Msg { return state.ScreenMenu }
		}
	}
	return m, nil
}

func (m *skillsModel) View() string {
	theme := styles.NewThemeFromName("default")

	// Styles
	titleStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true).
		Align(lipgloss.Center).
		MarginBottom(1)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Align(lipgloss.Center).
		Italic(true).
		MarginBottom(2)

	tabActiveStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true).
		Underline(true).
		Padding(0, 2)

	tabInactiveStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Padding(0, 2)

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Primary).
		Padding(2, 4)

	itemBoxStyle := lipgloss.NewStyle().
		Padding(1, 2).
		Margin(1, 0)

	selectedItemStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	normalItemStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary)

	cursorStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(theme.Help).
		MarginTop(2)

	// Build output
	var output string
	output += titleStyle.Render("🛠️  Skills & Technologies") + "\n"
	output += subtitleStyle.Render("Technical expertise across the stack") + "\n\n"

	// Category tabs - simpler, aligned
	var tabs string
	for i, c := range m.categories {
		if i == m.catIndex {
			tabs += tabActiveStyle.Render(c.Name)
		} else {
			tabs += tabInactiveStyle.Render(c.Name)
		}
		if i < len(m.categories)-1 {
			tabs += "  "
		}
	}
	output += tabs + "\n\n"

	// Items
	items := m.categories[m.catIndex].Items
	if len(items) == 0 {
		output += normalItemStyle.Render("(no items in this category)") + "\n"
		output += helpStyle.Render("\nPress ESC to go back")
		return containerStyle.Render(output)
	}

	// Item list - clean and aligned
	var itemsList string
	for i, it := range items {
		if i == m.cursor {
			itemsList += cursorStyle.Render("▸ ") + selectedItemStyle.Render("• "+it) + "\n"
		} else {
			itemsList += "  " + normalItemStyle.Render("• "+it) + "\n"
		}
	}

	output += itemBoxStyle.Render(itemsList)
	output += helpStyle.Render("\n←/→ switch tabs • ↑/↓ navigate • ESC back • q quit")

	return "\n" + containerStyle.Render(output)
}
