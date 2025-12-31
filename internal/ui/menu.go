package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
	cursor 		int
	choices 	[]components.ListItem
	search		textinput.Model
	theme 		styles.Theme
	width		int
	height 		int
}

func NewMenuModel(theme styles.Theme) MenuModel {
	ti := textinput.New()
	ti.Placeholder = "Search commands..."
	ti.CharLimit = 50
	ti.Width = 40

	choices := []components.ListItem{
		{
			Title: "Projects",
			Content: "",
			Icon: "📦",
			Badge: "GitHub",
		},
		{
			Title: "Skills",
			Content: "",
			Icon: "⚡",
			Badge: "Tech Stack",
		},
		{
			Title: "Experience",
			Content: "",
			Icon: "💼",
			Badge: "Career",
		},
		{
			Title: "Contact",
			Content: "",
			Icon: "📧",
			Badge: "Social",
		},
		{
			Title: "Themes",
			Content: "",
			Icon: "🎨",
			Badge: "Customize",
		},
		{
			Title: "About",
			Content: "",
			Icon: "ℹ️",
			Badge: "Info",
		},
		{
			Title: "Matrix Mode",
			Content: "",
			Icon: "🟢",
			Badge: "Easter Egg",
		},
	}
	return MenuModel{
		cursor: 0,
		choices: choices,
		search: ti,
		theme: theme,
	}
}

func (m MenuModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	
	case tea.KeyMsg:
		if m.search.Focused() {
			switch msg.String() {
			case "esc", km.Back:
				m.search.Blur()
				return m, nil
			case km.Confirm:
				m.search.Blur()
				return m, nil
			}
			m.search, cmd = m.search.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		case km.Up, "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case km.Down, "down":
			if m.cursor < len(m.choices) - 1 {
				m.cursor++
			}
		case km.Toggle:
			m.search.Focus()
			return m, textinput.Blink
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		case km.Confirm, " ":
			return m, func() tea.Msg {
				return NavigateMsg{Screen: m.getSelectedScreen()}
			}
		}
	}

	return m, nil
}

func (m MenuModel) View() string {
	if m.width == 0 { 
		return "Loading..."
	}

	var sections []string

	header := components.HeaderBox("COMMAND PALETTE", m.theme, m.width-4)
	sections = append(sections, header)

	subtitle := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Italic(true).
		Align(lipgloss.Center).
		Width(m.width).
		Render("Navigate through my portfolio")
	sections = append(sections, subtitle)

	sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

	if m.search.Focused() {
		searchBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(m.theme.Accent).
			Padding(0, 1).
			Width(m.width - 10).
			Align(lipgloss.Center).
			Render(m.search.View())
		sections = append(sections, lipgloss.PlaceHorizontal(m.width, lipgloss.Center, searchBox))
	}

	listStyle := components.ListStyle{
		ShowNumbers: false,
		ShowIcons: true,
		ShowBadges: true,
		CompactMode: false,
		HighlightColor: m.theme.Accent.(lipgloss.Color),
	}

	list := components.RenderList(m.choices, m.cursor, m.theme, listStyle)
	listBox := components.SectionBox("Available Commands", list, m.theme, m.width-8)
	sections = append(sections, lipgloss.PlaceHorizontal(m.width, lipgloss.Center, listBox))


	keyBindings := []components.KeyBind{
		{Key: "↑↓", Desc: "Navigate"},
        {Key: "Enter", Desc: "Select"},
        {Key: "/", Desc: "Search"},
        {Key: "q", Desc: "Quit"},
	}

	footer := components.RenderKeyBindings(keyBindings, m.theme, m.width)
	sections = append(sections, footer)

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)
	
	return lipgloss.Place(
		m.width,
		m.cursor,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (m MenuModel) getSelectedScreen() string {
	screens := map[int]string{
		0: "projects",
        1: "skills",
        2: "experience",
        3: "contact",
        4: "themes",
        5: "about",
        6: "matrix",
	}
	return screens[m.cursor]
}