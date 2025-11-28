package ui

import (
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
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