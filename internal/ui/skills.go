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
		Background(lipgloss.Color("#1a1a2e")).
		Bold(true).
		Padding(1, 3).
		Margin(0, 1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(theme.Accent)

	tabInactiveStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Padding(1, 3).
		Margin(0, 1)

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(theme.Primary).
		Padding(3, 5).
		Align(lipgloss.Center)

	itemBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Secondary).
		Padding(2, 3).
		Margin(1, 0).
		Width(55)

	selectedItemStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	normalItemStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary)

	helpStyle := lipgloss.NewStyle().
		Foreground(theme.Help).
		Italic(true).
		MarginTop(2).
		Align(lipgloss.Center)

	dividerStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Align(lipgloss.Center)

	// Title
	var output string
	output += titleStyle.Render("🛠️  Skills & Technologies") + "\n"
	output += subtitleStyle.Render("Technical expertise across the stack") + "\n"
	output += dividerStyle.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━") + "\n\n"

	// Category tabs
	var tabs string
	for i, c := range m.categories {
		if i == m.catIndex {
			tabs += tabActiveStyle.Render("[ " + c.Name + " ]")
		} else {
			tabs += tabInactiveStyle.Render(c.Name)
		}
	}
	output += lipgloss.NewStyle().Align(lipgloss.Center).Render(tabs) + "\n\n"

	// Items
	items := m.categories[m.catIndex].Items
	if len(items) == 0 {
		output += normalItemStyle.Render("(no items in this category)") + "\n"
		output += helpStyle.Render("\nPress ESC to go back\n")
		return containerStyle.Render(output)
	}

	// Item list with better formatting and columns
	var itemsList string
	for i, it := range items {
		var cursor string
		var style lipgloss.Style

		if i == m.cursor {
			cursor = "▸  "
			style = selectedItemStyle
		} else {
			cursor = "   "
			style = normalItemStyle
		}

		itemsList += cursor + style.Render("● "+it)
		if i < len(items)-1 {
			itemsList += "\n"
		}
	}

	output += itemBoxStyle.Render(itemsList) + "\n"
	output += dividerStyle.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━") + "\n"
	output += helpStyle.Render("←/→ switch tabs • ↑/↓ navigate • ESC back • q quit")

	return "\n" + containerStyle.Render(output)
}
