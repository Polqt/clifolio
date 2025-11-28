package ui

import (
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
)


type appModel struct {
	screen state.Screen

	intro tea.Model
	menu tea.Model
	projects tea.Model
	skills tea.Model
	experience tea.Model
	contact tea.Model
}

func AppModel() appModel {
	return appModel{
		screen: state.Intro,
		intro: IntroModel(),
		menu: MenuModel(),
		projects: ProjectsModel("Polqt"),
		skills: SkillsModel(),
		experience: ExperienceModel(),
		contact: ContactModel(),
	}
}

func (m appModel) Init() tea.Cmd {
	return m.intro.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	// Global exit
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "ctrl+c" || key.String() == "q" {
			return m, tea.Quit
		}
	}

	// Handle intro -> menu
	if _, ok := msg.(goToMenuMsg); ok {
		m.screen = state.Menu
		return m, nil
	}

	switch m.screen {
	case state.Intro:
		newIntro, cmd := m.intro.Update(msg)
		m.intro = newIntro
		return m, cmd
	
	case state.Menu:
		newMenu, cmd := m.menu.Update(msg)
		m.menu = newMenu
		if cmd != nil {
			return m, cmd
		}
		switch msg := msg.(type) {
		case state.Screen:
			m.screen = msg
			switch msg {
			case state.Projects:
				return m, m.projects.Init()
			}
		}
		return m, nil

	case state.Projects:
		newProjects, cmd := m.projects.Update(msg)
		m.projects = newProjects
		return m, cmd
	
	case state.Skills:
		newSkills, cmd := m.skills.Update(msg)
		m.skills = newSkills
		return m, cmd

	case state.Experience:
		newExperience, cmd := m.experience.Update(msg)
		m.experience = newExperience
		return m, cmd

	case state.Contact:
		newContact, cmd := m.contact.Update(msg)
		m.contact = newContact
		return m, cmd
	}

	return m, nil
}

func (m appModel) View() string {
	switch m.screen {
	case state.Intro:
		return m.intro.View()
	case state.Menu:
		return m.menu.View()
	case state.Projects:
		return m.projects.View()
	case state.Skills:
		return m.skills.View()
	case state.Experience:
		return m.experience.View()
	case state.Contact:
		return m.contact.View()
	default:
		return ""
	}
}

func App() {
	p := tea.NewProgram(AppModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}