package ui

import (
    "clifolio/internal/styles"
    "clifolio/internal/ui/components"
    "fmt"
    "strings"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type Experience struct {
    Type        string // "work", "education", "certification"
    Title       string
    Organization string
    Location    string
    StartDate   string
    EndDate     string
    Description []string
    Skills      []string
    Icon        string
}

type ExperienceModel struct {
    experiences []Experience
    cursor      int
    theme       styles.Theme
    width       int
    height      int
    keymap      components.Keymap
    viewType    string // "timeline", "detailed"
}

func NewExperienceModel(theme styles.Theme) ExperienceModel {
    experiences := []Experience{
        {
            Type:         "work",
            Title:        "Part Time Mobile Developer",
            Organization: "K92 Paints",
            Location:     "Philippines",
            StartDate:    "December 2024",
            EndDate:      "April 2025",
            Description: []string{
                "Developed and maintained a mobile application using Flutter",
                "Implemented features for paint color selection and visualization",
                "Collaborated with design team to create intuitive user interfaces",
                "Integrated backend APIs for real-time inventory management",
            },
            Skills: []string{"Flutter", "Dart", "Mobile Development", "API Integration"},
            Icon:   "💼",
        },
        {
            Type:         "education",
            Title:        "Bachelor of Science in Computer Science",
            Organization: "Your University",
            Location:     "Philippines",
            StartDate:    "2020",
            EndDate:      "2024",
            Description: []string{
                "Focused on software engineering and full-stack development",
                "Dean's List recipient for academic excellence",
                "Led student tech community initiatives",
            },
            Skills: []string{"Algorithms", "Data Structures", "Software Engineering"},
            Icon:   "🎓",
        },
        {
            Type:         "certification",
            Title:        "AWS Certified Developer",
            Organization: "Amazon Web Services",
            Location:     "Online",
            StartDate:    "2024",
            EndDate:      "2024",
            Description: []string{
                "Certified in cloud development and deployment",
                "Expertise in serverless architectures and microservices",
            },
            Skills: []string{"AWS", "Cloud Computing", "DevOps"},
            Icon:   "📜",
        },
    }

    return ExperienceModel{
        experiences: experiences,
        theme:       theme,
        keymap:      components.DefaultKeymap(),
        viewType:    "timeline",
    }
}

func (m ExperienceModel) Init() tea.Cmd {
    return nil
}

func (m ExperienceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        return m, nil

    case tea.KeyMsg:
        switch msg.String() {
        case m.keymap.Up, "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case m.keymap.Down, "j":
            if m.cursor < len(m.experiences)-1 {
                m.cursor++
            }
        case "v":
            if m.viewType == "timeline" {
                m.viewType = "detailed"
            } else {
                m.viewType = "timeline"
            }
        case m.keymap.Back, "esc":
            return m, func() tea.Msg {
                return NavigateMsg{Screen: "menu"}
            }
        case m.keymap.Quit, "ctrl+c":
            return m, tea.Quit
        }
    }

    return m, nil
}

func (m ExperienceModel) View() string {
    if m.width == 0 {
        return "Loading..."
    }

    var sections []string

    // Header
    header := components.HeaderBox("EXPERIENCE & EDUCATION", m.theme, m.width-4)
    sections = append(sections, header)

    // Stats
    stats := m.renderStats()
    sections = append(sections, stats)

    // Divider
    sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

    // Experience content
    if m.viewType == "timeline" {
        sections = append(sections, m.renderTimeline())
    } else {
        sections = append(sections, m.renderDetailed())
    }

    // Key bindings
    keyBindings := []components.KeyBind{
        {Key: "↑↓/k/j", Desc: "Navigate"},
        {Key: "v", Desc: "Toggle View"},
        {Key: "b/Esc", Desc: "Back"},
        {Key: "q", Desc: "Quit"},
    }
    footer := components.RenderKeyBindings(keyBindings, m.theme, m.width)
    sections = append(sections, footer)

    content := lipgloss.JoinVertical(lipgloss.Left, sections...)

    return lipgloss.Place(
        m.width,
        m.height,
        lipgloss.Center,
        lipgloss.Center,
        content,
    )
}

func (m ExperienceModel) renderStats() string {
    workCount := 0
    eduCount := 0
    certCount := 0

    for _, exp := range m.experiences {
        switch exp.Type {
        case "work":
            workCount++
        case "education":
            eduCount++
        case "certification":
            certCount++
        }
    }

    stats := []string{
        fmt.Sprintf("💼 %d Work", workCount),
        fmt.Sprintf("🎓 %d Education", eduCount),
        fmt.Sprintf("📜 %d Certifications", certCount),
        fmt.Sprintf("👁️  %s View", m.viewType),
    }

    statStyle := lipgloss.NewStyle().
        Foreground(m.theme.Secondary).
        Background(lipgloss.Color("#1a1a1a")).
        Padding(0, 2).
        Margin(1, 0)

    var statBoxes []string
    for _, stat := range stats {
        statBoxes = append(statBoxes, statStyle.Render(stat))
    }

    return lipgloss.PlaceHorizontal(
        m.width,
        lipgloss.Center,
        lipgloss.JoinHorizontal(lipgloss.Top, statBoxes...),
    )
}

