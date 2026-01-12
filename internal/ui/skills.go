package ui

import (
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Skill struct {
	Name     string
	Level    int
	Category string
	Years    int
	Icon     string
	Projects int
	Color    lipgloss.Color
}

type CategoryInfo struct {
	ID          string
	DisplayName string
	Icon        string
	Description string
}

type skillsModel struct {
	skills     []Skill
	categories []CategoryInfo
	cursor     int
	theme      styles.Theme
	width      int
	height     int
	keymap     components.Keymap
	category   string
}

func NewSkillsModel(theme styles.Theme) *skillsModel {
	categories := []CategoryInfo{
		{
			ID:          "frontend",
			DisplayName: "UI Mastery",
			Icon:        "🎨",
			Description: "Visual combat & interface arts",
		},
		{
			ID:          "backend",
			DisplayName: "Server Arts",
			Icon:        "⚙",
			Description: "Backend sorcery & API crafting",
		},
		{
			ID:          "mobile",
			DisplayName: "Mobile Tactics",
			Icon:        "📱",
			Description: "Portable realm creation",
		},
		{
			ID:          "devops",
			DisplayName: "War Engineering",
			Icon:        "🐳",
			Description: "Infrastructure & deployment tactics",
		},
		{
			ID:          "database",
			DisplayName: "Data Vaults",
			Icon:        "🗄",
			Description: "Knowledge storage mastery",
		},
		{
			ID:          "languages",
			DisplayName: "Code Tongues",
			Icon:        "💻",
			Description: "Ancient programming languages",
		},
	}

	skills := []Skill{
		// Frontend
		{Name: "React", Level: 5, Category: "frontend", Years: 3, Icon: "⚡", Projects: 20, Color: lipgloss.Color("#61DAFB")},
		{Name: "Vue", Level: 1, Category: "frontend", Years: 1, Icon: "⚡", Projects: 20, Color: lipgloss.Color("#61DAFB")},
		{Name: "Next.js", Level: 4, Category: "frontend", Years: 2, Icon: "▲", Projects: 15, Color: lipgloss.Color("#FFFFFF")},
		{Name: "TailwindCSS", Level: 5, Category: "frontend", Years: 2, Icon: "🎯", Projects: 18, Color: lipgloss.Color("#06B6D4")},
		// Backend
		{Name: "Node.js", Level: 4, Category: "backend", Years: 3, Icon: "🟩", Projects: 18, Color: lipgloss.Color("#339933")},
		{Name: "Go", Level: 2, Category: "backend", Years: 1, Icon: "🐹", Projects: 15, Color: lipgloss.Color("#00ADD8")},
		{Name: "Python", Level: 2, Category: "backend", Years: 2, Icon: "🐍", Projects: 10, Color: lipgloss.Color("#3776AB")},
		{Name: "Express.js", Level: 4, Category: "backend", Years: 3, Icon: "🚂", Projects: 16, Color: lipgloss.Color("#FFFFFF")},
		{Name: "REST APIs", Level: 5, Category: "backend", Years: 3, Icon: "🔌", Projects: 22, Color: lipgloss.Color("#00D9FF")},

		// Mobile
		{Name: "Flutter", Level: 2, Category: "mobile", Years: 1, Icon: "🎯", Projects: 8, Color: lipgloss.Color("#02569B")},
		{Name: "Dart", Level: 2, Category: "mobile", Years: 1, Icon: "💙", Projects: 8, Color: lipgloss.Color("#0175C2")},
		{Name: "React Native", Level: 2, Category: "mobile", Years: 1, Icon: "📱", Projects: 5, Color: lipgloss.Color("#61DAFB")},

		// DevOps
		{Name: "Docker", Level: 2, Category: "devops", Years: 1, Icon: "🐳", Projects: 12, Color: lipgloss.Color("#2496ED")},
		{Name: "Git", Level: 4, Category: "devops", Years: 3, Icon: "🔧", Projects: 50, Color: lipgloss.Color("#F05032")},
		{Name: "GitHub Actions", Level: 4, Category: "devops", Years: 2, Icon: "⚡", Projects: 10, Color: lipgloss.Color("#2088FF")},
		{Name: "AWS", Level: 2, Category: "devops", Years: 1, Icon: "🌐", Projects: 6, Color: lipgloss.Color("#FF9900")},

		// Database
		{Name: "PostgreSQL", Level: 4, Category: "database", Years: 3, Icon: "🐘", Projects: 15, Color: lipgloss.Color("#336791")},
		{Name: "MongoDB", Level: 4, Category: "database", Years: 1, Icon: "🍃", Projects: 12, Color: lipgloss.Color("#47A248")},
		{Name: "Redis", Level: 2, Category: "database", Years: 1, Icon: "🔴", Projects: 5, Color: lipgloss.Color("#DC382D")},
		{Name: "MySQL", Level: 4, Category: "database", Years: 3, Icon: "🐬", Projects: 14, Color: lipgloss.Color("#4479A1")},

		// Languages
		{Name: "JavaScript", Level: 5, Category: "languages", Years: 4, Icon: "🟨", Projects: 30, Color: lipgloss.Color("#F7DF1E")},
		{Name: "TypeScript", Level: 4, Category: "languages", Years: 3, Icon: "🔷", Projects: 25, Color: lipgloss.Color("#3178C6")},
		{Name: "Go", Level: 2, Category: "languages", Years: 2, Icon: "🐹", Projects: 15, Color: lipgloss.Color("#00ADD8")},
		{Name: "Python", Level: 3, Category: "languages", Years: 2, Icon: "🐍", Projects: 10, Color: lipgloss.Color("#3776AB")},
		{Name: "SQL", Level: 4, Category: "languages", Years: 3, Icon: "📊", Projects: 18, Color: lipgloss.Color("#CC2927")},
	}

	return &skillsModel{
		skills:     skills,
		categories: categories,
		theme:      theme,
		keymap:     components.DefaultKeymap(),
		category:   "frontend",
	}
}

// SkillsModel creates a new skills screen with default theme
func SkillsModel() tea.Model {
	theme := styles.NewThemeFromName("default")
	return NewSkillsModel(theme)
}

func (m *skillsModel) Init() tea.Cmd {
	return nil
}

func (m *skillsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case m.keymap.Left, "left":
			m.cycleCategoryBackward()
			m.cursor = 0
		case m.keymap.Right, "right":
			m.cycleCategoryForward()
			m.cursor = 0
		case "1":
			m.category = "frontend"
			m.cursor = 0
		case "2":
			m.category = "backend"
			m.cursor = 0
		case "3":
			m.category = "mobile"
			m.cursor = 0
		case "4":
			m.category = "devops"
			m.cursor = 0
		case "5":
			m.category = "database"
			m.cursor = 0
		case "6":
			m.category = "languages"
			m.cursor = 0
		case m.keymap.Back, "esc":
			return m, func() tea.Msg {
				return state.ScreenMenu
			}
		case m.keymap.Quit, "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *skillsModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sections []string

	// Header
	header := components.HeaderBox("WARRIOR'S ABILITIES", m.theme, m.width-4)
	sections = append(sections, header)

	// Category selector
	categorySelector := m.renderCategorySelector()
	sections = append(sections, categorySelector)

	sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

	// Current category info
	currentCat := m.getCurrentCategoryInfo()
	if currentCat != nil {
		catInfo := lipgloss.NewStyle().
			Foreground(m.theme.Secondary).
			Italic(true).
			Align(lipgloss.Center).
			Width(m.width).
			Render(fmt.Sprintf("%s %s - %s", currentCat.Icon, currentCat.DisplayName, currentCat.Description))
		sections = append(sections, catInfo)
	}

	// Stats overview
	stats := m.renderStatsOverview()
	sections = append(sections, stats)

	// Skills grid - all skills visible, no scrolling needed
	skillsGrid := m.renderSkillsCompactGrid()
	sections = append(sections, skillsGrid)

	// Key bindings
	keyBindings := []components.KeyBind{
		{Key: "←→/h/l", Desc: "Switch Abilities"},
		{Key: "1-6", Desc: "Quick Access"},
		{Key: "b/Esc", Desc: "Retreat"},
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

func (m *skillsModel) renderCategorySelector() string {
	var tabs []string

	activeStyle := lipgloss.NewStyle().
		Foreground(m.theme.Accent).
		Bold(true).
		Padding(0, 2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(m.theme.Accent)

	inactiveStyle := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Padding(0, 2).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("transparent"))

	for _, cat := range m.categories {
		label := cat.Icon + " " + cat.DisplayName
		if cat.ID == m.category {
			tabs = append(tabs, activeStyle.Render(label))
		} else {
			tabs = append(tabs, inactiveStyle.Render(label))
		}
	}

	tabsRow := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		tabsRow,
	)
}

func (m *skillsModel) renderStatsOverview() string {
	filteredSkills := m.getFilteredSkills()

	totalSkills := len(filteredSkills)
	avgLevel := 0.0
	totalProjects := 0
	maxYears := 0

	for _, skill := range filteredSkills {
		avgLevel += float64(skill.Level)
		totalProjects += skill.Projects
		if skill.Years > maxYears {
			maxYears = skill.Years
		}
	}

	if totalSkills > 0 {
		avgLevel /= float64(totalSkills)
	}

	stats := []string{
		fmt.Sprintf("Skills: %d", totalSkills),
		fmt.Sprintf("Mastery: %.1f/5", avgLevel),
		fmt.Sprintf("Battles: %d", totalProjects),
		fmt.Sprintf("Years: %d+", maxYears),
	}

	statStyle := lipgloss.NewStyle().
		Foreground(m.theme.Primary).
		Bold(true).
		Padding(0, 3)

	var statBoxes []string
	for _, stat := range stats {
		statBoxes = append(statBoxes, statStyle.Render(stat))
	}

	statsRow := lipgloss.JoinHorizontal(lipgloss.Top, statBoxes...)

	statsContainer := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(m.theme.Secondary).
		Padding(0, 1).
		Render(statsRow)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		statsContainer,
	)
}

func (m *skillsModel) renderSkillsCompactGrid() string {
	filteredSkills := m.getFilteredSkills()

	if len(filteredSkills) == 0 {
		emptyStyle := lipgloss.NewStyle().
			Foreground(m.theme.Secondary).
			Italic(true).
			Align(lipgloss.Center).
			Padding(2)
		return emptyStyle.Render("No skills in this category yet")
	}

	// Calculate optimal columns based on available width
	availableWidth := m.width - 24 // Account for margins and padding
	cols := 4
	if availableWidth < 120 {
		cols = 2
	} else if availableWidth < 160 {
		cols = 3
	}

	// Calculate card width to fill the space evenly
	margin := 4 // Total horizontal margin between cards
	cardWidth := (availableWidth - (margin * (cols - 1))) / cols

	var rows []string
	var currentRow []string

	for _, skill := range filteredSkills {
		card := m.renderSkillCard(skill, cardWidth)
		currentRow = append(currentRow, card)

		if len(currentRow) == cols {
			row := lipgloss.JoinHorizontal(lipgloss.Top, currentRow...)
			rows = append(rows, lipgloss.PlaceHorizontal(m.width-8, lipgloss.Center, row))
			currentRow = []string{}
		}
	}

	// Add remaining cards
	if len(currentRow) > 0 {
		row := lipgloss.JoinHorizontal(lipgloss.Top, currentRow...)
		rows = append(rows, lipgloss.PlaceHorizontal(m.width-8, lipgloss.Center, row))
	}

	grid := lipgloss.JoinVertical(lipgloss.Center, rows...)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		components.SectionBox("⚡ Abilities Acquired", grid, m.theme, m.width-8),
	)
}

