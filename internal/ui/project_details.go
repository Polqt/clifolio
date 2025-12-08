package ui

import (
	"clifolio/internal/services"
	"clifolio/internal/ui/components"
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

type backToProjectsMsg struct{}

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
			return err
		}
		return out
	}
}

func (m projectDetailsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()

	switch msg := msg.(type) {
	case string:
		m.rendered = msg
		m.loaded = true
	case error:
		m.err = msg
	case tea.KeyMsg:
		switch msg.String() {
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		case km.Back, "esc":
			return m, func() tea.Msg { return backToProjectsMsg{} }
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
