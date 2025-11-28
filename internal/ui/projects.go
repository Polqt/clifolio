package ui

import (
	"context"
	"fmt"
	"time"

	"clifolio/internal/services"
	"clifolio/internal/ui/components"

	tea "github.com/charmbracelet/bubbletea"
)

type projectsModel struct {
	username 	string
	projects 	[]services.Repo
	cursor  	int
	loading 	bool
	err 		error

	spin components.SpinnerComponent

	offset 		int
	pageSize 	int
}

type projectsLoadedMsg struct {
	projects []services.Repo
}

type projectsErrMsg struct {
	err error
}

func ProjectsModel(username string) *projectsModel {
	return &projectsModel{
		username: username,
		loading: true,
		spin:  components.NewSpinner(),
		cursor: 0,
		offset: 0,
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

func (m *projectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	newSpin, spinCmd := m.spin.Update(msg)
	m.spin = newSpin
	cmds = append(cmds, spinCmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h := msg.Height
		desired := h - 8
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
		case "ctrl+c", "q":
			return m, tea.Quit
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
			case "pgdown":
		if !m.loading {
			m.offset += m.pageSize
			if m.offset > len(m.projects) - 1 {
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
				m.cursor = len(m.projects)-1
			}
		}
	}

	// if m.loading {
	// 	cmds = append(cmds, m.spin.Model.Tick)
	// }

	return m, tea.Batch(cmds...)
}


func (m *projectsModel) View() string {
	if m.loading {
		return fmt.Sprintf("\n\n %s Loading Github repos...", m.spin.View())
	}

	if m.err != nil {
		return fmt.Sprintf("\n\n Error: %s", m.err)
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

    s := fmt.Sprintf("\n\n Projects of %s (showing %d-%d of %d) — Page %d/%d\n\n",
        m.username, start+1, end, len(m.projects), currentPage, totalPages)

    for i := start; i < end; i++ {
        repo := m.projects[i]
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }
        desc := repo.Description
        if desc == "" {
            desc = "(no description)"
        }
        s += fmt.Sprintf(" %s %s (%d ★)\n    %s\n\n", cursor, repo.Name, repo.Stars, desc)
    }

    s += "\nControls: j/k up/down, pgup/pgdown, home/end, q to quit.\n"
    return s
}
