package components

import (
	"clifolio/internal/styles"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type KeyBind struct {
	Key  string
	Desc string
}

func RenderKeyBindings(bindings []KeyBind, theme styles.Theme, width int) string {
	keyStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true).
		Background(lipgloss.Color("#1a1a1a")).
		Padding(0, 1)

	descStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary)

	var hints []string
	for _, binding := range bindings {
		hint := keyStyle.Render(binding.Key) + " " + descStyle.Render(binding.Desc)
		hints = append(hints, hint)
	}

	content := strings.Join(hints, "  •  ")

	// Add warrior-themed decorative border
	borderStyle := lipgloss.NewStyle().
		Foreground(theme.Primary)

	// Create a full-width line with centered COMMANDS text
	commandLabel := " ⚔️  COMMANDS ⚔️  "
	labelWidth := runewidth.StringWidth(commandLabel) // Calculate actual visual width
	remainingWidth := width - labelWidth
	if remainingWidth < 0 {
		remainingWidth = 0
	}
	leftPad := remainingWidth / 2
	rightPad := remainingWidth - leftPad

	topLine := strings.Repeat("━", leftPad) + commandLabel + strings.Repeat("━", rightPad)

	footerContent := lipgloss.JoinVertical(
		lipgloss.Left,
		borderStyle.Render(topLine),
		lipgloss.NewStyle().
			Width(width).
			Align(lipgloss.Center).
			Foreground(theme.Secondary).
			Padding(1, 0).
			Render(content),
	)

	return footerContent
}

func RenderHelpMenu(sections map[string][]KeyBind, theme styles.Theme) string {
	var output strings.Builder

	sectionStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true).
		MarginTop(1).
		MarginBottom(1)

	keyStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Width(15).
		Bold(true)

	descStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary)

	for section, bindings := range sections {
		output.WriteString(sectionStyle.Render("▸ "+section) + "\n")

		for _, binding := range bindings {
			line := lipgloss.JoinHorizontal(
				lipgloss.Top,
				keyStyle.Render(binding.Key),
				descStyle.Render(binding.Desc),
			)
			output.WriteString("  " + line + "\n")
		}
	}

	return output.String()
}

func GetNavigationBindings(km Keymap) []KeyBind {
	return []KeyBind{
		{Key: km.Up + "/" + km.Down, Desc: "Navigate"},
		{Key: km.Confirm, Desc: "Select"},
		{Key: km.Back, Desc: "Back"},
		{Key: km.Quit, Desc: "Quit"},
	}
}
