package ui

import (
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ThemeChangeMsg struct {
	ThemeName string
}

type themePickerModel struct {
	themes []string
	cursor int
}

func ThemePickerModel() *themePickerModel {
	return &themePickerModel{
		themes: []string{
			"default",
			"hacker",
			"dracula",
		},
		cursor: 0,
	}
}

func (m *themePickerModel) Init() tea.Cmd {
	return nil
}

func (m *themePickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		case km.Back, "esc":
			return m, func() tea.Msg { return state.ScreenMenu }
		
		}
	}
	return m, nil
}

func (m *themePickerModel) View() string {
	s := "\n\n🎨 Theme Picker\n\n"
	
	for i, theme := range m.themes {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s+= fmt.Sprintf("%s %s\n", cursor, theme)
	}

	s += "\nUse the arrow keys to navigate, Enter to select, Esc to go back.\n"
	return s
}