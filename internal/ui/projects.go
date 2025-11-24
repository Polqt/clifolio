package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type projectsModel struct {
	projects []string
	cursor  int
	selected map[int]struct{}
}

func ProjectsModel() projectsModel {
	return projectsModel{
		projects: []string{"Tandaan", "Pharmafetch", "AI Career Assistant"},
		selected: make(map[int]struct{}),
	}
}

func (m projectsModel) Init() tea.Cmd {
	return nil
}

func (m projectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		
		}
	}
	return m, nil
}

func (m projectsModel) View() string {
	s := "\n\n Here are some of my projects:\n\n"
	for i, project := range m.projects {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, project)
	}

	s += "\nPress q to quit.\n"
	return s
}

func RunProjects() {
	p := tea.NewProgram(ProjectsModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}