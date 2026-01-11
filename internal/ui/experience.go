package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Experience struct {
	Type         string
	Title        string
	Organization string
	Location     string
	StartDate    string
	EndDate      string
	Description  []string
	Skills       []string
	Icon         string
}

type experienceModel struct {
	experiences []Experience
	cursor      int
	theme       styles.Theme
	width       int
	height      int
	keymap      components.Keymap
	viewType    string
}

func NewExperienceModel(theme styles.Theme) *experienceModel {
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
			Organization: "University of St. La Salle - Bacolod",
			Location:     "Philippines",
			StartDate:    "August 2022",
			EndDate:      "April 2026",
			Description: []string{
				"Focused on game development, data structures and algorithms, and artificial intelligence",
				"Dean's List recipient for academic excellence",
			},
			Skills: []string{"Algorithms", "Data Structures", "Game Development", "AI", "ML", "Data Science"},
			Icon:   "🎓",
		},
	}

	return &experienceModel{
		experiences: experiences,
		theme:       theme,
		keymap:      components.DefaultKeymap(),
		viewType:    "timeline",
	}
}

func ExperienceModel() tea.Model {
	theme := styles.NewThemeFromName("default")
	return NewExperienceModel(theme)
}

func (m *experienceModel) Init() tea.Cmd {
	return nil
}

