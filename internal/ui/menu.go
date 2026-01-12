package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	cursor  int
	choices []components.ListItem
	search  textinput.Model
	theme   styles.Theme
	width   int
	height  int
	open    bool
}

func MenuModel() tea.Model {
	theme := styles.NewThemeFromName("default")
	return NewMenuModel(theme)
}

func NewMenuModel(theme styles.Theme) *menuModel {
	ti := textinput.New()
	ti.Placeholder = "Search commands..."
	ti.CharLimit = 50
	ti.Width = 40

	choices := []components.ListItem{
		{
			Title:   "Battle Records",
			Content: "Chronicles of completed quests",
			Icon:    "⚡",
			Badge:   "GitHub",
		},
		{
			Title:   "Abilities",
			Content: "Warrior's skills and mastery",
			Icon:    "✨",
			Badge:   "Tech Stack",
		},
		{
			Title:   "Combat History",
			Content: "Journey through battles past",
			Icon:    "📖",
			Badge:   "Career",
		},
		{
			Title:   "Summon Warrior",
			Content: "Reach out to the Dev-Warrior",
			Icon:    "📜",
			Badge:   "Social",
		},
		{
			Title:   "Warrior Stats",
			Content: "Live battle statistics",
			Icon:    "📊",
			Badge:   "Analytics",
		},
		{
			Title:   "Change Realm",
			Content: "Shift between realms",
			Icon:    "🌙",
			Badge:   "Customize",
		},
		{
			Title:   "Matrix Realm",
			Content: "Enter the forbidden realm...",
			Icon:    "🟢",
			Badge:   "Secret",
		},
	}

	return &menuModel{
		cursor:  0,
		choices: choices,
		search:  ti,
		theme:   theme,
		open:    true,
	}
}

func (m *menuModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		// Handle search input
		if m.search.Focused() {
			switch msg.String() {
			case "esc", km.Back:
				m.search.Blur()
				return m, nil
			case km.Confirm:
				m.search.Blur()
				// Filter choices based on search
				searchTerm := strings.ToLower(m.search.Value())
				if searchTerm != "" {
					for i, choice := range m.choices {
						if strings.Contains(strings.ToLower(choice.Title), searchTerm) {
							m.cursor = i
							break
						}
					}
				}
				return m, nil
			}
			m.search, cmd = m.search.Update(msg)
			return m, cmd
		}

		// Handle navigation
		switch msg.String() {
		case km.Up, "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}
		case km.Down, "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case km.Toggle, "/":
			m.search.Focus()
			return m, textinput.Blink
		case km.Quit, "ctrl+c", "q":
			return m, tea.Quit
		case km.Confirm, "enter", " ":
			selectedScreen := m.getSelectedScreen()
			return m, func() tea.Msg {
				return selectedScreen
			}
		}
	}

	return m, nil
}

func (m *menuModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sections []string

	// Header
	header := components.HeaderBox("WARRIOR'S COMMAND REALM", m.theme, m.width-2)
	sections = append(sections, header)

	// Subtitle
	subtitle := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Italic(true).
		Align(lipgloss.Center).
		Width(m.width).
		Render("Choose your path, Warrior")
	sections = append(sections, subtitle)

	// Divider
	sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

	// Search box (if focused)
	if m.search.Focused() {
		searchBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(m.theme.Accent).
			Padding(0, 1).
			Width(m.width - 10).
			Align(lipgloss.Center).
			Render(m.search.View())
		sections = append(sections, lipgloss.PlaceHorizontal(m.width, lipgloss.Center, searchBox))
		sections = append(sections, "\n")
	}

	// Menu list
	listStyle := components.ListStyle{
		ShowNumbers:    false,
		ShowIcons:      true,
		ShowBadges:     true,
		CompactMode:    false,
		HighlightColor: m.theme.Accent.(lipgloss.Color),
	}

	list := components.RenderList(m.choices, m.cursor, m.theme, listStyle)
	listBox := components.SectionBox("", list, m.theme, m.width-8)
	sections = append(sections, lipgloss.PlaceHorizontal(m.width, lipgloss.Center, listBox))

	keyBindings := []components.KeyBind{
		{Key: "↑↓/j/k", Desc: "Navigate"},
		{Key: "Enter", Desc: "Select"},
		{Key: "/", Desc: "Search"},
		{Key: "q", Desc: "Retreat"},
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

func (m *menuModel) getSelectedScreen() state.Screen {
	screens := []state.Screen{
		state.ScreenProjects,
		state.ScreenSkills,
		state.ScreenExperience,
		state.ScreenContact,
		state.ScreenStats,
		state.ScreenTheme,
		state.ScreenMatrix,
	}

	if m.cursor >= 0 && m.cursor < len(screens) {
		return screens[m.cursor]
	}
	return state.ScreenMenu
}
