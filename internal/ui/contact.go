package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Contact struct {
	Name string
	Link string
}

type contactModel struct {
	contacts []Contact
	cursor   int
	catIndex int
}

func ContactModel() *contactModel {
	return &contactModel{
		contacts: []Contact{
			{
				Name: "LinkedIn",
				Link: "https://www.linkedin.com/in/janpol-hidalgo-64174a241/",
			},
			{
				Name: "GitHub",
				Link: "github.com/Polqt",
			},
			{
				Name: "Email",
				Link: "poyhidalgo@gmail.com",
			},
			{
				Name: "Portfolio",
				Link: "https://yojepoy.vercel.app/",
			},
		},
		cursor:   0,
		catIndex: 0,
	}
}

func (m *contactModel) Init() tea.Cmd {
	return nil
}

func (m *contactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		case km.Up, "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case km.Down, "down", "j":
			if m.cursor < len(m.contacts)-1 {
				m.cursor++
			}
		case km.Back, "esc":
			return m, func() tea.Msg { return state.ScreenMenu }
		}
	}
	return m, nil
}

func (m *contactModel) View() string {
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

	nameStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true)

	normalContactStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary)

	cursorStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	linkStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00d4ff")).
		Underline(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(theme.Help).
		MarginTop(1)

	// Get contact icon
	getIcon := func(name string) string {
		switch name {
		case "LinkedIn":
			return "💼"
		case "GitHub":
			return "🐙"
		case "Email":
			return "📧"
		case "Portfolio":
			return "🌐"
		default:
			return "📱"
		}
	}

	// Build content
	var content string
	content += titleStyle.Render("📧 Get In Touch") + "\n"
	content += subtitleStyle.Render("Let's connect and collaborate") + "\n\n"

	// Contact items - simple list
	for i, contact := range m.contacts {
		icon := getIcon(contact.Name)

		if i == m.cursor {
			content += cursorStyle.Render("▸ ") + nameStyle.Render(icon+"  "+contact.Name) + "\n"
			content += "  " + linkStyle.Render(contact.Link) + "\n\n"
		} else {
			content += "  " + normalContactStyle.Render(icon+"  "+contact.Name) + "\n"
			content += "  " + linkStyle.Render(contact.Link) + "\n\n"
		}
	}

	content += helpStyle.Render("↑/↓ navigate • ESC back • q quit")

	return "\n" + containerStyle.Render(content)
}
