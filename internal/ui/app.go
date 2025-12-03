package ui

import (
	"clifolio/internal/services"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
)

type appModel struct {
	screen state.Screen
	
	introModel tea.Model
	menuModel  tea.Model		

	intro      tea.Model
	menu       tea.Model
	projects   tea.Model
	projectDetail tea.Model
	skills     tea.Model
	experience tea.Model
	contact    tea.Model

	theme 	   string
	menuOpen   bool
}

func AppWithTheme(themeName string) tea.Model {
	m := &appModel{
		screen: state.ScreenIntro,
		theme: themeName,
	}
	m.menuModel = MenuModel()
	m.introModel = nil
	return m
}


func AppModel() appModel {
	return appModel{
		screen:     state.ScreenIntro,
		intro:      IntroModel(),
		menu:       MenuModel(),
		projects:   ProjectsModel("Polqt"),
		projectDetail: ProjectDetailsModel(services.Repo{}, ""),
		skills:     SkillsModel(),
		experience: ExperienceModel(),
		contact:    ContactModel(),
		theme: "default",
		menuOpen: false,
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
		m.screen = state.ScreenMenu
		if mm, ok := m.menu.(*menuModel); ok {
			mm.open = true
		}
		return m, nil

	}

	if pm, ok := msg.(openProjectMsg); ok {
		m.projectDetail = ProjectDetailsModel(pm.repo, pm.md)
		m.screen = state.ScreenProjectDetail
		return m, nil
	}

	switch m.screen {
	case state.ScreenIntro:
		newIntro, cmd := m.intro.Update(msg)
		m.intro = newIntro
		return m, cmd

	case state.ScreenMenu:
		newMenu, cmd := m.menu.Update(msg)
		m.menu = newMenu
		if cmd != nil {
			return m, cmd
		}
		switch msg := msg.(type) {
		case state.Screen:
			m.screen = msg
			switch msg {
			case state.ScreenProjects:
				return m, m.projects.Init()
			}
		}
		return m, nil

	case state.ScreenProjects:
		newProjects, cmd := m.projects.Update(msg)
		m.projects = newProjects
		return m, cmd

	case state.ScreenProjectDetail:
		newProjectDetail, cmd := m.projectDetail.Update(msg)
		m.projectDetail = newProjectDetail
		return m, cmd

	case state.ScreenSkills:
		newSkills, cmd := m.skills.Update(msg)
		m.skills = newSkills
		return m, cmd

	case state.ScreenExperience:
		newExperience, cmd := m.experience.Update(msg)
		m.experience = newExperience
		return m, cmd

	case state.ScreenContact:
		newContact, cmd := m.contact.Update(msg)
		m.contact = newContact
		return m, cmd
	}

	return m, nil
}

func (m appModel) View() string {
	switch m.screen {
	case state.ScreenIntro:
		return m.intro.View()
	case state.ScreenMenu:
		return m.menu.View()
	case state.ScreenProjects:
		return m.projects.View()
	case state.ScreenProjectDetail:
		return m.projectDetail.View()
	case state.ScreenSkills:
		return m.skills.View()
	case state.ScreenExperience:
		return m.experience.View()
	case state.ScreenContact:
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
