package ui

import (
	"clifolio/internal/services"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type projectDetailsModel struct {
	project 		services.Repo
	rawMD			string
	rendered		string
	loaded 		    bool
	err				error
}

func ProjectDetailsModel(r services.Repo, md string) projectDetailsModel {
	return projectDetailsModel{
		project: r,
		rawMD: md,
		loaded: false,
	}
}

func (m projectDetailsModel) Init() tea.Cmd {
	return func() tea.Msg {
		out, err := services.GenerateMarkdown(m.rawMD)
		if err != nil {
			return fmt.Errorf("M")
		}
		return out
	}
}

func (m projectDetailsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case string:
		m.rendered = msg
		m.loaded = true
	case error:
		m.err = msg
	case tea.KeyMsg:
		if msg.String() == "b" || msg.String() == "esc" {
			return m, func() tea.Msg { return nil }
		}
	}
	return m, nil
}

func (m projectDetailsModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error loading project details: %v", m.err)
	}
	if !m.loaded {
		return "\n\n Rendering markdown..."
	}

	s := "\n\n" + m.project.Name + "\n\n"
	s += m.rendered
	s += "\n\nPress 'b' or 'esc' to go back."
	return s
}
