package ui

import (
	"context"
	"fmt"
	"time"

	"clifolio/internal/services"
	"clifolio/internal/ui/components"

	tea "github.com/charmbracelet/bubbletea"
)

type projectsModel struct {
	username string
	projects []services.Repo
	cursor  int
	loading bool
	err error

	spin components.SpinnerComponent
}

type projectsLoadedMsg struct {
	projects []services.Repo
}

type projectsErrMsg struct {
	err error
}

func ProjectsModel(username string) *projectsModel {
	return &projectsModel{
		username: username,
		loading: true,
		spin:  components.NewSpinner(),
	}
}

func fetchReposCmd(username string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		repos, err := services.FetchRepos(ctx, username)
		if err != nil {
			return projectsErrMsg{err}
		}
		return projectsLoadedMsg{projects: repos}
	}
}

func (m *projectsModel) Init() tea.Cmd {
	return tea.Batch(m.spin.Init(), fetchReposCmd(m.username))
}

func (m *projectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	newSpin, spinCmd := m.spin.Update(msg)
	m.spin = newSpin
	cmds = append(cmds, spinCmd)

	switch msg := msg.(type) {
	case projectsLoadedMsg:
		m.projects = msg.projects
		m.loading = false
		return m, nil

	case projectsErrMsg:
		m.err = msg.err
		m.loading = false
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if !m.loading && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if !m.loading && m.cursor < len(m.projects)-1 {
				m.cursor++
			}
		}
	}

	if m.loading {
		cmds = append(cmds, m.spin.Init())
	}

	return m, tea.Batch(cmds...)
}


func (m *projectsModel) View() string {
	if m.loading {
		return fmt.Sprintf("\n\n %s Loading Github repos...", m.spin.View())
	}

	if m.err != nil {
		return fmt.Sprintf("\n\n Error: %s", m.err)
	}

	if len(m.projects) == 0 {
		return "\n\n No repositories found."
	}

	s := "\n\n Projects of " + m.username + "\n\n"
	for i, repo := range m.projects {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf(" %s %s (%d ★)\n    %s\n\n", cursor, repo.Name, repo.Stars, repo.Description)
	}

	s += "\nPress q to quit.\n"
	return s
}
