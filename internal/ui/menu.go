package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	open   bool
	items  []state.Screen
	idx    int
	width  int
	height int
}

func MenuModel() tea.Model {
	return &menuModel{
		items: []state.Screen{
			state.ScreenProjects,
			state.ScreenSkills,
			state.ScreenExperience,
			state.ScreenContact,
			state.ScreenTheme,
			state.ScreenStats,
		},
	}
}

func (m *menuModel) Init() tea.Cmd {
	return nil
}

func (m *menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

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
		case km.Down, "down":
			if m.idx < len(m.items)-1 {
				m.idx++
			} else {
				m.idx = 0
			}
		case km.Up, "up":
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
	theme := styles.NewThemeFromName("default")

	// Enhanced styling
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

	selectedStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	normalStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary)

	cursorStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(theme.Help).
		Italic(true).
		MarginTop(2).
		Align(lipgloss.Center)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(theme.Primary).
		Padding(3, 6).
		Align(lipgloss.Center)

	dividerStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Align(lipgloss.Center)

	var content string

	// Title with better formatting
	content += titleStyle.Render("Portfolio Navigator") + "\n"
	content += subtitleStyle.Render("What would you like to explore?") + "\n"
	content += dividerStyle.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━") + "\n\n"

	// Menu items with better spacing
	for i, it := range m.items {
		icon := getScreenIcon(it)

		if i == m.idx {
			cursor := cursorStyle.Render("  ▸  ")
			content += cursor + selectedStyle.Render(icon+"  "+it.String()) + "\n"
		} else {
			content += "     " + normalStyle.Render(icon+"  "+it.String()) + "\n"
		}
		if i < len(m.items)-1 {
			content += "\n"
		}
	}

	// Help text
	content += "\n" + dividerStyle.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━") + "\n"
	content += helpStyle.Render("↑/↓ navigate • enter select • m matrix • q quit")

	box := boxStyle.Render(content)

	// Center the menu box
	if m.width > 0 && m.height > 0 {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
	}
	return "\n" + box
}

func getScreenIcon(s state.Screen) string {
	switch s {
	case state.ScreenProjects:
		return "📁"
	case state.ScreenSkills:
		return "🛠️"
	case state.ScreenExperience:
		return "💼"
	case state.ScreenContact:
		return "📧"
	case state.ScreenTheme:
		return "🎨"
	case state.ScreenStats:
		return "📊"
	default:
		return "•"
	}
}
