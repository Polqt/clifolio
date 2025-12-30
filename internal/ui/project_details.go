package ui

import (
	"clifolio/internal/services"
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type projectDetailsModel struct {
	project 		services.Repo
	rawMD			string
	rendered		string
	loaded 		    bool
	err				error

	width 			int
	height 			int
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
	theme := styles.NewThemeFromName("default")
	titleStyles := lipgloss.NewStyle().Foreground(theme.Primary).Bold(true).MarginBottom(1)
	metaStyle := lipgloss.NewStyle().Foreground(theme.Secondary)
	errorStyle := lipgloss.NewStyle().Foreground(theme.Error)
	loadingStyle := lipgloss.NewStyle().Foreground(theme.Accent)
	helpStyle := lipgloss.NewStyle().Foreground(theme.Help).MarginTop(1)
	boxStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(theme.Primary).Padding(1, 2)
	starStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))


	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error loading project details: %v", m.err))
	}
	
	if !m.loaded {
		loadingBox := boxStyle.Render(loadingStyle.Render(" Rendering markdown..."))
		if m.width > 0 && m.height > 0 {
			return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, loadingBox)
		}
		return "\n\n" + loadingBox
	}

	var s string

	lang := m.project.Language
	if lang == "" {
		lang = "Unknown"
	}

	langColor := styles.GetLanguageColor(lang)
	langStyle := lipgloss.NewStyle().Foreground(langColor).Bold(true)

	header := titleStyles.Render("📁 " + m.project.Name) + "\n"
	header += metaStyle.Render(m.project.Description) + "\n"
	header += langStyle.Render("● " + lang) + "  " + starStyle.Render(fmt.Sprintf("★ %d", m.project.Stars)) + "\n"
	header += metaStyle.Render("🔗 " + m.project.HTMLURL) + "\n"

	s += "\n" + header + "\n"

	if m.loaded && m.rendered != "" {
		s += m.rendered + "\n"
	}

	s += helpStyle.Render("\n↑/↓: navigate • enter: select • q: quit")
	return s
}
