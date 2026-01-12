package components

import (
	"clifolio/internal/styles"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

type BorderStyle int

const (
	BorderRounded BorderStyle = iota
	BorderDouble
	BorderThick
	BorderNormal
	BorderHidden
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

func TitleBox(theme styles.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(theme.Primary).
		Padding(0, 2).
		Bold(true)
}

func SubtleBorder(theme styles.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(theme.Secondary).
		Padding(1, 2)
}

func GlowBorder(theme styles.Theme) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(theme.Accent).
		Padding(1, 3).
		Bold(true)
}

func SectionBox(title, content string, theme styles.Theme, width int) string {
	// Scroll-like box with decorative borders
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.Border{
			Top:         "═",
			Bottom:      "═",
			Left:        "║",
			Right:       "║",
			TopLeft:     "╭",
			TopRight:    "╮",
			BottomLeft:  "╰",
			BottomRight: "╯",
		}).
		BorderForeground(theme.Primary).
		Padding(1, 2).
		Width(width)

	return boxStyle.Render(content)
}

func CardBox(content string, theme styles.Theme, selected bool) string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		MarginBottom(1)

	if selected {
		style = style.
			BorderForeground(theme.Accent).
			BorderStyle(lipgloss.ThickBorder()).
			Bold(true)
	} else {
		style = style.BorderForeground(theme.Secondary)
	}

	return style.Render(content)
}

func InfoPanel(label, value string, theme styles.Theme) string {
	labelStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Bold(true).
		Width(15).
		Align(lipgloss.Right)

	valueStyle := lipgloss.NewStyle().
		Foreground(theme.Primary)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		labelStyle.Render(label+":"),
		"  ",
		valueStyle.Render(value),
	)
}

func HeaderBox(title string, theme styles.Theme, width int) string {
	borderStyle := lipgloss.NewStyle().
		Foreground(theme.Primary)

	titleStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	// Scroll-like decorative top border - full width
	topBorder := borderStyle.Render("┌" + strings.Repeat("═", width-2) + "┐")

	// Title line with scroll decorations
	swordDecor := "⚔️"
	titleText := "  " + title + " "

	// Calculate visual width using runewidth to account for emojis
	swordWidth := runewidth.StringWidth(swordDecor)
	titleTextWidth := runewidth.StringWidth(titleText)
	titleVisualWidth := swordWidth + titleTextWidth + swordWidth
	contentWidth := width - 2 // Subtract borders

	// Calculate padding to fill the entire width
	padding := contentWidth - titleVisualWidth
	if padding < 0 {
		padding = 0
	}
	leftPad := padding / 2
	rightPad := padding - leftPad

	// Build the title line with symmetric padding
	titleLine := strings.Repeat(" ", leftPad) + swordDecor + titleText + swordDecor + strings.Repeat(" ", rightPad)

	titleLineWidth := runewidth.StringWidth(titleLine)
	if titleLineWidth < contentWidth {
		extraSpace := contentWidth - titleLineWidth
		extraLeft := extraSpace / 2
		extraRight := extraSpace - extraLeft
		titleLine = strings.Repeat(" ", extraLeft) + titleLine + strings.Repeat(" ", extraRight)
	}

	middleLine := borderStyle.Render("│") + titleStyle.Render(titleLine) + borderStyle.Render("│")

	// Scroll-like decorative bottom border - full width
	bottomBorder := borderStyle.Render("└" + strings.Repeat("═", width-2) + "┘")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		topBorder,
		middleLine,
		bottomBorder,
	)
}

func FooterBox(content string, theme styles.Theme, width int) string {
	style := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Background(theme.Background).
		Padding(0, 2).
		Width(width).
		Align(lipgloss.Center)

	return style.Render(content)
}

func DividerLine(theme styles.Theme, width int, char string) string {
	if char == "" {
		char = "─"
	}

	style := lipgloss.NewStyle().
		Foreground(theme.Secondary)

	swordDecor := " ⚔️   "
	swordVisualWidth := runewidth.StringWidth(swordDecor)

	// Ensure we have enough width for the sword decoration
	if width < swordVisualWidth {
		return style.Render(strings.Repeat(char, width))
	}

	remainingWidth := width - swordVisualWidth
	sideLength := remainingWidth / 2
	rightSideLength := remainingWidth - sideLength // Handle odd widths

	leftSide := strings.Repeat(char, sideLength)
	rightSide := strings.Repeat(char, rightSideLength)

	line := leftSide + swordDecor + rightSide

	return style.Render(line)
}

// WarriorStatusBar creates an HP/MP style bar for warrior theme
func WarriorStatusBar(label string, current, max int, theme styles.Theme, width int) string {
	barWidth := width - len(label) - 15 // Space for label and stats
	if barWidth < 10 {
		barWidth = 10
	}

	percentage := float64(current) / float64(max)
	filled := int(percentage * float64(barWidth))
	if filled > barWidth {
		filled = barWidth
	}

	filledBar := strings.Repeat("█", filled)
	emptyBar := strings.Repeat("░", barWidth-filled)

	barStyle := lipgloss.NewStyle().Foreground(theme.Accent)
	emptyStyle := lipgloss.NewStyle().Foreground(theme.Secondary)
	labelStyle := lipgloss.NewStyle().Foreground(theme.Primary).Bold(true)
	statsStyle := lipgloss.NewStyle().Foreground(theme.Secondary)

	return labelStyle.Render(label) + " [" +
		barStyle.Render(filledBar) + emptyStyle.Render(emptyBar) +
		"] " + statsStyle.Render(lipgloss.NewStyle().Render(string(rune(current)))+"/"+string(rune(max)))
}

// WarriorBox creates a warrior-themed decorated box
func WarriorBox(content string, theme styles.Theme, width int) string {
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.Border{
			Top:         "═",
			Bottom:      "═",
			Left:        "║",
			Right:       "║",
			TopLeft:     "╔",
			TopRight:    "╗",
			BottomLeft:  "╚",
			BottomRight: "╝",
		}).
		BorderForeground(theme.Primary).
		Padding(1, 2).
		Width(width)

	return borderStyle.Render(content)
}

func PixelDecoration(theme styles.Theme) string {
	decorStyle := lipgloss.NewStyle().Foreground(theme.Primary)
	return decorStyle.Render("▓▒░")
}
