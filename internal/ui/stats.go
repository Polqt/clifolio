package ui

import (
	"clifolio/internal/services"
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"
	"context"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type statsModel struct {
	stats 		*services.GitHubStats
	loading 	bool
	err			error
	spin 		components.SpinnerComponent
	width		int
	height		int
	username	string
}

type statsLoadedMsg struct {
	stats *services.GitHubStats
}

type statsErrorMsg struct {
	err error
}

type statsTickMsg struct{}

func StatsModel(username string) *statsModel {
	return &statsModel{
		loading: true,
		spin: components.NewSpinner(),
		username: username,
	}
}

func (m *statsModel) Init() tea.Cmd {
	return tea.Batch(
		m.spin.Init(),
		fetchStatsCmd(m.username),
		tickStats(),
	)
}

func fetchStatsCmd(username string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		stats, err := services.FetchGitHubStats(ctx, username)
		if err != nil {
			return statsErrorMsg{err}
		}
		return statsLoadedMsg{stats}
	}
}

func tickStats() tea.Cmd {
	return tea.Tick(30*time.Second, func(t time.Time) tea.Msg {
		return statsTickMsg{}
	})
}

func (m *statsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()

	var cmds []tea.Cmd


	newSpin, spinCmd := m.spin.Update(msg)
	m.spin = newSpin
	cmds = append(cmds, spinCmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		
	case statsLoadedMsg:
		m.stats = msg.stats
		m.loading = false

	case statsErrorMsg:
		m.err = msg.err
		m.loading = false

	case statsTickMsg:
		return m, tea.Batch(
			fetchStatsCmd(m.username),
			tickStats(),
		)
	case tea.KeyMsg:
		switch msg.String() {
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		case km.Back, "esc":
			return m, func() tea.Msg { return state.ScreenMenu }
		case "r":
			m.loading = true
			return m, fetchStatsCmd(m.username)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *statsModel) View() string {
	theme := styles.NewThemeFromName("default")

	titleStyle := lipgloss.NewStyle().
        Foreground(theme.Primary).
        Bold(true).
        MarginBottom(2)
    
    statBoxStyle := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(theme.Accent).
        Padding(1, 2).
        Margin(0, 1)
    
    labelStyle := lipgloss.NewStyle().
        Foreground(theme.Secondary).
        Bold(true)
    
    valueStyle := lipgloss.NewStyle().
        Foreground(theme.Primary).
        Bold(true).
        Align(lipgloss.Center)

	helpStyle := lipgloss.NewStyle().
		Foreground(theme.Help).
		MarginTop(2)

	if m.loading {
		return fmt.Sprintf("\n\n%s Fetching Github stats...\n", m.spin.View())
	}

	if m.err != nil {
		return fmt.Sprintf("\n\nError: %v\n\nPress ESC to go back\n", m.err)
	}

	if m.stats == nil {
		return "\n\nNo stats available.\n"
	}

	title := titleStyle.Render(fmt.Sprintf("GitHub stats for @%s", m.username))

	reposStat := statBoxStyle.Render(fmt.Sprintf("%s\n%s",
		labelStyle.Render("Repositories"),
		valueStyle.Render(fmt.Sprintf("%d", m.stats.TotalRepos)),
	))

	starsStat := statBoxStyle.Render(fmt.Sprint("%s\n%s",
		labelStyle.Render("Stars"),
		valueStyle.Render(fmt.Sprintf("⭐ %d", m.stats.TotalStars)),
	))

	followersStat := statBoxStyle.Render(fmt.Sprintf("%s\n%s",
		labelStyle.Render("Followers"),
		valueStyle.Render(fmt.Sprintf("👥 %d", m.stats.Followers)),
	))
	
	gistsStat := statBoxStyle.Render(fmt.Sprintf("%s\n%s",
        labelStyle.Render("Public Gists"),
        valueStyle.Render(fmt.Sprintf("📝 %d", m.stats.PublicGists)),
    ))

	row1 := lipgloss.JoinHorizontal(lipgloss.Top, reposStat, starsStat)
	row2 := lipgloss.JoinHorizontal(lipgloss.Top, followersStat, gistsStat)
	statsGrid := lipgloss.JoinVertical(lipgloss.Left, row1, row2)

	help := helpStyle.Render("Press 'r' to refresh • ESC to go back • q to quit")

	lastUpdate := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Italic(true).
		Render(fmt.Sprintf("Last updated: %s", m.stats.UpdatedAt.Format("15:04:05")))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		"\n",
		title,
		statsGrid,
		lastUpdate,
		help,
	)
}
