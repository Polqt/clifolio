package ui

import (
	"strings"
	"time"

	"clifolio/internal/services"
	"clifolio/internal/styles"
	"clifolio/internal/ui/components"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg struct{}
type goToMenuMsg struct{}

type introModel struct {
	fullRunes []rune
	pos       int
	lines     []string
	done      bool
	ascii     []string
	showASCII bool
	theme     styles.Theme
	width     int
	height    int
}

func IntroModel() introModel {
	introData, err := services.LoadASCII("assets/intro.txt")
	fullText := ""
	if err == nil {
		fullText = string(introData)
	} else {
		fullText = "Welcome to my portfolio!\n\nI'm a passionate developer building amazing applications."
	}

	fullText = strings.ReplaceAll(fullText, "\r\n", "\n")

	data, err := services.LoadASCII("assets/ascii.txt")
	var ascii []string
	if err == nil && len(data) > 0 {
		ascii = strings.Split(string(data), "\n")
	}

	theme := styles.NewThemeFromName("default")

	return introModel{
		fullRunes: []rune(fullText),
		lines:     []string{""},
		ascii:     ascii,
		theme:     theme,
	}
}

func (m introModel) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(30*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m introModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tickMsg:
		if m.pos < len(m.fullRunes) {
			r := m.fullRunes[m.pos]
			m.pos++

			if r == '\n' {
				m.lines = append(m.lines, "")
			} else {
				if len(m.lines) == 0 {
					m.lines = append(m.lines, "")
				}
				last := m.lines[len(m.lines)-1]
				last += string(r)
				m.lines[len(m.lines)-1] = last
			}
			return m, tick()
		}

		if !m.done {
			m.done = true
			if len(m.ascii) > 0 {
				m.showASCII = true
			}
			return m, nil
		}

		return m, nil

	case tea.KeyMsg:
		// Allow skipping animation or continuing when done
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// Skip animation or go to menu
		if !m.done {
			// Complete the animation instantly
			m.lines = []string{string(m.fullRunes)}
			m.pos = len(m.fullRunes)
			m.done = true
			if len(m.ascii) > 0 {
				m.showASCII = true
			}
			return m, nil
		}

		// If already done, go to menu
		return m, func() tea.Msg { return goToMenuMsg{} }
	}

	return m, nil
}

func (m introModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sections []string

	// Header
	header := components.HeaderBox("WELCOME", m.theme, m.width-4)
	sections = append(sections, header)

	// Animated intro text
	introContent := m.renderIntroText()
	sections = append(sections, introContent)

	// ASCII art if available
	if m.showASCII && len(m.ascii) > 0 {
		asciiArt := m.renderASCII()
		sections = append(sections, asciiArt)
	}

	// Prompt when done
	if m.done {
		sections = append(sections, components.DividerLine(m.theme, m.width-4, "─"))

		prompt := lipgloss.NewStyle().
			Foreground(m.theme.Accent).
			Bold(true).
			Align(lipgloss.Center).
			Width(m.width).
			Render("✨ Press any key to continue ✨")
		sections = append(sections, prompt)

		keyBindings := []components.KeyBind{
			{Key: "Any Key", Desc: "Enter Portfolio"},
			{Key: "Ctrl+C", Desc: "Exit"},
		}
		footer := components.RenderKeyBindings(keyBindings, m.theme, m.width)
		sections = append(sections, footer)
	}

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (m introModel) renderIntroText() string {
	// Join all lines into a single text
	text := strings.Join(m.lines, "\n")

	textStyle := lipgloss.NewStyle().
		Foreground(m.theme.Primary).
		Width(m.width-16).
		Align(lipgloss.Left).
		Padding(1, 2)

	highlightStyle := lipgloss.NewStyle().
		Foreground(m.theme.Accent).
		Bold(true)

	// Highlight name and key phrases
	text = strings.ReplaceAll(text, "Janpol Hidalgo", highlightStyle.Render("Janpol Hidalgo"))

	styledText := textStyle.Render(text)

	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.theme.Primary).
		Padding(1, 2).
		Width(m.width - 8)

	card := cardStyle.Render(styledText)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		card,
	)
}

func (m introModel) renderASCII() string {
	asciiContent := strings.Join(m.ascii, "\n")

	asciiStyle := lipgloss.NewStyle().
		Foreground(m.theme.Accent).
		Align(lipgloss.Center).
		Width(m.width - 16)

	styledASCII := asciiStyle.Render(asciiContent)

	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Center,
		components.SectionBox("", styledASCII, m.theme, m.width-8),
	)
}
