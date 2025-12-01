package styles

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Background lipgloss.TerminalColor
	Primary    lipgloss.TerminalColor
	Secondary  lipgloss.TerminalColor
	Accent     lipgloss.TerminalColor
	Title      lipgloss.Style
	Label      lipgloss.Style
}

func NewThemeFromName(name string) Theme {
	switch name {
	case "hacker":
		return Theme{
			Background: lipgloss.Color("#0f0f0f"),
			Primary: lipgloss.Color("#00ff00"),
			Secondary: lipgloss.Color("#007700"),
			Accent: lipgloss.Color("#33ff99"),
		}
	case "dracula":
		return Theme{
			Background: lipgloss.Color("#282a36"),
			Primary: lipgloss.Color("#ff79c6"),
			Secondary: lipgloss.Color("#6272a4"),
			Accent: lipgloss.Color("#8be9fd"),
		}
	default:
		return Theme{
			Background: lipgloss.Color("#002b36"),
			Primary: lipgloss.Color("#b58900"),
			Secondary: lipgloss.Color("#586e75"),
			Accent: lipgloss.Color("#2aa198"),
		}
	}
}

func (t Theme) BuildStyles() (title lipgloss.Style, label lipgloss.Style) {
	title = lipgloss.NewStyle().Foreground(t.Primary)
	label = lipgloss.NewStyle().Foreground(t.Secondary)
	return
}