package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ThemeChangeMsg struct {
	ThemeName string
}

type ThemeInfo struct {
	Name        string
	DisplayName string
	Icon        string
	Description string
	Preview     string
}

type themePickerModel struct {
	themes       []ThemeInfo
	cursor       int
	theme        styles.Theme
	width        int
	height       int
	keymap       components.Keymap
	previewTheme string
}

func ThemePickerModel() *themePickerModel {
	theme := styles.NewThemeFromName("default")
	return NewThemePickerModel(theme)
}

func NewThemePickerModel(theme styles.Theme) *themePickerModel {
	themes := []ThemeInfo{
		{
			Name:        "default",
			DisplayName: "Solarized Dark",
			Icon:        "🌙",
			Description: "Classic solarized dark theme - Easy on the eyes",
			Preview:     "Warm & Professional",
		},
		{
			Name:        "hacker",
			DisplayName: "Matrix Hacker",
			Icon:        "💻",
			Description: "Green terminal vibes - Enter the Matrix",
			Preview:     "Green & Bold",
		},
		{
			Name:        "dracula",
			DisplayName: "Dracula",
			Icon:        "🧛",
			Description: "Dark with vibrant accents - Modern & Stylish",
			Preview:     "Purple & Pink",
		},
	}

	return &themePickerModel{
		themes:       themes,
		cursor:       0,
		theme:        theme,
		keymap:       components.DefaultKeymap(),
		previewTheme: "default",
	}
}

func (m *themePickerModel) Init() tea.Cmd {
	return nil
}

func (m *themePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.cursor = len(m.themes) - 1
			}
			m.previewTheme = m.themes[m.cursor].Name
		case m.keymap.Down, "down", "j":
			if m.cursor < len(m.themes)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
			m.previewTheme = m.themes[m.cursor].Name
		case m.keymap.Confirm, "enter", " ":
			if m.cursor >= 0 && m.cursor < len(m.themes) {
				selectedTheme := m.themes[m.cursor].Name
				return m, func() tea.Msg { return ThemeChangeMsg{ThemeName: selectedTheme} }
			}
		case m.keymap.Back, "esc", "b":
			return m, func() tea.Msg { return state.ScreenMenu }
		}
	}
	return m, nil
}

func (m *themePickerModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sections []string

	header := components.HeaderBox("THEME SELECTOR", m.theme, m.width-4)
	sections = append(sections, header)

	intro := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Italic(true).
		Align(lipgloss.Center).
		Width(m.width).
		Render("Customize your portfolio experience!")

	sections = append(sections, intro)
	sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

	themeList := m.renderThemeList()
	sections = append(sections, themeList)

	preview := m.renderPreview()
	sections = append(sections, preview)

	keyBindings := []components.KeyBind{
		{Key: "↑↓/k/j", Desc: "Navigate"},
		{Key: "Enter", Desc: "Apply Theme"},
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

func (m *themePickerModel) renderThemeList() string {
	items := make([]components.ListItem, len(m.themes))

	for i, themeInfo := range m.themes {
		items[i] = components.ListItem{
			Title:   themeInfo.DisplayName,
			Content: themeInfo.Description,
			Icon:    themeInfo.Icon,
			Badge:   themeInfo.Preview,
		}
	}

	listStyle := components.ListStyle{
		ShowNumbers:    false,
		ShowIcons:      true,
		ShowBadges:     true,
		CompactMode:    false,
		HighlightColor: m.theme.Accent.(lipgloss.Color),
	}

	list := components.RenderList(items, m.cursor, m.theme, listStyle)
	return components.SectionBox("Available Themes", list, m.theme, m.width-8)
}

func (m *themePickerModel) renderPreview() string {
	previewTheme := styles.NewThemeFromName(m.previewTheme)

	titleStyle := lipgloss.NewStyle().
		Foreground(previewTheme.Primary).
		Bold(true).
		Align(lipgloss.Center)

	accentStyle := lipgloss.NewStyle().
		Foreground(previewTheme.Accent).
		Bold(true).
		Align(lipgloss.Center)

	secondaryStyle := lipgloss.NewStyle().
		Foreground(previewTheme.Secondary).
		Align(lipgloss.Center)

	previewContent := lipgloss.JoinVertical(
		lipgloss.Center,
		titleStyle.Render("● Primary Color"),
		accentStyle.Render("● Accent Color"),
		secondaryStyle.Render("● Secondary Color"),
	)

	preview := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(previewTheme.Primary).
		Padding(1, 2).
		Render(previewContent)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		components.SectionBox("Theme Preview", preview, m.theme, m.width-8),
	)
}
