package styles

import "github.com/charmbracelet/lipgloss"

const (
	PaddingSmall  = 1
	PaddingMedium = 2
	PaddingLarge  = 4
	MarginSmall   = 1
	MarginMedium  = 2
)

type TextStyles struct {
	Title 		lipgloss.Style
	Subtitle	lipgloss.Style	
	Body 		lipgloss.Style
	Code 		lipgloss.Style
	Highlighted	lipgloss.Style
	Dimmed 		lipgloss.Style
}

func NewStyles(theme Theme) TextStyles {
	return TextStyles{
		Title: lipgloss.NewStyle().
			Foreground(theme.Primary).
			Bold(true).
			MarginBottom(1),
		
		Subtitle: lipgloss.NewStyle().
			Foreground(theme.Secondary).
			Italic(true),

		Body: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")),
		
		Code: lipgloss.NewStyle().
			Foreground(theme.Accent).
			Background(lipgloss.Color("#1a1a1a")).
			Padding(0, 1),
		
		Highlighted: lipgloss.NewStyle().
            Foreground(theme.Accent).
            Bold(true),
        
        Dimmed: lipgloss.NewStyle().
            Foreground(theme.Secondary).
            Faint(true),
	}
}