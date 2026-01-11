package styles

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Background lipgloss.TerminalColor
	Primary    lipgloss.TerminalColor
	Secondary  lipgloss.TerminalColor
	Accent     lipgloss.TerminalColor
	Help       lipgloss.TerminalColor
	Error      lipgloss.TerminalColor
	Title      lipgloss.Style
	Label      lipgloss.Style
}

func NewThemeFromName(name string) Theme {
	switch name {
	case "warrior":
		return Theme{
			Background: lipgloss.Color("#0f0f0f"),
			Primary:    lipgloss.Color("#dc322f"), // Blood red for titles
			Secondary:  lipgloss.Color("#93a1a1"), // Steel gray for secondary text
			Accent:     lipgloss.Color("#2aa198"), // Cyan for stats/skills
			Help:       lipgloss.Color("#586e75"), // Muted gray for help
			Error:      lipgloss.Color("#b58900"), // Yellow for warnings
		}
	case "hacker":
		return Theme{
			Background: lipgloss.Color("#0f0f0f"),
			Primary:    lipgloss.Color("#00ff00"),
			Secondary:  lipgloss.Color("#007700"),
			Accent:     lipgloss.Color("#33ff99"),
			Help:       lipgloss.Color("#626262"),
			Error:      lipgloss.Color("#FF0000"),
		}
	case "dracula":
		return Theme{
			Background: lipgloss.Color("#282a36"),
			Primary:    lipgloss.Color("#ff79c6"),
			Secondary:  lipgloss.Color("#6272a4"),
			Accent:     lipgloss.Color("#8be9fd"),
			Help:       lipgloss.Color("#626262"),
			Error:      lipgloss.Color("#FF0000"),
		}
	case "space":
		return Theme{
            Background: lipgloss.Color("#0a0e27"), 
        	Primary:    lipgloss.Color("#ffffff"),
        	Secondary:  lipgloss.Color("#7c8fb5"), 
        	Accent:     lipgloss.Color("#9d7cff"),
        	Help:       lipgloss.Color("#4a5568"),
        	Error:      lipgloss.Color("#ff6b9d"), 
    	}
	case "digimon": 
		return Theme{
			Background: lipgloss.Color("#0a0e27"),  
			Primary:    lipgloss.Color("#ffffff"),  
			Secondary:  lipgloss.Color("#7c8fb5"),  
			Accent:     lipgloss.Color("#9d7cff"),  
			Help:       lipgloss.Color("#4a5568"),  
			Error:      lipgloss.Color("#ff5c5c"),  

		}
	default:
		// Default is now warrior theme
		return Theme{
			Background: lipgloss.Color("#0f0f0f"),
			Primary:    lipgloss.Color("#dc322f"), 
			Secondary:  lipgloss.Color("#93a1a1"), 
			Accent:     lipgloss.Color("#2aa198"),
			Help:       lipgloss.Color("#586e75"),
			Error:      lipgloss.Color("#b58900"), 
		}
	}
}

func (t Theme) BuildStyles() (title lipgloss.Style, label lipgloss.Style) {
	title = lipgloss.NewStyle().Foreground(t.Primary)
	label = lipgloss.NewStyle().Foreground(t.Secondary)
	return
}
