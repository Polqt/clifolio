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

type Skill struct {
	Name     string
	Level    int // 1-5
	Category string
	Years    int
	Icon     string
	Projects int
}

type skillsModel struct {
	skills   []Skill
	cursor   int
	theme    styles.Theme
	width    int
	height   int
	keymap   components.Keymap
	category string
}

func NewSkillsModel(theme styles.Theme) *skillsModel {
	skills := []Skill{
		{Name: "Go", Level: 4, Category: "languages", Years: 2, Icon: "🐹", Projects: 15},
		{Name: "JavaScript", Level: 5, Category: "languages", Years: 4, Icon: "🟨", Projects: 30},
		{Name: "TypeScript", Level: 4, Category: "languages", Years: 3, Icon: "🔷", Projects: 25},
		{Name: "Python", Level: 3, Category: "languages", Years: 2, Icon: "🐍", Projects: 10},
		{Name: "Flutter", Level: 4, Category: "frameworks", Years: 2, Icon: "🎯", Projects: 8},
		{Name: "React", Level: 5, Category: "frameworks", Years: 3, Icon: "⚛️", Projects: 20},
		{Name: "Node.js", Level: 4, Category: "frameworks", Years: 3, Icon: "🟩", Projects: 18},
		{Name: "Docker", Level: 4, Category: "tools", Years: 2, Icon: "🐳", Projects: 12},
		{Name: "Git", Level: 5, Category: "tools", Years: 4, Icon: "🔧", Projects: 50},
		{Name: "PostgreSQL", Level: 4, Category: "tools", Years: 3, Icon: "🐘", Projects: 15},
		{Name: "Redis", Level: 3, Category: "tools", Years: 1, Icon: "🔴", Projects: 5},
		{Name: "Problem Solving", Level: 5, Category: "soft", Years: 5, Icon: "🧩", Projects: 0},
		{Name: "Team Leadership", Level: 4, Category: "soft", Years: 3, Icon: "👥", Projects: 0},
		{Name: "Communication", Level: 5, Category: "soft", Years: 5, Icon: "💬", Projects: 0},
	}

	return &skillsModel{
		skills:   skills,
		theme:    theme,
		keymap:   components.DefaultKeymap(),
		category: "all",
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
		filteredSkills := m.getFilteredSkills()

		switch msg.String() {
		case m.keymap.Up, "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case m.keymap.Down, "j":
			if m.cursor < len(filteredSkills)-1 {
				m.cursor++
			}
		case m.keymap.Left, "h":
			m.cycleCategoryBackward()
			m.cursor = 0
		case m.keymap.Right, "l":
			m.cycleCategoryForward()
			m.cursor = 0
		case m.keymap.Back, "esc", "b":
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
	header := components.HeaderBox("SKILLS & EXPERTISE", m.theme, m.width-4)
	sections = append(sections, header)

	// Category tabs
	categoryTabs := m.renderCategoryTabs()
	sections = append(sections, categoryTabs)

	// Divider
	sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

	// Stats overview
	stats := m.renderStatsOverview()
	sections = append(sections, stats)

	// Skills grid/list
	skillsView := m.renderSkillsGrid()
	sections = append(sections, skillsView)

	// Skill details for selected
	filteredSkills := m.getFilteredSkills()
	if len(filteredSkills) > 0 && m.cursor < len(filteredSkills) {
		detail := m.renderSkillDetail(filteredSkills[m.cursor])
		sections = append(sections, detail)
	}

	// Key bindings
	keyBindings := []components.KeyBind{
		{Key: "←→/h/l", Desc: "Switch Category"},
		{Key: "↑↓/k/j", Desc: "Navigate"},
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

func (m *skillsModel) renderCategoryTabs() string {
	categories := []struct {
		ID    string
		Label string
		Icon  string
	}{
		{"all", "All Skills", "📚"},
		{"languages", "Languages", "💻"},
		{"frameworks", "Frameworks", "🔧"},
		{"tools", "Tools", "⚙️"},
		{"soft", "Soft Skills", "🎯"},
	}

	activeStyle := lipgloss.NewStyle().
		Foreground(m.theme.Accent).
		Background(lipgloss.Color("#1a1a1a")).
		Bold(true).
		Padding(0, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.theme.Accent)

	inactiveStyle := lipgloss.NewStyle().
		Foreground(m.theme.Secondary).
		Padding(0, 2)

	var tabs []string
	for _, cat := range categories {
		label := cat.Icon + " " + cat.Label
		if cat.ID == m.category {
			tabs = append(tabs, activeStyle.Render(label))
		} else {
			tabs = append(tabs, inactiveStyle.Render(label))
		}
	}

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, tabs...),
	)
}

func (m *skillsModel) renderStatsOverview() string {
	filteredSkills := m.getFilteredSkills()

	totalSkills := len(filteredSkills)
	avgLevel := 0.0
	totalProjects := 0
	totalYears := 0

	for _, skill := range filteredSkills {
		avgLevel += float64(skill.Level)
		totalProjects += skill.Projects
		totalYears += skill.Years
	}

	if totalSkills > 0 {
		avgLevel /= float64(totalSkills)
	}

	stats := []string{
		fmt.Sprintf("📊 %d Skills", totalSkills),
		fmt.Sprintf("⭐ %.1f Avg Level", avgLevel),
		fmt.Sprintf("📦 %d Projects", totalProjects),
		fmt.Sprintf("📅 %d Years Combined", totalYears),
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

func (m *skillsModel) renderSkillsGrid() string {
	filteredSkills := m.getFilteredSkills()

	items := make([]components.ListItem, len(filteredSkills))
	for i, skill := range filteredSkills {
		levelBar := m.renderLevelBar(skill.Level)

		desc := fmt.Sprintf("%s\n%d years experience", levelBar, skill.Years)
		if skill.Projects > 0 {
			desc += fmt.Sprintf(" • %d projects", skill.Projects)
		}

		items[i] = components.ListItem{
			Title:   skill.Name,
			Content: desc,
			Icon:    skill.Icon,
			Badge:   fmt.Sprintf("Lvl %d", skill.Level),
		}
	}

	return components.RenderCardList(items, m.cursor, m.theme, m.width-8)
}

func (m *skillsModel) renderLevelBar(level int) string {
	filled := strings.Repeat("█", level)
	empty := strings.Repeat("░", 5-level)

	filledStyle := lipgloss.NewStyle().Foreground(m.theme.Accent)
	emptyStyle := lipgloss.NewStyle().Foreground(m.theme.Secondary)

	return filledStyle.Render(filled) + emptyStyle.Render(empty)
}

func (m *skillsModel) renderSkillDetail(skill Skill) string {
	var details []string

	details = append(details, components.InfoPanel("Skill", skill.Name, m.theme))
	details = append(details, components.InfoPanel("Level", m.renderLevelBar(skill.Level), m.theme))
	details = append(details, components.InfoPanel("Experience", fmt.Sprintf("%d years", skill.Years), m.theme))

	if skill.Projects > 0 {
		details = append(details, components.InfoPanel("Projects", fmt.Sprintf("%d", skill.Projects), m.theme))
	}

	details = append(details, components.InfoPanel("Category", skill.Category, m.theme))

	content := strings.Join(details, "\n")

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		components.SectionBox("Selected Skill Details", content, m.theme, m.width-8),
	)
}

func (m *skillsModel) getFilteredSkills() []Skill {
	if m.category == "all" {
		return m.skills
	}

	var filtered []Skill
	for _, skill := range m.skills {
		if skill.Category == m.category {
			filtered = append(filtered, skill)
		}
	}
	return filtered
}

func (m *skillsModel) cycleCategoryForward() {
	categories := []string{"all", "languages", "frameworks", "tools", "soft"}
	for i, cat := range categories {
		if cat == m.category {
			m.category = categories[(i+1)%len(categories)]
			return
		}
	}
}

func (m *skillsModel) cycleCategoryBackward() {
	categories := []string{"all", "languages", "frameworks", "tools", "soft"}
	for i, cat := range categories {
		if cat == m.category {
			m.category = categories[(i-1+len(categories))%len(categories)]
			return
		}
	}
}