func (m *experienceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case m.keymap.Up, "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case m.keymap.Down, "down":
			if m.cursor < len(m.experiences)-1 {
				m.cursor++
			}
		case m.keymap.Confirm:
			if m.viewType == "timeline" {
				m.viewType = "detailed"
			} else {
				m.viewType = "timeline"
			}
		case m.keymap.Back, "esc", "b":
			// If in detailed view, go back to timeline first
			if m.viewType == "detailed" {
				m.viewType = "timeline"
				return m, nil
			}
			// Otherwise go back to menu
			m.cursor = 0
			return m, func() tea.Msg {
				return state.ScreenMenu
			}
		case m.keymap.Quit, "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *experienceModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sections []string

	// Header
	header := components.HeaderBox("COMBAT HISTORY & TRAINING", m.theme, m.width-4)
	sections = append(sections, header)

	// Stats - only show in timeline view to save space
	if m.viewType == "timeline" {
		stats := m.renderStats()
		sections = append(sections, stats)
		sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))
	}

	// Experience content
	if m.viewType == "timeline" {
		sections = append(sections, m.renderTimeline())
	} else {
		sections = append(sections, m.renderDetailed())
	}

	// Key bindings
	keyBindings := []components.KeyBind{
		{Key: "↑↓/k/j", Desc: "Navigate"},
		{Key: "enter", Desc: "Toggle View"},
		{Key: "b/esc", Desc: "Retreat"},
		{Key: "q", Desc: "Exit Realm"},
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

func (m *experienceModel) renderStats() string {
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

func (m *experienceModel) renderTimeline() string {
	var cards []string

	for i, exp := range m.experiences {
		isSelected := i == m.cursor

		headerStyle := lipgloss.NewStyle().
			Foreground(m.theme.Accent).
			Bold(true).
			PaddingBottom(1)

		cardHeader := headerStyle.Render(exp.Icon + "  " + exp.Title)

		orgStyle := lipgloss.NewStyle().
			Foreground(m.theme.Primary).
			Bold(true)

		locationStyle := lipgloss.NewStyle().
			Foreground(m.theme.Secondary).
			Italic(true)

		dateStyle := lipgloss.NewStyle().
			Foreground(m.theme.Secondary).
			Italic(true)

		orgLine := orgStyle.Render(exp.Organization)
		locationLine := locationStyle.Render(exp.Location)
		dateRange := dateStyle.Render(exp.StartDate + " → " + exp.EndDate)

		var descPreview string
		if len(exp.Description) > 0 {
			descStyle := lipgloss.NewStyle().
				Foreground(m.theme.Secondary).
				PaddingTop(1).
				Italic(true)

			desc := exp.Description[0]
			if len(desc) > 60 {
				desc = desc[:57] + "..."
			}
			descPreview = descStyle.Render("" + desc)
		}

		// Skills badges
		var skillsLine string
		if len(exp.Skills) > 0 {
			skillStyle := lipgloss.NewStyle().
				Foreground(m.theme.Background).
				Background(m.theme.Accent).
				Padding(0, 1).
				MarginRight(1).
				Bold(true)

			var skillBadges []string
			for _, skill := range exp.Skills {
				skillBadges = append(skillBadges, skillStyle.Render(skill))
			}
			skillsLine = "\n" + strings.Join(skillBadges, " ")
		}

		// Build card content
		cardContent := lipgloss.JoinVertical(
			lipgloss.Left,
			cardHeader,
			orgLine,
			locationLine,
			dateRange,
			descPreview,
			skillsLine,
		)

		// Apply card styling
		var card string
		if isSelected {
			cardStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(m.theme.Accent).
				Padding(1, 2).
				MarginBottom(1).
				Width(m.width - 20).
				BorderStyle(lipgloss.ThickBorder())

			card = cardStyle.Render(cardContent)
		} else {
			cardStyle := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(m.theme.Secondary).
				Padding(1, 2).
				MarginBottom(1).
				Width(m.width - 20)

			card = cardStyle.Render(cardContent)
		}

		cards = append(cards, card)
	}

	allCards := lipgloss.JoinVertical(lipgloss.Left, cards...)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		allCards,
	)
}

func (m *experienceModel) renderDetailed() string {
	if m.cursor >= len(m.experiences) {
		return ""
	}

	exp := m.experiences[m.cursor]

	var sections []string

	// Title section
	titleStyle := lipgloss.NewStyle().
		Foreground(m.theme.Accent).
		Bold(true).
		Underline(true)

	titleSection := titleStyle.Render(exp.Icon + "  " + exp.Title)
	sections = append(sections, titleSection)
	sections = append(sections, "")

	// Info in clean aligned format
	labelStyle := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Width(10).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(m.theme.Primary)

	orgLine := lipgloss.JoinHorizontal(lipgloss.Left,
		labelStyle.Render("Company:"),
		"  ",
		valueStyle.Render(exp.Organization))
	locationLine := lipgloss.JoinHorizontal(lipgloss.Left,
		labelStyle.Render("Location:"),
		"  ",
		valueStyle.Render(exp.Location))
	dateLine := lipgloss.JoinHorizontal(lipgloss.Left,
		labelStyle.Render("Period:"),
		"  ",
		valueStyle.Render(exp.StartDate+" → "+exp.EndDate))

	sections = append(sections, orgLine)
	sections = append(sections, locationLine)
	sections = append(sections, dateLine)
	sections = append(sections, "")

	// Quest Log (Description) with better styling
	if len(exp.Description) > 0 {
		divider := lipgloss.NewStyle().
			Foreground(m.theme.Accent).
			Bold(true).
			Render("━━━ QUEST LOG ━━━")

		sections = append(sections, divider)

		for _, desc := range exp.Description {
			// Add spacing to align with the values above (10 char label + 2 spaces)
			indent := strings.Repeat(" ", 12)
			descText := lipgloss.NewStyle().Foreground(m.theme.Primary).Render(desc)
			sections = append(sections, indent+descText)
		}
		sections = append(sections, "")
	}

	// Skills Arsenal with enhanced badges
	if len(exp.Skills) > 0 {
		divider := lipgloss.NewStyle().
			Foreground(m.theme.Accent).
			Bold(true).
			Render("━━━ SKILLS ━━━")

		sections = append(sections, divider)

		skillStyle := lipgloss.NewStyle().
			Foreground(m.theme.Background).
			Background(m.theme.Accent).
			Padding(0, 1).
			MarginRight(1).
			Bold(true)

		var skillBadges []string
		for _, skill := range exp.Skills {
			skillBadges = append(skillBadges, skillStyle.Render(skill))
		}
		sections = append(sections, strings.Join(skillBadges, " "))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		components.SectionBox("", content, m.theme, m.width-8),
	)
}
