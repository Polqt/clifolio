package ui

import (
	"context"
	"fmt"
	"time"

	"clifolio/internal/services"
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"
	"clifolio/internal/ui/state"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type projectsModel struct {
	username string
	projects []services.Repo
	cursor   int
	loading  bool
	err      error

	spin components.SpinnerComponent

	offset   int
	pageSize int

	width  int
	height int
}

type projectsLoadedMsg struct {
	projects []services.Repo
}

type projectsErrMsg struct {
	err error
}

type openProjectMsg struct {
	repo services.Repo
	md   string
	err  error
}

func ProjectsModel(username string) *projectsModel {
	return &projectsModel{
		username: username,
		loading:  true,
		spin:     components.NewSpinner(),
		cursor:   0,
		offset:   0,
		pageSize: 10,
	}
}

func (m *projectsModel) ensureCursorInWindow() {
	if m.cursor < m.offset {
		m.offset = m.cursor
	} else if m.cursor >= m.offset+m.pageSize {
		m.offset = m.cursor - m.pageSize + 1
	}

	if m.offset < 0 {
		m.offset = 0
	}
	if m.pageSize < 1 {
		m.pageSize = 3
	}
	if m.offset > max(0, len(m.projects)-m.pageSize) {
		m.offset = max(0, len(m.projects)-m.pageSize)
	}
}

func fetchReposCmd(username string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		repos, err := services.FetchRepos(ctx, username)
		if err != nil {
			fmt.Println("Fetch error:", err)
			return projectsErrMsg{err}
		}
		fmt.Println("Fetched repos:", len(repos))
		return projectsLoadedMsg{projects: repos}
	}
}

func (m *projectsModel) Init() tea.Cmd {
	return tea.Batch(m.spin.Init(), fetchReposCmd(m.username))
}

func fetchRepoReadmeCmd(owner string, r services.Repo) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		md, err := services.FetchRepoReadme(ctx, owner, r.Name)
		if err != nil {
			return openProjectMsg{repo: r, md: "", err: err}
		}
		return openProjectMsg{repo: r, md: md, err: nil}
	}
}

func (m *projectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	km := components.DefaultKeymap()

	var cmds []tea.Cmd

	newSpin, spinCmd := m.spin.Update(msg)
	m.spin = newSpin
	cmds = append(cmds, spinCmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		h := msg.Height
		desired := (h - 8) / 3
		if desired < 3 {
			desired = 3
		}
		m.pageSize = desired
		m.ensureCursorInWindow()

	case projectsLoadedMsg:
		m.projects = msg.projects
		m.loading = false
		if m.cursor >= len(m.projects) {
			m.cursor = 0
		}
		if m.offset > len(m.projects) {
			m.offset = 0
		}
		return m, nil

	case projectsErrMsg:
		m.err = msg.err
		m.loading = false
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case km.Quit, "ctrl+c":
			return m, tea.Quit
		case km.Back, "esc":
			return m, func() tea.Msg { return state.ScreenMenu }
		case "up", "k":
			if !m.loading && m.cursor > 0 {
				m.cursor--
				if m.cursor < m.offset {
					m.offset = m.cursor
				}
			}
		case "down", "j":
			if !m.loading && m.cursor < len(m.projects)-1 {
				m.cursor++
				if m.cursor >= m.offset+m.pageSize {
					m.offset = m.cursor - m.pageSize + 1
				}
			}
		case "enter":
			if !m.loading && len(m.projects) > 0 {
				repo := m.projects[m.cursor]
				return m, fetchRepoReadmeCmd(m.username, repo)
			}
			if !m.loading {
				m.offset += m.pageSize
				if m.offset > len(m.projects)-1 {
					m.offset = max(0, len(m.projects)-m.pageSize)
				}
				if m.cursor < m.offset {
					m.cursor = m.offset
				} else if m.cursor >= m.offset+m.pageSize {
					m.cursor = min(len(m.projects)-1, m.offset+m.pageSize-1)
				}
			}
		case "pgup":
			if !m.loading {
				m.offset -= m.pageSize
				if m.offset < 0 {
					m.offset = 0
				}
				if m.cursor < m.offset {
					m.cursor = m.offset
				}
			}
		case "home":
			if !m.loading {
				m.offset = 0
				m.cursor = 0
			}
		case "end":
			if !m.loading {
				if len(m.projects) > m.pageSize {
					m.offset = len(m.projects) - m.pageSize
				} else {
					m.offset = 0
				}
				m.cursor = len(m.projects) - 1
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *projectsModel) View() string {
	theme := styles.NewThemeFromName("default")
	titleStyles := lipgloss.NewStyle().Foreground(theme.Primary).Bold(true).MarginBottom(1)
	subtitleStyle := lipgloss.NewStyle().Foreground(theme.Secondary)
	selectedCardStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(theme.Accent).Padding(0, 1).MarginBottom(1)
	normalCardStyle := lipgloss.NewStyle().Border(lipgloss.HiddenBorder()).Padding(0, 1).MarginBottom(1)
	errorStyle := lipgloss.NewStyle().Foreground(theme.Error)
	helpStyle := lipgloss.NewStyle().Foreground(theme.Help).MarginTop(1)
	starStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))

	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("\n\n Error: %s", m.err))
	}

	if m.loading {
		loadingBox := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(theme.Primary).Padding(2, 4).Render(fmt.Sprintf("%s Loading GitHub repos...", m.spin.View()))

		if m.width > 0 && m.height > 0 {
			return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, loadingBox)
		}
		return "\n\n" + loadingBox
	}

	if len(m.projects) == 0 {
		return "\n\n No repositories found."
	}

	start := m.offset
	if start < 0 {
		start = 0
	}
	end := start + m.pageSize
	if end > len(m.projects) {
		end = len(m.projects)
	}

	totalPages := (len(m.projects) + m.pageSize - 1) / m.pageSize
	currentPage := (start / m.pageSize) + 1

	s := titleStyles.Render(fmt.Sprintf("📁 Projects of %s", m.username)) + "\n"
	s += subtitleStyle.Render(fmt.Sprintf("Showing %d-%d of %d • Page %d/%d",
		start+1, end, len(m.projects), currentPage, totalPages)) + "\n\n"

	for i := start; i < end; i++ {
		repo := m.projects[i]

		lang := repo.Language
		if lang == "" {
			lang = "Unknown"
		}
		langColor := styles.GetLanguageColor(lang)
		langStyle := lipgloss.NewStyle().Foreground(langColor).Bold(true)
		langIndicator := langStyle.Render("●") + " " + langStyle.Render(lang)

		desc := repo.Description
		if desc == "" {
			desc = "No description provided"
		}

		if len(desc) > 50 {
			desc = desc[:47] + "..."
		}

		stars := starStyle.Render(fmt.Sprintf("★ %d", repo.Stars))

		cardContent := fmt.Sprintf("%s %s\n%s",
			titleStyles.Render(repo.Name),
			stars,
			subtitleStyle.Render(desc)+"\n"+langIndicator,
		)

		if m.cursor == i {
			s += selectedCardStyle.Render("▸ "+cardContent) + "\n"
		} else {
			s += normalCardStyle.Render("  "+cardContent) + "\n"
		}
	}

	s += helpStyle.Render("\n↑/↓: navigate • enter: select • q: quit")

	return s
}
