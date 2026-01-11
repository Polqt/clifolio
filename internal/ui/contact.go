package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"fmt"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ContactInfo struct {
	Label string
	Value string
	Icon  string
	Link  string
}

type contactModel struct {
	contacts  []ContactInfo
	cursor    int
	theme     styles.Theme
	width     int
	height    int
	keymap    components.Keymap
	copiedMsg string
	showQR    bool
}

func ContactModel() tea.Model {
	theme := styles.NewThemeFromName("default")
	return NewContactModel(theme)
}

func NewContactModel(theme styles.Theme) *contactModel {
	contacts := []ContactInfo{
		{
			Label: "LinkedIn",
			Value: "https://www.linkedin.com/in/janpol-hidalgo",
			Icon:  "💼",
		},
		{
			Label: "GitHub",
			Value: "github.com/Polqt",
			Icon:  "🐙",
		},
		{
			Label: "Email",
			Value: "poyhidalgo@gmail.com",
			Icon:  "📧",
		},
		{
			Label: "Portfolio",
			Value: "https://yojepoy.vercel.app/",
			Icon:  "🌐",
		},
	}

	return &contactModel{
		contacts: contacts,
		theme:    theme,
		keymap:   components.DefaultKeymap(),
	}
}

func (m *contactModel) Init() tea.Cmd {
	return nil
}

func (m *contactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case m.keymap.Quit, "ctrl+c":
			return m, tea.Quit
		case m.keymap.Up, "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.contacts) - 1
			}
			m.copiedMsg = ""
		case m.keymap.Down, "down", "j":
			if m.cursor < len(m.contacts)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
			m.copiedMsg = ""
		case "c":
			if m.cursor < len(m.contacts) {
				contact := m.contacts[m.cursor]
				err := clipboard.WriteAll(contact.Value)
				if err == nil {
					m.copiedMsg = fmt.Sprintf("Copied %s to clipboard!", contact.Label)
				} else {
					m.copiedMsg = "Failed to copy to clipboard."
				}
			}
		case m.keymap.Back, "esc":
			return m, func() tea.Msg { return state.ScreenMenu }
		}
	}
	return m, nil
}

func (m *contactModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sections []string

	header := components.HeaderBox("SUMMON THE DEV-WARRIOR", m.theme, m.width-4)
	sections = append(sections, header)

	intro := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Italic(true).
		Align(lipgloss.Center).
		Width(m.width).
		Render("The warrior awaits your summons for new quests and alliances!")

	sections = append(sections, intro)

	sections = append(sections, components.DividerLine(m.theme, m.width-2, "─"))

	contactList := m.renderContactList()
	sections = append(sections, contactList)

	if m.copiedMsg != "" {
		msgStyle := lipgloss.NewStyle().
			Foreground(styles.Success).
			Bold(true).
			Align(lipgloss.Center).
			Width(m.width).
			Margin(1, 0)
		sections = append(sections, msgStyle.Render(m.copiedMsg))
	}

	keyBindings := []components.KeyBind{
		{Key: "↑↓/k/j", Desc: "Navigate"},
		{Key: "c", Desc: "Copy to Clipboard"},
		{Key: "b/Esc", Desc: "Retreat"},
	}

	footer := components.RenderKeyBindings(keyBindings, m.theme, m.width)
	sections = append(sections, footer)

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (m *contactModel) renderContactList() string {
	var cards []string

	for i, contact := range m.contacts {
		isSelected := i == m.cursor

		// Label with better styling
		labelStyle := lipgloss.NewStyle().
			Foreground(m.theme.Accent).
			Bold(true).
			Underline(isSelected)

		// Value with subtle styling
		valueStyle := lipgloss.NewStyle().
			Foreground(m.theme.Primary).
			Italic(true)

		// Build the contact content
		contactContent := lipgloss.JoinVertical(
			lipgloss.Left,
			labelStyle.Render(contact.Label),
			valueStyle.Render(contact.Value),
		)

		// Card styling
		var card string
		if isSelected {
			cardStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(m.theme.Accent).
				Padding(1, 3).
				MarginBottom(1).
				Width(m.width - 16).
				BorderStyle(lipgloss.ThickBorder())

			card = cardStyle.Render(contactContent)
		} else {
			cardStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(m.theme.Secondary).
				Padding(1, 3).
				MarginBottom(1).
				Width(m.width - 16)

			card = cardStyle.Render(contactContent)
		}

		cards = append(cards, card)
	}

	allCards := lipgloss.JoinVertical(lipgloss.Left, cards...)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		allCards,
	)
}
