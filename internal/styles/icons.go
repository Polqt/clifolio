package styles

import "clifolio/internal/ui/state"
//mankykykyto
func getScreenIcon(s state.Screen) string {
	switch s {
	case state.ScreenProjects:
		return "📁"
	case state.ScreenSkills:
		return "🛠️"
	case state.ScreenExperience:
		return "💼"
	case state.ScreenContact:
		return "📧"
	case state.ScreenTheme:
		return "🎨"
	default:
		return "•"
	}
}
