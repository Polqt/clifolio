package components

import (
	"clifolio/internal/styles"

	"github.com/charmbracelet/lipgloss"
)

func StandardBorder(theme styles.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Primary).
		Padding(1, 2)
}

func AccentBorder(theme styles.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(theme.Accent).
		Padding(1, 2)
}

func TitleBox(title string, theme styles.Theme) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(theme.Primary).
		Padding(0, 2).
		Bold(true).
		Align(lipgloss.Center)

	return style.Render(title)
}

func InfoBox(info string, theme styles.Theme) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Secondary).
		Padding(1, 2)

	return style.Render(info)
}
