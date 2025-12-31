package components

import (
	"clifolio/internal/styles"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type ListItem struct {
	Title   string
	Content string
	Icon    string
	Badge 	string
	Meta 	string
}
type ListStyle struct {
	ShowNumbers 	bool
	ShowIcons 		bool
	ShowBadges 		bool
	CompactMode 	bool
	HighlightColor	lipgloss.Color
}

func RenderList(items []ListItem, cursor int, theme styles.Theme, style ListStyle) string {
	if len(items) == 0 {
		return lipgloss.NewStyle().
			Foreground(theme.Secondary).
			Italic(true).
			Render("No items to display")
	}

	selectedStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)
	
	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff"))

	descStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Italic(true)

	badgeStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Background(lipgloss.Color("#1a1a1a")).
		Padding(0, 1).
		Bold(true)
	
	metaStyle := lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Faint(true)
		
	cursorIcon := lipgloss.NewStyle().
        Foreground(theme.Accent).
        Bold(true).
        Render("▸")
	

	var output strings.Builder

	for i, item := range items {
		isSelected := i == cursor

		var line strings.Builder
		
		if isSelected {
			line.WriteString(cursorIcon + " ")
		} else {
			if style.ShowNumbers {
				line.WriteString(fmt.Sprintf("%2d", i+1))
			} else {
				line.WriteString("   ")
			}
		}

		if style.ShowIcons && item.Icon != "" {
			icon := item.Icon
			if isSelected {
				icon = selectedStyle.Render(icon)
			} else {
				icon = normalStyle.Render(icon)
			}
			line.WriteString(icon + " ")
		}

		title := item.Title
		if isSelected {
			title = selectedStyle.Render(title)
		} else {
			title = normalStyle.Render(title)
		}
		line.WriteString(title)
		
		if style.ShowBadges && item.Badge != "" {
			line.WriteString(" " + badgeStyle.Render(item.Badge))
		}

		if item.Meta != "" {
			line.WriteString(" " + metaStyle.Render("("+item.Meta+")"))
		}

		output.WriteString(line.String() + "\n")

		if !style.CompactMode && item.Content != "" {
			desc := "     " + descStyle.Render(item.Content)
			output.WriteString(desc + "\n")
		}

		if !style.CompactMode && i < len(items)-1 {
			output.WriteString("\n")
		}
	}

	return output.String()
}

func RenderCardList(items []ListItem, cursor int, theme styles.Theme, width int) string {
	var cards []string
	for i, item := range items {
		isSelected := i == cursor

		var content strings.Builder

		titleStyle := lipgloss.NewStyle().
			Foreground(theme.Primary).
			Bold(true)

		if item.Icon != "" {
			content.WriteString(item.Icon + " ")
		}
		content.WriteString(titleStyle.Render(item.Title) + "\n\n")

		if item.Content != "" {
			descStyle := lipgloss.NewStyle().
				Foreground(theme.Secondary).
				Width(width - 8)
			content.WriteString(descStyle.Render(item.Content) + "\n")
		}

		if item.Badge != "" || item.Meta != "" {
			content.WriteString("\n")
			if item.Badge != "" {
				badgeStyle := lipgloss.NewStyle().
					Foreground(theme.Accent).
					Background(lipgloss.Color("#1a1a1a")).
					Padding(0, 1)
				content.WriteString(badgeStyle.Render(item.Badge) + " ")
			}
			if item.Meta != "" {
				metaStyle := lipgloss.NewStyle().
					Foreground(theme.Secondary).
					Faint(true)
				content.WriteString(metaStyle.Render(item.Meta))
			}
		}

		card := CardBox(content.String(), theme, isSelected)
		cards = append(cards, card)
	}

	return lipgloss.JoinVertical(lipgloss.Left, cards...)
}

func RenderGridList(items []ListItem, cursor int, theme styles.Theme, width int) string {
	cardWidth := (width - 4) / 2

	var rows []string 
	var currentRow []string

	for i, item := range items {
		isSelected := i == cursor

		titleStyle := lipgloss.NewStyle().
			Foreground(theme.Primary).
			Bold(true)
		
		content := titleStyle.Render(item.Icon + " " + item.Title) + "\n"
		if item.Content != "" {
			descStyle := lipgloss.NewStyle().
				Foreground(theme.Secondary).
				Width(cardWidth - 4)
			content += descStyle.Render(item.Content)
		}

		cardStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1).
			Width(cardWidth)

		if isSelected {
			cardStyle = cardStyle.
				BorderForeground(theme.Accent).
				BorderStyle(lipgloss.ThickBorder())
		} else {
			cardStyle = cardStyle.BorderForeground(theme.Secondary)
		}

		card := cardStyle.Render(content)
		currentRow = append(currentRow, card)

		if len(currentRow) == 2 {
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, currentRow...))
			currentRow = []string{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, currentRow...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func RenderTableList(headers []string, rows [][]string, cursor int, theme styles.Theme) string {
    headerStyle := lipgloss.NewStyle().
        Foreground(theme.Primary).
        Bold(true).
        BorderBottom(true).
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(theme.Secondary)

    selectedRowStyle := lipgloss.NewStyle().
        Foreground(theme.Accent).
        Bold(true).
        Background(lipgloss.Color("#1a1a1a"))

    normalRowStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#ffffff"))

    // Calculate column widths
    colWidths := make([]int, len(headers))
    for i, h := range headers {
        colWidths[i] = len(h)
    }
    for _, row := range rows {
        for i, cell := range row {
            if len(cell) > colWidths[i] {
                colWidths[i] = len(cell)
            }
        }
    }

    // Render header
    var headerCells []string
    for i, h := range headers {
        cell := headerStyle.Width(colWidths[i]).Render(h)
        headerCells = append(headerCells, cell)
    }
    output := lipgloss.JoinHorizontal(lipgloss.Top, headerCells...) + "\n"

    // Render rows
    for i, row := range rows {
        var cells []string
        style := normalRowStyle
        if i == cursor {
            style = selectedRowStyle
        }

        for j, cell := range row {
            cells = append(cells, style.Width(colWidths[j]).Render(cell))
        }
        output += lipgloss.JoinHorizontal(lipgloss.Top, cells...) + "\n"
    }

    return output
}