func (m *skillsModel) renderSkillCard(skill Skill, width int) string {
	// Icon with larger size
	iconStyle := lipgloss.NewStyle().
		Foreground(skill.Color).
		Bold(true).
		Width(width).
		Align(lipgloss.Center)

	// Bigger title with more presence
	titleStyle := lipgloss.NewStyle().
		Foreground(m.theme.Primary).
		Bold(true).
		Width(width).
		Align(lipgloss.Center)

	// Simple card without borders
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		iconStyle.Render(skill.Icon),
		"", // Empty line for spacing
		titleStyle.Render(skill.Name),
	)

	cardStyle := lipgloss.NewStyle().
		Width(width).
		Height(6).
		Margin(0, 1, 1, 0).
		Padding(1)

	return cardStyle.Render(content)
}

func (m *skillsModel) getFilteredSkills() []Skill {
	var filtered []Skill
	for _, skill := range m.skills {
		if skill.Category == m.category {
			filtered = append(filtered, skill)
		}
	}
	return filtered
}

func (m *skillsModel) getCurrentCategoryInfo() *CategoryInfo {
	for _, cat := range m.categories {
		if cat.ID == m.category {
			return &cat
		}
	}
	return nil
}

func (m *skillsModel) cycleCategoryForward() {
	for i, cat := range m.categories {
		if cat.ID == m.category {
			m.category = m.categories[(i+1)%len(m.categories)].ID
			return
		}
	}
}

func (m *skillsModel) cycleCategoryBackward() {
	for i, cat := range m.categories {
		if cat.ID == m.category {
			m.category = m.categories[(i-1+len(m.categories))%len(m.categories)].ID
			return
		}
	}
}
