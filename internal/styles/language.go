package styles

import "github.com/charmbracelet/lipgloss"

var LanguageColors = map[string]lipgloss.Color{
	"Go":         lipgloss.Color("#00ADD8"),
    "Python":     lipgloss.Color("#3572A5"),
    "JavaScript": lipgloss.Color("#F7DF1E"),
    "TypeScript": lipgloss.Color("#3178C6"),
    "Rust":       lipgloss.Color("#DEA584"),
    "Ruby":       lipgloss.Color("#CC342D"),
    "Java":       lipgloss.Color("#B07219"),
    "C++":        lipgloss.Color("#F34B7D"),
    "C":          lipgloss.Color("#555555"),
    "C#":         lipgloss.Color("#239120"),
    "PHP":        lipgloss.Color("#4F5D95"),
    "Swift":      lipgloss.Color("#FA7343"),
    "Kotlin":     lipgloss.Color("#A97BFF"),
    "Dart":       lipgloss.Color("#00B4AB"),
    "HTML":       lipgloss.Color("#E34C26"),
    "CSS":        lipgloss.Color("#1572B6"),
    "Shell":      lipgloss.Color("#89E051"),
    "Lua":        lipgloss.Color("#000080"),
    "Vue":        lipgloss.Color("#41B883"),
    "Svelte":     lipgloss.Color("#FF3E00"),
}

func GetLanguageColor(language string) lipgloss.Color {
	if color, ok := LanguageColors[language]; ok {
		return color
	}
	return lipgloss.Color("#858585") 
}

func GetLanguageStyle(language string) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(GetLanguageColor(language))
}