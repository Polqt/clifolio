package ui

import (
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
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
				Link: "",
			},
			{
				Name: "GitHub",
				Link: "",
			},
			{
				Name: "Email",
				Link: "",
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
			case km.Back, "esc":
				return m, func() tea.Msg { return state.ScreenMenu }
		}
	}
	return m, nil
}

func (m *contactModel) View() string {
	s := "\n\nContact\n\n"
	for i, contact := range m.contacts {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s: %s\n", cursor, contact.Name, contact.Link)
	}
	return s
}