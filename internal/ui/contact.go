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
	Label 	string
	Value 	string
	Icon 	string
	Link  	string
}

type contactModel struct {
	contacts 	[]ContactInfo
	cursor   	int
	theme 		styles.Theme
	width   	int
	height 		int
	keymap  	components.Keymap
	copiedMsg 	string
	showQR   	bool
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
			Icon: "💼",
			Link: "https://www.linkedin.com/in/janpol-hidalgo-64174a241/",
		},
		{
			Label: "GitHub",
			Value: "github.com/Polqt",
			Icon: "🐙",
			Link: "github.com/Polqt",
		},
		{
			Label: "Email",
			Value: "poyhidalgo@gmail.com",
			Icon: "📧",
			Link: "mailto:poyhidalgo@gmail.com",
		},
		{
			Label: "Portfolio",
			Value: "https://yojepoy.vercel.app/",
			Icon: "🌐",
			Link: "https://yojepoy.vercel.app/",
		},
	}

	return &contactModel{
		contacts: contacts,
		theme: theme,
		keymap: components.DefaultKeymap(),
	}
}

func (m *contactModel) Init() tea.Cmd {
	return nil
}

func (m *contactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case m.keymap.Quit, "ctrl+c":
			return m, tea.Quit
		case m.keymap.Up, "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			m.copiedMsg = ""
		case m.keymap.Down, "down", "j":
			if m.cursor < len(m.contacts)-1 {
				m.cursor++
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
		case "q":
			m.showQR = !m.showQR
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

	header := components.HeaderBox("GET IN TOUCH", m.theme, m.width - 4)
	sections = append(sections, header)

	intro := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Italic(true).
		Align(lipgloss.Center).
		Width(m.width).
		Render("I'm always open to new opportunities and collaborations!")
	
	sections = append(sections, intro)

	sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

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

	if m.showQR {
		qrCode := m.renderQRCode()
		sections = append(sections, qrCode)
	}

	socialCards := m.renderSocialCards()
	sections = append(sections, socialCards)

	keyBindings := []components.KeyBind{
		{Key: "↑↓/k/j", Desc: "Navigate"},
        {Key: "c", Desc: "Copy"},
        {Key: "q", Desc: "QR Code"},
        {Key: "b/Esc", Desc: "Back"},
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
	items := make([]components.ListItem, len(m.contacts))

	for i, contact := range m.contacts {
		desc := contact.Value
		if contact.Link != "" {
			desc += "\n🔗" + contact.Link
		}

		items[i] = components.ListItem{
			Title:   	contact.Label,
			Content: 	desc,
			Icon: 		contact.Icon,
			Badge:  	"Copy",		

		}
	}

	listStyle := components.ListStyle{
		ShowNumbers:		false,
		ShowIcons:	 		true,
		ShowBadges: 		true,
		CompactMode: 		false,
		HighlightColor: 	m.theme.Accent.(lipgloss.Color),
	}

	list := components.RenderList(items, m.cursor, m.theme, listStyle)
	return components.SectionBox("Contact Information", list, m.theme, m.width-8)
}

func (m *contactModel) renderSocialCards() string {
	cardWidth := (m.width - 12) / 2

	githubCard := m.createSocialCard("GitHub", "🐙", "@Polqt", "30+ repos", cardWidth, 0)
	linkedinCard := m.createSocialCard("LinkedIn", "💼", "Janpol Hidalgo", "500+ connections", cardWidth, 1)
	portfolioCard := m.createSocialCard("Portfolio", "🌐", "yojepoy.vercel.app", "Check out my work", cardWidth, 2)

	cards := lipgloss.JoinHorizontal(
		lipgloss.Top,
		githubCard,
		"  ",
		linkedinCard,
		"  ",
		portfolioCard,
	)

	return lipgloss.PlaceHorizontal(m.width, lipgloss.Center, cards)
}

func (m *contactModel) createSocialCard(title, icon, handle, meta string, width, index int) string {
    isSelected := m.cursor == index

    titleStyle := lipgloss.NewStyle().
        Foreground(m.theme.Primary).
        Bold(true).
        Align(lipgloss.Center)

    iconStyle := lipgloss.NewStyle().
        Align(lipgloss.Center)

    handleStyle := lipgloss.NewStyle().
        Foreground(m.theme.Accent).
        Align(lipgloss.Center)

    metaStyle := lipgloss.NewStyle().
        Foreground(m.theme.Secondary).
        Italic(true).
        Align(lipgloss.Center)

    content := lipgloss.JoinVertical(
        lipgloss.Center,
        iconStyle.Render(icon),
        titleStyle.Render(title),
        handleStyle.Render(handle),
        metaStyle.Render(meta),
    )

    return components.CardBox(content, m.theme, isSelected)
}

func (m *contactModel) renderQRCode() string {
	qrPlaceHolder := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Border(lipgloss.RoundedBorder()).
		Padding(2).
		BorderForeground(m.theme.Primary).
		Align(lipgloss.Center).
		Render("QR Code would appear here\n.")
	
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		components.SectionBox("Scan to Connect", qrPlaceHolder, m.theme, m.width-8),
	)
}