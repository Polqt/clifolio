package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Primary).
		Padding(2, 4)

	cardStyle := lipgloss.NewStyle().
		Padding(2, 3).
		Margin(1, 0)

	selectedCardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Accent).
		Padding(2, 3).
		Margin(1, 0)

	positionStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	companyStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true)

	dateStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Italic(true)

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c0c0c0")).
		MarginTop(1)

	helpStyle := lipgloss.NewStyle().
		Foreground(theme.Help).
		MarginTop(1)

	cursorStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	// Title
	var output string
	output += titleStyle.Render("💼 Professional Experience") + "\n"
	output += subtitleStyle.Render("My journey in software development") + "\n\n"

	// Experience cards
	for i, exp := range m.experiences {
		var card string

		// Build card content
		card += positionStyle.Render(exp.Position) + "\n"
		card += companyStyle.Render("@ "+exp.Company) + "\n"
		card += dateStyle.Render("📅 "+exp.Date) + "\n\n"
		card += descStyle.Render(exp.Description)

		// Apply card style based on selection
		if i == m.cursor {
			output += cursorStyle.Render("▸ ") + selectedCardStyle.Render(card) + "\n"
		} else {
			output += "  " + cardStyle.Render(card) + "\n"
		}
	}

	// Help text
	output += helpStyle.Render("\n↑/↓ navigate • ESC back • q quit")

	return "\n" + containerStyle.Render(output)
}
