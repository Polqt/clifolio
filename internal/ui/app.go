package ui

import (
    "clifolio/internal/services"
    "clifolio/internal/styles"
    "clifolio/internal/ui/state"

    tea "github.com/charmbracelet/bubbletea"
)

type appModel struct {
    screen state.Screen

    intro         tea.Model
    menu          tea.Model
    projects      tea.Model
    projectDetail tea.Model
    skills        tea.Model
    experience    tea.Model
    contact       tea.Model
    themePicker   tea.Model
    stats         tea.Model
    matrix        tea.Model

    theme    string
    menuOpen bool
    width    int
    height   int
}

func AppWithTheme(themeName string) tea.Model {
    m := &appModel{
        screen: state.ScreenIntro,
        theme:  themeName,
    }
    m.menu = MenuModel()
    m.intro = nil
    return m
}

func AppModel() appModel {
    return appModel{
        screen:        state.ScreenIntro,
        intro:         IntroModel(),
        menu:          MenuModel(),
        projects:      ProjectsModel("Polqt"),
        projectDetail: ProjectDetailsModel(services.Repo{}, ""),
        skills:        SkillsModel(),
        experience:    ExperienceModel(),
        contact:       ContactModel(),
        themePicker:   ThemePickerModel(),
        stats:         StatsModel("Polqt"),
        matrix:        MatrixModel(),
        theme:         "default",
        menuOpen:      false,
    }
}

func (m appModel) Init() tea.Cmd {
    return m.intro.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    // Handle WindowSizeMsg FIRST - forward to all models
    if wsm, ok := msg.(tea.WindowSizeMsg); ok {
        m.width = wsm.Width
        m.height = wsm.Height
        
        // Forward to all initialized models
        if m.intro != nil {
            m.intro, _ = m.intro.Update(msg)
        }
        if m.menu != nil {
            m.menu, _ = m.menu.Update(msg)
        }
        if m.projects != nil {
            m.projects, _ = m.projects.Update(msg)
        }
        if m.skills != nil {
            m.skills, _ = m.skills.Update(msg)
        }
        if m.experience != nil {
            m.experience, _ = m.experience.Update(msg)
        }
        if m.contact != nil {
            m.contact, _ = m.contact.Update(msg)
        }
        if m.themePicker != nil {
            m.themePicker, _ = m.themePicker.Update(msg)
        }
        if m.stats != nil {
            m.stats, _ = m.stats.Update(msg)
        }
        if m.matrix != nil {
            m.matrix, _ = m.matrix.Update(msg)
        }
        if m.projectDetail != nil {
            m.projectDetail, _ = m.projectDetail.Update(msg)
        }
        
        return m, nil
    }

    // Global exit
    if key, ok := msg.(tea.KeyMsg); ok {
        if key.String() == "ctrl+c" {
            return m, tea.Quit
        }

        // Global menu toggle (except during intro)
        if key.String() == "/" && m.screen != state.ScreenIntro {
            m.screen = state.ScreenMenu
            return m, nil
        }
    }

    // Handle intro -> menu transition
    if _, ok := msg.(goToMenuMsg); ok {
        m.screen = state.ScreenMenu
        return m, nil
    }

    // Handle back to projects
    if _, ok := msg.(backToProjectsMsg); ok {
        m.screen = state.ScreenProjects
        return m, nil
    }

    // Handle theme change
    if tc, ok := msg.(ThemeChangeMsg); ok {
        m.theme = tc.ThemeName
        _ = styles.NewThemeFromName(m.theme)
        m.screen = state.ScreenMenu
        return m, nil
    }

    // Handle project detail opening
    if pm, ok := msg.(openProjectMsg); ok {
        m.projectDetail = ProjectDetailsModel(pm.repo, pm.md)
        m.screen = state.ScreenProjectDetail
        return m, m.projectDetail.Init()
    }

    // Handle screen navigation
    if screen, ok := msg.(state.Screen); ok {
        m.screen = screen
        
        // Initialize the target screen if needed
        switch screen {
        case state.ScreenProjects:
            if m.projects == nil {
                m.projects = ProjectsModel("Polqt")
            }
            return m, m.projects.Init()
        case state.ScreenSkills:
            if m.skills == nil {
                m.skills = SkillsModel()
            }
            return m, m.skills.Init()
        case state.ScreenExperience:
            if m.experience == nil {
                m.experience = ExperienceModel()
            }
            return m, m.experience.Init()
        case state.ScreenContact:
            if m.contact == nil {
                m.contact = ContactModel()
            }
            return m, m.contact.Init()
        case state.ScreenStats:
            if m.stats == nil {
                m.stats = StatsModel("Polqt")
            }
            return m, m.stats.Init()
        case state.ScreenTheme:
            if m.themePicker == nil {
                m.themePicker = ThemePickerModel()
            }
            return m, m.themePicker.Init()
        case state.ScreenMatrix:
            if m.matrix == nil {
                m.matrix = MatrixModel()
            }
            return m, m.matrix.Init()
        case state.ScreenMenu:
            return m, nil
        }
        return m, nil
    }

    // Route to current screen
    switch m.screen {
    case state.ScreenIntro:
        m.intro, cmd = m.intro.Update(msg)
        return m, cmd

    case state.ScreenMenu:
        m.menu, cmd = m.menu.Update(msg)
        return m, cmd

    case state.ScreenProjects:
        m.projects, cmd = m.projects.Update(msg)
        return m, cmd

    case state.ScreenProjectDetail:
        m.projectDetail, cmd = m.projectDetail.Update(msg)
        return m, cmd

    case state.ScreenSkills:
        m.skills, cmd = m.skills.Update(msg)
        return m, cmd

    case state.ScreenExperience:
        m.experience, cmd = m.experience.Update(msg)
        return m, cmd

    case state.ScreenContact:
        m.contact, cmd = m.contact.Update(msg)
        return m, cmd

    case state.ScreenTheme:
        m.themePicker, cmd = m.themePicker.Update(msg)
        return m, cmd

    case state.ScreenStats:
        m.stats, cmd = m.stats.Update(msg)
        return m, cmd

    case state.ScreenMatrix:
        m.matrix, cmd = m.matrix.Update(msg)
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
    case state.ScreenTheme:
        return m.themePicker.View()
    case state.ScreenStats:
        return m.stats.View()
    case state.ScreenMatrix:
        return m.matrix.View()
    default:
        return "Unknown Screen"
    }
}

func App() {
    p := tea.NewProgram(AppModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        panic(err)
    }
}