func (m ExperienceModel) renderTimeline() string {
    var timeline strings.Builder

    connectorStyle := lipgloss.NewStyle().Foreground(m.theme.Secondary)
    dotStyle := lipgloss.NewStyle().Foreground(m.theme.Accent).Bold(true)
    selectedDotStyle := lipgloss.NewStyle().Foreground(m.theme.Accent).Bold(true)

    for i, exp := range m.experiences {
        isSelected := i == m.cursor

        // Timeline connector
        var dot string
        if isSelected {
            dot = selectedDotStyle.Render("●")
        } else {
            dot = dotStyle.Render("○")
        }

        // Date range
        dateStyle := lipgloss.NewStyle().
            Foreground(m.theme.Secondary).
            Width(20)
        dateRange := fmt.Sprintf("%s - %s", exp.StartDate, exp.EndDate)

        // Content
        titleStyle := lipgloss.NewStyle().
            Foreground(m.theme.Primary).
            Bold(true)

        orgStyle := lipgloss.NewStyle().
            Foreground(m.theme.Accent)

        locationStyle := lipgloss.NewStyle().
            Foreground(m.theme.Secondary).
            Italic(true)

        content := fmt.Sprintf("%s %s\n%s %s\n%s %s",
            exp.Icon,
            titleStyle.Render(exp.Title),
            "at",
            orgStyle.Render(exp.Organization),
            "📍",
            locationStyle.Render(exp.Location),
        )

        // Skills
        if len(exp.Skills) > 0 {
            skillStyle := lipgloss.NewStyle().
                Foreground(m.theme.Accent).
                Background(lipgloss.Color("#1a1a1a")).
                Padding(0, 1).
                MarginRight(1)

            var skillBadges []string
            for _, skill := range exp.Skills {
                skillBadges = append(skillBadges, skillStyle.Render(skill))
            }
            content += "\n" + strings.Join(skillBadges, " ")
        }

        // Build timeline entry
        var entry string
        if isSelected {
            entry = components.AccentBorder(m.theme).Render(
                lipgloss.JoinHorizontal(
                    lipgloss.Top,
                    dateStyle.Render(dateRange),
                    " "+dot+" ",
                    content,
                ),
            )
        } else {
            entry = components.SubtleBorder(m.theme).Render(
                lipgloss.JoinHorizontal(
                    lipgloss.Top,
                    dateStyle.Render(dateRange),
                    " "+dot+" ",
                    content,
                ),
            )
        }

        timeline.WriteString(entry + "\n")

        // Connector line
        if i < len(m.experiences)-1 {
            connector := connectorStyle.Render("     │")
            timeline.WriteString(connector + "\n")
        }
    }

    return lipgloss.PlaceHorizontal(
        m.width,
        lipgloss.Center,
        timeline.String(),
    )
}

func (m ExperienceModel) renderDetailed() string {
    if m.cursor >= len(m.experiences) {
        return ""
    }

    exp := m.experiences[m.cursor]

    var details strings.Builder

    // Title section
    titleStyle := lipgloss.NewStyle().
        Foreground(m.theme.Accent).
        Bold(true)
    
    details.WriteString(titleStyle.Render(exp.Icon+" "+exp.Title) + "\n\n")

    // Info panels
    details.WriteString(components.InfoPanel("Organization", exp.Organization, m.theme) + "\n")
    details.WriteString(components.InfoPanel("Location", exp.Location, m.theme) + "\n")
    details.WriteString(components.InfoPanel("Period", exp.StartDate+" - "+exp.EndDate, m.theme) + "\n")
    details.WriteString(components.InfoPanel("Type", strings.Title(exp.Type), m.theme) + "\n\n")

    // Description
    if len(exp.Description) > 0 {
        descStyle := lipgloss.NewStyle().
            Foreground(m.theme.Primary).
            Bold(true)
        details.WriteString(descStyle.Render("Description:") + "\n")

        for _, desc := range exp.Description {
            bullet := lipgloss.NewStyle().
                Foreground(m.theme.Accent).
                Render("  • ")
            details.WriteString(bullet + desc + "\n")
        }
        details.WriteString("\n")
    }

    // Skills
    if len(exp.Skills) > 0 {
        skillStyle := lipgloss.NewStyle().
            Foreground(m.theme.Accent).
            Background(lipgloss.Color("#1a1a1a")).
            Padding(0, 1).
            MarginRight(1)

        var skillBadges []string
        for _, skill := range exp.Skills {
            skillBadges = append(skillBadges, skillStyle.Render(skill))
        }
        
        skillsTitle := lipgloss.NewStyle().
            Foreground(m.theme.Primary).
            Bold(true).
            Render("Skills:")
        
        details.WriteString(skillsTitle + "\n")
        details.WriteString(strings.Join(skillBadges, " ") + "\n")
    }

    return lipgloss.PlaceHorizontal(
        m.width,
        lipgloss.Center,
        components.SectionBox("Detailed View", details.String(), m.theme, m.width-8),
    )
}