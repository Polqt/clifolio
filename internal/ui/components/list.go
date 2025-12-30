package components

import (
	"clifolio/internal/styles"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type ListItem struct {
	Title   string
	Content string
	Icon    string
}

func RenderList(items []ListItem, cursor int, theme styles.Theme) string {
	selectedStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)
	
	normalStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary)
	
	cursorStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	var output string
	for i, item := range items {
		icon := item.Icon
		if icon == "" {
            icon = "•"
        }

		line := fmt.Sprintf("%s %s", icon, item.Title)
		if item.Content != "" {
			line += "\n  " + lipgloss.NewStyle().
				Foreground(theme.Secondary).
				Italic(true).
				Render(item.Content)
		}

		if i == cursor {
			output += cursorStyle.Render("▸ ") + selectedStyle.Render(line) + "\n"
		} else {
			output += "  " + normalStyle.Render(line) + "\n"
		}
	}
	return output
}