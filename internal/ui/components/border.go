package components

import (
	"clifolio/internal/styles"

	"github.com/charmbracelet/lipgloss"
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
	titleStyle := lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true).
		Padding(0, 1).
		Background(theme.Background)
	
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(theme.Primary).
		Padding(1, 2).
		Width(width)
	
	titleBar := titleStyle.Render("┤ " + title + " ├")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleBar,
		boxStyle.Render(content),
	)
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
	style := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true).
		Align(lipgloss.Center).
		Width(width - 4)
	
	borderStyle := lipgloss.NewStyle().
		Foreground(theme.Primary)
	
	topBorder := borderStyle.Render("╔" + lipgloss.PlaceHorizontal(width-2, lipgloss.Center, "═") + "╗")
	bottomBorder := borderStyle.Render("╚" + lipgloss.PlaceHorizontal(width-2, lipgloss.Center, "═") + "╝")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		topBorder,
		"║ " + style.Render(title) +" ║",
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
		Foreground(theme.Secondary).
		Width(width)
	
	line := ""
	for i := 0; i < width; i++ {
		line += char
	}

	return style.Render(line)
}

func GradientBorder(theme styles.Theme, topColor, bottomColor lipgloss.Color) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(topColor).
		Padding(1, 2)
}
